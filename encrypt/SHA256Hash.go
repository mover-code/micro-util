/***************************
@File        : SHA256Hash.go
@Time        : 2022/07/04 17:40:29
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : sha256
****************************/
package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

//SHA256生成哈希值
// It takes a byte array as input, and returns a string as output
func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	_, _ = hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}
