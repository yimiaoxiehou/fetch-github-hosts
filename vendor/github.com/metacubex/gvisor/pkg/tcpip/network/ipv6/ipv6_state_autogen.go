// automatically generated by stateify.

package ipv6

import (
	"context"

	"github.com/metacubex/gvisor/pkg/state"
)

func (i *icmpv6DestinationUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.icmpv6DestinationUnreachableSockError"
}

func (i *icmpv6DestinationUnreachableSockError) StateFields() []string {
	return []string{}
}

func (i *icmpv6DestinationUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv6DestinationUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *icmpv6DestinationUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv6DestinationUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func (i *icmpv6DestinationNetworkUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.icmpv6DestinationNetworkUnreachableSockError"
}

func (i *icmpv6DestinationNetworkUnreachableSockError) StateFields() []string {
	return []string{
		"icmpv6DestinationUnreachableSockError",
	}
}

func (i *icmpv6DestinationNetworkUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv6DestinationNetworkUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv6DestinationUnreachableSockError)
}

func (i *icmpv6DestinationNetworkUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv6DestinationNetworkUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv6DestinationUnreachableSockError)
}

func (i *icmpv6DestinationPortUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.icmpv6DestinationPortUnreachableSockError"
}

func (i *icmpv6DestinationPortUnreachableSockError) StateFields() []string {
	return []string{
		"icmpv6DestinationUnreachableSockError",
	}
}

func (i *icmpv6DestinationPortUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv6DestinationPortUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv6DestinationUnreachableSockError)
}

func (i *icmpv6DestinationPortUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv6DestinationPortUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv6DestinationUnreachableSockError)
}

func (i *icmpv6DestinationAddressUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.icmpv6DestinationAddressUnreachableSockError"
}

func (i *icmpv6DestinationAddressUnreachableSockError) StateFields() []string {
	return []string{
		"icmpv6DestinationUnreachableSockError",
	}
}

func (i *icmpv6DestinationAddressUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv6DestinationAddressUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv6DestinationUnreachableSockError)
}

func (i *icmpv6DestinationAddressUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv6DestinationAddressUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv6DestinationUnreachableSockError)
}

func (e *icmpv6PacketTooBigSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.icmpv6PacketTooBigSockError"
}

func (e *icmpv6PacketTooBigSockError) StateFields() []string {
	return []string{
		"mtu",
	}
}

func (e *icmpv6PacketTooBigSockError) beforeSave() {}

// +checklocksignore
func (e *icmpv6PacketTooBigSockError) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.mtu)
}

func (e *icmpv6PacketTooBigSockError) afterLoad(context.Context) {}

// +checklocksignore
func (e *icmpv6PacketTooBigSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.mtu)
}

func (e *endpointMu) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.endpointMu"
}

func (e *endpointMu) StateFields() []string {
	return []string{
		"addressableEndpointState",
		"ndp",
		"mld",
	}
}

func (e *endpointMu) beforeSave() {}

// +checklocksignore
func (e *endpointMu) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.addressableEndpointState)
	stateSinkObject.Save(1, &e.ndp)
	stateSinkObject.Save(2, &e.mld)
}

func (e *endpointMu) afterLoad(context.Context) {}

// +checklocksignore
func (e *endpointMu) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.addressableEndpointState)
	stateSourceObject.Load(1, &e.ndp)
	stateSourceObject.Load(2, &e.mld)
}

func (d *dadMu) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.dadMu"
}

func (d *dadMu) StateFields() []string {
	return []string{
		"dad",
	}
}

func (d *dadMu) beforeSave() {}

// +checklocksignore
func (d *dadMu) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.dad)
}

func (d *dadMu) afterLoad(context.Context) {}

// +checklocksignore
func (d *dadMu) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.dad)
}

func (e *endpointDAD) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.endpointDAD"
}

func (e *endpointDAD) StateFields() []string {
	return []string{
		"mu",
	}
}

func (e *endpointDAD) beforeSave() {}

// +checklocksignore
func (e *endpointDAD) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.mu)
}

func (e *endpointDAD) afterLoad(context.Context) {}

// +checklocksignore
func (e *endpointDAD) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.mu)
}

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"nic",
		"dispatcher",
		"protocol",
		"stats",
		"enabled",
		"forwarding",
		"multicastForwarding",
		"mu",
		"dad",
	}
}

