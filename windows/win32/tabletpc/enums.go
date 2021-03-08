// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package tabletpc implements the Windows.Win32.TabletPC namespace.
package tabletpc

type PROPERTY_UNITS int32

const (
	PROPERTY_UNITS_DEFAULT = 0
	PROPERTY_UNITS_INCHES = 1
	PROPERTY_UNITS_CENTIMETERS = 2
	PROPERTY_UNITS_DEGREES = 3
	PROPERTY_UNITS_RADIANS = 4
	PROPERTY_UNITS_SECONDS = 5
	PROPERTY_UNITS_POUNDS = 6
	PROPERTY_UNITS_GRAMS = 7
	PROPERTY_UNITS_SILINEAR = 8
	PROPERTY_UNITS_SIROTATION = 9
	PROPERTY_UNITS_ENGLINEAR = 10
	PROPERTY_UNITS_ENGROTATION = 11
	PROPERTY_UNITS_SLUGS = 12
	PROPERTY_UNITS_KELVIN = 13
	PROPERTY_UNITS_FAHRENHEIT = 14
	PROPERTY_UNITS_AMPERE = 15
	PROPERTY_UNITS_CANDELA = 16
)

type enumINKMETRIC_FLAGS int32

const (
	IMF_FONT_SELECTED_IN_HDC = 1
	IMF_ITALIC = 2
	IMF_BOLD = 4
)

type enumGetCandidateFlags int32

const (
	TCF_ALLOW_RECOGNITION = 1
	TCF_FORCE_RECOGNITION = 2
)

type InkSelectionConstants int32

const (
	ISC_FirstElement = 0
	ISC_AllElements = -1
)

type InkBoundingBoxMode int32

const (
	IBBM_Default = 0
	IBBM_NoCurveFit = 1
	IBBM_CurveFit = 2
	IBBM_PointsOnly = 3
	IBBM_Union = 4
)

type InkExtractFlags int32

const (
	IEF_CopyFromOriginal = 0
	IEF_RemoveFromOriginal = 1
	IEF_Default = 1
)

type InkPersistenceFormat int32

const (
	IPF_InkSerializedFormat = 0
	IPF_Base64InkSerializedFormat = 1
	IPF_GIF = 2
	IPF_Base64GIF = 3
)

type InkPersistenceCompressionMode int32

const (
	IPCM_Default = 0
	IPCM_MaximumCompression = 1
	IPCM_NoCompression = 2
)

type InkPenTip int32

const (
	IPT_Ball = 0
	IPT_Rectangle = 1
)

type InkRasterOperation int32

const (
	IRO_Black = 1
	IRO_NotMergePen = 2
	IRO_MaskNotPen = 3
	IRO_NotCopyPen = 4
	IRO_MaskPenNot = 5
	IRO_Not = 6
	IRO_XOrPen = 7
	IRO_NotMaskPen = 8
	IRO_MaskPen = 9
	IRO_NotXOrPen = 10
	IRO_NoOperation = 11
	IRO_MergeNotPen = 12
	IRO_CopyPen = 13
	IRO_MergePenNot = 14
	IRO_MergePen = 15
	IRO_White = 16
)

type InkMousePointer int32

const (
	IMP_Default = 0
	IMP_Arrow = 1
	IMP_Crosshair = 2
	IMP_Ibeam = 3
	IMP_SizeNESW = 4
	IMP_SizeNS = 5
	IMP_SizeNWSE = 6
	IMP_SizeWE = 7
	IMP_UpArrow = 8
	IMP_Hourglass = 9
	IMP_NoDrop = 10
	IMP_ArrowHourglass = 11
	IMP_ArrowQuestion = 12
	IMP_SizeAll = 13
	IMP_Hand = 14
	IMP_Custom = 99
)

type InkClipboardModes int32

const (
	ICB_Copy = 0
	ICB_Cut = 1
	ICB_ExtractOnly = 48
	ICB_DelayedCopy = 32
	ICB_Default = 0
)

type InkClipboardFormats int32

const (
	ICF_None = 0
	ICF_InkSerializedFormat = 1
	ICF_SketchInk = 2
	ICF_TextInk = 6
	ICF_EnhancedMetafile = 8
	ICF_Metafile = 32
	ICF_Bitmap = 64
	ICF_PasteMask = 7
	ICF_CopyMask = 127
	ICF_Default = 127
)

type SelectionHitResult int32

const (
	SHR_None = 0
	SHR_NW = 1
	SHR_SE = 2
	SHR_NE = 3
	SHR_SW = 4
	SHR_E = 5
	SHR_W = 6
	SHR_N = 7
	SHR_S = 8
	SHR_Selection = 9
)

type InkRecognitionStatus int32

const (
	IRS_NoError = 0
	IRS_Interrupted = 1
	IRS_ProcessFailed = 2
	IRS_InkAddedFailed = 4
	IRS_SetAutoCompletionModeFailed = 8
	IRS_SetStrokesFailed = 16
	IRS_SetGuideFailed = 32
	IRS_SetFlagsFailed = 64
	IRS_SetFactoidFailed = 128
	IRS_SetPrefixSuffixFailed = 256
	IRS_SetWordListFailed = 512
)

type DISPID_InkRectangle int32

