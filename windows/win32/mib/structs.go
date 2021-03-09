// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package mib implements the Windows.Win32.Mib namespace.
package mib

type MIB_IF_ROW2 struct {
	InterfaceLuid               int
	InterfaceIndex              int
	InterfaceGuid               int
	Alias                       int
	Description                 int
	PhysicalAddressLength       int
	PhysicalAddress             int
	PermanentPhysicalAddress    int
	Mtu                         int
	Type                        int
	TunnelType                  int
	MediaType                   int
	PhysicalMediumType          int
	AccessType                  int
	DirectionType               int
	InterfaceAndOperStatusFlags int
	OperStatus                  int
	AdminStatus                 int
	MediaConnectState           int
	NetworkGuid                 int
	ConnectionType              int
	TransmitLinkSpeed           int
	ReceiveLinkSpeed            int
	InOctets                    int
	InUcastPkts                 int
	InNUcastPkts                int
	InDiscards                  int
	InErrors                    int
	InUnknownProtos             int
	InUcastOctets               int
	InMulticastOctets           int
	InBroadcastOctets           int
	OutOctets                   int
	OutUcastPkts                int
	OutNUcastPkts               int
	OutDiscards                 int
	OutErrors                   int
	OutUcastOctets              int
	OutMulticastOctets          int
	OutBroadcastOctets          int
	OutQLen                     int
}

type MIB_IF_TABLE2 struct {
	NumEntries int
	Table      int
}

type MIB_IPINTERFACE_ROW struct {
	Family                               int
	InterfaceLuid                        int
	InterfaceIndex                       int
	MaxReassemblySize                    int
	InterfaceIdentifier                  int
	MinRouterAdvertisementInterval       int
	MaxRouterAdvertisementInterval       int
	AdvertisingEnabled                   int
	ForwardingEnabled                    int
	WeakHostSend                         int
	WeakHostReceive                      int
	UseAutomaticMetric                   int
	UseNeighborUnreachabilityDetection   int
	ManagedAddressConfigurationSupported int
	OtherStatefulConfigurationSupported  int
	AdvertiseDefaultRoute                int
	RouterDiscoveryBehavior              int
	DadTransmits                         int
	BaseReachableTime                    int
	RetransmitTime                       int
	PathMtuDiscoveryTimeout              int
	LinkLocalAddressBehavior             int
	LinkLocalAddressTimeout              int
	ZoneIndices                          int
	SitePrefixLength                     int
	Metric                               int
	NlMtu                                int
	Connected                            int
	SupportsWakeUpPatterns               int
	SupportsNeighborDiscovery            int
	SupportsRouterDiscovery              int
	ReachableTime                        int
	TransmitOffload                      int
	ReceiveOffload                       int
	DisableDefaultRoutes                 int
}

type MIB_IPINTERFACE_TABLE struct {
	NumEntries int
	Table      int
}

type MIB_IFSTACK_ROW struct {
	HigherLayerInterfaceIndex int
	LowerLayerInterfaceIndex  int
}

type MIB_INVERTEDIFSTACK_ROW struct {
	LowerLayerInterfaceIndex  int
	HigherLayerInterfaceIndex int
}

type MIB_IFSTACK_TABLE struct {
	NumEntries int
	Table      int
}

type MIB_INVERTEDIFSTACK_TABLE struct {
	NumEntries int
	Table      int
}

type MIB_IP_NETWORK_CONNECTION_BANDWIDTH_ESTIMATES struct {
	InboundBandwidthInformation  int
	OutboundBandwidthInformation int
}

type MIB_UNICASTIPADDRESS_ROW struct {
	Address            int
	InterfaceLuid      int
	InterfaceIndex     int
	PrefixOrigin       int
	SuffixOrigin       int
	ValidLifetime      int
	PreferredLifetime  int
	OnLinkPrefixLength int
	SkipAsSource       int
	DadState           int
	ScopeId            int
	CreationTimeStamp  int
}

type MIB_UNICASTIPADDRESS_TABLE struct {
	NumEntries int
	Table      int
}

type MIB_ANYCASTIPADDRESS_ROW struct {
	Address        int
	InterfaceLuid  int
	InterfaceIndex int
	ScopeId        int
}

type MIB_ANYCASTIPADDRESS_TABLE struct {
	NumEntries int
	Table      int
}

type MIB_MULTICASTIPADDRESS_ROW struct {
	Address        int
	InterfaceIndex int
	InterfaceLuid  int
	ScopeId        int
}

