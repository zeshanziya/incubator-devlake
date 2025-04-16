package apimodels

import "encoding/json"

type ExampleUserItem struct {
		Gender string `json:"gender"`
		Name   struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			City  string `json:"city"`
			State string `json:"state"`
		} `json:"location"`
		Email string `json:"email"`
		Dob   struct {
			Age int `json:"age"`
		} `json:"dob"`
		Phone string `json:"phone"`
		Login struct {
			Uuid string `json:"uuid"`
		} `json:"login"`
}

type ExampleUserApiResult struct {
	Results []json.RawMessage `json:"results"`
	Info json.RawMessage `json:"info"`
}
