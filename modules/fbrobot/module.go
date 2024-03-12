package fbrobot

import (
	"fbrobot/comm"

	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/gin"
)

func NewModule() core.IModule {
	m := new(FBRobot)
	return m
}

type FBRobot struct {
	cbase.ModuleBase
	service base.IClusterService
	options *Options
	api     *apiComp
	gin     gin.ISys
}

func (this *FBRobot) GetType() core.M_Modules {
	return comm.SM_FBRobotModule
}

func (this *FBRobot) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

func (this *FBRobot) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	this.options = options.(*Options)
	if err = this.ModuleBase.Init(service, module, options); err != nil {
		return
	}
	if this.gin, err = gin.NewSys(
		gin.SetListenPort(this.options.ListenPort),
		gin.SetLetEncrypt(this.options.LetEncrypt),
		gin.SetDomain(this.options.DomainName),
		gin.SetDebug(this.options.Debug),
	); err != nil {
		return
	}
	return
}

func (this *FBRobot) OnInstallComp() {
	this.ModuleBase.OnInstallComp()
	this.api = this.RegisterComp(new(apiComp)).(*apiComp)
}