const (
	DISPID_IRTop = 1
	DISPID_IRLeft = 2
	DISPID_IRBottom = 3
	DISPID_IRRight = 4
	DISPID_IRGetRectangle = 5
	DISPID_IRSetRectangle = 6
	DISPID_IRData = 7
)

type DISPID_InkExtendedProperty int32

const (
	DISPID_IEPGuid = 1
	DISPID_IEPData = 2
)

type DISPID_InkExtendedProperties int32

const (
	DISPID_IEPs_NewEnum = -4
	DISPID_IEPsItem = 0
	DISPID_IEPsCount = 1
	DISPID_IEPsAdd = 2
	DISPID_IEPsRemove = 3
	DISPID_IEPsClear = 4
	DISPID_IEPsDoesPropertyExist = 5
)

type DISPID_InkDrawingAttributes int32

const (
	DISPID_DAHeight = 1
	DISPID_DAColor = 2
	DISPID_DAWidth = 3
	DISPID_DAFitToCurve = 4
	DISPID_DAIgnorePressure = 5
	DISPID_DAAntiAliased = 6
	DISPID_DATransparency = 7
	DISPID_DARasterOperation = 8
	DISPID_DAPenTip = 9
	DISPID_DAClone = 10
	DISPID_DAExtendedProperties = 11
)

type DISPID_InkTransform int32

const (
	DISPID_ITReset = 1
	DISPID_ITTranslate = 2
	DISPID_ITRotate = 3
	DISPID_ITReflect = 4
	DISPID_ITShear = 5
	DISPID_ITScale = 6
	DISPID_ITeM11 = 7
	DISPID_ITeM12 = 8
	DISPID_ITeM21 = 9
	DISPID_ITeM22 = 10
	DISPID_ITeDx = 11
	DISPID_ITeDy = 12
	DISPID_ITGetTransform = 13
	DISPID_ITSetTransform = 14
	DISPID_ITData = 15
)

type InkApplicationGesture int32

const (
	IAG_AllGestures = 0
	IAG_NoGesture = 61440
	IAG_Scratchout = 61441
	IAG_Triangle = 61442
	IAG_Square = 61443
	IAG_Star = 61444
	IAG_Check = 61445
	IAG_Curlicue = 61456
	IAG_DoubleCurlicue = 61457
	IAG_Circle = 61472
	IAG_DoubleCircle = 61473
	IAG_SemiCircleLeft = 61480
	IAG_SemiCircleRight = 61481
	IAG_ChevronUp = 61488
	IAG_ChevronDown = 61489
	IAG_ChevronLeft = 61490
	IAG_ChevronRight = 61491
	IAG_ArrowUp = 61496
	IAG_ArrowDown = 61497
	IAG_ArrowLeft = 61498
	IAG_ArrowRight = 61499
	IAG_Up = 61528
	IAG_Down = 61529
	IAG_Left = 61530
	IAG_Right = 61531
	IAG_UpDown = 61536
	IAG_DownUp = 61537
	IAG_LeftRight = 61538
	IAG_RightLeft = 61539
	IAG_UpLeftLong = 61540
	IAG_UpRightLong = 61541
	IAG_DownLeftLong = 61542
	IAG_DownRightLong = 61543
	IAG_UpLeft = 61544
	IAG_UpRight = 61545
	IAG_DownLeft = 61546
	IAG_DownRight = 61547
	IAG_LeftUp = 61548
	IAG_LeftDown = 61549
	IAG_RightUp = 61550
	IAG_RightDown = 61551
	IAG_Exclamation = 61604
	IAG_Tap = 61680
	IAG_DoubleTap = 61681
)

type InkSystemGesture int32

const (
	ISG_Tap = 16
	ISG_DoubleTap = 17
	ISG_RightTap = 18
	ISG_Drag = 19
	ISG_RightDrag = 20
	ISG_HoldEnter = 21
	ISG_HoldLeave = 22
	ISG_HoverEnter = 23
	ISG_HoverLeave = 24
	ISG_Flick = 31
)

type InkRecognitionConfidence int32

const (
	IRC_Strong = 0
	IRC_Intermediate = 1
	IRC_Poor = 2
)

type DISPID_InkGesture int32

const (
	DISPID_IGId = 0
	DISPID_IGGetHotPoint = 1
	DISPID_IGConfidence = 2
)

type DISPID_InkCursor int32

const (
	DISPID_ICsrName = 0
	DISPID_ICsrId = 1
	DISPID_ICsrDrawingAttributes = 2
	DISPID_ICsrButtons = 3
	DISPID_ICsrInverted = 4
	DISPID_ICsrTablet = 5
)

type DISPID_InkCursors int32

const (
	DISPID_ICs_NewEnum = -4
	DISPID_ICsItem = 0
	DISPID_ICsCount = 1
)

type InkCursorButtonState int32

const (
	ICBS_Unavailable = 0
	ICBS_Up = 1
	ICBS_Down = 2
)

type DISPID_InkCursorButton int32

const (
	DISPID_ICBName = 0
	DISPID_ICBId = 1
	DISPID_ICBState = 2
)

type DISPID_InkCursorButtons int32

const (
	DISPID_ICBs_NewEnum = -4
	DISPID_ICBsItem = 0
	DISPID_ICBsCount = 1
)

