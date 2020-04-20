package JpushGo

import (
	"testing"
)

// var jpush = New("7d431e42dfa6a6d693ac2d04", "5e987ac6d2e04d95a9d8f0d1", 30, false)

func TestScheduleCreateTask(t *testing.T) {
	req := &ScheduleRequest{
		Name:    "test",
		Enabled: true,
		Trigger: &ScheduleTrigger{
			Single: &ScheduleTriggerSingle{
				Timer: "2017-11-04 10:00:00",
			},
		},
		Push: getMsg(),
	}

	err, result := jpush.ScheduleCreateTask(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
