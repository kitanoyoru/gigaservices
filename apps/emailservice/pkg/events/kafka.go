package events

import "github.com/IBM/sarama"

// REFACTOR: interface

type KafkaProducer struct {
	client sarama.SyncProducer
}

func NewKafkaProducer(brokersUrl []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	client, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		client,
	}, nil
}

func (kp *KafkaProducer) ProduceEvent(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := kp.client.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
