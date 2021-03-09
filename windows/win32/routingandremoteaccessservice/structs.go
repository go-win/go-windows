// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package routingandremoteaccessservice implements the Windows.Win32.RoutingAndRemoteAccessService namespace.
package routingandremoteaccessservice

type RASIPADDR struct {
	A int
	B int
	C int
	D int
}

type TAGRASTUNNELENDPOINT struct {
	DWTYPE    int
	Anonymous int
}

type HRASCONN__ struct {
	UNUSED int
}

type TAGRASCONNW struct {
	DWSIZE            int
	HRASCONN          int
	SZENTRYNAME       int
	SZDEVICETYPE      int
	SZDEVICENAME      int
	SZPHONEBOOK       int
	DWSUBENTRY        int
	GUIDENTRY         int
	DWFLAGS           int
	LUID              int
	GUIDCORRELATIONID int
}

type TAGRASCONNA struct {
	DWSIZE            int
	HRASCONN          int
	SZENTRYNAME       int
	SZDEVICETYPE      int
	SZDEVICENAME      int
	SZPHONEBOOK       int
	DWSUBENTRY        int
	GUIDENTRY         int
	DWFLAGS           int
	LUID              int
	GUIDCORRELATIONID int
}

type TAGRASCONNSTATUSW struct {
	DWSIZE          int
	RASCONNSTATE    int
	DWERROR         int
	SZDEVICETYPE    int
	SZDEVICENAME    int
	SZPHONENUMBER   int
	LOCALENDPOINT   int
	REMOTEENDPOINT  int
	RASCONNSUBSTATE int
}

type TAGRASCONNSTATUSA struct {
	DWSIZE          int
	RASCONNSTATE    int
	DWERROR         int
	SZDEVICETYPE    int
	SZDEVICENAME    int
	SZPHONENUMBER   int
	LOCALENDPOINT   int
	REMOTEENDPOINT  int
	RASCONNSUBSTATE int
}

type TAGRASDIALPARAMSW struct {
	DWSIZE           int
	SZENTRYNAME      int
	SZPHONENUMBER    int
	SZCALLBACKNUMBER int
	SZUSERNAME       int
	SZPASSWORD       int
	SZDOMAIN         int
	DWSUBENTRY       int
	DWCALLBACKID     int
	DWIFINDEX        int
	SZENCPASSWORD    int
}

type TAGRASDIALPARAMSA struct {
	DWSIZE           int
	SZENTRYNAME      int
	SZPHONENUMBER    int
	SZCALLBACKNUMBER int
	SZUSERNAME       int
	SZPASSWORD       int
	SZDOMAIN         int
	DWSUBENTRY       int
	DWCALLBACKID     int
	DWIFINDEX        int
	SZENCPASSWORD    int
}

type TAGRASEAPINFO struct {
	DWSIZEOFEAPINFO int
	PBEAPINFO       int
}

type RASDEVSPECIFICINFO struct {
	DWSIZE            int
	PBDEVSPECIFICINFO int
}

type TAGRASDIALEXTENSIONS struct {
	DWSIZE             int
	DWFOPTIONS         int
	HWNDPARENT         int
	RESERVED           int
	RESERVED1          int
	RasEapInfo         int
	FSKIPPPPAUTH       int
	RasDevSpecificInfo int
}

type TAGRASENTRYNAMEW struct {
	DWSIZE          int
	SZENTRYNAME     int
	DWFLAGS         int
	SZPHONEBOOKPATH int
}

type TAGRASENTRYNAMEA struct {
	DWSIZE          int
	SZENTRYNAME     int
	DWFLAGS         int
	SZPHONEBOOKPATH int
}

type TAGRASAMBW struct {
	DWSIZE         int
	DWERROR        int
	SZNETBIOSERROR int
	BLANA          int
}

type TAGRASAMBA struct {
	DWSIZE         int
	DWERROR        int
	SZNETBIOSERROR int
	BLANA          int
}

type TAGRASPPPNBFW struct {
	DWSIZE            int
	DWERROR           int
	DWNETBIOSERROR    int
	SZNETBIOSERROR    int
	SZWORKSTATIONNAME int
	BLANA             int
}

type TAGRASPPPNBFA struct {
	DWSIZE            int
	DWERROR           int
	DWNETBIOSERROR    int
	SZNETBIOSERROR    int
	SZWORKSTATIONNAME int
	BLANA             int
}

type TAGRASIPXW struct {
	DWSIZE       int
	DWERROR      int
	SZIPXADDRESS int
}

type TAGRASPPPIPXA struct {
	DWSIZE       int
	DWERROR      int
	SZIPXADDRESS int
}

type TAGRASPPPIPW struct {
	DWSIZE            int
	DWERROR           int
	SZIPADDRESS       int
	SZSERVERIPADDRESS int
	DWOPTIONS         int
	DWSERVEROPTIONS   int
}

type TAGRASPPPIPA struct {
	DWSIZE            int
	DWERROR           int
	SZIPADDRESS       int
	SZSERVERIPADDRESS int
	DWOPTIONS         int
	DWSERVEROPTIONS   int
}

type TAGRASPPPIPV6 struct {
	DWSIZE                    int
	DWERROR                   int
	BLOCALINTERFACEIDENTIFIER int
	BPEERINTERFACEIDENTIFIER  int
	BLOCALCOMPRESSIONPROTOCOL int
	BPEERCOMPRESSIONPROTOCOL  int
}

type TAGRASPPPLCPW struct {
	DWSIZE                         int
	FBUNDLED                       int
	DWERROR                        int
	DWAUTHENTICATIONPROTOCOL       int
	DWAUTHENTICATIONDATA           int
	DWEAPTYPEID                    int
	DWSERVERAUTHENTICATIONPROTOCOL int
	DWSERVERAUTHENTICATIONDATA     int
	DWSERVEREAPTYPEID              int
	FMULTILINK                     int
	DWTERMINATEREASON              int
	DWSERVERTERMINATEREASON        int
	SZREPLYMESSAGE                 int
	DWOPTIONS                      int
	DWSERVEROPTIONS                int
}

