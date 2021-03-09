// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package tapi implements the Windows.Win32.Tapi namespace.
package tapi

type LINEADDRESSCAPS struct {
	DWTOTALSIZE                    int
	DWNEEDEDSIZE                   int
	DWUSEDSIZE                     int
	DWLINEDEVICEID                 int
	DWADDRESSSIZE                  int
	DWADDRESSOFFSET                int
	DWDEVSPECIFICSIZE              int
	DWDEVSPECIFICOFFSET            int
	DWADDRESSSHARING               int
	DWADDRESSSTATES                int
	DWCALLINFOSTATES               int
	DWCALLERIDFLAGS                int
	DWCALLEDIDFLAGS                int
	DWCONNECTEDIDFLAGS             int
	DWREDIRECTIONIDFLAGS           int
	DWREDIRECTINGIDFLAGS           int
	DWCALLSTATES                   int
	DWDIALTONEMODES                int
	DWBUSYMODES                    int
	DWSPECIALINFO                  int
	DWDISCONNECTMODES              int
	DWMAXNUMACTIVECALLS            int
	DWMAXNUMONHOLDCALLS            int
	DWMAXNUMONHOLDPENDINGCALLS     int
	DWMAXNUMCONFERENCE             int
	DWMAXNUMTRANSCONF              int
	DWADDRCAPFLAGS                 int
	DWCALLFEATURES                 int
	DWREMOVEFROMCONFCAPS           int
	DWREMOVEFROMCONFSTATE          int
	DWTRANSFERMODES                int
	DWPARKMODES                    int
	DWFORWARDMODES                 int
	DWMAXFORWARDENTRIES            int
	DWMAXSPECIFICENTRIES           int
	DWMINFWDNUMRINGS               int
	DWMAXFWDNUMRINGS               int
	DWMAXCALLCOMPLETIONS           int
	DWCALLCOMPLETIONCONDS          int
	DWCALLCOMPLETIONMODES          int
	DWNUMCOMPLETIONMESSAGES        int
	DWCOMPLETIONMSGTEXTENTRYSIZE   int
	DWCOMPLETIONMSGTEXTSIZE        int
	DWCOMPLETIONMSGTEXTOFFSET      int
	DWADDRESSFEATURES              int
	DWPREDICTIVEAUTOTRANSFERSTATES int
	DWNUMCALLTREATMENTS            int
	DWCALLTREATMENTLISTSIZE        int
	DWCALLTREATMENTLISTOFFSET      int
	DWDEVICECLASSESSIZE            int
	DWDEVICECLASSESOFFSET          int
	DWMAXCALLDATASIZE              int
	DWCALLFEATURES2                int
	DWMAXNOANSWERTIMEOUT           int
	DWCONNECTEDMODES               int
	DWOFFERINGMODES                int
	DWAVAILABLEMEDIAMODES          int
}

type LINEADDRESSSTATUS struct {
	DWTOTALSIZE           int
	DWNEEDEDSIZE          int
	DWUSEDSIZE            int
	DWNUMINUSE            int
	DWNUMACTIVECALLS      int
	DWNUMONHOLDCALLS      int
	DWNUMONHOLDPENDCALLS  int
	DWADDRESSFEATURES     int
	DWNUMRINGSNOANSWER    int
	DWFORWARDNUMENTRIES   int
	DWFORWARDSIZE         int
	DWFORWARDOFFSET       int
	DWTERMINALMODESSIZE   int
	DWTERMINALMODESOFFSET int
	DWDEVSPECIFICSIZE     int
	DWDEVSPECIFICOFFSET   int
}

type LINEAGENTACTIVITYENTRY struct {
	DWID         int
	DWNAMESIZE   int
	DWNAMEOFFSET int
}

type LINEAGENTACTIVITYLIST struct {
	DWTOTALSIZE  int
	DWNEEDEDSIZE int
	DWUSEDSIZE   int
	DWNUMENTRIES int
	DWLISTSIZE   int
	DWLISTOFFSET int
}

