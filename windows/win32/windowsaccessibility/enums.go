// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsaccessibility implements the Windows.Win32.WindowsAccessibility namespace.
package windowsaccessibility

type AnnoScope int32

const (
	ANNO_THIS = 0
	ANNO_CONTAINER = 1
)

type NavigateDirection int32

const (
	NavigateDirection_Parent = 0
	NavigateDirection_NextSibling = 1
	NavigateDirection_PreviousSibling = 2
	NavigateDirection_FirstChild = 3
	NavigateDirection_LastChild = 4
)

type ProviderOptions int32

const (
	ProviderOptions_ClientSideProvider = 1
	ProviderOptions_ServerSideProvider = 2
	ProviderOptions_NonClientAreaProvider = 4
	ProviderOptions_OverrideProvider = 8
	ProviderOptions_ProviderOwnsSetFocus = 16
	ProviderOptions_UseComThreading = 32
	ProviderOptions_RefuseNonClientSupport = 64
	ProviderOptions_HasNativeIAccessible = 128
	ProviderOptions_UseClientCoordinates = 256
)

type StructureChangeType int32

const (
	StructureChangeType_ChildAdded = 0
	StructureChangeType_ChildRemoved = 1
	StructureChangeType_ChildrenInvalidated = 2
	StructureChangeType_ChildrenBulkAdded = 3
	StructureChangeType_ChildrenBulkRemoved = 4
	StructureChangeType_ChildrenReordered = 5
)

type TextEditChangeType int32

const (
	TextEditChangeType_None = 0
	TextEditChangeType_AutoCorrect = 1
	TextEditChangeType_Composition = 2
	TextEditChangeType_CompositionFinalized = 3
	TextEditChangeType_AutoComplete = 4
)

type OrientationType int32

const (
	OrientationType_None = 0
	OrientationType_Horizontal = 1
	OrientationType_Vertical = 2
)

type DockPosition int32

const (
	DockPosition_Top = 0
	DockPosition_Left = 1
	DockPosition_Bottom = 2
	DockPosition_Right = 3
	DockPosition_Fill = 4
	DockPosition_None = 5
)

type ExpandCollapseState int32

const (
	ExpandCollapseState_Collapsed = 0
	ExpandCollapseState_Expanded = 1
	ExpandCollapseState_PartiallyExpanded = 2
	ExpandCollapseState_LeafNode = 3
)

type ScrollAmount int32

const (
	ScrollAmount_LargeDecrement = 0
	ScrollAmount_SmallDecrement = 1
	ScrollAmount_NoAmount = 2
	ScrollAmount_LargeIncrement = 3
	ScrollAmount_SmallIncrement = 4
)

type RowOrColumnMajor int32

const (
	RowOrColumnMajor_RowMajor = 0
	RowOrColumnMajor_ColumnMajor = 1
	RowOrColumnMajor_Indeterminate = 2
)

type ToggleState int32

const (
	ToggleState_Off = 0
	ToggleState_On = 1
	ToggleState_Indeterminate = 2
)

type WindowVisualState int32

const (
	WindowVisualState_Normal = 0
	WindowVisualState_Maximized = 1
	WindowVisualState_Minimized = 2
)

type SynchronizedInputType int32

const (
	SynchronizedInputType_KeyUp = 1
	SynchronizedInputType_KeyDown = 2
	SynchronizedInputType_LeftMouseUp = 4
	SynchronizedInputType_LeftMouseDown = 8
	SynchronizedInputType_RightMouseUp = 16
	SynchronizedInputType_RightMouseDown = 32
)

type WindowInteractionState int32

const (
	WindowInteractionState_Running = 0
	WindowInteractionState_Closing = 1
	WindowInteractionState_ReadyForUserInteraction = 2
	WindowInteractionState_BlockedByModalWindow = 3
	WindowInteractionState_NotResponding = 4
)

type SayAsInterpretAs int32

const (
	SayAsInterpretAs_None = 0
	SayAsInterpretAs_Spell = 1
	SayAsInterpretAs_Cardinal = 2
	SayAsInterpretAs_Ordinal = 3
	SayAsInterpretAs_Number = 4
	SayAsInterpretAs_Date = 5
	SayAsInterpretAs_Time = 6
	SayAsInterpretAs_Telephone = 7
	SayAsInterpretAs_Currency = 8
	SayAsInterpretAs_Net = 9
	SayAsInterpretAs_Url = 10
	SayAsInterpretAs_Address = 11
	SayAsInterpretAs_Alphanumeric = 12
	SayAsInterpretAs_Name = 13
	SayAsInterpretAs_Media = 14
	SayAsInterpretAs_Date_MonthDayYear = 15
	SayAsInterpretAs_Date_DayMonthYear = 16
	SayAsInterpretAs_Date_YearMonthDay = 17
	SayAsInterpretAs_Date_YearMonth = 18
	SayAsInterpretAs_Date_MonthYear = 19
	SayAsInterpretAs_Date_DayMonth = 20
	SayAsInterpretAs_Date_MonthDay = 21
	SayAsInterpretAs_Date_Year = 22
	SayAsInterpretAs_Time_HoursMinutesSeconds12 = 23
	SayAsInterpretAs_Time_HoursMinutes12 = 24
	SayAsInterpretAs_Time_HoursMinutesSeconds24 = 25
	SayAsInterpretAs_Time_HoursMinutes24 = 26
)

