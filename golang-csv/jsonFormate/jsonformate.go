package jsonFormate

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-csv/help"
	"golang-csv/modal"
	"io/ioutil"
	"os"
)

// write the data in file
func GetJsonFile(key string, users modal.Responder) error {
	userdata, _ := json.Marshal(users)
	filename := fmt.Sprintf("jsonfile/%s.json", key)
	err := ioutil.WriteFile(filename, userdata, 0644)
	err = help.WriteJsonError(err)
	return err
}

//read the data in file
func SetJsonFile(key string, user modal.Responder) (interface{}, error) {
	filename := fmt.Sprintf("jsonfile/%s.json", key)

	filedetails, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("this file not founded")
	}

	var usersDetails interface{}
	switch v := user.(type) {
	case []modal.FullDetails:
		err = json.Unmarshal(filedetails, &v)
		if err != nil {
			return nil, errors.New("error unmarshalling JSON")
		}
		if len(v) < 1 {
			return nil, errors.New(" empty file cannot read")
		}
		usersDetails = v
	case []modal.EmployeeDetails:
		err = json.Unmarshal(filedetails, &v)
		if err != nil {
			return nil, errors.New("error unmarshalling JSON")
		}
		if len(v) < 1 {
			return nil, errors.New(" empty file cannot read")
		}
		usersDetails = v
	case []modal.ManagerDetails:
		err = json.Unmarshal(filedetails, &v)
		if err != nil {
			return nil, errors.New("error unmarshalling JSON")
		}
		if len(v) < 1 {
			return nil, errors.New(" empty file cannot read")
		}
		usersDetails = v
	default:
		return nil, errors.New("unknown user type")
	}
	err = os.Remove(filename)
	if err != nil {
		return nil, errors.New("error to removing file")
	}
	return usersDetails, nil
}