type TabletHardwareCapabilities int32

const (
	THWC_Integrated = 1
	THWC_CursorMustTouch = 2
	THWC_HardProximity = 4
	THWC_CursorsHavePhysicalIds = 8
)

type TabletPropertyMetricUnit int32

const (
	TPMU_Default = 0
	TPMU_Inches = 1
	TPMU_Centimeters = 2
	TPMU_Degrees = 3
	TPMU_Radians = 4
	TPMU_Seconds = 5
	TPMU_Pounds = 6
	TPMU_Grams = 7
)

type DISPID_InkTablet int32

const (
	DISPID_ITName = 0
	DISPID_ITPlugAndPlayId = 1
	DISPID_ITPropertyMetrics = 2
	DISPID_ITIsPacketPropertySupported = 3
	DISPID_ITMaximumInputRectangle = 4
	DISPID_ITHardwareCapabilities = 5
)

type TabletDeviceKind int32

const (
	TDK_Mouse = 0
	TDK_Pen = 1
	TDK_Touch = 2
)

type DISPID_InkTablet2 int32

const (
	DISPID_IT2DeviceKind = 0
)

type DISPID_InkTablet3 int32

const (
	DISPID_IT3IsMultiTouch = 0
	DISPID_IT3MaximumCursors = 1
)

type DISPID_InkTablets int32

const (
	DISPID_ITs_NewEnum = -4
	DISPID_ITsItem = 0
	DISPID_ITsDefaultTablet = 1
	DISPID_ITsCount = 2
	DISPID_ITsIsPacketPropertySupported = 3
)

type DISPID_InkStrokeDisp int32

const (
	DISPID_ISDInkIndex = 1
	DISPID_ISDID = 2
	DISPID_ISDGetBoundingBox = 3
	DISPID_ISDDrawingAttributes = 4
	DISPID_ISDFindIntersections = 5
	DISPID_ISDGetRectangleIntersections = 6
	DISPID_ISDClip = 7
	DISPID_ISDHitTestCircle = 8
	DISPID_ISDNearestPoint = 9
	DISPID_ISDSplit = 10
	DISPID_ISDExtendedProperties = 11
	DISPID_ISDInk = 12
	DISPID_ISDBezierPoints = 13
	DISPID_ISDPolylineCusps = 14
	DISPID_ISDBezierCusps = 15
	DISPID_ISDSelfIntersections = 16
	DISPID_ISDPacketCount = 17
	DISPID_ISDPacketSize = 18
	DISPID_ISDPacketDescription = 19
	DISPID_ISDDeleted = 20
	DISPID_ISDGetPacketDescriptionPropertyMetrics = 21
	DISPID_ISDGetPoints = 22
	DISPID_ISDSetPoints = 23
	DISPID_ISDGetPacketData = 24
	DISPID_ISDGetPacketValuesByProperty = 25
	DISPID_ISDSetPacketValuesByProperty = 26
	DISPID_ISDGetFlattenedBezierPoints = 27
	DISPID_ISDScaleToRectangle = 28
	DISPID_ISDTransform = 29
	DISPID_ISDMove = 30
	DISPID_ISDRotate = 31
	DISPID_ISDShear = 32
	DISPID_ISDScale = 33
)

type DISPID_InkStrokes int32

const (
	DISPID_ISs_NewEnum = -4
	DISPID_ISsItem = 0
	DISPID_ISsCount = 1
	DISPID_ISsValid = 2
	DISPID_ISsInk = 3
	DISPID_ISsAdd = 4
	DISPID_ISsAddStrokes = 5
	DISPID_ISsRemove = 6
	DISPID_ISsRemoveStrokes = 7
	DISPID_ISsToString = 8
	DISPID_ISsModifyDrawingAttributes = 9
	DISPID_ISsGetBoundingBox = 10
	DISPID_ISsScaleToRectangle = 11
	DISPID_ISsTransform = 12
	DISPID_ISsMove = 13
	DISPID_ISsRotate = 14
	DISPID_ISsShear = 15
	DISPID_ISsScale = 16
	DISPID_ISsClip = 17
	DISPID_ISsRecognitionResult = 18
	DISPID_ISsRemoveRecognitionResult = 19
)

type DISPID_InkCustomStrokes int32

const (
	DISPID_ICSs_NewEnum = -4
	DISPID_ICSsItem = 0
	DISPID_ICSsCount = 1
	DISPID_ICSsAdd = 2
	DISPID_ICSsRemove = 3
	DISPID_ICSsClear = 4
)

type DISPID_StrokeEvent int32

const (
	DISPID_SEStrokesAdded = 1
	DISPID_SEStrokesRemoved = 2
)

type DISPID_Ink int32

