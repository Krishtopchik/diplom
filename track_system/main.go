package main

import (
	"./service"
	"flag"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := service.InitLog()
	if err != nil {
		log.Error("Cannot init logs: ", err)
	}
	log.Info("Starting...")
	var cnf = flag.String("c", "conf.json", "Config file name (in a current dir)")
	flag.Parse()
	log.Info("Loading with config: ", *cnf)

	service.InitConfiguration(*cnf, service.Conf)

	service.RunRest()
}


