package main

import (
	"fmt"
	"global"

	"webserver"

	"github.com/BurntSushi/toml"
	logger "github.com/shengkehua/xlog4go"
)

func initService() (err error) {

	if err := initConf(); err != nil {
		return fmt.Errorf("msg=[service init fail] detail=[init config fail] err=[%s]", err.Error())
	}

	if err := initLog(); err != nil {
		return fmt.Errorf("msg=[service init fail] detail=[init log fail] err=[%s]", err.Error())
	}

	logger.Info("global.Conf=%v", global.Conf)
	return nil
}

func initConf() error {
	_, err := toml.DecodeFile(global.CONF_FILE, &global.Conf)
	if err != nil {
		fmt.Printf("msg=[init config fail] err=[%s]\n", err.Error())
		return err
	}

	return nil
}

func initLog() error {

	logFile := global.Conf.Service.LogFile

	return logger.SetupLogWithConf(logFile)
}

func initHttpServer() {
	hPort := global.Conf.Http.Port
	go webserver.RunHttpServer(hPort)
}
