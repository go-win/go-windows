// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package qualityofservice implements the Windows.Win32.QualityOfService namespace.
package qualityofservice

type int_serv_wkp int32

const (
	IS_WKP_HOP_CNT      int_serv_wkp = 4
	IS_WKP_PATH_BW      int_serv_wkp = 6
	IS_WKP_MIN_LATENCY  int_serv_wkp = 8
	IS_WKP_COMPOSED_MTU int_serv_wkp = 10
	IS_WKP_TB_TSPEC     int_serv_wkp = 127
	IS_WKP_Q_TSPEC      int_serv_wkp = 128
)

type QOS_TRAFFIC_TYPE int32

const (
	QOSTrafficTypeBestEffort      QOS_TRAFFIC_TYPE = 0
	QOSTrafficTypeBackground      QOS_TRAFFIC_TYPE = 1
	QOSTrafficTypeExcellentEffort QOS_TRAFFIC_TYPE = 2
	QOSTrafficTypeAudioVideo      QOS_TRAFFIC_TYPE = 3
	QOSTrafficTypeVoice           QOS_TRAFFIC_TYPE = 4
	QOSTrafficTypeControl         QOS_TRAFFIC_TYPE = 5
)

type QOS_SET_FLOW int32

const (
	QOSSetTrafficType       QOS_SET_FLOW = 0
	QOSSetOutgoingRate      QOS_SET_FLOW = 1
	QOSSetOutgoingDSCPValue QOS_SET_FLOW = 2
)

type QOS_FLOWRATE_REASON int32

const (
	QOSFlowRateNotApplicable         QOS_FLOWRATE_REASON = 0
	QOSFlowRateContentChange         QOS_FLOWRATE_REASON = 1
	QOSFlowRateCongestion            QOS_FLOWRATE_REASON = 2
	QOSFlowRateHigherContentEncoding QOS_FLOWRATE_REASON = 3
	QOSFlowRateUserCaused            QOS_FLOWRATE_REASON = 4
)

type QOS_SHAPING int32

const (
	QOSShapeOnly                QOS_SHAPING = 0
	QOSShapeAndMark             QOS_SHAPING = 1
	QOSUseNonConformantMarkings QOS_SHAPING = 2
)

type QOS_QUERY_FLOW int32

const (
	QOSQueryFlowFundamentals QOS_QUERY_FLOW = 0
	QOSQueryPacketPriority   QOS_QUERY_FLOW = 1
	QOSQueryOutgoingRate     QOS_QUERY_FLOW = 2
)

type QOS_NOTIFY_FLOW int32

const (
	QOSNotifyCongested   QOS_NOTIFY_FLOW = 0
	QOSNotifyUncongested QOS_NOTIFY_FLOW = 1
	QOSNotifyAvailable   QOS_NOTIFY_FLOW = 2
)