func (e *endpoint) beforeSave() {}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.nic)
	stateSinkObject.Save(1, &e.dispatcher)
	stateSinkObject.Save(2, &e.protocol)
	stateSinkObject.Save(3, &e.stats)
	stateSinkObject.Save(4, &e.enabled)
	stateSinkObject.Save(5, &e.forwarding)
	stateSinkObject.Save(6, &e.multicastForwarding)
	stateSinkObject.Save(7, &e.mu)
	stateSinkObject.Save(8, &e.dad)
}

func (e *endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.nic)
	stateSourceObject.Load(1, &e.dispatcher)
	stateSourceObject.Load(2, &e.protocol)
	stateSourceObject.Load(3, &e.stats)
	stateSourceObject.Load(4, &e.enabled)
	stateSourceObject.Load(5, &e.forwarding)
	stateSourceObject.Load(6, &e.multicastForwarding)
	stateSourceObject.Load(7, &e.mu)
	stateSourceObject.Load(8, &e.dad)
}

func (o *OpaqueInterfaceIdentifierOptions) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.OpaqueInterfaceIdentifierOptions"
}

func (o *OpaqueInterfaceIdentifierOptions) StateFields() []string {
	return []string{
		"SecretKey",
	}
}

func (o *OpaqueInterfaceIdentifierOptions) beforeSave() {}

// +checklocksignore
func (o *OpaqueInterfaceIdentifierOptions) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.SecretKey)
}

func (o *OpaqueInterfaceIdentifierOptions) afterLoad(context.Context) {}

// +checklocksignore
func (o *OpaqueInterfaceIdentifierOptions) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.SecretKey)
}

func (p *protocolMu) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.protocolMu"
}

func (p *protocolMu) StateFields() []string {
	return []string{
		"eps",
		"icmpRateLimitedTypes",
		"multicastForwardingDisp",
	}
}

func (p *protocolMu) beforeSave() {}

// +checklocksignore
func (p *protocolMu) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.eps)
	stateSinkObject.Save(1, &p.icmpRateLimitedTypes)
	stateSinkObject.Save(2, &p.multicastForwardingDisp)
}

func (p *protocolMu) afterLoad(context.Context) {}

// +checklocksignore
func (p *protocolMu) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.eps)
	stateSourceObject.Load(1, &p.icmpRateLimitedTypes)
	stateSourceObject.Load(2, &p.multicastForwardingDisp)
}

func (p *protocol) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.protocol"
}

func (p *protocol) StateFields() []string {
	return []string{
		"stack",
		"options",
		"mu",
		"defaultTTL",
		"fragmentation",
		"icmpRateLimiter",
		"multicastRouteTable",
	}
}

func (p *protocol) beforeSave() {}

// +checklocksignore
func (p *protocol) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.stack)
	stateSinkObject.Save(1, &p.options)
	stateSinkObject.Save(2, &p.mu)
	stateSinkObject.Save(3, &p.defaultTTL)
	stateSinkObject.Save(4, &p.fragmentation)
	stateSinkObject.Save(5, &p.icmpRateLimiter)
	stateSinkObject.Save(6, &p.multicastRouteTable)
}

func (p *protocol) afterLoad(context.Context) {}

// +checklocksignore
func (p *protocol) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.stack)
	stateSourceObject.Load(1, &p.options)
	stateSourceObject.Load(2, &p.mu)
	stateSourceObject.Load(3, &p.defaultTTL)
	stateSourceObject.Load(4, &p.fragmentation)
	stateSourceObject.Load(5, &p.icmpRateLimiter)
	stateSourceObject.Load(6, &p.multicastRouteTable)
}

func (o *Options) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.Options"
}

func (o *Options) StateFields() []string {
	return []string{
		"NDPConfigs",
		"AutoGenLinkLocal",
		"NDPDisp",
		"OpaqueIIDOpts",
		"TempIIDSeed",
		"MLD",
		"DADConfigs",
		"AllowExternalLoopbackTraffic",
	}
}

func (o *Options) beforeSave() {}

