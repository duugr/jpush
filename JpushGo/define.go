package JpushGo

import (
	"encoding/base64"
	"fmt"
	// "io"
	// "io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

const (
	APP_KEY    = "7d431e42dfa6a6d693ac2d04"
	APP_SECRET = "5e987ac6d2e04d95a9d8f0d1"
	// 推送消息 API
	URL_PUSH = "https://api.jpush.cn"
	// 获取统计数据
	URL_RECEIVED = "https://report.jpush.cn"
	// 操作 tag, alias
	URL_DEVICES = "https://device.jpush.cn"

	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"

	IOS      = "ios"
	ANDROID  = "android"
	WINPHONE = "winphone"

	CHARSET                    = "UTF-8"
	CONTENT_TYPE_JSON          = "application/json"
	DEFAULT_CONNECTION_TIMEOUT = 20 //seconds
	DEFAULT_SOCKET_TIMEOUT     = 30 // seconds

	// 获取当前设备的所有属性，包含 tags, alias，手机号码 mobile。
	get_devices = "{scheme}/v3/devices/{registration_id}"
	// 设置设备的别名与标签
	post_devices = "{scheme}/v3/devices/{registration_id}"
	// 获取用户在线状态（VIP 专属接口）
	post_devices_status = "{scheme}/v3/devices/status/"

	// 获取指定 alias 下的设备，最多输出 10 个；
	get_aliases = "{scheme}/v3/aliases/{alias_value}"
	// 删除一个别名，以及该别名与设备的绑定关系。
	delete_aliases = "{scheme}/v3/aliases/{alias_value}"
	// 批量解绑设备与别名之间的关系。
	post_aliases = "{scheme}/v3/aliases/{alias_value}"

	// 查询标签列表
	get_tags = "{scheme}/v3/tags/"
	// 判断设备与标签绑定关系,查询某个设备是否在 tag 下。
	get_tags_reg = "{scheme}/v3/tags/{tag_value}/registration_ids/{registration_id}"
	// 更新标签,为一个标签添加或者删除设备。
	post_tags = "{scheme}/v3/tags/{tag_value}"
	// 删除一个标签，以及标签与设备之间的关联关系。
	delete_tags = "{scheme}/v3/tags/{tag_value}"
)

type (
	Hyper struct {
		Url       string
		Params    url.Values
		UserAgent map[string]string
		Headers   map[string]string
		Client    *http.Client

		Result []byte
	}

	Platform struct {
		Os     interface{}
		osArry []string
	}

	Message struct {
		Content     string                 `json:"msg_content"`
		Title       string                 `json:"title,omitempty"`
		ContentType string                 `json:"content_type,omitempty"`
		Extras      map[string]interface{} `json:"extras,omitempty"`
	}

	Schedule struct {
		Cid     string                 `json:"cid"`
		Name    string                 `json:"name"`
		Enabled bool                   `json:"enabled"`
		Trigger map[string]interface{} `json:"trigger"`
		Push    *PayLoad               `json:"push"`
	}

	Notice struct {
		Alert    string          `json:"alert,omitempty"`
		Android  *AndroidNotice  `json:"android,omitempty"`
		IOS      *IOSNotice      `json:"ios,omitempty"`
		WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
	}

	AndroidNotice struct {
		Alert     string                 `json:"alert"`
		Title     string                 `json:"title,omitempty"`
		BuilderId int                    `json:"builder_id,omitempty"`
		Extras    map[string]interface{} `json:"extras,omitempty"`
	}

	IOSNotice struct {
		Alert            interface{}            `json:"alert"`
		Sound            string                 `json:"sound,omitempty"`
		Badge            string                 `json:"badge,omitempty"`
		ContentAvailable bool                   `json:"content-available,omitempty"`
		MutableContent   bool                   `json:"mutable-content,omitempty"`
		Category         string                 `json:"category,omitempty"`
		Extras           map[string]interface{} `json:"extras,omitempty"`
	}

	WinPhoneNotice struct {
		Alert    string                 `json:"alert"`
		Title    string                 `json:"title,omitempty"`
		OpenPage string                 `json:"_open_page,omitempty"`
		Extras   map[string]interface{} `json:"extras,omitempty"`
	}
	PayLoad struct {
		Platform     interface{} `json:"platform"`
		Audience     interface{} `json:"audience"`
		Notification interface{} `json:"notification,omitempty"`
		Message      interface{} `json:"message,omitempty"`
		Options      *Option     `json:"options,omitempty"`
	}
	Option struct {
		SendNo          int   `json:"sendno,omitempty"`
		TimeLive        int   `json:"time_to_live,omitempty"`
		ApnsProduction  bool  `json:"apns_production"`
		OverrideMsgId   int64 `json:"override_msg_id,omitempty"`
		BigPushDuration int   `json:"big_push_duration,omitempty"`
	}

	DeviceRequest struct {
		Tags   *DeviceRequestTags `json:"tags"`
		Alias  string             `json:"alias"`
		Mobile string             `json:"mobile"`
	}
	DeviceEmptyTagsRequest struct {
		Tags   string `json:"tags"`
		Alias  string `json:"alias"`
		Mobile string `json:"mobile"`
	}
	DeviceRequestTags struct {
		Add    []string `json:"add,omitempty"`
		Remove []string `json:"remove,omitempty"`
	}
	DeviceBindTagsRequest struct {
		Add    []string `json:"add,omitempty"`
		Remove []string `json:"remove,omitempty"`
	}
)

func getAuthorization(isGroup bool) string {
	str := APP_KEY + ":" + APP_SECRET
	if isGroup {
		str = "group-" + str
	}
	buf := []byte(str)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(buf))
}
func getUserAgent() string {
	return fmt.Sprintf("(%s) go/%s", runtime.GOOS, runtime.Version())
}

func New(timeOut time.Duration, isGroup bool) *Hyper {
	hyper := &Hyper{}
	hyper.Client = &http.Client{Timeout: timeOut * time.Second}

	hyper.UserAgent = map[string]string{"User-Agent": getUserAgent()}
	hyper.Headers = map[string]string{
		"Authorization": getAuthorization(isGroup),
		"Content-Type":  "application/json",
	}
	hyper.Params = url.Values{}

	return hyper
}
