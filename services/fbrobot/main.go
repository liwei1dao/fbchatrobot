package main

import (
	"fbrobot/modules/fbrobot"
	"fbrobot/services"
	"fbrobot/sys/fbmessager"
	"flag"
	"fmt"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/single"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/log"
)

var (
	conf = flag.String("conf", "./conf/fbrobot.yaml", "获取需要启动的服务配置文件")
)

func main() {
	flag.Parse()
	s := NewService(
		single.SetConfPath(*conf),
		single.SetVersion("1.0.0"),
	)
	s.OnInstallComp( //装备组件
	)
	lego.Run(s, //运行模块
		fbrobot.NewModule(),
	)
}

func NewService(ops ...single.Option) core.IService {
	s := new(Service)
	s.Configure(ops...)
	return s
}

type Service struct {
	services.ServiceBase
}

func (this *Service) InitSys() {
	this.ServiceBase.InitSys()

	if err := fbmessager.OnInit(this.GetSettings().Sys["fbmessager"]); err != nil {
		panic(fmt.Sprintf("init sys.fbmessager err: %s", err.Error()))
	} else {
		log.Infof("init sys.fbmessager success!")
	}
}
