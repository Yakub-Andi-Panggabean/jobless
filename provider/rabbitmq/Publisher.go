package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"siji/sms-api/actor"
	"siji/sms-api/util"
	"strings"
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
	conn, err := amqp.Dial(generateRabbitMQUrl())
	if err != nil {
		panic("could not make connection to rabbitmq," + err.Error())
	} else {
		connection = conn
	}

}

func (m *MessagePublisher) publish(message string) error {

	var errorVal error
	exchangeName := util.GetConfig().GetString("queue.exchange.incoming")

	channel, err := connection.Channel()

	if err != nil {

		log.Warn("failed to create channel, err:", err.Error())
		errorVal = err

	} else {

		errExchange := channel.ExchangeDeclare(exchangeName, amqp.ExchangeFanout, true, false, false, false, nil)

		if errExchange != nil {

			log.Warn("failed to declare exchange, err:", errExchange.Error())
			errorVal = errExchange

		} else {

			channel.Publish(exchangeName, "", false, false, amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         []byte(message),
			})

		}

	}

	return errorVal

}

func (m *MessagePublisher) SaveMessage(u actor.UserMessageStatus) (int, error) {

	var result int

	err := m.publish(util.ConvertToJson(u))

	if err != nil {
		result = 0
	} else {
		result = 1
	}

	return result, err

}

func generateRabbitMQUrl() string {

	url := "amqp://{username}:{password}@{host}:{port}/"

	url = strings.Replace(url, "{username}", util.GetConfig().GetString("queue.user"), 1)
	url = strings.Replace(url, "{password}", util.GetConfig().GetString("queue.password"), 1)
	url = strings.Replace(url, "{host}", util.GetConfig().GetString("queue.host"), 1)
	url = strings.Replace(url, "{port}", util.GetConfig().GetString("queue.port"), 1)

	return url

}
