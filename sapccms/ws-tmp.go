
package sapccms



import "encoding/xml"
import "strings"
import "io/ioutil"
import "net/http"


type SAPCCMS struct {
	Url string
}

func NewSAPCCMS(url string)(svc SAPCCMS){
	svc.Url=url
	return
}


type InstanceType struct{
	InstanceType string `xml:"InstanceType"`
}

func (t InstanceType) GetInstanceType() string {
	return t.InstanceType
}

func (t *InstanceType) SetInstanceType (val string) {
	t.InstanceType = val
}



type  ALGLOBTID struct {
Mtsysid string `xml:"Mtsysid"`
Mtmcname string `xml:"Mtmcname"`
Mtnumrange string `xml:"Mtnumrange"`
Mtuid string `xml:"Mtuid"`
Mtclass string `xml:"Mtclass"`
Mtindex string `xml:"Mtindex"`
Extindex string `xml:"Extindex"`
}

func (t ALGLOBTID)GetMtsysid()string{
	return t.Mtsysid
}

func (t *ALGLOBTID)SetMtsysid(val string){
	t.Mtsysid = val
}

func (t ALGLOBTID)GetMtmcname()string{
	return t.Mtmcname
}

func (t *ALGLOBTID)SetMtmcname(val string){
	t.Mtmcname = val
}

func (t ALGLOBTID)GetMtnumrange()string{
	return t.Mtnumrange
}

func (t *ALGLOBTID)SetMtnumrange(val string){
	t.Mtnumrange = val
}

func (t ALGLOBTID)GetMtuid()string{
	return t.Mtuid
}

func (t *ALGLOBTID)SetMtuid(val string){
	t.Mtuid = val
}

func (t ALGLOBTID)GetMtclass()string{
	return t.Mtclass
}

func (t *ALGLOBTID)SetMtclass(val string){
	t.Mtclass = val
}

func (t ALGLOBTID)GetMtindex()string{
	return t.Mtindex
}

func (t *ALGLOBTID)SetMtindex(val string){
	t.Mtindex = val
}

func (t ALGLOBTID)GetExtindex()string{
	return t.Extindex
}

func (t *ALGLOBTID)SetExtindex(val string){
	t.Extindex = val
}


type  ALGLOBAID struct {
Alsysid string `xml:"Alsysid"`
Msegname string `xml:"Msegname"`
Aluniqnum string `xml:"Aluniqnum"`
Alindex string `xml:"Alindex"`
Alertdate string `xml:"Alertdate"`
Alerttime string `xml:"Alerttime"`
}

func (t ALGLOBAID)GetAlsysid()string{
	return t.Alsysid
}

func (t *ALGLOBAID)SetAlsysid(val string){
	t.Alsysid = val
}

func (t ALGLOBAID)GetMsegname()string{
	return t.Msegname
}

func (t *ALGLOBAID)SetMsegname(val string){
	t.Msegname = val
}

func (t ALGLOBAID)GetAluniqnum()string{
	return t.Aluniqnum
}

func (t *ALGLOBAID)SetAluniqnum(val string){
	t.Aluniqnum = val
}

func (t ALGLOBAID)GetAlindex()string{
	return t.Alindex
}

func (t *ALGLOBAID)SetAlindex(val string){
	t.Alindex = val
}

func (t ALGLOBAID)GetAlertdate()string{
	return t.Alertdate
}

func (t *ALGLOBAID)SetAlertdate(val string){
	t.Alertdate = val
}

func (t ALGLOBAID)GetAlerttime()string{
	return t.Alerttime
}

func (t *ALGLOBAID)SetAlerttime(val string){
	t.Alerttime = val
}


type  ALMTCUST struct {
CustStatus string `xml:"CustStatus"`
Useclass string `xml:"Useclass"`
Severity string `xml:"Severity"`
StatRecord string `xml:"StatRecord"`
VisibilityLevel string `xml:"VisibilityLevel"`
TypeOfAlertsToKeep string `xml:"TypeOfAlertsToKeep"`
NumOfAlertsToKeep string `xml:"NumOfAlertsToKeep"`
CollInterval string `xml:"CollInterval"`
InactiveAfter string `xml:"InactiveAfter"`
Warmuptime string `xml:"Warmuptime"`
TextClass string `xml:"TextClass"`
TextId string `xml:"TextId"`
AttrGroupName string `xml:"AttrGroupName"`
}

func (t ALMTCUST)GetCustStatus()string{
	return t.CustStatus
}

func (t *ALMTCUST)SetCustStatus(val string){
	t.CustStatus = val
}

func (t ALMTCUST)GetUseclass()string{
	return t.Useclass
}

func (t *ALMTCUST)SetUseclass(val string){
	t.Useclass = val
}

func (t ALMTCUST)GetSeverity()string{
	return t.Severity
}

func (t *ALMTCUST)SetSeverity(val string){
	t.Severity = val
}

func (t ALMTCUST)GetStatRecord()string{
	return t.StatRecord
}

func (t *ALMTCUST)SetStatRecord(val string){
	t.StatRecord = val
}

func (t ALMTCUST)GetVisibilityLevel()string{
	return t.VisibilityLevel
}

func (t *ALMTCUST)SetVisibilityLevel(val string){
	t.VisibilityLevel = val
}

func (t ALMTCUST)GetTypeOfAlertsToKeep()string{
	return t.TypeOfAlertsToKeep
}

func (t *ALMTCUST)SetTypeOfAlertsToKeep(val string){
	t.TypeOfAlertsToKeep = val
}

func (t ALMTCUST)GetNumOfAlertsToKeep()string{
	return t.NumOfAlertsToKeep
}

func (t *ALMTCUST)SetNumOfAlertsToKeep(val string){
	t.NumOfAlertsToKeep = val
}

func (t ALMTCUST)GetCollInterval()string{
	return t.CollInterval
}

func (t *ALMTCUST)SetCollInterval(val string){
	t.CollInterval = val
}

func (t ALMTCUST)GetInactiveAfter()string{
	return t.InactiveAfter
}

func (t *ALMTCUST)SetInactiveAfter(val string){
	t.InactiveAfter = val
}

func (t ALMTCUST)GetWarmuptime()string{
	return t.Warmuptime
}

func (t *ALMTCUST)SetWarmuptime(val string){
	t.Warmuptime = val
}

func (t ALMTCUST)GetTextClass()string{
	return t.TextClass
}

func (t *ALMTCUST)SetTextClass(val string){
	t.TextClass = val
}

func (t ALMTCUST)GetTextId()string{
	return t.TextId
}

func (t *ALMTCUST)SetTextId(val string){
	t.TextId = val
}

func (t ALMTCUST)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALMTCUST)SetAttrGroupName(val string){
	t.AttrGroupName = val
}


type  ALPERFCUS struct {
CustStatus string `xml:"CustStatus"`
RelevantValueTyp string `xml:"RelevantValueTyp"`
ThresholdDirection string `xml:"ThresholdDirection"`
ThresholdStatus string `xml:"ThresholdStatus"`
ActiveThresholdG2Y string `xml:"ActiveThresholdG2Y"`
ActiveThresholdY2R string `xml:"ActiveThresholdY2R"`
ActiveThresholdY2G string `xml:"ActiveThresholdY2G"`
ActiveThresholdR2Y string `xml:"ActiveThresholdR2Y"`
DefaultThresholdG2Y string `xml:"DefaultThresholdG2Y"`
DefaultThresholdY2R string `xml:"DefaultThresholdY2R"`
DefaultThresholdY2G string `xml:"DefaultThresholdY2G"`
DefaultThresholdR2Y string `xml:"DefaultThresholdR2Y"`
MsgClass string `xml:"MsgClass"`
MsgId string `xml:"MsgId"`
AttrGroupName string `xml:"AttrGroupName"`
Unit string `xml:"Unit"`
Decimals string `xml:"Decimals"`
}

func (t ALPERFCUS)GetCustStatus()string{
	return t.CustStatus
}

func (t *ALPERFCUS)SetCustStatus(val string){
	t.CustStatus = val
}

func (t ALPERFCUS)GetRelevantValueTyp()string{
	return t.RelevantValueTyp
}

func (t *ALPERFCUS)SetRelevantValueTyp(val string){
	t.RelevantValueTyp = val
}

func (t ALPERFCUS)GetThresholdDirection()string{
	return t.ThresholdDirection
}

func (t *ALPERFCUS)SetThresholdDirection(val string){
	t.ThresholdDirection = val
}

func (t ALPERFCUS)GetThresholdStatus()string{
	return t.ThresholdStatus
}

func (t *ALPERFCUS)SetThresholdStatus(val string){
	t.ThresholdStatus = val
}

func (t ALPERFCUS)GetActiveThresholdG2Y()string{
	return t.ActiveThresholdG2Y
}

func (t *ALPERFCUS)SetActiveThresholdG2Y(val string){
	t.ActiveThresholdG2Y = val
}

func (t ALPERFCUS)GetActiveThresholdY2R()string{
	return t.ActiveThresholdY2R
}

func (t *ALPERFCUS)SetActiveThresholdY2R(val string){
	t.ActiveThresholdY2R = val
}

func (t ALPERFCUS)GetActiveThresholdY2G()string{
	return t.ActiveThresholdY2G
}

func (t *ALPERFCUS)SetActiveThresholdY2G(val string){
	t.ActiveThresholdY2G = val
}

func (t ALPERFCUS)GetActiveThresholdR2Y()string{
	return t.ActiveThresholdR2Y
}

func (t *ALPERFCUS)SetActiveThresholdR2Y(val string){
	t.ActiveThresholdR2Y = val
}

func (t ALPERFCUS)GetDefaultThresholdG2Y()string{
	return t.DefaultThresholdG2Y
}

func (t *ALPERFCUS)SetDefaultThresholdG2Y(val string){
	t.DefaultThresholdG2Y = val
}

func (t ALPERFCUS)GetDefaultThresholdY2R()string{
	return t.DefaultThresholdY2R
}

func (t *ALPERFCUS)SetDefaultThresholdY2R(val string){
	t.DefaultThresholdY2R = val
}

func (t ALPERFCUS)GetDefaultThresholdY2G()string{
	return t.DefaultThresholdY2G
}

func (t *ALPERFCUS)SetDefaultThresholdY2G(val string){
	t.DefaultThresholdY2G = val
}

func (t ALPERFCUS)GetDefaultThresholdR2Y()string{
	return t.DefaultThresholdR2Y
}

func (t *ALPERFCUS)SetDefaultThresholdR2Y(val string){
	t.DefaultThresholdR2Y = val
}

func (t ALPERFCUS)GetMsgClass()string{
	return t.MsgClass
}

func (t *ALPERFCUS)SetMsgClass(val string){
	t.MsgClass = val
}

func (t ALPERFCUS)GetMsgId()string{
	return t.MsgId
}

func (t *ALPERFCUS)SetMsgId(val string){
	t.MsgId = val
}

func (t ALPERFCUS)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALPERFCUS)SetAttrGroupName(val string){
	t.AttrGroupName = val
}

func (t ALPERFCUS)GetUnit()string{
	return t.Unit
}

func (t *ALPERFCUS)SetUnit(val string){
	t.Unit = val
}

func (t ALPERFCUS)GetDecimals()string{
	return t.Decimals
}

func (t *ALPERFCUS)SetDecimals(val string){
	t.Decimals = val
}


type  ALPERFVAL struct {
PerfTotalHigh string `xml:"PerfTotalHigh"`
PerfTotalLow string `xml:"PerfTotalLow"`
PerfCountHigh string `xml:"PerfCountHigh"`
PerfCountLow string `xml:"PerfCountLow"`
AlertRelevantValue string `xml:"AlertRelevantValue"`
AlertRelevantDate string `xml:"AlertRelevantDate"`
AlertRelevantTime string `xml:"AlertRelevantTime"`
LastAlertStatus string `xml:"LastAlertStatus"`
LastPerfValue string `xml:"LastPerfValue"`
Avg00PerfValue string `xml:"Avg00PerfValue"`
Avg01PerfValue string `xml:"Avg01PerfValue"`
Avg05PerfValue string `xml:"Avg05PerfValue"`
Avg15PerfValue string `xml:"Avg15PerfValue"`
Avg01SumValue string `xml:"Avg01SumValue"`
Avg05SumValue string `xml:"Avg05SumValue"`
Avg15SumValue string `xml:"Avg15SumValue"`
Avg01CountValue string `xml:"Avg01CountValue"`
Avg05CountValue string `xml:"Avg05CountValue"`
Avg15CountValue string `xml:"Avg15CountValue"`
MaxPerfValueSeen string `xml:"MaxPerfValueSeen"`
MaxPerfValueDate string `xml:"MaxPerfValueDate"`
MaxPerfValueTime string `xml:"MaxPerfValueTime"`
MinPerfValueSeen string `xml:"MinPerfValueSeen"`
MinPerfValueDate string `xml:"MinPerfValueDate"`
MinPerfValueTime string `xml:"MinPerfValueTime"`
}

func (t ALPERFVAL)GetPerfTotalHigh()string{
	return t.PerfTotalHigh
}

func (t *ALPERFVAL)SetPerfTotalHigh(val string){
	t.PerfTotalHigh = val
}

func (t ALPERFVAL)GetPerfTotalLow()string{
	return t.PerfTotalLow
}

func (t *ALPERFVAL)SetPerfTotalLow(val string){
	t.PerfTotalLow = val
}

func (t ALPERFVAL)GetPerfCountHigh()string{
	return t.PerfCountHigh
}

func (t *ALPERFVAL)SetPerfCountHigh(val string){
	t.PerfCountHigh = val
}

func (t ALPERFVAL)GetPerfCountLow()string{
	return t.PerfCountLow
}

func (t *ALPERFVAL)SetPerfCountLow(val string){
	t.PerfCountLow = val
}

func (t ALPERFVAL)GetAlertRelevantValue()string{
	return t.AlertRelevantValue
}

func (t *ALPERFVAL)SetAlertRelevantValue(val string){
	t.AlertRelevantValue = val
}

func (t ALPERFVAL)GetAlertRelevantDate()string{
	return t.AlertRelevantDate
}

func (t *ALPERFVAL)SetAlertRelevantDate(val string){
	t.AlertRelevantDate = val
}

func (t ALPERFVAL)GetAlertRelevantTime()string{
	return t.AlertRelevantTime
}

func (t *ALPERFVAL)SetAlertRelevantTime(val string){
	t.AlertRelevantTime = val
}

func (t ALPERFVAL)GetLastAlertStatus()string{
	return t.LastAlertStatus
}

func (t *ALPERFVAL)SetLastAlertStatus(val string){
	t.LastAlertStatus = val
}

func (t ALPERFVAL)GetLastPerfValue()string{
	return t.LastPerfValue
}

func (t *ALPERFVAL)SetLastPerfValue(val string){
	t.LastPerfValue = val
}

func (t ALPERFVAL)GetAvg00PerfValue()string{
	return t.Avg00PerfValue
}

func (t *ALPERFVAL)SetAvg00PerfValue(val string){
	t.Avg00PerfValue = val
}

func (t ALPERFVAL)GetAvg01PerfValue()string{
	return t.Avg01PerfValue
}

func (t *ALPERFVAL)SetAvg01PerfValue(val string){
	t.Avg01PerfValue = val
}

func (t ALPERFVAL)GetAvg05PerfValue()string{
	return t.Avg05PerfValue
}

func (t *ALPERFVAL)SetAvg05PerfValue(val string){
	t.Avg05PerfValue = val
}

func (t ALPERFVAL)GetAvg15PerfValue()string{
	return t.Avg15PerfValue
}

func (t *ALPERFVAL)SetAvg15PerfValue(val string){
	t.Avg15PerfValue = val
}

func (t ALPERFVAL)GetAvg01SumValue()string{
	return t.Avg01SumValue
}

func (t *ALPERFVAL)SetAvg01SumValue(val string){
	t.Avg01SumValue = val
}

func (t ALPERFVAL)GetAvg05SumValue()string{
	return t.Avg05SumValue
}

func (t *ALPERFVAL)SetAvg05SumValue(val string){
	t.Avg05SumValue = val
}

func (t ALPERFVAL)GetAvg15SumValue()string{
	return t.Avg15SumValue
}

func (t *ALPERFVAL)SetAvg15SumValue(val string){
	t.Avg15SumValue = val
}

func (t ALPERFVAL)GetAvg01CountValue()string{
	return t.Avg01CountValue
}

func (t *ALPERFVAL)SetAvg01CountValue(val string){
	t.Avg01CountValue = val
}

func (t ALPERFVAL)GetAvg05CountValue()string{
	return t.Avg05CountValue
}

func (t *ALPERFVAL)SetAvg05CountValue(val string){
	t.Avg05CountValue = val
}

func (t ALPERFVAL)GetAvg15CountValue()string{
	return t.Avg15CountValue
}

func (t *ALPERFVAL)SetAvg15CountValue(val string){
	t.Avg15CountValue = val
}

func (t ALPERFVAL)GetMaxPerfValueSeen()string{
	return t.MaxPerfValueSeen
}

func (t *ALPERFVAL)SetMaxPerfValueSeen(val string){
	t.MaxPerfValueSeen = val
}

func (t ALPERFVAL)GetMaxPerfValueDate()string{
	return t.MaxPerfValueDate
}

func (t *ALPERFVAL)SetMaxPerfValueDate(val string){
	t.MaxPerfValueDate = val
}

func (t ALPERFVAL)GetMaxPerfValueTime()string{
	return t.MaxPerfValueTime
}

func (t *ALPERFVAL)SetMaxPerfValueTime(val string){
	t.MaxPerfValueTime = val
}

func (t ALPERFVAL)GetMinPerfValueSeen()string{
	return t.MinPerfValueSeen
}

func (t *ALPERFVAL)SetMinPerfValueSeen(val string){
	t.MinPerfValueSeen = val
}

func (t ALPERFVAL)GetMinPerfValueDate()string{
	return t.MinPerfValueDate
}

func (t *ALPERFVAL)SetMinPerfValueDate(val string){
	t.MinPerfValueDate = val
}

func (t ALPERFVAL)GetMinPerfValueTime()string{
	return t.MinPerfValueTime
}

func (t *ALPERFVAL)SetMinPerfValueTime(val string){
	t.MinPerfValueTime = val
}


type  SXMIMSGRAW struct {
MsgClass string `xml:"MsgClass"`
MsgId string `xml:"MsgId"`
Msgarg1 string `xml:"Msgarg1"`
Argtype1 string `xml:"Argtype1"`
Msgarg2 string `xml:"Msgarg2"`
Argtype2 string `xml:"Argtype2"`
Msgarg3 string `xml:"Msgarg3"`
Argtype3 string `xml:"Argtype3"`
Msgarg4 string `xml:"Msgarg4"`
Argtype4 string `xml:"Argtype4"`
Msgtext string `xml:"Msgtext"`
}

func (t SXMIMSGRAW)GetMsgClass()string{
	return t.MsgClass
}

func (t *SXMIMSGRAW)SetMsgClass(val string){
	t.MsgClass = val
}

func (t SXMIMSGRAW)GetMsgId()string{
	return t.MsgId
}

func (t *SXMIMSGRAW)SetMsgId(val string){
	t.MsgId = val
}

func (t SXMIMSGRAW)GetMsgarg1()string{
	return t.Msgarg1
}

func (t *SXMIMSGRAW)SetMsgarg1(val string){
	t.Msgarg1 = val
}

func (t SXMIMSGRAW)GetArgtype1()string{
	return t.Argtype1
}

func (t *SXMIMSGRAW)SetArgtype1(val string){
	t.Argtype1 = val
}

func (t SXMIMSGRAW)GetMsgarg2()string{
	return t.Msgarg2
}

func (t *SXMIMSGRAW)SetMsgarg2(val string){
	t.Msgarg2 = val
}

func (t SXMIMSGRAW)GetArgtype2()string{
	return t.Argtype2
}

func (t *SXMIMSGRAW)SetArgtype2(val string){
	t.Argtype2 = val
}

func (t SXMIMSGRAW)GetMsgarg3()string{
	return t.Msgarg3
}

func (t *SXMIMSGRAW)SetMsgarg3(val string){
	t.Msgarg3 = val
}

func (t SXMIMSGRAW)GetArgtype3()string{
	return t.Argtype3
}

func (t *SXMIMSGRAW)SetArgtype3(val string){
	t.Argtype3 = val
}

func (t SXMIMSGRAW)GetMsgarg4()string{
	return t.Msgarg4
}

func (t *SXMIMSGRAW)SetMsgarg4(val string){
	t.Msgarg4 = val
}

func (t SXMIMSGRAW)GetArgtype4()string{
	return t.Argtype4
}

func (t *SXMIMSGRAW)SetArgtype4(val string){
	t.Argtype4 = val
}

func (t SXMIMSGRAW)GetMsgtext()string{
	return t.Msgtext
}

func (t *SXMIMSGRAW)SetMsgtext(val string){
	t.Msgtext = val
}


type  ALTOOL struct {
ToolName string `xml:"ToolName"`
ToolStatus string `xml:"ToolStatus"`
StartDate string `xml:"StartDate"`
StartTime string `xml:"StartTime"`
ToolDispatcher string `xml:"ToolDispatcher"`
}

func (t ALTOOL)GetToolName()string{
	return t.ToolName
}

func (t *ALTOOL)SetToolName(val string){
	t.ToolName = val
}

func (t ALTOOL)GetToolStatus()string{
	return t.ToolStatus
}

func (t *ALTOOL)SetToolStatus(val string){
	t.ToolStatus = val
}

func (t ALTOOL)GetStartDate()string{
	return t.StartDate
}

func (t *ALTOOL)SetStartDate(val string){
	t.StartDate = val
}

func (t ALTOOL)GetStartTime()string{
	return t.StartTime
}

func (t *ALTOOL)SetStartTime(val string){
	t.StartTime = val
}

func (t ALTOOL)GetToolDispatcher()string{
	return t.ToolDispatcher
}

func (t *ALTOOL)SetToolDispatcher(val string){
	t.ToolDispatcher = val
}


type  ALMTNAMEL struct {
CompleteName string `xml:"CompleteName"`
}

func (t ALMTNAMEL)GetCompleteName()string{
	return t.CompleteName
}

func (t *ALMTNAMEL)SetCompleteName(val string){
	t.CompleteName = val
}


type  ArrayOfALMTNAMEL struct {
Item []ALMTNAMEL `xml:"item"`
}

func (t ArrayOfALMTNAMEL)GetItem()[]ALMTNAMEL{
	return t.Item
}

func (t *ArrayOfALMTNAMEL)SetItem(val []ALMTNAMEL){
	t.Item = val
}


type  ALGTIDLNRC struct {
Tid ALGLOBTID `xml:"Tid"`
CompleteName string `xml:"CompleteName"`
Rc string `xml:"Rc"`
}

func (t ALGTIDLNRC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALGTIDLNRC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALGTIDLNRC)GetCompleteName()string{
	return t.CompleteName
}

func (t *ALGTIDLNRC)SetCompleteName(val string){
	t.CompleteName = val
}

func (t ALGTIDLNRC)GetRc()string{
	return t.Rc
}

func (t *ALGTIDLNRC)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALGTIDLNRC struct {
Item []ALGTIDLNRC `xml:"item"`
}

func (t ArrayOfALGTIDLNRC)GetItem()[]ALGTIDLNRC{
	return t.Item
}

func (t *ArrayOfALGTIDLNRC)SetItem(val []ALGTIDLNRC){
	t.Item = val
}


type  ALMTEIRC struct {
Rc string `xml:"Rc"`
Tid ALGLOBTID `xml:"Tid"`
MteSubtype string `xml:"MteSubtype"`
Mtename string `xml:"Mtename"`
MteClass string `xml:"MteClass"`
ParentName string `xml:"ParentName"`
ParentClass string `xml:"ParentClass"`
ParentTid ALGLOBTID `xml:"ParentTid"`
}

func (t ALMTEIRC)GetRc()string{
	return t.Rc
}

func (t *ALMTEIRC)SetRc(val string){
	t.Rc = val
}

func (t ALMTEIRC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMTEIRC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMTEIRC)GetMteSubtype()string{
	return t.MteSubtype
}

func (t *ALMTEIRC)SetMteSubtype(val string){
	t.MteSubtype = val
}

func (t ALMTEIRC)GetMtename()string{
	return t.Mtename
}

func (t *ALMTEIRC)SetMtename(val string){
	t.Mtename = val
}

func (t ALMTEIRC)GetMteClass()string{
	return t.MteClass
}

func (t *ALMTEIRC)SetMteClass(val string){
	t.MteClass = val
}

func (t ALMTEIRC)GetParentName()string{
	return t.ParentName
}

func (t *ALMTEIRC)SetParentName(val string){
	t.ParentName = val
}

func (t ALMTEIRC)GetParentClass()string{
	return t.ParentClass
}

func (t *ALMTEIRC)SetParentClass(val string){
	t.ParentClass = val
}

func (t ALMTEIRC)GetParentTid()ALGLOBTID{
	return t.ParentTid
}

func (t *ALMTEIRC)SetParentTid(val ALGLOBTID){
	t.ParentTid = val
}


type  ArrayOfALMTEIRC struct {
Item []ALMTEIRC `xml:"item"`
}

func (t ArrayOfALMTEIRC)GetItem()[]ALMTEIRC{
	return t.Item
}

func (t *ArrayOfALMTEIRC)SetItem(val []ALMTEIRC){
	t.Item = val
}


type  ALTDEFRC struct {
Rc string `xml:"Rc"`
Tid ALGLOBTID `xml:"Tid"`
GeneralCust ALMTCUST `xml:"GeneralCust"`
DeliveryStatus string `xml:"DeliveryStatus"`
HighAlertValue string `xml:"HighAlertValue"`
HighAlertSeverity string `xml:"HighAlertSeverity"`
Aid ALGLOBAID `xml:"Aid"`
CurrentValue string `xml:"CurrentValue"`
CurrentSeverity string `xml:"CurrentSeverity"`
CurrentValueAid ALGLOBAID `xml:"CurrentValueAid"`
ActiveAlerts string `xml:"ActiveAlerts"`
AlertsSinceStartup string `xml:"AlertsSinceStartup"`
TypDefStatus string `xml:"TypDefStatus"`
CollToolDefName string `xml:"CollToolDefName"`
CollToolDefDispatcher string `xml:"CollToolDefDispatcher"`
CollToolDefStatus string `xml:"CollToolDefStatus"`
CollToolDefUseClass string `xml:"CollToolDefUseClass"`
AnalysisToolName string `xml:"AnalysisToolName"`
AnalysisToolDispatcher string `xml:"AnalysisToolDispatcher"`
AnalysisToolStatus string `xml:"AnalysisToolStatus"`
AnalysisToolUseClass string `xml:"AnalysisToolUseClass"`
AutoreactionDefName string `xml:"AutoreactionDefName"`
AutoreactionDefDispatcher string `xml:"AutoreactionDefDispatcher"`
AutoreactionDefStatus string `xml:"AutoreactionDefStatus"`
AutoreactionDefUseClass string `xml:"AutoreactionDefUseClass"`
CollToolRunName string `xml:"CollToolRunName"`
CollToolRunStatus string `xml:"CollToolRunStatus"`
CollToolRunDate string `xml:"CollToolRunDate"`
CollToolRunTime string `xml:"CollToolRunTime"`
CollToolRunDispatcher string `xml:"CollToolRunDispatcher"`
AutoreactionRunName string `xml:"AutoreactionRunName"`
AutoreactionRunStatus string `xml:"AutoreactionRunStatus"`
AutoreactionRunDate string `xml:"AutoreactionRunDate"`
AutoreactionRunTime string `xml:"AutoreactionRunTime"`
AutoreactionRunDispatcher string `xml:"AutoreactionRunDispatcher"`
Objectname string `xml:"Objectname"`
ShortName string `xml:"ShortName"`
MteSubtype string `xml:"MteSubtype"`
AlertsInList string `xml:"AlertsInList"`
AlertTypChildCount string `xml:"AlertTypChildCount"`
}

func (t ALTDEFRC)GetRc()string{
	return t.Rc
}

func (t *ALTDEFRC)SetRc(val string){
	t.Rc = val
}

func (t ALTDEFRC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTDEFRC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTDEFRC)GetGeneralCust()ALMTCUST{
	return t.GeneralCust
}

func (t *ALTDEFRC)SetGeneralCust(val ALMTCUST){
	t.GeneralCust = val
}

func (t ALTDEFRC)GetDeliveryStatus()string{
	return t.DeliveryStatus
}

func (t *ALTDEFRC)SetDeliveryStatus(val string){
	t.DeliveryStatus = val
}

func (t ALTDEFRC)GetHighAlertValue()string{
	return t.HighAlertValue
}

func (t *ALTDEFRC)SetHighAlertValue(val string){
	t.HighAlertValue = val
}

func (t ALTDEFRC)GetHighAlertSeverity()string{
	return t.HighAlertSeverity
}

func (t *ALTDEFRC)SetHighAlertSeverity(val string){
	t.HighAlertSeverity = val
}

func (t ALTDEFRC)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALTDEFRC)SetAid(val ALGLOBAID){
	t.Aid = val
}

func (t ALTDEFRC)GetCurrentValue()string{
	return t.CurrentValue
}

func (t *ALTDEFRC)SetCurrentValue(val string){
	t.CurrentValue = val
}

func (t ALTDEFRC)GetCurrentSeverity()string{
	return t.CurrentSeverity
}

func (t *ALTDEFRC)SetCurrentSeverity(val string){
	t.CurrentSeverity = val
}

func (t ALTDEFRC)GetCurrentValueAid()ALGLOBAID{
	return t.CurrentValueAid
}

func (t *ALTDEFRC)SetCurrentValueAid(val ALGLOBAID){
	t.CurrentValueAid = val
}

func (t ALTDEFRC)GetActiveAlerts()string{
	return t.ActiveAlerts
}

func (t *ALTDEFRC)SetActiveAlerts(val string){
	t.ActiveAlerts = val
}

func (t ALTDEFRC)GetAlertsSinceStartup()string{
	return t.AlertsSinceStartup
}

func (t *ALTDEFRC)SetAlertsSinceStartup(val string){
	t.AlertsSinceStartup = val
}

func (t ALTDEFRC)GetTypDefStatus()string{
	return t.TypDefStatus
}

func (t *ALTDEFRC)SetTypDefStatus(val string){
	t.TypDefStatus = val
}

func (t ALTDEFRC)GetCollToolDefName()string{
	return t.CollToolDefName
}

func (t *ALTDEFRC)SetCollToolDefName(val string){
	t.CollToolDefName = val
}

func (t ALTDEFRC)GetCollToolDefDispatcher()string{
	return t.CollToolDefDispatcher
}

func (t *ALTDEFRC)SetCollToolDefDispatcher(val string){
	t.CollToolDefDispatcher = val
}