const (
	DISPID_IStrokes = 1
	DISPID_IExtendedProperties = 2
	DISPID_IGetBoundingBox = 3
	DISPID_IDeleteStrokes = 4
	DISPID_IDeleteStroke = 5
	DISPID_IExtractStrokes = 6
	DISPID_IExtractWithRectangle = 7
	DISPID_IDirty = 8
	DISPID_ICustomStrokes = 9
	DISPID_IClone = 10
	DISPID_IHitTestCircle = 11
	DISPID_IHitTestWithRectangle = 12
	DISPID_IHitTestWithLasso = 13
	DISPID_INearestPoint = 14
	DISPID_ICreateStrokes = 15
	DISPID_ICreateStroke = 16
	DISPID_IAddStrokesAtRectangle = 17
	DISPID_IClip = 18
	DISPID_ISave = 19
	DISPID_ILoad = 20
	DISPID_ICreateStrokeFromPoints = 21
	DISPID_IClipboardCopyWithRectangle = 22
	DISPID_IClipboardCopy = 23
	DISPID_ICanPaste = 24
	DISPID_IClipboardPaste = 25
)

type DISPID_InkEvent int32

const (
	DISPID_IEInkAdded = 1
	DISPID_IEInkDeleted = 2
)

type DISPID_InkRenderer int32

const (
	DISPID_IRGetViewTransform = 1
	DISPID_IRSetViewTransform = 2
	DISPID_IRGetObjectTransform = 3
	DISPID_IRSetObjectTransform = 4
	DISPID_IRDraw = 5
	DISPID_IRDrawStroke = 6
	DISPID_IRPixelToInkSpace = 7
	DISPID_IRInkSpaceToPixel = 8
	DISPID_IRPixelToInkSpaceFromPoints = 9
	DISPID_IRInkSpaceToPixelFromPoints = 10
	DISPID_IRMeasure = 11
	DISPID_IRMeasureStroke = 12
	DISPID_IRMove = 13
	DISPID_IRRotate = 14
	DISPID_IRScale = 15
)

type InkCollectorEventInterest int32

const (
	ICEI_DefaultEvents = -1
	ICEI_CursorDown = 0
	ICEI_Stroke = 1
	ICEI_NewPackets = 2
	ICEI_NewInAirPackets = 3
	ICEI_CursorButtonDown = 4
	ICEI_CursorButtonUp = 5
	ICEI_CursorInRange = 6
	ICEI_CursorOutOfRange = 7
	ICEI_SystemGesture = 8
	ICEI_TabletAdded = 9
	ICEI_TabletRemoved = 10
	ICEI_MouseDown = 11
	ICEI_MouseMove = 12
	ICEI_MouseUp = 13
	ICEI_MouseWheel = 14
	ICEI_DblClick = 15
	ICEI_AllEvents = 16
)

type InkMouseButton int32

const (
	IMF_Left = 1
	IMF_Right = 2
	IMF_Middle = 4
)

type InkShiftKeyModifierFlags int32

const (
	IKM_Shift = 1
	IKM_Control = 2
	IKM_Alt = 4
)

type DISPID_InkCollectorEvent int32

const (
	DISPID_ICEStroke = 1
	DISPID_ICECursorDown = 2
	DISPID_ICENewPackets = 3
	DISPID_ICENewInAirPackets = 4
	DISPID_ICECursorButtonDown = 5
	DISPID_ICECursorButtonUp = 6
	DISPID_ICECursorInRange = 7
	DISPID_ICECursorOutOfRange = 8
	DISPID_ICESystemGesture = 9
	DISPID_ICEGesture = 10
	DISPID_ICETabletAdded = 11
	DISPID_ICETabletRemoved = 12
	DISPID_IOEPainting = 13
	DISPID_IOEPainted = 14
	DISPID_IOESelectionChanging = 15
	DISPID_IOESelectionChanged = 16
	DISPID_IOESelectionMoving = 17
	DISPID_IOESelectionMoved = 18
	DISPID_IOESelectionResizing = 19
	DISPID_IOESelectionResized = 20
	DISPID_IOEStrokesDeleting = 21
	DISPID_IOEStrokesDeleted = 22
	DISPID_IPEChangeUICues = 23
	DISPID_IPEClick = 24
	DISPID_IPEDblClick = 25
	DISPID_IPEInvalidated = 26
	DISPID_IPEMouseDown = 27
	DISPID_IPEMouseEnter = 28
	DISPID_IPEMouseHover = 29
	DISPID_IPEMouseLeave = 30
	DISPID_IPEMouseMove = 31
	DISPID_IPEMouseUp = 32
	DISPID_IPEMouseWheel = 33
	DISPID_IPESizeModeChanged = 34
	DISPID_IPEStyleChanged = 35
	DISPID_IPESystemColorsChanged = 36
	DISPID_IPEKeyDown = 37
	DISPID_IPEKeyPress = 38
	DISPID_IPEKeyUp = 39
	DISPID_IPEResize = 40
	DISPID_IPESizeChanged = 41
)

type InkOverlayEditingMode int32

const (
	IOEM_Ink = 0
	IOEM_Delete = 1
	IOEM_Select = 2
)

type InkOverlayAttachMode int32

const (
	IOAM_Behind = 0
	IOAM_InFront = 1
)

type InkPictureSizeMode int32

const (
	IPSM_AutoSize = 0
	IPSM_CenterImage = 1
	IPSM_Normal = 2
	IPSM_StretchImage = 3
)

type InkOverlayEraserMode int32

const (
	IOERM_StrokeErase = 0
	IOERM_PointErase = 1
)

type InkCollectionMode int32

const (
	ICM_InkOnly = 0
	ICM_GestureOnly = 1
	ICM_InkAndGesture = 2
)

