package repository

import (
	"encoding/json"
	"io/ioutil"
)

func loadData(f string, vv interface{}) error {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, &vv)
	if err != nil {
		return err
	}
	return nil
}

func saveData(f string, vv interface{}) error {
	raw, err := json.MarshalIndent(vv, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f, raw, 0644)
}