type TAGRASPPPLCPA struct {
	DWSIZE                         int
	FBUNDLED                       int
	DWERROR                        int
	DWAUTHENTICATIONPROTOCOL       int
	DWAUTHENTICATIONDATA           int
	DWEAPTYPEID                    int
	DWSERVERAUTHENTICATIONPROTOCOL int
	DWSERVERAUTHENTICATIONDATA     int
	DWSERVEREAPTYPEID              int
	FMULTILINK                     int
	DWTERMINATEREASON              int
	DWSERVERTERMINATEREASON        int
	SZREPLYMESSAGE                 int
	DWOPTIONS                      int
	DWSERVEROPTIONS                int
}

type TAGRASPPPCCP struct {
	DWSIZE                       int
	DWERROR                      int
	DWCOMPRESSIONALGORITHM       int
	DWOPTIONS                    int
	DWSERVERCOMPRESSIONALGORITHM int
	DWSERVEROPTIONS              int
}

type RASPPP_PROJECTION_INFO struct {
	DWIPV4NEGOTIATIONERROR          int
	IPV4ADDRESS                     int
	IPV4SERVERADDRESS               int
	DWIPV4OPTIONS                   int
	DWIPV4SERVEROPTIONS             int
	DWIPV6NEGOTIATIONERROR          int
	BINTERFACEIDENTIFIER            int
	BSERVERINTERFACEIDENTIFIER      int
	FBUNDLED                        int
	FMULTILINK                      int
	DWAUTHENTICATIONPROTOCOL        int
	DWAUTHENTICATIONDATA            int
	DWSERVERAUTHENTICATIONPROTOCOL  int
	DWSERVERAUTHENTICATIONDATA      int
	DWEAPTYPEID                     int
	DWSERVEREAPTYPEID               int
	DWLCPOPTIONS                    int
	DWLCPSERVEROPTIONS              int
	DWCCPERROR                      int
	DWCCPCOMPRESSIONALGORITHM       int
	DWCCPSERVERCOMPRESSIONALGORITHM int
	DWCCPOPTIONS                    int
	DWCCPSERVEROPTIONS              int
}

type RASIKEV2_PROJECTION_INFO struct {
	DWIPV4NEGOTIATIONERROR   int
	IPV4ADDRESS              int
	IPV4SERVERADDRESS        int
	DWIPV6NEGOTIATIONERROR   int
	IPV6ADDRESS              int
	IPV6SERVERADDRESS        int
	DWPREFIXLENGTH           int
	DWAUTHENTICATIONPROTOCOL int
	DWEAPTYPEID              int
	DWFLAGS                  int
	DWENCRYPTIONMETHOD       int
	NUMIPV4SERVERADDRESSES   int
	IPV4SERVERADDRESSES      int
	NUMIPV6SERVERADDRESSES   int
	IPV6SERVERADDRESSES      int
}

type RAS_PROJECTION_INFO struct {
	VERSION   int
	TYPE      int
	Anonymous int
}

type TAGRASDEVINFOW struct {
	DWSIZE       int
	SZDEVICETYPE int
	SZDEVICENAME int
}

type TAGRASDEVINFOA struct {
	DWSIZE       int
	SZDEVICETYPE int
	SZDEVICENAME int
}

type RASCTRYINFO struct {
	DWSIZE              int
	DWCOUNTRYID         int
	DWNEXTCOUNTRYID     int
	DWCOUNTRYCODE       int
	DWCOUNTRYNAMEOFFSET int
}

type TAGRASENTRYA struct {
	DWSIZE                     int
	DWFOPTIONS                 int
	DWCOUNTRYID                int
	DWCOUNTRYCODE              int
	SZAREACODE                 int
	SZLOCALPHONENUMBER         int
	DWALTERNATEOFFSET          int
	IPADDR                     int
	IPADDRDNS                  int
	IPADDRDNSALT               int
	IPADDRWINS                 int
	IPADDRWINSALT              int
	DWFRAMESIZE                int
	DWFNETPROTOCOLS            int
	DWFRAMINGPROTOCOL          int
	SZSCRIPT                   int
	SZAUTODIALDLL              int
	SZAUTODIALFUNC             int
	SZDEVICETYPE               int
	SZDEVICENAME               int
	SZX25PADTYPE               int
	SZX25ADDRESS               int
	SZX25FACILITIES            int
	SZX25USERDATA              int
	DWCHANNELS                 int
	DWRESERVED1                int
	DWRESERVED2                int
	DWSUBENTRIES               int
	DWDIALMODE                 int
	DWDIALEXTRAPERCENT         int
	DWDIALEXTRASAMPLESECONDS   int
	DWHANGUPEXTRAPERCENT       int
	DWHANGUPEXTRASAMPLESECONDS int
	DWIDLEDISCONNECTSECONDS    int
	DWTYPE                     int
	DWENCRYPTIONTYPE           int
	DWCUSTOMAUTHKEY            int
	GUIDID                     int
	SZCUSTOMDIALDLL            int
	DWVPNSTRATEGY              int
	DWFOPTIONS2                int
	DWFOPTIONS3                int
	SZDNSSUFFIX                int
	DWTCPWINDOWSIZE            int
	SZPREREQUISITEPBK          int
	SZPREREQUISITEENTRY        int
	DWREDIALCOUNT              int
	DWREDIALPAUSE              int
	IPV6ADDRDNS                int
	IPV6ADDRDNSALT             int
	DWIPV4INTERFACEMETRIC      int
	DWIPV6INTERFACEMETRIC      int
	IPV6ADDR                   int
	DWIPV6PREFIXLENGTH         int
	DWNETWORKOUTAGETIME        int
	SZIDI                      int
	SZIDR                      int
	FISIMSCONFIG               int
	IdiType                    int
	IdrType                    int
	FDISABLEIKEV2FRAGMENTATION int
}

