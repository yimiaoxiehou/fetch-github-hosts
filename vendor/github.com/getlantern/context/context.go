// Package context provides a mechanism for transparently tracking contextual
// state associated to the current goroutine and even across goroutines.
package context

import (
	"sync"
)

// Manager provides the ability to create and access Contexts.
type Manager interface {
	// Enter enters a new level on the current Context stack, creating a new Context
	// if necessary.
	Enter() Context

	// Go starts the given function on a new goroutine but sharing the context of
	// the current goroutine (if it has one).
	Go(func())

	// PutGlobal puts the given key->value pair into the global context.
	PutGlobal(key string, value interface{})

	// PutGlobalDynamic puts a key->value pair into the global context where the
	// value is generated by a function that gets evaluated at every Read. If the
	// value is a map[string]interface{}, we will unpack the map and set each
	// contained key->value pair independently.
	PutGlobalDynamic(key string, valueFN func() interface{})

	// AsMap returns a map containing all values from the supplied obj if it is a
	// Contextual, plus any addition values from along the stack, plus globals if so
	// specified.
	AsMap(obj interface{}, includeGlobals bool) Map
}

type manager struct {
	contexts   map[uint64]*context
	mxContexts sync.RWMutex
	global     Map
	mxGlobal   sync.RWMutex
}

// NewManager creates a new Manager
func NewManager() Manager {
	return &manager{
		contexts: make(map[uint64]*context),
		global:   make(Map),
	}
}

// Contextual is an interface for anything that maintains its own context.
type Contextual interface {
	// Fill fills the given Map with all of this Contextual's context
	Fill(m Map)
}

// Map is a map of key->value pairs.
type Map map[string]interface{}

// Fill implements the method from the Contextual interface.
func (_m Map) Fill(m Map) {
	for key, value := range _m {
		m[key] = value
	}
}

// Context is a context containing key->value pairs
type Context interface {
	// Enter enters a new level on this Context stack.
	Enter() Context

	// Go starts the given function on a new goroutine.
	Go(fn func())

	// Exit exits the current level on this Context stack.
	Exit()

	// Put puts a key->value pair into the current level of the context stack.
	Put(key string, value interface{}) Context

	// PutIfAbsent puts the given key->value pair into the current level of the
	// context stack if and only if that key is defined nowhere within the context
	// stack (including parent contexts).
	PutIfAbsent(key string, value interface{}) Context

	// PutDynamic puts a key->value pair into the current level of the context
	// stack where the value is generated by a function that gets evaluated at
	// every Read. If the value is a map[string]interface{}, we will unpack the
	// map and set each contained key->value pair independently.
	PutDynamic(key string, valueFN func() interface{}) Context

	// Fill fills the given map with data from this Context
	Fill(m Map)

	// AsMap returns a map containing all values from the supplied obj if it is a
	// Contextual, plus any addition values from along the stack, plus globals if
	// so specified.
	AsMap(obj interface{}, includeGlobals bool) Map
}

type context struct {
	cm           *manager
	id           uint64
	parent       *context
	branchedFrom *context
	data         Map
	mx           sync.RWMutex
}

type dynval struct {
	fn func() interface{}
}

func (cm *manager) Enter() Context {
	return cm.enter(curGoroutineID())
}

func (cm *manager) enter(id uint64) *context {
	cm.mxContexts.Lock()
	parentOrNil := cm.contexts[id]
	c := cm.makeContext(id, parentOrNil, nil)
	cm.contexts[id] = c
	cm.mxContexts.Unlock()
	return c
}

func (cm *manager) exit(id uint64, parent *context) {
	cm.mxContexts.Lock()
	if parent == nil {
		delete(cm.contexts, id)
	} else {
		cm.contexts[id] = parent
	}
	cm.mxContexts.Unlock()
}

func (cm *manager) branch(id uint64, from *context) {
	next := cm.makeContext(id, nil, from)
	cm.mxContexts.Lock()
	cm.contexts[id] = next
	cm.mxContexts.Unlock()
}

