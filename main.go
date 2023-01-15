package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func main() {

	warnings, forecasts := CreateRuntimes()

	jsonStr, _ := json.MarshalIndent(warnings, "", "  ")
	_ = forecasts

	log.Infof("json=\n%s", string(jsonStr))

}