type TAGRASENTRYW struct {
	DWSIZE                     int
	DWFOPTIONS                 int
	DWCOUNTRYID                int
	DWCOUNTRYCODE              int
	SZAREACODE                 int
	SZLOCALPHONENUMBER         int
	DWALTERNATEOFFSET          int
	IPADDR                     int
	IPADDRDNS                  int
	IPADDRDNSALT               int
	IPADDRWINS                 int
	IPADDRWINSALT              int
	DWFRAMESIZE                int
	DWFNETPROTOCOLS            int
	DWFRAMINGPROTOCOL          int
	SZSCRIPT                   int
	SZAUTODIALDLL              int
	SZAUTODIALFUNC             int
	SZDEVICETYPE               int
	SZDEVICENAME               int
	SZX25PADTYPE               int
	SZX25ADDRESS               int
	SZX25FACILITIES            int
	SZX25USERDATA              int
	DWCHANNELS                 int
	DWRESERVED1                int
	DWRESERVED2                int
	DWSUBENTRIES               int
	DWDIALMODE                 int
	DWDIALEXTRAPERCENT         int
	DWDIALEXTRASAMPLESECONDS   int
	DWHANGUPEXTRAPERCENT       int
	DWHANGUPEXTRASAMPLESECONDS int
	DWIDLEDISCONNECTSECONDS    int
	DWTYPE                     int
	DWENCRYPTIONTYPE           int
	DWCUSTOMAUTHKEY            int
	GUIDID                     int
	SZCUSTOMDIALDLL            int
	DWVPNSTRATEGY              int
	DWFOPTIONS2                int
	DWFOPTIONS3                int
	SZDNSSUFFIX                int
	DWTCPWINDOWSIZE            int
	SZPREREQUISITEPBK          int
	SZPREREQUISITEENTRY        int
	DWREDIALCOUNT              int
	DWREDIALPAUSE              int
	IPV6ADDRDNS                int
	IPV6ADDRDNSALT             int
	DWIPV4INTERFACEMETRIC      int
	DWIPV6INTERFACEMETRIC      int
	IPV6ADDR                   int
	DWIPV6PREFIXLENGTH         int
	DWNETWORKOUTAGETIME        int
	SZIDI                      int
	SZIDR                      int
	FISIMSCONFIG               int
	IdiType                    int
	IdrType                    int
	FDISABLEIKEV2FRAGMENTATION int
}

type TAGRASADPARAMS struct {
	DWSIZE    int
	HWNDOWNER int
	DWFLAGS   int
	XDLG      int
	YDLG      int
}

type TAGRASSUBENTRYA struct {
	DWSIZE             int
	DWFFLAGS           int
	SZDEVICETYPE       int
	SZDEVICENAME       int
	SZLOCALPHONENUMBER int
	DWALTERNATEOFFSET  int
}

type TAGRASSUBENTRYW struct {
	DWSIZE             int
	DWFFLAGS           int
	SZDEVICETYPE       int
	SZDEVICENAME       int
	SZLOCALPHONENUMBER int
	DWALTERNATEOFFSET  int
}

type TAGRASCREDENTIALSA struct {
	DWSIZE     int
	DWMASK     int
	SZUSERNAME int
	SZPASSWORD int
	SZDOMAIN   int
}

type TAGRASCREDENTIALSW struct {
	DWSIZE     int
	DWMASK     int
	SZUSERNAME int
	SZPASSWORD int
	SZDOMAIN   int
}

type TAGRASAUTODIALENTRYA struct {
	DWSIZE            int
	DWFLAGS           int
	DWDIALINGLOCATION int
	SZENTRY           int
}

type TAGRASAUTODIALENTRYW struct {
	DWSIZE            int
	DWFLAGS           int
	DWDIALINGLOCATION int
	SZENTRY           int
}

type TAGRASEAPUSERIDENTITYA struct {
	SZUSERNAME      int
	DWSIZEOFEAPINFO int
	PBEAPINFO       int
}

type TAGRASEAPUSERIDENTITYW struct {
	SZUSERNAME      int
	DWSIZEOFEAPINFO int
	PBEAPINFO       int
}

type TAGRASCOMMSETTINGS struct {
	DWSIZE    int
	BPARITY   int
	BSTOP     int
	BBYTESIZE int
	BALIGN    int
}

type TAGRASCUSTOMSCRIPTEXTENSIONS struct {
	DWSIZE                int
	PFNRASSETCOMMSETTINGS int
}

type RAS_STATS struct {
	DWSIZE                int
	DWBYTESXMITED         int
	DWBYTESRCVED          int
	DWFRAMESXMITED        int
	DWFRAMESRCVED         int
	DWCRCERR              int
	DWTIMEOUTERR          int
	DWALIGNMENTERR        int
	DWHARDWAREOVERRUNERR  int
	DWFRAMINGERR          int
	DWBUFFEROVERRUNERR    int
	DWCOMPRESSIONRATIOIN  int
	DWCOMPRESSIONRATIOOUT int
	DWBPS                 int
	DWCONNECTDURATION     int
}

type TAGRASUPDATECONN struct {
	VERSION        int
	DWSIZE         int
	DWFLAGS        int
	DWIFINDEX      int
	LOCALENDPOINT  int
	REMOTEENDPOINT int
}

type TAGRASNOUSERW struct {
	DWSIZE      int
	DWFLAGS     int
	DWTIMEOUTMS int
	SZUSERNAME  int
	SZPASSWORD  int
	SZDOMAIN    int
}

type TAGRASNOUSERA struct {
	DWSIZE      int
	DWFLAGS     int
	DWTIMEOUTMS int
	SZUSERNAME  int
	SZPASSWORD  int
	SZDOMAIN    int
}

type TAGRASPBDLGW struct {
	DWSIZE       int
	HWNDOWNER    int
	DWFLAGS      int
	XDLG         int
	YDLG         int
	DWCALLBACKID int
	PCALLBACK    int
	DWERROR      int
	RESERVED     int
	RESERVED2    int
}

type TAGRASPBDLGA struct {
	DWSIZE       int
	HWNDOWNER    int
	DWFLAGS      int
	XDLG         int
	YDLG         int
	DWCALLBACKID int
	PCALLBACK    int
	DWERROR      int
	RESERVED     int
	RESERVED2    int
}

type TAGRASENTRYDLGW struct {
	DWSIZE    int
	HWNDOWNER int
	DWFLAGS   int
	XDLG      int
	YDLG      int
	SZENTRY   int
	DWERROR   int
	RESERVED  int
	RESERVED2 int
}

