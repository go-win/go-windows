// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package intl implements the Windows.Win32.Intl namespace.
package intl

type FONTSIGNATURE struct {
	FSUSB int
	FSCSB int
}

type CHARSETINFO struct {
	CICHARSET int
	CIACP     int
	FS        int
}

type LOCALESIGNATURE struct {
	LSUSB          int
	LSCSBDEFAULT   int
	LSCSBSUPPORTED int
}

type CPINFO struct {
	MaxCharSize int
	DefaultChar int
	LeadByte    int
}

type CPINFOEXA struct {
	MaxCharSize        int
	DefaultChar        int
	LeadByte           int
	UnicodeDefaultChar int
	CodePage           int
	CodePageName       int
}

type CPINFOEXW struct {
	MaxCharSize        int
	DefaultChar        int
	LeadByte           int
	UnicodeDefaultChar int
	CodePage           int
	CodePageName       int
}

type NUMBERFMTA struct {
	NumDigits     int
	LeadingZero   int
	Grouping      int
	LPDECIMALSEP  int
	LPTHOUSANDSEP int
	NegativeOrder int
}

type NUMBERFMTW struct {
	NumDigits     int
	LeadingZero   int
	Grouping      int
	LPDECIMALSEP  int
	LPTHOUSANDSEP int
	NegativeOrder int
}

type CURRENCYFMTA struct {
	NumDigits        int
	LeadingZero      int
	Grouping         int
	LPDECIMALSEP     int
	LPTHOUSANDSEP    int
	NegativeOrder    int
	PositiveOrder    int
	LPCURRENCYSYMBOL int
}

type CURRENCYFMTW struct {
	NumDigits        int
	LeadingZero      int
	Grouping         int
	LPDECIMALSEP     int
	LPTHOUSANDSEP    int
	NegativeOrder    int
	PositiveOrder    int
	LPCURRENCYSYMBOL int
}

type NLSVERSIONINFO struct {
	DWNLSVERSIONINFOSIZE int
	DWNLSVERSION         int
	DWDEFINEDVERSION     int
	DWEFFECTIVEID        int
	GUIDCUSTOMVERSION    int
}

type NLSVERSIONINFOEX struct {
	DWNLSVERSIONINFOSIZE int
	DWNLSVERSION         int
	DWDEFINEDVERSION     int
	DWEFFECTIVEID        int
	GUIDCUSTOMVERSION    int
}

type FILEMUIINFO struct {
	DWSIZE               int
	DWVERSION            int
	DWFILETYPE           int
	PCHECKSUM            int
	PSERVICECHECKSUM     int
	DWLANGUAGENAMEOFFSET int
	DWTYPEIDMAINSIZE     int
	DWTYPEIDMAINOFFSET   int
	DWTYPENAMEMAINOFFSET int
	DWTYPEIDMUISIZE      int
	DWTYPEIDMUIOFFSET    int
	DWTYPENAMEMUIOFFSET  int
	ABBUFFER             int
}

type HSAVEDUILANGUAGES__ struct {
	UNUSED int
}

type HIMC__ struct {
	UNUSED int
}

type HIMCC__ struct {
	UNUSED int
}

type COMPOSITIONFORM struct {
	DWSTYLE      int
	PTCURRENTPOS int
	RCAREA       int
}

type CANDIDATEFORM struct {
	DWINDEX      int
	DWSTYLE      int
	PTCURRENTPOS int
	RCAREA       int
}

type CANDIDATELIST struct {
	DWSIZE      int
	DWSTYLE     int
	DWCOUNT     int
	DWSELECTION int
	DWPAGESTART int
	DWPAGESIZE  int
	DWOFFSET    int
}

type REGISTERWORDA struct {
	LPREADING int
	LPWORD    int
}

type REGISTERWORDW struct {
	LPREADING int
	LPWORD    int
}

type RECONVERTSTRING struct {
	DWSIZE            int
	DWVERSION         int
	DWSTRLEN          int
	DWSTROFFSET       int
	DWCOMPSTRLEN      int
	DWCOMPSTROFFSET   int
	DWTARGETSTRLEN    int
	DWTARGETSTROFFSET int
}