func (t ALTDEFRC)GetCollToolDefStatus()string{
	return t.CollToolDefStatus
}

func (t *ALTDEFRC)SetCollToolDefStatus(val string){
	t.CollToolDefStatus = val
}

func (t ALTDEFRC)GetCollToolDefUseClass()string{
	return t.CollToolDefUseClass
}

func (t *ALTDEFRC)SetCollToolDefUseClass(val string){
	t.CollToolDefUseClass = val
}

func (t ALTDEFRC)GetAnalysisToolName()string{
	return t.AnalysisToolName
}

func (t *ALTDEFRC)SetAnalysisToolName(val string){
	t.AnalysisToolName = val
}

func (t ALTDEFRC)GetAnalysisToolDispatcher()string{
	return t.AnalysisToolDispatcher
}

func (t *ALTDEFRC)SetAnalysisToolDispatcher(val string){
	t.AnalysisToolDispatcher = val
}

func (t ALTDEFRC)GetAnalysisToolStatus()string{
	return t.AnalysisToolStatus
}

func (t *ALTDEFRC)SetAnalysisToolStatus(val string){
	t.AnalysisToolStatus = val
}

func (t ALTDEFRC)GetAnalysisToolUseClass()string{
	return t.AnalysisToolUseClass
}

func (t *ALTDEFRC)SetAnalysisToolUseClass(val string){
	t.AnalysisToolUseClass = val
}

func (t ALTDEFRC)GetAutoreactionDefName()string{
	return t.AutoreactionDefName
}

func (t *ALTDEFRC)SetAutoreactionDefName(val string){
	t.AutoreactionDefName = val
}

func (t ALTDEFRC)GetAutoreactionDefDispatcher()string{
	return t.AutoreactionDefDispatcher
}

func (t *ALTDEFRC)SetAutoreactionDefDispatcher(val string){
	t.AutoreactionDefDispatcher = val
}

func (t ALTDEFRC)GetAutoreactionDefStatus()string{
	return t.AutoreactionDefStatus
}

func (t *ALTDEFRC)SetAutoreactionDefStatus(val string){
	t.AutoreactionDefStatus = val
}

func (t ALTDEFRC)GetAutoreactionDefUseClass()string{
	return t.AutoreactionDefUseClass
}

func (t *ALTDEFRC)SetAutoreactionDefUseClass(val string){
	t.AutoreactionDefUseClass = val
}

func (t ALTDEFRC)GetCollToolRunName()string{
	return t.CollToolRunName
}

func (t *ALTDEFRC)SetCollToolRunName(val string){
	t.CollToolRunName = val
}

func (t ALTDEFRC)GetCollToolRunStatus()string{
	return t.CollToolRunStatus
}

func (t *ALTDEFRC)SetCollToolRunStatus(val string){
	t.CollToolRunStatus = val
}

func (t ALTDEFRC)GetCollToolRunDate()string{
	return t.CollToolRunDate
}

func (t *ALTDEFRC)SetCollToolRunDate(val string){
	t.CollToolRunDate = val
}

func (t ALTDEFRC)GetCollToolRunTime()string{
	return t.CollToolRunTime
}

func (t *ALTDEFRC)SetCollToolRunTime(val string){
	t.CollToolRunTime = val
}

func (t ALTDEFRC)GetCollToolRunDispatcher()string{
	return t.CollToolRunDispatcher
}

func (t *ALTDEFRC)SetCollToolRunDispatcher(val string){
	t.CollToolRunDispatcher = val
}

func (t ALTDEFRC)GetAutoreactionRunName()string{
	return t.AutoreactionRunName
}

func (t *ALTDEFRC)SetAutoreactionRunName(val string){
	t.AutoreactionRunName = val
}

func (t ALTDEFRC)GetAutoreactionRunStatus()string{
	return t.AutoreactionRunStatus
}

func (t *ALTDEFRC)SetAutoreactionRunStatus(val string){
	t.AutoreactionRunStatus = val
}

func (t ALTDEFRC)GetAutoreactionRunDate()string{
	return t.AutoreactionRunDate
}

func (t *ALTDEFRC)SetAutoreactionRunDate(val string){
	t.AutoreactionRunDate = val
}

func (t ALTDEFRC)GetAutoreactionRunTime()string{
	return t.AutoreactionRunTime
}

func (t *ALTDEFRC)SetAutoreactionRunTime(val string){
	t.AutoreactionRunTime = val
}

func (t ALTDEFRC)GetAutoreactionRunDispatcher()string{
	return t.AutoreactionRunDispatcher
}

func (t *ALTDEFRC)SetAutoreactionRunDispatcher(val string){
	t.AutoreactionRunDispatcher = val
}

func (t ALTDEFRC)GetObjectname()string{
	return t.Objectname
}

func (t *ALTDEFRC)SetObjectname(val string){
	t.Objectname = val
}

func (t ALTDEFRC)GetShortName()string{
	return t.ShortName
}

func (t *ALTDEFRC)SetShortName(val string){
	t.ShortName = val
}

func (t ALTDEFRC)GetMteSubtype()string{
	return t.MteSubtype
}

func (t *ALTDEFRC)SetMteSubtype(val string){
	t.MteSubtype = val
}

func (t ALTDEFRC)GetAlertsInList()string{
	return t.AlertsInList
}

func (t *ALTDEFRC)SetAlertsInList(val string){
	t.AlertsInList = val
}

func (t ALTDEFRC)GetAlertTypChildCount()string{
	return t.AlertTypChildCount
}

func (t *ALTDEFRC)SetAlertTypChildCount(val string){
	t.AlertTypChildCount = val
}


type  ArrayOfALTDEFRC struct {
Item []ALTDEFRC `xml:"item"`
}

func (t ArrayOfALTDEFRC)GetItem()[]ALTDEFRC{
	return t.Item
}

func (t *ArrayOfALTDEFRC)SetItem(val []ALTDEFRC){
	t.Item = val
}


type  ALPERFTYPE struct {
Tid ALGLOBTID `xml:"Tid"`
PerfCustomizing ALPERFCUS `xml:"PerfCustomizing"`
PerfValue ALPERFVAL `xml:"PerfValue"`
Rc string `xml:"Rc"`
}

func (t ALPERFTYPE)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALPERFTYPE)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALPERFTYPE)GetPerfCustomizing()ALPERFCUS{
	return t.PerfCustomizing
}

func (t *ALPERFTYPE)SetPerfCustomizing(val ALPERFCUS){
	t.PerfCustomizing = val
}

func (t ALPERFTYPE)GetPerfValue()ALPERFVAL{
	return t.PerfValue
}

func (t *ALPERFTYPE)SetPerfValue(val ALPERFVAL){
	t.PerfValue = val
}

func (t ALPERFTYPE)GetRc()string{
	return t.Rc
}

func (t *ALPERFTYPE)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALPERFTYPE struct {
Item []ALPERFTYPE `xml:"item"`
}

func (t ArrayOfALPERFTYPE)GetItem()[]ALPERFTYPE{
	return t.Item
}

func (t *ArrayOfALPERFTYPE)SetItem(val []ALPERFTYPE){
	t.Item = val
}


type  ALGTIDSMO struct {
Tid ALGLOBTID `xml:"Tid"`
SmoothingType string `xml:"SmoothingType"`
SmoothDate string `xml:"SmoothDate"`
SmoothTime string `xml:"SmoothTime"`
SmoothSumValue string `xml:"SmoothSumValue"`
SmoothNumValues string `xml:"SmoothNumValues"`
SmoothAvgValue string `xml:"SmoothAvgValue"`
Rc string `xml:"Rc"`
}

func (t ALGTIDSMO)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALGTIDSMO)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALGTIDSMO)GetSmoothingType()string{
	return t.SmoothingType
}

func (t *ALGTIDSMO)SetSmoothingType(val string){
	t.SmoothingType = val
}

func (t ALGTIDSMO)GetSmoothDate()string{
	return t.SmoothDate
}

func (t *ALGTIDSMO)SetSmoothDate(val string){
	t.SmoothDate = val
}

func (t ALGTIDSMO)GetSmoothTime()string{
	return t.SmoothTime
}

func (t *ALGTIDSMO)SetSmoothTime(val string){
	t.SmoothTime = val
}

func (t ALGTIDSMO)GetSmoothSumValue()string{
	return t.SmoothSumValue
}

func (t *ALGTIDSMO)SetSmoothSumValue(val string){
	t.SmoothSumValue = val
}

func (t ALGTIDSMO)GetSmoothNumValues()string{
	return t.SmoothNumValues
}

func (t *ALGTIDSMO)SetSmoothNumValues(val string){
	t.SmoothNumValues = val
}

func (t ALGTIDSMO)GetSmoothAvgValue()string{
	return t.SmoothAvgValue
}

func (t *ALGTIDSMO)SetSmoothAvgValue(val string){
	t.SmoothAvgValue = val
}

func (t ALGTIDSMO)GetRc()string{
	return t.Rc
}

func (t *ALGTIDSMO)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALGTIDSMO struct {
Item []ALGTIDSMO `xml:"item"`
}

func (t ArrayOfALGTIDSMO)GetItem()[]ALGTIDSMO{
	return t.Item
}

func (t *ArrayOfALGTIDSMO)SetItem(val []ALGTIDSMO){
	t.Item = val
}


type  ALMSCTIDRL struct {
TidNumber string `xml:"TidNumber"`
Tid ALGLOBTID `xml:"Tid"`
MscLineId string `xml:"MscLineId"`
MscDate string `xml:"MscDate"`
MscTime string `xml:"MscTime"`
ValueOrig string `xml:"ValueOrig"`
SeverityOrig string `xml:"SeverityOrig"`
ValueFilter string `xml:"ValueFilter"`
SeverityFilter string `xml:"SeverityFilter"`
ABAPClient string `xml:"ABAPClient"`
User string `xml:"User"`
Message SXMIMSGRAW `xml:"Message"`
}

func (t ALMSCTIDRL)GetTidNumber()string{
	return t.TidNumber
}

func (t *ALMSCTIDRL)SetTidNumber(val string){
	t.TidNumber = val
}

func (t ALMSCTIDRL)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMSCTIDRL)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMSCTIDRL)GetMscLineId()string{
	return t.MscLineId
}

func (t *ALMSCTIDRL)SetMscLineId(val string){
	t.MscLineId = val
}

func (t ALMSCTIDRL)GetMscDate()string{
	return t.MscDate
}

func (t *ALMSCTIDRL)SetMscDate(val string){
	t.MscDate = val
}

func (t ALMSCTIDRL)GetMscTime()string{
	return t.MscTime
}

func (t *ALMSCTIDRL)SetMscTime(val string){
	t.MscTime = val
}

func (t ALMSCTIDRL)GetValueOrig()string{
	return t.ValueOrig
}

func (t *ALMSCTIDRL)SetValueOrig(val string){
	t.ValueOrig = val
}

func (t ALMSCTIDRL)GetSeverityOrig()string{
	return t.SeverityOrig
}

func (t *ALMSCTIDRL)SetSeverityOrig(val string){
	t.SeverityOrig = val
}

func (t ALMSCTIDRL)GetValueFilter()string{
	return t.ValueFilter
}

func (t *ALMSCTIDRL)SetValueFilter(val string){
	t.ValueFilter = val
}

func (t ALMSCTIDRL)GetSeverityFilter()string{
	return t.SeverityFilter
}

func (t *ALMSCTIDRL)SetSeverityFilter(val string){
	t.SeverityFilter = val
}

func (t ALMSCTIDRL)GetABAPClient()string{
	return t.ABAPClient
}

func (t *ALMSCTIDRL)SetABAPClient(val string){
	t.ABAPClient = val
}

func (t ALMSCTIDRL)GetUser()string{
	return t.User
}

func (t *ALMSCTIDRL)SetUser(val string){
	t.User = val
}

func (t ALMSCTIDRL)GetMessage()SXMIMSGRAW{
	return t.Message
}

func (t *ALMSCTIDRL)SetMessage(val SXMIMSGRAW){
	t.Message = val
}


type  ArrayOfALMSCTIDRL struct {
Item []ALMSCTIDRL `xml:"item"`
}

func (t ArrayOfALMSCTIDRL)GetItem()[]ALMSCTIDRL{
	return t.Item
}

func (t *ArrayOfALMSCTIDRL)SetItem(val []ALMSCTIDRL){
	t.Item = val
}


type  ALMSCSELEC struct {
Tid ALGLOBTID `xml:"Tid"`
FirstLineNum string `xml:"FirstLineNum"`
LastLineNum string `xml:"LastLineNum"`
StartDate string `xml:"StartDate"`
StartTime string `xml:"StartTime"`
EndDate string `xml:"EndDate"`
EndTime string `xml:"EndTime"`
SecBeforeFirstLineNum string `xml:"SecBeforeFirstLineNum"`
SecAfterLastLineNum string `xml:"SecAfterLastLineNum"`
ABAPClient string `xml:"ABAPClient"`
User string `xml:"User"`
}

func (t ALMSCSELEC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMSCSELEC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMSCSELEC)GetFirstLineNum()string{
	return t.FirstLineNum
}

func (t *ALMSCSELEC)SetFirstLineNum(val string){
	t.FirstLineNum = val
}

func (t ALMSCSELEC)GetLastLineNum()string{
	return t.LastLineNum
}

func (t *ALMSCSELEC)SetLastLineNum(val string){
	t.LastLineNum = val
}

func (t ALMSCSELEC)GetStartDate()string{
	return t.StartDate
}

func (t *ALMSCSELEC)SetStartDate(val string){
	t.StartDate = val
}

func (t ALMSCSELEC)GetStartTime()string{
	return t.StartTime
}

func (t *ALMSCSELEC)SetStartTime(val string){
	t.StartTime = val
}

func (t ALMSCSELEC)GetEndDate()string{
	return t.EndDate
}

func (t *ALMSCSELEC)SetEndDate(val string){
	t.EndDate = val
}

func (t ALMSCSELEC)GetEndTime()string{
	return t.EndTime
}

func (t *ALMSCSELEC)SetEndTime(val string){
	t.EndTime = val
}

func (t ALMSCSELEC)GetSecBeforeFirstLineNum()string{
	return t.SecBeforeFirstLineNum
}

func (t *ALMSCSELEC)SetSecBeforeFirstLineNum(val string){
	t.SecBeforeFirstLineNum = val
}

func (t ALMSCSELEC)GetSecAfterLastLineNum()string{
	return t.SecAfterLastLineNum
}

func (t *ALMSCSELEC)SetSecAfterLastLineNum(val string){
	t.SecAfterLastLineNum = val
}

func (t ALMSCSELEC)GetABAPClient()string{
	return t.ABAPClient
}

func (t *ALMSCSELEC)SetABAPClient(val string){
	t.ABAPClient = val
}

func (t ALMSCSELEC)GetUser()string{
	return t.User
}

func (t *ALMSCSELEC)SetUser(val string){
	t.User = val
}


type  ArrayOfALMSCSELEC struct {
Item []ALMSCSELEC `xml:"item"`
}

func (t ArrayOfALMSCSELEC)GetItem()[]ALMSCSELEC{
	return t.Item
}

func (t *ArrayOfALMSCSELEC)SetItem(val []ALMSCSELEC){
	t.Item = val
}


type  ALGTIDRC struct {
Tid ALGLOBTID `xml:"Tid"`
Rc string `xml:"Rc"`
}

func (t ALGTIDRC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALGTIDRC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALGTIDRC)GetRc()string{
	return t.Rc
}

func (t *ALGTIDRC)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALGTIDRC struct {
Item []ALGTIDRC `xml:"item"`
}

func (t ArrayOfALGTIDRC)GetItem()[]ALGTIDRC{
	return t.Item
}

func (t *ArrayOfALGTIDRC)SetItem(val []ALGTIDRC){
	t.Item = val
}


type  ALSMSGRAWT struct {
Tid ALGLOBTID `xml:"Tid"`
CustStatus string `xml:"CustStatus"`
AlertMode string `xml:"AlertMode"`
AlertShift string `xml:"AlertShift"`
AttrGroupName string `xml:"AttrGroupName"`
MsgValue string `xml:"MsgValue"`
MsgDate string `xml:"MsgDate"`
MsgTime string `xml:"MsgTime"`
Message SXMIMSGRAW `xml:"Message"`
Rc string `xml:"Rc"`
}

func (t ALSMSGRAWT)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALSMSGRAWT)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALSMSGRAWT)GetCustStatus()string{
	return t.CustStatus
}

func (t *ALSMSGRAWT)SetCustStatus(val string){
	t.CustStatus = val
}

func (t ALSMSGRAWT)GetAlertMode()string{
	return t.AlertMode
}

func (t *ALSMSGRAWT)SetAlertMode(val string){
	t.AlertMode = val
}

func (t ALSMSGRAWT)GetAlertShift()string{
	return t.AlertShift
}

func (t *ALSMSGRAWT)SetAlertShift(val string){
	t.AlertShift = val
}

func (t ALSMSGRAWT)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALSMSGRAWT)SetAttrGroupName(val string){
	t.AttrGroupName = val
}

func (t ALSMSGRAWT)GetMsgValue()string{
	return t.MsgValue
}

func (t *ALSMSGRAWT)SetMsgValue(val string){
	t.MsgValue = val
}

func (t ALSMSGRAWT)GetMsgDate()string{
	return t.MsgDate
}

func (t *ALSMSGRAWT)SetMsgDate(val string){
	t.MsgDate = val
}

func (t ALSMSGRAWT)GetMsgTime()string{
	return t.MsgTime
}

func (t *ALSMSGRAWT)SetMsgTime(val string){
	t.MsgTime = val
}

func (t ALSMSGRAWT)GetMessage()SXMIMSGRAW{
	return t.Message
}

func (t *ALSMSGRAWT)SetMessage(val SXMIMSGRAW){
	t.Message = val
}

func (t ALSMSGRAWT)GetRc()string{
	return t.Rc
}

func (t *ALSMSGRAWT)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALSMSGRAWT struct {
Item []ALSMSGRAWT `xml:"item"`
}

func (t ArrayOfALSMSGRAWT)GetItem()[]ALSMSGRAWT{
	return t.Item
}

func (t *ArrayOfALSMSGRAWT)SetItem(val []ALSMSGRAWT){
	t.Item = val
}


type  ALTEXTATTR struct {
Tid ALGLOBTID `xml:"Tid"`
AttributeText string `xml:"AttributeText"`
Rc string `xml:"Rc"`
}

func (t ALTEXTATTR)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTEXTATTR)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTEXTATTR)GetAttributeText()string{
	return t.AttributeText
}

func (t *ALTEXTATTR)SetAttributeText(val string){
	t.AttributeText = val
}

func (t ALTEXTATTR)GetRc()string{
	return t.Rc
}

func (t *ALTEXTATTR)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALTEXTATTR struct {
Item []ALTEXTATTR `xml:"item"`
}

func (t ArrayOfALTEXTATTR)GetItem()[]ALTEXTATTR{
	return t.Item
}

func (t *ArrayOfALTEXTATTR)SetItem(val []ALTEXTATTR){
	t.Item = val
}


type  ALTXTLNKAT struct {
Tid ALGLOBTID `xml:"Tid"`
MteSubtype string `xml:"MteSubtype"`
AttributeText string `xml:"AttributeText"`
LinkedTid ALGLOBTID `xml:"LinkedTid"`
Rc string `xml:"Rc"`
}

func (t ALTXTLNKAT)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTXTLNKAT)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTXTLNKAT)GetMteSubtype()string{
	return t.MteSubtype
}

func (t *ALTXTLNKAT)SetMteSubtype(val string){
	t.MteSubtype = val
}

func (t ALTXTLNKAT)GetAttributeText()string{
	return t.AttributeText
}

func (t *ALTXTLNKAT)SetAttributeText(val string){
	t.AttributeText = val
}

func (t ALTXTLNKAT)GetLinkedTid()ALGLOBTID{
	return t.LinkedTid
}

func (t *ALTXTLNKAT)SetLinkedTid(val ALGLOBTID){
	t.LinkedTid = val
}

func (t ALTXTLNKAT)GetRc()string{
	return t.Rc
}

func (t *ALTXTLNKAT)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALTXTLNKAT struct {
Item []ALTXTLNKAT `xml:"item"`
}

func (t ArrayOfALTXTLNKAT)GetItem()[]ALTXTLNKAT{
	return t.Item
}

func (t *ArrayOfALTXTLNKAT)SetItem(val []ALTXTLNKAT){
	t.Item = val
}


type  ALCCTIDAID struct {
TidNumber string `xml:"TidNumber"`
Tid ALGLOBTID `xml:"Tid"`
Aid ALGLOBAID `xml:"Aid"`
}

func (t ALCCTIDAID)GetTidNumber()string{
	return t.TidNumber
}

func (t *ALCCTIDAID)SetTidNumber(val string){
	t.TidNumber = val
}

func (t ALCCTIDAID)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALCCTIDAID)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALCCTIDAID)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALCCTIDAID)SetAid(val ALGLOBAID){
	t.Aid = val
}


type  ArrayOfALCCTIDAID struct {
Item []ALCCTIDAID `xml:"item"`
}

func (t ArrayOfALCCTIDAID)GetItem()[]ALCCTIDAID{
	return t.Item
}

func (t *ArrayOfALCCTIDAID)SetItem(val []ALCCTIDAID){
	t.Item = val
}


type  ALALRAWRC struct {
Rc string `xml:"Rc"`
Aid ALGLOBAID `xml:"Aid"`
Tid ALGLOBTID `xml:"Tid"`
Value string `xml:"Value"`
Severity string `xml:"Severity"`
Status string `xml:"Status"`
Objectname string `xml:"Objectname"`
Fieldname string `xml:"Fieldname"`
GoneDate string `xml:"GoneDate"`
GoneTime string `xml:"GoneTime"`
ReportedBy string `xml:"ReportedBy"`
StatChangeDate string `xml:"StatChangeDate"`
StatChangeTime string `xml:"StatChangeTime"`
ChangedBy string `xml:"ChangedBy"`
TimeoutData string `xml:"TimeoutData"`
TimeoutTime string `xml:"TimeoutTime"`
Message SXMIMSGRAW `xml:"Message"`
ABAPClient string `xml:"ABAPClient"`
User string `xml:"User"`
MscLineId string `xml:"MscLineId"`
}

func (t ALALRAWRC)GetRc()string{
	return t.Rc
}

func (t *ALALRAWRC)SetRc(val string){
	t.Rc = val
}

func (t ALALRAWRC)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALALRAWRC)SetAid(val ALGLOBAID){
	t.Aid = val
}

func (t ALALRAWRC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALALRAWRC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALALRAWRC)GetValue()string{
	return t.Value
}

func (t *ALALRAWRC)SetValue(val string){
	t.Value = val
}

func (t ALALRAWRC)GetSeverity()string{
	return t.Severity
}

func (t *ALALRAWRC)SetSeverity(val string){
	t.Severity = val
}

func (t ALALRAWRC)GetStatus()string{
	return t.Status
}

func (t *ALALRAWRC)SetStatus(val string){
	t.Status = val
}

func (t ALALRAWRC)GetObjectname()string{
	return t.Objectname
}

func (t *ALALRAWRC)SetObjectname(val string){
	t.Objectname = val
}

func (t ALALRAWRC)GetFieldname()string{
	return t.Fieldname
}

func (t *ALALRAWRC)SetFieldname(val string){
	t.Fieldname = val
}

func (t ALALRAWRC)GetGoneDate()string{
	return t.GoneDate
}

func (t *ALALRAWRC)SetGoneDate(val string){
	t.GoneDate = val
}

func (t ALALRAWRC)GetGoneTime()string{
	return t.GoneTime
}

func (t *ALALRAWRC)SetGoneTime(val string){
	t.GoneTime = val
}

func (t ALALRAWRC)GetReportedBy()string{
	return t.ReportedBy
}

func (t *ALALRAWRC)SetReportedBy(val string){
	t.ReportedBy = val
}

func (t ALALRAWRC)GetStatChangeDate()string{
	return t.StatChangeDate
}

func (t *ALALRAWRC)SetStatChangeDate(val string){
	t.StatChangeDate = val
}

func (t ALALRAWRC)GetStatChangeTime()string{
	return t.StatChangeTime
}

func (t *ALALRAWRC)SetStatChangeTime(val string){
	t.StatChangeTime = val
}

func (t ALALRAWRC)GetChangedBy()string{
	return t.ChangedBy
}

func (t *ALALRAWRC)SetChangedBy(val string){
	t.ChangedBy = val
}

func (t ALALRAWRC)GetTimeoutData()string{
	return t.TimeoutData
}

func (t *ALALRAWRC)SetTimeoutData(val string){
	t.TimeoutData = val
}

func (t ALALRAWRC)GetTimeoutTime()string{
	return t.TimeoutTime
}

func (t *ALALRAWRC)SetTimeoutTime(val string){
	t.TimeoutTime = val
}

func (t ALALRAWRC)GetMessage()SXMIMSGRAW{
	return t.Message
}

func (t *ALALRAWRC)SetMessage(val SXMIMSGRAW){
	t.Message = val
}

func (t ALALRAWRC)GetABAPClient()string{
	return t.ABAPClient
}

func (t *ALALRAWRC)SetABAPClient(val string){
	t.ABAPClient = val
}

func (t ALALRAWRC)GetUser()string{
	return t.User
}

func (t *ALALRAWRC)SetUser(val string){
	t.User = val
}

func (t ALALRAWRC)GetMscLineId()string{
	return t.MscLineId
}

func (t *ALALRAWRC)SetMscLineId(val string){
	t.MscLineId = val
}


type  ArrayOfALALRAWRC struct {
Item []ALALRAWRC `xml:"item"`
}

func (t ArrayOfALALRAWRC)GetItem()[]ALALRAWRC{
	return t.Item
}

func (t *ArrayOfALALRAWRC)SetItem(val []ALALRAWRC){
	t.Item = val
}


type  ArrayOfALGLOBAID struct {
Item []ALGLOBAID `xml:"item"`
}

func (t ArrayOfALGLOBAID)GetItem()[]ALGLOBAID{
	return t.Item
}

func (t *ArrayOfALGLOBAID)SetItem(val []ALGLOBAID){
	t.Item = val
}


type  ALGTIDWHTL struct {
Tid ALGLOBTID `xml:"Tid"`
Whichtool string `xml:"Whichtool"`
}

func (t ALGTIDWHTL)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALGTIDWHTL)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALGTIDWHTL)GetWhichtool()string{
	return t.Whichtool
}

func (t *ALGTIDWHTL)SetWhichtool(val string){
	t.Whichtool = val
}


type  ArrayOfALGTIDWHTL struct {
Item []ALGTIDWHTL `xml:"item"`
}

func (t ArrayOfALGTIDWHTL)GetItem()[]ALGTIDWHTL{
	return t.Item
}

func (t *ArrayOfALGTIDWHTL)SetItem(val []ALGTIDWHTL){
	t.Item = val
}


type  ALTOOLEFRC struct {
Tid ALGLOBTID `xml:"Tid"`
ChildClass string `xml:"ChildClass"`
Whichtool string `xml:"Whichtool"`
Toolname string `xml:"Toolname"`
ToolDispatcher string `xml:"ToolDispatcher"`
DefStatus string `xml:"DefStatus"`
Useclass string `xml:"Useclass"`
ParentTid ALGLOBTID `xml:"ParentTid"`
ParentObjectName string `xml:"ParentObjectName"`
ParentShortName string `xml:"ParentShortName"`
ParentClass string `xml:"ParentClass"`
Rc string `xml:"Rc"`
}

func (t ALTOOLEFRC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTOOLEFRC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTOOLEFRC)GetChildClass()string{
	return t.ChildClass
}

func (t *ALTOOLEFRC)SetChildClass(val string){
	t.ChildClass = val
}

func (t ALTOOLEFRC)GetWhichtool()string{
	return t.Whichtool
}

func (t *ALTOOLEFRC)SetWhichtool(val string){
	t.Whichtool = val
}

func (t ALTOOLEFRC)GetToolname()string{
	return t.Toolname
}

func (t *ALTOOLEFRC)SetToolname(val string){
	t.Toolname = val
}

func (t ALTOOLEFRC)GetToolDispatcher()string{
	return t.ToolDispatcher
}

func (t *ALTOOLEFRC)SetToolDispatcher(val string){
	t.ToolDispatcher = val
}

func (t ALTOOLEFRC)GetDefStatus()string{
	return t.DefStatus
}

func (t *ALTOOLEFRC)SetDefStatus(val string){
	t.DefStatus = val
}

func (t ALTOOLEFRC)GetUseclass()string{
	return t.Useclass
}

func (t *ALTOOLEFRC)SetUseclass(val string){
	t.Useclass = val
}

func (t ALTOOLEFRC)GetParentTid()ALGLOBTID{
	return t.ParentTid
}

func (t *ALTOOLEFRC)SetParentTid(val ALGLOBTID){
	t.ParentTid = val
}

func (t ALTOOLEFRC)GetParentObjectName()string{
	return t.ParentObjectName
}

func (t *ALTOOLEFRC)SetParentObjectName(val string){
	t.ParentObjectName = val
}

func (t ALTOOLEFRC)GetParentShortName()string{
	return t.ParentShortName
}

func (t *ALTOOLEFRC)SetParentShortName(val string){
	t.ParentShortName = val
}

func (t ALTOOLEFRC)GetParentClass()string{
	return t.ParentClass
}

func (t *ALTOOLEFRC)SetParentClass(val string){
	t.ParentClass = val
}

func (t ALTOOLEFRC)GetRc()string{
	return t.Rc
}

func (t *ALTOOLEFRC)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALTOOLEFRC struct {
Item []ALTOOLEFRC `xml:"item"`
}

func (t ArrayOfALTOOLEFRC)GetItem()[]ALTOOLEFRC{
	return t.Item
}

func (t *ArrayOfALTOOLEFRC)SetItem(val []ALTOOLEFRC){
	t.Item = val
}


type  ALGTIDSTAT struct {
Tid ALGLOBTID `xml:"Tid"`
TypDefStatus string `xml:"TypDefStatus"`
}

func (t ALGTIDSTAT)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALGTIDSTAT)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALGTIDSTAT)GetTypDefStatus()string{
	return t.TypDefStatus
}

func (t *ALGTIDSTAT)SetTypDefStatus(val string){
	t.TypDefStatus = val
}


type  ArrayOfALGTIDSTAT struct {
Item []ALGTIDSTAT `xml:"item"`
}

func (t ArrayOfALGTIDSTAT)GetItem()[]ALGTIDSTAT{
	return t.Item
}

func (t *ArrayOfALGTIDSTAT)SetItem(val []ALGTIDSTAT){
	t.Item = val
}


type  ALMTCUSTWR struct {
Tid ALGLOBTID `xml:"Tid"`
GeneralCust ALMTCUST `xml:"GeneralCust"`
}

