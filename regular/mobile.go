/***************************
@File        : mobile.go
@Time        : 2022/06/17 11:24:48
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 校验手机号
****************************/

package regular

import "regexp"

func VerifyMobile(mobile string) bool {
    regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
    reg := regexp.MustCompile(regular)
    return reg.MatchString(mobile)
}
