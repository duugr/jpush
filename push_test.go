package JpushGo

import (
	"testing"
)

// var jpush = New("7d431e42dfa6a6d693ac2d04", "5e987ac6d2e04d95a9d8f0d1", 30, false)

func getMsg() *PushRequest {
	params := make(map[string]interface{})
	params["url"] = "https://www.jpush.cn"
	req := &PushRequest{
		Cid:      "7d431e42dfa6a6d693ac2d04-11ea3aaf-6b80-4acc-9b97-70cc2085e03b",
		Platform: Platform_All,
		Audience: &PushAudience{
			RegistrationId: []string{"0815fd9e991"},
		},
		Message: &PushMessage{
			MsgContent:  "Message Content",
			Title:       "Message Title",
			ContentType: "text",
			Extras:      params,
		},
		Options: &PushOptions{
			TimeToLive:     60,
			ApnsCollapseId: "jiguang_test_201706011100",
			ApnsProduction: true,
		},
	}
	return req
}

func TestPushGetCid(t *testing.T) {
	err, result := jpush.GetCid(1, "push")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
