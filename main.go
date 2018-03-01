package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"siji/sms-api/adapter/server"
	"siji/sms-api/util"
)

var (
	log                     *logrus.Logger
	versatile               = "sendSmsVersatile.do"
	latin                   = "sendSmsLatin.do"
	latinConcat             = "sendSmsLatinConcat.do"
	wapPush                 = "sendSmsWapPush.do"
	broadcast               = "sendSmsBroadcast.do"
	deliveryStatus          = "smsDeliveryStatus.do"
	versatileDeliveryStatus = "versatileSmsDeliveryStatus.do"
	fetchReply              = "fetchReply.do"
)

func init() {

	log = logrus.New()
}

func main() {

	router := mux.NewRouter()

	prefix := util.GetConfig().GetString("service.prefix")

	log.Infof("prefix : %s",prefix)

	router.HandleFunc("/", server.IndexHandler)
	router.HandleFunc(prefix+versatile, server.IncomingSmsHandler)
	router.HandleFunc(prefix+latin, server.IncomingSmsHandler)
	router.HandleFunc(prefix+latinConcat, server.IncomingSmsHandler)
	router.HandleFunc(prefix+wapPush, server.IncomingSmsHandler)
	router.HandleFunc(prefix+broadcast, server.IncomingSmsHandler)
	router.HandleFunc(prefix+deliveryStatus, server.IncomingSmsHandler)
	router.HandleFunc(prefix+versatileDeliveryStatus, server.IncomingSmsHandler)
	router.HandleFunc(prefix+fetchReply, server.IncomingSmsHandler)

	http.ListenAndServe(":8070", router)

}
