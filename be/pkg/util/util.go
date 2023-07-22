package util

import (
	"encoding/json"
	"os"
)

type CommonResponse struct {
	Success bool        `json:"success"`
	Reason  string      `json:"reason,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ConstructErrResponse(err error) CommonResponse {
	return CommonResponse{
		Success: false,
		Reason:  err.Error(),
	}
}

func ConstructSuccessResponse(data interface{}) CommonResponse {
	return CommonResponse{
		Success: true,
		Data:    data,
	}
}

func LoadJSONFile(file string, res interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &res)
}
