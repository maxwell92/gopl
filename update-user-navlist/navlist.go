package main

import (
	"encoding/json"
	"log"
)

type NavList struct {
	List []NavElements `json:"list"`
}

type NavElements struct {
	HashKey      string `json:"$$hashKey"`
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	State        string `json:"state"`
	IncludeState string `json:"includeStat"`
	Href         string `json:"href"`

	ClassName string         `json:"className,omitempty"`
	Operation interface{}    `json:"operation,omitempty"`
	Item      []*NavElements `json:"item,omitempty"`
}

type OldOperationConfig bool

type NewOperationConfig struct {
	Handle bool `json:"handle"`
	Delete bool `json:"delete"`
}

func DecodeNavList(navList string) *NavList {
	nl := new(NavList)

	err := json.Unmarshal([]byte(navList), nl)
	if err != nil {
		log.Printf("DecodeNavList Error: navList=%s, error=%s\n", navList, err)
		panic(err)
	}

	return nl
}

func (nl *NavList) EncodeNavList() string {
	data, err := json.Marshal(nl)
	if err != nil {
		log.Printf("EncodeNavList Error: error=%s\n", err)
		panic(err)
	}

	return string(data)
}