func (t ALMTCUSTWR)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMTCUSTWR)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMTCUSTWR)GetGeneralCust()ALMTCUST{
	return t.GeneralCust
}

func (t *ALMTCUSTWR)SetGeneralCust(val ALMTCUST){
	t.GeneralCust = val
}


type  ArrayOfALMTCUSTWR struct {
Item []ALMTCUSTWR `xml:"item"`
}

func (t ArrayOfALMTCUSTWR)GetItem()[]ALMTCUSTWR{
	return t.Item
}

func (t *ArrayOfALMTCUSTWR)SetItem(val []ALMTCUSTWR){
	t.Item = val
}


type  ALTIDUSER struct {
Tid ALGLOBTID `xml:"Tid"`
Aluser string `xml:"Aluser"`
}

func (t ALTIDUSER)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTIDUSER)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTIDUSER)GetAluser()string{
	return t.Aluser
}

func (t *ALTIDUSER)SetAluser(val string){
	t.Aluser = val
}


type  ArrayOfALTIDUSER struct {
Item []ALTIDUSER `xml:"item"`
}

func (t ArrayOfALTIDUSER)GetItem()[]ALTIDUSER{
	return t.Item
}

func (t *ArrayOfALTIDUSER)SetItem(val []ALTIDUSER){
	t.Item = val
}


type  ALPERFCUSW struct {
Tid ALGLOBTID `xml:"Tid"`
RelevantValueTyp string `xml:"RelevantValueTyp"`
ThresholdDirection string `xml:"ThresholdDirection"`
ThresholdG2Y string `xml:"ThresholdG2Y"`
ThresholdY2R string `xml:"ThresholdY2R"`
ThresholdY2G string `xml:"ThresholdY2G"`
ThresholdR2Y string `xml:"ThresholdR2Y"`
MsgClass string `xml:"MsgClass"`
MsgId string `xml:"MsgId"`
AttrGroupName string `xml:"AttrGroupName"`
NewCustStatus string `xml:"NewCustStatus"`
Unit string `xml:"Unit"`
Decimals string `xml:"Decimals"`
}

func (t ALPERFCUSW)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALPERFCUSW)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALPERFCUSW)GetRelevantValueTyp()string{
	return t.RelevantValueTyp
}

func (t *ALPERFCUSW)SetRelevantValueTyp(val string){
	t.RelevantValueTyp = val
}

func (t ALPERFCUSW)GetThresholdDirection()string{
	return t.ThresholdDirection
}

func (t *ALPERFCUSW)SetThresholdDirection(val string){
	t.ThresholdDirection = val
}

func (t ALPERFCUSW)GetThresholdG2Y()string{
	return t.ThresholdG2Y
}

func (t *ALPERFCUSW)SetThresholdG2Y(val string){
	t.ThresholdG2Y = val
}

func (t ALPERFCUSW)GetThresholdY2R()string{
	return t.ThresholdY2R
}

func (t *ALPERFCUSW)SetThresholdY2R(val string){
	t.ThresholdY2R = val
}

func (t ALPERFCUSW)GetThresholdY2G()string{
	return t.ThresholdY2G
}

func (t *ALPERFCUSW)SetThresholdY2G(val string){
	t.ThresholdY2G = val
}

func (t ALPERFCUSW)GetThresholdR2Y()string{
	return t.ThresholdR2Y
}

func (t *ALPERFCUSW)SetThresholdR2Y(val string){
	t.ThresholdR2Y = val
}

func (t ALPERFCUSW)GetMsgClass()string{
	return t.MsgClass
}

func (t *ALPERFCUSW)SetMsgClass(val string){
	t.MsgClass = val
}

func (t ALPERFCUSW)GetMsgId()string{
	return t.MsgId
}

func (t *ALPERFCUSW)SetMsgId(val string){
	t.MsgId = val
}

func (t ALPERFCUSW)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALPERFCUSW)SetAttrGroupName(val string){
	t.AttrGroupName = val
}

func (t ALPERFCUSW)GetNewCustStatus()string{
	return t.NewCustStatus
}

func (t *ALPERFCUSW)SetNewCustStatus(val string){
	t.NewCustStatus = val
}

func (t ALPERFCUSW)GetUnit()string{
	return t.Unit
}

func (t *ALPERFCUSW)SetUnit(val string){
	t.Unit = val
}

func (t ALPERFCUSW)GetDecimals()string{
	return t.Decimals
}

func (t *ALPERFCUSW)SetDecimals(val string){
	t.Decimals = val
}


type  ArrayOfALPERFCUSW struct {
Item []ALPERFCUSW `xml:"item"`
}

func (t ArrayOfALPERFCUSW)GetItem()[]ALPERFCUSW{
	return t.Item
}

func (t *ArrayOfALPERFCUSW)SetItem(val []ALPERFCUSW){
	t.Item = val
}


type  ALMSCCUSW struct {
Tid ALGLOBTID `xml:"Tid"`
CustStatus string `xml:"CustStatus"`
RaiseValue string `xml:"RaiseValue"`
RaiseSeverity string `xml:"RaiseSeverity"`
MscValMode string `xml:"MscValMode"`
MscValModeTimeSpan string `xml:"MscValModeTimeSpan"`
MaxAlertsPerID string `xml:"MaxAlertsPerID"`
KeepLinesTyp string `xml:"KeepLinesTyp"`
KeepLinesMax string `xml:"KeepLinesMax"`
AttrGroupName string `xml:"AttrGroupName"`
}

func (t ALMSCCUSW)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMSCCUSW)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMSCCUSW)GetCustStatus()string{
	return t.CustStatus
}

func (t *ALMSCCUSW)SetCustStatus(val string){
	t.CustStatus = val
}

func (t ALMSCCUSW)GetRaiseValue()string{
	return t.RaiseValue
}

func (t *ALMSCCUSW)SetRaiseValue(val string){
	t.RaiseValue = val
}

func (t ALMSCCUSW)GetRaiseSeverity()string{
	return t.RaiseSeverity
}

func (t *ALMSCCUSW)SetRaiseSeverity(val string){
	t.RaiseSeverity = val
}

func (t ALMSCCUSW)GetMscValMode()string{
	return t.MscValMode
}

func (t *ALMSCCUSW)SetMscValMode(val string){
	t.MscValMode = val
}

func (t ALMSCCUSW)GetMscValModeTimeSpan()string{
	return t.MscValModeTimeSpan
}

func (t *ALMSCCUSW)SetMscValModeTimeSpan(val string){
	t.MscValModeTimeSpan = val
}

func (t ALMSCCUSW)GetMaxAlertsPerID()string{
	return t.MaxAlertsPerID
}

func (t *ALMSCCUSW)SetMaxAlertsPerID(val string){
	t.MaxAlertsPerID = val
}

func (t ALMSCCUSW)GetKeepLinesTyp()string{
	return t.KeepLinesTyp
}

func (t *ALMSCCUSW)SetKeepLinesTyp(val string){
	t.KeepLinesTyp = val
}

func (t ALMSCCUSW)GetKeepLinesMax()string{
	return t.KeepLinesMax
}

func (t *ALMSCCUSW)SetKeepLinesMax(val string){
	t.KeepLinesMax = val
}

func (t ALMSCCUSW)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALMSCCUSW)SetAttrGroupName(val string){
	t.AttrGroupName = val
}


type  ArrayOfALMSCCUSW struct {
Item []ALMSCCUSW `xml:"item"`
}

func (t ArrayOfALMSCCUSW)GetItem()[]ALMSCCUSW{
	return t.Item
}

func (t *ArrayOfALMSCCUSW)SetItem(val []ALMSCCUSW){
	t.Item = val
}


type  ALMSCTIDFL struct {
Tid ALGLOBTID `xml:"Tid"`
Linenumber string `xml:"Linenumber"`
FromClass string `xml:"FromClass"`
FromID string `xml:"FromID"`
ToClass string `xml:"ToClass"`
ToID string `xml:"ToID"`
MsgValue string `xml:"MsgValue"`
Severity string `xml:"Severity"`
}

func (t ALMSCTIDFL)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMSCTIDFL)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMSCTIDFL)GetLinenumber()string{
	return t.Linenumber
}

func (t *ALMSCTIDFL)SetLinenumber(val string){
	t.Linenumber = val
}

func (t ALMSCTIDFL)GetFromClass()string{
	return t.FromClass
}

func (t *ALMSCTIDFL)SetFromClass(val string){
	t.FromClass = val
}

func (t ALMSCTIDFL)GetFromID()string{
	return t.FromID
}

func (t *ALMSCTIDFL)SetFromID(val string){
	t.FromID = val
}

func (t ALMSCTIDFL)GetToClass()string{
	return t.ToClass
}

func (t *ALMSCTIDFL)SetToClass(val string){
	t.ToClass = val
}

func (t ALMSCTIDFL)GetToID()string{
	return t.ToID
}

func (t *ALMSCTIDFL)SetToID(val string){
	t.ToID = val
}

func (t ALMSCTIDFL)GetMsgValue()string{
	return t.MsgValue
}

func (t *ALMSCTIDFL)SetMsgValue(val string){
	t.MsgValue = val
}

func (t ALMSCTIDFL)GetSeverity()string{
	return t.Severity
}

func (t *ALMSCTIDFL)SetSeverity(val string){
	t.Severity = val
}


type  ArrayOfALMSCTIDFL struct {
Item []ALMSCTIDFL `xml:"item"`
}

func (t ArrayOfALMSCTIDFL)GetItem()[]ALMSCTIDFL{
	return t.Item
}

func (t *ArrayOfALMSCTIDFL)SetItem(val []ALMSCTIDFL){
	t.Item = val
}


type  ALSMSGCUSW struct {
Tid ALGLOBTID `xml:"Tid"`
CustStatus string `xml:"CustStatus"`
AlertMode string `xml:"AlertMode"`
AlertShift string `xml:"AlertShift"`
AttrGroupName string `xml:"AttrGroupName"`
}

func (t ALSMSGCUSW)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALSMSGCUSW)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALSMSGCUSW)GetCustStatus()string{
	return t.CustStatus
}

func (t *ALSMSGCUSW)SetCustStatus(val string){
	t.CustStatus = val
}

func (t ALSMSGCUSW)GetAlertMode()string{
	return t.AlertMode
}

func (t *ALSMSGCUSW)SetAlertMode(val string){
	t.AlertMode = val
}

func (t ALSMSGCUSW)GetAlertShift()string{
	return t.AlertShift
}

func (t *ALSMSGCUSW)SetAlertShift(val string){
	t.AlertShift = val
}

func (t ALSMSGCUSW)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALSMSGCUSW)SetAttrGroupName(val string){
	t.AttrGroupName = val
}


type  ArrayOfALSMSGCUSW struct {
Item []ALSMSGCUSW `xml:"item"`
}

func (t ArrayOfALSMSGCUSW)GetItem()[]ALSMSGCUSW{
	return t.Item
}

func (t *ArrayOfALSMSGCUSW)SetItem(val []ALSMSGCUSW){
	t.Item = val
}


type  ALGTIDTLRC struct {
Tid ALGLOBTID `xml:"Tid"`
Whichtool string `xml:"Whichtool"`
Rc string `xml:"Rc"`
}

func (t ALGTIDTLRC)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALGTIDTLRC)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALGTIDTLRC)GetWhichtool()string{
	return t.Whichtool
}

func (t *ALGTIDTLRC)SetWhichtool(val string){
	t.Whichtool = val
}

func (t ALGTIDTLRC)GetRc()string{
	return t.Rc
}

func (t *ALGTIDTLRC)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALGTIDTLRC struct {
Item []ALGTIDTLRC `xml:"item"`
}

func (t ArrayOfALGTIDTLRC)GetItem()[]ALGTIDTLRC{
	return t.Item
}

func (t *ArrayOfALGTIDTLRC)SetItem(val []ALGTIDTLRC){
	t.Item = val
}


type  ALTOOLASSG struct {
Tid ALGLOBTID `xml:"Tid"`
Whichtool string `xml:"Whichtool"`
ToolName string `xml:"ToolName"`
ToolDispatcher string `xml:"ToolDispatcher"`
ToolStatus string `xml:"ToolStatus"`
Useclass string `xml:"Useclass"`
}

func (t ALTOOLASSG)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTOOLASSG)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTOOLASSG)GetWhichtool()string{
	return t.Whichtool
}

func (t *ALTOOLASSG)SetWhichtool(val string){
	t.Whichtool = val
}

func (t ALTOOLASSG)GetToolName()string{
	return t.ToolName
}

func (t *ALTOOLASSG)SetToolName(val string){
	t.ToolName = val
}

func (t ALTOOLASSG)GetToolDispatcher()string{
	return t.ToolDispatcher
}

func (t *ALTOOLASSG)SetToolDispatcher(val string){
	t.ToolDispatcher = val
}

func (t ALTOOLASSG)GetToolStatus()string{
	return t.ToolStatus
}

func (t *ALTOOLASSG)SetToolStatus(val string){
	t.ToolStatus = val
}

func (t ALTOOLASSG)GetUseclass()string{
	return t.Useclass
}

func (t *ALTOOLASSG)SetUseclass(val string){
	t.Useclass = val
}


type  ArrayOfALTOOLASSG struct {
Item []ALTOOLASSG `xml:"item"`
}

func (t ArrayOfALTOOLASSG)GetItem()[]ALTOOLASSG{
	return t.Item
}

func (t *ArrayOfALTOOLASSG)SetItem(val []ALTOOLASSG){
	t.Item = val
}


type  ALTOOLCHST struct {
Tid ALGLOBTID `xml:"Tid"`
Whichtool string `xml:"Whichtool"`
ToolStatus string `xml:"ToolStatus"`
}

func (t ALTOOLCHST)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTOOLCHST)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTOOLCHST)GetWhichtool()string{
	return t.Whichtool
}

func (t *ALTOOLCHST)SetWhichtool(val string){
	t.Whichtool = val
}

func (t ALTOOLCHST)GetToolStatus()string{
	return t.ToolStatus
}

func (t *ALTOOLCHST)SetToolStatus(val string){
	t.ToolStatus = val
}


type  ArrayOfALTOOLCHST struct {
Item []ALTOOLCHST `xml:"item"`
}

func (t ArrayOfALTOOLCHST)GetItem()[]ALTOOLCHST{
	return t.Item
}

func (t *ArrayOfALTOOLCHST)SetItem(val []ALTOOLCHST){
	t.Item = val
}


type  ALAIDALSTA struct {
Aid ALGLOBAID `xml:"Aid"`
Status string `xml:"Status"`
Aluser string `xml:"Aluser"`
}

func (t ALAIDALSTA)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALAIDALSTA)SetAid(val ALGLOBAID){
	t.Aid = val
}

func (t ALAIDALSTA)GetStatus()string{
	return t.Status
}

func (t *ALAIDALSTA)SetStatus(val string){
	t.Status = val
}

func (t ALAIDALSTA)GetAluser()string{
	return t.Aluser
}

func (t *ALAIDALSTA)SetAluser(val string){
	t.Aluser = val
}


type  ArrayOfALAIDALSTA struct {
Item []ALAIDALSTA `xml:"item"`
}

func (t ArrayOfALAIDALSTA)GetItem()[]ALAIDALSTA{
	return t.Item
}

func (t *ArrayOfALAIDALSTA)SetItem(val []ALAIDALSTA){
	t.Item = val
}


type  ALGAIDRC struct {
Aid ALGLOBAID `xml:"Aid"`
Rc string `xml:"Rc"`
}

func (t ALGAIDRC)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALGAIDRC)SetAid(val ALGLOBAID){
	t.Aid = val
}

func (t ALGAIDRC)GetRc()string{
	return t.Rc
}

func (t *ALGAIDRC)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALGAIDRC struct {
Item []ALGAIDRC `xml:"item"`
}

func (t ArrayOfALGAIDRC)GetItem()[]ALGAIDRC{
	return t.Item
}

func (t *ArrayOfALGAIDRC)SetItem(val []ALGAIDRC){
	t.Item = val
}


type  ALREQUESTCUC struct {
Guid string `xml:"Guid"`
SysOrGrp string `xml:"SysOrGrp"`
Class1 string `xml:"Class1"`
Class2 string `xml:"Class2"`
Client string `xml:"Client"`
TargetFilter string `xml:"TargetFilter"`
NoWildcards string `xml:"NoWildcards"`
RequestOption string `xml:"RequestOption"`
}

func (t ALREQUESTCUC)GetGuid()string{
	return t.Guid
}

func (t *ALREQUESTCUC)SetGuid(val string){
	t.Guid = val
}

func (t ALREQUESTCUC)GetSysOrGrp()string{
	return t.SysOrGrp
}

func (t *ALREQUESTCUC)SetSysOrGrp(val string){
	t.SysOrGrp = val
}

func (t ALREQUESTCUC)GetClass1()string{
	return t.Class1
}

func (t *ALREQUESTCUC)SetClass1(val string){
	t.Class1 = val
}

func (t ALREQUESTCUC)GetClass2()string{
	return t.Class2
}

func (t *ALREQUESTCUC)SetClass2(val string){
	t.Class2 = val
}

func (t ALREQUESTCUC)GetClient()string{
	return t.Client
}

func (t *ALREQUESTCUC)SetClient(val string){
	t.Client = val
}

func (t ALREQUESTCUC)GetTargetFilter()string{
	return t.TargetFilter
}

func (t *ALREQUESTCUC)SetTargetFilter(val string){
	t.TargetFilter = val
}

func (t ALREQUESTCUC)GetNoWildcards()string{
	return t.NoWildcards
}

func (t *ALREQUESTCUC)SetNoWildcards(val string){
	t.NoWildcards = val
}

func (t ALREQUESTCUC)GetRequestOption()string{
	return t.RequestOption
}

func (t *ALREQUESTCUC)SetRequestOption(val string){
	t.RequestOption = val
}


type  ArrayOfALREQUESTCUC struct {
Item []ALREQUESTCUC `xml:"item"`
}

func (t ArrayOfALREQUESTCUC)GetItem()[]ALREQUESTCUC{
	return t.Item
}

func (t *ArrayOfALREQUESTCUC)SetItem(val []ALREQUESTCUC){
	t.Item = val
}


type  ALRESULTCUC struct {
Guid string `xml:"Guid"`
Rc string `xml:"Rc"`
PTid ALGLOBTID `xml:"PTid"`
MTid ALGLOBTID `xml:"MTid"`
}

func (t ALRESULTCUC)GetGuid()string{
	return t.Guid
}

func (t *ALRESULTCUC)SetGuid(val string){
	t.Guid = val
}

func (t ALRESULTCUC)GetRc()string{
	return t.Rc
}

func (t *ALRESULTCUC)SetRc(val string){
	t.Rc = val
}

func (t ALRESULTCUC)GetPTid()ALGLOBTID{
	return t.PTid
}

func (t *ALRESULTCUC)SetPTid(val ALGLOBTID){
	t.PTid = val
}

func (t ALRESULTCUC)GetMTid()ALGLOBTID{
	return t.MTid
}

func (t *ALRESULTCUC)SetMTid(val ALGLOBTID){
	t.MTid = val
}


type  ArrayOfALRESULTCUC struct {
Item []ALRESULTCUC `xml:"item"`
}

func (t ArrayOfALRESULTCUC)GetItem()[]ALRESULTCUC{
	return t.Item
}

func (t *ArrayOfALRESULTCUC)SetItem(val []ALRESULTCUC){
	t.Item = val
}


type  ALRSLTCUCL struct {
Guid string `xml:"Guid"`
Rc string `xml:"Rc"`
PTid ALGLOBTID `xml:"PTid"`
PMtename string `xml:"PMtename"`
MTid ALGLOBTID `xml:"MTid"`
Mtename string `xml:"Mtename"`
}

func (t ALRSLTCUCL)GetGuid()string{
	return t.Guid
}

func (t *ALRSLTCUCL)SetGuid(val string){
	t.Guid = val
}

func (t ALRSLTCUCL)GetRc()string{
	return t.Rc
}

func (t *ALRSLTCUCL)SetRc(val string){
	t.Rc = val
}

func (t ALRSLTCUCL)GetPTid()ALGLOBTID{
	return t.PTid
}

func (t *ALRSLTCUCL)SetPTid(val ALGLOBTID){
	t.PTid = val
}

func (t ALRSLTCUCL)GetPMtename()string{
	return t.PMtename
}

func (t *ALRSLTCUCL)SetPMtename(val string){
	t.PMtename = val
}

func (t ALRSLTCUCL)GetMTid()ALGLOBTID{
	return t.MTid
}

func (t *ALRSLTCUCL)SetMTid(val ALGLOBTID){
	t.MTid = val
}

func (t ALRSLTCUCL)GetMtename()string{
	return t.Mtename
}

func (t *ALRSLTCUCL)SetMtename(val string){
	t.Mtename = val
}


type  ArrayOfALRSLTCUCL struct {
Item []ALRSLTCUCL `xml:"item"`
}

func (t ArrayOfALRSLTCUCL)GetItem()[]ALRSLTCUCL{
	return t.Item
}

func (t *ArrayOfALRSLTCUCL)SetItem(val []ALRSLTCUCL){
	t.Item = val
}


type  ALMONCTX2 struct {
Tid ALGLOBTID `xml:"Tid"`
Owner string `xml:"Owner"`
ContxtStatus string `xml:"ContxtStatus"`
ContxtType string `xml:"ContxtType"`
ABAPClient string `xml:"ABAPClient"`
MteClass string `xml:"MteClass"`
}

func (t ALMONCTX2)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMONCTX2)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMONCTX2)GetOwner()string{
	return t.Owner
}

func (t *ALMONCTX2)SetOwner(val string){
	t.Owner = val
}

func (t ALMONCTX2)GetContxtStatus()string{
	return t.ContxtStatus
}

func (t *ALMONCTX2)SetContxtStatus(val string){
	t.ContxtStatus = val
}

func (t ALMONCTX2)GetContxtType()string{
	return t.ContxtType
}

func (t *ALMONCTX2)SetContxtType(val string){
	t.ContxtType = val
}

func (t ALMONCTX2)GetABAPClient()string{
	return t.ABAPClient
}

func (t *ALMONCTX2)SetABAPClient(val string){
	t.ABAPClient = val
}

func (t ALMONCTX2)GetMteClass()string{
	return t.MteClass
}

func (t *ALMONCTX2)SetMteClass(val string){
	t.MteClass = val
}


type  ArrayOfALMONCTX2 struct {
Item []ALMONCTX2 `xml:"item"`
}

func (t ArrayOfALMONCTX2)GetItem()[]ALMONCTX2{
	return t.Item
}

func (t *ArrayOfALMONCTX2)SetItem(val []ALMONCTX2){
	t.Item = val
}


type  ALMONCTX3 struct {
Tid ALGLOBTID `xml:"Tid"`
Owner string `xml:"Owner"`
ContxtStatus string `xml:"ContxtStatus"`
ContxtType string `xml:"ContxtType"`
ABAPClient string `xml:"ABAPClient"`
MteClass string `xml:"MteClass"`
SystemId string `xml:"SystemId"`
}

func (t ALMONCTX3)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMONCTX3)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMONCTX3)GetOwner()string{
	return t.Owner
}

func (t *ALMONCTX3)SetOwner(val string){
	t.Owner = val
}

func (t ALMONCTX3)GetContxtStatus()string{
	return t.ContxtStatus
}

func (t *ALMONCTX3)SetContxtStatus(val string){
	t.ContxtStatus = val
}

func (t ALMONCTX3)GetContxtType()string{
	return t.ContxtType
}

func (t *ALMONCTX3)SetContxtType(val string){
	t.ContxtType = val
}

func (t ALMONCTX3)GetABAPClient()string{
	return t.ABAPClient
}

func (t *ALMONCTX3)SetABAPClient(val string){
	t.ABAPClient = val
}

func (t ALMONCTX3)GetMteClass()string{
	return t.MteClass
}

func (t *ALMONCTX3)SetMteClass(val string){
	t.MteClass = val
}

func (t ALMONCTX3)GetSystemId()string{
	return t.SystemId
}

func (t *ALMONCTX3)SetSystemId(val string){
	t.SystemId = val
}


type  ArrayOfALMONCTX3 struct {
Item []ALMONCTX3 `xml:"item"`
}

func (t ArrayOfALMONCTX3)GetItem()[]ALMONCTX3{
	return t.Item
}

func (t *ArrayOfALMONCTX3)SetItem(val []ALMONCTX3){
	t.Item = val
}


type  ALCUSTGRP struct {
AttrGroupName string `xml:"AttrGroupName"`
}

func (t ALCUSTGRP)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALCUSTGRP)SetAttrGroupName(val string){
	t.AttrGroupName = val
}


type  ArrayOfALCUSTGRP struct {
Item []ALCUSTGRP `xml:"item"`
}

func (t ArrayOfALCUSTGRP)GetItem()[]ALCUSTGRP{
	return t.Item
}

func (t *ArrayOfALCUSTGRP)SetItem(val []ALCUSTGRP){
	t.Item = val
}


type  ALMONSEG struct {
SegmentStatus string `xml:"SegmentStatus"`
SegmentType string `xml:"SegmentType"`
Hostname string `xml:"Hostname"`
Version string `xml:"Version"`
SegmentStartDate string `xml:"SegmentStartDate"`
SegmentStartTime string `xml:"SegmentStartTime"`
SegmentName string `xml:"SegmentName"`
SegmentSysid string `xml:"SegmentSysid"`
SegmentOwner string `xml:"SegmentOwner"`
SegmentLongtxt string `xml:"SegmentLongtxt"`
Tool ALTOOL `xml:"Tool"`
Tid ALGLOBTID `xml:"Tid"`
Aid ALGLOBAID `xml:"Aid"`
Whichtool string `xml:"Whichtool"`
SyAgentStatusInfo2 string `xml:"SyAgentStatusInfo2"`
SyStartDate string `xml:"SyStartDate"`
SyStartTime string `xml:"SyStartTime"`
SyStopDate string `xml:"SyStopDate"`
SyStopTime string `xml:"SyStopTime"`
}

func (t ALMONSEG)GetSegmentStatus()string{
	return t.SegmentStatus
}

func (t *ALMONSEG)SetSegmentStatus(val string){
	t.SegmentStatus = val
}

func (t ALMONSEG)GetSegmentType()string{
	return t.SegmentType
}

func (t *ALMONSEG)SetSegmentType(val string){
	t.SegmentType = val
}

func (t ALMONSEG)GetHostname()string{
	return t.Hostname
}

func (t *ALMONSEG)SetHostname(val string){
	t.Hostname = val
}

func (t ALMONSEG)GetVersion()string{
	return t.Version
}

func (t *ALMONSEG)SetVersion(val string){
	t.Version = val
}

func (t ALMONSEG)GetSegmentStartDate()string{
	return t.SegmentStartDate
}

func (t *ALMONSEG)SetSegmentStartDate(val string){
	t.SegmentStartDate = val
}

func (t ALMONSEG)GetSegmentStartTime()string{
	return t.SegmentStartTime
}

func (t *ALMONSEG)SetSegmentStartTime(val string){
	t.SegmentStartTime = val
}

func (t ALMONSEG)GetSegmentName()string{
	return t.SegmentName
}

func (t *ALMONSEG)SetSegmentName(val string){
	t.SegmentName = val
}

func (t ALMONSEG)GetSegmentSysid()string{
	return t.SegmentSysid
}

func (t *ALMONSEG)SetSegmentSysid(val string){
	t.SegmentSysid = val
}

func (t ALMONSEG)GetSegmentOwner()string{
	return t.SegmentOwner
}

func (t *ALMONSEG)SetSegmentOwner(val string){
	t.SegmentOwner = val
}

func (t ALMONSEG)GetSegmentLongtxt()string{
	return t.SegmentLongtxt
}

func (t *ALMONSEG)SetSegmentLongtxt(val string){
	t.SegmentLongtxt = val
}

func (t ALMONSEG)GetTool()ALTOOL{
	return t.Tool
}

func (t *ALMONSEG)SetTool(val ALTOOL){
	t.Tool = val
}

func (t ALMONSEG)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMONSEG)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMONSEG)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALMONSEG)SetAid(val ALGLOBAID){
	t.Aid = val
}

func (t ALMONSEG)GetWhichtool()string{
	return t.Whichtool
}

func (t *ALMONSEG)SetWhichtool(val string){
	t.Whichtool = val
}

func (t ALMONSEG)GetSyAgentStatusInfo2()string{
	return t.SyAgentStatusInfo2
}

func (t *ALMONSEG)SetSyAgentStatusInfo2(val string){
	t.SyAgentStatusInfo2 = val
}

func (t ALMONSEG)GetSyStartDate()string{
	return t.SyStartDate
}

func (t *ALMONSEG)SetSyStartDate(val string){
	t.SyStartDate = val
}

func (t ALMONSEG)GetSyStartTime()string{
	return t.SyStartTime
}

func (t *ALMONSEG)SetSyStartTime(val string){
	t.SyStartTime = val
}

func (t ALMONSEG)GetSyStopDate()string{
	return t.SyStopDate
}

func (t *ALMONSEG)SetSyStopDate(val string){
	t.SyStopDate = val
}

func (t ALMONSEG)GetSyStopTime()string{
	return t.SyStopTime
}

func (t *ALMONSEG)SetSyStopTime(val string){
	t.SyStopTime = val
}


type  ArrayOfALMONSEG struct {
Item []ALMONSEG `xml:"item"`
}

func (t ArrayOfALMONSEG)GetItem()[]ALMONSEG{
	return t.Item
}

func (t *ArrayOfALMONSEG)SetItem(val []ALMONSEG){
	t.Item = val
}


type  ALTIDMAXLV struct {
Tid ALGLOBTID `xml:"Tid"`
Maxlevel string `xml:"Maxlevel"`
VisibilityLevel string `xml:"VisibilityLevel"`
}

func (t ALTIDMAXLV)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTIDMAXLV)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTIDMAXLV)GetMaxlevel()string{
	return t.Maxlevel
}

func (t *ALTIDMAXLV)SetMaxlevel(val string){
	t.Maxlevel = val
}

func (t ALTIDMAXLV)GetVisibilityLevel()string{
	return t.VisibilityLevel
}

func (t *ALTIDMAXLV)SetVisibilityLevel(val string){
	t.VisibilityLevel = val
}


type  ArrayOfALTIDMAXLV struct {
Item []ALTIDMAXLV `xml:"item"`
}

func (t ArrayOfALTIDMAXLV)GetItem()[]ALTIDMAXLV{
	return t.Item
}

func (t *ArrayOfALTIDMAXLV)SetItem(val []ALTIDMAXLV){
	t.Item = val
}


