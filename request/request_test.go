/***************************
@File        : request_test.go
@Time        : 2022/07/06 18:06:56
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : test request send
****************************/

package request

import "testing"

func TestSend(t *testing.T) {
    cli := NewCli()
    // cli.SetTimeout(5)
    req := NewReq("GET", "http://www.baidu.com", nil)
    body, err := cli.Do(req)
    if err != nil {
        t.Error(err)
    }
    t.Log(string(body))
}
