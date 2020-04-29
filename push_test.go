package JpushGo

import (
	"testing"
)

// var jpush = New("7d431e42dfa6a6d693ac2d04", "5e987ac6d2e04d95a9d8f0d1", 30, false)

var cidString string

func TestPushGetCid(t *testing.T) {
	err, result := jpush.GetCid(1, "push")
	if err != nil {
		t.Error(err)
		return
	}

	cidString = result["cidlist"][0]

	t.Log(result)
}

func getMsg(isAll bool) *PushRequest {
	params := make(map[string]interface{})
	params["url"] = "https://www.jpush.cn"

	req := NewPush("Title", "Content")
	req.AddCid(cidString)
	req.PlatformAll()
	req.AddExtras(params)
	req.AddContentType("text")
	req.AddTimeToLive(60)
	req.AddApnsCollapseId("jpush_user_158803817123456")
	if isAll {
		req.AudienceAll()
	} else {
		req.AddAlias([]string{"697c98b25a2920178b66cd0d"})
	}

	return req
}

func TestPushMessage(t *testing.T) {
	req := getMsg(false)
	t.Logf("req %+v, %#v", req, req)
	err, result := jpush.Push(req, false)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