// +checklocksignore
func (o *Options) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.NDPConfigs)
	stateSinkObject.Save(1, &o.AutoGenLinkLocal)
	stateSinkObject.Save(2, &o.NDPDisp)
	stateSinkObject.Save(3, &o.OpaqueIIDOpts)
	stateSinkObject.Save(4, &o.TempIIDSeed)
	stateSinkObject.Save(5, &o.MLD)
	stateSinkObject.Save(6, &o.DADConfigs)
	stateSinkObject.Save(7, &o.AllowExternalLoopbackTraffic)
}

func (o *Options) afterLoad(context.Context) {}

// +checklocksignore
func (o *Options) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.NDPConfigs)
	stateSourceObject.Load(1, &o.AutoGenLinkLocal)
	stateSourceObject.Load(2, &o.NDPDisp)
	stateSourceObject.Load(3, &o.OpaqueIIDOpts)
	stateSourceObject.Load(4, &o.TempIIDSeed)
	stateSourceObject.Load(5, &o.MLD)
	stateSourceObject.Load(6, &o.DADConfigs)
	stateSourceObject.Load(7, &o.AllowExternalLoopbackTraffic)
}

func (m *MLDOptions) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.MLDOptions"
}

func (m *MLDOptions) StateFields() []string {
	return []string{
		"Enabled",
	}
}

func (m *MLDOptions) beforeSave() {}

// +checklocksignore
func (m *MLDOptions) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.Enabled)
}

func (m *MLDOptions) afterLoad(context.Context) {}

// +checklocksignore
func (m *MLDOptions) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.Enabled)
}

func (mld *mldState) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.mldState"
}

func (mld *mldState) StateFields() []string {
	return []string{
		"ep",
		"genericMulticastProtocol",
	}
}

func (mld *mldState) beforeSave() {}

// +checklocksignore
func (mld *mldState) StateSave(stateSinkObject state.Sink) {
	mld.beforeSave()
	stateSinkObject.Save(0, &mld.ep)
	stateSinkObject.Save(1, &mld.genericMulticastProtocol)
}

func (mld *mldState) afterLoad(context.Context) {}

// +checklocksignore
func (mld *mldState) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mld.ep)
	stateSourceObject.Load(1, &mld.genericMulticastProtocol)
}

func (c *NDPConfigurations) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.NDPConfigurations"
}

func (c *NDPConfigurations) StateFields() []string {
	return []string{
		"MaxRtrSolicitations",
		"RtrSolicitationInterval",
		"MaxRtrSolicitationDelay",
		"HandleRAs",
		"DiscoverDefaultRouters",
		"DiscoverMoreSpecificRoutes",
		"DiscoverOnLinkPrefixes",
		"AutoGenGlobalAddresses",
		"AutoGenAddressConflictRetries",
		"AutoGenTempGlobalAddresses",
		"MaxTempAddrValidLifetime",
		"MaxTempAddrPreferredLifetime",
		"RegenAdvanceDuration",
	}
}

func (c *NDPConfigurations) beforeSave() {}

// +checklocksignore
func (c *NDPConfigurations) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.MaxRtrSolicitations)
	stateSinkObject.Save(1, &c.RtrSolicitationInterval)
	stateSinkObject.Save(2, &c.MaxRtrSolicitationDelay)
	stateSinkObject.Save(3, &c.HandleRAs)
	stateSinkObject.Save(4, &c.DiscoverDefaultRouters)
	stateSinkObject.Save(5, &c.DiscoverMoreSpecificRoutes)
	stateSinkObject.Save(6, &c.DiscoverOnLinkPrefixes)
	stateSinkObject.Save(7, &c.AutoGenGlobalAddresses)
	stateSinkObject.Save(8, &c.AutoGenAddressConflictRetries)
	stateSinkObject.Save(9, &c.AutoGenTempGlobalAddresses)
	stateSinkObject.Save(10, &c.MaxTempAddrValidLifetime)
	stateSinkObject.Save(11, &c.MaxTempAddrPreferredLifetime)
	stateSinkObject.Save(12, &c.RegenAdvanceDuration)
}

func (c *NDPConfigurations) afterLoad(context.Context) {}

