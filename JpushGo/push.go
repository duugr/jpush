package JpushGo

import (
	"strconv"
	"strings"
)

// cid：推送唯一标识符
func (hyper *Hyper) GetCid(count int, citType string) (error, map[string][]string) {
	hyper.Url = strings.ReplaceAll(PUSH_CID_GET, "{count}", strconv.Itoa(count))
	hyper.Url = strings.ReplaceAll(hyper.Url, "{type}", citType)
	err := hyper.Get()
	if err != nil {
		return err, nil
	}

	return hyper.MapStrings()
}

func (hyper *Hyper) Push(push *PushRequest, isGroup bool) (error, map[string]interface{}) {
	if isGroup {
		// 应用分组推送
		hyper.Url = PUSH_GROUP_POST
	} else {
		hyper.Url = PUSH_POST
	}

	err := hyper.Post(push)
	if err != nil {
		return err, nil
	}

	return hyper.Map()
}

func (hyper *Hyper) Validate(req *PushRequest) (error, map[string]interface{}) {
	hyper.Url = PUSH_VALIDATE_POST
	err := hyper.Post(req)
	if err != nil {
		return err, nil
	}

	return hyper.Map()
}