type LINEAGENTCAPS struct {
	DWTOTALSIZE                  int
	DWNEEDEDSIZE                 int
	DWUSEDSIZE                   int
	DWAGENTHANDLERINFOSIZE       int
	DWAGENTHANDLERINFOOFFSET     int
	DWCAPSVERSION                int
	DWFEATURES                   int
	DWSTATES                     int
	DWNEXTSTATES                 int
	DWMAXNUMGROUPENTRIES         int
	DWAGENTSTATUSMESSAGES        int
	DWNUMAGENTEXTENSIONIDS       int
	DWAGENTEXTENSIONIDLISTSIZE   int
	DWAGENTEXTENSIONIDLISTOFFSET int
	ProxyGUID                    int
}

type LINEAGENTGROUPENTRY struct {
	GroupID      int
	DWNAMESIZE   int
	DWNAMEOFFSET int
}

type LINEAGENTGROUPLIST struct {
	DWTOTALSIZE  int
	DWNEEDEDSIZE int
	DWUSEDSIZE   int
	DWNUMENTRIES int
	DWLISTSIZE   int
	DWLISTOFFSET int
}

type LINEAGENTSTATUS struct {
	DWTOTALSIZE       int
	DWNEEDEDSIZE      int
	DWUSEDSIZE        int
	DWNUMENTRIES      int
	DWGROUPLISTSIZE   int
	DWGROUPLISTOFFSET int
	DWSTATE           int
	DWNEXTSTATE       int
	DWACTIVITYID      int
	DWACTIVITYSIZE    int
	DWACTIVITYOFFSET  int
	DWAGENTFEATURES   int
	DWVALIDSTATES     int
	DWVALIDNEXTSTATES int
}

type LINEAPPINFO struct {
	DWMACHINENAMESIZE      int
	DWMACHINENAMEOFFSET    int
	DWUSERNAMESIZE         int
	DWUSERNAMEOFFSET       int
	DWMODULEFILENAMESIZE   int
	DWMODULEFILENAMEOFFSET int
	DWFRIENDLYNAMESIZE     int
	DWFRIENDLYNAMEOFFSET   int
	DWMEDIAMODES           int
	DWADDRESSID            int
}

type LINEAGENTENTRY struct {
	HAGENT       int
	DWNAMESIZE   int
	DWNAMEOFFSET int
	DWIDSIZE     int
	DWIDOFFSET   int
	DWPINSIZE    int
	DWPINOFFSET  int
}

type LINEAGENTLIST struct {
	DWTOTALSIZE  int
	DWNEEDEDSIZE int
	DWUSEDSIZE   int
	DWNUMENTRIES int
	DWLISTSIZE   int
	DWLISTOFFSET int
}

type LINEAGENTINFO struct {
	DWTOTALSIZE             int
	DWNEEDEDSIZE            int
	DWUSEDSIZE              int
	DWAGENTSTATE            int
	DWNEXTAGENTSTATE        int
	DWMEASUREMENTPERIOD     int
	CYOVERALLCALLRATE       int
	DWNUMBEROFACDCALLS      int
	DWNUMBEROFINCOMINGCALLS int
	DWNUMBEROFOUTGOINGCALLS int
	DWTOTALACDTALKTIME      int
	DWTOTALACDCALLTIME      int
	DWTOTALACDWRAPUPTIME    int
}

type LINEAGENTSESSIONENTRY struct {
	HAGENTSESSION      int
	HAGENT             int
	GroupID            int
	DWWORKINGADDRESSID int
}

type LINEAGENTSESSIONLIST struct {
	DWTOTALSIZE  int
	DWNEEDEDSIZE int
	DWUSEDSIZE   int
	DWNUMENTRIES int
	DWLISTSIZE   int
	DWLISTOFFSET int
}

type LINEAGENTSESSIONINFO struct {
	DWTOTALSIZE             int
	DWNEEDEDSIZE            int
	DWUSEDSIZE              int
	DWAGENTSESSIONSTATE     int
	DWNEXTAGENTSESSIONSTATE int
	DATESESSIONSTARTTIME    int
	DWSESSIONDURATION       int
	DWNUMBEROFCALLS         int
	DWTOTALTALKTIME         int
	DWAVERAGETALKTIME       int
	DWTOTALCALLTIME         int
	DWAVERAGECALLTIME       int
	DWTOTALWRAPUPTIME       int
	DWAVERAGEWRAPUPTIME     int
	CYACDCALLRATE           int
	DWLONGESTTIMETOANSWER   int
	DWAVERAGETIMETOANSWER   int
}