type TAGRASENTRYDLGA struct {
	DWSIZE    int
	HWNDOWNER int
	DWFLAGS   int
	XDLG      int
	YDLG      int
	SZENTRY   int
	DWERROR   int
	RESERVED  int
	RESERVED2 int
}

type TAGRASDIALDLG struct {
	DWSIZE     int
	HWNDOWNER  int
	DWFLAGS    int
	XDLG       int
	YDLG       int
	DWSUBENTRY int
	DWERROR    int
	RESERVED   int
	RESERVED2  int
}

type MPR_INTERFACE_0 struct {
	WSZINTERFACENAME       int
	HINTERFACE             int
	FENABLED               int
	DWIFTYPE               int
	DWCONNECTIONSTATE      int
	FUNREACHABILITYREASONS int
	DWLASTERROR            int
}

type MPR_IPINIP_INTERFACE_0 struct {
	WSZFRIENDLYNAME int
	Guid            int
}

type MPR_INTERFACE_1 struct {
	WSZINTERFACENAME            int
	HINTERFACE                  int
	FENABLED                    int
	DWIFTYPE                    int
	DWCONNECTIONSTATE           int
	FUNREACHABILITYREASONS      int
	DWLASTERROR                 int
	LPWSDIALOUTHOURSRESTRICTION int
}

type MPR_INTERFACE_2 struct {
	WSZINTERFACENAME           int
	HINTERFACE                 int
	FENABLED                   int
	DWIFTYPE                   int
	DWCONNECTIONSTATE          int
	FUNREACHABILITYREASONS     int
	DWLASTERROR                int
	DWFOPTIONS                 int
	SZLOCALPHONENUMBER         int
	SZALTERNATES               int
	IPADDR                     int
	IPADDRDNS                  int
	IPADDRDNSALT               int
	IPADDRWINS                 int
	IPADDRWINSALT              int
	DWFNETPROTOCOLS            int
	SZDEVICETYPE               int
	SZDEVICENAME               int
	SZX25PADTYPE               int
	SZX25ADDRESS               int
	SZX25FACILITIES            int
	SZX25USERDATA              int
	DWCHANNELS                 int
	DWSUBENTRIES               int
	DWDIALMODE                 int
	DWDIALEXTRAPERCENT         int
	DWDIALEXTRASAMPLESECONDS   int
	DWHANGUPEXTRAPERCENT       int
	DWHANGUPEXTRASAMPLESECONDS int
	DWIDLEDISCONNECTSECONDS    int
	DWTYPE                     int
	DWENCRYPTIONTYPE           int
	DWCUSTOMAUTHKEY            int
	DWCUSTOMAUTHDATASIZE       int
	LPBCUSTOMAUTHDATA          int
	GUIDID                     int
	DWVPNSTRATEGY              int
}

type MPR_INTERFACE_3 struct {
	WSZINTERFACENAME           int
	HINTERFACE                 int
	FENABLED                   int
	DWIFTYPE                   int
	DWCONNECTIONSTATE          int
	FUNREACHABILITYREASONS     int
	DWLASTERROR                int
	DWFOPTIONS                 int
	SZLOCALPHONENUMBER         int
	SZALTERNATES               int
	IPADDR                     int
	IPADDRDNS                  int
	IPADDRDNSALT               int
	IPADDRWINS                 int
	IPADDRWINSALT              int
	DWFNETPROTOCOLS            int
	SZDEVICETYPE               int
	SZDEVICENAME               int
	SZX25PADTYPE               int
	SZX25ADDRESS               int
	SZX25FACILITIES            int
	SZX25USERDATA              int
	DWCHANNELS                 int
	DWSUBENTRIES               int
	DWDIALMODE                 int
	DWDIALEXTRAPERCENT         int
	DWDIALEXTRASAMPLESECONDS   int
	DWHANGUPEXTRAPERCENT       int
	DWHANGUPEXTRASAMPLESECONDS int
	DWIDLEDISCONNECTSECONDS    int
	DWTYPE                     int
	DWENCRYPTIONTYPE           int
	DWCUSTOMAUTHKEY            int
	DWCUSTOMAUTHDATASIZE       int
	LPBCUSTOMAUTHDATA          int
	GUIDID                     int
	DWVPNSTRATEGY              int
	AddressCount               int
	IPV6ADDRDNS                int
	IPV6ADDRDNSALT             int
	IPV6ADDR                   int
}

type MPR_DEVICE_0 struct {
	SZDEVICETYPE int
	SZDEVICENAME int
}

type MPR_DEVICE_1 struct {
	SZDEVICETYPE       int
	SZDEVICENAME       int
	SZLOCALPHONENUMBER int
	SZALTERNATES       int
}

type MPR_CREDENTIALSEX_0 struct {
	DWSIZE             int
	LPBCREDENTIALSINFO int
}

type MPR_CREDENTIALSEX_1 struct {
	DWSIZE             int
	LPBCREDENTIALSINFO int
}

type MPR_TRANSPORT_0 struct {
	DWTRANSPORTID    int
	HTRANSPORT       int
	WSZTRANSPORTNAME int
}

type MPR_IFTRANSPORT_0 struct {
	DWTRANSPORTID      int
	HIFTRANSPORT       int
	WSZIFTRANSPORTNAME int
}

type MPR_SERVER_0 struct {
	FLANONLYMODE int
	DWUPTIME     int
	DWTOTALPORTS int
	DWPORTSINUSE int
}

type MPR_SERVER_1 struct {
	DWNUMPPTPPORTS  int
	DWPPTPPORTFLAGS int
	DWNUML2TPPORTS  int
	DWL2TPPORTFLAGS int
}

type MPR_SERVER_2 struct {
	DWNUMPPTPPORTS  int
	DWPPTPPORTFLAGS int
	DWNUML2TPPORTS  int
	DWL2TPPORTFLAGS int
	DWNUMSSTPPORTS  int
	DWSSTPPORTFLAGS int
}

