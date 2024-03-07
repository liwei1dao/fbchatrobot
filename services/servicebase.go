package services

import (
	"github.com/liwei1dao/lego/base/single"
)

type ServiceBase struct {
	single.SingleService
}

func (this *ServiceBase) InitSys() {
	this.SingleService.InitSys()
}
