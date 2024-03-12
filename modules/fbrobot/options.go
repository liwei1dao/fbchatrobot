package fbrobot

import (
	"errors"

	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type (
	Options struct {
		ListenPort int      //监听端口
		LetEncrypt bool     //是否支持LetEncrypt 证书
		DomainName []string //域名
		Debug      bool     //日志是否开启
		Log        log.ILogger
	}
)

func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	err = mapstructure.Decode(settings, this)
	if this.Log = log.NewTurnlog(this.Debug, log.Clone("", 4)); this.Log == nil {
		err = errors.New("log is nil")
	}
	return
}