// +checklocksignore
func (c *NDPConfigurations) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.MaxRtrSolicitations)
	stateSourceObject.Load(1, &c.RtrSolicitationInterval)
	stateSourceObject.Load(2, &c.MaxRtrSolicitationDelay)
	stateSourceObject.Load(3, &c.HandleRAs)
	stateSourceObject.Load(4, &c.DiscoverDefaultRouters)
	stateSourceObject.Load(5, &c.DiscoverMoreSpecificRoutes)
	stateSourceObject.Load(6, &c.DiscoverOnLinkPrefixes)
	stateSourceObject.Load(7, &c.AutoGenGlobalAddresses)
	stateSourceObject.Load(8, &c.AutoGenAddressConflictRetries)
	stateSourceObject.Load(9, &c.AutoGenTempGlobalAddresses)
	stateSourceObject.Load(10, &c.MaxTempAddrValidLifetime)
	stateSourceObject.Load(11, &c.MaxTempAddrPreferredLifetime)
	stateSourceObject.Load(12, &c.RegenAdvanceDuration)
}

func (t *timer) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.timer"
}

func (t *timer) StateFields() []string {
	return []string{
		"done",
		"timer",
	}
}

func (t *timer) beforeSave() {}

// +checklocksignore
func (t *timer) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.done)
	stateSinkObject.Save(1, &t.timer)
}

func (t *timer) afterLoad(context.Context) {}

// +checklocksignore
func (t *timer) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.done)
	stateSourceObject.Load(1, &t.timer)
}

func (o *offLinkRoute) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.offLinkRoute"
}

func (o *offLinkRoute) StateFields() []string {
	return []string{
		"dest",
		"router",
	}
}

func (o *offLinkRoute) beforeSave() {}

// +checklocksignore
func (o *offLinkRoute) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.dest)
	stateSinkObject.Save(1, &o.router)
}

func (o *offLinkRoute) afterLoad(context.Context) {}

// +checklocksignore
func (o *offLinkRoute) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.dest)
	stateSourceObject.Load(1, &o.router)
}

func (ndp *ndpState) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.ndpState"
}

func (ndp *ndpState) StateFields() []string {
	return []string{
		"ep",
		"configs",
		"dad",
		"offLinkRoutes",
		"rtrSolicitTimer",
		"onLinkPrefixes",
		"slaacPrefixes",
		"dhcpv6Configuration",
		"temporaryIIDHistory",
		"temporaryAddressDesyncFactor",
	}
}

func (ndp *ndpState) beforeSave() {}

// +checklocksignore
func (ndp *ndpState) StateSave(stateSinkObject state.Sink) {
	ndp.beforeSave()
	stateSinkObject.Save(0, &ndp.ep)
	stateSinkObject.Save(1, &ndp.configs)
	stateSinkObject.Save(2, &ndp.dad)
	stateSinkObject.Save(3, &ndp.offLinkRoutes)
	stateSinkObject.Save(4, &ndp.rtrSolicitTimer)
	stateSinkObject.Save(5, &ndp.onLinkPrefixes)
	stateSinkObject.Save(6, &ndp.slaacPrefixes)
	stateSinkObject.Save(7, &ndp.dhcpv6Configuration)
	stateSinkObject.Save(8, &ndp.temporaryIIDHistory)
	stateSinkObject.Save(9, &ndp.temporaryAddressDesyncFactor)
}

func (ndp *ndpState) afterLoad(context.Context) {}

// +checklocksignore
func (ndp *ndpState) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ndp.ep)
	stateSourceObject.Load(1, &ndp.configs)
	stateSourceObject.Load(2, &ndp.dad)
	stateSourceObject.Load(3, &ndp.offLinkRoutes)
	stateSourceObject.Load(4, &ndp.rtrSolicitTimer)
	stateSourceObject.Load(5, &ndp.onLinkPrefixes)
	stateSourceObject.Load(6, &ndp.slaacPrefixes)
	stateSourceObject.Load(7, &ndp.dhcpv6Configuration)
	stateSourceObject.Load(8, &ndp.temporaryIIDHistory)
	stateSourceObject.Load(9, &ndp.temporaryAddressDesyncFactor)
}

func (o *offLinkRouteState) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.offLinkRouteState"
}

func (o *offLinkRouteState) StateFields() []string {
	return []string{
		"prf",
		"invalidationJob",
	}
}

func (o *offLinkRouteState) beforeSave() {}

// +checklocksignore
func (o *offLinkRouteState) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.prf)
	stateSinkObject.Save(1, &o.invalidationJob)
}

