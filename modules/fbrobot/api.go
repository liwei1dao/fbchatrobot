package fbrobot

import (
	"fbrobot/sys/fbmessager"
	"net/http"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/gin/engine"
)

///API组件
type apiComp struct {
	cbase.ModuleCompBase
	module *FBRobot
}

func (this *apiComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.module = module.(*FBRobot)
	return
}

func (this *apiComp) Start() (err error) {
	err = this.ModuleCompBase.Start()
	this.module.gin.GET("/test", this.Test)
	this.module.gin.Any("/webhook", this.Webhook)
	return
}

//验证码
func (this *apiComp) Test(c *engine.Context) {
	c.JSON(http.StatusOK, "hello! you nice")
}

//验证码
func (this *apiComp) Webhook(c *engine.Context) {
	fbmessager.Handler(c.Writer, c.Request)
}