type LINEQUEUEENTRY struct {
	DWQUEUEID    int
	DWNAMESIZE   int
	DWNAMEOFFSET int
}

type LINEQUEUELIST struct {
	DWTOTALSIZE  int
	DWNEEDEDSIZE int
	DWUSEDSIZE   int
	DWNUMENTRIES int
	DWLISTSIZE   int
	DWLISTOFFSET int
}

type LINEQUEUEINFO struct {
	DWTOTALSIZE              int
	DWNEEDEDSIZE             int
	DWUSEDSIZE               int
	DWMEASUREMENTPERIOD      int
	DWTOTALCALLSQUEUED       int
	DWCURRENTCALLSQUEUED     int
	DWTOTALCALLSABANDONED    int
	DWTOTALCALLSFLOWEDIN     int
	DWTOTALCALLSFLOWEDOUT    int
	DWLONGESTEVERWAITTIME    int
	DWCURRENTLONGESTWAITTIME int
	DWAVERAGEWAITTIME        int
	DWFINALDISPOSITION       int
}

type LINEPROXYREQUESTLIST struct {
	DWTOTALSIZE  int
	DWNEEDEDSIZE int
	DWUSEDSIZE   int
	DWNUMENTRIES int
	DWLISTSIZE   int
	DWLISTOFFSET int
}

type LINEDIALPARAMS struct {
	DWDIALPAUSE       int
	DWDIALSPEED       int
	DWDIGITDURATION   int
	DWWAITFORDIALTONE int
}

type LINECALLINFO struct {
	DWTOTALSIZE                int
	DWNEEDEDSIZE               int
	DWUSEDSIZE                 int
	HLINE                      int
	DWLINEDEVICEID             int
	DWADDRESSID                int
	DWBEARERMODE               int
	DWRATE                     int
	DWMEDIAMODE                int
	DWAPPSPECIFIC              int
	DWCALLID                   int
	DWRELATEDCALLID            int
	DWCALLPARAMFLAGS           int
	DWCALLSTATES               int
	DWMONITORDIGITMODES        int
	DWMONITORMEDIAMODES        int
	DialParams                 int
	DWORIGIN                   int
	DWREASON                   int
	DWCOMPLETIONID             int
	DWNUMOWNERS                int
	DWNUMMONITORS              int
	DWCOUNTRYCODE              int
	DWTRUNK                    int
	DWCALLERIDFLAGS            int
	DWCALLERIDSIZE             int
	DWCALLERIDOFFSET           int
	DWCALLERIDNAMESIZE         int
	DWCALLERIDNAMEOFFSET       int
	DWCALLEDIDFLAGS            int
	DWCALLEDIDSIZE             int
	DWCALLEDIDOFFSET           int
	DWCALLEDIDNAMESIZE         int
	DWCALLEDIDNAMEOFFSET       int
	DWCONNECTEDIDFLAGS         int
	DWCONNECTEDIDSIZE          int
	DWCONNECTEDIDOFFSET        int
	DWCONNECTEDIDNAMESIZE      int
	DWCONNECTEDIDNAMEOFFSET    int
	DWREDIRECTIONIDFLAGS       int
	DWREDIRECTIONIDSIZE        int
	DWREDIRECTIONIDOFFSET      int
	DWREDIRECTIONIDNAMESIZE    int
	DWREDIRECTIONIDNAMEOFFSET  int
	DWREDIRECTINGIDFLAGS       int
	DWREDIRECTINGIDSIZE        int
	DWREDIRECTINGIDOFFSET      int
	DWREDIRECTINGIDNAMESIZE    int
	DWREDIRECTINGIDNAMEOFFSET  int
	DWAPPNAMESIZE              int
	DWAPPNAMEOFFSET            int
	DWDISPLAYABLEADDRESSSIZE   int
	DWDISPLAYABLEADDRESSOFFSET int
	DWCALLEDPARTYSIZE          int
	DWCALLEDPARTYOFFSET        int
	DWCOMMENTSIZE              int
	DWCOMMENTOFFSET            int
	DWDISPLAYSIZE              int
	DWDISPLAYOFFSET            int
	DWUSERUSERINFOSIZE         int
	DWUSERUSERINFOOFFSET       int
	DWHIGHLEVELCOMPSIZE        int
	DWHIGHLEVELCOMPOFFSET      int
	DWLOWLEVELCOMPSIZE         int
	DWLOWLEVELCOMPOFFSET       int
	DWCHARGINGINFOSIZE         int
	DWCHARGINGINFOOFFSET       int
	DWTERMINALMODESSIZE        int
	DWTERMINALMODESOFFSET      int
	DWDEVSPECIFICSIZE          int
	DWDEVSPECIFICOFFSET        int
	DWCALLTREATMENT            int
	DWCALLDATASIZE             int
	DWCALLDATAOFFSET           int
	DWSENDINGFLOWSPECSIZE      int
	DWSENDINGFLOWSPECOFFSET    int
	DWRECEIVINGFLOWSPECSIZE    int
	DWRECEIVINGFLOWSPECOFFSET  int
}