func (cm *manager) merge(id uint64) {
	cm.mxContexts.Lock()
	delete(cm.contexts, id)
	cm.mxContexts.Unlock()
}

func (c *context) Enter() Context {
	c.mx.RLock()
	id := c.id
	c.mx.RUnlock()
	return c.cm.enter(id)
}

func (c *context) Go(fn func()) {
	go func() {
		id := curGoroutineID()
		c.cm.branch(id, c)
		fn()
		c.cm.merge(id)
	}()
}

func (cm *manager) Go(fn func()) {
	c := cm.currentContext()
	if c != nil {
		c.Go(fn)
	} else {
		go fn()
	}
}

func (cm *manager) makeContext(id uint64, parent *context, branchedFrom *context) *context {
	return &context{
		cm:           cm,
		id:           id,
		parent:       parent,
		branchedFrom: branchedFrom,
		data:         make(Map),
	}
}

func (c *context) Exit() {
	c.mx.RLock()
	id := c.id
	parent := c.parent
	c.mx.RUnlock()
	c.cm.exit(id, parent)
}

func (c *context) Put(key string, value interface{}) Context {
	c.mx.Lock()
	c.data[key] = value
	c.mx.Unlock()
	return c
}

func (c *context) PutIfAbsent(key string, value interface{}) Context {
	for ctx := c; ctx != nil; {
		ctx.mx.RLock()
		_, exists := ctx.data[key]
		next := ctx.parent
		if next == nil {
			next = ctx.branchedFrom
		}
		ctx.mx.RUnlock()
		if exists {
			return c
		}
		ctx = next
	}

	// Value not set, set it
	return c.Put(key, value)
}

func (c *context) PutDynamic(key string, valueFN func() interface{}) Context {
	value := &dynval{valueFN}
	c.mx.Lock()
	c.data[key] = value
	c.mx.Unlock()
	return c
}

func (cm *manager) PutGlobal(key string, value interface{}) {
	cm.mxGlobal.Lock()
	cm.global[key] = value
	cm.mxGlobal.Unlock()
}

func (cm *manager) PutGlobalDynamic(key string, valueFN func() interface{}) {
	value := &dynval{valueFN}
	cm.mxGlobal.Lock()
	cm.global[key] = value
	cm.mxGlobal.Unlock()
}

func (c *context) Fill(m Map) {
	for ctx := c; ctx != nil; {
		ctx.mx.RLock()
		fill(m, ctx.data)
		next := ctx.parent
		if next == nil {
			next = ctx.branchedFrom
		}
		ctx.mx.RUnlock()
		ctx = next
	}
}

func (cm *manager) AsMap(obj interface{}, includeGlobals bool) Map {
	return cm.currentContext().asMap(cm, obj, includeGlobals)
}

func (c *context) AsMap(obj interface{}, includeGlobals bool) Map {
	return c.asMap(c.cm, obj, includeGlobals)
}

func (c *context) asMap(cm *manager, obj interface{}, includeGlobals bool) Map {
	result := make(Map, 0)
	cl, ok := obj.(Contextual)
	if ok {
		cl.Fill(result)
	}
	if c != nil {
		c.Fill(result)
	}
	if includeGlobals {
		cm.mxGlobal.RLock()
		fill(result, cm.global)
		cm.mxGlobal.RUnlock()
	}
	return result
}

func fill(m Map, from Map) {
	if m != nil {
		doFill := func(key string, _value interface{}) {
			switch value := _value.(type) {
			case map[string]interface{}:
				for k, v := range value {
					m[k] = v
				}
			default:
				m[key] = value
			}
		}

		for key, value := range from {
			_, alreadyRead := m[key]
			if !alreadyRead {
				switch v := value.(type) {
				case *dynval:
					doFill(key, v.fn())
				default:
					doFill(key, v)
				}
			}
		}
	}
}

func (cm *manager) currentContext() *context {
	id := curGoroutineID()
	cm.mxContexts.RLock()
	c := cm.contexts[id]
	cm.mxContexts.RUnlock()
	return c
}
