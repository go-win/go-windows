// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsaddressbook implements the Windows.Win32.WindowsAddressBook namespace.
package windowsaddressbook

type ENTRYID struct {
	ABFLAGS int
	AB      int
}

type MAPIUID struct {
	AB int
}

type SPropTagArray struct {
	CVALUES    int
	AULPROPTAG int
}

type SBinary struct {
	CB  int
	LPB int
}

type SShortArray struct {
	CVALUES int
	LPI     int
}

type SGuidArray struct {
	CVALUES int
	LPGUID  int
}

type SRealArray struct {
	CVALUES int
	LPFLT   int
}

type SLongArray struct {
	CVALUES int
	LPL     int
}

type SLargeIntegerArray struct {
	CVALUES int
	LPLI    int
}

type SDateTimeArray struct {
	CVALUES int
	LPFT    int
}

type SAppTimeArray struct {
	CVALUES int
	LPAT    int
}

type SCurrencyArray struct {
	CVALUES int
	LPCUR   int
}

type SBinaryArray struct {
	CVALUES int
	LPBIN   int
}

type SDoubleArray struct {
	CVALUES int
	LPDBL   int
}

type SWStringArray struct {
	CVALUES int
	LPPSZW  int
}

type SLPSTRArray struct {
	CVALUES int
	LPPSZA  int
}

type PV struct {
	I      int
	L      int
	UL     int
	FLT    int
	DBL    int
	B      int
	CUR    int
	AT     int
	FT     int
	LPSZA  int
	BIN    int
	LPSZW  int
	LPGUID int
	LI     int
	MVi    int
	MVl    int
	MVflt  int
	MVdbl  int
	MVcur  int
	MVat   int
	MVft   int
	MVbin  int
	MVszA  int
	MVszW  int
	MVguid int
	MVli   int
	ERR    int
	X      int
}

type SPropValue struct {
	ULPROPTAG  int
	DWALIGNPAD int
	Value      int
}

type SPropProblem struct {
	ULINDEX   int
	ULPROPTAG int
	SCODE     int
}

type SPropProblemArray struct {
	CPROBLEM int
	APROBLEM int
}

type FLATENTRY struct {
	CB      int
	ABENTRY int
}

type FLATENTRYLIST struct {
	CENTRIES  int
	CBENTRIES int
	ABENTRIES int
}

type MTSID struct {
	CB int
	AB int
}

type FLATMTSIDLIST struct {
	CMTSIDS  int
	CBMTSIDS int
	ABMTSIDS int
}

type ADRENTRY struct {
	ULRESERVED1 int
	CVALUES     int
	RGPROPVALS  int
}

type ADRLIST struct {
	CENTRIES int
	AENTRIES int
}

type SRow struct {
	ULADRENTRYPAD int
	CVALUES       int
	LPPROPS       int
}

type SRowSet struct {
	CROWS int
	AROW  int
}

type MAPIERROR struct {
	ULVERSION       int
	LPSZERROR       int
	LPSZCOMPONENT   int
	ULLOWLEVELERROR int
	ULCONTEXT       int
}

type ERROR_NOTIFICATION struct {
	CBENTRYID   int
	LPENTRYID   int
	SCODE       int
	ULFLAGS     int
	LPMAPIERROR int
}

type NEWMAIL_NOTIFICATION struct {
	CBENTRYID        int
	LPENTRYID        int
	CBPARENTID       int
	LPPARENTID       int
	ULFLAGS          int
	LPSZMESSAGECLASS int
	ULMESSAGEFLAGS   int
}

type OBJECT_NOTIFICATION struct {
	CBENTRYID      int
	LPENTRYID      int
	ULOBJTYPE      int
	CBPARENTID     int
	LPPARENTID     int
	CBOLDID        int
	LPOLDID        int
	CBOLDPARENTID  int
	LPOLDPARENTID  int
	LPPROPTAGARRAY int
}

type TABLE_NOTIFICATION struct {
	ULTABLEEVENT int
	HRESULT      int
	PROPINDEX    int
	PROPPRIOR    int
	ROW          int
	ULPAD        int
}

type EXTENDED_NOTIFICATION struct {
	ULEVENT           int
	CB                int
	PBEVENTPARAMETERS int
}

type STATUS_OBJECT_NOTIFICATION struct {
	CBENTRYID  int
	LPENTRYID  int
	CVALUES    int
	LPPROPVALS int
}

type NOTIFICATION struct {
	ULEVENTTYPE int
	ULALIGNPAD  int
	INFO        int
}

type MAPINAMEID struct {
	LPGUID int
	ULKIND int
	Kind   int
}

type SSortOrder struct {
	ULPROPTAG int
	ULORDER   int
}