type TextUnit int32

const (
	TextUnit_Character = 0
	TextUnit_Format = 1
	TextUnit_Word = 2
	TextUnit_Line = 3
	TextUnit_Paragraph = 4
	TextUnit_Page = 5
	TextUnit_Document = 6
)

type TextPatternRangeEndpoint int32

const (
	TextPatternRangeEndpoint_Start = 0
	TextPatternRangeEndpoint_End = 1
)

type SupportedTextSelection int32

const (
	SupportedTextSelection_None = 0
	SupportedTextSelection_Single = 1
	SupportedTextSelection_Multiple = 2
)

type LiveSetting int32

const (
	Off = 0
	Polite = 1
	Assertive = 2
)

type ActiveEnd int32

const (
	ActiveEnd_None = 0
	ActiveEnd_Start = 1
	ActiveEnd_End = 2
)

type CaretPosition int32

const (
	CaretPosition_Unknown = 0
	CaretPosition_EndOfLine = 1
	CaretPosition_BeginningOfLine = 2
)

type CaretBidiMode int32

const (
	CaretBidiMode_LTR = 0
	CaretBidiMode_RTL = 1
)

type ZoomUnit int32

const (
	ZoomUnit_NoAmount = 0
	ZoomUnit_LargeDecrement = 1
	ZoomUnit_SmallDecrement = 2
	ZoomUnit_LargeIncrement = 3
	ZoomUnit_SmallIncrement = 4
)

type AnimationStyle int32

const (
	AnimationStyle_None = 0
	AnimationStyle_LasVegasLights = 1
	AnimationStyle_BlinkingBackground = 2
	AnimationStyle_SparkleText = 3
	AnimationStyle_MarchingBlackAnts = 4
	AnimationStyle_MarchingRedAnts = 5
	AnimationStyle_Shimmer = 6
	AnimationStyle_Other = -1
)

type BulletStyle int32

const (
	BulletStyle_None = 0
	BulletStyle_HollowRoundBullet = 1
	BulletStyle_FilledRoundBullet = 2
	BulletStyle_HollowSquareBullet = 3
	BulletStyle_FilledSquareBullet = 4
	BulletStyle_DashBullet = 5
	BulletStyle_Other = -1
)

type CapStyle int32

const (
	CapStyle_None = 0
	CapStyle_SmallCap = 1
	CapStyle_AllCap = 2
	CapStyle_AllPetiteCaps = 3
	CapStyle_PetiteCaps = 4
	CapStyle_Unicase = 5
	CapStyle_Titling = 6
	CapStyle_Other = -1
)

type FillType int32

const (
	FillType_None = 0
	FillType_Color = 1
	FillType_Gradient = 2
	FillType_Picture = 3
	FillType_Pattern = 4
)

type FlowDirections int32

const (
	FlowDirections_Default = 0
	FlowDirections_RightToLeft = 1
	FlowDirections_BottomToTop = 2
	FlowDirections_Vertical = 4
)

type HorizontalTextAlignment int32

const (
	HorizontalTextAlignment_Left = 0
	HorizontalTextAlignment_Centered = 1
	HorizontalTextAlignment_Right = 2
	HorizontalTextAlignment_Justified = 3
)

type OutlineStyles int32

const (
	OutlineStyles_None = 0
	OutlineStyles_Outline = 1
	OutlineStyles_Shadow = 2
	OutlineStyles_Engraved = 4
	OutlineStyles_Embossed = 8
)

type TextDecorationLineStyle int32

const (
	TextDecorationLineStyle_None = 0
	TextDecorationLineStyle_Single = 1
	TextDecorationLineStyle_WordsOnly = 2
	TextDecorationLineStyle_Double = 3
	TextDecorationLineStyle_Dot = 4
	TextDecorationLineStyle_Dash = 5
	TextDecorationLineStyle_DashDot = 6
	TextDecorationLineStyle_DashDotDot = 7
	TextDecorationLineStyle_Wavy = 8
	TextDecorationLineStyle_ThickSingle = 9
	TextDecorationLineStyle_DoubleWavy = 11
	TextDecorationLineStyle_ThickWavy = 12
	TextDecorationLineStyle_LongDash = 13
	TextDecorationLineStyle_ThickDash = 14
	TextDecorationLineStyle_ThickDashDot = 15
	TextDecorationLineStyle_ThickDashDotDot = 16
	TextDecorationLineStyle_ThickDot = 17
	TextDecorationLineStyle_ThickLongDash = 18
	TextDecorationLineStyle_Other = -1
)