type MIB_MULTICASTIPADDRESS_TABLE struct {
	NumEntries int
	Table      int
}

type MIB_IPFORWARD_ROW2 struct {
	InterfaceLuid        int
	InterfaceIndex       int
	DestinationPrefix    int
	NextHop              int
	SitePrefixLength     int
	ValidLifetime        int
	PreferredLifetime    int
	Metric               int
	Protocol             int
	Loopback             int
	AutoconfigureAddress int
	Publish              int
	Immortal             int
	Age                  int
	Origin               int
}

type MIB_IPFORWARD_TABLE2 struct {
	NumEntries int
	Table      int
}

type MIB_IPPATH_ROW struct {
	Source            int
	Destination       int
	InterfaceLuid     int
	InterfaceIndex    int
	CurrentNextHop    int
	PathMtu           int
	RttMean           int
	RttDeviation      int
	Anonymous         int
	IsReachable       int
	LinkTransmitSpeed int
	LinkReceiveSpeed  int
}

type MIB_IPPATH_TABLE struct {
	NumEntries int
	Table      int
}

type MIB_IPNET_ROW2 struct {
	Address               int
	InterfaceIndex        int
	InterfaceLuid         int
	PhysicalAddress       int
	PhysicalAddressLength int
	State                 int
	Anonymous             int
	ReachabilityTime      int
}

type MIB_IPNET_TABLE2 struct {
	NumEntries int
	Table      int
}

type MIB_OPAQUE_QUERY struct {
	DWVARID      int
	RGDWVARINDEX int
}

type MIB_IFNUMBER struct {
	DWVALUE int
}

type MIB_IFROW struct {
	WSZNAME           int
	DWINDEX           int
	DWTYPE            int
	DWMTU             int
	DWSPEED           int
	DWPHYSADDRLEN     int
	BPHYSADDR         int
	DWADMINSTATUS     int
	DWOPERSTATUS      int
	DWLASTCHANGE      int
	DWINOCTETS        int
	DWINUCASTPKTS     int
	DWINNUCASTPKTS    int
	DWINDISCARDS      int
	DWINERRORS        int
	DWINUNKNOWNPROTOS int
	DWOUTOCTETS       int
	DWOUTUCASTPKTS    int
	DWOUTNUCASTPKTS   int
	DWOUTDISCARDS     int
	DWOUTERRORS       int
	DWOUTQLEN         int
	DWDESCRLEN        int
	BDESCR            int
}

type MIB_IFTABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_IPADDRROW_XP struct {
	DWADDR      int
	DWINDEX     int
	DWMASK      int
	DWBCASTADDR int
	DWREASMSIZE int
	UNUSED1     int
	WTYPE       int
}

type MIB_IPADDRROW_W2K struct {
	DWADDR      int
	DWINDEX     int
	DWMASK      int
	DWBCASTADDR int
	DWREASMSIZE int
	UNUSED1     int
	UNUSED2     int
}

type MIB_IPADDRTABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_IPFORWARDNUMBER struct {
	DWVALUE int
}

type MIB_IPFORWARDROW struct {
	DWFORWARDDEST      int
	DWFORWARDMASK      int
	DWFORWARDPOLICY    int
	DWFORWARDNEXTHOP   int
	DWFORWARDIFINDEX   int
	Anonymous1         int
	Anonymous2         int
	DWFORWARDAGE       int
	DWFORWARDNEXTHOPAS int
	DWFORWARDMETRIC1   int
	DWFORWARDMETRIC2   int
	DWFORWARDMETRIC3   int
	DWFORWARDMETRIC4   int
	DWFORWARDMETRIC5   int
}

type MIB_IPFORWARDTABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_IPNETROW_LH struct {
	DWINDEX       int
	DWPHYSADDRLEN int
	BPHYSADDR     int
	DWADDR        int
	Anonymous     int
}

type MIB_IPNETROW_W2K struct {
	DWINDEX       int
	DWPHYSADDRLEN int
	BPHYSADDR     int
	DWADDR        int
	DWTYPE        int
}

type MIB_IPNETTABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_IPSTATS_LH struct {
	Anonymous         int
	DWDEFAULTTTL      int
	DWINRECEIVES      int
	DWINHDRERRORS     int
	DWINADDRERRORS    int
	DWFORWDATAGRAMS   int
	DWINUNKNOWNPROTOS int
	DWINDISCARDS      int
	DWINDELIVERS      int
	DWOUTREQUESTS     int
	DWROUTINGDISCARDS int
	DWOUTDISCARDS     int
	DWOUTNOROUTES     int
	DWREASMTIMEOUT    int
	DWREASMREQDS      int
	DWREASMOKS        int
	DWREASMFAILS      int
	DWFRAGOKS         int
	DWFRAGFAILS       int
	DWFRAGCREATES     int
	DWNUMIF           int
	DWNUMADDR         int
	DWNUMROUTES       int
}