type LINECALLLIST struct {
	DWTOTALSIZE       int
	DWNEEDEDSIZE      int
	DWUSEDSIZE        int
	DWCALLSNUMENTRIES int
	DWCALLSSIZE       int
	DWCALLSOFFSET     int
}

type LINECALLPARAMS struct {
	DWTOTALSIZE                    int
	DWBEARERMODE                   int
	DWMINRATE                      int
	DWMAXRATE                      int
	DWMEDIAMODE                    int
	DWCALLPARAMFLAGS               int
	DWADDRESSMODE                  int
	DWADDRESSID                    int
	DialParams                     int
	DWORIGADDRESSSIZE              int
	DWORIGADDRESSOFFSET            int
	DWDISPLAYABLEADDRESSSIZE       int
	DWDISPLAYABLEADDRESSOFFSET     int
	DWCALLEDPARTYSIZE              int
	DWCALLEDPARTYOFFSET            int
	DWCOMMENTSIZE                  int
	DWCOMMENTOFFSET                int
	DWUSERUSERINFOSIZE             int
	DWUSERUSERINFOOFFSET           int
	DWHIGHLEVELCOMPSIZE            int
	DWHIGHLEVELCOMPOFFSET          int
	DWLOWLEVELCOMPSIZE             int
	DWLOWLEVELCOMPOFFSET           int
	DWDEVSPECIFICSIZE              int
	DWDEVSPECIFICOFFSET            int
	DWPREDICTIVEAUTOTRANSFERSTATES int
	DWTARGETADDRESSSIZE            int
	DWTARGETADDRESSOFFSET          int
	DWSENDINGFLOWSPECSIZE          int
	DWSENDINGFLOWSPECOFFSET        int
	DWRECEIVINGFLOWSPECSIZE        int
	DWRECEIVINGFLOWSPECOFFSET      int
	DWDEVICECLASSSIZE              int
	DWDEVICECLASSOFFSET            int
	DWDEVICECONFIGSIZE             int
	DWDEVICECONFIGOFFSET           int
	DWCALLDATASIZE                 int
	DWCALLDATAOFFSET               int
	DWNOANSWERTIMEOUT              int
	DWCALLINGPARTYIDSIZE           int
	DWCALLINGPARTYIDOFFSET         int
}

type LINECALLSTATUS struct {
	DWTOTALSIZE         int
	DWNEEDEDSIZE        int
	DWUSEDSIZE          int
	DWCALLSTATE         int
	DWCALLSTATEMODE     int
	DWCALLPRIVILEGE     int
	DWCALLFEATURES      int
	DWDEVSPECIFICSIZE   int
	DWDEVSPECIFICOFFSET int
	DWCALLFEATURES2     int
	TSTATEENTRYTIME     int
}

type LINECALLTREATMENTENTRY struct {
	DWCALLTREATMENTID         int
	DWCALLTREATMENTNAMESIZE   int
	DWCALLTREATMENTNAMEOFFSET int
}