type VisualEffects int32

const (
	VisualEffects_None = 0
	VisualEffects_Shadow = 1
	VisualEffects_Reflection = 2
	VisualEffects_Glow = 4
	VisualEffects_SoftEdges = 8
	VisualEffects_Bevel = 16
)

type NotificationProcessing int32

const (
	NotificationProcessing_ImportantAll = 0
	NotificationProcessing_ImportantMostRecent = 1
	NotificationProcessing_All = 2
	NotificationProcessing_MostRecent = 3
	NotificationProcessing_CurrentThenMostRecent = 4
)

type NotificationKind int32

const (
	NotificationKind_ItemAdded = 0
	NotificationKind_ItemRemoved = 1
	NotificationKind_ActionCompleted = 2
	NotificationKind_ActionAborted = 3
	NotificationKind_Other = 4
)

type UIAutomationType int32

const (
	UIAutomationType_Int = 1
	UIAutomationType_Bool = 2
	UIAutomationType_String = 3
	UIAutomationType_Double = 4
	UIAutomationType_Point = 5
	UIAutomationType_Rect = 6
	UIAutomationType_Element = 7
	UIAutomationType_Array = 65536
	UIAutomationType_Out = 131072
	UIAutomationType_IntArray = 65537
	UIAutomationType_BoolArray = 65538
	UIAutomationType_StringArray = 65539
	UIAutomationType_DoubleArray = 65540
	UIAutomationType_PointArray = 65541
	UIAutomationType_RectArray = 65542
	UIAutomationType_ElementArray = 65543
	UIAutomationType_OutInt = 131073
	UIAutomationType_OutBool = 131074
	UIAutomationType_OutString = 131075
	UIAutomationType_OutDouble = 131076
	UIAutomationType_OutPoint = 131077
	UIAutomationType_OutRect = 131078
	UIAutomationType_OutElement = 131079
	UIAutomationType_OutIntArray = 196609
	UIAutomationType_OutBoolArray = 196610
	UIAutomationType_OutStringArray = 196611
	UIAutomationType_OutDoubleArray = 196612
	UIAutomationType_OutPointArray = 196613
	UIAutomationType_OutRectArray = 196614
	UIAutomationType_OutElementArray = 196615
)

type TreeScope int32

const (
	TreeScope_None = 0
	TreeScope_Element = 1
	TreeScope_Children = 2
	TreeScope_Descendants = 4
	TreeScope_Parent = 8
	TreeScope_Ancestors = 16
	TreeScope_Subtree = 7
)

type ConditionType int32

const (
	ConditionType_True = 0
	ConditionType_False = 1
	ConditionType_Property = 2
	ConditionType_And = 3
	ConditionType_Or = 4
	ConditionType_Not = 5
)

type PropertyConditionFlags int32

const (
	PropertyConditionFlags_None = 0
	PropertyConditionFlags_IgnoreCase = 1
	PropertyConditionFlags_MatchSubstring = 2
)

type AutomationElementMode int32

const (
	AutomationElementMode_None = 0
	AutomationElementMode_Full = 1
)

type NormalizeState int32

const (
	NormalizeState_None = 0
	NormalizeState_View = 1
	NormalizeState_Custom = 2
)

type TreeTraversalOptions int32

const (
	TreeTraversalOptions_Default = 0
	TreeTraversalOptions_PostOrder = 1
	TreeTraversalOptions_LastToFirstOrder = 2
)

type ProviderType int32

const (
	ProviderType_BaseHwnd = 0
	ProviderType_Proxy = 1
	ProviderType_NonClientArea = 2
)

type AutomationIdentifierType int32

const (
	AutomationIdentifierType_Property = 0
	AutomationIdentifierType_Pattern = 1
	AutomationIdentifierType_Event = 2
	AutomationIdentifierType_ControlType = 3
	AutomationIdentifierType_TextAttribute = 4
	AutomationIdentifierType_LandmarkType = 5
	AutomationIdentifierType_Annotation = 6
	AutomationIdentifierType_Changes = 7
	AutomationIdentifierType_Style = 8
)

type EventArgsType int32

const (
	EventArgsType_Simple = 0
	EventArgsType_PropertyChanged = 1
	EventArgsType_StructureChanged = 2
	EventArgsType_AsyncContentLoaded = 3
	EventArgsType_WindowClosed = 4
	EventArgsType_TextEditTextChanged = 5
	EventArgsType_Changes = 6
	EventArgsType_Notification = 7
	EventArgsType_ActiveTextPositionChanged = 8
	EventArgsType_StructuredMarkup = 9
)

type AsyncContentLoadedState int32

const (
	AsyncContentLoadedState_Beginning = 0
	AsyncContentLoadedState_Progress = 1
	AsyncContentLoadedState_Completed = 2
)
