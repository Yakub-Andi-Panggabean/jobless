package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger

func init() {

	log = logrus.New()

}

func ConvertToJson(p interface{}) string {

	b, err := json.Marshal(p)

	if err != nil {

		log.Error("failed to convert object to json ,err :", err)
		return nil

	} else {

		return string(b)

	}

}