type LINECARDENTRY struct {
	DWPERMANENTCARDID         int
	DWCARDNAMESIZE            int
	DWCARDNAMEOFFSET          int
	DWCARDNUMBERDIGITS        int
	DWSAMEAREARULESIZE        int
	DWSAMEAREARULEOFFSET      int
	DWLONGDISTANCERULESIZE    int
	DWLONGDISTANCERULEOFFSET  int
	DWINTERNATIONALRULESIZE   int
	DWINTERNATIONALRULEOFFSET int
	DWOPTIONS                 int
}

type LINECOUNTRYENTRY struct {
	DWCOUNTRYID               int
	DWCOUNTRYCODE             int
	DWNEXTCOUNTRYID           int
	DWCOUNTRYNAMESIZE         int
	DWCOUNTRYNAMEOFFSET       int
	DWSAMEAREARULESIZE        int
	DWSAMEAREARULEOFFSET      int
	DWLONGDISTANCERULESIZE    int
	DWLONGDISTANCERULEOFFSET  int
	DWINTERNATIONALRULESIZE   int
	DWINTERNATIONALRULEOFFSET int
}

type LINECOUNTRYLIST struct {
	DWTOTALSIZE         int
	DWNEEDEDSIZE        int
	DWUSEDSIZE          int
	DWNUMCOUNTRIES      int
	DWCOUNTRYLISTSIZE   int
	DWCOUNTRYLISTOFFSET int
}

type LINEDEVCAPS struct {
	DWTOTALSIZE                  int
	DWNEEDEDSIZE                 int
	DWUSEDSIZE                   int
	DWPROVIDERINFOSIZE           int
	DWPROVIDERINFOOFFSET         int
	DWSWITCHINFOSIZE             int
	DWSWITCHINFOOFFSET           int
	DWPERMANENTLINEID            int
	DWLINENAMESIZE               int
	DWLINENAMEOFFSET             int
	DWSTRINGFORMAT               int
	DWADDRESSMODES               int
	DWNUMADDRESSES               int
	DWBEARERMODES                int
	DWMAXRATE                    int
	DWMEDIAMODES                 int
	DWGENERATETONEMODES          int
	DWGENERATETONEMAXNUMFREQ     int
	DWGENERATEDIGITMODES         int
	DWMONITORTONEMAXNUMFREQ      int
	DWMONITORTONEMAXNUMENTRIES   int
	DWMONITORDIGITMODES          int
	DWGATHERDIGITSMINTIMEOUT     int
	DWGATHERDIGITSMAXTIMEOUT     int
	DWMEDCTLDIGITMAXLISTSIZE     int
	DWMEDCTLMEDIAMAXLISTSIZE     int
	DWMEDCTLTONEMAXLISTSIZE      int
	DWMEDCTLCALLSTATEMAXLISTSIZE int
	DWDEVCAPFLAGS                int
	DWMAXNUMACTIVECALLS          int
	DWANSWERMODE                 int
	DWRINGMODES                  int
	DWLINESTATES                 int
	DWUUIACCEPTSIZE              int
	DWUUIANSWERSIZE              int
	DWUUIMAKECALLSIZE            int
	DWUUIDROPSIZE                int
	DWUUISENDUSERUSERINFOSIZE    int
	DWUUICALLINFOSIZE            int
	MinDialParams                int
	MaxDialParams                int
	DefaultDialParams            int
	DWNUMTERMINALS               int
	DWTERMINALCAPSSIZE           int
	DWTERMINALCAPSOFFSET         int
	DWTERMINALTEXTENTRYSIZE      int
	DWTERMINALTEXTSIZE           int
	DWTERMINALTEXTOFFSET         int
	DWDEVSPECIFICSIZE            int
	DWDEVSPECIFICOFFSET          int
	DWLINEFEATURES               int
	DWSETTABLEDEVSTATUS          int
	DWDEVICECLASSESSIZE          int
	DWDEVICECLASSESOFFSET        int
	PermanentLineGuid            int
}

