package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {

	logger = logrus.New()

}

func ConvertToJson(p interface{}) string {

	b, err := json.Marshal(p)

	if err != nil {

		logger.Error("failed to convert object to json ,err :", err)
		return ""

	} else {

		return string(b)

	}

}
