// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package ldap implements the Windows.Win32.Ldap namespace.
package ldap

type LDAP struct {
	LD_SB           int
	LD_HOST         int
	LD_VERSION      int
	LD_LBEROPTIONS  int
	LD_DEREF        int
	LD_TIMELIMIT    int
	LD_SIZELIMIT    int
	LD_ERRNO        int
	LD_MATCHED      int
	LD_ERROR        int
	LD_MSGID        int
	Reserved3       int
	LD_CLDAPTRIES   int
	LD_CLDAPTIMEOUT int
	LD_REFHOPLIMIT  int
	LD_OPTIONS      int
}

type LDAP_TIMEVAL struct {
	TV_SEC  int
	TV_USEC int
}

type LDAP_BERVAL struct {
	BV_LEN int
	BV_VAL int
}

type LDAPMessage struct {
	LM_MSGID             int
	LM_MSGTYPE           int
	LM_BER               int
	LM_CHAIN             int
	LM_NEXT              int
	LM_TIME              int
	Connection           int
	Request              int
	LM_RETURNCODE        int
	LM_REFERRAL          int
	LM_CHASED            int
	LM_EOM               int
	ConnectionReferenced int
}

type LDAPCONTROLA struct {
	LDCTL_OID        int
	LDCTL_VALUE      int
	LDCTL_ISCRITICAL int
}

type LDAPCONTROLW struct {
	LDCTL_OID        int
	LDCTL_VALUE      int
	LDCTL_ISCRITICAL int
}

type LDAPMODW struct {
	MOD_OP   int
	MOD_TYPE int
	MOD_VALS int
}

type LDAPMODA struct {
	MOD_OP   int
	MOD_TYPE int
	MOD_VALS int
}

type BERELEMENT struct {
	OPAQUE int
}

type LDAP_VERSION_INFO struct {
	LV_SIZE  int
	LV_MAJOR int
	LV_MINOR int
}

type LDAPAPIINFOA struct {
	LDAPAI_INFO_VERSION     int
	LDAPAI_API_VERSION      int
	LDAPAI_PROTOCOL_VERSION int
	LDAPAI_EXTENSIONS       int
	LDAPAI_VENDOR_NAME      int
	LDAPAI_VENDOR_VERSION   int
}

type LDAPAPIINFOW struct {
	LDAPAI_INFO_VERSION     int
	LDAPAI_API_VERSION      int
	LDAPAI_PROTOCOL_VERSION int
	LDAPAI_EXTENSIONS       int
	LDAPAI_VENDOR_NAME      int
	LDAPAI_VENDOR_VERSION   int
}

type LDAPAPIFeatureInfoA struct {
	LDAPAIF_INFO_VERSION int
	LDAPAIF_NAME         int
	LDAPAIF_VERSION      int
}

type LDAPAPIFeatureInfoW struct {
	LDAPAIF_INFO_VERSION int
	LDAPAIF_NAME         int
	LDAPAIF_VERSION      int
}

type LDAPSEARCH struct {
}

type LDAPSORTKEYW struct {
	SK_ATTRTYPE     int
	SK_MATCHRULEOID int
	SK_REVERSEORDER int
}

type LDAPSORTKEYA struct {
	SK_ATTRTYPE     int
	SK_MATCHRULEOID int
	SK_REVERSEORDER int
}

type LDAPVLVINFO struct {
	LDVLV_VERSION      int
	LDVLV_BEFORE_COUNT int
	LDVLV_AFTER_COUNT  int
	LDVLV_OFFSET       int
	LDVLV_COUNT        int
	LDVLV_ATTRVALUE    int
	LDVLV_CONTEXT      int
	LDVLV_EXTRADATA    int
}

type LDAP_REFERRAL_CALLBACK struct {
	SizeOfCallbacks    int
	QueryForConnection int
	NotifyRoutine      int
	DereferenceRoutine int
}