type MIB_IPSTATS_W2K struct {
	DWFORWARDING      int
	DWDEFAULTTTL      int
	DWINRECEIVES      int
	DWINHDRERRORS     int
	DWINADDRERRORS    int
	DWFORWDATAGRAMS   int
	DWINUNKNOWNPROTOS int
	DWINDISCARDS      int
	DWINDELIVERS      int
	DWOUTREQUESTS     int
	DWROUTINGDISCARDS int
	DWOUTDISCARDS     int
	DWOUTNOROUTES     int
	DWREASMTIMEOUT    int
	DWREASMREQDS      int
	DWREASMOKS        int
	DWREASMFAILS      int
	DWFRAGOKS         int
	DWFRAGFAILS       int
	DWFRAGCREATES     int
	DWNUMIF           int
	DWNUMADDR         int
	DWNUMROUTES       int
}

type MIBICMPSTATS struct {
	DWMSGS          int
	DWERRORS        int
	DWDESTUNREACHS  int
	DWTIMEEXCDS     int
	DWPARMPROBS     int
	DWSRCQUENCHS    int
	DWREDIRECTS     int
	DWECHOS         int
	DWECHOREPS      int
	DWTIMESTAMPS    int
	DWTIMESTAMPREPS int
	DWADDRMASKS     int
	DWADDRMASKREPS  int
}

type MIBICMPINFO struct {
	ICMPINSTATS  int
	ICMPOUTSTATS int
}

type MIB_ICMP struct {
	STATS int
}

type MIBICMPSTATS_EX_XPSP1 struct {
	DWMSGS        int
	DWERRORS      int
	RGDWTYPECOUNT int
}

type MIB_ICMP_EX_XPSP1 struct {
	ICMPINSTATS  int
	ICMPOUTSTATS int
}

type MIB_IPMCAST_OIF_XP struct {
	DWOUTIFINDEX  int
	DWNEXTHOPADDR int
	DWRESERVED    int
	DWRESERVED1   int
}

type MIB_IPMCAST_OIF_W2K struct {
	DWOUTIFINDEX  int
	DWNEXTHOPADDR int
	PVRESERVED    int
	DWRESERVED    int
}

type MIB_IPMCAST_MFE struct {
	DWGROUP         int
	DWSOURCE        int
	DWSRCMASK       int
	DWUPSTRMNGBR    int
	DWINIFINDEX     int
	DWINIFPROTOCOL  int
	DWROUTEPROTOCOL int
	DWROUTENETWORK  int
	DWROUTEMASK     int
	ULUPTIME        int
	ULEXPIRYTIME    int
	ULTIMEOUT       int
	ULNUMOUTIF      int
	FFLAGS          int
	DWRESERVED      int
	RGMIOOUTINFO    int
}

type MIB_MFE_TABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_IPMCAST_OIF_STATS_LH struct {
	DWOUTIFINDEX  int
	DWNEXTHOPADDR int
	DWDIALCONTEXT int
	ULTTLTOOLOW   int
	ULFRAGNEEDED  int
	ULOUTPACKETS  int
	ULOUTDISCARDS int
}

type MIB_IPMCAST_OIF_STATS_W2K struct {
	DWOUTIFINDEX  int
	DWNEXTHOPADDR int
	PVDIALCONTEXT int
	ULTTLTOOLOW   int
	ULFRAGNEEDED  int
	ULOUTPACKETS  int
	ULOUTDISCARDS int
}

type MIB_IPMCAST_MFE_STATS struct {
	DWGROUP           int
	DWSOURCE          int
	DWSRCMASK         int
	DWUPSTRMNGBR      int
	DWINIFINDEX       int
	DWINIFPROTOCOL    int
	DWROUTEPROTOCOL   int
	DWROUTENETWORK    int
	DWROUTEMASK       int
	ULUPTIME          int
	ULEXPIRYTIME      int
	ULNUMOUTIF        int
	ULINPKTS          int
	ULINOCTETS        int
	ULPKTSDIFFERENTIF int
	ULQUEUEOVERFLOW   int
	RGMIOSOUTSTATS    int
}

