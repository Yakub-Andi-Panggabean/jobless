package util

import (
	"crypto/md5"
	"encoding/hex"
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

func GetMd5Hash(text string) string {

	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))

}