func (o *offLinkRouteState) afterLoad(context.Context) {}

// +checklocksignore
func (o *offLinkRouteState) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.prf)
	stateSourceObject.Load(1, &o.invalidationJob)
}

func (o *onLinkPrefixState) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.onLinkPrefixState"
}

func (o *onLinkPrefixState) StateFields() []string {
	return []string{
		"invalidationJob",
	}
}

func (o *onLinkPrefixState) beforeSave() {}

// +checklocksignore
func (o *onLinkPrefixState) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.invalidationJob)
}

func (o *onLinkPrefixState) afterLoad(context.Context) {}

// +checklocksignore
func (o *onLinkPrefixState) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.invalidationJob)
}

func (t *tempSLAACAddrState) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.tempSLAACAddrState"
}

func (t *tempSLAACAddrState) StateFields() []string {
	return []string{
		"deprecationJob",
		"invalidationJob",
		"regenJob",
		"createdAt",
		"addressEndpoint",
		"regenerated",
	}
}

func (t *tempSLAACAddrState) beforeSave() {}

// +checklocksignore
func (t *tempSLAACAddrState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.deprecationJob)
	stateSinkObject.Save(1, &t.invalidationJob)
	stateSinkObject.Save(2, &t.regenJob)
	stateSinkObject.Save(3, &t.createdAt)
	stateSinkObject.Save(4, &t.addressEndpoint)
	stateSinkObject.Save(5, &t.regenerated)
}

func (t *tempSLAACAddrState) afterLoad(context.Context) {}

// +checklocksignore
func (t *tempSLAACAddrState) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.deprecationJob)
	stateSourceObject.Load(1, &t.invalidationJob)
	stateSourceObject.Load(2, &t.regenJob)
	stateSourceObject.Load(3, &t.createdAt)
	stateSourceObject.Load(4, &t.addressEndpoint)
	stateSourceObject.Load(5, &t.regenerated)
}

func (s *stableAddrState) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.stableAddrState"
}

func (s *stableAddrState) StateFields() []string {
	return []string{
		"addressEndpoint",
		"localGenerationFailures",
	}
}

func (s *stableAddrState) beforeSave() {}

// +checklocksignore
func (s *stableAddrState) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.addressEndpoint)
	stateSinkObject.Save(1, &s.localGenerationFailures)
}

func (s *stableAddrState) afterLoad(context.Context) {}

// +checklocksignore
func (s *stableAddrState) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.addressEndpoint)
	stateSourceObject.Load(1, &s.localGenerationFailures)
}

func (s *slaacPrefixState) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.slaacPrefixState"
}

func (s *slaacPrefixState) StateFields() []string {
	return []string{
		"deprecationJob",
		"invalidationJob",
		"validUntil",
		"preferredUntil",
		"stableAddr",
		"tempAddrs",
		"generationAttempts",
		"maxGenerationAttempts",
	}
}

func (s *slaacPrefixState) beforeSave() {}

// +checklocksignore
func (s *slaacPrefixState) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.deprecationJob)
	stateSinkObject.Save(1, &s.invalidationJob)
	stateSinkObject.Save(2, &s.validUntil)
	stateSinkObject.Save(3, &s.preferredUntil)
	stateSinkObject.Save(4, &s.stableAddr)
	stateSinkObject.Save(5, &s.tempAddrs)
	stateSinkObject.Save(6, &s.generationAttempts)
	stateSinkObject.Save(7, &s.maxGenerationAttempts)
}

func (s *slaacPrefixState) afterLoad(context.Context) {}

// +checklocksignore
func (s *slaacPrefixState) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.deprecationJob)
	stateSourceObject.Load(1, &s.invalidationJob)
	stateSourceObject.Load(2, &s.validUntil)
	stateSourceObject.Load(3, &s.preferredUntil)
	stateSourceObject.Load(4, &s.stableAddr)
	stateSourceObject.Load(5, &s.tempAddrs)
	stateSourceObject.Load(6, &s.generationAttempts)
	stateSourceObject.Load(7, &s.maxGenerationAttempts)
}

func (s *Stats) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.Stats"
}

func (s *Stats) StateFields() []string {
	return []string{
		"IP",
		"ICMP",
		"UnhandledRouterAdvertisements",
	}
}