type LINEDEVSTATUS struct {
	DWTOTALSIZE           int
	DWNEEDEDSIZE          int
	DWUSEDSIZE            int
	DWNUMOPENS            int
	DWOPENMEDIAMODES      int
	DWNUMACTIVECALLS      int
	DWNUMONHOLDCALLS      int
	DWNUMONHOLDPENDCALLS  int
	DWLINEFEATURES        int
	DWNUMCALLCOMPLETIONS  int
	DWRINGMODE            int
	DWSIGNALLEVEL         int
	DWBATTERYLEVEL        int
	DWROAMMODE            int
	DWDEVSTATUSFLAGS      int
	DWTERMINALMODESSIZE   int
	DWTERMINALMODESOFFSET int
	DWDEVSPECIFICSIZE     int
	DWDEVSPECIFICOFFSET   int
	DWAVAILABLEMEDIAMODES int
	DWAPPINFOSIZE         int
	DWAPPINFOOFFSET       int
}

type LINEEXTENSIONID struct {
	DWEXTENSIONID0 int
	DWEXTENSIONID1 int
	DWEXTENSIONID2 int
	DWEXTENSIONID3 int
}

type LINEFORWARD struct {
	DWFORWARDMODE         int
	DWCALLERADDRESSSIZE   int
	DWCALLERADDRESSOFFSET int
	DWDESTCOUNTRYCODE     int
	DWDESTADDRESSSIZE     int
	DWDESTADDRESSOFFSET   int
}

type LINEFORWARDLIST struct {
	DWTOTALSIZE  int
	DWNUMENTRIES int
	ForwardList  int
}

type LINEGENERATETONE struct {
	DWFREQUENCY  int
	DWCADENCEON  int
	DWCADENCEOFF int
	DWVOLUME     int
}

type LINEINITIALIZEEXPARAMS struct {
	DWTOTALSIZE     int
	DWNEEDEDSIZE    int
	DWUSEDSIZE      int
	DWOPTIONS       int
	Handles         int
	DWCOMPLETIONKEY int
}

type LINELOCATIONENTRY struct {
	DWPERMANENTLOCATIONID          int
	DWLOCATIONNAMESIZE             int
	DWLOCATIONNAMEOFFSET           int
	DWCOUNTRYCODE                  int
	DWCITYCODESIZE                 int
	DWCITYCODEOFFSET               int
	DWPREFERREDCARDID              int
	DWLOCALACCESSCODESIZE          int
	DWLOCALACCESSCODEOFFSET        int
	DWLONGDISTANCEACCESSCODESIZE   int
	DWLONGDISTANCEACCESSCODEOFFSET int
	DWTOLLPREFIXLISTSIZE           int
	DWTOLLPREFIXLISTOFFSET         int
	DWCOUNTRYID                    int
	DWOPTIONS                      int
	DWCANCELCALLWAITINGSIZE        int
	DWCANCELCALLWAITINGOFFSET      int
}

type LINEMEDIACONTROLCALLSTATE struct {
	DWCALLSTATES   int
	DWMEDIACONTROL int
}

type LINEMEDIACONTROLDIGIT struct {
	DWDIGIT        int
	DWDIGITMODES   int
	DWMEDIACONTROL int
}

type LINEMEDIACONTROLMEDIA struct {
	DWMEDIAMODES   int
	DWDURATION     int
	DWMEDIACONTROL int
}

type LINEMEDIACONTROLTONE struct {
	DWAPPSPECIFIC  int
	DWDURATION     int
	DWFREQUENCY1   int
	DWFREQUENCY2   int
	DWFREQUENCY3   int
	DWMEDIACONTROL int
}

type LINEMESSAGE struct {
	HDEVICE            int
	DWMESSAGEID        int
	DWCALLBACKINSTANCE int
	DWPARAM1           int
	DWPARAM2           int
	DWPARAM3           int
}

type LINEMONITORTONE struct {
	DWAPPSPECIFIC int
	DWDURATION    int
	DWFREQUENCY1  int
	DWFREQUENCY2  int
	DWFREQUENCY3  int
}

type LINEPROVIDERENTRY struct {
	DWPERMANENTPROVIDERID    int
	DWPROVIDERFILENAMESIZE   int
	DWPROVIDERFILENAMEOFFSET int
}