type STYLEBUFA struct {
	DWSTYLE       int
	SZDESCRIPTION int
}

type STYLEBUFW struct {
	DWSTYLE       int
	SZDESCRIPTION int
}

type IMEMENUITEMINFOA struct {
	CBSIZE        int
	FTYPE         int
	FSTATE        int
	WID           int
	HBMPCHECKED   int
	HBMPUNCHECKED int
	DWITEMDATA    int
	SZSTRING      int
	HBMPITEM      int
}

type IMEMENUITEMINFOW struct {
	CBSIZE        int
	FTYPE         int
	FSTATE        int
	WID           int
	HBMPCHECKED   int
	HBMPUNCHECKED int
	DWITEMDATA    int
	SZSTRING      int
	HBMPITEM      int
}

type IMECHARPOSITION struct {
	DWSIZE      int
	DWCHARPOS   int
	PT          int
	CLINEHEIGHT int
	RCDOCUMENT  int
}

type MAPPING_SERVICE_INFO struct {
	Size                      int
	PSZCOPYRIGHT              int
	WMAJORVERSION             int
	WMINORVERSION             int
	WBUILDVERSION             int
	WSTEPVERSION              int
	DWINPUTCONTENTTYPESCOUNT  int
	PRGINPUTCONTENTTYPES      int
	DWOUTPUTCONTENTTYPESCOUNT int
	PRGOUTPUTCONTENTTYPES     int
	DWINPUTLANGUAGESCOUNT     int
	PRGINPUTLANGUAGES         int
	DWOUTPUTLANGUAGESCOUNT    int
	PRGOUTPUTLANGUAGES        int
	DWINPUTSCRIPTSCOUNT       int
	PRGINPUTSCRIPTS           int
	DWOUTPUTSCRIPTSCOUNT      int
	PRGOUTPUTSCRIPTS          int
	GUID                      int
	PSZCATEGORY               int
	PSZDESCRIPTION            int
	DWPRIVATEDATASIZE         int
	PPRIVATEDATA              int
	PCONTEXT                  int
	BITFIELD                  int
}

type MAPPING_ENUM_OPTIONS struct {
	Size                 int
	PSZCATEGORY          int
	PSZINPUTLANGUAGE     int
	PSZOUTPUTLANGUAGE    int
	PSZINPUTSCRIPT       int
	PSZOUTPUTSCRIPT      int
	PSZINPUTCONTENTTYPE  int
	PSZOUTPUTCONTENTTYPE int
	PGUID                int
	BITFIELD             int
}

type MAPPING_OPTIONS struct {
	Size                      int
	PSZINPUTLANGUAGE          int
	PSZOUTPUTLANGUAGE         int
	PSZINPUTSCRIPT            int
	PSZOUTPUTSCRIPT           int
	PSZINPUTCONTENTTYPE       int
	PSZOUTPUTCONTENTTYPE      int
	PSZUILANGUAGE             int
	PFNRECOGNIZECALLBACK      int
	PRECOGNIZECALLERDATA      int
	DWRECOGNIZECALLERDATASIZE int
	PFNACTIONCALLBACK         int
	PACTIONCALLERDATA         int
	DWACTIONCALLERDATASIZE    int
	DWSERVICEFLAG             int
	BITFIELD                  int
}

type MAPPING_DATA_RANGE struct {
	DWSTARTINDEX          int
	DWENDINDEX            int
	PSZDESCRIPTION        int
	DWDESCRIPTIONLENGTH   int
	PDATA                 int
	DWDATASIZE            int
	PSZCONTENTTYPE        int
	PRGACTIONIDS          int
	DWACTIONSCOUNT        int
	PRGACTIONDISPLAYNAMES int
}

type MAPPING_PROPERTY_BAG struct {
	Size              int
	PRGRESULTRANGES   int
	DWRANGESCOUNT     int
	PSERVICEDATA      int
	DWSERVICEDATASIZE int
	PCALLERDATA       int
	DWCALLERDATASIZE  int
	PCONTEXT          int
}

type SpellCheckerFactory struct {
}

type IMEDLG struct {
	CBIMEDLG   int
	HWND       int
	LPWSTRWORD int
	NTABID     int
}

