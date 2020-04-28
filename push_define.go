package JpushGo

const (
	// 推送消息 API
	PUSH_POST = "https://api.jpush.cn/v3/push"
	// 应用分组推送
	PUSH_GROUP_POST = "https://api.jpush.cn/v3/grouppush"
	// 批量单推（VIP专属接口）
	PUSH_BATCH_REGID_POST = "https://api.jpush.cn/v3/push/batch/regid/single"
	PUSH_BATCH_ALIAS_POST = "https://api.jpush.cn/v3/push/batch/alias/single"
	// 推送校验 API
	PUSH_VALIDATE_POST = "https://api.jpush.cn/v3/push/validate"
	// 文件推送
	PUSH_FILE_POST = "https://api.jpush.cn/v3/push/file"
	// 推送撤销
	PUSH_DELETE  = "https://api.jpush.cn/v3/push/{msgid}"
	PUSH_CID_GET = "https://api.jpush.cn/v3/push/cid?count={count}&type={type}"

	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"

	IOS     = "ios"
	ANDROID = "android"

	PUSH_AUDIENCE_ALL = "all"

	CONTENT_TYPE_JSON = "application/json"
)

type (
	PushRequest struct {
		Cid          string            `json:"cid"`
		Platform     []string          `json:"platform"`
		Audience     interface{}       `json:"audience"`
		Notification *PushNotification `json:"notification,omitempty"`
		Message      *PushMessage      `json:"message,omitempty"`
		SmsMessage   *SmsMessage       `json:"sms_message,omitempty"`
		Options      *PushOptions      `json:"options,omitempty"`
	}
	PushAudience struct {
		Tag            []string `json:"tag,omitempty"`
		TagAnd         []string `json:"tag_and,omitempty"`
		TagNot         []string `json:"tag_not,omitempty"`
		Alias          []string `json:"alias,omitempty"`
		RegistrationId []string `json:"registration_id,omitempty"`
		Segment        []string `json:"segment,omitempty"`
		ABTest         []string `json:"abtest,omitempty"`
	}
	PushNotification struct {
		Alert   string               `json:"alert,omitempty"`
		Android *NotificationAndroid `json:"android,omitempty"`
		IOS     *NotificationIOS     `json:"ios,omitempty"`
	}
	NotificationAndroid struct {
		Alert      string                 `json:"alert"`
		Title      string                 `json:"title,omitempty"`
		BuilderId  int                    `json:"builder_id,int,omitempty"`
		Priority   int                    `json:"priority,omitempty"`
		Category   string                 `json:"category,omitempty"`
		Style      int                    `json:"style,int,omitempty"`
		AlertType  int                    `json:"alert_type,int,omitempty"`
		BigText    string                 `json:"big_text,omitempty"`
		Inbox      map[string]interface{} `json:"inbox,omitempty"`
		BigPicPath string                 `json:"big_pic_path,omitempty"`
		Extras     map[string]interface{} `json:"extras,omitempty"`
	}

	NotificationIOS struct {
		Alert            interface{}            `json:"alert"`
		Sound            string                 `json:"sound,omitempty"`
		Badge            int                    `json:"badge,int,omitempty"`
		ContentAvailable bool                   `json:"content-available,omitempty"`
		MutableContent   bool                   `json:"mutable-content,omitempty"`
		Category         string                 `json:"category,omitempty"`
		Extras           map[string]interface{} `json:"extras,omitempty"`
	}

	PushMessage struct {
		MsgContent  string                 `json:"msg_content"`
		Title       string                 `json:"title,omitempty"`
		ContentType string                 `json:"content_type,omitempty"`
		Extras      map[string]interface{} `json:"extras,omitempty"`
	}

	SmsMessage struct {
		Content   string `json:"content"`
		DelayTime int    `json:"delay_time,int,omitempty"`
	}

	PushOptions struct {
		SendNo          int64  `json:"sendno,int,omitempty"`
		TimeToLive      int    `json:"time_to_live,int,omitempty"`
		OverrideMsgId   int64  `json:"override_msg_id,int64,omitempty"`
		ApnsProduction  bool   `json:"apns_production"`
		ApnsCollapseId  string `json:"apns_collapse_id,omitempty"`
		BigPushDuration int    `json:"big_push_duration,int,omitempty"`
	}
)

