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
	PUSH_DELETE = "https://api.jpush.cn/v3/push/{msgid}"

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

)

type ()
