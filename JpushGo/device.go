package JpushGo

import (
	"strings"
)

// 查询设备的别名与标签
func (hyper *Hyper) GetDevices(registrationId string) (error, []byte) {
	hyper.Url = strings.ReplaceAll(DERVICE_GET, "{registration_id}", registrationId)
	err := hyper.Get()
	if err != nil {
		return err, nil
	}

	return err, hyper.Result
}

// 设置设备的别名与标签
func (hyper *Hyper) PostDevices(registrationId string, req *DeviceRequest) (error, []byte) {
	hyper.Url = strings.ReplaceAll(DERVICE_POST, "{registration_id}", registrationId)

	err := hyper.Post(req)
	if err != nil {
		return err, nil
	}

	return err, hyper.Result
}

// 查询别名
func (hyper *Hyper) GetAliases(aliasValue string) (error, []byte) {
	hyper.Url = strings.ReplaceAll(ALIAS_GET, "{alias_value}", aliasValue)

	err := hyper.Get()
	if err != nil {
		return err, nil
	}

	return err, hyper.Result
}
