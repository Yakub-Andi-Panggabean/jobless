package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"siji/sms-api/actor"
	"siji/sms-api/provider/mysql"
	"siji/sms-api/provider/rabbitmq"
	"siji/sms-api/usecase"
	"siji/sms-api/util"
	"strconv"
)

var log *logrus.Logger

var usecaseFactory usecase.Factory

func init() {

	publisher := rabbitmq.MessagePublisher{
		Password: util.GetConfig().GetString("queue.user"),
		Username: util.GetConfig().GetString("queue.password"),
		Address:  util.GetConfig().GetString("queue.address"),
		Port:     util.GetConfig().GetInt("queue.port"),
	}

	storageFactory := mysql.NewStorage()

	usecaseFactory = usecase.New(publisher, storageFactory)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome to sms api")
}

func IncomingSmsHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	messageType, err := strconv.ParseInt(q.Get("type"), 10, 8)

	if err != nil {

		messageType = 0

	}

	request := usecase.MessageRequest{
		Username:       q.Get("g3p4i"),
		Password:       q.Get("G4PIpw"),
		Destination:    q.Get("dst"),
		SenderId:       q.Get("src"),
		Type:           actor.MessageType(messageType),
		MessageContent: q.Get("msg"),
		ApiVersion:     "",
		IsSplitSms:     false,
		Id:             q.Get("ID"),
		IPOrigin:       r.RemoteAddr,
	}

	messageUsecase := usecaseFactory.NewMessageUsecase()

	messageUsecase.SendMessage(request)

}

func VersatileSmsHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	messageType, err := strconv.ParseInt(q.Get("type"), 10, 8)

	if err != nil {

		messageType = 0

	}

	request := usecase.MessageRequest{
		Username:       q.Get("g3p4i"),
		Password:       q.Get("G4PIpw"),
		Destination:    q.Get("dst"),
		SenderId:       q.Get("src"),
		Type:           actor.MessageType(messageType),
		MessageContent: q.Get("msg"),
		ApiVersion:     "",
		IsSplitSms:     false,
		Id:             q.Get("ID"),
		IPOrigin:       r.RemoteAddr,
	}

	messageUsecase := usecaseFactory.NewMessageUsecase()

	messageUsecase.SendMessage(request)

}