type LINEPROVIDERLIST struct {
	DWTOTALSIZE          int
	DWNEEDEDSIZE         int
	DWUSEDSIZE           int
	DWNUMPROVIDERS       int
	DWPROVIDERLISTSIZE   int
	DWPROVIDERLISTOFFSET int
}

type LINEPROXYREQUEST struct {
	DWSIZE                    int
	DWCLIENTMACHINENAMESIZE   int
	DWCLIENTMACHINENAMEOFFSET int
	DWCLIENTUSERNAMESIZE      int
	DWCLIENTUSERNAMEOFFSET    int
	DWCLIENTAPPAPIVERSION     int
	DWREQUESTTYPE             int
	Anonymous                 int
}

type LINEREQMAKECALL struct {
	SZDESTADDRESS int
	SZAPPNAME     int
	SZCALLEDPARTY int
	SZCOMMENT     int
}

type LINEREQMAKECALLW_TAG struct {
	SZDESTADDRESS int
	SZAPPNAME     int
	SZCALLEDPARTY int
	SZCOMMENT     int
}

type LINEREQMEDIACALL struct {
	HWND          int
	WREQUESTID    int
	SZDEVICECLASS int
	UCDEVICEID    int
	DWSIZE        int
	DWSECURE      int
	SZDESTADDRESS int
	SZAPPNAME     int
	SZCALLEDPARTY int
	SZCOMMENT     int
}

type LINEREQMEDIACALLW_TAG struct {
	HWND          int
	WREQUESTID    int
	SZDEVICECLASS int
	UCDEVICEID    int
	DWSIZE        int
	DWSECURE      int
	SZDESTADDRESS int
	SZAPPNAME     int
	SZCALLEDPARTY int
	SZCOMMENT     int
}

type LINETERMCAPS struct {
	DWTERMDEV     int
	DWTERMMODES   int
	DWTERMSHARING int
}

type LINETRANSLATECAPS struct {
	DWTOTALSIZE              int
	DWNEEDEDSIZE             int
	DWUSEDSIZE               int
	DWNUMLOCATIONS           int
	DWLOCATIONLISTSIZE       int
	DWLOCATIONLISTOFFSET     int
	DWCURRENTLOCATIONID      int
	DWNUMCARDS               int
	DWCARDLISTSIZE           int
	DWCARDLISTOFFSET         int
	DWCURRENTPREFERREDCARDID int
}

type LINETRANSLATEOUTPUT struct {
	DWTOTALSIZE               int
	DWNEEDEDSIZE              int
	DWUSEDSIZE                int
	DWDIALABLESTRINGSIZE      int
	DWDIALABLESTRINGOFFSET    int
	DWDISPLAYABLESTRINGSIZE   int
	DWDISPLAYABLESTRINGOFFSET int
	DWCURRENTCOUNTRY          int
	DWDESTCOUNTRY             int
	DWTRANSLATERESULTS        int
}

type PHONEBUTTONINFO struct {
	DWTOTALSIZE         int
	DWNEEDEDSIZE        int
	DWUSEDSIZE          int
	DWBUTTONMODE        int
	DWBUTTONFUNCTION    int
	DWBUTTONTEXTSIZE    int
	DWBUTTONTEXTOFFSET  int
	DWDEVSPECIFICSIZE   int
	DWDEVSPECIFICOFFSET int
	DWBUTTONSTATE       int
}

