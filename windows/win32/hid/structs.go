// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package hid implements the Windows.Win32.Hid namespace.
package hid

type DICONSTANTFORCE struct {
	LMAGNITUDE int
}

type DIRAMPFORCE struct {
	LSTART int
	LEND   int
}

type DIPERIODIC struct {
	DWMAGNITUDE int
	LOFFSET     int
	DWPHASE     int
	DWPERIOD    int
}

type DICONDITION struct {
	LOFFSET              int
	LPOSITIVECOEFFICIENT int
	LNEGATIVECOEFFICIENT int
	DWPOSITIVESATURATION int
	DWNEGATIVESATURATION int
	LDEADBAND            int
}

type DICUSTOMFORCE struct {
	CCHANNELS      int
	DWSAMPLEPERIOD int
	CSAMPLES       int
	RGLFORCEDATA   int
}

type DIENVELOPE struct {
	DWSIZE        int
	DWATTACKLEVEL int
	DWATTACKTIME  int
	DWFADELEVEL   int
	DWFADETIME    int
}

type DIEFFECT_DX5 struct {
	DWSIZE                  int
	DWFLAGS                 int
	DWDURATION              int
	DWSAMPLEPERIOD          int
	DWGAIN                  int
	DWTRIGGERBUTTON         int
	DWTRIGGERREPEATINTERVAL int
	CAXES                   int
	RGDWAXES                int
	RGLDIRECTION            int
	LPENVELOPE              int
	CBTYPESPECIFICPARAMS    int
	LPVTYPESPECIFICPARAMS   int
}

type DIEFFECT struct {
	DWSIZE                  int
	DWFLAGS                 int
	DWDURATION              int
	DWSAMPLEPERIOD          int
	DWGAIN                  int
	DWTRIGGERBUTTON         int
	DWTRIGGERREPEATINTERVAL int
	CAXES                   int
	RGDWAXES                int
	RGLDIRECTION            int
	LPENVELOPE              int
	CBTYPESPECIFICPARAMS    int
	LPVTYPESPECIFICPARAMS   int
	DWSTARTDELAY            int
}

type DIFILEEFFECT struct {
	DWSIZE         int
	GuidEffect     int
	LPDIEFFECT     int
	SZFRIENDLYNAME int
}

type DIEFFESCAPE struct {
	DWSIZE       int
	DWCOMMAND    int
	LPVINBUFFER  int
	CBINBUFFER   int
	LPVOUTBUFFER int
	CBOUTBUFFER  int
}

type DIDEVCAPS_DX3 struct {
	DWSIZE    int
	DWFLAGS   int
	DWDEVTYPE int
	DWAXES    int
	DWBUTTONS int
	DWPOVS    int
}

type DIDEVCAPS struct {
	DWSIZE                int
	DWFLAGS               int
	DWDEVTYPE             int
	DWAXES                int
	DWBUTTONS             int
	DWPOVS                int
	DWFFSAMPLEPERIOD      int
	DWFFMINTIMERESOLUTION int
	DWFIRMWAREREVISION    int
	DWHARDWAREREVISION    int
	DWFFDRIVERVERSION     int
}

type DIOBJECTDATAFORMAT struct {
	PGUID   int
	DWOFS   int
	DWTYPE  int
	DWFLAGS int
}

type DIDATAFORMAT struct {
	DWSIZE     int
	DWOBJSIZE  int
	DWFLAGS    int
	DWDATASIZE int
	DWNUMOBJS  int
	RGODF      int
}

type DIACTIONA struct {
	UAPPDATA     int
	DWSEMANTIC   int
	DWFLAGS      int
	Anonymous    int
	GUIDINSTANCE int
	DWOBJID      int
	DWHOW        int
}

type DIACTIONW struct {
	UAPPDATA     int
	DWSEMANTIC   int
	DWFLAGS      int
	Anonymous    int
	GUIDINSTANCE int
	DWOBJID      int
	DWHOW        int
}

