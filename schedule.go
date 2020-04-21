package JpushGo

import (
	"strconv"
	"strings"
)

func (hyper *Hyper) ScheduleCreateTask(push *ScheduleRequest) (error, map[string]interface{}) {
	hyper.Url = SCHEDULE_POST

	err := hyper.Post(push)
	if err != nil {
		return err, nil
	}

	return hyper.Map()
}

func (hyper *Hyper) ScheduleGetList(page int) (error, map[string]interface{}) {
	hyper.Url = strings.ReplaceAll(SCHEDULE_GET, "{page}", strconv.Itoa(page))
	err := hyper.Get()
	if err != nil {
		return err, nil
	}

	return hyper.Map()
}

func (hyper *Hyper) ScheduleView(id string) (error, map[string]interface{}) {
	hyper.Url = strings.ReplaceAll(SCHEDULE_ID_GET, "{schedule_id}", id)
	err := hyper.Get()
	if err != nil {
		return err, nil
	}

	return hyper.Map()
}

func (hyper *Hyper) ScheduleUpdate(id string, req *ScheduleRequest) (error, map[string]interface{}) {
	hyper.Url = strings.ReplaceAll(SCHEDULE_ID_PUT, "{schedule_id}", id)

	err := hyper.Put(req)
	if err != nil {
		return err, nil
	}

	return hyper.Map()
}

func (hyper *Hyper) ScheduleDelete(id string) (error, []byte) {
	hyper.Url = strings.ReplaceAll(SCHEDULE_ID_DELETE, "{schedule_id}", id)

	err := hyper.Delete()
	if err != nil {
		return err, nil
	}

	return nil, hyper.Result
}
