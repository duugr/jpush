package JpushGo

import (
	"strings"
)

// 查询设备的别名与标签
func (hyper *Hyper) GetDevices(registrationId string) (error, []byte) {
	url := strings.ReplaceAll(get_devices, "{scheme}", URL_DEVICES)
	url = strings.ReplaceAll(url, "{registration_id}", registrationId)
	hyper.Url = url
	err := hyper.Get()
	if err != nil {
		return err, nil
	}

	return err, hyper.Result
}

// 设置设备的别名与标签
func (hyper *Hyper) PostDevices(registrationId string) (error, []byte) {
	url := strings.ReplaceAll(get_devices, "{scheme}", URL_DEVICES)
	url = strings.ReplaceAll(url, "{registration_id}", registrationId)
	hyper.Url = url
	err := hyper.Post()
	if err != nil {
		return err, nil
	}

	return err, hyper.Result
}
