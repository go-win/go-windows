// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package backgroundintelligenttransferservice implements the Windows.Win32.BackgroundIntelligentTransferService namespace.
package backgroundintelligenttransferservice

type BackgroundCopyManager struct {
}

type BG_FILE_PROGRESS struct {
	BytesTotal       int
	BytesTransferred int
	Completed        int
}

type BG_FILE_INFO struct {
	RemoteName int
	LocalName  int
}

type BG_JOB_PROGRESS struct {
	BytesTotal       int
	BytesTransferred int
	FilesTotal       int
	FilesTransferred int
}

type BG_JOB_TIMES struct {
	CreationTime           int
	ModificationTime       int
	TransferCompletionTime int
}

type BackgroundCopyManager1_5 struct {
}

type BG_JOB_REPLY_PROGRESS struct {
	BytesTotal       int
	BytesTransferred int
}

type BG_BASIC_CREDENTIALS struct {
	UserName int
	Password int
}

type BG_AUTH_CREDENTIALS_UNION struct {
	Basic int
}

type BG_AUTH_CREDENTIALS struct {
	Target      int
	Scheme      int
	Credentials int
}

type BackgroundCopyManager2_0 struct {
}

type BG_FILE_RANGE struct {
	InitialOffset int
	Length        int
}

type BackgroundCopyManager2_5 struct {
}

type BackgroundCopyManager3_0 struct {
}

type BackgroundCopyManager4_0 struct {
}

type BackgroundCopyManager5_0 struct {
}

type BITS_JOB_PROPERTY_VALUE struct {
	Dword  int
	ClsID  int
	Enable int
	Uint64 int
	Target int
}

type BITS_FILE_PROPERTY_VALUE struct {
	String int
}

type BackgroundCopyManager10_1 struct {
}

type BackgroundCopyManager10_2 struct {
}

type BackgroundCopyManager10_3 struct {
}

type BITSExtensionSetupFactory struct {
}

type BackgroundCopyQMgr struct {
}

type FILESETINFO struct {
	BSTRREMOTEFILE int
	BSTRLOCALFILE  int
	DWSIZEHINT     int
}