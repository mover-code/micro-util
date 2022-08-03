/***************************
@File        : translate_test.go
@Time        : 2022/08/03 17:45:05
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : translate_test
****************************/

package translate

import (
    "fmt"
    "testing"
    "time"
)

func TestT(t *testing.T) {
    fmt.Println(time.Now())
    fmt.Println(Translate("审核通知", "zh-CN", "en"))
    fmt.Println(time.Now())
}
