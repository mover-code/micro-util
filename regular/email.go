/***************************
@File        : email.go
@Time        : 2022/06/17 11:22:36
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 正则校验邮箱
****************************/

package regular

import "regexp"

// VerigyEmail 校验邮箱
func VerifyEmail(email string) bool {
    pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

    reg := regexp.MustCompile(pattern)
    return reg.MatchString(email)
}
