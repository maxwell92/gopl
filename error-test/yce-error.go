package main

type log interface {
}

type error interface {
}

type YceError struct {
	Err error
	Log log
}

type YceLog struct {
	Logmsg string
	Who    string
}
