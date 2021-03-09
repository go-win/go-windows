// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package pointerinput implements the Windows.Win32.PointerInput namespace.
package pointerinput

type POINTER_INFO struct {
	POINTERTYPE           int
	POINTERID             int
	FRAMEID               int
	POINTERFLAGS          int
	SOURCEDEVICE          int
	HWNDTARGET            int
	PTPIXELLOCATION       int
	PTHIMETRICLOCATION    int
	PTPIXELLOCATIONRAW    int
	PTHIMETRICLOCATIONRAW int
	DWTIME                int
	HISTORYCOUNT          int
	InputData             int
	DWKEYSTATES           int
	PerformanceCount      int
	ButtonChangeType      int
}

type POINTER_TOUCH_INFO struct {
	POINTERINFO  int
	TOUCHFLAGS   int
	TOUCHMASK    int
	RCCONTACT    int
	RCCONTACTRAW int
	ORIENTATION  int
	PRESSURE     int
}

type POINTER_PEN_INFO struct {
	POINTERINFO int
	PENFLAGS    int
	PENMASK     int
	PRESSURE    int
	ROTATION    int
	TILTX       int
	TILTY       int
}

type INPUT_TRANSFORM struct {
	Anonymous int
}
