// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package rightsmanagement implements the Windows.Win32.RightsManagement namespace.
package rightsmanagement

type DRMID struct {
	UVERSION  int
	WSZIDTYPE int
	WSZID     int
}

type DRMBOUNDLICENSEPARAMS struct {
	UVERSION                               int
	HENABLINGPRINCIPAL                     int
	HSECURESTORE                           int
	WSZRIGHTSREQUESTED                     int
	WSZRIGHTSGROUP                         int
	IDRESOURCE                             int
	CAUTHENTICATORCOUNT                    int
	RGHAUTHENTICATORS                      int
	WSZDEFAULTENABLINGPRINCIPALCREDENTIALS int
	DWFLAGS                                int
}

type DRM_LICENSE_ACQ_DATA struct {
	UVERSION         int
	WSZURL           int
	WSZLOCALFILENAME int
	PBPOSTDATA       int
	DWPOSTDATASIZE   int
	WSZFRIENDLYNAME  int
}

type DRM_ACTSERV_INFO struct {
	UVERSION  int
	WSZPUBKEY int
	WSZURL    int
}

type DRM_CLIENT_VERSION_INFO struct {
	USTRUCTVERSION        int
	DWVERSION             int
	WSZHIERARCHY          int
	WSZPRODUCTID          int
	WSZPRODUCTDESCRIPTION int
}
