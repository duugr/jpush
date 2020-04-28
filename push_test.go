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

func getMsg() *PushRequest {
	params := make(map[string]interface{})
	params["url"] = "https://www.jpush.cn"

	content := "Message Content"
	title := "Message Title"

	req := &PushRequest{
		Cid:      cidString,
		Platform: PlatformAll,
		Audience: &PushAudience{
			Alias: []string{"697c98b25a2920178b66cd0d"},
		},
		Notification: &PushNotification{
			Alert: content,
			Android: &NotificationAndroid{
				Alert:  content,
				Title:  title,
				Extras: params,
			},
			IOS: &NotificationIOS{
				Alert:  content,
				Extras: params,
			},
		},
		Message: &PushMessage{
			MsgContent:  content,
			Title:       title,
			ContentType: "text",
			Extras:      params,
		},
		Options: &PushOptions{
			TimeToLive:     60,
			ApnsCollapseId: "jpush_user_158803817123456",
			// ApnsProduction: false,
		},
	}
	return req
}

func TestPushMessage(t *testing.T) {
	req := getMsg()
	t.Logf("req %+v, %#v", req, req)
	err, result := jpush.Push(req, false)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
