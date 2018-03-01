package server

import (
	"net/http"
	"siji/sms-api/actor"
	"siji/sms-api/provider/mysql"
	"siji/sms-api/provider/rabbitmq"
	"siji/sms-api/usecase"
)

var usecaseFactory usecase.Factory

func init() {

	publisher := rabbitmq.MessagePublisher{
		Password: "",
		Username: "",
		Address:  "",
		Port:     15271,
	}

	storageFactory := mysql.NewStorage()

	usecaseFactory = usecase.New(publisher, storageFactory)

}

func IncomingSmsHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	request := usecase.MessageRequest{
		Username:       q.Get(""),
		Password:       q.Get(""),
		Broadcast:      nil,
		Destination:    q.Get(""),
		SenderId:       q.Get(""),
		Type:           actor.LATIN,
		MessageContent: q.Get(""),
		ApiVersion:     nil,
		IsSplitSms:     false,
		Id:             q.Get(""),
		IPOrigin:       q.Get(""),
	}

	messageUsecase := usecaseFactory.NewMessageUsecase()

	messageUsecase.SendMessage(request)

}
