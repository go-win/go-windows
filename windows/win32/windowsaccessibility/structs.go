// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsaccessibility implements the Windows.Win32.WindowsAccessibility namespace.
package windowsaccessibility

type HWINEVENTHOOK struct {
	Value int
}

type SERIALKEYSA struct {
	CBSIZE         int
	DWFLAGS        int
	LPSZACTIVEPORT int
	LPSZPORT       int
	IBAUDRATE      int
	IPORTSTATE     int
	IACTIVE        int
}

type SERIALKEYSW struct {
	CBSIZE         int
	DWFLAGS        int
	LPSZACTIVEPORT int
	LPSZPORT       int
	IBAUDRATE      int
	IPORTSTATE     int
	IACTIVE        int
}

type HIGHCONTRASTA struct {
	CBSIZE            int
	DWFLAGS           int
	LPSZDEFAULTSCHEME int
}

type HIGHCONTRASTW struct {
	CBSIZE            int
	DWFLAGS           int
	LPSZDEFAULTSCHEME int
}

type FILTERKEYS struct {
	CBSIZE      int
	DWFLAGS     int
	IWAITMSEC   int
	IDELAYMSEC  int
	IREPEATMSEC int
	IBOUNCEMSEC int
}

type STICKYKEYS struct {
	CBSIZE  int
	DWFLAGS int
}

type MOUSEKEYS struct {
	CBSIZE          int
	DWFLAGS         int
	IMAXSPEED       int
	ITIMETOMAXSPEED int
	ICTRLSPEED      int
	DWRESERVED1     int
	DWRESERVED2     int
}

type ACCESSTIMEOUT struct {
	CBSIZE       int
	DWFLAGS      int
	ITIMEOUTMSEC int
}

type SOUNDSENTRYA struct {
	CBSIZE                 int
	DWFLAGS                int
	IFSTEXTEFFECT          int
	IFSTEXTEFFECTMSEC      int
	IFSTEXTEFFECTCOLORBITS int
	IFSGRAFEFFECT          int
	IFSGRAFEFFECTMSEC      int
	IFSGRAFEFFECTCOLOR     int
	IWINDOWSEFFECT         int
	IWINDOWSEFFECTMSEC     int
	LPSZWINDOWSEFFECTDLL   int
	IWINDOWSEFFECTORDINAL  int
}

type SOUNDSENTRYW struct {
	CBSIZE                 int
	DWFLAGS                int
	IFSTEXTEFFECT          int
	IFSTEXTEFFECTMSEC      int
	IFSTEXTEFFECTCOLORBITS int
	IFSGRAFEFFECT          int
	IFSGRAFEFFECTMSEC      int
	IFSGRAFEFFECTCOLOR     int
	IWINDOWSEFFECT         int
	IWINDOWSEFFECTMSEC     int
	LPSZWINDOWSEFFECTDLL   int
	IWINDOWSEFFECTORDINAL  int
}

type TOGGLEKEYS struct {
	CBSIZE  int
	DWFLAGS int
}

type MSAAControl struct {
}

type AccStore struct {
}

type AccDictionary struct {
}

type AccServerDocMgr struct {
}

type AccClientDocMgr struct {
}

type DocWrap struct {
}

type CAccPropServices struct {
}

type MSAAMENUINFO struct {
	DWMSAASIGNATURE int
	CCHWTEXT        int
	PSZWTEXT        int
}

type CUIAutomationRegistrar struct {
}

type UiaRect struct {
	LEFT   int
	TOP    int
	WIDTH  int
	HEIGHT int
}

type UiaPoint struct {
	X int
	Y int
}

type UiaChangeInfo struct {
	UIAID     int
	PAYLOAD   int
	EXTRAINFO int
}

type UIAutomationParameter struct {
	TYPE  int
	PDATA int
}

type UIAutomationPropertyInfo struct {
	GUID              int
	PPROGRAMMATICNAME int
	TYPE              int
}

type UIAutomationEventInfo struct {
	GUID              int
	PPROGRAMMATICNAME int
}

type UIAutomationMethodInfo struct {
	PPROGRAMMATICNAME int
	DOSETFOCUS        int
	CINPARAMETERS     int
	COUTPARAMETERS    int
	PPARAMETERTYPES   int
	PPARAMETERNAMES   int
}

type UIAutomationPatternInfo struct {
	GUID                int
	PPROGRAMMATICNAME   int
	PROVIDERINTERFACEID int
	CLIENTINTERFACEID   int
	CPROPERTIES         int
	PPROPERTIES         int
	CMETHODS            int
	PMETHODS            int
	CEVENTS             int
	PEVENTS             int
	PPATTERNHANDLER     int
}

type HUIANODE__ struct {
	UNUSED int
}

type HUIAPATTERNOBJECT__ struct {
	UNUSED int
}

type HUIATEXTRANGE__ struct {
	UNUSED int
}

type HUIAEVENT__ struct {
	UNUSED int
}

type UiaCondition struct {
	ConditionType int
}

type UiaPropertyCondition struct {
	ConditionType int
	PropertyId    int
	Value         int
	Flags         int
}

type UiaAndOrCondition struct {
	ConditionType int
	PPCONDITIONS  int
	CCONDITIONS   int
}

type UiaNotCondition struct {
	ConditionType int
	PCONDITION    int
}

type UiaCacheRequest struct {
	PVIEWCONDITION        int
	Scope                 int
	PPROPERTIES           int
	CPROPERTIES           int
	PPATTERNS             int
	CPATTERNS             int
	AUTOMATIONELEMENTMODE int
}

type UiaFindParams struct {
	MaxDepth       int
	FindFirst      int
	ExcludeRoot    int
	PFINDCONDITION int
}

type UiaEventArgs struct {
	Type    int
	EventId int
}

type UiaPropertyChangedEventArgs struct {
	Type       int
	EventId    int
	PropertyId int
	OldValue   int
	NewValue   int
}

type UiaStructureChangedEventArgs struct {
	Type                int
	EventId             int
	StructureChangeType int
	PRUNTIMEID          int
	CRUNTIMEIDLEN       int
}

type UiaTextEditTextChangedEventArgs struct {
	Type               int
	EventId            int
	TextEditChangeType int
	PTEXTCHANGE        int
}

type UiaChangesEventArgs struct {
	Type         int
	EventId      int
	EventIdCount int
	PUIACHANGES  int
}

type UiaAsyncContentLoadedEventArgs struct {
	Type                    int
	EventId                 int
	AsyncContentLoadedState int
	PercentComplete         int
}

type UiaWindowClosedEventArgs struct {
	Type          int
	EventId       int
	PRUNTIMEID    int
	CRUNTIMEIDLEN int
}

type CUIAutomation struct {
}

type CUIAutomation8 struct {
}

type ExtendedProperty struct {
	PropertyName  int
	PropertyValue int
}
