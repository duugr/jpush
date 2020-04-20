package JpushGo

const (
	ScheduleDay   = "day"
	ScheduleWeek  = "week"
	ScheduleMonth = "month"

	SCHEDULE_POST = "https://api.jpush.cn/v3/schedules"
	// 获取有效的 Schedule 列表
	SCHEDULE_GET = "https://api.jpush.cn/v3/schedules?page={page}"
	// 获取指定的定时任务
	SCHEDULE_ID_GET = "https://api.jpush.cn/v3/schedules/{schedule_id}"
	// 获取定时任务对应的所有 msg_id
	SCHEDULE_ID_MSG_GET = "https://api.jpush.cn/v3/schedules/{schedule_id}/msg_ids"
	// 修改指定的 Schedule
	SCHEDULE_ID_PUT = "https://api.jpush.cn/v3/schedules/{schedule_id}"
	// 删除指定的 Schedule 任务
	SCHEDULE_ID_DELETE = "https://api.jpush.cn/v3/schedules/{schedule_id}"
)

type (
	ScheduleRequest struct {
		Cid     string           `json:"cid,omitempty"`
		Name    string           `json:"name"`
		Enabled bool             `json:"enabled"`
		Push    *PushRequest     `json:"push"`
		Trigger *ScheduleTrigger `json:"trigger"`
	}

	ScheduleTrigger struct {
		Single     *ScheduleTriggerSingle     `json:"single,omitempty"`
		Periodical *ScheduleTriggerPeriodical `json:"periodical,omitempty"`
	}

	ScheduleTriggerSingle struct {
		Timer string `json:"time,omitempty"`
	}

	ScheduleTriggerPeriodical struct {
		Start     string      `json:"start,omitempty"`
		End       string      `json:"end,omitempty"`
		Time      string      `json:"time,omitempty"`
		TimeUnit  string      `json:"time_unit,omitempty"`
		Frequency int         `json:"frequency,int,omitempty"`
		Point     interface{} `json:"point,omitempty"`
	}
)