type DISPID_InkCollector int32

const (
	DISPID_ICEnabled = 1
	DISPID_ICHwnd = 2
	DISPID_ICPaint = 3
	DISPID_ICText = 4
	DISPID_ICDefaultDrawingAttributes = 5
	DISPID_ICRenderer = 6
	DISPID_ICInk = 7
	DISPID_ICAutoRedraw = 8
	DISPID_ICCollectingInk = 9
	DISPID_ICSetEventInterest = 10
	DISPID_ICGetEventInterest = 11
	DISPID_IOEditingMode = 12
	DISPID_IOSelection = 13
	DISPID_IOAttachMode = 14
	DISPID_IOHitTestSelection = 15
	DISPID_IODraw = 16
	DISPID_IPPicture = 17
	DISPID_IPSizeMode = 18
	DISPID_IPBackColor = 19
	DISPID_ICCursors = 20
	DISPID_ICMarginX = 21
	DISPID_ICMarginY = 22
	DISPID_ICSetWindowInputRectangle = 23
	DISPID_ICGetWindowInputRectangle = 24
	DISPID_ICTablet = 25
	DISPID_ICSetAllTabletsMode = 26
	DISPID_ICSetSingleTabletIntegratedMode = 27
	DISPID_ICCollectionMode = 28
	DISPID_ICSetGestureStatus = 29
	DISPID_ICGetGestureStatus = 30
	DISPID_ICDynamicRendering = 31
	DISPID_ICDesiredPacketDescription = 32
	DISPID_IOEraserMode = 33
	DISPID_IOEraserWidth = 34
	DISPID_ICMouseIcon = 35
	DISPID_ICMousePointer = 36
	DISPID_IPInkEnabled = 37
	DISPID_ICSupportHighContrastInk = 38
	DISPID_IOSupportHighContrastSelectionUI = 39
)

type DISPID_InkRecognizer int32

const (
	DISPID_RecoClsid = 1
	DISPID_RecoName = 2
	DISPID_RecoVendor = 3
	DISPID_RecoCapabilities = 4
	DISPID_RecoLanguageID = 5
	DISPID_RecoPreferredPacketDescription = 6
	DISPID_RecoCreateRecognizerContext = 7
	DISPID_RecoSupportedProperties = 8
)

type InkRecognizerCapabilities int32

const (
	IRC_DontCare = 1
	IRC_Object = 2
	IRC_FreeInput = 4
	IRC_LinedInput = 8
	IRC_BoxedInput = 16
	IRC_CharacterAutoCompletionInput = 32
	IRC_RightAndDown = 64
	IRC_LeftAndDown = 128
	IRC_DownAndLeft = 256
	IRC_DownAndRight = 512
	IRC_ArbitraryAngle = 1024
	IRC_Lattice = 2048
	IRC_AdviseInkChange = 4096
	IRC_StrokeReorder = 8192
	IRC_Personalizable = 16384
	IRC_PrefersArbitraryAngle = 32768
	IRC_PrefersParagraphBreaking = 65536
	IRC_PrefersSegmentation = 131072
	IRC_Cursive = 262144
	IRC_TextPrediction = 524288
	IRC_Alpha = 1048576
	IRC_Beta = 2097152
)

type DISPID_InkRecognizer2 int32

const (
	DISPID_RecoId = 0
	DISPID_RecoUnicodeRanges = 1
)

type DISPID_InkRecognizers int32

const (
	DISPID_IRecos_NewEnum = -4
	DISPID_IRecosItem = 0
	DISPID_IRecosCount = 1
	DISPID_IRecosGetDefaultRecognizer = 2
)

type InkRecognizerCharacterAutoCompletionMode int32

const (
	IRCACM_Full = 0
	IRCACM_Prefix = 1
	IRCACM_Random = 2
)

type InkRecognitionModes int32

const (
	IRM_None = 0
	IRM_WordModeOnly = 1
	IRM_Coerce = 2
	IRM_TopInkBreaksOnly = 4
	IRM_PrefixOk = 8
	IRM_LineMode = 16
	IRM_DisablePersonalization = 32
	IRM_AutoSpace = 64
	IRM_Max = 128
)

type DISPID_InkRecognitionEvent int32

const (
	DISPID_IRERecognitionWithAlternates = 1
	DISPID_IRERecognition = 2
)

type DISPID_InkRecoContext int32

const (
	DISPID_IRecoCtx_Strokes = 1
	DISPID_IRecoCtx_CharacterAutoCompletionMode = 2
	DISPID_IRecoCtx_Factoid = 3
	DISPID_IRecoCtx_WordList = 4
	DISPID_IRecoCtx_Recognizer = 5
	DISPID_IRecoCtx_Guide = 6
	DISPID_IRecoCtx_Flags = 7
	DISPID_IRecoCtx_PrefixText = 8
	DISPID_IRecoCtx_SuffixText = 9
	DISPID_IRecoCtx_StopRecognition = 10
	DISPID_IRecoCtx_Clone = 11
	DISPID_IRecoCtx_Recognize = 12
	DISPID_IRecoCtx_StopBackgroundRecognition = 13
	DISPID_IRecoCtx_EndInkInput = 14
	DISPID_IRecoCtx_BackgroundRecognize = 15
	DISPID_IRecoCtx_BackgroundRecognizeWithAlternates = 16
	DISPID_IRecoCtx_IsStringSupported = 17
)

