// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package softwaredevice implements the Windows.Win32.SoftwareDevice namespace.
package softwaredevice

type SW_DEVICE_CREATE_INFO struct {
	CBSIZE               int
	PSZINSTANCEID        int
	PSZZHARDWAREIDS      int
	PSZZCOMPATIBLEIDS    int
	PCONTAINERID         int
	CapabilityFlags      int
	PSZDEVICEDESCRIPTION int
	PSZDEVICELOCATION    int
	PSECURITYDESCRIPTOR  int
}

type HSWDEVICE__ struct {
	UNUSED int
}