type  ALMTTREE struct {
Tid ALGLOBTID `xml:"Tid"`
TidNumber string `xml:"TidNumber"`
IndexInTree string `xml:"IndexInTree"`
LevelInTree string `xml:"LevelInTree"`
ParentInTree string `xml:"ParentInTree"`
Objectname string `xml:"Objectname"`
ShortName string `xml:"ShortName"`
ClassName string `xml:"ClassName"`
DeliveryStatus string `xml:"DeliveryStatus"`
HighAlertValue string `xml:"HighAlertValue"`
HighAlertSeverity string `xml:"HighAlertSeverity"`
Aid ALGLOBAID `xml:"Aid"`
LastValueDate string `xml:"LastValueDate"`
LastValueTime string `xml:"LastValueTime"`
CurrentValue string `xml:"CurrentValue"`
CurrentSeverity string `xml:"CurrentSeverity"`
CurrentValueAid ALGLOBAID `xml:"CurrentValueAid"`
ActiveAlerts string `xml:"ActiveAlerts"`
AlertsSinceStartup string `xml:"AlertsSinceStartup"`
VisibilityLevel string `xml:"VisibilityLevel"`
TypDefStatus string `xml:"TypDefStatus"`
Rc string `xml:"Rc"`
}

func (t ALMTTREE)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMTTREE)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMTTREE)GetTidNumber()string{
	return t.TidNumber
}

func (t *ALMTTREE)SetTidNumber(val string){
	t.TidNumber = val
}

func (t ALMTTREE)GetIndexInTree()string{
	return t.IndexInTree
}

func (t *ALMTTREE)SetIndexInTree(val string){
	t.IndexInTree = val
}

func (t ALMTTREE)GetLevelInTree()string{
	return t.LevelInTree
}

func (t *ALMTTREE)SetLevelInTree(val string){
	t.LevelInTree = val
}

func (t ALMTTREE)GetParentInTree()string{
	return t.ParentInTree
}

func (t *ALMTTREE)SetParentInTree(val string){
	t.ParentInTree = val
}

func (t ALMTTREE)GetObjectname()string{
	return t.Objectname
}

func (t *ALMTTREE)SetObjectname(val string){
	t.Objectname = val
}

func (t ALMTTREE)GetShortName()string{
	return t.ShortName
}

func (t *ALMTTREE)SetShortName(val string){
	t.ShortName = val
}

func (t ALMTTREE)GetClassName()string{
	return t.ClassName
}

func (t *ALMTTREE)SetClassName(val string){
	t.ClassName = val
}

func (t ALMTTREE)GetDeliveryStatus()string{
	return t.DeliveryStatus
}

func (t *ALMTTREE)SetDeliveryStatus(val string){
	t.DeliveryStatus = val
}

func (t ALMTTREE)GetHighAlertValue()string{
	return t.HighAlertValue
}

func (t *ALMTTREE)SetHighAlertValue(val string){
	t.HighAlertValue = val
}

func (t ALMTTREE)GetHighAlertSeverity()string{
	return t.HighAlertSeverity
}

func (t *ALMTTREE)SetHighAlertSeverity(val string){
	t.HighAlertSeverity = val
}

func (t ALMTTREE)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALMTTREE)SetAid(val ALGLOBAID){
	t.Aid = val
}

func (t ALMTTREE)GetLastValueDate()string{
	return t.LastValueDate
}

func (t *ALMTTREE)SetLastValueDate(val string){
	t.LastValueDate = val
}

func (t ALMTTREE)GetLastValueTime()string{
	return t.LastValueTime
}

func (t *ALMTTREE)SetLastValueTime(val string){
	t.LastValueTime = val
}

func (t ALMTTREE)GetCurrentValue()string{
	return t.CurrentValue
}

func (t *ALMTTREE)SetCurrentValue(val string){
	t.CurrentValue = val
}

func (t ALMTTREE)GetCurrentSeverity()string{
	return t.CurrentSeverity
}

func (t *ALMTTREE)SetCurrentSeverity(val string){
	t.CurrentSeverity = val
}

func (t ALMTTREE)GetCurrentValueAid()ALGLOBAID{
	return t.CurrentValueAid
}

func (t *ALMTTREE)SetCurrentValueAid(val ALGLOBAID){
	t.CurrentValueAid = val
}

func (t ALMTTREE)GetActiveAlerts()string{
	return t.ActiveAlerts
}

func (t *ALMTTREE)SetActiveAlerts(val string){
	t.ActiveAlerts = val
}

func (t ALMTTREE)GetAlertsSinceStartup()string{
	return t.AlertsSinceStartup
}

func (t *ALMTTREE)SetAlertsSinceStartup(val string){
	t.AlertsSinceStartup = val
}

func (t ALMTTREE)GetVisibilityLevel()string{
	return t.VisibilityLevel
}

func (t *ALMTTREE)SetVisibilityLevel(val string){
	t.VisibilityLevel = val
}

func (t ALMTTREE)GetTypDefStatus()string{
	return t.TypDefStatus
}

func (t *ALMTTREE)SetTypDefStatus(val string){
	t.TypDefStatus = val
}

func (t ALMTTREE)GetRc()string{
	return t.Rc
}

func (t *ALMTTREE)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALMTTREE struct {
Item []ALMTTREE `xml:"item"`
}

func (t ArrayOfALMTTREE)GetItem()[]ALMTTREE{
	return t.Item
}

func (t *ArrayOfALMTTREE)SetItem(val []ALMTTREE){
	t.Item = val
}


type  ALSMSGTYPE struct {
Tid ALGLOBTID `xml:"Tid"`
CustStatus string `xml:"CustStatus"`
AlertMode string `xml:"AlertMode"`
AlertShift string `xml:"AlertShift"`
AttrGroupName string `xml:"AttrGroupName"`
MsgValue string `xml:"MsgValue"`
MsgDate string `xml:"MsgDate"`
MsgTime string `xml:"MsgTime"`
Message string `xml:"Message"`
Rc string `xml:"Rc"`
}

func (t ALSMSGTYPE)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALSMSGTYPE)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALSMSGTYPE)GetCustStatus()string{
	return t.CustStatus
}

func (t *ALSMSGTYPE)SetCustStatus(val string){
	t.CustStatus = val
}

func (t ALSMSGTYPE)GetAlertMode()string{
	return t.AlertMode
}

func (t *ALSMSGTYPE)SetAlertMode(val string){
	t.AlertMode = val
}

func (t ALSMSGTYPE)GetAlertShift()string{
	return t.AlertShift
}

func (t *ALSMSGTYPE)SetAlertShift(val string){
	t.AlertShift = val
}

func (t ALSMSGTYPE)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALSMSGTYPE)SetAttrGroupName(val string){
	t.AttrGroupName = val
}

func (t ALSMSGTYPE)GetMsgValue()string{
	return t.MsgValue
}

func (t *ALSMSGTYPE)SetMsgValue(val string){
	t.MsgValue = val
}

func (t ALSMSGTYPE)GetMsgDate()string{
	return t.MsgDate
}

func (t *ALSMSGTYPE)SetMsgDate(val string){
	t.MsgDate = val
}

func (t ALSMSGTYPE)GetMsgTime()string{
	return t.MsgTime
}

func (t *ALSMSGTYPE)SetMsgTime(val string){
	t.MsgTime = val
}

func (t ALSMSGTYPE)GetMessage()string{
	return t.Message
}

func (t *ALSMSGTYPE)SetMessage(val string){
	t.Message = val
}

func (t ALSMSGTYPE)GetRc()string{
	return t.Rc
}

func (t *ALSMSGTYPE)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALSMSGTYPE struct {
Item []ALSMSGTYPE `xml:"item"`
}

func (t ArrayOfALSMSGTYPE)GetItem()[]ALSMSGTYPE{
	return t.Item
}

func (t *ArrayOfALSMSGTYPE)SetItem(val []ALSMSGTYPE){
	t.Item = val
}


type  ArrayOfALGLOBTID struct {
Item []ALGLOBTID `xml:"item"`
}

func (t ArrayOfALGLOBTID)GetItem()[]ALGLOBTID{
	return t.Item
}

func (t *ArrayOfALGLOBTID)SetItem(val []ALGLOBTID){
	t.Item = val
}


type  ALTOOLCTRL struct {
ToolDesc ALTOOL `xml:"ToolDesc"`
Tid ALGLOBTID `xml:"Tid"`
Aid ALGLOBAID `xml:"Aid"`
Whichtool string `xml:"Whichtool"`
}

func (t ALTOOLCTRL)GetToolDesc()ALTOOL{
	return t.ToolDesc
}

func (t *ALTOOLCTRL)SetToolDesc(val ALTOOL){
	t.ToolDesc = val
}

func (t ALTOOLCTRL)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALTOOLCTRL)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALTOOLCTRL)GetAid()ALGLOBAID{
	return t.Aid
}

func (t *ALTOOLCTRL)SetAid(val ALGLOBAID){
	t.Aid = val
}

func (t ALTOOLCTRL)GetWhichtool()string{
	return t.Whichtool
}

func (t *ALTOOLCTRL)SetWhichtool(val string){
	t.Whichtool = val
}


type  ArrayOfALTOOLCTRL struct {
Item []ALTOOLCTRL `xml:"item"`
}

func (t ArrayOfALTOOLCTRL)GetItem()[]ALTOOLCTRL{
	return t.Item
}

func (t *ArrayOfALTOOLCTRL)SetItem(val []ALTOOLCTRL){
	t.Item = val
}


type  ALTLDPNM struct {
ToolDispatcher string `xml:"ToolDispatcher"`
}

func (t ALTLDPNM)GetToolDispatcher()string{
	return t.ToolDispatcher
}

func (t *ALTLDPNM)SetToolDispatcher(val string){
	t.ToolDispatcher = val
}


type  ArrayOfALTLDPNM struct {
Item []ALTLDPNM `xml:"item"`
}

func (t ArrayOfALTLDPNM)GetItem()[]ALTLDPNM{
	return t.Item
}

func (t *ArrayOfALTLDPNM)SetItem(val []ALTLDPNM){
	t.Item = val
}


type  ALTOOLDPRC struct {
ToolDispatcher string `xml:"ToolDispatcher"`
Rc string `xml:"Rc"`
}

func (t ALTOOLDPRC)GetToolDispatcher()string{
	return t.ToolDispatcher
}

func (t *ALTOOLDPRC)SetToolDispatcher(val string){
	t.ToolDispatcher = val
}

func (t ALTOOLDPRC)GetRc()string{
	return t.Rc
}

func (t *ALTOOLDPRC)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALTOOLDPRC struct {
Item []ALTOOLDPRC `xml:"item"`
}

func (t ArrayOfALTOOLDPRC)GetItem()[]ALTOOLDPRC{
	return t.Item
}

func (t *ArrayOfALTOOLDPRC)SetItem(val []ALTOOLDPRC){
	t.Item = val
}


type  ALMSCRAWT struct {
Tid ALGLOBTID `xml:"Tid"`
CustStatus string `xml:"CustStatus"`
RaiseValue string `xml:"RaiseValue"`
RaiseSeverity string `xml:"RaiseSeverity"`
MscValMode string `xml:"MscValMode"`
MscValModeTimeSpan string `xml:"MscValModeTimeSpan"`
MaxAlertsPerID string `xml:"MaxAlertsPerID"`
KeepLinesTyp string `xml:"KeepLinesTyp"`
KeepLinesMax string `xml:"KeepLinesMax"`
AttrGroupName string `xml:"AttrGroupName"`
MscLineId string `xml:"MscLineId"`
MscDate string `xml:"MscDate"`
MscTime string `xml:"MscTime"`
ValueOrig string `xml:"ValueOrig"`
SeverityOrig string `xml:"SeverityOrig"`
ValueFilter string `xml:"ValueFilter"`
SeverityFilter string `xml:"SeverityFilter"`
ABAPClient string `xml:"ABAPClient"`
User string `xml:"User"`
Message SXMIMSGRAW `xml:"Message"`
Rc string `xml:"Rc"`
}

func (t ALMSCRAWT)GetTid()ALGLOBTID{
	return t.Tid
}

func (t *ALMSCRAWT)SetTid(val ALGLOBTID){
	t.Tid = val
}

func (t ALMSCRAWT)GetCustStatus()string{
	return t.CustStatus
}

func (t *ALMSCRAWT)SetCustStatus(val string){
	t.CustStatus = val
}

func (t ALMSCRAWT)GetRaiseValue()string{
	return t.RaiseValue
}

func (t *ALMSCRAWT)SetRaiseValue(val string){
	t.RaiseValue = val
}

func (t ALMSCRAWT)GetRaiseSeverity()string{
	return t.RaiseSeverity
}

func (t *ALMSCRAWT)SetRaiseSeverity(val string){
	t.RaiseSeverity = val
}

func (t ALMSCRAWT)GetMscValMode()string{
	return t.MscValMode
}

func (t *ALMSCRAWT)SetMscValMode(val string){
	t.MscValMode = val
}

func (t ALMSCRAWT)GetMscValModeTimeSpan()string{
	return t.MscValModeTimeSpan
}

func (t *ALMSCRAWT)SetMscValModeTimeSpan(val string){
	t.MscValModeTimeSpan = val
}

func (t ALMSCRAWT)GetMaxAlertsPerID()string{
	return t.MaxAlertsPerID
}

func (t *ALMSCRAWT)SetMaxAlertsPerID(val string){
	t.MaxAlertsPerID = val
}

func (t ALMSCRAWT)GetKeepLinesTyp()string{
	return t.KeepLinesTyp
}

func (t *ALMSCRAWT)SetKeepLinesTyp(val string){
	t.KeepLinesTyp = val
}

func (t ALMSCRAWT)GetKeepLinesMax()string{
	return t.KeepLinesMax
}

func (t *ALMSCRAWT)SetKeepLinesMax(val string){
	t.KeepLinesMax = val
}

func (t ALMSCRAWT)GetAttrGroupName()string{
	return t.AttrGroupName
}

func (t *ALMSCRAWT)SetAttrGroupName(val string){
	t.AttrGroupName = val
}

func (t ALMSCRAWT)GetMscLineId()string{
	return t.MscLineId
}

func (t *ALMSCRAWT)SetMscLineId(val string){
	t.MscLineId = val
}

func (t ALMSCRAWT)GetMscDate()string{
	return t.MscDate
}

func (t *ALMSCRAWT)SetMscDate(val string){
	t.MscDate = val
}

func (t ALMSCRAWT)GetMscTime()string{
	return t.MscTime
}

func (t *ALMSCRAWT)SetMscTime(val string){
	t.MscTime = val
}

func (t ALMSCRAWT)GetValueOrig()string{
	return t.ValueOrig
}

func (t *ALMSCRAWT)SetValueOrig(val string){
	t.ValueOrig = val
}

func (t ALMSCRAWT)GetSeverityOrig()string{
	return t.SeverityOrig
}

func (t *ALMSCRAWT)SetSeverityOrig(val string){
	t.SeverityOrig = val
}

func (t ALMSCRAWT)GetValueFilter()string{
	return t.ValueFilter
}

func (t *ALMSCRAWT)SetValueFilter(val string){
	t.ValueFilter = val
}

func (t ALMSCRAWT)GetSeverityFilter()string{
	return t.SeverityFilter
}

func (t *ALMSCRAWT)SetSeverityFilter(val string){
	t.SeverityFilter = val
}

func (t ALMSCRAWT)GetABAPClient()string{
	return t.ABAPClient
}

func (t *ALMSCRAWT)SetABAPClient(val string){
	t.ABAPClient = val
}

func (t ALMSCRAWT)GetUser()string{
	return t.User
}

func (t *ALMSCRAWT)SetUser(val string){
	t.User = val
}

func (t ALMSCRAWT)GetMessage()SXMIMSGRAW{
	return t.Message
}

func (t *ALMSCRAWT)SetMessage(val SXMIMSGRAW){
	t.Message = val
}

func (t ALMSCRAWT)GetRc()string{
	return t.Rc
}

func (t *ALMSCRAWT)SetRc(val string){
	t.Rc = val
}


type  ArrayOfALMSCRAWT struct {
Item []ALMSCRAWT `xml:"item"`
}

func (t ArrayOfALMSCRAWT)GetItem()[]ALMSCRAWT{
	return t.Item
}

func (t *ArrayOfALMSCRAWT)SetItem(val []ALMSCRAWT){
	t.Item = val
}


type  CentralSystem struct {
SystemName string `xml:"SystemName"`
DBHost string `xml:"DBHost"`
MessageServer string `xml:"MessageServer"`
}

func (t CentralSystem)GetSystemName()string{
	return t.SystemName
}

func (t *CentralSystem)SetSystemName(val string){
	t.SystemName = val
}

func (t CentralSystem)GetDBHost()string{
	return t.DBHost
}

func (t *CentralSystem)SetDBHost(val string){
	t.DBHost = val
}

func (t CentralSystem)GetMessageServer()string{
	return t.MessageServer
}

func (t *CentralSystem)SetMessageServer(val string){
	t.MessageServer = val
}


type  ClientInfo struct {
CsmregClient string `xml:"CsmregClient"`
CsmregUser string `xml:"CsmregUser"`
CsmregPassword string `xml:"CsmregPassword"`
CsmregLanguage string `xml:"CsmregLanguage"`
UseLogonBalancing string `xml:"UseLogonBalancing"`
LogonGroup string `xml:"LogonGroup"`
Hostname string `xml:"Hostname"`
SystemNumber string `xml:"SystemNumber"`
RouteString string `xml:"RouteString"`
TraceLevel string `xml:"TraceLevel"`
}

func (t ClientInfo)GetCsmregClient()string{
	return t.CsmregClient
}

func (t *ClientInfo)SetCsmregClient(val string){
	t.CsmregClient = val
}

func (t ClientInfo)GetCsmregUser()string{
	return t.CsmregUser
}

func (t *ClientInfo)SetCsmregUser(val string){
	t.CsmregUser = val
}

func (t ClientInfo)GetCsmregPassword()string{
	return t.CsmregPassword
}

func (t *ClientInfo)SetCsmregPassword(val string){
	t.CsmregPassword = val
}

func (t ClientInfo)GetCsmregLanguage()string{
	return t.CsmregLanguage
}

func (t *ClientInfo)SetCsmregLanguage(val string){
	t.CsmregLanguage = val
}

func (t ClientInfo)GetUseLogonBalancing()string{
	return t.UseLogonBalancing
}

func (t *ClientInfo)SetUseLogonBalancing(val string){
	t.UseLogonBalancing = val
}

func (t ClientInfo)GetLogonGroup()string{
	return t.LogonGroup
}

func (t *ClientInfo)SetLogonGroup(val string){
	t.LogonGroup = val
}

func (t ClientInfo)GetHostname()string{
	return t.Hostname
}

func (t *ClientInfo)SetHostname(val string){
	t.Hostname = val
}

func (t ClientInfo)GetSystemNumber()string{
	return t.SystemNumber
}

func (t *ClientInfo)SetSystemNumber(val string){
	t.SystemNumber = val
}

func (t ClientInfo)GetRouteString()string{
	return t.RouteString
}

func (t *ClientInfo)SetRouteString(val string){
	t.RouteString = val
}

func (t ClientInfo)GetTraceLevel()string{
	return t.TraceLevel
}

func (t *ClientInfo)SetTraceLevel(val string){
	t.TraceLevel = val
}


type  CEN struct {
SLDInfo CentralSystem `xml:"SLDInfo"`
RfcClientInfo ClientInfo `xml:"RfcClientInfo"`
}

func (t CEN)GetSLDInfo()CentralSystem{
	return t.SLDInfo
}

func (t *CEN)SetSLDInfo(val CentralSystem){
	t.SLDInfo = val
}

func (t CEN)GetRfcClientInfo()ClientInfo{
	return t.RfcClientInfo
}

func (t *CEN)SetRfcClientInfo(val ClientInfo){
	t.RfcClientInfo = val
}


type  ArrayOfCEN struct {
Item []CEN `xml:"item"`
}

func (t ArrayOfCEN)GetItem()[]CEN{
	return t.Item
}

func (t *ArrayOfCEN)SetItem(val []CEN){
	t.Item = val
}


type  RegistrationInfo struct {
CEN CentralSystem `xml:"CEN"`
ABAPClient ClientInfo `xml:"ABAPClient"`
ManagedSystemName string `xml:"ManagedSystemName"`
ForceRegistration string `xml:"ForceRegistration"`
}

func (t RegistrationInfo)GetCEN()CentralSystem{
	return t.CEN
}

func (t *RegistrationInfo)SetCEN(val CentralSystem){
	t.CEN = val
}

func (t RegistrationInfo)GetABAPClient()ClientInfo{
	return t.ABAPClient
}

func (t *RegistrationInfo)SetABAPClient(val ClientInfo){
	t.ABAPClient = val
}

func (t RegistrationInfo)GetManagedSystemName()string{
	return t.ManagedSystemName
}

func (t *RegistrationInfo)SetManagedSystemName(val string){
	t.ManagedSystemName = val
}

func (t RegistrationInfo)GetForceRegistration()string{
	return t.ForceRegistration
}

func (t *RegistrationInfo)SetForceRegistration(val string){
	t.ForceRegistration = val
}


type  LogFile struct {
Filename string `xml:"filename"`
Size string `xml:"size"`
Owner string `xml:"owner"`
Group string `xml:"group"`
Modtime string `xml:"modtime"`
Type string `xml:"type"`
Format string `xml:"format"`
}

func (t LogFile)GetFilename()string{
	return t.Filename
}

func (t *LogFile)SetFilename(val string){
	t.Filename = val
}

func (t LogFile)GetSize()string{
	return t.Size
}

func (t *LogFile)SetSize(val string){
	t.Size = val
}

func (t LogFile)GetOwner()string{
	return t.Owner
}

func (t *LogFile)SetOwner(val string){
	t.Owner = val
}

func (t LogFile)GetGroup()string{
	return t.Group
}

func (t *LogFile)SetGroup(val string){
	t.Group = val
}

func (t LogFile)GetModtime()string{
	return t.Modtime
}

func (t *LogFile)SetModtime(val string){
	t.Modtime = val
}

func (t LogFile)GetType()string{
	return t.Type
}

func (t *LogFile)SetType(val string){
	t.Type = val
}

func (t LogFile)GetFormat()string{
	return t.Format
}

func (t *LogFile)SetFormat(val string){
	t.Format = val
}


type  ArrayOfLogFile struct {
Item []LogFile `xml:"item"`
}

func (t ArrayOfLogFile)GetItem()[]LogFile{
	return t.Item
}

func (t *ArrayOfLogFile)SetItem(val []LogFile){
	t.Item = val
}


type  ArrayOfString struct {
Item []string `xml:"item"`
}

func (t ArrayOfString)GetItem()[]string{
	return t.Item
}

func (t *ArrayOfString)SetItem(val []string){
	t.Item = val
}


type  FileContent struct {
Lines string `xml:"lines"`
}

func (t FileContent)GetLines()string{
	return t.Lines
}

func (t *FileContent)SetLines(val string){
	t.Lines = val
}


type  ReadFileRequest struct {
Filename string `xml:"filename"`
Filter string `xml:"filter"`
Language string `xml:"language"`
Maxentries string `xml:"maxentries"`
Statecookie string `xml:"statecookie"`
}

func (t ReadFileRequest)GetFilename()string{
	return t.Filename
}

func (t *ReadFileRequest)SetFilename(val string){
	t.Filename = val
}

func (t ReadFileRequest)GetFilter()string{
	return t.Filter
}

func (t *ReadFileRequest)SetFilter(val string){
	t.Filter = val
}

func (t ReadFileRequest)GetLanguage()string{
	return t.Language
}

func (t *ReadFileRequest)SetLanguage(val string){
	t.Language = val
}

func (t ReadFileRequest)GetMaxentries()string{
	return t.Maxentries
}

func (t *ReadFileRequest)SetMaxentries(val string){
	t.Maxentries = val
}

func (t ReadFileRequest)GetStatecookie()string{
	return t.Statecookie
}

func (t *ReadFileRequest)SetStatecookie(val string){
	t.Statecookie = val
}


type  ProfileParameter struct {
ParamName string `xml:"ParamName"`
ParamValue string `xml:"ParamValue"`
Rc string `xml:"rc"`
}

func (t ProfileParameter)GetParamName()string{
	return t.ParamName
}

func (t *ProfileParameter)SetParamName(val string){
	t.ParamName = val
}

func (t ProfileParameter)GetParamValue()string{
	return t.ParamValue
}

func (t *ProfileParameter)SetParamValue(val string){
	t.ParamValue = val
}

func (t ProfileParameter)GetRc()string{
	return t.Rc
}

func (t *ProfileParameter)SetRc(val string){
	t.Rc = val
}


type  ArrayOfProfileParameter struct {
Item []ProfileParameter `xml:"item"`
}

func (t ArrayOfProfileParameter)GetItem()[]ProfileParameter{
	return t.Item
}

func (t *ArrayOfProfileParameter)SetItem(val []ProfileParameter){
	t.Item = val
}




type ReadFile struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SoapRequest ReadFileRequest  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct ReadFile
func (t ReadFile) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct ReadFile to value val
func (t *ReadFile) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct ReadFile
func (t ReadFile) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct ReadFile to value val
func (t *ReadFile) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SoapRequest of struct ReadFile
func (t ReadFile) GetSoapRequest()ReadFileRequest{
	return t.SoapRequest
}

// Set SoapRequest of struct ReadFile to value val
func (t *ReadFile) SetSoapRequest(val ReadFileRequest){
	t.SoapRequest = val
}


type ReadFileResponse struct{
	Format string  `xml:"format"`
	Startcookie string  `xml:"startcookie"`
	Endcookie string  `xml:"endcookie"`
	FileContent FileContent  `xml:"FileContent"`
	
}

// Returns Format of struct ReadFileResponse
func (t ReadFileResponse) GetFormat()string{
	return t.Format
}

// Set Format of struct ReadFileResponse to value val
func (t *ReadFileResponse) SetFormat(val string){
	t.Format = val
}

// Returns Startcookie of struct ReadFileResponse
func (t ReadFileResponse) GetStartcookie()string{
	return t.Startcookie
}

// Set Startcookie of struct ReadFileResponse to value val
func (t *ReadFileResponse) SetStartcookie(val string){
	t.Startcookie = val
}

// Returns Endcookie of struct ReadFileResponse
func (t ReadFileResponse) GetEndcookie()string{
	return t.Endcookie
}

// Set Endcookie of struct ReadFileResponse to value val
func (t *ReadFileResponse) SetEndcookie(val string){
	t.Endcookie = val
}

// Returns FileContent of struct ReadFileResponse
func (t ReadFileResponse) GetFileContent()FileContent{
	return t.FileContent
}

// Set FileContent of struct ReadFileResponse to value val
func (t *ReadFileResponse) SetFileContent(val FileContent){
	t.FileContent = val
}


type ReadDirectory struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	DirectoryName string  `xml:"DirectoryName"`
	
}

// Returns CallingSID of struct ReadDirectory
func (t ReadDirectory) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct ReadDirectory to value val
func (t *ReadDirectory) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct ReadDirectory
func (t ReadDirectory) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct ReadDirectory to value val
func (t *ReadDirectory) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns DirectoryName of struct ReadDirectory
func (t ReadDirectory) GetDirectoryName()string{
	return t.DirectoryName
}

// Set DirectoryName of struct ReadDirectory to value val
func (t *ReadDirectory) SetDirectoryName(val string){
	t.DirectoryName = val
}


type ReadDirectoryResponse struct{
	FileList ArrayOfLogFile  `xml:"FileList"`
	
}

// Returns FileList of struct ReadDirectoryResponse
func (t ReadDirectoryResponse) GetFileList()ArrayOfLogFile{
	return t.FileList
}

// Set FileList of struct ReadDirectoryResponse to value val
func (t *ReadDirectoryResponse) SetFileList(val ArrayOfLogFile){
	t.FileList = val
}


type ReadProfileParameters struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SoapRequest ArrayOfString  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct ReadProfileParameters
func (t ReadProfileParameters) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct ReadProfileParameters to value val
func (t *ReadProfileParameters) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct ReadProfileParameters
func (t ReadProfileParameters) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct ReadProfileParameters to value val
func (t *ReadProfileParameters) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SoapRequest of struct ReadProfileParameters
func (t ReadProfileParameters) GetSoapRequest()ArrayOfString{
	return t.SoapRequest
}

// Set SoapRequest of struct ReadProfileParameters to value val
func (t *ReadProfileParameters) SetSoapRequest(val ArrayOfString){
	t.SoapRequest = val
}


type ReadProfileParametersResponse struct{
	Result ArrayOfProfileParameter  `xml:"Result"`
	
}

// Returns Result of struct ReadProfileParametersResponse
func (t ReadProfileParametersResponse) GetResult()ArrayOfProfileParameter{
	return t.Result
}

// Set Result of struct ReadProfileParametersResponse to value val
func (t *ReadProfileParametersResponse) SetResult(val ArrayOfProfileParameter){
	t.Result = val
}


type Register struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest RegistrationInfo  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct Register
func (t Register) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct Register to value val
func (t *Register) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct Register
func (t Register) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct Register to value val
func (t *Register) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct Register
func (t Register) GetSoapRequest()RegistrationInfo{
	return t.SoapRequest
}

// Set SoapRequest of struct Register to value val
func (t *Register) SetSoapRequest(val RegistrationInfo){
	t.SoapRequest = val
}


type RegisterResponse struct{
	RegistrationID string  `xml:"RegistrationID"`
	IsPrimaryCen string  `xml:"IsPrimaryCen"`
	ReturnValue string  `xml:"ReturnValue"`
	
}

// Returns RegistrationID of struct RegisterResponse
func (t RegisterResponse) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct RegisterResponse to value val
func (t *RegisterResponse) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns IsPrimaryCen of struct RegisterResponse
func (t RegisterResponse) GetIsPrimaryCen()string{
	return t.IsPrimaryCen
}

// Set IsPrimaryCen of struct RegisterResponse to value val
func (t *RegisterResponse) SetIsPrimaryCen(val string){
	t.IsPrimaryCen = val
}

// Returns ReturnValue of struct RegisterResponse
func (t RegisterResponse) GetReturnValue()string{
	return t.ReturnValue
}

// Set ReturnValue of struct RegisterResponse to value val
func (t *RegisterResponse) SetReturnValue(val string){
	t.ReturnValue = val
}


type Unregister struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest RegistrationInfo  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct Unregister
func (t Unregister) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct Unregister to value val
func (t *Unregister) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct Unregister
func (t Unregister) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct Unregister to value val
func (t *Unregister) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct Unregister
func (t Unregister) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct Unregister to value val
func (t *Unregister) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct Unregister
func (t Unregister) GetSoapRequest()RegistrationInfo{
	return t.SoapRequest
}

// Set SoapRequest of struct Unregister to value val
func (t *Unregister) SetSoapRequest(val RegistrationInfo){
	t.SoapRequest = val
}


