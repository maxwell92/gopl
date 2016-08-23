package main

import (
	//	"encoding/json"
	"fmt"
)

type Errno uintptr

const (
	EDBCONN Errno = 1 /*MySQL connection error*/
	EYCE    Errno = 2 /*Yce internal error*/
	EKUBE   Errno = 3 /*Kubernetes error*/

	ETIMEOUT Errno = 4 /*Operation time out*/

)

var errors = [...]string{
	EDBCONN:  "MySQL connection error",
	EYCE:     "Yce internal error",
	EKUBE:    "Kubernetes error",
	ETIMEOUT: "Operation time out",
}

type error interface {
	EncodeJSON()
	Error() string
}

type YceError struct {
	ErrMsg string
}

func New(errmsg string) *YceError {
	return &YceError{ErrMsg: errmsg}
}

func (e *YceError) Error() string {
	return e.ErrMsg
}

func (e Errno) Error() string {
	if 0 <= int(e) && int(e) < len(errors) {
		return errors[e]
	}
	return fmt.Sprintf("errno %d", e)
}

/*
func (e *YceError) EncodeJSON() []byte {
	return json.Marshal(e.Error)
}
*/
func main() {
	fmt.Println(New("ListDeployController error").Error())
	fmt.Println(ETIMEOUT.Error())
}