type MIB_MFE_STATS_TABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_IPMCAST_MFE_STATS_EX_XP struct {
	DWGROUP           int
	DWSOURCE          int
	DWSRCMASK         int
	DWUPSTRMNGBR      int
	DWINIFINDEX       int
	DWINIFPROTOCOL    int
	DWROUTEPROTOCOL   int
	DWROUTENETWORK    int
	DWROUTEMASK       int
	ULUPTIME          int
	ULEXPIRYTIME      int
	ULNUMOUTIF        int
	ULINPKTS          int
	ULINOCTETS        int
	ULPKTSDIFFERENTIF int
	ULQUEUEOVERFLOW   int
	ULUNINITMFE       int
	ULNEGATIVEMFE     int
	ULINDISCARDS      int
	ULINHDRERRORS     int
	ULTOTALOUTPACKETS int
	RGMIOSOUTSTATS    int
}

type MIB_MFE_STATS_TABLE_EX_XP struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_IPMCAST_GLOBAL struct {
	DWENABLE int
}

type MIB_IPMCAST_IF_ENTRY struct {
	DWIFINDEX        int
	DWTTL            int
	DWPROTOCOL       int
	DWRATELIMIT      int
	ULINMCASTOCTETS  int
	ULOUTMCASTOCTETS int
}

type MIB_IPMCAST_IF_TABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCPROW_LH struct {
	Anonymous    int
	DWLOCALADDR  int
	DWLOCALPORT  int
	DWREMOTEADDR int
	DWREMOTEPORT int
}

type MIB_TCPROW_W2K struct {
	DWSTATE      int
	DWLOCALADDR  int
	DWLOCALPORT  int
	DWREMOTEADDR int
	DWREMOTEPORT int
}

type MIB_TCPTABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCPROW2 struct {
	DWSTATE        int
	DWLOCALADDR    int
	DWLOCALPORT    int
	DWREMOTEADDR   int
	DWREMOTEPORT   int
	DWOWNINGPID    int
	DWOFFLOADSTATE int
}

type MIB_TCPTABLE2 struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCPROW_OWNER_PID struct {
	DWSTATE      int
	DWLOCALADDR  int
	DWLOCALPORT  int
	DWREMOTEADDR int
	DWREMOTEPORT int
	DWOWNINGPID  int
}

type MIB_TCPTABLE_OWNER_PID struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCPROW_OWNER_MODULE struct {
	DWSTATE           int
	DWLOCALADDR       int
	DWLOCALPORT       int
	DWREMOTEADDR      int
	DWREMOTEPORT      int
	DWOWNINGPID       int
	LICREATETIMESTAMP int
	OwningModuleInfo  int
}

type MIB_TCPTABLE_OWNER_MODULE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCP6ROW struct {
	State           int
	LocalAddr       int
	DWLOCALSCOPEID  int
	DWLOCALPORT     int
	RemoteAddr      int
	DWREMOTESCOPEID int
	DWREMOTEPORT    int
}

type MIB_TCP6TABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCP6ROW2 struct {
	LocalAddr       int
	DWLOCALSCOPEID  int
	DWLOCALPORT     int
	RemoteAddr      int
	DWREMOTESCOPEID int
	DWREMOTEPORT    int
	State           int
	DWOWNINGPID     int
	DWOFFLOADSTATE  int
}

type MIB_TCP6TABLE2 struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCP6ROW_OWNER_PID struct {
	UCLOCALADDR     int
	DWLOCALSCOPEID  int
	DWLOCALPORT     int
	UCREMOTEADDR    int
	DWREMOTESCOPEID int
	DWREMOTEPORT    int
	DWSTATE         int
	DWOWNINGPID     int
}

type MIB_TCP6TABLE_OWNER_PID struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCP6ROW_OWNER_MODULE struct {
	UCLOCALADDR       int
	DWLOCALSCOPEID    int
	DWLOCALPORT       int
	UCREMOTEADDR      int
	DWREMOTESCOPEID   int
	DWREMOTEPORT      int
	DWSTATE           int
	DWOWNINGPID       int
	LICREATETIMESTAMP int
	OwningModuleInfo  int
}

type MIB_TCP6TABLE_OWNER_MODULE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_TCPSTATS_LH struct {
	Anonymous      int
	DWRTOMIN       int
	DWRTOMAX       int
	DWMAXCONN      int
	DWACTIVEOPENS  int
	DWPASSIVEOPENS int
	DWATTEMPTFAILS int
	DWESTABRESETS  int
	DWCURRESTAB    int
	DWINSEGS       int
	DWOUTSEGS      int
	DWRETRANSSEGS  int
	DWINERRS       int
	DWOUTRSTS      int
	DWNUMCONNS     int
}