type PHONECAPS struct {
	DWTOTALSIZE                       int
	DWNEEDEDSIZE                      int
	DWUSEDSIZE                        int
	DWPROVIDERINFOSIZE                int
	DWPROVIDERINFOOFFSET              int
	DWPHONEINFOSIZE                   int
	DWPHONEINFOOFFSET                 int
	DWPERMANENTPHONEID                int
	DWPHONENAMESIZE                   int
	DWPHONENAMEOFFSET                 int
	DWSTRINGFORMAT                    int
	DWPHONESTATES                     int
	DWHOOKSWITCHDEVS                  int
	DWHANDSETHOOKSWITCHMODES          int
	DWSPEAKERHOOKSWITCHMODES          int
	DWHEADSETHOOKSWITCHMODES          int
	DWVOLUMEFLAGS                     int
	DWGAINFLAGS                       int
	DWDISPLAYNUMROWS                  int
	DWDISPLAYNUMCOLUMNS               int
	DWNUMRINGMODES                    int
	DWNUMBUTTONLAMPS                  int
	DWBUTTONMODESSIZE                 int
	DWBUTTONMODESOFFSET               int
	DWBUTTONFUNCTIONSSIZE             int
	DWBUTTONFUNCTIONSOFFSET           int
	DWLAMPMODESSIZE                   int
	DWLAMPMODESOFFSET                 int
	DWNUMSETDATA                      int
	DWSETDATASIZE                     int
	DWSETDATAOFFSET                   int
	DWNUMGETDATA                      int
	DWGETDATASIZE                     int
	DWGETDATAOFFSET                   int
	DWDEVSPECIFICSIZE                 int
	DWDEVSPECIFICOFFSET               int
	DWDEVICECLASSESSIZE               int
	DWDEVICECLASSESOFFSET             int
	DWPHONEFEATURES                   int
	DWSETTABLEHANDSETHOOKSWITCHMODES  int
	DWSETTABLESPEAKERHOOKSWITCHMODES  int
	DWSETTABLEHEADSETHOOKSWITCHMODES  int
	DWMONITOREDHANDSETHOOKSWITCHMODES int
	DWMONITOREDSPEAKERHOOKSWITCHMODES int
	DWMONITOREDHEADSETHOOKSWITCHMODES int
	PermanentPhoneGuid                int
}

type PHONEEXTENSIONID struct {
	DWEXTENSIONID0 int
	DWEXTENSIONID1 int
	DWEXTENSIONID2 int
	DWEXTENSIONID3 int
}

type PHONEINITIALIZEEXPARAMS struct {
	DWTOTALSIZE     int
	DWNEEDEDSIZE    int
	DWUSEDSIZE      int
	DWOPTIONS       int
	Handles         int
	DWCOMPLETIONKEY int
}

type PHONEMESSAGE struct {
	HDEVICE            int
	DWMESSAGEID        int
	DWCALLBACKINSTANCE int
	DWPARAM1           int
	DWPARAM2           int
	DWPARAM3           int
}

type PHONESTATUS struct {
	DWTOTALSIZE             int
	DWNEEDEDSIZE            int
	DWUSEDSIZE              int
	DWSTATUSFLAGS           int
	DWNUMOWNERS             int
	DWNUMMONITORS           int
	DWRINGMODE              int
	DWRINGVOLUME            int
	DWHANDSETHOOKSWITCHMODE int
	DWHANDSETVOLUME         int
	DWHANDSETGAIN           int
	DWSPEAKERHOOKSWITCHMODE int
	DWSPEAKERVOLUME         int
	DWSPEAKERGAIN           int
	DWHEADSETHOOKSWITCHMODE int
	DWHEADSETVOLUME         int
	DWHEADSETGAIN           int
	DWDISPLAYSIZE           int
	DWDISPLAYOFFSET         int
	DWLAMPMODESSIZE         int
	DWLAMPMODESOFFSET       int
	DWOWNERNAMESIZE         int
	DWOWNERNAMEOFFSET       int
	DWDEVSPECIFICSIZE       int
	DWDEVSPECIFICOFFSET     int
	DWPHONEFEATURES         int
}

type VARSTRING struct {
	DWTOTALSIZE    int
	DWNEEDEDSIZE   int
	DWUSEDSIZE     int
	DWSTRINGFORMAT int
	DWSTRINGSIZE   int
	DWSTRINGOFFSET int
}

type TAPI struct {
}

type DispatchMapper struct {
}

type RequestMakeCall struct {
}

type TAPI_CUSTOMTONE struct {
	DWFREQUENCY  int
	DWCADENCEON  int
	DWCADENCEOFF int
	DWVOLUME     int
}

type TAPI_DETECTTONE struct {
	DWAPPSPECIFIC int
	DWDURATION    int
	DWFREQUENCY1  int
	DWFREQUENCY2  int
	DWFREQUENCY3  int
}

type MSP_EVENT_INFO struct {
	DWSIZE    int
	Event     int
	HCALL     int
	Anonymous int
}

type Rendezvous struct {
}

type McastAddressAllocation struct {
}