func NewPush() *PushRequest {
	return &PushRequest{
		Notification: &PushNotification{
			Android: &NotificationAndroid{},
			IOS:     &NotificationIOS{},
		},
		Message: &PushMessage{},
		// SmsMessage: &SmsMessage{},
		Options: &PushOptions{},
	}
}
func (p *PushRequest) AddCid(cid string) {
	p.Cid = cid
}
func (p *PushRequest) PlatformAll() {
	p.Platform = []string{IOS, ANDROID}
}
func (p *PushRequest) PlatformIOS() {
	p.Platform = []string{IOS}
}
func (p *PushRequest) PlatformAndroid() {
	p.Platform = []string{ANDROID}
}

// start audience
func (p *PushRequest) AudienceAll() {
	p.Audience = PUSH_AUDIENCE_ALL
}
func (p *PushRequest) AddAlias(val []string) {
	p.Audience = &PushAudience{Alias: val}
}
func (p *PushRequest) AddTag(val []string) {
	p.Audience = &PushAudience{Tag: val}
}
func (p *PushRequest) AddTagAnd(val []string) {
	p.Audience = &PushAudience{TagAnd: val}
}
func (p *PushRequest) AddTagNot(val []string) {
	p.Audience = &PushAudience{TagNot: val}
}
func (p *PushRequest) AddRegistrationId(val []string) {
	p.Audience = &PushAudience{RegistrationId: val}
}
func (p *PushRequest) AddSegment(val []string) {
	p.Audience = &PushAudience{Segment: val}
}
func (p *PushRequest) AddABTest(val []string) {
	p.Audience = &PushAudience{ABTest: val}
}

// end audience

func (p *PushRequest) AddTitle(title string) {
	p.Notification.Android.Title = title
	p.Message.Title = title
}
func (p *PushRequest) AddAlert(alert string) {
	p.Notification.Alert = alert
	p.Notification.Android.Alert = alert
	p.Notification.IOS.Alert = alert
	p.Message.MsgContent = alert
}
func (p *PushRequest) AddExtras(extras map[string]interface{}) {
	p.Notification.Android.Extras = extras
	p.Notification.IOS.Extras = extras
	p.Message.Extras = extras
}

// android start
func (p *PushRequest) AddBuilderId(val int) {
	p.Notification.Android.BuilderId = val
}
func (p *PushRequest) AddPriority(val int) {
	p.Notification.Android.Priority = val
}
func (p *PushRequest) AddAndroidCategory(val string) {
	p.Notification.Android.Category = val
}
func (p *PushRequest) AddStyle(val int) {
	p.Notification.Android.Style = val
}
func (p *PushRequest) AddAlertType(val int) {
	p.Notification.Android.AlertType = val
}
func (p *PushRequest) AddBigText(val string) {
	p.Notification.Android.BigText = val
}
func (p *PushRequest) AddInbox(val map[string]interface{}) {
	p.Notification.Android.Inbox = val
}
func (p *PushRequest) AddBigPicPath(val string) {
	p.Notification.Android.BigPicPath = val
}

// android end

// ios start
func (p *PushRequest) AddSound(val string) {
	p.Notification.IOS.Sound = val
}
func (p *PushRequest) AddBadge(val int) {
	p.Notification.IOS.Badge = val
}
func (p *PushRequest) AddContentAvailable(val bool) {
	p.Notification.IOS.ContentAvailable = val
}
func (p *PushRequest) AddMutableContent(val bool) {
	p.Notification.IOS.MutableContent = val
}
func (p *PushRequest) AddIOSCategory(val string) {
	p.Notification.IOS.Category = val
}

// ios end

// message start
func (p *PushRequest) AddContentType(val string) {
	p.Message.ContentType = val
}

// message en

// options start
func (p *PushRequest) AddSendNo(sendNo int64) {
	p.Options.SendNo = sendNo
}
func (p *PushRequest) AddTimeToLive(live int) {
	p.Options.TimeToLive = live
}
func (p *PushRequest) AddApnsCollapseId(coll string) {
	p.Options.ApnsCollapseId = coll
}
func (p *PushRequest) AddApnsProduction(env bool) {
	p.Options.ApnsProduction = env
}
func (p *PushRequest) AddOverrideMsgId(val int64) {
	p.Options.OverrideMsgId = val
}
func (p *PushRequest) AddBigPushDuration(val int) {
	p.Options.BigPushDuration = val
}

// options end
