package JpushGo

var Platform []string

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

	PLATFORM_IOS     Platform = []string{"ios"}
	PLATFORM_ANDROID Platform = []string{"android"}
	PLATFORM_ALL     Platform = []string{"ios", "android"}

	CHARSET                    = "UTF-8"
	CONTENT_TYPE_JSON          = "application/json"
	DEFAULT_CONNECTION_TIMEOUT = 20 //seconds
	DEFAULT_SOCKET_TIMEOUT     = 30 // seconds
)

type (
	PushRequest struct {
		Cid          string            `json: "cid,omitempty"`
		Platform     Platform          `json: platform`
		Audience     *PushAudience     `json:"audience,omitempty"`
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
		Alert    string                `json:"alert,omitempty"`
		Android  *NotificationAndroid  `json:"android,omitempty"`
		IOS      *NotificationIOS      `json:"ios,omitempty"`
		WinPhone *NotificationWinPhone `json:"winphone,omitempty"`
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

	NotificationWinPhone struct {
		Alert    string                 `json:"alert"`
		Title    string                 `json:"title,omitempty"`
		OpenPage string                 `json:"_open_page,omitempty"`
		Extras   map[string]interface{} `json:"extras,omitempty"`
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
		SendNo          int    `json:"sendno,int,omitempty"`
		TimeToLive      int    `json:"time_to_live,int,omitempty"`
		OverrideMsgId   int64  `json:"override_msg_id,int64,omitempty"`
		ApnsProduction  bool   `json:"apns_production"`
		ApnsCollapseId  string `json:"apns_collapse_id,omitempty"`
		BigPushDuration int    `json:"big_push_duration,int,omitempty"`
	}
)