type DIACTIONFORMATA struct {
	DWSIZE        int
	DWACTIONSIZE  int
	DWDATASIZE    int
	DWNUMACTIONS  int
	RGOACTION     int
	GUIDACTIONMAP int
	DWGENRE       int
	DWBUFFERSIZE  int
	LAXISMIN      int
	LAXISMAX      int
	HINSTSTRING   int
	FTTIMESTAMP   int
	DWCRC         int
	TSZACTIONMAP  int
}

type DIACTIONFORMATW struct {
	DWSIZE        int
	DWACTIONSIZE  int
	DWDATASIZE    int
	DWNUMACTIONS  int
	RGOACTION     int
	GUIDACTIONMAP int
	DWGENRE       int
	DWBUFFERSIZE  int
	LAXISMIN      int
	LAXISMAX      int
	HINSTSTRING   int
	FTTIMESTAMP   int
	DWCRC         int
	TSZACTIONMAP  int
}

type DICOLORSET struct {
	DWSIZE            int
	CTEXTFORE         int
	CTEXTHIGHLIGHT    int
	CCALLOUTLINE      int
	CCALLOUTHIGHLIGHT int
	CBORDER           int
	CCONTROLFILL      int
	CHIGHLIGHTFILL    int
	CAREAFILL         int
}

type DICONFIGUREDEVICESPARAMSA struct {
	DWSIZE         int
	DWCUSERS       int
	LPTSZUSERNAMES int
	DWCFORMATS     int
	LPRGFORMATS    int
	HWND           int
	DICS           int
	LPUNKDDSTARGET int
}

type DICONFIGUREDEVICESPARAMSW struct {
	DWSIZE         int
	DWCUSERS       int
	LPTSZUSERNAMES int
	DWCFORMATS     int
	LPRGFORMATS    int
	HWND           int
	DICS           int
	LPUNKDDSTARGET int
}

type DIDEVICEIMAGEINFOA struct {
	TSZIMAGEPATH    int
	DWFLAGS         int
	DWVIEWID        int
	RCOVERLAY       int
	DWOBJID         int
	DWCVALIDPTS     int
	RGPTCALLOUTLINE int
	RCCALLOUTRECT   int
	DWTEXTALIGN     int
}

type DIDEVICEIMAGEINFOW struct {
	TSZIMAGEPATH    int
	DWFLAGS         int
	DWVIEWID        int
	RCOVERLAY       int
	DWOBJID         int
	DWCVALIDPTS     int
	RGPTCALLOUTLINE int
	RCCALLOUTRECT   int
	DWTEXTALIGN     int
}

type DIDEVICEIMAGEINFOHEADERA struct {
	DWSIZE             int
	DWSIZEIMAGEINFO    int
	DWCVIEWS           int
	DWCBUTTONS         int
	DWCAXES            int
	DWCPOVS            int
	DWBUFFERSIZE       int
	DWBUFFERUSED       int
	LPRGIMAGEINFOARRAY int
}

type DIDEVICEIMAGEINFOHEADERW struct {
	DWSIZE             int
	DWSIZEIMAGEINFO    int
	DWCVIEWS           int
	DWCBUTTONS         int
	DWCAXES            int
	DWCPOVS            int
	DWBUFFERSIZE       int
	DWBUFFERUSED       int
	LPRGIMAGEINFOARRAY int
}

type DIDEVICEOBJECTINSTANCE_DX3A struct {
	DWSIZE   int
	GUIDTYPE int
	DWOFS    int
	DWTYPE   int
	DWFLAGS  int
	TSZNAME  int
}

type DIDEVICEOBJECTINSTANCE_DX3W struct {
	DWSIZE   int
	GUIDTYPE int
	DWOFS    int
	DWTYPE   int
	DWFLAGS  int
	TSZNAME  int
}

type DIDEVICEOBJECTINSTANCEA struct {
	DWSIZE              int
	GUIDTYPE            int
	DWOFS               int
	DWTYPE              int
	DWFLAGS             int
	TSZNAME             int
	DWFFMAXFORCE        int
	DWFFFORCERESOLUTION int
	WCOLLECTIONNUMBER   int
	WDESIGNATORINDEX    int
	WUSAGEPAGE          int
	WUSAGE              int
	DWDIMENSION         int
	WEXPONENT           int
	WREPORTID           int
}

