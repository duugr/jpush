package JpushGo

import (
	"encoding/json"
	"testing"
)

// var jpush = New("7d431e42dfa6a6d693ac2d04", "5e987ac6d2e04d95a9d8f0d1", 30, false)
var jpush = New("697c98b25a2920178b66cd0d", "3caabeeecdb2533776a51394", 30, false)

func TestGetDevices(t *testing.T) {
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
func TestGetAliases(t *testing.T) {
	err, data := jpush.GetAliases("123")
	if err != nil {
		t.Error(err)
		return
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
func TestPostDevices(t *testing.T) {
	tags := &DeviceTagsRequest{
		Add: []string{"test"},
	}
	req := &DeviceRequest{
		Alias:  "xialei",
		Mobile: "13333333333",
		Tags:   tags,
	}

	err, devices := jpush.PostDevices("123", req)
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
