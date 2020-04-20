package JpushGo

import (
	"encoding/json"
	"testing"
)

// var jpush = New("7d431e42dfa6a6d693ac2d04", "5e987ac6d2e04d95a9d8f0d1", 30, false)

func TestPush(t *testing.T) {
	err, devices := jpush.GetDevices("123")
	if err != nil {
		t.Error(err)
		return
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(devices, &result)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