type DIDEVICEOBJECTINSTANCEW struct {
	DWSIZE              int
	GUIDTYPE            int
	DWOFS               int
	DWTYPE              int
	DWFLAGS             int
	TSZNAME             int
	DWFFMAXFORCE        int
	DWFFFORCERESOLUTION int
	WCOLLECTIONNUMBER   int
	WDESIGNATORINDEX    int
	WUSAGEPAGE          int
	WUSAGE              int
	DWDIMENSION         int
	WEXPONENT           int
	WREPORTID           int
}

type DIPROPHEADER struct {
	DWSIZE       int
	DWHEADERSIZE int
	DWOBJ        int
	DWHOW        int
}

type DIPROPDWORD struct {
	DIPH   int
	DWDATA int
}

type DIPROPPOINTER struct {
	DIPH  int
	UDATA int
}

type DIPROPRANGE struct {
	DIPH int
	LMIN int
	LMAX int
}

type DIPROPCAL struct {
	DIPH    int
	LMIN    int
	LCENTER int
	LMAX    int
}

type DIPROPCALPOV struct {
	DIPH int
	LMIN int
	LMAX int
}

type DIPROPGUIDANDPATH struct {
	DIPH      int
	GUIDCLASS int
	WSZPATH   int
}

type DIPROPSTRING struct {
	DIPH int
	WSZ  int
}

type CPOINT struct {
	LP    int
	DWLOG int
}

type DIPROPCPOINTS struct {
	DIPH         int
	DWCPOINTSNUM int
	CP           int
}

type DIDEVICEOBJECTDATA_DX3 struct {
	DWOFS       int
	DWDATA      int
	DWTIMESTAMP int
	DWSEQUENCE  int
}

type DIDEVICEOBJECTDATA struct {
	DWOFS       int
	DWDATA      int
	DWTIMESTAMP int
	DWSEQUENCE  int
	UAPPDATA    int
}

type DIDEVICEINSTANCE_DX3A struct {
	DWSIZE          int
	GUIDINSTANCE    int
	GUIDPRODUCT     int
	DWDEVTYPE       int
	TSZINSTANCENAME int
	TSZPRODUCTNAME  int
}

type DIDEVICEINSTANCE_DX3W struct {
	DWSIZE          int
	GUIDINSTANCE    int
	GUIDPRODUCT     int
	DWDEVTYPE       int
	TSZINSTANCENAME int
	TSZPRODUCTNAME  int
}

type DIDEVICEINSTANCEA struct {
	DWSIZE          int
	GUIDINSTANCE    int
	GUIDPRODUCT     int
	DWDEVTYPE       int
	TSZINSTANCENAME int
	TSZPRODUCTNAME  int
	GUIDFFDRIVER    int
	WUSAGEPAGE      int
	WUSAGE          int
}

type DIDEVICEINSTANCEW struct {
	DWSIZE          int
	GUIDINSTANCE    int
	GUIDPRODUCT     int
	DWDEVTYPE       int
	TSZINSTANCENAME int
	TSZPRODUCTNAME  int
	GUIDFFDRIVER    int
	WUSAGEPAGE      int
	WUSAGE          int
}

type DIEFFECTINFOA struct {
	DWSIZE          int
	GUID            int
	DWEFFTYPE       int
	DWSTATICPARAMS  int
	DWDYNAMICPARAMS int
	TSZNAME         int
}

type DIEFFECTINFOW struct {
	DWSIZE          int
	GUID            int
	DWEFFTYPE       int
	DWSTATICPARAMS  int
	DWDYNAMICPARAMS int
	TSZNAME         int
}

type DIMOUSESTATE struct {
	LX         int
	LY         int
	LZ         int
	RGBBUTTONS int
}

type DIMOUSESTATE2 struct {
	LX         int
	LY         int
	LZ         int
	RGBBUTTONS int
}

type DIJOYSTATE struct {
	LX         int
	LY         int
	LZ         int
	LRX        int
	LRY        int
	LRZ        int
	RGLSLIDER  int
	RGDWPOV    int
	RGBBUTTONS int
}

