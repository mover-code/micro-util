/***************************
@File        : base64.go
@Time        : 2022/07/04 17:34:42
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : base64 encrypt
****************************/

package encrypt

import (
	"encoding/base64"
)

//const Base64Salt = "Base64salt"

// base64 编码
// Base64Encode takes a byte slice and returns a string.
func Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

//  base64 解码
// Decode a base64 string into a string and a byte array.
func Base64Decode(str string) (string, []byte) {
	resBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(resBytes), resBytes
}