type DISPID_InkRecoContext2 int32

const (
	DISPID_IRecoCtx2_EnabledUnicodeRanges = 0
)

type InkRecognitionAlternatesSelection int32

const (
	IRAS_Start = 0
	IRAS_DefaultCount = 10
	IRAS_All = -1
)

type DISPID_InkRecognitionResult int32

const (
	DISPID_InkRecognitionResult_TopString = 1
	DISPID_InkRecognitionResult_TopAlternate = 2
	DISPID_InkRecognitionResult_Strokes = 3
	DISPID_InkRecognitionResult_TopConfidence = 4
	DISPID_InkRecognitionResult_AlternatesFromSelection = 5
	DISPID_InkRecognitionResult_ModifyTopAlternate = 6
	DISPID_InkRecognitionResult_SetResultOnStrokes = 7
)

type DISPID_InkRecoAlternate int32

const (
	DISPID_InkRecoAlternate_String = 1
	DISPID_InkRecoAlternate_LineNumber = 2
	DISPID_InkRecoAlternate_Baseline = 3
	DISPID_InkRecoAlternate_Midline = 4
	DISPID_InkRecoAlternate_Ascender = 5
	DISPID_InkRecoAlternate_Descender = 6
	DISPID_InkRecoAlternate_Confidence = 7
	DISPID_InkRecoAlternate_Strokes = 8
	DISPID_InkRecoAlternate_GetStrokesFromStrokeRanges = 9
	DISPID_InkRecoAlternate_GetStrokesFromTextRange = 10
	DISPID_InkRecoAlternate_GetTextRangeFromStrokes = 11
	DISPID_InkRecoAlternate_GetPropertyValue = 12
	DISPID_InkRecoAlternate_LineAlternates = 13
	DISPID_InkRecoAlternate_ConfidenceAlternates = 14
	DISPID_InkRecoAlternate_AlternatesWithConstantPropertyValues = 15
)

type DISPID_InkRecognitionAlternates int32

const (
	DISPID_InkRecognitionAlternates_NewEnum = -4
	DISPID_InkRecognitionAlternates_Item = 0
	DISPID_InkRecognitionAlternates_Count = 1
	DISPID_InkRecognitionAlternates_Strokes = 2
)

type DISPID_InkRecognizerGuide int32

const (
	DISPID_IRGWritingBox = 1
	DISPID_IRGDrawnBox = 2
	DISPID_IRGRows = 3
	DISPID_IRGColumns = 4
	DISPID_IRGMidline = 5
	DISPID_IRGGuideData = 6
)

type DISPID_InkWordList int32

const (
	DISPID_InkWordList_AddWord = 0
	DISPID_InkWordList_RemoveWord = 1
	DISPID_InkWordList_Merge = 2
)

type DISPID_InkWordList2 int32

const (
	DISPID_InkWordList2_AddWords = 3
)

type InkDivisionType int32

const (
	IDT_Segment = 0
	IDT_Line = 1
	IDT_Paragraph = 2
	IDT_Drawing = 3
)

type DISPID_InkDivider int32

const (
	DISPID_IInkDivider_Strokes = 1
	DISPID_IInkDivider_RecognizerContext = 2
	DISPID_IInkDivider_LineHeight = 3
	DISPID_IInkDivider_Divide = 4
)

type DISPID_InkDivisionResult int32

const (
	DISPID_IInkDivisionResult_Strokes = 1
	DISPID_IInkDivisionResult_ResultByType = 2
)

type DISPID_InkDivisionUnit int32

const (
	DISPID_IInkDivisionUnit_Strokes = 1
	DISPID_IInkDivisionUnit_DivisionType = 2
	DISPID_IInkDivisionUnit_RecognizedString = 3
	DISPID_IInkDivisionUnit_RotationTransform = 4
)

type DISPID_InkDivisionUnits int32

const (
	DISPID_IInkDivisionUnits_NewEnum = -4
	DISPID_IInkDivisionUnits_Item = 0
	DISPID_IInkDivisionUnits_Count = 1
)

type DISPID_PenInputPanel int32

const (
	DISPID_PIPAttachedEditWindow = 0
	DISPID_PIPFactoid = 1
	DISPID_PIPCurrentPanel = 2
	DISPID_PIPDefaultPanel = 3
	DISPID_PIPVisible = 4
	DISPID_PIPTop = 5
	DISPID_PIPLeft = 6
	DISPID_PIPWidth = 7
	DISPID_PIPHeight = 8
	DISPID_PIPMoveTo = 9
	DISPID_PIPCommitPendingInput = 10
	DISPID_PIPRefresh = 11
	DISPID_PIPBusy = 12
	DISPID_PIPVerticalOffset = 13
	DISPID_PIPHorizontalOffset = 14
	DISPID_PIPEnableTsf = 15
	DISPID_PIPAutoShow = 16
)

type DISPID_PenInputPanelEvents int32

const (
	DISPID_PIPEVisibleChanged = 0
	DISPID_PIPEPanelChanged = 1
	DISPID_PIPEInputFailed = 2
	DISPID_PIPEPanelMoving = 3
)