type DIJOYSTATE2 struct {
	LX         int
	LY         int
	LZ         int
	LRX        int
	LRY        int
	LRZ        int
	RGLSLIDER  int
	RGDWPOV    int
	RGBBUTTONS int
	LVX        int
	LVY        int
	LVZ        int
	LVRX       int
	LVRY       int
	LVRZ       int
	RGLVSLIDER int
	LAX        int
	LAY        int
	LAZ        int
	LARX       int
	LARY       int
	LARZ       int
	RGLASLIDER int
	LFX        int
	LFY        int
	LFZ        int
	LFRX       int
	LFRY       int
	LFRZ       int
	RGLFSLIDER int
}

type DIOBJECTATTRIBUTES struct {
	DWFLAGS    int
	WUSAGEPAGE int
	WUSAGE     int
}

type DIFFOBJECTATTRIBUTES struct {
	DWFFMAXFORCE        int
	DWFFFORCERESOLUTION int
}

type DIOBJECTCALIBRATION struct {
	LMIN    int
	LCENTER int
	LMAX    int
}

type DIPOVCALIBRATION struct {
	LMIN int
	LMAX int
}

type DIEFFECTATTRIBUTES struct {
	DWEFFECTID      int
	DWEFFTYPE       int
	DWSTATICPARAMS  int
	DWDYNAMICPARAMS int
	DWCOORDS        int
}

type DIFFDEVICEATTRIBUTES struct {
	DWFLAGS               int
	DWFFSAMPLEPERIOD      int
	DWFFMINTIMERESOLUTION int
}

type DIDRIVERVERSIONS struct {
	DWSIZE             int
	DWFIRMWAREREVISION int
	DWHARDWAREREVISION int
	DWFFDRIVERVERSION  int
}

type DIDEVICESTATE struct {
	DWSIZE  int
	DWSTATE int
	DWLOAD  int
}

type DIHIDFFINITINFO struct {
	DWSIZE              int
	PWSZDEVICEINTERFACE int
	GuidInstance        int
}

type DIJOYTYPEINFO_DX5 struct {
	DWSIZE         int
	HWS            int
	CLSIDCONFIG    int
	WSZDISPLAYNAME int
	WSZCALLOUT     int
}

type DIJOYTYPEINFO_DX6 struct {
	DWSIZE         int
	HWS            int
	CLSIDCONFIG    int
	WSZDISPLAYNAME int
	WSZCALLOUT     int
	WSZHARDWAREID  int
	DWFLAGS1       int
}

type DIJOYTYPEINFO struct {
	DWSIZE         int
	HWS            int
	CLSIDCONFIG    int
	WSZDISPLAYNAME int
	WSZCALLOUT     int
	WSZHARDWAREID  int
	DWFLAGS1       int
	DWFLAGS2       int
	WSZMAPFILE     int
}

type DIJOYCONFIG_DX5 struct {
	DWSIZE       int
	GUIDINSTANCE int
	HWC          int
	DWGAIN       int
	WSZTYPE      int
	WSZCALLOUT   int
}

type DIJOYCONFIG struct {
	DWSIZE       int
	GUIDINSTANCE int
	HWC          int
	DWGAIN       int
	WSZTYPE      int
	WSZCALLOUT   int
	GUIDGAMEPORT int
}

type DIJOYUSERVALUES struct {
	DWSIZE              int
	RUV                 int
	WSZGLOBALDRIVER     int
	WSZGAMEPORTEMULATOR int
}

type KEYBOARD_INPUT_DATA struct {
	UnitId           int
	MakeCode         int
	Flags            int
	Reserved         int
	ExtraInformation int
}

type KEYBOARD_TYPEMATIC_PARAMETERS struct {
	UnitId int
	Rate   int
	Delay  int
}

type KEYBOARD_ID struct {
	Type    int
	Subtype int
}

type KEYBOARD_ATTRIBUTES struct {
	KeyboardIdentifier   int
	KeyboardMode         int
	NumberOfFunctionKeys int
	NumberOfIndicators   int
	NumberOfKeysTotal    int
	InputDataQueueLength int
	KeyRepeatMinimum     int
	KeyRepeatMaximum     int
}

