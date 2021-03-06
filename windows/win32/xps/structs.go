// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package xps implements the Windows.Win32.Xps namespace.
package xps

type DRAWPATRECT struct {
	PTPOSITION int
	PTSIZE     int
	WSTYLE     int
	WPATTERN   int
}

type PSINJECTDATA struct {
	DataBytes      int
	InjectionPoint int
	PageNumber     int
}

type PSFEATURE_OUTPUT struct {
	BPAGEINDEPENDENT int
	BSETPAGEDEVICE   int
}

type PSFEATURE_CUSTPAPER struct {
	LORIENTATION  int
	LWIDTH        int
	LHEIGHT       int
	LWIDTHOFFSET  int
	LHEIGHTOFFSET int
}

type DEVMODEA struct {
	DMDEVICENAME       int
	DMSPECVERSION      int
	DMDRIVERVERSION    int
	DMSIZE             int
	DMDRIVEREXTRA      int
	DMFIELDS           int
	Anonymous1         int
	DMCOLOR            int
	DMDUPLEX           int
	DMYRESOLUTION      int
	DMTTOPTION         int
	DMCOLLATE          int
	DMFORMNAME         int
	DMLOGPIXELS        int
	DMBITSPERPEL       int
	DMPELSWIDTH        int
	DMPELSHEIGHT       int
	Anonymous2         int
	DMDISPLAYFREQUENCY int
	DMICMMETHOD        int
	DMICMINTENT        int
	DMMEDIATYPE        int
	DMDITHERTYPE       int
	DMRESERVED1        int
	DMRESERVED2        int
	DMPANNINGWIDTH     int
	DMPANNINGHEIGHT    int
}

type DOCINFOA struct {
	CBSIZE       int
	LPSZDOCNAME  int
	LPSZOUTPUT   int
	LPSZDATATYPE int
	FWTYPE       int
}

type DOCINFOW struct {
	CBSIZE       int
	LPSZDOCNAME  int
	LPSZOUTPUT   int
	LPSZDATATYPE int
	FWTYPE       int
}

type XpsOMObjectFactory struct {
}

type XpsOMThumbnailGenerator struct {
}

type XPS_POINT struct {
	X int
	Y int
}

type XPS_SIZE struct {
	WIDTH  int
	HEIGHT int
}

type XPS_RECT struct {
	X      int
	Y      int
	WIDTH  int
	HEIGHT int
}

type XPS_DASH struct {
	LENGTH int
	GAP    int
}

type XPS_GLYPH_MAPPING struct {
	UNICODESTRINGSTART  int
	UNICODESTRINGLENGTH int
	GLYPHINDICESSTART   int
	GLYPHINDICESLENGTH  int
}

type XPS_MATRIX struct {
	M11 int
	M12 int
	M21 int
	M22 int
	M31 int
	M32 int
}

type XPS_COLOR struct {
	COLORTYPE int
	VALUE     int
}

type XpsSignatureManager struct {
}

type PrintDocumentPackageTarget struct {
}

type PrintDocumentPackageTargetFactory struct {
}

type PrintDocumentPackageStatus struct {
	JobId            int
	CurrentDocument  int
	CurrentPage      int
	CurrentPageTotal int
	Completion       int
	PackageStatus    int
}

type HPTPROVIDER__ struct {
	UNUSED int
}