type RAS_PORT_0 struct {
	HPORT                int
	HCONNECTION          int
	DWPORTCONDITION      int
	DWTOTALNUMBEROFCALLS int
	DWCONNECTDURATION    int
	WSZPORTNAME          int
	WSZMEDIANAME         int
	WSZDEVICENAME        int
	WSZDEVICETYPE        int
}

type RAS_PORT_1 struct {
	HPORT                 int
	HCONNECTION           int
	DWHARDWARECONDITION   int
	DWLINESPEED           int
	DWBYTESXMITED         int
	DWBYTESRCVED          int
	DWFRAMESXMITED        int
	DWFRAMESRCVED         int
	DWCRCERR              int
	DWTIMEOUTERR          int
	DWALIGNMENTERR        int
	DWHARDWAREOVERRUNERR  int
	DWFRAMINGERR          int
	DWBUFFEROVERRUNERR    int
	DWCOMPRESSIONRATIOIN  int
	DWCOMPRESSIONRATIOOUT int
}

type RAS_PORT_2 struct {
	HPORT                   int
	HCONNECTION             int
	DWCONN_STATE            int
	WSZPORTNAME             int
	WSZMEDIANAME            int
	WSZDEVICENAME           int
	WSZDEVICETYPE           int
	DWHARDWARECONDITION     int
	DWLINESPEED             int
	DWCRCERR                int
	DWSERIALOVERRUNERRS     int
	DWTIMEOUTERR            int
	DWALIGNMENTERR          int
	DWHARDWAREOVERRUNERR    int
	DWFRAMINGERR            int
	DWBUFFEROVERRUNERR      int
	DWCOMPRESSIONRATIOIN    int
	DWCOMPRESSIONRATIOOUT   int
	DWTOTALERRORS           int
	ULLBYTESXMITED          int
	ULLBYTESRCVED           int
	ULLFRAMESXMITED         int
	ULLFRAMESRCVED          int
	ULLBYTESTXUNCOMPRESSED  int
	ULLBYTESTXCOMPRESSED    int
	ULLBYTESRCVUNCOMPRESSED int
	ULLBYTESRCVCOMPRESSED   int
}

type PPP_NBFCP_INFO struct {
	DWERROR  int
	WSZWKSTA int
}

type PPP_IPCP_INFO struct {
	DWERROR          int
	WSZADDRESS       int
	WSZREMOTEADDRESS int
}

type PPP_IPCP_INFO2 struct {
	DWERROR          int
	WSZADDRESS       int
	WSZREMOTEADDRESS int
	DWOPTIONS        int
	DWREMOTEOPTIONS  int
}

type PPP_IPXCP_INFO struct {
	DWERROR    int
	WSZADDRESS int
}

type PPP_ATCP_INFO struct {
	DWERROR    int
	WSZADDRESS int
}

type PPP_IPV6_CP_INFO struct {
	DWVERSION                  int
	DWSIZE                     int
	DWERROR                    int
	BINTERFACEIDENTIFIER       int
	BREMOTEINTERFACEIDENTIFIER int
	DWOPTIONS                  int
	DWREMOTEOPTIONS            int
	BPREFIX                    int
	DWPREFIXLENGTH             int
}

type PPP_INFO struct {
	NBF int
	IP  int
	IPX int
	AT  int
}

type PPP_CCP_INFO struct {
	DWERROR                      int
	DWCOMPRESSIONALGORITHM       int
	DWOPTIONS                    int
	DWREMOTECOMPRESSIONALGORITHM int
	DWREMOTEOPTIONS              int
}

type PPP_LCP_INFO struct {
	DWERROR                        int
	DWAUTHENTICATIONPROTOCOL       int
	DWAUTHENTICATIONDATA           int
	DWREMOTEAUTHENTICATIONPROTOCOL int
	DWREMOTEAUTHENTICATIONDATA     int
	DWTERMINATEREASON              int
	DWREMOTETERMINATEREASON        int
	DWOPTIONS                      int
	DWREMOTEOPTIONS                int
	DWEAPTYPEID                    int
	DWREMOTEEAPTYPEID              int
}

type PPP_INFO_2 struct {
	NBF int
	IP  int
	IPX int
	AT  int
	CCP int
	LCP int
}

type PPP_INFO_3 struct {
	NBF  int
	IP   int
	IPV6 int
	CCP  int
	LCP  int
}

type RAS_CONNECTION_0 struct {
	HCONNECTION       int
	HINTERFACE        int
	DWCONNECTDURATION int
	DWINTERFACETYPE   int
	DWCONNECTIONFLAGS int
	WSZINTERFACENAME  int
	WSZUSERNAME       int
	WSZLOGONDOMAIN    int
	WSZREMOTECOMPUTER int
}

type RAS_CONNECTION_1 struct {
	HCONNECTION           int
	HINTERFACE            int
	PppInfo               int
	DWBYTESXMITED         int
	DWBYTESRCVED          int
	DWFRAMESXMITED        int
	DWFRAMESRCVED         int
	DWCRCERR              int
	DWTIMEOUTERR          int
	DWALIGNMENTERR        int
	DWHARDWAREOVERRUNERR  int
	DWFRAMINGERR          int
	DWBUFFEROVERRUNERR    int
	DWCOMPRESSIONRATIOIN  int
	DWCOMPRESSIONRATIOOUT int
}

type RAS_CONNECTION_2 struct {
	HCONNECTION     int
	WSZUSERNAME     int
	DWINTERFACETYPE int
	GUID            int
	PppInfo2        int
}

type RAS_CONNECTION_3 struct {
	DWVERSION       int
	DWSIZE          int
	HCONNECTION     int
	WSZUSERNAME     int
	DWINTERFACETYPE int
	GUID            int
	PppInfo3        int
	RASQUARSTATE    int
	TIMER           int
}

type RAS_USER_0 struct {
	BFPRIVILEGE    int
	WSZPHONENUMBER int
}

type RAS_USER_1 struct {
	BFPRIVILEGE    int
	WSZPHONENUMBER int
	BFPRIVILEGE2   int
}

type MPR_FILTER_0 struct {
	FENABLE int
}

type MPRAPI_OBJECT_HEADER struct {
	REVISION int
	TYPE     int
	SIZE     int
}

