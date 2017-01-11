package controller

import (
	myerror "tools/error"
	"github.com/kataras/iris"
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
	log.Errorf("Controller Response YceError: controller=%p, code=%d, msg=%s", c, c.Ye.Code, myerror.Errors[c.Ye.Code].LogMsg)
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
	c.Ye = myerror.NewYceError(myerror.EOK, msg)
	c.Response.Header.Set("Cache-Control", "no-store")
	log.Infof("Controller Response OK: controller=%p, code=%d, msg=%s", c, c.Ye.Code, myerror.Errors[c.Ye.Code].LogMsg)
	c.Write(c.Ye.String())

}