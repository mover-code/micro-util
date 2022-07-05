/***************************
@File        : product.go
@Time        : 2022/06/16 15:50:23
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : kafka producer
****************************/
package mq

import (
    "fmt"
    "time"

    "github.com/Shopify/sarama"
)

func NewProducerMes(mes, topic, key string, timestamp time.Time) *sarama.ProducerMessage {
    msg := &sarama.ProducerMessage{
        Topic:     topic,
        Key:       sarama.StringEncoder(key),
        Value:     sarama.StringEncoder(mes),
        Timestamp: timestamp,
    }
    msg.Value = sarama.ByteEncoder(mes)
    return msg
}

func (c *Config) ProducerSend(mes *sarama.ProducerMessage) error {
    producer, err := sarama.NewAsyncProducer(c.Addr, config)
    if err != nil {
        return err
    }
    defer producer.AsyncClose()
    for {
        producer.Input() <- mes
        select {
        case suc := <-producer.Successes():
            fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
            return nil
        case fail := <-producer.Errors():
            return fail.Err
        }
    }
}