type UnregisterResponse struct{
	ReturnValue string  `xml:"ReturnValue"`
	
}

// Returns ReturnValue of struct UnregisterResponse
func (t UnregisterResponse) GetReturnValue()string{
	return t.ReturnValue
}

// Set ReturnValue of struct UnregisterResponse to value val
func (t *UnregisterResponse) SetReturnValue(val string){
	t.ReturnValue = val
}


type GetAgentConfig struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	
}

// Returns CallingSID of struct GetAgentConfig
func (t GetAgentConfig) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct GetAgentConfig to value val
func (t *GetAgentConfig) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct GetAgentConfig
func (t GetAgentConfig) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct GetAgentConfig to value val
func (t *GetAgentConfig) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct GetAgentConfig
func (t GetAgentConfig) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct GetAgentConfig to value val
func (t *GetAgentConfig) SetSegmentType(val string){
	t.SegmentType = val
}


type AgentConfig struct{
	AgentPluginEnabled string  `xml:"AgentPluginEnabled"`
	IsInit string  `xml:"IsInit"`
	AgentType string  `xml:"AgentType"`
	AgentName string  `xml:"AgentName"`
	Version string  `xml:"Version"`
	PatchText string  `xml:"PatchText"`
	OperatingSystem string  `xml:"OperatingSystem"`
	WorkDir string  `xml:"WorkDir"`
	LogFileName string  `xml:"LogFileName"`
	CsmconfName string  `xml:"CsmconfName"`
	InifileName string  `xml:"InifileName"`
	LocalHostName string  `xml:"LocalHostName"`
	InstanceType InstanceType  `xml:"InstanceType"`
	ManagedSapSysNr string  `xml:"ManagedSapSysNr"`
	MgdSystemName string  `xml:"MgdSystemName"`
	MgdSystemNameMultipleC11s string  `xml:"MgdSystemNameMultipleC11s"`
	ConnectedToSharedMemory string  `xml:"ConnectedToSharedMemory"`
	SharedMemorySize string  `xml:"SharedMemorySize"`
	CentralCCMSAttached string  `xml:"CentralCCMSAttached"`
	NumberOfCENs string  `xml:"NumberOfCENs"`
	ListOfCentralSystems ArrayOfCEN  `xml:"ListOfCentralSystems"`
	CustomizingInLocalSystem string  `xml:"CustomizingInLocalSystem"`
	SemaphoresEnabled string  `xml:"SemaphoresEnabled"`
	
}

// Returns AgentPluginEnabled of struct AgentConfig
func (t AgentConfig) GetAgentPluginEnabled()string{
	return t.AgentPluginEnabled
}

// Set AgentPluginEnabled of struct AgentConfig to value val
func (t *AgentConfig) SetAgentPluginEnabled(val string){
	t.AgentPluginEnabled = val
}

// Returns IsInit of struct AgentConfig
func (t AgentConfig) GetIsInit()string{
	return t.IsInit
}

// Set IsInit of struct AgentConfig to value val
func (t *AgentConfig) SetIsInit(val string){
	t.IsInit = val
}

// Returns AgentType of struct AgentConfig
func (t AgentConfig) GetAgentType()string{
	return t.AgentType
}

// Set AgentType of struct AgentConfig to value val
func (t *AgentConfig) SetAgentType(val string){
	t.AgentType = val
}

// Returns AgentName of struct AgentConfig
func (t AgentConfig) GetAgentName()string{
	return t.AgentName
}

// Set AgentName of struct AgentConfig to value val
func (t *AgentConfig) SetAgentName(val string){
	t.AgentName = val
}

// Returns Version of struct AgentConfig
func (t AgentConfig) GetVersion()string{
	return t.Version
}

// Set Version of struct AgentConfig to value val
func (t *AgentConfig) SetVersion(val string){
	t.Version = val
}

// Returns PatchText of struct AgentConfig
func (t AgentConfig) GetPatchText()string{
	return t.PatchText
}

// Set PatchText of struct AgentConfig to value val
func (t *AgentConfig) SetPatchText(val string){
	t.PatchText = val
}

// Returns OperatingSystem of struct AgentConfig
func (t AgentConfig) GetOperatingSystem()string{
	return t.OperatingSystem
}

// Set OperatingSystem of struct AgentConfig to value val
func (t *AgentConfig) SetOperatingSystem(val string){
	t.OperatingSystem = val
}

// Returns WorkDir of struct AgentConfig
func (t AgentConfig) GetWorkDir()string{
	return t.WorkDir
}

// Set WorkDir of struct AgentConfig to value val
func (t *AgentConfig) SetWorkDir(val string){
	t.WorkDir = val
}

// Returns LogFileName of struct AgentConfig
func (t AgentConfig) GetLogFileName()string{
	return t.LogFileName
}

// Set LogFileName of struct AgentConfig to value val
func (t *AgentConfig) SetLogFileName(val string){
	t.LogFileName = val
}

// Returns CsmconfName of struct AgentConfig
func (t AgentConfig) GetCsmconfName()string{
	return t.CsmconfName
}

// Set CsmconfName of struct AgentConfig to value val
func (t *AgentConfig) SetCsmconfName(val string){
	t.CsmconfName = val
}

// Returns InifileName of struct AgentConfig
func (t AgentConfig) GetInifileName()string{
	return t.InifileName
}

// Set InifileName of struct AgentConfig to value val
func (t *AgentConfig) SetInifileName(val string){
	t.InifileName = val
}

// Returns LocalHostName of struct AgentConfig
func (t AgentConfig) GetLocalHostName()string{
	return t.LocalHostName
}

// Set LocalHostName of struct AgentConfig to value val
func (t *AgentConfig) SetLocalHostName(val string){
	t.LocalHostName = val
}

// Returns InstanceType of struct AgentConfig
func (t AgentConfig) GetInstanceType()InstanceType{
	return t.InstanceType
}

// Set InstanceType of struct AgentConfig to value val
func (t *AgentConfig) SetInstanceType(val InstanceType){
	t.InstanceType = val
}

// Returns ManagedSapSysNr of struct AgentConfig
func (t AgentConfig) GetManagedSapSysNr()string{
	return t.ManagedSapSysNr
}

// Set ManagedSapSysNr of struct AgentConfig to value val
func (t *AgentConfig) SetManagedSapSysNr(val string){
	t.ManagedSapSysNr = val
}

// Returns MgdSystemName of struct AgentConfig
func (t AgentConfig) GetMgdSystemName()string{
	return t.MgdSystemName
}

// Set MgdSystemName of struct AgentConfig to value val
func (t *AgentConfig) SetMgdSystemName(val string){
	t.MgdSystemName = val
}

// Returns MgdSystemNameMultipleC11s of struct AgentConfig
func (t AgentConfig) GetMgdSystemNameMultipleC11s()string{
	return t.MgdSystemNameMultipleC11s
}

// Set MgdSystemNameMultipleC11s of struct AgentConfig to value val
func (t *AgentConfig) SetMgdSystemNameMultipleC11s(val string){
	t.MgdSystemNameMultipleC11s = val
}

// Returns ConnectedToSharedMemory of struct AgentConfig
func (t AgentConfig) GetConnectedToSharedMemory()string{
	return t.ConnectedToSharedMemory
}

// Set ConnectedToSharedMemory of struct AgentConfig to value val
func (t *AgentConfig) SetConnectedToSharedMemory(val string){
	t.ConnectedToSharedMemory = val
}

// Returns SharedMemorySize of struct AgentConfig
func (t AgentConfig) GetSharedMemorySize()string{
	return t.SharedMemorySize
}

// Set SharedMemorySize of struct AgentConfig to value val
func (t *AgentConfig) SetSharedMemorySize(val string){
	t.SharedMemorySize = val
}

// Returns CentralCCMSAttached of struct AgentConfig
func (t AgentConfig) GetCentralCCMSAttached()string{
	return t.CentralCCMSAttached
}

// Set CentralCCMSAttached of struct AgentConfig to value val
func (t *AgentConfig) SetCentralCCMSAttached(val string){
	t.CentralCCMSAttached = val
}

// Returns NumberOfCENs of struct AgentConfig
func (t AgentConfig) GetNumberOfCENs()string{
	return t.NumberOfCENs
}

// Set NumberOfCENs of struct AgentConfig to value val
func (t *AgentConfig) SetNumberOfCENs(val string){
	t.NumberOfCENs = val
}

// Returns ListOfCentralSystems of struct AgentConfig
func (t AgentConfig) GetListOfCentralSystems()ArrayOfCEN{
	return t.ListOfCentralSystems
}

// Set ListOfCentralSystems of struct AgentConfig to value val
func (t *AgentConfig) SetListOfCentralSystems(val ArrayOfCEN){
	t.ListOfCentralSystems = val
}

// Returns CustomizingInLocalSystem of struct AgentConfig
func (t AgentConfig) GetCustomizingInLocalSystem()string{
	return t.CustomizingInLocalSystem
}

// Set CustomizingInLocalSystem of struct AgentConfig to value val
func (t *AgentConfig) SetCustomizingInLocalSystem(val string){
	t.CustomizingInLocalSystem = val
}

// Returns SemaphoresEnabled of struct AgentConfig
func (t AgentConfig) GetSemaphoresEnabled()string{
	return t.SemaphoresEnabled
}

// Set SemaphoresEnabled of struct AgentConfig to value val
func (t *AgentConfig) SetSemaphoresEnabled(val string){
	t.SemaphoresEnabled = val
}


type MtGetTidByName struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALMTNAMEL  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtGetTidByName
func (t MtGetTidByName) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtGetTidByName to value val
func (t *MtGetTidByName) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MtGetTidByName
func (t MtGetTidByName) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtGetTidByName to value val
func (t *MtGetTidByName) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MtGetTidByName
func (t MtGetTidByName) GetSoapRequest()ArrayOfALMTNAMEL{
	return t.SoapRequest
}

// Set SoapRequest of struct MtGetTidByName to value val
func (t *MtGetTidByName) SetSoapRequest(val ArrayOfALMTNAMEL){
	t.SoapRequest = val
}


type MtGetTidByNameResponse struct{
	TidTable ArrayOfALGTIDLNRC  `xml:"TidTable"`
	
}

// Returns TidTable of struct MtGetTidByNameResponse
func (t MtGetTidByNameResponse) GetTidTable()ArrayOfALGTIDLNRC{
	return t.TidTable
}

// Set TidTable of struct MtGetTidByNameResponse to value val
func (t *MtGetTidByNameResponse) SetTidTable(val ArrayOfALGTIDLNRC){
	t.TidTable = val
}


type MtGetMteinfo struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtGetMteinfo
func (t MtGetMteinfo) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtGetMteinfo to value val
func (t *MtGetMteinfo) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MtGetMteinfo
func (t MtGetMteinfo) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtGetMteinfo to value val
func (t *MtGetMteinfo) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MtGetMteinfo
func (t MtGetMteinfo) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct MtGetMteinfo to value val
func (t *MtGetMteinfo) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type MtGetMteinfoResponse struct{
	MteInfo ArrayOfALMTEIRC  `xml:"MteInfo"`
	LongName ArrayOfALGTIDLNRC  `xml:"LongName"`
	
}

// Returns MteInfo of struct MtGetMteinfoResponse
func (t MtGetMteinfoResponse) GetMteInfo()ArrayOfALMTEIRC{
	return t.MteInfo
}

// Set MteInfo of struct MtGetMteinfoResponse to value val
func (t *MtGetMteinfoResponse) SetMteInfo(val ArrayOfALMTEIRC){
	t.MteInfo = val
}

// Returns LongName of struct MtGetMteinfoResponse
func (t MtGetMteinfoResponse) GetLongName()ArrayOfALGTIDLNRC{
	return t.LongName
}

// Set LongName of struct MtGetMteinfoResponse to value val
func (t *MtGetMteinfoResponse) SetLongName(val ArrayOfALGTIDLNRC){
	t.LongName = val
}


