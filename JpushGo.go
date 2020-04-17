package main

import (
	"JpushGo/JpushGo"

	"encoding/json"
	"fmt"
)

func main() {
	jpush := JpushGo.New(30, false)

	err, devices := jpush.GetDevices("123")
	if err != nil {
		fmt.Println(err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(devices, &result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
