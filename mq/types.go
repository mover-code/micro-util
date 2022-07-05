/***************************
@File        : types.go
@Time        : 2022/06/16 15:57:27
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        :
****************************/

package mq

import "github.com/Shopify/sarama"

var config = sarama.NewConfig()

func init() {
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true
    config.Producer.Return.Errors = true
    config.Version = sarama.V0_11_0_2
}

type Config struct {
    Addr   []string // kafka地址
    Topic  string   // kafka topic
    Key    string   // kafka key
    Expire int64    // 过期时间
}
