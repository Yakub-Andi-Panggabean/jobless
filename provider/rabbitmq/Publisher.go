package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"siji/sms-api/actor"
	"siji/sms-api/util"
)

var log logrus.Logger

type (
	MessagePublisher struct {
		Address  string
		port     int
		Username string
		Password string
	}
)

var connection *amqp.Connection

func init() {
	log = logrus.New()
	conn, err := amqp.Dial("")
	if err != nil {
		panic("could not make connection to rabbitmq," + err.Error())
	} else {
		connection = conn
	}

}

func (m *MessagePublisher) publish(message string) {

	exchangeName := util.GetConfig().GetString("queue.exchange.incoming")

	channel, err := connection.Channel()

	if err != nil {

		log.Warn("failed to create channel, err:", err.Error())

	} else {

		errExchange := channel.ExchangeDeclare(exchangeName, amqp.ExchangeFanout, true, false, false, false, nil)

		if errExchange != nil {

			log.Warn("failed to declare exchange, err:", err.Error())

		} else {

			channel.Publish(exchangeName, "", false, false, amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         []byte(message),
			})

		}

	}

}

func (m *MessagePublisher) save(u actor.UserMessageStatus) (int, error) {

	var err error
	var result int

	return result, err

}
