/***************************
@File        : conmuster.go
@Time        : 2022/06/16 17:37:52
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : kafka conmuster
****************************/

package mq

import (
	"github.com/Shopify/sarama"
)

// Consumer is a sarama consumer 消费队列
func (c *Config) Consumer() ([]byte, error) {
    consumer, err := sarama.NewConsumer(c.Addr, config)
    if err != nil {
        return nil, err
    }
    defer consumer.Close()

    partition_consumer, err := consumer.ConsumePartition(c.Topic, 0, sarama.OffsetOldest)
    if err != nil {
        return nil, err
    }
    defer partition_consumer.Close()

    for {
        select {
        case msg := <-partition_consumer.Messages():
            return msg.Value, nil
        case err := <-partition_consumer.Errors():
            return nil, err
        }
    }
}

// Metadata client
func (c *Config) Metadata() (*sarama.Client, error) {
    client, err := sarama.NewClient(c.Addr, config)
    if err != nil {
        return nil, err
    }
    return &client, nil
}