type VisualState int32

const (
	InPlace = 0
	Floating = 1
	DockedTop = 2
	DockedBottom = 3
	Closed = 4
)

type InteractionMode int32

const (
	InteractionMode_InPlace = 0
	InteractionMode_Floating = 1
	InteractionMode_DockedTop = 2
	InteractionMode_DockedBottom = 3
)

type InPlaceState int32

const (
	InPlaceState_Auto = 0
	InPlaceState_HoverTarget = 1
	InPlaceState_Expanded = 2
)

type PanelInputArea int32

const (
	PanelInputArea_Auto = 0
	PanelInputArea_Keyboard = 1
	PanelInputArea_WritingPad = 2
	PanelInputArea_CharacterPad = 3
)

type CorrectionMode int32

const (
	CorrectionMode_NotVisible = 0
	CorrectionMode_PreInsertion = 1
	CorrectionMode_PostInsertionCollapsed = 2
	CorrectionMode_PostInsertionExpanded = 3
)

type CorrectionPosition int32

const (
	CorrectionPosition_Auto = 0
	CorrectionPosition_Bottom = 1
	CorrectionPosition_Top = 2
)

type InPlaceDirection int32

const (
	InPlaceDirection_Auto = 0
	InPlaceDirection_Bottom = 1
	InPlaceDirection_Top = 2
)

type EventMask int32

const (
	EventMask_InPlaceStateChanging = 1
	EventMask_InPlaceStateChanged = 2
	EventMask_InPlaceSizeChanging = 4
	EventMask_InPlaceSizeChanged = 8
	EventMask_InputAreaChanging = 16
	EventMask_InputAreaChanged = 32
	EventMask_CorrectionModeChanging = 64
	EventMask_CorrectionModeChanged = 128
	EventMask_InPlaceVisibilityChanging = 256
	EventMask_InPlaceVisibilityChanged = 512
	EventMask_TextInserting = 1024
	EventMask_TextInserted = 2048
	EventMask_All = 4095
)

type PanelType int32

const (
	PT_Default = 0
	PT_Inactive = 1
	PT_Handwriting = 2
	PT_Keyboard = 3
)

type FLICKDIRECTION int32

const (
	FLICKDIRECTION_MIN = 0
	FLICKDIRECTION_RIGHT = 0
	FLICKDIRECTION_UPRIGHT = 1
	FLICKDIRECTION_UP = 2
	FLICKDIRECTION_UPLEFT = 3
	FLICKDIRECTION_LEFT = 4
	FLICKDIRECTION_DOWNLEFT = 5
	FLICKDIRECTION_DOWN = 6
	FLICKDIRECTION_DOWNRIGHT = 7
	FLICKDIRECTION_INVALID = 8
)

type FLICKMODE int32

const (
	FLICKMODE_MIN = 0
	FLICKMODE_OFF = 0
	FLICKMODE_ON = 1
	FLICKMODE_LEARNING = 2
	FLICKMODE_MAX = 2
	FLICKMODE_DEFAULT = 1
)

type FLICKACTION_COMMANDCODE int32

const (
	FLICKACTION_COMMANDCODE_NULL = 0
	FLICKACTION_COMMANDCODE_SCROLL = 1
	FLICKACTION_COMMANDCODE_APPCOMMAND = 2
	FLICKACTION_COMMANDCODE_CUSTOMKEY = 3
	FLICKACTION_COMMANDCODE_KEYMODIFIER = 4
)

type SCROLLDIRECTION int32

const (
	SCROLLDIRECTION_UP = 0
	SCROLLDIRECTION_DOWN = 1
)

type KEYMODIFIER int32

const (
	KEYMODIFIER_CONTROL = 1
	KEYMODIFIER_MENU = 2
	KEYMODIFIER_SHIFT = 4
	KEYMODIFIER_WIN = 8
	KEYMODIFIER_ALTGR = 16
	KEYMODIFIER_EXT = 32
)

type MouseButton int32

const (
	NO_BUTTON = 0
	LEFT_BUTTON = 1
	RIGHT_BUTTON = 2
	MIDDLE_BUTTON = 4
)

type SelAlignmentConstants int32

const (
	rtfLeft = 0
	rtfRight = 1
	rtfCenter = 2
)

type DISPID_InkEdit int32

const (
	DISPID_Text = 0
	DISPID_TextRTF = 1
	DISPID_Hwnd = 2
	DISPID_DisableNoScroll = 3
	DISPID_Locked = 4
	DISPID_Enabled = 5
	DISPID_MaxLength = 6
	DISPID_MultiLine = 7
	DISPID_ScrollBars = 8
	DISPID_RTSelStart = 9
	DISPID_RTSelLength = 10
	DISPID_RTSelText = 11
	DISPID_SelAlignment = 12
	DISPID_SelBold = 13
	DISPID_SelCharOffset = 14
	DISPID_SelColor = 15
	DISPID_SelFontName = 16
	DISPID_SelFontSize = 17
	DISPID_SelItalic = 18
	DISPID_SelRTF = 19
	DISPID_SelUnderline = 20
	DISPID_DragIcon = 21
	DISPID_Status = 22
	DISPID_UseMouseForInput = 23
	DISPID_InkMode = 24
	DISPID_InkInsertMode = 25
	DISPID_RecoTimeout = 26
	DISPID_DrawAttr = 27
	DISPID_Recognizer = 28
	DISPID_Factoid = 29
	DISPID_SelInk = 30
	DISPID_SelInksDisplayMode = 31
	DISPID_Recognize = 32
	DISPID_GetGestStatus = 33
	DISPID_SetGestStatus = 34
	DISPID_Refresh = 35
)

