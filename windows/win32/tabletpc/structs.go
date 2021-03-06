// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package tabletpc implements the Windows.Win32.TabletPC namespace.
package tabletpc

type InkDisp struct {
}

type InkOverlay struct {
}

type InkPicture struct {
}

type InkCollector struct {
}

type InkDrawingAttributes struct {
}

type InkRectangle struct {
}

type InkRenderer struct {
}

type InkTransform struct {
}

type InkRecognizers struct {
}

type InkRecognizerContext struct {
}

type InkRecognizerGuide struct {
}

type InkTablets struct {
}

type InkWordList struct {
}

type InkStrokes struct {
}

type Ink struct {
}

type SketchInk struct {
}

type SYSTEM_EVENT_DATA struct {
	BMODIFIER     int
	WKEY          int
	XPOS          int
	YPOS          int
	BCURSORMODE   int
	DWBUTTONSTATE int
}

type STROKE_RANGE struct {
	ISTROKEBEGIN int
	ISTROKEEND   int
}

type PROPERTY_METRICS struct {
	NLOGICALMIN int
	NLOGICALMAX int
	Units       int
	FRESOLUTION int
}

type PACKET_PROPERTY struct {
	GUID            int
	PropertyMetrics int
}

type PACKET_DESCRIPTION struct {
	CBPACKETSIZE      int
	CPACKETPROPERTIES int
	PPACKETPROPERTIES int
	CBUTTONS          int
	PGUIDBUTTONS      int
}

type INKMETRIC struct {
	IHEIGHT      int
	IFONTASCENT  int
	IFONTDESCENT int
	DWFLAGS      int
	COLOR        int
}

type InkRecoGuide struct {
	RECTWRITINGBOX int
	RECTDRAWNBOX   int
	CROWS          int
	CCOLUMNS       int
	MIDLINE        int
}

type InkDivider struct {
}

type HandwrittenTextInsertion struct {
}

type PenInputPanel struct {
}

type TextInputPanel struct {
}

type PenInputPanel_Internal struct {
}

type FLICK_POINT struct {
	BITFIELD int
}

type FLICK_DATA struct {
	BITFIELD int
}

type InkEdit struct {
}

type IEC_STROKEINFO struct {
	NMHDR  int
	Cursor int
	Stroke int
}

type IEC_GESTUREINFO struct {
	NMHDR    int
	Cursor   int
	Strokes  int
	Gestures int
}

type IEC_RECOGNITIONRESULTINFO struct {
	NMHDR             int
	RecognitionResult int
}

type MathInputControl struct {
}

type RealTimeStylus struct {
}

type DynamicRenderer struct {
}

type GestureRecognizer struct {
}

type StrokeBuilder struct {
}

type StylusInfo struct {
	TCID              int
	CID               int
	BISINVERTEDCURSOR int
}

type GESTURE_DATA struct {
	GESTUREID      int
	RECOCONFIDENCE int
	STROKECOUNT    int
}

type DYNAMIC_RENDERER_CACHED_DATA struct {
	STROKEID        int
	DYNAMICRENDERER int
}

type RECO_GUIDE struct {
	XORIGIN  int
	YORIGIN  int
	CXBOX    int
	CYBOX    int
	CXBASE   int
	CYBASE   int
	CHORZBOX int
	CVERTBOX int
	CYMID    int
}

type RECO_ATTRS struct {
	DWRECOCAPABILITYFLAGS int
	AWCVENDORNAME         int
	AWCFRIENDLYNAME       int
	AWLANGUAGEID          int
}

type RECO_RANGE struct {
	IWCBEGIN int
	CCOUNT   int
}

type LINE_SEGMENT struct {
	PtA int
	PtB int
}

type LATTICE_METRICS struct {
	LSBASELINE     int
	IMIDLINEOFFSET int
}

type RECO_LATTICE_PROPERTY struct {
	GUIDPROPERTY    int
	CBPROPERTYVALUE int
	PPROPERTYVALUE  int
}

type RECO_LATTICE_PROPERTIES struct {
	CPROPERTIES int
	APPROPS     int
}

type RECO_LATTICE_ELEMENT struct {
	SCORE          int
	TYPE           int
	PDATA          int
	ULNEXTCOLUMN   int
	ULSTROKENUMBER int
	EPPROP         int
}

type RECO_LATTICE_COLUMN struct {
	KEY              int
	CPPROP           int
	CSTROKES         int
	PSTROKES         int
	CLATTICEELEMENTS int
	PLATTICEELEMENTS int
}

type RECO_LATTICE struct {
	ULCOLUMNCOUNT           int
	PLATTICECOLUMNS         int
	ULPROPERTYCOUNT         int
	PGUIDPROPERTIES         int
	ULBESTRESULTCOLUMNCOUNT int
	PULBESTRESULTCOLUMNS    int
	PULBESTRESULTINDEXES    int
}

type CHARACTER_RANGE struct {
	WCLOW  int
	CCHARS int
}

type HRECOALT__ struct {
	UNUSED int
}

type HRECOCONTEXT__ struct {
	UNUSED int
}

type HRECOGNIZER__ struct {
	UNUSED int
}

type HRECOLATTICE__ struct {
	UNUSED int
}

type HRECOWORDLIST__ struct {
	UNUSED int
}

type TipAutoCompleteClient struct {
}
