package helper

import (
	"auto_healer/configs"
	"auto_healer/internal/pkg/env"
	"encoding/json"
	"strings"

	"github.com/common-nighthawk/go-figure"

	log "logger"
)

var (
	servicePackageName = env.GetEnv("SERVICE_NAME", configs.SERVICE_NAME)
	serverDebugLevel   = env.GetEnv("DEBUG_LEVEL", configs.DEBUG_LEVEL)
)

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := perror
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	log.Debug().Msgf("%s \n", p)
}

func GetServicePackageName() string {
	return servicePackageName
}

func ShowServicelogoPrint() {
	serviceLogo := strings.ToUpper("hooker-client")
	myFigure := figure.NewColorFigure(serviceLogo, "doom", "cyan", true)
	myFigure.Print()
}