type KEYBOARD_EXTENDED_ATTRIBUTES struct {
	Version                        int
	FormFactor                     int
	KeyType                        int
	PhysicalLayout                 int
	VendorSpecificPhysicalLayout   int
	IETFLanguageTagIndex           int
	ImplementedInputAssistControls int
}

type KEYBOARD_INDICATOR_PARAMETERS struct {
	UnitId   int
	LedFlags int
}

type INDICATOR_LIST struct {
	MakeCode       int
	IndicatorFlags int
}

type KEYBOARD_INDICATOR_TRANSLATION struct {
	NumberOfIndicatorKeys int
	IndicatorList         int
}

type KEYBOARD_UNIT_ID_PARAMETER struct {
	UnitId int
}

type KEYBOARD_IME_STATUS struct {
	UnitId      int
	ImeOpen     int
	ImeConvMode int
}

type MOUSE_INPUT_DATA struct {
	UnitId           int
	Flags            int
	Anonymous        int
	RawButtons       int
	LastX            int
	LastY            int
	ExtraInformation int
}

type MOUSE_ATTRIBUTES struct {
	MouseIdentifier      int
	NumberOfButtons      int
	SampleRate           int
	InputDataQueueLength int
}

type MOUSE_UNIT_ID_PARAMETER struct {
	UnitId int
}

type USAGE_AND_PAGE struct {
	Usage     int
	UsagePage int
}

type HIDP_BUTTON_CAPS struct {
	UsagePage         int
	ReportID          int
	IsAlias           int
	BitField          int
	LinkCollection    int
	LinkUsage         int
	LinkUsagePage     int
	IsRange           int
	IsStringRange     int
	IsDesignatorRange int
	IsAbsolute        int
	Reserved          int
	Anonymous         int
}

type HIDP_VALUE_CAPS struct {
	UsagePage         int
	ReportID          int
	IsAlias           int
	BitField          int
	LinkCollection    int
	LinkUsage         int
	LinkUsagePage     int
	IsRange           int
	IsStringRange     int
	IsDesignatorRange int
	IsAbsolute        int
	HasNull           int
	Reserved          int
	BitSize           int
	ReportCount       int
	Reserved2         int
	UnitsExp          int
	Units             int
	LogicalMin        int
	LogicalMax        int
	PhysicalMin       int
	PhysicalMax       int
	Anonymous         int
}

type HIDP_LINK_COLLECTION_NODE struct {
	LinkUsage        int
	LinkUsagePage    int
	Parent           int
	NumberOfChildren int
	NextSibling      int
	FirstChild       int
	BITFIELD         int
	UserContext      int
}

type HIDP_PREPARSED_DATA struct {
}

type HIDP_CAPS struct {
	Usage                     int
	UsagePage                 int
	InputReportByteLength     int
	OutputReportByteLength    int
	FeatureReportByteLength   int
	Reserved                  int
	NumberLinkCollectionNodes int
	NumberInputButtonCaps     int
	NumberInputValueCaps      int
	NumberInputDataIndices    int
	NumberOutputButtonCaps    int
	NumberOutputValueCaps     int
	NumberOutputDataIndices   int
	NumberFeatureButtonCaps   int
	NumberFeatureValueCaps    int
	NumberFeatureDataIndices  int
}

type HIDP_DATA struct {
	DataIndex int
	Reserved  int
	Anonymous int
}

type HIDP_UNKNOWN_TOKEN struct {
	Token    int
	Reserved int
	BitField int
}

type HIDP_EXTENDED_ATTRIBUTES struct {
	NumGlobalUnknowns int
	Reserved          int
	GlobalUnknowns    int
	Data              int
}

type HIDP_KEYBOARD_MODIFIER_STATE struct {
	Anonymous int
}

type HIDD_CONFIGURATION struct {
	COOKIE         int
	SIZE           int
	RingBufferSize int
}

type HIDD_ATTRIBUTES struct {
	Size          int
	VendorID      int
	ProductID     int
	VersionNumber int
}

type JOYREGHWVALUES struct {
	JRVHARDWARE int
	DWPOVVALUES int
	DWCALFLAGS  int
}
