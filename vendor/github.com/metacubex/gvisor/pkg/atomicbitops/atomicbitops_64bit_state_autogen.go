// automatically generated by stateify.

//go:build !arm && !mips && !mipsle && !386 && !arm && !mips && !mipsle && !386
// +build !arm,!mips,!mipsle,!386,!arm,!mips,!mipsle,!386

package atomicbitops

import (
	"context"

	"github.com/metacubex/gvisor/pkg/state"
)

func (i *Int32) StateTypeName() string {
	return "pkg/atomicbitops.Int32"
}

func (i *Int32) StateFields() []string {
	return []string{
		"value",
	}
}

func (i *Int32) beforeSave() {}

// +checklocksignore
func (i *Int32) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.value)
}

func (i *Int32) afterLoad(context.Context) {}

// +checklocksignore
func (i *Int32) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.value)
}

func (u *Uint32) StateTypeName() string {
	return "pkg/atomicbitops.Uint32"
}

func (u *Uint32) StateFields() []string {
	return []string{
		"value",
	}
}

func (u *Uint32) beforeSave() {}

// +checklocksignore
func (u *Uint32) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.value)
}

func (u *Uint32) afterLoad(context.Context) {}

// +checklocksignore
func (u *Uint32) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.value)
}

func (b *Bool) StateTypeName() string {
	return "pkg/atomicbitops.Bool"
}

func (b *Bool) StateFields() []string {
	return []string{
		"Uint32",
	}
}

func (b *Bool) beforeSave() {}

// +checklocksignore
func (b *Bool) StateSave(stateSinkObject state.Sink) {
	b.beforeSave()
	stateSinkObject.Save(0, &b.Uint32)
}

func (b *Bool) afterLoad(context.Context) {}

// +checklocksignore
func (b *Bool) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &b.Uint32)
}

func (i *Int64) StateTypeName() string {
	return "pkg/atomicbitops.Int64"
}

func (i *Int64) StateFields() []string {
	return []string{
		"value",
	}
}

func (i *Int64) beforeSave() {}

// +checklocksignore
func (i *Int64) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.value)
}

func (i *Int64) afterLoad(context.Context) {}

// +checklocksignore
func (i *Int64) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.value)
}

func (u *Uint64) StateTypeName() string {
	return "pkg/atomicbitops.Uint64"
}

func (u *Uint64) StateFields() []string {
	return []string{
		"value",
	}
}

func (u *Uint64) beforeSave() {}

// +checklocksignore
func (u *Uint64) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.value)
}

func (u *Uint64) afterLoad(context.Context) {}

// +checklocksignore
func (u *Uint64) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.value)
}

func init() {
	state.Register((*Int32)(nil))
	state.Register((*Uint32)(nil))
	state.Register((*Bool)(nil))
	state.Register((*Int64)(nil))
	state.Register((*Uint64)(nil))
}
