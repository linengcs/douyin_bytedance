package services

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var MqService = newMqService()

func newMqService() *mqService {
	return &mqService{}
}

type mqService struct {
}

type Callback func(msg string)

func (s *mqService) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

func (s *mqService) Publish(exchange string, queueName string, body string) error {
	conn, err := MqService.Connect()

	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return err
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return err
}

func (s *mqService) Consumer(exchange string, queueName string, callback Callback) {
	conn, err := MqService.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		log.Println(err)
		return
	}

	q, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_ = channel.Qos(
		// 每次队列只消费一个消息 这个消息处理不完服务器不会发送第二个消息过来
		// 当前消费者一次能接受的最大消息数量
		1,
		// 服务器传递的最大容量
		0,
		// 如果为true 对channel可用 false则只对当前队列可用
		false,
	)

	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false)
		}
	}()
	fmt.Println("Waiting for messages")
	<-forever
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
