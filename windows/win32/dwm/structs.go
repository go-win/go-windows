// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package dwm implements the Windows.Win32.Dwm namespace.
package dwm

type DWM_BLURBEHIND struct {
	DWFLAGS                int
	FENABLE                int
	HRGNBLUR               int
	FTRANSITIONONMAXIMIZED int
}

type DWM_THUMBNAIL_PROPERTIES struct {
	DWFLAGS               int
	RCDESTINATION         int
	RCSOURCE              int
	OPACITY               int
	FVISIBLE              int
	FSOURCECLIENTAREAONLY int
}

type UNSIGNED_RATIO struct {
	UINUMERATOR   int
	UIDENOMINATOR int
}

type DWM_TIMING_INFO struct {
	CBSIZE                 int
	RATEREFRESH            int
	QPCREFRESHPERIOD       int
	RATECOMPOSE            int
	QPCVBLANK              int
	CREFRESH               int
	CDXREFRESH             int
	QPCCOMPOSE             int
	CFRAME                 int
	CDXPRESENT             int
	CREFRESHFRAME          int
	CFRAMESUBMITTED        int
	CDXPRESENTSUBMITTED    int
	CFRAMECONFIRMED        int
	CDXPRESENTCONFIRMED    int
	CREFRESHCONFIRMED      int
	CDXREFRESHCONFIRMED    int
	CFRAMESLATE            int
	CFRAMESOUTSTANDING     int
	CFRAMEDISPLAYED        int
	QPCFRAMEDISPLAYED      int
	CREFRESHFRAMEDISPLAYED int
	CFRAMECOMPLETE         int
	QPCFRAMECOMPLETE       int
	CFRAMEPENDING          int
	QPCFRAMEPENDING        int
	CFRAMESDISPLAYED       int
	CFRAMESCOMPLETE        int
	CFRAMESPENDING         int
	CFRAMESAVAILABLE       int
	CFRAMESDROPPED         int
	CFRAMESMISSED          int
	CREFRESHNEXTDISPLAYED  int
	CREFRESHNEXTPRESENTED  int
	CREFRESHESDISPLAYED    int
	CREFRESHESPRESENTED    int
	CREFRESHSTARTED        int
	CPIXELSRECEIVED        int
	CPIXELSDRAWN           int
	CBUFFERSEMPTY          int
}

type DWM_PRESENT_PARAMETERS struct {
	CBSIZE             int
	FQUEUE             int
	CREFRESHSTART      int
	CBUFFER            int
	FUSESOURCERATE     int
	RATESOURCE         int
	CREFRESHESPERFRAME int
	ESAMPLING          int
}