func (s *Stats) beforeSave() {}

// +checklocksignore
func (s *Stats) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.IP)
	stateSinkObject.Save(1, &s.ICMP)
	stateSinkObject.Save(2, &s.UnhandledRouterAdvertisements)
}

func (s *Stats) afterLoad(context.Context) {}

// +checklocksignore
func (s *Stats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.IP)
	stateSourceObject.Load(1, &s.ICMP)
	stateSourceObject.Load(2, &s.UnhandledRouterAdvertisements)
}

func (s *sharedStats) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.sharedStats"
}

func (s *sharedStats) StateFields() []string {
	return []string{
		"localStats",
		"ip",
		"icmp",
	}
}

func (s *sharedStats) beforeSave() {}

// +checklocksignore
func (s *sharedStats) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.localStats)
	stateSinkObject.Save(1, &s.ip)
	stateSinkObject.Save(2, &s.icmp)
}

func (s *sharedStats) afterLoad(context.Context) {}

// +checklocksignore
func (s *sharedStats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.localStats)
	stateSourceObject.Load(1, &s.ip)
	stateSourceObject.Load(2, &s.icmp)
}

func (m *multiCounterICMPv6PacketStats) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.multiCounterICMPv6PacketStats"
}

func (m *multiCounterICMPv6PacketStats) StateFields() []string {
	return []string{
		"echoRequest",
		"echoReply",
		"dstUnreachable",
		"packetTooBig",
		"timeExceeded",
		"paramProblem",
		"routerSolicit",
		"routerAdvert",
		"neighborSolicit",
		"neighborAdvert",
		"redirectMsg",
		"multicastListenerQuery",
		"multicastListenerReport",
		"multicastListenerReportV2",
		"multicastListenerDone",
	}
}

func (m *multiCounterICMPv6PacketStats) beforeSave() {}

// +checklocksignore
func (m *multiCounterICMPv6PacketStats) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.echoRequest)
	stateSinkObject.Save(1, &m.echoReply)
	stateSinkObject.Save(2, &m.dstUnreachable)
	stateSinkObject.Save(3, &m.packetTooBig)
	stateSinkObject.Save(4, &m.timeExceeded)
	stateSinkObject.Save(5, &m.paramProblem)
	stateSinkObject.Save(6, &m.routerSolicit)
	stateSinkObject.Save(7, &m.routerAdvert)
	stateSinkObject.Save(8, &m.neighborSolicit)
	stateSinkObject.Save(9, &m.neighborAdvert)
	stateSinkObject.Save(10, &m.redirectMsg)
	stateSinkObject.Save(11, &m.multicastListenerQuery)
	stateSinkObject.Save(12, &m.multicastListenerReport)
	stateSinkObject.Save(13, &m.multicastListenerReportV2)
	stateSinkObject.Save(14, &m.multicastListenerDone)
}

func (m *multiCounterICMPv6PacketStats) afterLoad(context.Context) {}

// +checklocksignore
func (m *multiCounterICMPv6PacketStats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.echoRequest)
	stateSourceObject.Load(1, &m.echoReply)
	stateSourceObject.Load(2, &m.dstUnreachable)
	stateSourceObject.Load(3, &m.packetTooBig)
	stateSourceObject.Load(4, &m.timeExceeded)
	stateSourceObject.Load(5, &m.paramProblem)
	stateSourceObject.Load(6, &m.routerSolicit)
	stateSourceObject.Load(7, &m.routerAdvert)
	stateSourceObject.Load(8, &m.neighborSolicit)
	stateSourceObject.Load(9, &m.neighborAdvert)
	stateSourceObject.Load(10, &m.redirectMsg)
	stateSourceObject.Load(11, &m.multicastListenerQuery)
	stateSourceObject.Load(12, &m.multicastListenerReport)
	stateSourceObject.Load(13, &m.multicastListenerReportV2)
	stateSourceObject.Load(14, &m.multicastListenerDone)
}

func (m *multiCounterICMPv6SentPacketStats) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.multiCounterICMPv6SentPacketStats"
}

func (m *multiCounterICMPv6SentPacketStats) StateFields() []string {
	return []string{
		"multiCounterICMPv6PacketStats",
		"dropped",
		"rateLimited",
	}
}

func (m *multiCounterICMPv6SentPacketStats) beforeSave() {}

