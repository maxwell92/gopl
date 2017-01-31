package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now() // Time
	current.Format("xxx")
	current.Unix("...")

	current.Minutes() // wrong, different type. current: Time. Minutes(): Duration

	time.Now().Format("xxx")
	time.Now().Unix("...")

	time.Now().Unix("2006-01-02 15:04:05")

}
