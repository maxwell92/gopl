package error

import (
	"encoding/json"
	mylog "tools/mylog"
)

var log = mylog.Log

type YceError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func NewYceError(code int32, data string) *YceError {
	return &YceError{
		Code:    code,
		Message: Errors[code].ErrMsg,
		Data:    data,
	}
}

func (ye *YceError) DecodeJson(data string) error {
	err := json.Unmarshal([]byte(data), ye)

	if err != nil {
		log.Errorf("YceError DecodeJson Error: err=%s", err)
		return err
	}

	return nil
}

func (ye *YceError) EncodeJson() (string, error) {
	data, err := json.Marshal(ye)

	if err != nil {
		log.Errorf("YceError EncodeJson Error: err=%s", err)
		return "", err
	}
	return string(data), nil
}

func (ye *YceError) String() string {
	errJSON, _ := json.Marshal(ye)
	return string(errJSON)
}

func (ye *YceError) SetError(code int32, message, data string) {
	ye.Code = code
	ye.Message = message
	ye.Data = data
}

func (ye *YceError) GetCode() int32 {
	return ye.Code
}

func (ye *YceError) GetMessage() string {
	return ye.Message
}

func (ye *YceError) GetData() string {
	return ye.Data
}