type WDD struct {
	WDISPPOS      int
	Anonymous1    int
	CCHDISP       int
	Anonymous2    int
	WDD_nReserve1 int
	NPOS          int
	BITFIELD      int
	PRESERVED     int
}

type MORRSLT struct {
	DWSIZE          int
	PWCHOUTPUT      int
	CCHOUTPUT       int
	Anonymous1      int
	Anonymous2      int
	PCHINPUTPOS     int
	PCHOUTPUTIDXWDD int
	Anonymous3      int
	PAMONORUBYPOS   int
	PWDD            int
	CWDD            int
	PPRIVATE        int
	BLKBuff         int
}

type IMEWRD struct {
	PWCHREADING int
	PWCHDISPLAY int
	Anonymous   int
	RGULATTRS   int
	CBCOMMENT   int
	UCT         int
	PVCOMMENT   int
}

type IMESHF struct {
	CBSHF         int
	VERDIC        int
	SZTITLE       int
	SZDESCRIPTION int
	SZCOPYRIGHT   int
}

type POSTBL struct {
	NPOS   int
	SZNAME int
}

type IMEDP struct {
	WRDMODIFIER int
	WRDMODIFIEE int
	RELID       int
}

type IMEKMSINIT struct {
	CBSIZE int
	HWND   int
}

type IMEKMSKEY struct {
	DWSTATUS     int
	DWCOMPSTATUS int
	DWVKEY       int
	Anonymous1   int
	Anonymous2   int
}

type IMEKMS struct {
	CBSIZE   int
	HIMC     int
	CKEYLIST int
	PKEYLIST int
}

type IMEKMSNTFY struct {
	CBSIZE  int
	HIMC    int
	FSELECT int
}

type IMEKMSKMP struct {
	CBSIZE   int
	HIMC     int
	IDLANG   int
	WVKSTART int
	WVKEND   int
	CKEYLIST int
	PKEYLIST int
}

type IMEKMSINVK struct {
	CBSIZE    int
	HIMC      int
	DWCONTROL int
}

type IMEKMSFUNCDESC struct {
	CBSIZE          int
	IDLANG          int
	DWCONTROL       int
	PWSZDESCRIPTION int
}

type COMPOSITIONSTRING struct {
	DWSIZE                   int
	DWCOMPREADATTRLEN        int
	DWCOMPREADATTROFFSET     int
	DWCOMPREADCLAUSELEN      int
	DWCOMPREADCLAUSEOFFSET   int
	DWCOMPREADSTRLEN         int
	DWCOMPREADSTROFFSET      int
	DWCOMPATTRLEN            int
	DWCOMPATTROFFSET         int
	DWCOMPCLAUSELEN          int
	DWCOMPCLAUSEOFFSET       int
	DWCOMPSTRLEN             int
	DWCOMPSTROFFSET          int
	DWCURSORPOS              int
	DWDELTASTART             int
	DWRESULTREADCLAUSELEN    int
	DWRESULTREADCLAUSEOFFSET int
	DWRESULTREADSTRLEN       int
	DWRESULTREADSTROFFSET    int
	DWRESULTCLAUSELEN        int
	DWRESULTCLAUSEOFFSET     int
	DWRESULTSTRLEN           int
	DWRESULTSTROFFSET        int
	DWPRIVATESIZE            int
	DWPRIVATEOFFSET          int
}

type GUIDELINE struct {
	DWSIZE          int
	DWLEVEL         int
	DWINDEX         int
	DWSTRLEN        int
	DWSTROFFSET     int
	DWPRIVATESIZE   int
	DWPRIVATEOFFSET int
}

type TRANSMSG struct {
	MESSAGE int
	WPARAM  int
	LPARAM  int
}

type TRANSMSGLIST struct {
	UMSGCOUNT int
	TransMsg  int
}

type CANDIDATEINFO struct {
	DWSIZE          int
	DWCOUNT         int
	DWOFFSET        int
	DWPRIVATESIZE   int
	DWPRIVATEOFFSET int
}

