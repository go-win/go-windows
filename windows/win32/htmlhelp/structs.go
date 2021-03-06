// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package htmlhelp implements the Windows.Win32.HtmlHelp namespace.
package htmlhelp

type HHN_NOTIFY struct {
	HDR    int
	PSZURL int
}

type HH_POPUP struct {
	CBSTRUCT      int
	HINST         int
	IDSTRING      int
	PSZTEXT       int
	PT            int
	CLRFOREGROUND int
	CLRBACKGROUND int
	RCMARGINS     int
	PSZFONT       int
}

type HH_AKLINK struct {
	CBSTRUCT     int
	FRESERVED    int
	PSZKEYWORDS  int
	PSZURL       int
	PSZMSGTEXT   int
	PSZMSGTITLE  int
	PSZWINDOW    int
	FINDEXONFAIL int
}

type HH_ENUM_IT struct {
	CBSTRUCT         int
	ITYPE            int
	PSZCATNAME       int
	PSZITNAME        int
	PSZITDESCRIPTION int
}

type HH_ENUM_CAT struct {
	CBSTRUCT          int
	PSZCATNAME        int
	PSZCATDESCRIPTION int
}

type HH_SET_INFOTYPE struct {
	CBSTRUCT        int
	PSZCATNAME      int
	PSZINFOTYPENAME int
}

type HH_FTS_QUERY struct {
	CBSTRUCT        int
	FUNICODESTRINGS int
	PSZSEARCHQUERY  int
	IPROXIMITY      int
	FSTEMMEDSEARCH  int
	FTITLEONLY      int
	FEXECUTE        int
	PSZWINDOW       int
}

type HH_WINTYPE struct {
	CBSTRUCT        int
	FUNICODESTRINGS int
	PSZTYPE         int
	FSVALIDMEMBERS  int
	FSWINPROPERTIES int
	PSZCAPTION      int
	DWSTYLES        int
	DWEXSTYLES      int
	RCWINDOWPOS     int
	NSHOWSTATE      int
	HWNDHELP        int
	HWNDCALLER      int
	PAINFOTYPES     int
	HWNDTOOLBAR     int
	HWNDNAVIGATION  int
	HWNDHTML        int
	INAVWIDTH       int
	RCHTML          int
	PSZTOC          int
	PSZINDEX        int
	PSZFILE         int
	PSZHOME         int
	FSTOOLBARFLAGS  int
	FNOTEXPANDED    int
	CURNAVTYPE      int
	TABPOS          int
	IDNOTIFY        int
	TABORDER        int
	CHISTORY        int
	PSZJUMP1        int
	PSZJUMP2        int
	PSZURLJUMP1     int
	PSZURLJUMP2     int
	RCMINSIZE       int
	CBINFOTYPES     int
	PSZCUSTOMTABS   int
}

type HHNTRACK struct {
	HDR        int
	PSZCURURL  int
	IDACTION   int
	PHHWINTYPE int
}

type HH_GLOBAL_PROPERTY struct {
	ID  int
	VAR int
}

type CProperty struct {
	DWPROPID  int
	CBDATA    int
	DWTYPE    int
	Anonymous int
	FPERSIST  int
}

type IITGroup struct {
}

type IITQuery struct {
}

type IITStopWordList struct {
}

type ROWSTATUS struct {
	LROWFIRST   int
	CROWS       int
	CPROPERTIES int
	CROWSTOTAL  int
}

type COLUMNSTATUS struct {
	CPROPCOUNT   int
	CPROPSLOADED int
}
