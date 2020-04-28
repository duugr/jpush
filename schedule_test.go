package JpushGo

import (
	"testing"
)

// var jpush = New("7d431e42dfa6a6d693ac2d04", "5e987ac6d2e04d95a9d8f0d1", 30, false)

func TestScheduleCreateTask(t *testing.T) {
	req := NewSchedule(true)
	req.AddName("test")
	req.AddEnabled(true)
	req.AddTimer("2020-04-30 10:00:00")
	req.Push.AddCid(cidString)
	req.Push.PlatformAll()
	req.Push.AddTitle("Title")
	req.Push.AddAlert("Content")
	req.Push.AddContentType("text")
	req.Push.AddTimeToLive(60)
	req.Push.AddApnsCollapseId("jpush_user_158803817123456")
	req.Push.AudienceAll()
	t.Logf("ScheduleCreateTask req : %+v", req)
	err, result := jpush.ScheduleCreateTask(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