type PPP_PROJECTION_INFO struct {
	DWIPV4NEGOTIATIONERROR         int
	WSZADDRESS                     int
	WSZREMOTEADDRESS               int
	DWIPV4OPTIONS                  int
	DWIPV4REMOTEOPTIONS            int
	IPv4SubInterfaceIndex          int
	DWIPV6NEGOTIATIONERROR         int
	BINTERFACEIDENTIFIER           int
	BREMOTEINTERFACEIDENTIFIER     int
	BPREFIX                        int
	DWPREFIXLENGTH                 int
	IPv6SubInterfaceIndex          int
	DWLCPERROR                     int
	DWAUTHENTICATIONPROTOCOL       int
	DWAUTHENTICATIONDATA           int
	DWREMOTEAUTHENTICATIONPROTOCOL int
	DWREMOTEAUTHENTICATIONDATA     int
	DWLCPTERMINATEREASON           int
	DWLCPREMOTETERMINATEREASON     int
	DWLCPOPTIONS                   int
	DWLCPREMOTEOPTIONS             int
	DWEAPTYPEID                    int
	DWREMOTEEAPTYPEID              int
	DWCCPERROR                     int
	DWCOMPRESSIONALGORITHM         int
	DWCCPOPTIONS                   int
	DWREMOTECOMPRESSIONALGORITHM   int
	DWCCPREMOTEOPTIONS             int
}

type PPP_PROJECTION_INFO2 struct {
	DWIPV4NEGOTIATIONERROR         int
	WSZADDRESS                     int
	WSZREMOTEADDRESS               int
	DWIPV4OPTIONS                  int
	DWIPV4REMOTEOPTIONS            int
	IPv4SubInterfaceIndex          int
	DWIPV6NEGOTIATIONERROR         int
	BINTERFACEIDENTIFIER           int
	BREMOTEINTERFACEIDENTIFIER     int
	BPREFIX                        int
	DWPREFIXLENGTH                 int
	IPv6SubInterfaceIndex          int
	DWLCPERROR                     int
	DWAUTHENTICATIONPROTOCOL       int
	DWAUTHENTICATIONDATA           int
	DWREMOTEAUTHENTICATIONPROTOCOL int
	DWREMOTEAUTHENTICATIONDATA     int
	DWLCPTERMINATEREASON           int
	DWLCPREMOTETERMINATEREASON     int
	DWLCPOPTIONS                   int
	DWLCPREMOTEOPTIONS             int
	DWEAPTYPEID                    int
	DWEMBEDDEDEAPTYPEID            int
	DWREMOTEEAPTYPEID              int
	DWCCPERROR                     int
	DWCOMPRESSIONALGORITHM         int
	DWCCPOPTIONS                   int
	DWREMOTECOMPRESSIONALGORITHM   int
	DWCCPREMOTEOPTIONS             int
}

type IKEV2_PROJECTION_INFO struct {
	DWIPV4NEGOTIATIONERROR     int
	WSZADDRESS                 int
	WSZREMOTEADDRESS           int
	IPv4SubInterfaceIndex      int
	DWIPV6NEGOTIATIONERROR     int
	BINTERFACEIDENTIFIER       int
	BREMOTEINTERFACEIDENTIFIER int
	BPREFIX                    int
	DWPREFIXLENGTH             int
	IPv6SubInterfaceIndex      int
	DWOPTIONS                  int
	DWAUTHENTICATIONPROTOCOL   int
	DWEAPTYPEID                int
	DWCOMPRESSIONALGORITHM     int
	DWENCRYPTIONMETHOD         int
}

type IKEV2_PROJECTION_INFO2 struct {
	DWIPV4NEGOTIATIONERROR     int
	WSZADDRESS                 int
	WSZREMOTEADDRESS           int
	IPv4SubInterfaceIndex      int
	DWIPV6NEGOTIATIONERROR     int
	BINTERFACEIDENTIFIER       int
	BREMOTEINTERFACEIDENTIFIER int
	BPREFIX                    int
	DWPREFIXLENGTH             int
	IPv6SubInterfaceIndex      int
	DWOPTIONS                  int
	DWAUTHENTICATIONPROTOCOL   int
	DWEAPTYPEID                int
	DWEMBEDDEDEAPTYPEID        int
	DWCOMPRESSIONALGORITHM     int
	DWENCRYPTIONMETHOD         int
}

type PROJECTION_INFO struct {
	PROJECTIONINFOTYPE int
	Anonymous          int
}

type PROJECTION_INFO2 struct {
	PROJECTIONINFOTYPE int
	Anonymous          int
}

type RAS_CONNECTION_EX struct {
	Header                   int
	DWCONNECTDURATION        int
	DWINTERFACETYPE          int
	DWCONNECTIONFLAGS        int
	WSZINTERFACENAME         int
	WSZUSERNAME              int
	WSZLOGONDOMAIN           int
	WSZREMOTECOMPUTER        int
	GUID                     int
	RASQUARSTATE             int
	PROBATIONTIME            int
	DWBYTESXMITED            int
	DWBYTESRCVED             int
	DWFRAMESXMITED           int
	DWFRAMESRCVED            int
	DWCRCERR                 int
	DWTIMEOUTERR             int
	DWALIGNMENTERR           int
	DWHARDWAREOVERRUNERR     int
	DWFRAMINGERR             int
	DWBUFFEROVERRUNERR       int
	DWCOMPRESSIONRATIOIN     int
	DWCOMPRESSIONRATIOOUT    int
	DWNUMSWITCHOVERS         int
	WSZREMOTEENDPOINTADDRESS int
	WSZLOCALENDPOINTADDRESS  int
	ProjectionInfo           int
	HCONNECTION              int
	HINTERFACE               int
}

