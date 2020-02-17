package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBodyJson(holder interface{}, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return err
	}

	if err = json.Unmarshal(body, &holder); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return err
	}

	return nil
}