type INPUTCONTEXT struct {
	HWND           int
	FOPEN          int
	PTSTATUSWNDPOS int
	PTSOFTKBDPOS   int
	FDWCONVERSION  int
	FDWSENTENCE    int
	LFFONT         int
	CFCOMPFORM     int
	CFCANDFORM     int
	HCOMPSTR       int
	HCANDINFO      int
	HGUIDELINE     int
	HPRIVATE       int
	DWNUMMSGBUF    int
	HMSGBUF        int
	FDWINIT        int
	DWRESERVE      int
}

type IMEINFO struct {
	DWPRIVATEDATASIZE int
	FDWPROPERTY       int
	FDWCONVERSIONCAPS int
	FDWSENTENCECAPS   int
	FDWUICAPS         int
	FDWSCSCAPS        int
	FDWSELECTCAPS     int
}

type SOFTKBDDATA struct {
	UCOUNT int
	WCODE  int
}

type APPLETIDLIST struct {
	COUNT    int
	PIIDLIST int
}

type IMESTRINGCANDIDATE struct {
	UCOUNT int
	LPWSTR int
}

type IMEITEM struct {
	CBSIZE     int
	ITYPE      int
	LPITEMDATA int
}

type IMEITEMCANDIDATE struct {
	UCOUNT  int
	IMEITEM int
}

type TABIMESTRINGINFO struct {
	DWFAREASTID int
	LPWSTR      int
}

type TABIMEFAREASTINFO struct {
	DWSIZE int
	DWTYPE int
	DWDATA int
}

type IMESTRINGCANDIDATEINFO struct {
	DWFAREASTID   int
	LPFAREASTINFO int
	FINFOMASK     int
	ISELINDEX     int
	UCOUNT        int
	LPWSTR        int
}

type IMECOMPOSITIONSTRINGINFO struct {
	ICOMPSTRLEN  int
	ICARETPOS    int
	IEDITSTART   int
	IEDITLEN     int
	ITARGETSTART int
	ITARGETLEN   int
}

type IMECHARINFO struct {
	WCH        int
	DWCHARINFO int
}

type IMEAPPLETCFG struct {
	DWCONFIG         int
	WCHTITLE         int
	WCHTITLEFONTFACE int
	DWCHARSET        int
	ICATEGORY        int
	HICON            int
	LANGID           int
	DUMMY            int
	LRESERVED1       int
}

type IMEAPPLETUI struct {
	HWND       int
	DWSTYLE    int
	WIDTH      int
	HEIGHT     int
	MINWIDTH   int
	MINHEIGHT  int
	MAXWIDTH   int
	MAXHEIGHT  int
	LRESERVED1 int
	LRESERVED2 int
}

type APPLYCANDEXPARAM struct {
	DWSIZE        int
	LPWSTRDISPLAY int
	LPWSTRREADING int
	DWRESERVED    int
}

type SCRIPT_CONTROL struct {
	BITFIELD int
}

type SCRIPT_STATE struct {
	BITFIELD int
}

type SCRIPT_ANALYSIS struct {
	BITFIELD int
	S        int
}

type SCRIPT_ITEM struct {
	ICHARPOS int
	A        int
}

type SCRIPT_VISATTR struct {
	BITFIELD int
}

type GOFFSET struct {
	DU int
	DV int
}

type SCRIPT_LOGATTR struct {
	BITFIELD int
}

type SCRIPT_PROPERTIES struct {
	BITFIELD1 int
	BITFIELD2 int
}

type SCRIPT_FONTPROPERTIES struct {
	CBYTES        int
	WGBLANK       int
	WGDEFAULT     int
	WGINVALID     int
	WGKASHIDA     int
	IKASHIDAWIDTH int
}

type SCRIPT_TABDEF struct {
	CTABSTOPS  int
	ISCALE     int
	PTABSTOPS  int
	ITABORIGIN int
}

type SCRIPT_DIGITSUBSTITUTE struct {
	BITFIELD1  int
	BITFIELD2  int
	DWRESERVED int
}

type OPENTYPE_FEATURE_RECORD struct {
	TAGFEATURE int
	LPARAMETER int
}

type TEXTRANGE_PROPERTIES struct {
	POTFRECORDS int
	COTFRECORDS int
}

type SCRIPT_CHARPROP struct {
	BITFIELD int
}

type SCRIPT_GLYPHPROP struct {
	SVA      int
	RESERVED int
}