type RAS_CONNECTION_4 struct {
	DWCONNECTDURATION        int
	DWINTERFACETYPE          int
	DWCONNECTIONFLAGS        int
	WSZINTERFACENAME         int
	WSZUSERNAME              int
	WSZLOGONDOMAIN           int
	WSZREMOTECOMPUTER        int
	GUID                     int
	RASQUARSTATE             int
	PROBATIONTIME            int
	CONNECTIONSTARTTIME      int
	ULLBYTESXMITED           int
	ULLBYTESRCVED            int
	DWFRAMESXMITED           int
	DWFRAMESRCVED            int
	DWCRCERR                 int
	DWTIMEOUTERR             int
	DWALIGNMENTERR           int
	DWHARDWAREOVERRUNERR     int
	DWFRAMINGERR             int
	DWBUFFEROVERRUNERR       int
	DWCOMPRESSIONRATIOIN     int
	DWCOMPRESSIONRATIOOUT    int
	DWNUMSWITCHOVERS         int
	WSZREMOTEENDPOINTADDRESS int
	WSZLOCALENDPOINTADDRESS  int
	ProjectionInfo           int
	HCONNECTION              int
	HINTERFACE               int
	DWDEVICETYPE             int
}

type ROUTER_CUSTOM_IKEv2_POLICY0 struct {
	DWINTEGRITYMETHOD         int
	DWENCRYPTIONMETHOD        int
	DWCIPHERTRANSFORMCONSTANT int
	DWAUTHTRANSFORMCONSTANT   int
	DWPFSGROUP                int
	DWDHGROUP                 int
}

type ROUTER_IKEv2_IF_CUSTOM_CONFIG0 struct {
	DWSALIFETIME    int
	DWSADATASIZE    int
	CERTIFICATENAME int
	CUSTOMPOLICY    int
}

type MPR_IF_CUSTOMINFOEX0 struct {
	Header            int
	DWFLAGS           int
	CUSTOMIKEV2CONFIG int
}

type MPR_CERT_EKU struct {
	DWSIZE   int
	IsEKUOID int
	PWSZEKU  int
}

type VPN_TS_IP_ADDRESS struct {
	Type      int
	Anonymous int
}

type MPR_VPN_SELECTOR struct {
	TYPE        int
	PROTOCOLID  int
	PORTSTART   int
	PORTEND     int
	TSPAYLOADID int
	ADDRSTART   int
	ADDREND     int
}

type MPR_VPN_TRAFFIC_SELECTORS struct {
	NUMTSI int
	NUMTSR int
	TSI    int
	TSR    int
}

type ROUTER_IKEv2_IF_CUSTOM_CONFIG2 struct {
	DWSALIFETIME        int
	DWSADATASIZE        int
	CERTIFICATENAME     int
	CUSTOMPOLICY        int
	CERTIFICATEHASH     int
	DWMMSALIFETIME      int
	VPNTRAFFICSELECTORS int
}

type MPR_IF_CUSTOMINFOEX2 struct {
	Header            int
	DWFLAGS           int
	CUSTOMIKEV2CONFIG int
}

type IKEV2_TUNNEL_CONFIG_PARAMS4 struct {
	DWIDLETIMEOUT                int
	DWNETWORKBLACKOUTTIME        int
	DWSALIFETIME                 int
	DWSADATASIZEFORRENEGOTIATION int
	DWCONFIGOPTIONS              int
	DWTOTALCERTIFICATES          int
	CERTIFICATENAMES             int
	MACHINECERTIFICATENAME       int
	DWENCRYPTIONTYPE             int
	CUSTOMPOLICY                 int
	DWTOTALEKUS                  int
	CERTIFICATEEKUS              int
	MACHINECERTIFICATEHASH       int
	DWMMSALIFETIME               int
}

type ROUTER_IKEv2_IF_CUSTOM_CONFIG1 struct {
	DWSALIFETIME    int
	DWSADATASIZE    int
	CERTIFICATENAME int
	CUSTOMPOLICY    int
	CERTIFICATEHASH int
}

type MPR_IF_CUSTOMINFOEX1 struct {
	Header            int
	DWFLAGS           int
	CUSTOMIKEV2CONFIG int
}

type IKEV2_TUNNEL_CONFIG_PARAMS3 struct {
	DWIDLETIMEOUT                int
	DWNETWORKBLACKOUTTIME        int
	DWSALIFETIME                 int
	DWSADATASIZEFORRENEGOTIATION int
	DWCONFIGOPTIONS              int
	DWTOTALCERTIFICATES          int
	CERTIFICATENAMES             int
	MACHINECERTIFICATENAME       int
	DWENCRYPTIONTYPE             int
	CUSTOMPOLICY                 int
	DWTOTALEKUS                  int
	CERTIFICATEEKUS              int
	MACHINECERTIFICATEHASH       int
}

type IKEV2_TUNNEL_CONFIG_PARAMS2 struct {
	DWIDLETIMEOUT                int
	DWNETWORKBLACKOUTTIME        int
	DWSALIFETIME                 int
	DWSADATASIZEFORRENEGOTIATION int
	DWCONFIGOPTIONS              int
	DWTOTALCERTIFICATES          int
	CERTIFICATENAMES             int
	MACHINECERTIFICATENAME       int
	DWENCRYPTIONTYPE             int
	CUSTOMPOLICY                 int
}

type L2TP_TUNNEL_CONFIG_PARAMS2 struct {
	DWIDLETIMEOUT                int
	DWENCRYPTIONTYPE             int
	DWSALIFETIME                 int
	DWSADATASIZEFORRENEGOTIATION int
	CUSTOMPOLICY                 int
	DWMMSALIFETIME               int
}

type L2TP_TUNNEL_CONFIG_PARAMS1 struct {
	DWIDLETIMEOUT                int
	DWENCRYPTIONTYPE             int
	DWSALIFETIME                 int
	DWSADATASIZEFORRENEGOTIATION int
	CUSTOMPOLICY                 int
}

type IKEV2_CONFIG_PARAMS struct {
	DWNUMPORTS               int
	DWPORTFLAGS              int
	DWTUNNELCONFIGPARAMFLAGS int
	TunnelConfigParams       int
}

type PPTP_CONFIG_PARAMS struct {
	DWNUMPORTS  int
	DWPORTFLAGS int
}

type L2TP_CONFIG_PARAMS1 struct {
	DWNUMPORTS               int
	DWPORTFLAGS              int
	DWTUNNELCONFIGPARAMFLAGS int
	TunnelConfigParams       int
}