type MIB_TCPSTATS_W2K struct {
	DWRTOALGORITHM int
	DWRTOMIN       int
	DWRTOMAX       int
	DWMAXCONN      int
	DWACTIVEOPENS  int
	DWPASSIVEOPENS int
	DWATTEMPTFAILS int
	DWESTABRESETS  int
	DWCURRESTAB    int
	DWINSEGS       int
	DWOUTSEGS      int
	DWRETRANSSEGS  int
	DWINERRS       int
	DWOUTRSTS      int
	DWNUMCONNS     int
}

type MIB_TCPSTATS2 struct {
	RtoAlgorithm   int
	DWRTOMIN       int
	DWRTOMAX       int
	DWMAXCONN      int
	DWACTIVEOPENS  int
	DWPASSIVEOPENS int
	DWATTEMPTFAILS int
	DWESTABRESETS  int
	DWCURRESTAB    int
	DW64INSEGS     int
	DW64OUTSEGS    int
	DWRETRANSSEGS  int
	DWINERRS       int
	DWOUTRSTS      int
	DWNUMCONNS     int
}

type MIB_UDPROW struct {
	DWLOCALADDR int
	DWLOCALPORT int
}

type MIB_UDPTABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_UDPROW_OWNER_PID struct {
	DWLOCALADDR int
	DWLOCALPORT int
	DWOWNINGPID int
}

type MIB_UDPTABLE_OWNER_PID struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_UDPROW_OWNER_MODULE struct {
	DWLOCALADDR       int
	DWLOCALPORT       int
	DWOWNINGPID       int
	LICREATETIMESTAMP int
	Anonymous         int
	OwningModuleInfo  int
}

type MIB_UDPTABLE_OWNER_MODULE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_UDP6ROW struct {
	DWLOCALADDR    int
	DWLOCALSCOPEID int
	DWLOCALPORT    int
}

type MIB_UDP6TABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_UDP6ROW_OWNER_PID struct {
	UCLOCALADDR    int
	DWLOCALSCOPEID int
	DWLOCALPORT    int
	DWOWNINGPID    int
}

type MIB_UDP6TABLE_OWNER_PID struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_UDP6ROW_OWNER_MODULE struct {
	UCLOCALADDR       int
	DWLOCALSCOPEID    int
	DWLOCALPORT       int
	DWOWNINGPID       int
	LICREATETIMESTAMP int
	Anonymous         int
	OwningModuleInfo  int
}

type MIB_UDP6TABLE_OWNER_MODULE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_UDPSTATS struct {
	DWINDATAGRAMS  int
	DWNOPORTS      int
	DWINERRORS     int
	DWOUTDATAGRAMS int
	DWNUMADDRS     int
}

type MIB_UDPSTATS2 struct {
	DW64INDATAGRAMS  int
	DWNOPORTS        int
	DWINERRORS       int
	DW64OUTDATAGRAMS int
	DWNUMADDRS       int
}

type MIB_IPMCAST_BOUNDARY struct {
	DWIFINDEX      int
	DWGROUPADDRESS int
	DWGROUPMASK    int
	DWSTATUS       int
}

type MIB_IPMCAST_BOUNDARY_TABLE struct {
	DWNUMENTRIES int
	TABLE        int
}

type MIB_BOUNDARYROW struct {
	DWGROUPADDRESS int
	DWGROUPMASK    int
}

type MIB_MCAST_LIMIT_ROW struct {
	DWTTL       int
	DWRATELIMIT int
}

type MIB_IPMCAST_SCOPE struct {
	DWGROUPADDRESS int
	DWGROUPMASK    int
	SNNAMEBUFFER   int
	DWSTATUS       int
}

type MIB_BEST_IF struct {
	DWDESTADDR int
	DWIFINDEX  int
}

type MIB_PROXYARP struct {
	DWADDRESS int
	DWMASK    int
	DWIFINDEX int
}

type MIB_IFSTATUS struct {
	DWIFINDEX           int
	DWADMINSTATUS       int
	DWOPERATIONALSTATUS int
	BMHBEATACTIVE       int
	BMHBEATALIVE        int
}

type MIB_OPAQUE_INFO struct {
	DWID      int
	Anonymous int
}

type NL_INTERFACE_OFFLOAD_ROD struct {
	BITFIELD int
}