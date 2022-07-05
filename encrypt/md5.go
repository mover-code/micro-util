/***************************
@File        : md5.go
@Time        : 2022/07/04 17:39:55
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : md5 encrypt
****************************/
package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// const md5Salt = "md5salt"

// 生成32位md5 salt加密字符串
// It takes a string and a salt, and returns the MD5 hash of the string concatenated with the salt
func Md5SaltString(str, md5Salt string) string {
	h := md5.New()
	_, _ = h.Write([]byte(str + md5Salt))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成32位md5字符串
// It takes a string, creates a new MD5 hash, writes the string to the hash, and then returns the
// hexadecimal representation of the hash
func Md5String(str string) string {
	h := md5.New()
	_, _ = h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// It takes a byte slice and returns a string
func Md5Bytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