// +checklocksignore
func (m *multiCounterICMPv6SentPacketStats) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.multiCounterICMPv6PacketStats)
	stateSinkObject.Save(1, &m.dropped)
	stateSinkObject.Save(2, &m.rateLimited)
}

func (m *multiCounterICMPv6SentPacketStats) afterLoad(context.Context) {}

// +checklocksignore
func (m *multiCounterICMPv6SentPacketStats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.multiCounterICMPv6PacketStats)
	stateSourceObject.Load(1, &m.dropped)
	stateSourceObject.Load(2, &m.rateLimited)
}

func (m *multiCounterICMPv6ReceivedPacketStats) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.multiCounterICMPv6ReceivedPacketStats"
}

func (m *multiCounterICMPv6ReceivedPacketStats) StateFields() []string {
	return []string{
		"multiCounterICMPv6PacketStats",
		"unrecognized",
		"invalid",
		"routerOnlyPacketsDroppedByHost",
	}
}

func (m *multiCounterICMPv6ReceivedPacketStats) beforeSave() {}

// +checklocksignore
func (m *multiCounterICMPv6ReceivedPacketStats) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.multiCounterICMPv6PacketStats)
	stateSinkObject.Save(1, &m.unrecognized)
	stateSinkObject.Save(2, &m.invalid)
	stateSinkObject.Save(3, &m.routerOnlyPacketsDroppedByHost)
}

func (m *multiCounterICMPv6ReceivedPacketStats) afterLoad(context.Context) {}

// +checklocksignore
func (m *multiCounterICMPv6ReceivedPacketStats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.multiCounterICMPv6PacketStats)
	stateSourceObject.Load(1, &m.unrecognized)
	stateSourceObject.Load(2, &m.invalid)
	stateSourceObject.Load(3, &m.routerOnlyPacketsDroppedByHost)
}

func (m *multiCounterICMPv6Stats) StateTypeName() string {
	return "pkg/tcpip/network/ipv6.multiCounterICMPv6Stats"
}

func (m *multiCounterICMPv6Stats) StateFields() []string {
	return []string{
		"packetsSent",
		"packetsReceived",
	}
}

func (m *multiCounterICMPv6Stats) beforeSave() {}

// +checklocksignore
func (m *multiCounterICMPv6Stats) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.packetsSent)
	stateSinkObject.Save(1, &m.packetsReceived)
}

func (m *multiCounterICMPv6Stats) afterLoad(context.Context) {}

// +checklocksignore
func (m *multiCounterICMPv6Stats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.packetsSent)
	stateSourceObject.Load(1, &m.packetsReceived)
}

func init() {
	state.Register((*icmpv6DestinationUnreachableSockError)(nil))
	state.Register((*icmpv6DestinationNetworkUnreachableSockError)(nil))
	state.Register((*icmpv6DestinationPortUnreachableSockError)(nil))
	state.Register((*icmpv6DestinationAddressUnreachableSockError)(nil))
	state.Register((*icmpv6PacketTooBigSockError)(nil))
	state.Register((*endpointMu)(nil))
	state.Register((*dadMu)(nil))
	state.Register((*endpointDAD)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*OpaqueInterfaceIdentifierOptions)(nil))
	state.Register((*protocolMu)(nil))
	state.Register((*protocol)(nil))
	state.Register((*Options)(nil))
	state.Register((*MLDOptions)(nil))
	state.Register((*mldState)(nil))
	state.Register((*NDPConfigurations)(nil))
	state.Register((*timer)(nil))
	state.Register((*offLinkRoute)(nil))
	state.Register((*ndpState)(nil))
	state.Register((*offLinkRouteState)(nil))
	state.Register((*onLinkPrefixState)(nil))
	state.Register((*tempSLAACAddrState)(nil))
	state.Register((*stableAddrState)(nil))
	state.Register((*slaacPrefixState)(nil))
	state.Register((*Stats)(nil))
	state.Register((*sharedStats)(nil))
	state.Register((*multiCounterICMPv6PacketStats)(nil))
	state.Register((*multiCounterICMPv6SentPacketStats)(nil))
	state.Register((*multiCounterICMPv6ReceivedPacketStats)(nil))
	state.Register((*multiCounterICMPv6Stats)(nil))
}