type UReplaceableCallbacks struct {
	LENGTH   int
	CHARAT   int
	CHAR32AT int
	REPLACE  int
	EXTRACT  int
	COPY     int
}

type UFieldPosition struct {
	FIELD      int
	BEGININDEX int
	ENDINDEX   int
}

type UCharIterator struct {
	CONTEXT       int
	LENGTH        int
	START         int
	INDEX         int
	LIMIT         int
	RESERVEDFIELD int
	GETINDEX      int
	MOVE          int
	HASNEXT       int
	HASPREVIOUS   int
	CURRENT       int
	NEXT          int
	PREVIOUS      int
	RESERVEDFN    int
	GETSTATE      int
	SETSTATE      int
}

type UEnumeration struct {
}

type UResourceBundle struct {
}

type ULocaleDisplayNames struct {
}

type UConverter struct {
}

type UConverterFromUnicodeArgs struct {
	SIZE        int
	FLUSH       int
	CONVERTER   int
	SOURCE      int
	SOURCELIMIT int
	TARGET      int
	TARGETLIMIT int
	OFFSETS     int
}

type UConverterToUnicodeArgs struct {
	SIZE        int
	FLUSH       int
	CONVERTER   int
	SOURCE      int
	SOURCELIMIT int
	TARGET      int
	TARGETLIMIT int
	OFFSETS     int
}

type USet struct {
}

type UBiDi struct {
}

type UBiDiTransform struct {
}

type UTextFuncs struct {
	TABLESIZE             int
	RESERVED1             int
	RESERVED2             int
	RESERVED3             int
	CLONE                 int
	NATIVELENGTH          int
	ACCESS                int
	EXTRACT               int
	REPLACE               int
	COPY                  int
	MAPOFFSETTONATIVE     int
	MAPNATIVEINDEXTOUTF16 int
	CLOSE                 int
	SPARE1                int
	SPARE2                int
	SPARE3                int
}

type UText struct {
	MAGIC               int
	FLAGS               int
	PROVIDERPROPERTIES  int
	SIZEOFSTRUCT        int
	CHUNKNATIVELIMIT    int
	EXTRASIZE           int
	NATIVEINDEXINGLIMIT int
	CHUNKNATIVESTART    int
	CHUNKOFFSET         int
	CHUNKLENGTH         int
	CHUNKCONTENTS       int
	PFUNCS              int
	PEXTRA              int
	CONTEXT             int
	P                   int
	Q                   int
	R                   int
	PRIVP               int
	A                   int
	B                   int
	C                   int
	PRIVA               int
	PRIVB               int
	PRIVC               int
}

type USerializedSet struct {
	ARRAY       int
	BMPLENGTH   int
	LENGTH      int
	STATICARRAY int
}

type UNormalizer2 struct {
}

type UConverterSelector struct {
}

type UBreakIterator struct {
}

type UCaseMap struct {
}

type UParseError struct {
	LINE        int
	OFFSET      int
	PRECONTEXT  int
	POSTCONTEXT int
}

type UStringPrepProfile struct {
}

type UIDNA struct {
}

type UIDNAInfo struct {
	SIZE                    int
	ISTRANSITIONALDIFFERENT int
	RESERVEDB3              int
	ERRORS                  int
	RESERVEDI2              int
	RESERVEDI3              int
}

type UCollator struct {
}

type UCollationElements struct {
}

type UCharsetDetector struct {
}

type UCharsetMatch struct {
}

type UFieldPositionIterator struct {
}

type UDateIntervalFormat struct {
}

type UGenderInfo struct {
}

type UListFormatter struct {
}

type ULocaleData struct {
}

type UDateFormatSymbols struct {
}

type UNumberFormatter struct {
}

type UFormattedNumber struct {
}

type UNumberingSystem struct {
}

type UPluralRules struct {
}

type URegularExpression struct {
}

type URegion struct {
}

type URelativeDateTimeFormatter struct {
}

type UStringSearch struct {
}

type USpoofChecker struct {
}

type USpoofCheckResult struct {
}

type UTransPosition struct {
	CONTEXTSTART int
	CONTEXTLIMIT int
	START        int
	LIMIT        int
}
