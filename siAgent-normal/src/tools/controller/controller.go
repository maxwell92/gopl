package controller

import (
	"github.com/kataras/iris"
	myerror "tools/error"
)

type Controller struct {
	*iris.Context
	Ye *myerror.YceError
}

type IController interface {
	WriteError()
	CheckError() bool // if error true, else false
	WriteOk()
}

func (c *Controller) WriteError() {
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Response.Header.Set("Access-Control-Allow-Methods", "POST")
	c.Response.Header.Set("Access-Control-Allow-Methods", "GET")
	c.Response.Header.Set("Access-Control-Allow-Headers", "x-requested-with, content-type")
	c.Response.Header.Set("Cache-Control", "no-store")
	c.Write(c.Ye.String())
}

func (c *Controller) CheckError() bool {
	if c.Ye != nil {
		c.WriteError()
		return true
	}
	return false
}

func (c *Controller) WriteOk(msg string) {
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Response.Header.Set("Access-Control-Allow-Methods", "POST")
	c.Response.Header.Set("Access-Control-Allow-Methods", "GET")
	c.Response.Header.Set("Access-Control-Allow-Headers", "x-requested-with, content-type")
	c.Ye = myerror.NewYceError(myerror.EOK, msg)
	c.Response.Header.Set("Cache-Control", "no-store")
	c.Write(c.Ye.String())

}
