package JpushGo

const (
	// 获取当前设备的所有属性，包含 tags, alias，手机号码 mobile。
	DERVICE_GET = "https://device.jpush.cn/v3/devices/{registration_id}"
	// 设置设备的别名与标签
	DERVICE_POST = "https://device.jpush.cn/v3/devices/{registration_id}"
	// 获取用户在线状态（VIP 专属接口）
	DERVICE_STATUS_POST = "https://device.jpush.cn/v3/devices/status/"

	// 获取指定 alias 下的设备，最多输出 10 个；
	ALIAS_GET = "https://device.jpush.cn/v3/aliases/{alias_value}?platform=android,ios"
	// 删除一个别名，以及该别名与设备的绑定关系。
	ALIAS_DELETE = "https://device.jpush.cn/v3/aliases/{alias_value}"
	// 批量解绑设备与别名之间的关系。
	ALIAS_POST = "https://device.jpush.cn/v3/aliases/{alias_value}"

	// 查询标签列表
	TAGS_GET = "https://device.jpush.cn/v3/tags/"
	// 判断设备与标签绑定关系,查询某个设备是否在 tag 下。
	TAGS_REG_GET = "https://device.jpush.cn/v3/tags/{tag_value}/registration_ids/{registration_id}"
	// 更新标签,为一个标签添加或者删除设备。
	TAGS_POST = "https://device.jpush.cn/v3/tags/{tag_value}"
	// 删除一个标签，以及标签与设备之间的关联关系。
	TAGS_DELETE = "https://device.jpush.cn/v3/tags/{tag_value}"
)

type (
	DeviceRequest struct {
		Tags   *DeviceTagsRequest `json:"tags"`
		Alias  string             `json:"alias"`
		Mobile string             `json:"mobile"`
	}
	DeviceEmptyTagsRequest struct {
		Tags   string `json:"tags"`
		Alias  string `json:"alias"`
		Mobile string `json:"mobile"`
	}
	DeviceTagsRequest struct {
		Add    []string `json:"add,omitempty"`
		Remove []string `json:"remove,omitempty"`
	}
)