type MtRead struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	IGetEffTools string  `xml:"iGetEffTools"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtRead
func (t MtRead) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtRead to value val
func (t *MtRead) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MtRead
func (t MtRead) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtRead to value val
func (t *MtRead) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns IGetEffTools of struct MtRead
func (t MtRead) GetIGetEffTools()string{
	return t.IGetEffTools
}

// Set IGetEffTools of struct MtRead to value val
func (t *MtRead) SetIGetEffTools(val string){
	t.IGetEffTools = val
}

// Returns SoapRequest of struct MtRead
func (t MtRead) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct MtRead to value val
func (t *MtRead) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type MtReadResponse struct{
	TidTable ArrayOfALTDEFRC  `xml:"TidTable"`
	
}

// Returns TidTable of struct MtReadResponse
func (t MtReadResponse) GetTidTable()ArrayOfALTDEFRC{
	return t.TidTable
}

// Set TidTable of struct MtReadResponse to value val
func (t *MtReadResponse) SetTidTable(val ArrayOfALTDEFRC){
	t.TidTable = val
}


type PerfRead struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct PerfRead
func (t PerfRead) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct PerfRead to value val
func (t *PerfRead) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct PerfRead
func (t PerfRead) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct PerfRead to value val
func (t *PerfRead) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct PerfRead
func (t PerfRead) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct PerfRead to value val
func (t *PerfRead) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type PerfReadResponse struct{
	TidTable ArrayOfALPERFTYPE  `xml:"TidTable"`
	
}

// Returns TidTable of struct PerfReadResponse
func (t PerfReadResponse) GetTidTable()ArrayOfALPERFTYPE{
	return t.TidTable
}

// Set TidTable of struct PerfReadResponse to value val
func (t *PerfReadResponse) SetTidTable(val ArrayOfALPERFTYPE){
	t.TidTable = val
}


type PerfReadSmoothData struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct PerfReadSmoothData
func (t PerfReadSmoothData) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct PerfReadSmoothData to value val
func (t *PerfReadSmoothData) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct PerfReadSmoothData
func (t PerfReadSmoothData) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct PerfReadSmoothData to value val
func (t *PerfReadSmoothData) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct PerfReadSmoothData
func (t PerfReadSmoothData) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct PerfReadSmoothData to value val
func (t *PerfReadSmoothData) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type PerfReadSmoothDataResponse struct{
	TidTable ArrayOfALGTIDSMO  `xml:"TidTable"`
	
}

// Returns TidTable of struct PerfReadSmoothDataResponse
func (t PerfReadSmoothDataResponse) GetTidTable()ArrayOfALGTIDSMO{
	return t.TidTable
}

// Set TidTable of struct PerfReadSmoothDataResponse to value val
func (t *PerfReadSmoothDataResponse) SetTidTable(val ArrayOfALGTIDSMO){
	t.TidTable = val
}


type MscReadCache struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALMSCSELEC  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MscReadCache
func (t MscReadCache) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MscReadCache to value val
func (t *MscReadCache) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MscReadCache
func (t MscReadCache) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MscReadCache to value val
func (t *MscReadCache) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MscReadCache
func (t MscReadCache) GetSoapRequest()ArrayOfALMSCSELEC{
	return t.SoapRequest
}

// Set SoapRequest of struct MscReadCache to value val
func (t *MscReadCache) SetSoapRequest(val ArrayOfALMSCSELEC){
	t.SoapRequest = val
}


type MscReadCacheResponse struct{
	TidRc ArrayOfALGTIDRC  `xml:"TidRc"`
	CacheLines ArrayOfALMSCTIDRL  `xml:"CacheLines"`
	
}

// Returns TidRc of struct MscReadCacheResponse
func (t MscReadCacheResponse) GetTidRc()ArrayOfALGTIDRC{
	return t.TidRc
}

// Set TidRc of struct MscReadCacheResponse to value val
func (t *MscReadCacheResponse) SetTidRc(val ArrayOfALGTIDRC){
	t.TidRc = val
}

// Returns CacheLines of struct MscReadCacheResponse
func (t MscReadCacheResponse) GetCacheLines()ArrayOfALMSCTIDRL{
	return t.CacheLines
}

// Set CacheLines of struct MscReadCacheResponse to value val
func (t *MscReadCacheResponse) SetCacheLines(val ArrayOfALMSCTIDRL){
	t.CacheLines = val
}


type UtilSnglmsgReadRawdata struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct UtilSnglmsgReadRawdata
func (t UtilSnglmsgReadRawdata) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct UtilSnglmsgReadRawdata to value val
func (t *UtilSnglmsgReadRawdata) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct UtilSnglmsgReadRawdata
func (t UtilSnglmsgReadRawdata) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct UtilSnglmsgReadRawdata to value val
func (t *UtilSnglmsgReadRawdata) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct UtilSnglmsgReadRawdata
func (t UtilSnglmsgReadRawdata) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct UtilSnglmsgReadRawdata to value val
func (t *UtilSnglmsgReadRawdata) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type UtilSnglmsgReadRawdataResponse struct{
	SnglMsgList ArrayOfALSMSGRAWT  `xml:"SnglMsgList"`
	
}

// Returns SnglMsgList of struct UtilSnglmsgReadRawdataResponse
func (t UtilSnglmsgReadRawdataResponse) GetSnglMsgList()ArrayOfALSMSGRAWT{
	return t.SnglMsgList
}

// Set SnglMsgList of struct UtilSnglmsgReadRawdataResponse to value val
func (t *UtilSnglmsgReadRawdataResponse) SetSnglMsgList(val ArrayOfALSMSGRAWT){
	t.SnglMsgList = val
}


type UtilMtReadAll struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	GetTree string  `xml:"GetTree"`
	GetGeneralData string  `xml:"GetGeneralData"`
	GetMscData string  `xml:"GetMscData"`
	GetPerfData string  `xml:"GetPerfData"`
	GetPerfSmoothData string  `xml:"GetPerfSmoothData"`
	GetSmesData string  `xml:"GetSmesData"`
	GetTextAttr string  `xml:"GetTextAttr"`
	SoapRequest ArrayOfALTIDMAXLV  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct UtilMtReadAll
func (t UtilMtReadAll) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct UtilMtReadAll
func (t UtilMtReadAll) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns GetTree of struct UtilMtReadAll
func (t UtilMtReadAll) GetGetTree()string{
	return t.GetTree
}

// Set GetTree of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetGetTree(val string){
	t.GetTree = val
}

// Returns GetGeneralData of struct UtilMtReadAll
func (t UtilMtReadAll) GetGetGeneralData()string{
	return t.GetGeneralData
}

// Set GetGeneralData of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetGetGeneralData(val string){
	t.GetGeneralData = val
}

// Returns GetMscData of struct UtilMtReadAll
func (t UtilMtReadAll) GetGetMscData()string{
	return t.GetMscData
}

// Set GetMscData of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetGetMscData(val string){
	t.GetMscData = val
}

// Returns GetPerfData of struct UtilMtReadAll
func (t UtilMtReadAll) GetGetPerfData()string{
	return t.GetPerfData
}

// Set GetPerfData of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetGetPerfData(val string){
	t.GetPerfData = val
}

// Returns GetPerfSmoothData of struct UtilMtReadAll
func (t UtilMtReadAll) GetGetPerfSmoothData()string{
	return t.GetPerfSmoothData
}

// Set GetPerfSmoothData of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetGetPerfSmoothData(val string){
	t.GetPerfSmoothData = val
}

// Returns GetSmesData of struct UtilMtReadAll
func (t UtilMtReadAll) GetGetSmesData()string{
	return t.GetSmesData
}

// Set GetSmesData of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetGetSmesData(val string){
	t.GetSmesData = val
}

// Returns GetTextAttr of struct UtilMtReadAll
func (t UtilMtReadAll) GetGetTextAttr()string{
	return t.GetTextAttr
}

// Set GetTextAttr of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetGetTextAttr(val string){
	t.GetTextAttr = val
}

// Returns SoapRequest of struct UtilMtReadAll
func (t UtilMtReadAll) GetSoapRequest()ArrayOfALTIDMAXLV{
	return t.SoapRequest
}

// Set SoapRequest of struct UtilMtReadAll to value val
func (t *UtilMtReadAll) SetSoapRequest(val ArrayOfALTIDMAXLV){
	t.SoapRequest = val
}


type UtilMtReadAllResponse struct{
	TreeInfo ArrayOfALMTTREE  `xml:"TreeInfo"`
	TidRc ArrayOfALGTIDRC  `xml:"TidRc"`
	GeneralInfo ArrayOfALTDEFRC  `xml:"GeneralInfo"`
	PerfInfo ArrayOfALPERFTYPE  `xml:"PerfInfo"`
	PerfSmoothInfo ArrayOfALGTIDSMO  `xml:"PerfSmoothInfo"`
	MscInfo ArrayOfALMSCRAWT  `xml:"MscInfo"`
	MscFilterInfo ArrayOfALMSCTIDFL  `xml:"MscFilterInfo"`
	SmesInfo ArrayOfALSMSGRAWT  `xml:"SmesInfo"`
	TextAttrInfo ArrayOfALTXTLNKAT  `xml:"TextAttrInfo"`
	
}

// Returns TreeInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetTreeInfo()ArrayOfALMTTREE{
	return t.TreeInfo
}

// Set TreeInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetTreeInfo(val ArrayOfALMTTREE){
	t.TreeInfo = val
}

// Returns TidRc of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetTidRc()ArrayOfALGTIDRC{
	return t.TidRc
}

// Set TidRc of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetTidRc(val ArrayOfALGTIDRC){
	t.TidRc = val
}

// Returns GeneralInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetGeneralInfo()ArrayOfALTDEFRC{
	return t.GeneralInfo
}

// Set GeneralInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetGeneralInfo(val ArrayOfALTDEFRC){
	t.GeneralInfo = val
}

// Returns PerfInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetPerfInfo()ArrayOfALPERFTYPE{
	return t.PerfInfo
}

// Set PerfInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetPerfInfo(val ArrayOfALPERFTYPE){
	t.PerfInfo = val
}

// Returns PerfSmoothInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetPerfSmoothInfo()ArrayOfALGTIDSMO{
	return t.PerfSmoothInfo
}

// Set PerfSmoothInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetPerfSmoothInfo(val ArrayOfALGTIDSMO){
	t.PerfSmoothInfo = val
}

// Returns MscInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetMscInfo()ArrayOfALMSCRAWT{
	return t.MscInfo
}

// Set MscInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetMscInfo(val ArrayOfALMSCRAWT){
	t.MscInfo = val
}

// Returns MscFilterInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetMscFilterInfo()ArrayOfALMSCTIDFL{
	return t.MscFilterInfo
}

// Set MscFilterInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetMscFilterInfo(val ArrayOfALMSCTIDFL){
	t.MscFilterInfo = val
}

// Returns SmesInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetSmesInfo()ArrayOfALSMSGRAWT{
	return t.SmesInfo
}

// Set SmesInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetSmesInfo(val ArrayOfALSMSGRAWT){
	t.SmesInfo = val
}

// Returns TextAttrInfo of struct UtilMtReadAllResponse
func (t UtilMtReadAllResponse) GetTextAttrInfo()ArrayOfALTXTLNKAT{
	return t.TextAttrInfo
}

// Set TextAttrInfo of struct UtilMtReadAllResponse to value val
func (t *UtilMtReadAllResponse) SetTextAttrInfo(val ArrayOfALTXTLNKAT){
	t.TextAttrInfo = val
}


type TextAttrRead struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct TextAttrRead
func (t TextAttrRead) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct TextAttrRead to value val
func (t *TextAttrRead) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct TextAttrRead
func (t TextAttrRead) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct TextAttrRead to value val
func (t *TextAttrRead) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct TextAttrRead
func (t TextAttrRead) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct TextAttrRead to value val
func (t *TextAttrRead) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type TextAttrReadResponse struct{
	TextAttrList ArrayOfALTEXTATTR  `xml:"TextAttrList"`
	
}

// Returns TextAttrList of struct TextAttrReadResponse
func (t TextAttrReadResponse) GetTextAttrList()ArrayOfALTEXTATTR{
	return t.TextAttrList
}

// Set TextAttrList of struct TextAttrReadResponse to value val
func (t *TextAttrReadResponse) SetTextAttrList(val ArrayOfALTEXTATTR){
	t.TextAttrList = val
}


type ReferenceRead struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct ReferenceRead
func (t ReferenceRead) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct ReferenceRead to value val
func (t *ReferenceRead) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct ReferenceRead
func (t ReferenceRead) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct ReferenceRead to value val
func (t *ReferenceRead) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct ReferenceRead
func (t ReferenceRead) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct ReferenceRead to value val
func (t *ReferenceRead) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type ReferenceReadResponse struct{
	ReferenceList ArrayOfALTXTLNKAT  `xml:"ReferenceList"`
	
}

// Returns ReferenceList of struct ReferenceReadResponse
func (t ReferenceReadResponse) GetReferenceList()ArrayOfALTXTLNKAT{
	return t.ReferenceList
}

// Set ReferenceList of struct ReferenceReadResponse to value val
func (t *ReferenceReadResponse) SetReferenceList(val ArrayOfALTXTLNKAT){
	t.ReferenceList = val
}


type UtilMtGetAidByTid struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct UtilMtGetAidByTid
func (t UtilMtGetAidByTid) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct UtilMtGetAidByTid to value val
func (t *UtilMtGetAidByTid) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct UtilMtGetAidByTid
func (t UtilMtGetAidByTid) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct UtilMtGetAidByTid to value val
func (t *UtilMtGetAidByTid) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct UtilMtGetAidByTid
func (t UtilMtGetAidByTid) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct UtilMtGetAidByTid to value val
func (t *UtilMtGetAidByTid) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type UtilMtGetAidByTidResponse struct{
	TidAid ArrayOfALCCTIDAID  `xml:"TidAid"`
	TidRc ArrayOfALGTIDRC  `xml:"TidRc"`
	
}

// Returns TidAid of struct UtilMtGetAidByTidResponse
func (t UtilMtGetAidByTidResponse) GetTidAid()ArrayOfALCCTIDAID{
	return t.TidAid
}

// Set TidAid of struct UtilMtGetAidByTidResponse to value val
func (t *UtilMtGetAidByTidResponse) SetTidAid(val ArrayOfALCCTIDAID){
	t.TidAid = val
}

// Returns TidRc of struct UtilMtGetAidByTidResponse
func (t UtilMtGetAidByTidResponse) GetTidRc()ArrayOfALGTIDRC{
	return t.TidRc
}

// Set TidRc of struct UtilMtGetAidByTidResponse to value val
func (t *UtilMtGetAidByTidResponse) SetTidRc(val ArrayOfALGTIDRC){
	t.TidRc = val
}


type UtilReadRawalertByAid struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBAID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct UtilReadRawalertByAid
func (t UtilReadRawalertByAid) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct UtilReadRawalertByAid to value val
func (t *UtilReadRawalertByAid) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct UtilReadRawalertByAid
func (t UtilReadRawalertByAid) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct UtilReadRawalertByAid to value val
func (t *UtilReadRawalertByAid) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct UtilReadRawalertByAid
func (t UtilReadRawalertByAid) GetSoapRequest()ArrayOfALGLOBAID{
	return t.SoapRequest
}

// Set SoapRequest of struct UtilReadRawalertByAid to value val
func (t *UtilReadRawalertByAid) SetSoapRequest(val ArrayOfALGLOBAID){
	t.SoapRequest = val
}


type UtilReadRawalertByAidResponse struct{
	RawAlertList ArrayOfALALRAWRC  `xml:"RawAlertList"`
	
}

// Returns RawAlertList of struct UtilReadRawalertByAidResponse
func (t UtilReadRawalertByAidResponse) GetRawAlertList()ArrayOfALALRAWRC{
	return t.RawAlertList
}

// Set RawAlertList of struct UtilReadRawalertByAidResponse to value val
func (t *UtilReadRawalertByAidResponse) SetRawAlertList(val ArrayOfALALRAWRC){
	t.RawAlertList = val
}


type ToolGetEffective struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGTIDWHTL  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct ToolGetEffective
func (t ToolGetEffective) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct ToolGetEffective to value val
func (t *ToolGetEffective) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct ToolGetEffective
func (t ToolGetEffective) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct ToolGetEffective to value val
func (t *ToolGetEffective) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct ToolGetEffective
func (t ToolGetEffective) GetSoapRequest()ArrayOfALGTIDWHTL{
	return t.SoapRequest
}

// Set SoapRequest of struct ToolGetEffective to value val
func (t *ToolGetEffective) SetSoapRequest(val ArrayOfALGTIDWHTL){
	t.SoapRequest = val
}


type ToolGetEffectiveResponse struct{
	ToolList ArrayOfALTOOLEFRC  `xml:"ToolList"`
	
}

// Returns ToolList of struct ToolGetEffectiveResponse
func (t ToolGetEffectiveResponse) GetToolList()ArrayOfALTOOLEFRC{
	return t.ToolList
}

// Set ToolList of struct ToolGetEffectiveResponse to value val
func (t *ToolGetEffectiveResponse) SetToolList(val ArrayOfALTOOLEFRC){
	t.ToolList = val
}


type TriggerDataCollection struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALTOOLASSG  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct TriggerDataCollection
func (t TriggerDataCollection) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct TriggerDataCollection to value val
func (t *TriggerDataCollection) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct TriggerDataCollection
func (t TriggerDataCollection) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct TriggerDataCollection to value val
func (t *TriggerDataCollection) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct TriggerDataCollection
func (t TriggerDataCollection) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct TriggerDataCollection to value val
func (t *TriggerDataCollection) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct TriggerDataCollection
func (t TriggerDataCollection) GetSoapRequest()ArrayOfALTOOLASSG{
	return t.SoapRequest
}

// Set SoapRequest of struct TriggerDataCollection to value val
func (t *TriggerDataCollection) SetSoapRequest(val ArrayOfALTOOLASSG){
	t.SoapRequest = val
}


type TriggerDataCollectionResponse struct{
	Rc string  `xml:"Rc"`
	
}

// Returns Rc of struct TriggerDataCollectionResponse
func (t TriggerDataCollectionResponse) GetRc()string{
	return t.Rc
}

// Set Rc of struct TriggerDataCollectionResponse to value val
func (t *TriggerDataCollectionResponse) SetRc(val string){
	t.Rc = val
}


type MtChangeStatus struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGTIDSTAT  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtChangeStatus
func (t MtChangeStatus) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtChangeStatus to value val
func (t *MtChangeStatus) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct MtChangeStatus
func (t MtChangeStatus) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct MtChangeStatus to value val
func (t *MtChangeStatus) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct MtChangeStatus
func (t MtChangeStatus) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtChangeStatus to value val
func (t *MtChangeStatus) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MtChangeStatus
func (t MtChangeStatus) GetSoapRequest()ArrayOfALGTIDSTAT{
	return t.SoapRequest
}

// Set SoapRequest of struct MtChangeStatus to value val
func (t *MtChangeStatus) SetSoapRequest(val ArrayOfALGTIDSTAT){
	t.SoapRequest = val
}


type MtChangeStatusResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct MtChangeStatusResponse
func (t MtChangeStatusResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct MtChangeStatusResponse to value val
func (t *MtChangeStatusResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type MtCustomizeWrite struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALMTCUSTWR  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtCustomizeWrite
func (t MtCustomizeWrite) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtCustomizeWrite to value val
func (t *MtCustomizeWrite) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct MtCustomizeWrite
func (t MtCustomizeWrite) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct MtCustomizeWrite to value val
func (t *MtCustomizeWrite) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct MtCustomizeWrite
func (t MtCustomizeWrite) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtCustomizeWrite to value val
func (t *MtCustomizeWrite) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MtCustomizeWrite
func (t MtCustomizeWrite) GetSoapRequest()ArrayOfALMTCUSTWR{
	return t.SoapRequest
}

// Set SoapRequest of struct MtCustomizeWrite to value val
func (t *MtCustomizeWrite) SetSoapRequest(val ArrayOfALMTCUSTWR){
	t.SoapRequest = val
}


type MtCustomizeWriteResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct MtCustomizeWriteResponse
func (t MtCustomizeWriteResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct MtCustomizeWriteResponse to value val
func (t *MtCustomizeWriteResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type MtDestroyMarkNTry struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtDestroyMarkNTry
func (t MtDestroyMarkNTry) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtDestroyMarkNTry to value val
func (t *MtDestroyMarkNTry) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct MtDestroyMarkNTry
func (t MtDestroyMarkNTry) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct MtDestroyMarkNTry to value val
func (t *MtDestroyMarkNTry) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct MtDestroyMarkNTry
func (t MtDestroyMarkNTry) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtDestroyMarkNTry to value val
func (t *MtDestroyMarkNTry) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MtDestroyMarkNTry
func (t MtDestroyMarkNTry) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct MtDestroyMarkNTry to value val
func (t *MtDestroyMarkNTry) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type MtDestroyMarkNTryResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct MtDestroyMarkNTryResponse
func (t MtDestroyMarkNTryResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct MtDestroyMarkNTryResponse to value val
func (t *MtDestroyMarkNTryResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type MtReset struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALTIDUSER  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtReset
func (t MtReset) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtReset to value val
func (t *MtReset) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct MtReset
func (t MtReset) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct MtReset to value val
func (t *MtReset) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct MtReset
func (t MtReset) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtReset to value val
func (t *MtReset) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MtReset
func (t MtReset) GetSoapRequest()ArrayOfALTIDUSER{
	return t.SoapRequest
}

// Set SoapRequest of struct MtReset to value val
func (t *MtReset) SetSoapRequest(val ArrayOfALTIDUSER){
	t.SoapRequest = val
}


type MtResetResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct MtResetResponse
func (t MtResetResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct MtResetResponse to value val
func (t *MtResetResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type PerfCustomizeWrite struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALPERFCUSW  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct PerfCustomizeWrite
func (t PerfCustomizeWrite) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct PerfCustomizeWrite to value val
func (t *PerfCustomizeWrite) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct PerfCustomizeWrite
func (t PerfCustomizeWrite) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct PerfCustomizeWrite to value val
func (t *PerfCustomizeWrite) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct PerfCustomizeWrite
func (t PerfCustomizeWrite) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct PerfCustomizeWrite to value val
func (t *PerfCustomizeWrite) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct PerfCustomizeWrite
func (t PerfCustomizeWrite) GetSoapRequest()ArrayOfALPERFCUSW{
	return t.SoapRequest
}

// Set SoapRequest of struct PerfCustomizeWrite to value val
func (t *PerfCustomizeWrite) SetSoapRequest(val ArrayOfALPERFCUSW){
	t.SoapRequest = val
}


type PerfCustomizeWriteResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct PerfCustomizeWriteResponse
func (t PerfCustomizeWriteResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct PerfCustomizeWriteResponse to value val
func (t *PerfCustomizeWriteResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type MscCustomizeWrite struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	ThMsccustwrite ArrayOfALMSCCUSW  `xml:"thMsccustwrite"`
	ThMsctidfilter ArrayOfALMSCTIDFL  `xml:"thMsctidfilter"`
	
}

// Returns CallingSID of struct MscCustomizeWrite
func (t MscCustomizeWrite) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MscCustomizeWrite to value val
func (t *MscCustomizeWrite) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct MscCustomizeWrite
func (t MscCustomizeWrite) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct MscCustomizeWrite to value val
func (t *MscCustomizeWrite) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct MscCustomizeWrite
func (t MscCustomizeWrite) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MscCustomizeWrite to value val
func (t *MscCustomizeWrite) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns ThMsccustwrite of struct MscCustomizeWrite
func (t MscCustomizeWrite) GetThMsccustwrite()ArrayOfALMSCCUSW{
	return t.ThMsccustwrite
}

// Set ThMsccustwrite of struct MscCustomizeWrite to value val
func (t *MscCustomizeWrite) SetThMsccustwrite(val ArrayOfALMSCCUSW){
	t.ThMsccustwrite = val
}

// Returns ThMsctidfilter of struct MscCustomizeWrite
func (t MscCustomizeWrite) GetThMsctidfilter()ArrayOfALMSCTIDFL{
	return t.ThMsctidfilter
}

// Set ThMsctidfilter of struct MscCustomizeWrite to value val
func (t *MscCustomizeWrite) SetThMsctidfilter(val ArrayOfALMSCTIDFL){
	t.ThMsctidfilter = val
}


type MscCustomizeWriteResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct MscCustomizeWriteResponse
func (t MscCustomizeWriteResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct MscCustomizeWriteResponse to value val
func (t *MscCustomizeWriteResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type MscDeleteLines struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALMSCSELEC  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MscDeleteLines
func (t MscDeleteLines) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MscDeleteLines to value val
func (t *MscDeleteLines) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct MscDeleteLines
func (t MscDeleteLines) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct MscDeleteLines to value val
func (t *MscDeleteLines) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct MscDeleteLines
func (t MscDeleteLines) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MscDeleteLines to value val
func (t *MscDeleteLines) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MscDeleteLines
func (t MscDeleteLines) GetSoapRequest()ArrayOfALMSCSELEC{
	return t.SoapRequest
}

// Set SoapRequest of struct MscDeleteLines to value val
func (t *MscDeleteLines) SetSoapRequest(val ArrayOfALMSCSELEC){
	t.SoapRequest = val
}


type MscDeleteLinesResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct MscDeleteLinesResponse
func (t MscDeleteLinesResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct MscDeleteLinesResponse to value val
func (t *MscDeleteLinesResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type SnglmgsCustomizeWrite struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALSMSGCUSW  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct SnglmgsCustomizeWrite
func (t SnglmgsCustomizeWrite) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct SnglmgsCustomizeWrite to value val
func (t *SnglmgsCustomizeWrite) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct SnglmgsCustomizeWrite
func (t SnglmgsCustomizeWrite) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct SnglmgsCustomizeWrite to value val
func (t *SnglmgsCustomizeWrite) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct SnglmgsCustomizeWrite
func (t SnglmgsCustomizeWrite) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct SnglmgsCustomizeWrite to value val
func (t *SnglmgsCustomizeWrite) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct SnglmgsCustomizeWrite
func (t SnglmgsCustomizeWrite) GetSoapRequest()ArrayOfALSMSGCUSW{
	return t.SoapRequest
}

// Set SoapRequest of struct SnglmgsCustomizeWrite to value val
func (t *SnglmgsCustomizeWrite) SetSoapRequest(val ArrayOfALSMSGCUSW){
	t.SoapRequest = val
}


type SnglmgsCustomizeWriteResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct SnglmgsCustomizeWriteResponse
func (t SnglmgsCustomizeWriteResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct SnglmgsCustomizeWriteResponse to value val
func (t *SnglmgsCustomizeWriteResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type ToolSet struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	IDeleteAllOtherCenauto string  `xml:"iDeleteAllOtherCenauto"`
	SoapRequest ArrayOfALTOOLASSG  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct ToolSet
func (t ToolSet) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct ToolSet to value val
func (t *ToolSet) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct ToolSet
func (t ToolSet) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct ToolSet to value val
func (t *ToolSet) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct ToolSet
func (t ToolSet) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct ToolSet to value val
func (t *ToolSet) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns IDeleteAllOtherCenauto of struct ToolSet
func (t ToolSet) GetIDeleteAllOtherCenauto()string{
	return t.IDeleteAllOtherCenauto
}

// Set IDeleteAllOtherCenauto of struct ToolSet to value val
func (t *ToolSet) SetIDeleteAllOtherCenauto(val string){
	t.IDeleteAllOtherCenauto = val
}

// Returns SoapRequest of struct ToolSet
func (t ToolSet) GetSoapRequest()ArrayOfALTOOLASSG{
	return t.SoapRequest
}

// Set SoapRequest of struct ToolSet to value val
func (t *ToolSet) SetSoapRequest(val ArrayOfALTOOLASSG){
	t.SoapRequest = val
}


type ToolSetResponse struct{
	ResultList ArrayOfALGTIDTLRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct ToolSetResponse
func (t ToolSetResponse) GetResultList()ArrayOfALGTIDTLRC{
	return t.ResultList
}

// Set ResultList of struct ToolSetResponse to value val
func (t *ToolSetResponse) SetResultList(val ArrayOfALGTIDTLRC){
	t.ResultList = val
}


type ToolSetRuntimeStatus struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALTOOLCHST  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct ToolSetRuntimeStatus
func (t ToolSetRuntimeStatus) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct ToolSetRuntimeStatus to value val
func (t *ToolSetRuntimeStatus) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct ToolSetRuntimeStatus
func (t ToolSetRuntimeStatus) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct ToolSetRuntimeStatus to value val
func (t *ToolSetRuntimeStatus) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct ToolSetRuntimeStatus
func (t ToolSetRuntimeStatus) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct ToolSetRuntimeStatus to value val
func (t *ToolSetRuntimeStatus) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct ToolSetRuntimeStatus
func (t ToolSetRuntimeStatus) GetSoapRequest()ArrayOfALTOOLCHST{
	return t.SoapRequest
}

// Set SoapRequest of struct ToolSetRuntimeStatus to value val
func (t *ToolSetRuntimeStatus) SetSoapRequest(val ArrayOfALTOOLCHST){
	t.SoapRequest = val
}


type ToolSetRuntimeStatusResponse struct{
	ResultList ArrayOfALGTIDTLRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct ToolSetRuntimeStatusResponse
func (t ToolSetRuntimeStatusResponse) GetResultList()ArrayOfALGTIDTLRC{
	return t.ResultList
}

// Set ResultList of struct ToolSetRuntimeStatusResponse to value val
func (t *ToolSetRuntimeStatusResponse) SetResultList(val ArrayOfALGTIDTLRC){
	t.ResultList = val
}


type UtilAlChangeStatus struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALAIDALSTA  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct UtilAlChangeStatus
func (t UtilAlChangeStatus) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct UtilAlChangeStatus to value val
func (t *UtilAlChangeStatus) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct UtilAlChangeStatus
func (t UtilAlChangeStatus) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct UtilAlChangeStatus to value val
func (t *UtilAlChangeStatus) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct UtilAlChangeStatus
func (t UtilAlChangeStatus) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct UtilAlChangeStatus to value val
func (t *UtilAlChangeStatus) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct UtilAlChangeStatus
func (t UtilAlChangeStatus) GetSoapRequest()ArrayOfALAIDALSTA{
	return t.SoapRequest
}

// Set SoapRequest of struct UtilAlChangeStatus to value val
func (t *UtilAlChangeStatus) SetSoapRequest(val ArrayOfALAIDALSTA){
	t.SoapRequest = val
}


type UtilAlChangeStatusResponse struct{
	ResultList ArrayOfALGAIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct UtilAlChangeStatusResponse
func (t UtilAlChangeStatusResponse) GetResultList()ArrayOfALGAIDRC{
	return t.ResultList
}

// Set ResultList of struct UtilAlChangeStatusResponse to value val
func (t *UtilAlChangeStatusResponse) SetResultList(val ArrayOfALGAIDRC){
	t.ResultList = val
}


type MtDbsetToWpsetByTid struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MtDbsetToWpsetByTid
func (t MtDbsetToWpsetByTid) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtDbsetToWpsetByTid to value val
func (t *MtDbsetToWpsetByTid) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct MtDbsetToWpsetByTid
func (t MtDbsetToWpsetByTid) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct MtDbsetToWpsetByTid to value val
func (t *MtDbsetToWpsetByTid) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct MtDbsetToWpsetByTid
func (t MtDbsetToWpsetByTid) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtDbsetToWpsetByTid to value val
func (t *MtDbsetToWpsetByTid) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct MtDbsetToWpsetByTid
func (t MtDbsetToWpsetByTid) GetSoapRequest()ArrayOfALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct MtDbsetToWpsetByTid to value val
func (t *MtDbsetToWpsetByTid) SetSoapRequest(val ArrayOfALGLOBTID){
	t.SoapRequest = val
}


type MtDbsetToWpsetByTidResponse struct{
	ResultList ArrayOfALGTIDRC  `xml:"ResultList"`
	
}

// Returns ResultList of struct MtDbsetToWpsetByTidResponse
func (t MtDbsetToWpsetByTidResponse) GetResultList()ArrayOfALGTIDRC{
	return t.ResultList
}

// Set ResultList of struct MtDbsetToWpsetByTidResponse to value val
func (t *MtDbsetToWpsetByTidResponse) SetResultList(val ArrayOfALGTIDRC){
	t.ResultList = val
}


type SystemObjectSetValue struct{
	CallingSID string  `xml:"CallingSID"`
	RegistrationID string  `xml:"RegistrationID"`
	SegmentType string  `xml:"SegmentType"`
	Object string  `xml:"Object"`
	Value string  `xml:"Value"`
	Text string  `xml:"Text"`
	
}

// Returns CallingSID of struct SystemObjectSetValue
func (t SystemObjectSetValue) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct SystemObjectSetValue to value val
func (t *SystemObjectSetValue) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns RegistrationID of struct SystemObjectSetValue
func (t SystemObjectSetValue) GetRegistrationID()string{
	return t.RegistrationID
}

// Set RegistrationID of struct SystemObjectSetValue to value val
func (t *SystemObjectSetValue) SetRegistrationID(val string){
	t.RegistrationID = val
}

// Returns SegmentType of struct SystemObjectSetValue
func (t SystemObjectSetValue) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct SystemObjectSetValue to value val
func (t *SystemObjectSetValue) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns Object of struct SystemObjectSetValue
func (t SystemObjectSetValue) GetObject()string{
	return t.Object
}

// Set Object of struct SystemObjectSetValue to value val
func (t *SystemObjectSetValue) SetObject(val string){
	t.Object = val
}

// Returns Value of struct SystemObjectSetValue
func (t SystemObjectSetValue) GetValue()string{
	return t.Value
}

// Set Value of struct SystemObjectSetValue to value val
func (t *SystemObjectSetValue) SetValue(val string){
	t.Value = val
}

// Returns Text of struct SystemObjectSetValue
func (t SystemObjectSetValue) GetText()string{
	return t.Text
}

// Set Text of struct SystemObjectSetValue to value val
func (t *SystemObjectSetValue) SetText(val string){
	t.Text = val
}


type SystemObjectSetValueResponse struct{
	Dummy string  `xml:"dummy"`
	
}

// Returns Dummy of struct SystemObjectSetValueResponse
func (t SystemObjectSetValueResponse) GetDummy()string{
	return t.Dummy
}

// Set Dummy of struct SystemObjectSetValueResponse to value val
func (t *SystemObjectSetValueResponse) SetDummy(val string){
	t.Dummy = val
}


type GetMtesByRequestTable struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	IGetExtendedResult string  `xml:"iGetExtendedResult"`
	SoapRequest ArrayOfALREQUESTCUC  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct GetMtesByRequestTable
func (t GetMtesByRequestTable) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct GetMtesByRequestTable to value val
func (t *GetMtesByRequestTable) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct GetMtesByRequestTable
func (t GetMtesByRequestTable) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct GetMtesByRequestTable to value val
func (t *GetMtesByRequestTable) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns IGetExtendedResult of struct GetMtesByRequestTable
func (t GetMtesByRequestTable) GetIGetExtendedResult()string{
	return t.IGetExtendedResult
}

// Set IGetExtendedResult of struct GetMtesByRequestTable to value val
func (t *GetMtesByRequestTable) SetIGetExtendedResult(val string){
	t.IGetExtendedResult = val
}

// Returns SoapRequest of struct GetMtesByRequestTable
func (t GetMtesByRequestTable) GetSoapRequest()ArrayOfALREQUESTCUC{
	return t.SoapRequest
}

// Set SoapRequest of struct GetMtesByRequestTable to value val
func (t *GetMtesByRequestTable) SetSoapRequest(val ArrayOfALREQUESTCUC){
	t.SoapRequest = val
}


type GetMtesByRequestTableResponse struct{
	Result ArrayOfALRESULTCUC  `xml:"Result"`
	ResultLong ArrayOfALRSLTCUCL  `xml:"ResultLong"`
	
}

// Returns Result of struct GetMtesByRequestTableResponse
func (t GetMtesByRequestTableResponse) GetResult()ArrayOfALRESULTCUC{
	return t.Result
}

// Set Result of struct GetMtesByRequestTableResponse to value val
func (t *GetMtesByRequestTableResponse) SetResult(val ArrayOfALRESULTCUC){
	t.Result = val
}

// Returns ResultLong of struct GetMtesByRequestTableResponse
func (t GetMtesByRequestTableResponse) GetResultLong()ArrayOfALRSLTCUCL{
	return t.ResultLong
}

// Set ResultLong of struct GetMtesByRequestTableResponse to value val
func (t *GetMtesByRequestTableResponse) SetResultLong(val ArrayOfALRSLTCUCL){
	t.ResultLong = val
}


type GetMcInLocalMs struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	
}

// Returns CallingSID of struct GetMcInLocalMs
func (t GetMcInLocalMs) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct GetMcInLocalMs to value val
func (t *GetMcInLocalMs) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct GetMcInLocalMs
func (t GetMcInLocalMs) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct GetMcInLocalMs to value val
func (t *GetMcInLocalMs) SetSegmentType(val string){
	t.SegmentType = val
}


type GetMcInLocalMsResponse struct{
	ContextsV2 ArrayOfALMONCTX2  `xml:"ContextsV2"`
	ContextsV3 ArrayOfALMONCTX3  `xml:"ContextsV3"`
	
}

// Returns ContextsV2 of struct GetMcInLocalMsResponse
func (t GetMcInLocalMsResponse) GetContextsV2()ArrayOfALMONCTX2{
	return t.ContextsV2
}

// Set ContextsV2 of struct GetMcInLocalMsResponse to value val
func (t *GetMcInLocalMsResponse) SetContextsV2(val ArrayOfALMONCTX2){
	t.ContextsV2 = val
}

// Returns ContextsV3 of struct GetMcInLocalMsResponse
func (t GetMcInLocalMsResponse) GetContextsV3()ArrayOfALMONCTX3{
	return t.ContextsV3
}

// Set ContextsV3 of struct GetMcInLocalMsResponse to value val
func (t *GetMcInLocalMsResponse) SetContextsV3(val ArrayOfALMONCTX3){
	t.ContextsV3 = val
}


type MsGetMteclsInLocalMs struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	IMtgroup string  `xml:"iMtgroup"`
	
}

// Returns CallingSID of struct MsGetMteclsInLocalMs
func (t MsGetMteclsInLocalMs) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MsGetMteclsInLocalMs to value val
func (t *MsGetMteclsInLocalMs) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MsGetMteclsInLocalMs
func (t MsGetMteclsInLocalMs) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MsGetMteclsInLocalMs to value val
func (t *MsGetMteclsInLocalMs) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns IMtgroup of struct MsGetMteclsInLocalMs
func (t MsGetMteclsInLocalMs) GetIMtgroup()string{
	return t.IMtgroup
}

// Set IMtgroup of struct MsGetMteclsInLocalMs to value val
func (t *MsGetMteclsInLocalMs) SetIMtgroup(val string){
	t.IMtgroup = val
}


type MsGetMteclsInLocalMsResponse struct{
	MTEClassList ArrayOfALCUSTGRP  `xml:"MTEClassList"`
	
}

// Returns MTEClassList of struct MsGetMteclsInLocalMsResponse
func (t MsGetMteclsInLocalMsResponse) GetMTEClassList()ArrayOfALCUSTGRP{
	return t.MTEClassList
}

// Set MTEClassList of struct MsGetMteclsInLocalMsResponse to value val
func (t *MsGetMteclsInLocalMsResponse) SetMTEClassList(val ArrayOfALCUSTGRP){
	t.MTEClassList = val
}


type MsGetLocalMsInfo struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	
}

// Returns CallingSID of struct MsGetLocalMsInfo
func (t MsGetLocalMsInfo) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MsGetLocalMsInfo to value val
func (t *MsGetLocalMsInfo) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MsGetLocalMsInfo
func (t MsGetLocalMsInfo) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MsGetLocalMsInfo to value val
func (t *MsGetLocalMsInfo) SetSegmentType(val string){
	t.SegmentType = val
}


type MsGetLocalMsInfoResponse struct{
	SegmentInfo ArrayOfALMONSEG  `xml:"SegmentInfo"`
	
}

// Returns SegmentInfo of struct MsGetLocalMsInfoResponse
func (t MsGetLocalMsInfoResponse) GetSegmentInfo()ArrayOfALMONSEG{
	return t.SegmentInfo
}

// Set SegmentInfo of struct MsGetLocalMsInfoResponse to value val
func (t *MsGetLocalMsInfoResponse) SetSegmentInfo(val ArrayOfALMONSEG){
	t.SegmentInfo = val
}


type UtilMtGetTreeLocal struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SoapRequest ArrayOfALTIDMAXLV  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct UtilMtGetTreeLocal
func (t UtilMtGetTreeLocal) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct UtilMtGetTreeLocal to value val
func (t *UtilMtGetTreeLocal) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct UtilMtGetTreeLocal
func (t UtilMtGetTreeLocal) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct UtilMtGetTreeLocal to value val
func (t *UtilMtGetTreeLocal) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SoapRequest of struct UtilMtGetTreeLocal
func (t UtilMtGetTreeLocal) GetSoapRequest()ArrayOfALTIDMAXLV{
	return t.SoapRequest
}

// Set SoapRequest of struct UtilMtGetTreeLocal to value val
func (t *UtilMtGetTreeLocal) SetSoapRequest(val ArrayOfALTIDMAXLV){
	t.SoapRequest = val
}


type UtilMtGetTreeLocalResponse struct{
	TIDList ArrayOfALMTTREE  `xml:"TIDList"`
	RCList ArrayOfALGTIDRC  `xml:"RCList"`
	
}

// Returns TIDList of struct UtilMtGetTreeLocalResponse
func (t UtilMtGetTreeLocalResponse) GetTIDList()ArrayOfALMTTREE{
	return t.TIDList
}

// Set TIDList of struct UtilMtGetTreeLocalResponse to value val
func (t *UtilMtGetTreeLocalResponse) SetTIDList(val ArrayOfALMTTREE){
	t.TIDList = val
}

// Returns RCList of struct UtilMtGetTreeLocalResponse
func (t UtilMtGetTreeLocalResponse) GetRCList()ArrayOfALGTIDRC{
	return t.RCList
}

// Set RCList of struct UtilMtGetTreeLocalResponse to value val
func (t *UtilMtGetTreeLocalResponse) SetRCList(val ArrayOfALGTIDRC){
	t.RCList = val
}


type InfoGetTree struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	SytemId string  `xml:"SytemId"`
	ContextName string  `xml:"ContextName"`
	Object string  `xml:"Object"`
	AttributeName string  `xml:"AttributeName"`
	Tid ALGLOBTID  `xml:"Tid"`
	
}

// Returns CallingSID of struct InfoGetTree
func (t InfoGetTree) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct InfoGetTree to value val
func (t *InfoGetTree) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct InfoGetTree
func (t InfoGetTree) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct InfoGetTree to value val
func (t *InfoGetTree) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns SytemId of struct InfoGetTree
func (t InfoGetTree) GetSytemId()string{
	return t.SytemId
}

// Set SytemId of struct InfoGetTree to value val
func (t *InfoGetTree) SetSytemId(val string){
	t.SytemId = val
}

// Returns ContextName of struct InfoGetTree
func (t InfoGetTree) GetContextName()string{
	return t.ContextName
}

// Set ContextName of struct InfoGetTree to value val
func (t *InfoGetTree) SetContextName(val string){
	t.ContextName = val
}

// Returns Object of struct InfoGetTree
func (t InfoGetTree) GetObject()string{
	return t.Object
}

// Set Object of struct InfoGetTree to value val
func (t *InfoGetTree) SetObject(val string){
	t.Object = val
}

// Returns AttributeName of struct InfoGetTree
func (t InfoGetTree) GetAttributeName()string{
	return t.AttributeName
}

// Set AttributeName of struct InfoGetTree to value val
func (t *InfoGetTree) SetAttributeName(val string){
	t.AttributeName = val
}

// Returns Tid of struct InfoGetTree
func (t InfoGetTree) GetTid()ALGLOBTID{
	return t.Tid
}

// Set Tid of struct InfoGetTree to value val
func (t *InfoGetTree) SetTid(val ALGLOBTID){
	t.Tid = val
}


type InfoGetTreeResponse struct{
	TreeInfo ArrayOfALMTTREE  `xml:"TreeInfo"`
	PerfInfo ArrayOfALPERFTYPE  `xml:"PerfInfo"`
	SmesInfo ArrayOfALSMSGRAWT  `xml:"SmesInfo"`
	LinkedTxtAttr ArrayOfALTXTLNKAT  `xml:"LinkedTxtAttr"`
	
}

// Returns TreeInfo of struct InfoGetTreeResponse
func (t InfoGetTreeResponse) GetTreeInfo()ArrayOfALMTTREE{
	return t.TreeInfo
}

// Set TreeInfo of struct InfoGetTreeResponse to value val
func (t *InfoGetTreeResponse) SetTreeInfo(val ArrayOfALMTTREE){
	t.TreeInfo = val
}

// Returns PerfInfo of struct InfoGetTreeResponse
func (t InfoGetTreeResponse) GetPerfInfo()ArrayOfALPERFTYPE{
	return t.PerfInfo
}

// Set PerfInfo of struct InfoGetTreeResponse to value val
func (t *InfoGetTreeResponse) SetPerfInfo(val ArrayOfALPERFTYPE){
	t.PerfInfo = val
}

// Returns SmesInfo of struct InfoGetTreeResponse
func (t InfoGetTreeResponse) GetSmesInfo()ArrayOfALSMSGRAWT{
	return t.SmesInfo
}

// Set SmesInfo of struct InfoGetTreeResponse to value val
func (t *InfoGetTreeResponse) SetSmesInfo(val ArrayOfALSMSGRAWT){
	t.SmesInfo = val
}

// Returns LinkedTxtAttr of struct InfoGetTreeResponse
func (t InfoGetTreeResponse) GetLinkedTxtAttr()ArrayOfALTXTLNKAT{
	return t.LinkedTxtAttr
}

// Set LinkedTxtAttr of struct InfoGetTreeResponse to value val
func (t *InfoGetTreeResponse) SetLinkedTxtAttr(val ArrayOfALTXTLNKAT){
	t.LinkedTxtAttr = val
}


type GetMtListByMtclass struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	IGetMembers string  `xml:"iGetMembers"`
	IGetUsingForCus string  `xml:"iGetUsingForCus"`
	IGetUsingForTools string  `xml:"iGetUsingForTools"`
	ClassName string  `xml:"ClassName"`
	SoapRequest ALGLOBTID  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct GetMtListByMtclass
func (t GetMtListByMtclass) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct GetMtListByMtclass to value val
func (t *GetMtListByMtclass) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct GetMtListByMtclass
func (t GetMtListByMtclass) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct GetMtListByMtclass to value val
func (t *GetMtListByMtclass) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns IGetMembers of struct GetMtListByMtclass
func (t GetMtListByMtclass) GetIGetMembers()string{
	return t.IGetMembers
}

// Set IGetMembers of struct GetMtListByMtclass to value val
func (t *GetMtListByMtclass) SetIGetMembers(val string){
	t.IGetMembers = val
}

// Returns IGetUsingForCus of struct GetMtListByMtclass
func (t GetMtListByMtclass) GetIGetUsingForCus()string{
	return t.IGetUsingForCus
}

// Set IGetUsingForCus of struct GetMtListByMtclass to value val
func (t *GetMtListByMtclass) SetIGetUsingForCus(val string){
	t.IGetUsingForCus = val
}

// Returns IGetUsingForTools of struct GetMtListByMtclass
func (t GetMtListByMtclass) GetIGetUsingForTools()string{
	return t.IGetUsingForTools
}

// Set IGetUsingForTools of struct GetMtListByMtclass to value val
func (t *GetMtListByMtclass) SetIGetUsingForTools(val string){
	t.IGetUsingForTools = val
}

// Returns ClassName of struct GetMtListByMtclass
func (t GetMtListByMtclass) GetClassName()string{
	return t.ClassName
}

// Set ClassName of struct GetMtListByMtclass to value val
func (t *GetMtListByMtclass) SetClassName(val string){
	t.ClassName = val
}

// Returns SoapRequest of struct GetMtListByMtclass
func (t GetMtListByMtclass) GetSoapRequest()ALGLOBTID{
	return t.SoapRequest
}

// Set SoapRequest of struct GetMtListByMtclass to value val
func (t *GetMtListByMtclass) SetSoapRequest(val ALGLOBTID){
	t.SoapRequest = val
}


type GetMtListByMtclassResponse struct{
	MTEsFoundForClass ArrayOfALGLOBTID  `xml:"MTEsFoundForClass"`
	MTEsWithGeneralPropFromClass ArrayOfALGLOBTID  `xml:"MTEsWithGeneralPropFromClass"`
	MTEsWithDataCollFromClass ArrayOfALGLOBTID  `xml:"MTEsWithDataCollFromClass"`
	MTEsWithAutoreactionFromClass ArrayOfALGLOBTID  `xml:"MTEsWithAutoreactionFromClass"`
	MTEsWithAnalysisMethodFromClass ArrayOfALGLOBTID  `xml:"MTEsWithAnalysisMethodFromClass"`
	
}

// Returns MTEsFoundForClass of struct GetMtListByMtclassResponse
func (t GetMtListByMtclassResponse) GetMTEsFoundForClass()ArrayOfALGLOBTID{
	return t.MTEsFoundForClass
}

// Set MTEsFoundForClass of struct GetMtListByMtclassResponse to value val
func (t *GetMtListByMtclassResponse) SetMTEsFoundForClass(val ArrayOfALGLOBTID){
	t.MTEsFoundForClass = val
}

// Returns MTEsWithGeneralPropFromClass of struct GetMtListByMtclassResponse
func (t GetMtListByMtclassResponse) GetMTEsWithGeneralPropFromClass()ArrayOfALGLOBTID{
	return t.MTEsWithGeneralPropFromClass
}

// Set MTEsWithGeneralPropFromClass of struct GetMtListByMtclassResponse to value val
func (t *GetMtListByMtclassResponse) SetMTEsWithGeneralPropFromClass(val ArrayOfALGLOBTID){
	t.MTEsWithGeneralPropFromClass = val
}

// Returns MTEsWithDataCollFromClass of struct GetMtListByMtclassResponse
func (t GetMtListByMtclassResponse) GetMTEsWithDataCollFromClass()ArrayOfALGLOBTID{
	return t.MTEsWithDataCollFromClass
}

// Set MTEsWithDataCollFromClass of struct GetMtListByMtclassResponse to value val
func (t *GetMtListByMtclassResponse) SetMTEsWithDataCollFromClass(val ArrayOfALGLOBTID){
	t.MTEsWithDataCollFromClass = val
}

// Returns MTEsWithAutoreactionFromClass of struct GetMtListByMtclassResponse
func (t GetMtListByMtclassResponse) GetMTEsWithAutoreactionFromClass()ArrayOfALGLOBTID{
	return t.MTEsWithAutoreactionFromClass
}

// Set MTEsWithAutoreactionFromClass of struct GetMtListByMtclassResponse to value val
func (t *GetMtListByMtclassResponse) SetMTEsWithAutoreactionFromClass(val ArrayOfALGLOBTID){
	t.MTEsWithAutoreactionFromClass = val
}

// Returns MTEsWithAnalysisMethodFromClass of struct GetMtListByMtclassResponse
func (t GetMtListByMtclassResponse) GetMTEsWithAnalysisMethodFromClass()ArrayOfALGLOBTID{
	return t.MTEsWithAnalysisMethodFromClass
}

// Set MTEsWithAnalysisMethodFromClass of struct GetMtListByMtclassResponse to value val
func (t *GetMtListByMtclassResponse) SetMTEsWithAnalysisMethodFromClass(val ArrayOfALGLOBTID){
	t.MTEsWithAnalysisMethodFromClass = val
}


type MteGetByToolRunstatus struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	ILookAheadSeconds string  `xml:"iLookAheadSeconds"`
	IRunTimeStatus string  `xml:"iRunTimeStatus"`
	SoapRequest ArrayOfALTLDPNM  `xml:"SoapRequest"`
	
}

// Returns CallingSID of struct MteGetByToolRunstatus
func (t MteGetByToolRunstatus) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MteGetByToolRunstatus to value val
func (t *MteGetByToolRunstatus) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MteGetByToolRunstatus
func (t MteGetByToolRunstatus) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MteGetByToolRunstatus to value val
func (t *MteGetByToolRunstatus) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns ILookAheadSeconds of struct MteGetByToolRunstatus
func (t MteGetByToolRunstatus) GetILookAheadSeconds()string{
	return t.ILookAheadSeconds
}

// Set ILookAheadSeconds of struct MteGetByToolRunstatus to value val
func (t *MteGetByToolRunstatus) SetILookAheadSeconds(val string){
	t.ILookAheadSeconds = val
}

// Returns IRunTimeStatus of struct MteGetByToolRunstatus
func (t MteGetByToolRunstatus) GetIRunTimeStatus()string{
	return t.IRunTimeStatus
}

// Set IRunTimeStatus of struct MteGetByToolRunstatus to value val
func (t *MteGetByToolRunstatus) SetIRunTimeStatus(val string){
	t.IRunTimeStatus = val
}

// Returns SoapRequest of struct MteGetByToolRunstatus
func (t MteGetByToolRunstatus) GetSoapRequest()ArrayOfALTLDPNM{
	return t.SoapRequest
}

// Set SoapRequest of struct MteGetByToolRunstatus to value val
func (t *MteGetByToolRunstatus) SetSoapRequest(val ArrayOfALTLDPNM){
	t.SoapRequest = val
}


type MteGetByToolRunstatusResponse struct{
	DispatcherRc ArrayOfALTOOLDPRC  `xml:"DispatcherRc"`
	ToolControl ArrayOfALTOOLCTRL  `xml:"ToolControl"`
	
}

// Returns DispatcherRc of struct MteGetByToolRunstatusResponse
func (t MteGetByToolRunstatusResponse) GetDispatcherRc()ArrayOfALTOOLDPRC{
	return t.DispatcherRc
}

// Set DispatcherRc of struct MteGetByToolRunstatusResponse to value val
func (t *MteGetByToolRunstatusResponse) SetDispatcherRc(val ArrayOfALTOOLDPRC){
	t.DispatcherRc = val
}

// Returns ToolControl of struct MteGetByToolRunstatusResponse
func (t MteGetByToolRunstatusResponse) GetToolControl()ArrayOfALTOOLCTRL{
	return t.ToolControl
}

// Set ToolControl of struct MteGetByToolRunstatusResponse to value val
func (t *MteGetByToolRunstatusResponse) SetToolControl(val ArrayOfALTOOLCTRL){
	t.ToolControl = val
}


type MtGetAllToolsToSet struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	
}

// Returns CallingSID of struct MtGetAllToolsToSet
func (t MtGetAllToolsToSet) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtGetAllToolsToSet to value val
func (t *MtGetAllToolsToSet) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MtGetAllToolsToSet
func (t MtGetAllToolsToSet) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtGetAllToolsToSet to value val
func (t *MtGetAllToolsToSet) SetSegmentType(val string){
	t.SegmentType = val
}


type MtGetAllToolsToSetResponse struct{
	TidList ArrayOfALGLOBTID  `xml:"TidList"`
	
}

// Returns TidList of struct MtGetAllToolsToSetResponse
func (t MtGetAllToolsToSetResponse) GetTidList()ArrayOfALGLOBTID{
	return t.TidList
}

// Set TidList of struct MtGetAllToolsToSetResponse to value val
func (t *MtGetAllToolsToSetResponse) SetTidList(val ArrayOfALGLOBTID){
	t.TidList = val
}


type MtGetAllToCust struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	
}

// Returns CallingSID of struct MtGetAllToCust
func (t MtGetAllToCust) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct MtGetAllToCust to value val
func (t *MtGetAllToCust) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct MtGetAllToCust
func (t MtGetAllToCust) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct MtGetAllToCust to value val
func (t *MtGetAllToCust) SetSegmentType(val string){
	t.SegmentType = val
}


type MtGetAllToCustResponse struct{
	GeneralProperties ArrayOfALGLOBTID  `xml:"GeneralProperties"`
	PerformanceProperties ArrayOfALGLOBTID  `xml:"PerformanceProperties"`
	StatusMessageProperties ArrayOfALGLOBTID  `xml:"StatusMessageProperties"`
	MessageContainerProperties ArrayOfALGLOBTID  `xml:"MessageContainerProperties"`
	
}

// Returns GeneralProperties of struct MtGetAllToCustResponse
func (t MtGetAllToCustResponse) GetGeneralProperties()ArrayOfALGLOBTID{
	return t.GeneralProperties
}

// Set GeneralProperties of struct MtGetAllToCustResponse to value val
func (t *MtGetAllToCustResponse) SetGeneralProperties(val ArrayOfALGLOBTID){
	t.GeneralProperties = val
}

// Returns PerformanceProperties of struct MtGetAllToCustResponse
func (t MtGetAllToCustResponse) GetPerformanceProperties()ArrayOfALGLOBTID{
	return t.PerformanceProperties
}

// Set PerformanceProperties of struct MtGetAllToCustResponse to value val
func (t *MtGetAllToCustResponse) SetPerformanceProperties(val ArrayOfALGLOBTID){
	t.PerformanceProperties = val
}

// Returns StatusMessageProperties of struct MtGetAllToCustResponse
func (t MtGetAllToCustResponse) GetStatusMessageProperties()ArrayOfALGLOBTID{
	return t.StatusMessageProperties
}

// Set StatusMessageProperties of struct MtGetAllToCustResponse to value val
func (t *MtGetAllToCustResponse) SetStatusMessageProperties(val ArrayOfALGLOBTID){
	t.StatusMessageProperties = val
}

// Returns MessageContainerProperties of struct MtGetAllToCustResponse
func (t MtGetAllToCustResponse) GetMessageContainerProperties()ArrayOfALGLOBTID{
	return t.MessageContainerProperties
}

// Set MessageContainerProperties of struct MtGetAllToCustResponse to value val
func (t *MtGetAllToCustResponse) SetMessageContainerProperties(val ArrayOfALGLOBTID){
	t.MessageContainerProperties = val
}


type GetListOfMaByCusGrp struct{
	CallingSID string  `xml:"CallingSID"`
	SegmentType string  `xml:"SegmentType"`
	AttrGroupName string  `xml:"AttrGroupName"`
	
}

// Returns CallingSID of struct GetListOfMaByCusGrp
func (t GetListOfMaByCusGrp) GetCallingSID()string{
	return t.CallingSID
}

// Set CallingSID of struct GetListOfMaByCusGrp to value val
func (t *GetListOfMaByCusGrp) SetCallingSID(val string){
	t.CallingSID = val
}

// Returns SegmentType of struct GetListOfMaByCusGrp
func (t GetListOfMaByCusGrp) GetSegmentType()string{
	return t.SegmentType
}

// Set SegmentType of struct GetListOfMaByCusGrp to value val
func (t *GetListOfMaByCusGrp) SetSegmentType(val string){
	t.SegmentType = val
}

// Returns AttrGroupName of struct GetListOfMaByCusGrp
func (t GetListOfMaByCusGrp) GetAttrGroupName()string{
	return t.AttrGroupName
}

// Set AttrGroupName of struct GetListOfMaByCusGrp to value val
func (t *GetListOfMaByCusGrp) SetAttrGroupName(val string){
	t.AttrGroupName = val
}


type GetListOfMaByCusGrpResponse struct{
	GeneralProperties ArrayOfALGLOBTID  `xml:"GeneralProperties"`
	PerformanceProperties ArrayOfALGLOBTID  `xml:"PerformanceProperties"`
	StatusMessageProperties ArrayOfALGLOBTID  `xml:"StatusMessageProperties"`
	MessageContainerProperties ArrayOfALGLOBTID  `xml:"MessageContainerProperties"`
	
}

// Returns GeneralProperties of struct GetListOfMaByCusGrpResponse
func (t GetListOfMaByCusGrpResponse) GetGeneralProperties()ArrayOfALGLOBTID{
	return t.GeneralProperties
}

// Set GeneralProperties of struct GetListOfMaByCusGrpResponse to value val
func (t *GetListOfMaByCusGrpResponse) SetGeneralProperties(val ArrayOfALGLOBTID){
	t.GeneralProperties = val
}

// Returns PerformanceProperties of struct GetListOfMaByCusGrpResponse
func (t GetListOfMaByCusGrpResponse) GetPerformanceProperties()ArrayOfALGLOBTID{
	return t.PerformanceProperties
}

// Set PerformanceProperties of struct GetListOfMaByCusGrpResponse to value val
func (t *GetListOfMaByCusGrpResponse) SetPerformanceProperties(val ArrayOfALGLOBTID){
	t.PerformanceProperties = val
}

// Returns StatusMessageProperties of struct GetListOfMaByCusGrpResponse
func (t GetListOfMaByCusGrpResponse) GetStatusMessageProperties()ArrayOfALGLOBTID{
	return t.StatusMessageProperties
}

// Set StatusMessageProperties of struct GetListOfMaByCusGrpResponse to value val
func (t *GetListOfMaByCusGrpResponse) SetStatusMessageProperties(val ArrayOfALGLOBTID){
	t.StatusMessageProperties = val
}

// Returns MessageContainerProperties of struct GetListOfMaByCusGrpResponse
func (t GetListOfMaByCusGrpResponse) GetMessageContainerProperties()ArrayOfALGLOBTID{
	return t.MessageContainerProperties
}

// Set MessageContainerProperties of struct GetListOfMaByCusGrpResponse to value val
func (t *GetListOfMaByCusGrpResponse) SetMessageContainerProperties(val ArrayOfALGLOBTID){
	t.MessageContainerProperties = val
}


type RequestLogonFile struct{
	User string  `xml:"user"`
	
}

// Returns User of struct RequestLogonFile
func (t RequestLogonFile) GetUser()string{
	return t.User
}

// Set User of struct RequestLogonFile to value val
func (t *RequestLogonFile) SetUser(val string){
	t.User = val
}


type RequestLogonFileResponse struct{
	Filename string  `xml:"filename"`
	
}

// Returns Filename of struct RequestLogonFileResponse
func (t RequestLogonFileResponse) GetFilename()string{
	return t.Filename
}

// Set Filename of struct RequestLogonFileResponse to value val
func (t *RequestLogonFileResponse) SetFilename(val string){
	t.Filename = val
}



// input message name is Name of Operation
// output message name is Name of a Message

//type SAPCCMS interface {
//	
//	ReadFile(inMessage *MsgReadFileRequest, outMessage *MsgReadFileResponse) error
//	
//	ReadDirectory(inMessage *MsgReadDirectoryRequest, outMessage *MsgReadDirectoryResponse) error
//	
//	ReadProfileParameters(inMessage *MsgReadProfileParametersRequest, outMessage *MsgReadProfileParametersResponse) error
//	
//	Register(inMessage *MsgRegisterRequest, outMessage *MsgRegisterResponse) error
//	
//	Unregister(inMessage *MsgUnregisterRequest, outMessage *MsgUnregisterResponse) error
//	
//	GetAgentConfig(inMessage *MsgGetAgentConfigRequest, outMessage *MsgAgentConfig) error
//	
//	MtGetTidByName(inMessage *MsgMtGetTidByNameRequest, outMessage *MsgMtGetTidByNameResponse) error
//	
//	MtGetMteinfo(inMessage *MsgMtGetMteinfoRequest, outMessage *MsgMtGetMteinfoResponse) error
//	
//	MtRead(inMessage *MsgMtReadRequest, outMessage *MsgMtReadResponse) error
//	
//	PerfRead(inMessage *MsgPerfReadRequest, outMessage *MsgPerfReadResponse) error
//	
//	PerfReadSmoothData(inMessage *MsgPerfReadSmoothDataRequest, outMessage *MsgPerfReadSmoothDataResponse) error
//	
//	MscReadCache(inMessage *MsgMscReadCacheRequest, outMessage *MsgMscReadCacheResponse) error
//	
//	UtilSnglmsgReadRawdata(inMessage *MsgUtilSnglmsgReadRawdataRequest, outMessage *MsgUtilSnglmsgReadRawdataResponse) error
//	
//	UtilMtReadAll(inMessage *MsgUtilMtReadAllRequest, outMessage *MsgUtilMtReadAllResponse) error
//	
//	TextAttrRead(inMessage *MsgTextAttrReadRequest, outMessage *MsgTextAttrReadResponse) error
//	
//	ReferenceRead(inMessage *MsgReferenceReadRequest, outMessage *MsgReferenceReadResponse) error
//	
//	UtilMtGetAidByTid(inMessage *MsgUtilMtGetAidByTidRequest, outMessage *MsgUtilMtGetAidByTidResponse) error
//	
//	UtilReadRawalertByAid(inMessage *MsgUtilReadRawalertByAidRequest, outMessage *MsgUtilReadRawalertByAidResponse) error
//	
//	ToolGetEffective(inMessage *MsgToolGetEffectiveRequest, outMessage *MsgToolGetEffectiveResponse) error
//	
//	TriggerDataCollection(inMessage *MsgTriggerDataCollectionRequest, outMessage *MsgTriggerDataCollectionResponse) error
//	
//	MtChangeStatus(inMessage *MsgMtChangeStatusRequest, outMessage *MsgMtChangeStatusResponse) error
//	
//	MtCustomizeWrite(inMessage *MsgMtCustomizeWriteRequest, outMessage *MsgMtCustomizeWriteResponse) error
//	
//	MtDestroyMarkNTry(inMessage *MsgMtDestroyMarkNTryRequest, outMessage *MsgMtDestroyMarkNTryResponse) error
//	
//	MtReset(inMessage *MsgMtResetRequest, outMessage *MsgMtResetResponse) error
//	
//	PerfCustomizeWrite(inMessage *MsgPerfCustomizeWriteRequest, outMessage *MsgPerfCustomizeWriteResponse) error
//	
//	MscCustomizeWrite(inMessage *MsgMscCustomizeWriteRequest, outMessage *MsgMscCustomizeWriteResponse) error
//	
//	MscDeleteLines(inMessage *MsgMscDeleteLinesRequest, outMessage *MsgMscDeleteLinesResponse) error
//	
//	SnglmgsCustomizeWrite(inMessage *MsgSnglmgsCustomizeWriteRequest, outMessage *MsgSnglmgsCustomizeWriteResponse) error
//	
//	ToolSet(inMessage *MsgToolSetRequest, outMessage *MsgToolSetResponse) error
//	
//	ToolSetRuntimeStatus(inMessage *MsgToolSetRuntimeStatusRequest, outMessage *MsgToolSetRuntimeStatusResponse) error
//	
//	UtilAlChangeStatus(inMessage *MsgUtilAlChangeStatusRequest, outMessage *MsgUtilAlChangeStatusResponse) error
//	
//	MtDbsetToWpsetByTid(inMessage *MsgMtDbsetToWpsetByTidRequest, outMessage *MsgMtDbsetToWpsetByTidResponse) error
//	
//	SystemObjectSetValue(inMessage *MsgSystemObjectSetValueRequest, outMessage *MsgSystemObjectSetValueResponse) error
//	
//	GetMtesByRequestTable(inMessage *MsgGetMtesByRequestTableRequest, outMessage *MsgGetMtesByRequestTableResponse) error
//	
//	GetMcInLocalMs(inMessage *MsgGetMcInLocalMsRequest, outMessage *MsgGetMcInLocalMsResponse) error
//	
//	MsGetMteclsInLocalMs(inMessage *MsgMsGetMteclsInLocalMsRequest, outMessage *MsgMsGetMteclsInLocalMsResponse) error
//	
//	MsGetLocalMsInfo(inMessage *MsgMsGetLocalMsInfoRequest, outMessage *MsgMsGetLocalMsInfoResponse) error
//	
//	UtilMtGetTreeLocal(inMessage *MsgUtilMtGetTreeLocalRequest, outMessage *MsgUtilMtGetTreeLocalResponse) error
//	
//	InfoGetTree(inMessage *MsgInfoGetTreeRequest, outMessage *MsgInfoGetTreeResponse) error
//	
//	GetMtListByMtclass(inMessage *MsgGetMtListByMtclassRequest, outMessage *MsgGetMtListByMtclassResponse) error
//	
//	MteGetByToolRunstatus(inMessage *MsgMteGetByToolRunstatusRequest, outMessage *MsgMteGetByToolRunstatusResponse) error
//	
//	MtGetAllToolsToSet(inMessage *MsgMtGetAllToolsToSetRequest, outMessage *MsgMtGetAllToolsToSetResponse) error
//	
//	MtGetAllToCust(inMessage *MsgMtGetAllToCustRequest, outMessage *MsgMtGetAllToCustResponse) error
//	
//	GetListOfMaByCusGrp(inMessage *MsgGetListOfMaByCusGrpRequest, outMessage *MsgGetListOfMaByCusGrpResponse) error
//	
//	RequestLogonFile(inMessage *MsgRequestLogonFileRequest, outMessage *MsgRequestLogonFileResponse) error
//	
//}


func (s *SAPCCMS) ReadFile(inMs *MsgReadFileRequest, outMs *MsgReadFileResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgReadFileResponse)
	return nil
}

func (s *SAPCCMS) ReadDirectory(inMs *MsgReadDirectoryRequest, outMs *MsgReadDirectoryResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgReadDirectoryResponse)
	return nil
}

func (s *SAPCCMS) ReadProfileParameters(inMs *MsgReadProfileParametersRequest, outMs *MsgReadProfileParametersResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgReadProfileParametersResponse)
	return nil
}

func (s *SAPCCMS) Register(inMs *MsgRegisterRequest, outMs *MsgRegisterResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgRegisterResponse)
	return nil
}

func (s *SAPCCMS) Unregister(inMs *MsgUnregisterRequest, outMs *MsgUnregisterResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgUnregisterResponse)
	return nil
}

func (s *SAPCCMS) GetAgentConfig(inMs *MsgGetAgentConfigRequest, outMs *MsgAgentConfig) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgAgentConfig)
	return nil
}

func (s *SAPCCMS) MtGetTidByName(inMs *MsgMtGetTidByNameRequest, outMs *MsgMtGetTidByNameResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtGetTidByNameResponse)
	return nil
}

func (s *SAPCCMS) MtGetMteinfo(inMs *MsgMtGetMteinfoRequest, outMs *MsgMtGetMteinfoResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtGetMteinfoResponse)
	return nil
}

func (s *SAPCCMS) MtRead(inMs *MsgMtReadRequest, outMs *MsgMtReadResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtReadResponse)
	return nil
}

func (s *SAPCCMS) PerfRead(inMs *MsgPerfReadRequest, outMs *MsgPerfReadResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgPerfReadResponse)
	return nil
}

func (s *SAPCCMS) PerfReadSmoothData(inMs *MsgPerfReadSmoothDataRequest, outMs *MsgPerfReadSmoothDataResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgPerfReadSmoothDataResponse)
	return nil
}

func (s *SAPCCMS) MscReadCache(inMs *MsgMscReadCacheRequest, outMs *MsgMscReadCacheResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMscReadCacheResponse)
	return nil
}

func (s *SAPCCMS) UtilSnglmsgReadRawdata(inMs *MsgUtilSnglmsgReadRawdataRequest, outMs *MsgUtilSnglmsgReadRawdataResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgUtilSnglmsgReadRawdataResponse)
	return nil
}

func (s *SAPCCMS) UtilMtReadAll(inMs *MsgUtilMtReadAllRequest, outMs *MsgUtilMtReadAllResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgUtilMtReadAllResponse)
	return nil
}

func (s *SAPCCMS) TextAttrRead(inMs *MsgTextAttrReadRequest, outMs *MsgTextAttrReadResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgTextAttrReadResponse)
	return nil
}

func (s *SAPCCMS) ReferenceRead(inMs *MsgReferenceReadRequest, outMs *MsgReferenceReadResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgReferenceReadResponse)
	return nil
}

func (s *SAPCCMS) UtilMtGetAidByTid(inMs *MsgUtilMtGetAidByTidRequest, outMs *MsgUtilMtGetAidByTidResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgUtilMtGetAidByTidResponse)
	return nil
}

func (s *SAPCCMS) UtilReadRawalertByAid(inMs *MsgUtilReadRawalertByAidRequest, outMs *MsgUtilReadRawalertByAidResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgUtilReadRawalertByAidResponse)
	return nil
}

func (s *SAPCCMS) ToolGetEffective(inMs *MsgToolGetEffectiveRequest, outMs *MsgToolGetEffectiveResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgToolGetEffectiveResponse)
	return nil
}

func (s *SAPCCMS) TriggerDataCollection(inMs *MsgTriggerDataCollectionRequest, outMs *MsgTriggerDataCollectionResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgTriggerDataCollectionResponse)
	return nil
}

func (s *SAPCCMS) MtChangeStatus(inMs *MsgMtChangeStatusRequest, outMs *MsgMtChangeStatusResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtChangeStatusResponse)
	return nil
}

func (s *SAPCCMS) MtCustomizeWrite(inMs *MsgMtCustomizeWriteRequest, outMs *MsgMtCustomizeWriteResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtCustomizeWriteResponse)
	return nil
}

func (s *SAPCCMS) MtDestroyMarkNTry(inMs *MsgMtDestroyMarkNTryRequest, outMs *MsgMtDestroyMarkNTryResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtDestroyMarkNTryResponse)
	return nil
}

func (s *SAPCCMS) MtReset(inMs *MsgMtResetRequest, outMs *MsgMtResetResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtResetResponse)
	return nil
}

func (s *SAPCCMS) PerfCustomizeWrite(inMs *MsgPerfCustomizeWriteRequest, outMs *MsgPerfCustomizeWriteResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgPerfCustomizeWriteResponse)
	return nil
}

func (s *SAPCCMS) MscCustomizeWrite(inMs *MsgMscCustomizeWriteRequest, outMs *MsgMscCustomizeWriteResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMscCustomizeWriteResponse)
	return nil
}

func (s *SAPCCMS) MscDeleteLines(inMs *MsgMscDeleteLinesRequest, outMs *MsgMscDeleteLinesResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMscDeleteLinesResponse)
	return nil
}

func (s *SAPCCMS) SnglmgsCustomizeWrite(inMs *MsgSnglmgsCustomizeWriteRequest, outMs *MsgSnglmgsCustomizeWriteResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgSnglmgsCustomizeWriteResponse)
	return nil
}

func (s *SAPCCMS) ToolSet(inMs *MsgToolSetRequest, outMs *MsgToolSetResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgToolSetResponse)
	return nil
}

func (s *SAPCCMS) ToolSetRuntimeStatus(inMs *MsgToolSetRuntimeStatusRequest, outMs *MsgToolSetRuntimeStatusResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgToolSetRuntimeStatusResponse)
	return nil
}

func (s *SAPCCMS) UtilAlChangeStatus(inMs *MsgUtilAlChangeStatusRequest, outMs *MsgUtilAlChangeStatusResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgUtilAlChangeStatusResponse)
	return nil
}

func (s *SAPCCMS) MtDbsetToWpsetByTid(inMs *MsgMtDbsetToWpsetByTidRequest, outMs *MsgMtDbsetToWpsetByTidResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtDbsetToWpsetByTidResponse)
	return nil
}

func (s *SAPCCMS) SystemObjectSetValue(inMs *MsgSystemObjectSetValueRequest, outMs *MsgSystemObjectSetValueResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgSystemObjectSetValueResponse)
	return nil
}

func (s *SAPCCMS) GetMtesByRequestTable(inMs *MsgGetMtesByRequestTableRequest, outMs *MsgGetMtesByRequestTableResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgGetMtesByRequestTableResponse)
	return nil
}

func (s *SAPCCMS) GetMcInLocalMs(inMs *MsgGetMcInLocalMsRequest, outMs *MsgGetMcInLocalMsResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgGetMcInLocalMsResponse)
	return nil
}

func (s *SAPCCMS) MsGetMteclsInLocalMs(inMs *MsgMsGetMteclsInLocalMsRequest, outMs *MsgMsGetMteclsInLocalMsResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMsGetMteclsInLocalMsResponse)
	return nil
}

func (s *SAPCCMS) MsGetLocalMsInfo(inMs *MsgMsGetLocalMsInfoRequest, outMs *MsgMsGetLocalMsInfoResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMsGetLocalMsInfoResponse)
	return nil
}

func (s *SAPCCMS) UtilMtGetTreeLocal(inMs *MsgUtilMtGetTreeLocalRequest, outMs *MsgUtilMtGetTreeLocalResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgUtilMtGetTreeLocalResponse)
	return nil
}

func (s *SAPCCMS) InfoGetTree(inMs *MsgInfoGetTreeRequest, outMs *MsgInfoGetTreeResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgInfoGetTreeResponse)
	return nil
}

func (s *SAPCCMS) GetMtListByMtclass(inMs *MsgGetMtListByMtclassRequest, outMs *MsgGetMtListByMtclassResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgGetMtListByMtclassResponse)
	return nil
}

func (s *SAPCCMS) MteGetByToolRunstatus(inMs *MsgMteGetByToolRunstatusRequest, outMs *MsgMteGetByToolRunstatusResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMteGetByToolRunstatusResponse)
	return nil
}

func (s *SAPCCMS) MtGetAllToolsToSet(inMs *MsgMtGetAllToolsToSetRequest, outMs *MsgMtGetAllToolsToSetResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtGetAllToolsToSetResponse)
	return nil
}

func (s *SAPCCMS) MtGetAllToCust(inMs *MsgMtGetAllToCustRequest, outMs *MsgMtGetAllToCustResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgMtGetAllToCustResponse)
	return nil
}

func (s *SAPCCMS) GetListOfMaByCusGrp(inMs *MsgGetListOfMaByCusGrpRequest, outMs *MsgGetListOfMaByCusGrpResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgGetListOfMaByCusGrpResponse)
	return nil
}

func (s *SAPCCMS) RequestLogonFile(inMs *MsgRequestLogonFileRequest, outMs *MsgRequestLogonFileResponse) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*MsgRequestLogonFileResponse)
	return nil
}




type MsgReadFileRequest struct {
	XMLName xml.Name
	ReadFile `xml:"ReadFileRequest"`
}

func NewMsgReadFileRequest() *MsgReadFileRequest {
	var msg MsgReadFileRequest
	msg.XMLName.Local = "ReadFile"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgReadFileResponse struct {
	XMLName xml.Name
	ReadFileResponse `xml:"ReadFileResponse"`
}

func NewMsgReadFileResponse() *MsgReadFileResponse {
	var msg MsgReadFileResponse
	msg.XMLName.Local = "ReadFileResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgReadDirectoryRequest struct {
	XMLName xml.Name
	ReadDirectory `xml:"ReadDirectoryRequest"`
}

func NewMsgReadDirectoryRequest() *MsgReadDirectoryRequest {
	var msg MsgReadDirectoryRequest
	msg.XMLName.Local = "ReadDirectory"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgReadDirectoryResponse struct {
	XMLName xml.Name
	ReadDirectoryResponse `xml:"ReadDirectoryResponse"`
}

func NewMsgReadDirectoryResponse() *MsgReadDirectoryResponse {
	var msg MsgReadDirectoryResponse
	msg.XMLName.Local = "ReadDirectoryResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgReadProfileParametersRequest struct {
	XMLName xml.Name
	ReadProfileParameters `xml:"ReadProfileParametersRequest"`
}

func NewMsgReadProfileParametersRequest() *MsgReadProfileParametersRequest {
	var msg MsgReadProfileParametersRequest
	msg.XMLName.Local = "ReadProfileParameters"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgReadProfileParametersResponse struct {
	XMLName xml.Name
	ReadProfileParametersResponse `xml:"ReadProfileParametersResponse"`
}

func NewMsgReadProfileParametersResponse() *MsgReadProfileParametersResponse {
	var msg MsgReadProfileParametersResponse
	msg.XMLName.Local = "ReadProfileParametersResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgRegisterRequest struct {
	XMLName xml.Name
	Register `xml:"RegisterRequest"`
}

func NewMsgRegisterRequest() *MsgRegisterRequest {
	var msg MsgRegisterRequest
	msg.XMLName.Local = "Register"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgRegisterResponse struct {
	XMLName xml.Name
	RegisterResponse `xml:"RegisterResponse"`
}

func NewMsgRegisterResponse() *MsgRegisterResponse {
	var msg MsgRegisterResponse
	msg.XMLName.Local = "RegisterResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUnregisterRequest struct {
	XMLName xml.Name
	Unregister `xml:"UnregisterRequest"`
}

func NewMsgUnregisterRequest() *MsgUnregisterRequest {
	var msg MsgUnregisterRequest
	msg.XMLName.Local = "Unregister"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUnregisterResponse struct {
	XMLName xml.Name
	UnregisterResponse `xml:"UnregisterResponse"`
}

func NewMsgUnregisterResponse() *MsgUnregisterResponse {
	var msg MsgUnregisterResponse
	msg.XMLName.Local = "UnregisterResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetAgentConfigRequest struct {
	XMLName xml.Name
	GetAgentConfig `xml:"GetAgentConfigRequest"`
}

func NewMsgGetAgentConfigRequest() *MsgGetAgentConfigRequest {
	var msg MsgGetAgentConfigRequest
	msg.XMLName.Local = "GetAgentConfig"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgAgentConfig struct {
	XMLName xml.Name
	AgentConfig `xml:"AgentConfig"`
}

func NewMsgAgentConfig() *MsgAgentConfig {
	var msg MsgAgentConfig
	msg.XMLName.Local = "AgentConfig"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetTidByNameRequest struct {
	XMLName xml.Name
	MtGetTidByName `xml:"MtGetTidByNameRequest"`
}

func NewMsgMtGetTidByNameRequest() *MsgMtGetTidByNameRequest {
	var msg MsgMtGetTidByNameRequest
	msg.XMLName.Local = "MtGetTidByName"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetTidByNameResponse struct {
	XMLName xml.Name
	MtGetTidByNameResponse `xml:"MtGetTidByNameResponse"`
}

func NewMsgMtGetTidByNameResponse() *MsgMtGetTidByNameResponse {
	var msg MsgMtGetTidByNameResponse
	msg.XMLName.Local = "MtGetTidByNameResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetMteinfoRequest struct {
	XMLName xml.Name
	MtGetMteinfo `xml:"MtGetMteinfoRequest"`
}

func NewMsgMtGetMteinfoRequest() *MsgMtGetMteinfoRequest {
	var msg MsgMtGetMteinfoRequest
	msg.XMLName.Local = "MtGetMteinfo"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetMteinfoResponse struct {
	XMLName xml.Name
	MtGetMteinfoResponse `xml:"MtGetMteinfoResponse"`
}

func NewMsgMtGetMteinfoResponse() *MsgMtGetMteinfoResponse {
	var msg MsgMtGetMteinfoResponse
	msg.XMLName.Local = "MtGetMteinfoResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtReadRequest struct {
	XMLName xml.Name
	MtRead `xml:"MtReadRequest"`
}

func NewMsgMtReadRequest() *MsgMtReadRequest {
	var msg MsgMtReadRequest
	msg.XMLName.Local = "MtRead"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtReadResponse struct {
	XMLName xml.Name
	MtReadResponse `xml:"MtReadResponse"`
}

func NewMsgMtReadResponse() *MsgMtReadResponse {
	var msg MsgMtReadResponse
	msg.XMLName.Local = "MtReadResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgPerfReadRequest struct {
	XMLName xml.Name
	PerfRead `xml:"PerfReadRequest"`
}

func NewMsgPerfReadRequest() *MsgPerfReadRequest {
	var msg MsgPerfReadRequest
	msg.XMLName.Local = "PerfRead"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgPerfReadResponse struct {
	XMLName xml.Name
	PerfReadResponse `xml:"PerfReadResponse"`
}

func NewMsgPerfReadResponse() *MsgPerfReadResponse {
	var msg MsgPerfReadResponse
	msg.XMLName.Local = "PerfReadResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgPerfReadSmoothDataRequest struct {
	XMLName xml.Name
	PerfReadSmoothData `xml:"PerfReadSmoothDataRequest"`
}

func NewMsgPerfReadSmoothDataRequest() *MsgPerfReadSmoothDataRequest {
	var msg MsgPerfReadSmoothDataRequest
	msg.XMLName.Local = "PerfReadSmoothData"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgPerfReadSmoothDataResponse struct {
	XMLName xml.Name
	PerfReadSmoothDataResponse `xml:"PerfReadSmoothDataResponse"`
}

func NewMsgPerfReadSmoothDataResponse() *MsgPerfReadSmoothDataResponse {
	var msg MsgPerfReadSmoothDataResponse
	msg.XMLName.Local = "PerfReadSmoothDataResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMscReadCacheRequest struct {
	XMLName xml.Name
	MscReadCache `xml:"MscReadCacheRequest"`
}

func NewMsgMscReadCacheRequest() *MsgMscReadCacheRequest {
	var msg MsgMscReadCacheRequest
	msg.XMLName.Local = "MscReadCache"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMscReadCacheResponse struct {
	XMLName xml.Name
	MscReadCacheResponse `xml:"MscReadCacheResponse"`
}

func NewMsgMscReadCacheResponse() *MsgMscReadCacheResponse {
	var msg MsgMscReadCacheResponse
	msg.XMLName.Local = "MscReadCacheResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilSnglmsgReadRawdataRequest struct {
	XMLName xml.Name
	UtilSnglmsgReadRawdata `xml:"UtilSnglmsgReadRawdataRequest"`
}

func NewMsgUtilSnglmsgReadRawdataRequest() *MsgUtilSnglmsgReadRawdataRequest {
	var msg MsgUtilSnglmsgReadRawdataRequest
	msg.XMLName.Local = "UtilSnglmsgReadRawdata"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilSnglmsgReadRawdataResponse struct {
	XMLName xml.Name
	UtilSnglmsgReadRawdataResponse `xml:"UtilSnglmsgReadRawdataResponse"`
}

func NewMsgUtilSnglmsgReadRawdataResponse() *MsgUtilSnglmsgReadRawdataResponse {
	var msg MsgUtilSnglmsgReadRawdataResponse
	msg.XMLName.Local = "UtilSnglmsgReadRawdataResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilMtReadAllRequest struct {
	XMLName xml.Name
	UtilMtReadAll `xml:"UtilMtReadAllRequest"`
}

func NewMsgUtilMtReadAllRequest() *MsgUtilMtReadAllRequest {
	var msg MsgUtilMtReadAllRequest
	msg.XMLName.Local = "UtilMtReadAll"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilMtReadAllResponse struct {
	XMLName xml.Name
	UtilMtReadAllResponse `xml:"UtilMtReadAllResponse"`
}

func NewMsgUtilMtReadAllResponse() *MsgUtilMtReadAllResponse {
	var msg MsgUtilMtReadAllResponse
	msg.XMLName.Local = "UtilMtReadAllResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgTextAttrReadRequest struct {
	XMLName xml.Name
	TextAttrRead `xml:"TextAttrReadRequest"`
}

func NewMsgTextAttrReadRequest() *MsgTextAttrReadRequest {
	var msg MsgTextAttrReadRequest
	msg.XMLName.Local = "TextAttrRead"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgTextAttrReadResponse struct {
	XMLName xml.Name
	TextAttrReadResponse `xml:"TextAttrReadResponse"`
}

func NewMsgTextAttrReadResponse() *MsgTextAttrReadResponse {
	var msg MsgTextAttrReadResponse
	msg.XMLName.Local = "TextAttrReadResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgReferenceReadRequest struct {
	XMLName xml.Name
	ReferenceRead `xml:"ReferenceReadRequest"`
}

func NewMsgReferenceReadRequest() *MsgReferenceReadRequest {
	var msg MsgReferenceReadRequest
	msg.XMLName.Local = "ReferenceRead"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgReferenceReadResponse struct {
	XMLName xml.Name
	ReferenceReadResponse `xml:"ReferenceReadResponse"`
}

func NewMsgReferenceReadResponse() *MsgReferenceReadResponse {
	var msg MsgReferenceReadResponse
	msg.XMLName.Local = "ReferenceReadResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilMtGetAidByTidRequest struct {
	XMLName xml.Name
	UtilMtGetAidByTid `xml:"UtilMtGetAidByTidRequest"`
}

func NewMsgUtilMtGetAidByTidRequest() *MsgUtilMtGetAidByTidRequest {
	var msg MsgUtilMtGetAidByTidRequest
	msg.XMLName.Local = "UtilMtGetAidByTid"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilMtGetAidByTidResponse struct {
	XMLName xml.Name
	UtilMtGetAidByTidResponse `xml:"UtilMtGetAidByTidResponse"`
}

func NewMsgUtilMtGetAidByTidResponse() *MsgUtilMtGetAidByTidResponse {
	var msg MsgUtilMtGetAidByTidResponse
	msg.XMLName.Local = "UtilMtGetAidByTidResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilReadRawalertByAidRequest struct {
	XMLName xml.Name
	UtilReadRawalertByAid `xml:"UtilReadRawalertByAidRequest"`
}

func NewMsgUtilReadRawalertByAidRequest() *MsgUtilReadRawalertByAidRequest {
	var msg MsgUtilReadRawalertByAidRequest
	msg.XMLName.Local = "UtilReadRawalertByAid"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilReadRawalertByAidResponse struct {
	XMLName xml.Name
	UtilReadRawalertByAidResponse `xml:"UtilReadRawalertByAidResponse"`
}

func NewMsgUtilReadRawalertByAidResponse() *MsgUtilReadRawalertByAidResponse {
	var msg MsgUtilReadRawalertByAidResponse
	msg.XMLName.Local = "UtilReadRawalertByAidResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgToolGetEffectiveRequest struct {
	XMLName xml.Name
	ToolGetEffective `xml:"ToolGetEffectiveRequest"`
}

func NewMsgToolGetEffectiveRequest() *MsgToolGetEffectiveRequest {
	var msg MsgToolGetEffectiveRequest
	msg.XMLName.Local = "ToolGetEffective"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgToolGetEffectiveResponse struct {
	XMLName xml.Name
	ToolGetEffectiveResponse `xml:"ToolGetEffectiveResponse"`
}

func NewMsgToolGetEffectiveResponse() *MsgToolGetEffectiveResponse {
	var msg MsgToolGetEffectiveResponse
	msg.XMLName.Local = "ToolGetEffectiveResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgTriggerDataCollectionRequest struct {
	XMLName xml.Name
	TriggerDataCollection `xml:"TriggerDataCollectionRequest"`
}

func NewMsgTriggerDataCollectionRequest() *MsgTriggerDataCollectionRequest {
	var msg MsgTriggerDataCollectionRequest
	msg.XMLName.Local = "TriggerDataCollection"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgTriggerDataCollectionResponse struct {
	XMLName xml.Name
	TriggerDataCollectionResponse `xml:"TriggerDataCollectionResponse"`
}

func NewMsgTriggerDataCollectionResponse() *MsgTriggerDataCollectionResponse {
	var msg MsgTriggerDataCollectionResponse
	msg.XMLName.Local = "TriggerDataCollectionResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtChangeStatusRequest struct {
	XMLName xml.Name
	MtChangeStatus `xml:"MtChangeStatusRequest"`
}

func NewMsgMtChangeStatusRequest() *MsgMtChangeStatusRequest {
	var msg MsgMtChangeStatusRequest
	msg.XMLName.Local = "MtChangeStatus"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtChangeStatusResponse struct {
	XMLName xml.Name
	MtChangeStatusResponse `xml:"MtChangeStatusResponse"`
}

func NewMsgMtChangeStatusResponse() *MsgMtChangeStatusResponse {
	var msg MsgMtChangeStatusResponse
	msg.XMLName.Local = "MtChangeStatusResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtCustomizeWriteRequest struct {
	XMLName xml.Name
	MtCustomizeWrite `xml:"MtCustomizeWriteRequest"`
}

func NewMsgMtCustomizeWriteRequest() *MsgMtCustomizeWriteRequest {
	var msg MsgMtCustomizeWriteRequest
	msg.XMLName.Local = "MtCustomizeWrite"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtCustomizeWriteResponse struct {
	XMLName xml.Name
	MtCustomizeWriteResponse `xml:"MtCustomizeWriteResponse"`
}

func NewMsgMtCustomizeWriteResponse() *MsgMtCustomizeWriteResponse {
	var msg MsgMtCustomizeWriteResponse
	msg.XMLName.Local = "MtCustomizeWriteResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtDestroyMarkNTryRequest struct {
	XMLName xml.Name
	MtDestroyMarkNTry `xml:"MtDestroyMarkNTryRequest"`
}

func NewMsgMtDestroyMarkNTryRequest() *MsgMtDestroyMarkNTryRequest {
	var msg MsgMtDestroyMarkNTryRequest
	msg.XMLName.Local = "MtDestroyMarkNTry"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtDestroyMarkNTryResponse struct {
	XMLName xml.Name
	MtDestroyMarkNTryResponse `xml:"MtDestroyMarkNTryResponse"`
}

func NewMsgMtDestroyMarkNTryResponse() *MsgMtDestroyMarkNTryResponse {
	var msg MsgMtDestroyMarkNTryResponse
	msg.XMLName.Local = "MtDestroyMarkNTryResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtResetRequest struct {
	XMLName xml.Name
	MtReset `xml:"MtResetRequest"`
}

func NewMsgMtResetRequest() *MsgMtResetRequest {
	var msg MsgMtResetRequest
	msg.XMLName.Local = "MtReset"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtResetResponse struct {
	XMLName xml.Name
	MtResetResponse `xml:"MtResetResponse"`
}

func NewMsgMtResetResponse() *MsgMtResetResponse {
	var msg MsgMtResetResponse
	msg.XMLName.Local = "MtResetResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgPerfCustomizeWriteRequest struct {
	XMLName xml.Name
	PerfCustomizeWrite `xml:"PerfCustomizeWriteRequest"`
}

func NewMsgPerfCustomizeWriteRequest() *MsgPerfCustomizeWriteRequest {
	var msg MsgPerfCustomizeWriteRequest
	msg.XMLName.Local = "PerfCustomizeWrite"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgPerfCustomizeWriteResponse struct {
	XMLName xml.Name
	PerfCustomizeWriteResponse `xml:"PerfCustomizeWriteResponse"`
}

func NewMsgPerfCustomizeWriteResponse() *MsgPerfCustomizeWriteResponse {
	var msg MsgPerfCustomizeWriteResponse
	msg.XMLName.Local = "PerfCustomizeWriteResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMscCustomizeWriteRequest struct {
	XMLName xml.Name
	MscCustomizeWrite `xml:"MscCustomizeWriteRequest"`
}

func NewMsgMscCustomizeWriteRequest() *MsgMscCustomizeWriteRequest {
	var msg MsgMscCustomizeWriteRequest
	msg.XMLName.Local = "MscCustomizeWrite"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMscCustomizeWriteResponse struct {
	XMLName xml.Name
	MscCustomizeWriteResponse `xml:"MscCustomizeWriteResponse"`
}

func NewMsgMscCustomizeWriteResponse() *MsgMscCustomizeWriteResponse {
	var msg MsgMscCustomizeWriteResponse
	msg.XMLName.Local = "MscCustomizeWriteResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMscDeleteLinesRequest struct {
	XMLName xml.Name
	MscDeleteLines `xml:"MscDeleteLinesRequest"`
}

func NewMsgMscDeleteLinesRequest() *MsgMscDeleteLinesRequest {
	var msg MsgMscDeleteLinesRequest
	msg.XMLName.Local = "MscDeleteLines"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMscDeleteLinesResponse struct {
	XMLName xml.Name
	MscDeleteLinesResponse `xml:"MscDeleteLinesResponse"`
}

func NewMsgMscDeleteLinesResponse() *MsgMscDeleteLinesResponse {
	var msg MsgMscDeleteLinesResponse
	msg.XMLName.Local = "MscDeleteLinesResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgSnglmgsCustomizeWriteRequest struct {
	XMLName xml.Name
	SnglmgsCustomizeWrite `xml:"SnglmgsCustomizeWriteRequest"`
}

func NewMsgSnglmgsCustomizeWriteRequest() *MsgSnglmgsCustomizeWriteRequest {
	var msg MsgSnglmgsCustomizeWriteRequest
	msg.XMLName.Local = "SnglmgsCustomizeWrite"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgSnglmgsCustomizeWriteResponse struct {
	XMLName xml.Name
	SnglmgsCustomizeWriteResponse `xml:"SnglmgsCustomizeWriteResponse"`
}

func NewMsgSnglmgsCustomizeWriteResponse() *MsgSnglmgsCustomizeWriteResponse {
	var msg MsgSnglmgsCustomizeWriteResponse
	msg.XMLName.Local = "SnglmgsCustomizeWriteResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgToolSetRequest struct {
	XMLName xml.Name
	ToolSet `xml:"ToolSetRequest"`
}

func NewMsgToolSetRequest() *MsgToolSetRequest {
	var msg MsgToolSetRequest
	msg.XMLName.Local = "ToolSet"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgToolSetResponse struct {
	XMLName xml.Name
	ToolSetResponse `xml:"ToolSetResponse"`
}

func NewMsgToolSetResponse() *MsgToolSetResponse {
	var msg MsgToolSetResponse
	msg.XMLName.Local = "ToolSetResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgToolSetRuntimeStatusRequest struct {
	XMLName xml.Name
	ToolSetRuntimeStatus `xml:"ToolSetRuntimeStatusRequest"`
}

func NewMsgToolSetRuntimeStatusRequest() *MsgToolSetRuntimeStatusRequest {
	var msg MsgToolSetRuntimeStatusRequest
	msg.XMLName.Local = "ToolSetRuntimeStatus"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgToolSetRuntimeStatusResponse struct {
	XMLName xml.Name
	ToolSetRuntimeStatusResponse `xml:"ToolSetRuntimeStatusResponse"`
}

func NewMsgToolSetRuntimeStatusResponse() *MsgToolSetRuntimeStatusResponse {
	var msg MsgToolSetRuntimeStatusResponse
	msg.XMLName.Local = "ToolSetRuntimeStatusResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilAlChangeStatusRequest struct {
	XMLName xml.Name
	UtilAlChangeStatus `xml:"UtilAlChangeStatusRequest"`
}

func NewMsgUtilAlChangeStatusRequest() *MsgUtilAlChangeStatusRequest {
	var msg MsgUtilAlChangeStatusRequest
	msg.XMLName.Local = "UtilAlChangeStatus"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilAlChangeStatusResponse struct {
	XMLName xml.Name
	UtilAlChangeStatusResponse `xml:"UtilAlChangeStatusResponse"`
}

func NewMsgUtilAlChangeStatusResponse() *MsgUtilAlChangeStatusResponse {
	var msg MsgUtilAlChangeStatusResponse
	msg.XMLName.Local = "UtilAlChangeStatusResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtDbsetToWpsetByTidRequest struct {
	XMLName xml.Name
	MtDbsetToWpsetByTid `xml:"MtDbsetToWpsetByTidRequest"`
}

func NewMsgMtDbsetToWpsetByTidRequest() *MsgMtDbsetToWpsetByTidRequest {
	var msg MsgMtDbsetToWpsetByTidRequest
	msg.XMLName.Local = "MtDbsetToWpsetByTid"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtDbsetToWpsetByTidResponse struct {
	XMLName xml.Name
	MtDbsetToWpsetByTidResponse `xml:"MtDbsetToWpsetByTidResponse"`
}

func NewMsgMtDbsetToWpsetByTidResponse() *MsgMtDbsetToWpsetByTidResponse {
	var msg MsgMtDbsetToWpsetByTidResponse
	msg.XMLName.Local = "MtDbsetToWpsetByTidResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgSystemObjectSetValueRequest struct {
	XMLName xml.Name
	SystemObjectSetValue `xml:"SystemObjectSetValueRequest"`
}

func NewMsgSystemObjectSetValueRequest() *MsgSystemObjectSetValueRequest {
	var msg MsgSystemObjectSetValueRequest
	msg.XMLName.Local = "SystemObjectSetValue"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgSystemObjectSetValueResponse struct {
	XMLName xml.Name
	SystemObjectSetValueResponse `xml:"SystemObjectSetValueResponse"`
}

func NewMsgSystemObjectSetValueResponse() *MsgSystemObjectSetValueResponse {
	var msg MsgSystemObjectSetValueResponse
	msg.XMLName.Local = "SystemObjectSetValueResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetMtesByRequestTableRequest struct {
	XMLName xml.Name
	GetMtesByRequestTable `xml:"GetMtesByRequestTableRequest"`
}

func NewMsgGetMtesByRequestTableRequest() *MsgGetMtesByRequestTableRequest {
	var msg MsgGetMtesByRequestTableRequest
	msg.XMLName.Local = "GetMtesByRequestTable"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetMtesByRequestTableResponse struct {
	XMLName xml.Name
	GetMtesByRequestTableResponse `xml:"GetMtesByRequestTableResponse"`
}

func NewMsgGetMtesByRequestTableResponse() *MsgGetMtesByRequestTableResponse {
	var msg MsgGetMtesByRequestTableResponse
	msg.XMLName.Local = "GetMtesByRequestTableResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetMcInLocalMsRequest struct {
	XMLName xml.Name
	GetMcInLocalMs `xml:"GetMcInLocalMsRequest"`
}

func NewMsgGetMcInLocalMsRequest() *MsgGetMcInLocalMsRequest {
	var msg MsgGetMcInLocalMsRequest
	msg.XMLName.Local = "GetMcInLocalMs"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetMcInLocalMsResponse struct {
	XMLName xml.Name
	GetMcInLocalMsResponse `xml:"GetMcInLocalMsResponse"`
}

func NewMsgGetMcInLocalMsResponse() *MsgGetMcInLocalMsResponse {
	var msg MsgGetMcInLocalMsResponse
	msg.XMLName.Local = "GetMcInLocalMsResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMsGetMteclsInLocalMsRequest struct {
	XMLName xml.Name
	MsGetMteclsInLocalMs `xml:"MsGetMteclsInLocalMsRequest"`
}

func NewMsgMsGetMteclsInLocalMsRequest() *MsgMsGetMteclsInLocalMsRequest {
	var msg MsgMsGetMteclsInLocalMsRequest
	msg.XMLName.Local = "MsGetMteclsInLocalMs"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMsGetMteclsInLocalMsResponse struct {
	XMLName xml.Name
	MsGetMteclsInLocalMsResponse `xml:"MsGetMteclsInLocalMsResponse"`
}

func NewMsgMsGetMteclsInLocalMsResponse() *MsgMsGetMteclsInLocalMsResponse {
	var msg MsgMsGetMteclsInLocalMsResponse
	msg.XMLName.Local = "MsGetMteclsInLocalMsResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMsGetLocalMsInfoRequest struct {
	XMLName xml.Name
	MsGetLocalMsInfo `xml:"MsGetLocalMsInfoRequest"`
}

func NewMsgMsGetLocalMsInfoRequest() *MsgMsGetLocalMsInfoRequest {
	var msg MsgMsGetLocalMsInfoRequest
	msg.XMLName.Local = "MsGetLocalMsInfo"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMsGetLocalMsInfoResponse struct {
	XMLName xml.Name
	MsGetLocalMsInfoResponse `xml:"MsGetLocalMsInfoResponse"`
}

func NewMsgMsGetLocalMsInfoResponse() *MsgMsGetLocalMsInfoResponse {
	var msg MsgMsGetLocalMsInfoResponse
	msg.XMLName.Local = "MsGetLocalMsInfoResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilMtGetTreeLocalRequest struct {
	XMLName xml.Name
	UtilMtGetTreeLocal `xml:"UtilMtGetTreeLocalRequest"`
}

func NewMsgUtilMtGetTreeLocalRequest() *MsgUtilMtGetTreeLocalRequest {
	var msg MsgUtilMtGetTreeLocalRequest
	msg.XMLName.Local = "UtilMtGetTreeLocal"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgUtilMtGetTreeLocalResponse struct {
	XMLName xml.Name
	UtilMtGetTreeLocalResponse `xml:"UtilMtGetTreeLocalResponse"`
}

func NewMsgUtilMtGetTreeLocalResponse() *MsgUtilMtGetTreeLocalResponse {
	var msg MsgUtilMtGetTreeLocalResponse
	msg.XMLName.Local = "UtilMtGetTreeLocalResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgInfoGetTreeRequest struct {
	XMLName xml.Name
	InfoGetTree `xml:"InfoGetTreeRequest"`
}

func NewMsgInfoGetTreeRequest() *MsgInfoGetTreeRequest {
	var msg MsgInfoGetTreeRequest
	msg.XMLName.Local = "InfoGetTree"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgInfoGetTreeResponse struct {
	XMLName xml.Name
	InfoGetTreeResponse `xml:"InfoGetTreeResponse"`
}

func NewMsgInfoGetTreeResponse() *MsgInfoGetTreeResponse {
	var msg MsgInfoGetTreeResponse
	msg.XMLName.Local = "InfoGetTreeResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetMtListByMtclassRequest struct {
	XMLName xml.Name
	GetMtListByMtclass `xml:"GetMtListByMtclassRequest"`
}

func NewMsgGetMtListByMtclassRequest() *MsgGetMtListByMtclassRequest {
	var msg MsgGetMtListByMtclassRequest
	msg.XMLName.Local = "GetMtListByMtclass"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetMtListByMtclassResponse struct {
	XMLName xml.Name
	GetMtListByMtclassResponse `xml:"GetMtListByMtclassResponse"`
}

func NewMsgGetMtListByMtclassResponse() *MsgGetMtListByMtclassResponse {
	var msg MsgGetMtListByMtclassResponse
	msg.XMLName.Local = "GetMtListByMtclassResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMteGetByToolRunstatusRequest struct {
	XMLName xml.Name
	MteGetByToolRunstatus `xml:"MteGetByToolRunstatusRequest"`
}

func NewMsgMteGetByToolRunstatusRequest() *MsgMteGetByToolRunstatusRequest {
	var msg MsgMteGetByToolRunstatusRequest
	msg.XMLName.Local = "MteGetByToolRunstatus"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMteGetByToolRunstatusResponse struct {
	XMLName xml.Name
	MteGetByToolRunstatusResponse `xml:"MteGetByToolRunstatusResponse"`
}

func NewMsgMteGetByToolRunstatusResponse() *MsgMteGetByToolRunstatusResponse {
	var msg MsgMteGetByToolRunstatusResponse
	msg.XMLName.Local = "MteGetByToolRunstatusResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetAllToolsToSetRequest struct {
	XMLName xml.Name
	MtGetAllToolsToSet `xml:"MtGetAllToolsToSetRequest"`
}

func NewMsgMtGetAllToolsToSetRequest() *MsgMtGetAllToolsToSetRequest {
	var msg MsgMtGetAllToolsToSetRequest
	msg.XMLName.Local = "MtGetAllToolsToSet"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetAllToolsToSetResponse struct {
	XMLName xml.Name
	MtGetAllToolsToSetResponse `xml:"MtGetAllToolsToSetResponse"`
}

func NewMsgMtGetAllToolsToSetResponse() *MsgMtGetAllToolsToSetResponse {
	var msg MsgMtGetAllToolsToSetResponse
	msg.XMLName.Local = "MtGetAllToolsToSetResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetAllToCustRequest struct {
	XMLName xml.Name
	MtGetAllToCust `xml:"MtGetAllToCustRequest"`
}

func NewMsgMtGetAllToCustRequest() *MsgMtGetAllToCustRequest {
	var msg MsgMtGetAllToCustRequest
	msg.XMLName.Local = "MtGetAllToCust"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgMtGetAllToCustResponse struct {
	XMLName xml.Name
	MtGetAllToCustResponse `xml:"MtGetAllToCustResponse"`
}

func NewMsgMtGetAllToCustResponse() *MsgMtGetAllToCustResponse {
	var msg MsgMtGetAllToCustResponse
	msg.XMLName.Local = "MtGetAllToCustResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetListOfMaByCusGrpRequest struct {
	XMLName xml.Name
	GetListOfMaByCusGrp `xml:"GetListOfMaByCusGrpRequest"`
}

func NewMsgGetListOfMaByCusGrpRequest() *MsgGetListOfMaByCusGrpRequest {
	var msg MsgGetListOfMaByCusGrpRequest
	msg.XMLName.Local = "GetListOfMaByCusGrp"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgGetListOfMaByCusGrpResponse struct {
	XMLName xml.Name
	GetListOfMaByCusGrpResponse `xml:"GetListOfMaByCusGrpResponse"`
}

func NewMsgGetListOfMaByCusGrpResponse() *MsgGetListOfMaByCusGrpResponse {
	var msg MsgGetListOfMaByCusGrpResponse
	msg.XMLName.Local = "GetListOfMaByCusGrpResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgRequestLogonFileRequest struct {
	XMLName xml.Name
	RequestLogonFile `xml:"RequestLogonFileRequest"`
}

func NewMsgRequestLogonFileRequest() *MsgRequestLogonFileRequest {
	var msg MsgRequestLogonFileRequest
	msg.XMLName.Local = "RequestLogonFile"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type MsgRequestLogonFileResponse struct {
	XMLName xml.Name
	RequestLogonFileResponse `xml:"RequestLogonFileResponse"`
}

func NewMsgRequestLogonFileResponse() *MsgRequestLogonFileResponse {
	var msg MsgRequestLogonFileResponse
	msg.XMLName.Local = "RequestLogonFileResponse"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}



type SoapEnvelope struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Soap    string      `xml:"xmlns soap,attr"`
	Body    Body 		`xml:"Body"`
}

func NewSoapEnvelope() *SoapEnvelope {
	se := &SoapEnvelope{}
	se.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	return se
}

type Body struct {
	XMLName xml.Name    `xml:"Body"`
	Content interface{} `xml:",any"`
}

type SoapFault struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
	Detail      string   `xml:"detail"`
}

func (s *SAPCCMS)SendRequest(send, receive *SoapEnvelope) error {
	sendMarshalled, err := xml.MarshalIndent(&send, "", "  ")
	if err != nil {
		return nil
	}
	var soapreader = strings.NewReader(string(sendMarshalled))
	resp, err := http.Post(s.Url, "text/xml; charset=utf-8", soapreader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	xml.Unmarshal(buffer, receive)
	return nil
}