type SSortOrderSet struct {
	CSORTS      int
	CCATEGORIES int
	CEXPANDED   int
	ASORT       int
}

type SAndRestriction struct {
	CRES  int
	LPRES int
}

type SOrRestriction struct {
	CRES  int
	LPRES int
}

type SNotRestriction struct {
	ULRESERVED int
	LPRES      int
}

type SContentRestriction struct {
	ULFUZZYLEVEL int
	ULPROPTAG    int
	LPPROP       int
}

type SBitMaskRestriction struct {
	RELBMR    int
	ULPROPTAG int
	ULMASK    int
}

type SPropertyRestriction struct {
	RELOP     int
	ULPROPTAG int
	LPPROP    int
}

type SComparePropsRestriction struct {
	RELOP      int
	ULPROPTAG1 int
	ULPROPTAG2 int
}

type SSizeRestriction struct {
	RELOP     int
	ULPROPTAG int
	CB        int
}

type SExistRestriction struct {
	ULRESERVED1 int
	ULPROPTAG   int
	ULRESERVED2 int
}

type SSubRestriction struct {
	ULSUBOBJECT int
	LPRES       int
}

type SCommentRestriction struct {
	CVALUES int
	LPRES   int
	LPPROP  int
}

type SRestriction struct {
	RT  int
	RES int
}

type FLAGLIST struct {
	CFLAGS int
	ULFLAG int
}

type ADRPARM struct {
	CBABCONTENTRYID    int
	LPABCONTENTRYID    int
	ULFLAGS            int
	LPRESERVED         int
	ULHELPCONTEXT      int
	LPSZHELPFILENAME   int
	LPFNABSDI          int
	LPFNDISMISS        int
	LPVDISMISSCONTEXT  int
	LPSZCAPTION        int
	LPSZNEWENTRYTITLE  int
	LPSZDESTWELLSTITLE int
	CDESTFIELDS        int
	NDESTFIELDFOCUS    int
	LPPSZDESTTITLES    int
	LPULDESTCOMPS      int
	LPCONTRESTRICTION  int
	LPHIERRESTRICTION  int
}

type DTBLLABEL struct {
	ULBLPSZLABELNAME int
	ULFLAGS          int
}

type DTBLEDIT struct {
	ULBLPSZCHARSALLOWED int
	ULFLAGS             int
	ULNUMCHARSALLOWED   int
	ULPROPTAG           int
}

type DTBLLBX struct {
	ULFLAGS         int
	ULPRSETPROPERTY int
	ULPRTABLENAME   int
}

type DTBLCOMBOBOX struct {
	ULBLPSZCHARSALLOWED int
	ULFLAGS             int
	ULNUMCHARSALLOWED   int
	ULPRPROPERTYNAME    int
	ULPRTABLENAME       int
}

type DTBLDDLBX struct {
	ULFLAGS             int
	ULPRDISPLAYPROPERTY int
	ULPRSETPROPERTY     int
	ULPRTABLENAME       int
}

type DTBLCHECKBOX struct {
	ULBLPSZLABEL     int
	ULFLAGS          int
	ULPRPROPERTYNAME int
}

type DTBLGROUPBOX struct {
	ULBLPSZLABEL int
	ULFLAGS      int
}

type DTBLBUTTON struct {
	ULBLPSZLABEL int
	ULFLAGS      int
	ULPRCONTROL  int
}

type DTBLPAGE struct {
	ULBLPSZLABEL     int
	ULFLAGS          int
	ULBLPSZCOMPONENT int
	ULCONTEXT        int
}

type DTBLRADIOBUTTON struct {
	ULBLPSZLABEL int
	ULFLAGS      int
	ULCBUTTONS   int
	ULPROPTAG    int
	LRETURNVALUE int
}

type DTBLMVLISTBOX struct {
	ULFLAGS     int
	ULMVPROPTAG int
}

type DTBLMVDDLBX struct {
	ULFLAGS     int
	ULMVPROPTAG int
}

type WABACTIONITEM struct {
}

type WAB_PARAM struct {
	CBSIZE     int
	HWND       int
	SZFILENAME int
	ULFLAGS    int
	GUIDPSEXT  int
}

type WABIMPORTPARAM struct {
	CBSIZE       int
	LPADRBOOK    int
	HWND         int
	ULFLAGS      int
	LPSZFILENAME int
}

type WABEXTDISPLAY struct {
	CBSIZE       int
	LPWABOBJECT  int
	LPADRBOOK    int
	LPPROPOBJ    int
	FREADONLY    int
	FDATACHANGED int
	ULFLAGS      int
	LPV          int
	LPSZ         int
}