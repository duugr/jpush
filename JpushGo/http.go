package JpushGo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (hyper *Hyper) Do(req *http.Request) (err error) {
	if len(hyper.UserAgent) > 0 {
		for key, value := range hyper.UserAgent {
			req.Header.Add(key, value)
		}
	}

	if len(hyper.Headers) > 0 {
		for key, value := range hyper.Headers {
			req.Header.Set(key, value)
		}
	}
	fmt.Println(req)
	resp, err := hyper.Client.Do(req)
	if err != nil {
		fmt.Printf("hyper.Client : %v", err)
		return
	}
	defer resp.Body.Close()

	hyper.Result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("hyper ioutil.ReadAll : %v", err)
	}
	return
}

func (hyper *Hyper) Post() error {

	data := bytes.NewReader([]byte(hyper.Params.Encode()))

	req, err := http.NewRequest("POST", hyper.Url, data)
	if err != nil {
		fmt.Printf("hyper http.NewRequest Error: %#v", err)
		return err
	}

	err = hyper.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (hyper *Hyper) Get() error {
	req, err := http.NewRequest("GET", hyper.Url, nil)
	if err != nil {
		fmt.Printf("hyper http.NewRequest Error: %#v", err)
		return err
	}

	err = hyper.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (hyper *Hyper) Array() (error, []interface{}) {
	list := make([]interface{}, 0)
	err := json.Unmarshal(hyper.Result, &list)
	return err, list
}

func (hyper *Hyper) Map() (error, map[string]interface{}) {
	result := make(map[string]interface{})
	err := json.Unmarshal(hyper.Result, &result)
	return err, result
}

// func (hyper *Hyper) jsonDecode() Result {
// 	var info Result

// 	decoder := json.NewDecoder(bytes.NewReader(hyper.Result))
// 	decoder.UseNumber()

// 	err := decoder.Decode(&info)
// 	if err != nil {
// 		fmt.Printf("jsonDecode : %v", err)
// 		fmt.Printf("jsonDecode : %T", err)
// 		fmt.Printf("jsonDecode : %s", err)
// 	}

// 	return info
// }
