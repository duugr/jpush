package JpushGo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type Hyper struct {
	Key       string
	Secret    string
	Url       string
	UserAgent map[string]string
	Headers   map[string]string
	Client    *http.Client

	Result []byte
}

func New(key, secret string, timeOut time.Duration, isGroup bool) *Hyper {
	hyper := &Hyper{}
	hyper.Key = key
	hyper.Secret = secret

	hyper.UserAgent = map[string]string{"User-Agent": hyper.getUserAgent()}
	hyper.Headers = map[string]string{
		"Authorization": hyper.getAuthorization(isGroup),
		"Content-Type":  "application/json",
	}
	hyper.Client = &http.Client{Timeout: timeOut * time.Second}

	return hyper
}
func (hyper *Hyper) getAuthorization(isGroup bool) string {
	str := hyper.Key + ":" + hyper.Secret
	if isGroup {
		str = "group-" + str
	}
	buf := []byte(str)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(buf))
}
func (hyper *Hyper) getUserAgent() string {
	return fmt.Sprintf("(%s) go/%s", runtime.GOOS, runtime.Version())
}

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

func (hyper *Hyper) Post(data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", hyper.Url, bytes.NewReader(buf))
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

func (hyper *Hyper) Put(data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", hyper.Url, bytes.NewReader(buf))
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

func (hyper *Hyper) Delete() error {
	req, err := http.NewRequest("DELETE", hyper.Url, nil)
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
func (hyper *Hyper) MapStrings() (error, map[string][]string) {
	result := make(map[string][]string)
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
