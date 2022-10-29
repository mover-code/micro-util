/***************************
@File        : add_test.go
@Time        : 2022/03/30 13:59:25
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 测试ipfs
****************************/

package ipfs

import (
	"log"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	log.Println(Init("addr").Sh.Add(strings.NewReader("hello world")))
}

func TestAddNftStorage(t *testing.T) {
	log.Println(AddStorage(strings.NewReader("today is rain"), "token"))
}