type DISPID_InkEditEvents int32

const (
	DISPID_IeeChange = 1
	DISPID_IeeSelChange = 2
	DISPID_IeeKeyDown = 3
	DISPID_IeeKeyUp = 4
	DISPID_IeeMouseUp = 5
	DISPID_IeeMouseDown = 6
	DISPID_IeeKeyPress = 7
	DISPID_IeeDblClick = 8
	DISPID_IeeClick = 9
	DISPID_IeeMouseMove = 10
	DISPID_IeeCursorDown = 21
	DISPID_IeeStroke = 22
	DISPID_IeeGesture = 23
	DISPID_IeeRecognitionResult = 24
)

type InkMode int32

const (
	IEM_Disabled = 0
	IEM_Ink = 1
	IEM_InkAndGesture = 2
)

type InkInsertMode int32

const (
	IEM_InsertText = 0
	IEM_InsertInk = 1
)

type InkEditStatus int32

const (
	IES_Idle = 0
	IES_Collecting = 1
	IES_Recognizing = 2
)

type InkDisplayMode int32

const (
	IDM_Ink = 0
	IDM_Text = 1
)

type AppearanceConstants int32

const (
	rtfFlat = 0
	rtfThreeD = 1
)

type BorderStyleConstants int32

const (
	rtfNoBorder = 0
	rtfFixedSingle = 1
)

type ScrollBarsConstants int32

const (
	rtfNone = 0
	rtfHorizontal = 1
	rtfVertical = 2
	rtfBoth = 3
)

type MICUIELEMENT int32

const (
	MICUIELEMENT_BUTTON_WRITE = 1
	MICUIELEMENT_BUTTON_ERASE = 2
	MICUIELEMENT_BUTTON_CORRECT = 4
	MICUIELEMENT_BUTTON_CLEAR = 8
	MICUIELEMENT_BUTTON_UNDO = 16
	MICUIELEMENT_BUTTON_REDO = 32
	MICUIELEMENT_BUTTON_INSERT = 64
	MICUIELEMENT_BUTTON_CANCEL = 128
	MICUIELEMENT_INKPANEL_BACKGROUND = 256
	MICUIELEMENT_RESULTPANEL_BACKGROUND = 512
)

type MICUIELEMENTSTATE int32

const (
	MICUIELEMENTSTATE_NORMAL = 1
	MICUIELEMENTSTATE_HOT = 2
	MICUIELEMENTSTATE_PRESSED = 3
	MICUIELEMENTSTATE_DISABLED = 4
)

type DISPID_MathInputControlEvents int32

const (
	DISPID_MICInsert = 0
	DISPID_MICClose = 1
	DISPID_MICPaint = 2
	DISPID_MICClear = 3
)

type RealTimeStylusDataInterest int32

const (
	RTSDI_AllData = -1
	RTSDI_None = 0
	RTSDI_Error = 1
	RTSDI_RealTimeStylusEnabled = 2
	RTSDI_RealTimeStylusDisabled = 4
	RTSDI_StylusNew = 8
	RTSDI_StylusInRange = 16
	RTSDI_InAirPackets = 32
	RTSDI_StylusOutOfRange = 64
	RTSDI_StylusDown = 128
	RTSDI_Packets = 256
	RTSDI_StylusUp = 512
	RTSDI_StylusButtonUp = 1024
	RTSDI_StylusButtonDown = 2048
	RTSDI_SystemEvents = 4096
	RTSDI_TabletAdded = 8192
	RTSDI_TabletRemoved = 16384
	RTSDI_CustomStylusDataAdded = 32768
	RTSDI_UpdateMapping = 65536
	RTSDI_DefaultEvents = 37766
)

type StylusQueue int32

const (
	SyncStylusQueue = 1
	AsyncStylusQueueImmediate = 2
	AsyncStylusQueue = 3
)

type RealTimeStylusLockType int32

const (
	RTSLT_ObjLock = 1
	RTSLT_SyncEventLock = 2
	RTSLT_AsyncEventLock = 4
	RTSLT_ExcludeCallback = 8
	RTSLT_SyncObjLock = 11
	RTSLT_AsyncObjLock = 13
)

type LINE_METRICS int32

const (
	LM_BASELINE = 0
	LM_MIDLINE = 1
	LM_ASCENDER = 2
	LM_DESCENDER = 3
)

type CONFIDENCE_LEVEL int32

const (
	CFL_STRONG = 0
	CFL_INTERMEDIATE = 1
	CFL_POOR = 2
)

type ALT_BREAKS int32

const (
	ALT_BREAKS_SAME = 0
	ALT_BREAKS_UNIQUE = 1
	ALT_BREAKS_FULL = 2
)

type enumRECO_TYPE int32

const (
	RECO_TYPE_WSTRING = 0
	RECO_TYPE_WCHAR = 1
)

