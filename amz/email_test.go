/***************************
@File        : email_test.go
@Time        : 2022/06/21 17:52:46
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 测试邮箱发送
****************************/

package amz

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
    conf := Init("email-smtp.ap-northeast-1.amazonaws.com", "AKIARNISI25Q7N27POXP", "BFYk2bwZ89Chdr4cIB3RXQ/NRuv4WfFh2LTd97VDKgwZ", "xms.chnb@gmail.com", 2587)
    fmt.Println(conf.Send("shuo582@163.com", "hello hello"))
}