type GRE_CONFIG_PARAMS0 struct {
	DWNUMPORTS  int
	DWPORTFLAGS int
}

type L2TP_CONFIG_PARAMS0 struct {
	DWNUMPORTS  int
	DWPORTFLAGS int
}

type SSTP_CERT_INFO struct {
	ISDEFAULT int
	CERTBLOB  int
}

type SSTP_CONFIG_PARAMS struct {
	DWNUMPORTS      int
	DWPORTFLAGS     int
	ISUSEHTTPS      int
	CERTALGORITHM   int
	SSTPCERTDETAILS int
}

type MPRAPI_TUNNEL_CONFIG_PARAMS0 struct {
	IkeConfigParams  int
	PptpConfigParams int
	L2tpConfigParams int
	SstpConfigParams int
}

type MPRAPI_TUNNEL_CONFIG_PARAMS1 struct {
	IkeConfigParams  int
	PptpConfigParams int
	L2tpConfigParams int
	SstpConfigParams int
	GREConfigParams  int
}

type MPR_SERVER_EX0 struct {
	Header       int
	FLANONLYMODE int
	DWUPTIME     int
	DWTOTALPORTS int
	DWPORTSINUSE int
	Reserved     int
	ConfigParams int
}

type MPR_SERVER_EX1 struct {
	Header       int
	FLANONLYMODE int
	DWUPTIME     int
	DWTOTALPORTS int
	DWPORTSINUSE int
	Reserved     int
	ConfigParams int
}

type MPR_SERVER_SET_CONFIG_EX0 struct {
	Header                int
	SETCONFIGFORPROTOCOLS int
	ConfigParams          int
}

type MPR_SERVER_SET_CONFIG_EX1 struct {
	Header                int
	SETCONFIGFORPROTOCOLS int
	ConfigParams          int
}

type AUTH_VALIDATION_EX struct {
	Header         int
	HRASCONNECTION int
	WSZUSERNAME    int
	WSZLOGONDOMAIN int
	AuthInfoSize   int
	AuthInfo       int
}

type RAS_UPDATE_CONNECTION struct {
	Header                   int
	DWIFINDEX                int
	WSZLOCALENDPOINTADDRESS  int
	WSZREMOTEENDPOINTADDRESS int
}

type MPRAPI_ADMIN_DLL_CALLBACKS struct {
	REVISION                                    int
	LPFNMPRADMINGETIPADDRESSFORUSER             int
	LPFNMPRADMINRELEASEIPADDRESS                int
	LPFNMPRADMINGETIPV6ADDRESSFORUSER           int
	LPFNMPRADMINRELEASEIPV6ADDRESSFORUSER       int
	LPFNRASADMINACCEPTNEWLINK                   int
	LPFNRASADMINLINKHANGUPNOTIFICATION          int
	LPFNRASADMINTERMINATEDLL                    int
	LPFNRASADMINACCEPTNEWCONNECTIONEX           int
	LPFNRASADMINACCEPTENDPOINTCHANGEEX          int
	LPFNRASADMINACCEPTREAUTHENTICATIONEX        int
	LPFNRASADMINCONNECTIONHANGUPNOTIFICATIONEX  int
	LPFNRASVALIDATEPREAUTHENTICATEDCONNECTIONEX int
}

type SECURITY_MESSAGE struct {
	DWMSGID  int
	HPORT    int
	DWERROR  int
	UserName int
	Domain   int
}

type RAS_SECURITY_INFO struct {
	LastError     int
	BytesReceived int
	DeviceName    int
}

type MGM_IF_ENTRY struct {
	DWIFINDEX       int
	DWIFNEXTHOPADDR int
	BIGMP           int
	BISENABLED      int
}

type ROUTING_PROTOCOL_CONFIG struct {
	DWCALLBACKFLAGS          int
	PFNRPFCALLBACK           int
	PFNCREATIONALERTCALLBACK int
	PFNPRUNEALERTCALLBACK    int
	PFNJOINALERTCALLBACK     int
	PFNWRONGIFCALLBACK       int
	PFNLOCALJOINCALLBACK     int
	PFNLOCALLEAVECALLBACK    int
	PFNDISABLEIGMPCALLBACK   int
	PFNENABLEIGMPCALLBACK    int
}

type SOURCE_GROUP_ENTRY struct {
	DWSOURCEADDR int
	DWSOURCEMASK int
	DWGROUPADDR  int
	DWGROUPMASK  int
}

type RTM_REGN_PROFILE struct {
	MaxNextHopsInRoute int
	MaxHandlesInEnum   int
	ViewsSupported     int
	NumberOfViews      int
}

type RTM_NET_ADDRESS struct {
	AddressFamily int
	NumBits       int
	AddrBits      int
}

type RTM_PREF_INFO struct {
	Metric     int
	Preference int
}

type RTM_NEXTHOP_LIST struct {
	NumNextHops int
	NextHops    int
}

type RTM_DEST_INFO struct {
	DestHandle     int
	DestAddress    int
	LastChanged    int
	BelongsToViews int
	NumberOfViews  int
	ViewInfo       int
}

type RTM_ROUTE_INFO struct {
	DestHandle         int
	RouteOwner         int
	Neighbour          int
	State              int
	Flags1             int
	Flags              int
	PrefInfo           int
	BelongsToViews     int
	EntitySpecificInfo int
	NextHopsList       int
}

type RTM_NEXTHOP_INFO struct {
	NextHopAddress     int
	NextHopOwner       int
	InterfaceIndex     int
	State              int
	Flags              int
	EntitySpecificInfo int
	RemoteNextHop      int
}

type RTM_ENTITY_ID struct {
	Anonymous int
}

type RTM_ENTITY_INFO struct {
	RtmInstanceId int
	AddressFamily int
	EntityId      int
}

type RTM_ENTITY_METHOD_INPUT struct {
	MethodType int
	InputSize  int
	InputData  int
}

type RTM_ENTITY_METHOD_OUTPUT struct {
	MethodType   int
	MethodStatus int
	OutputSize   int
	OutputData   int
}

type RTM_ENTITY_EXPORT_METHODS struct {
	NumMethods int
	Methods    int
}