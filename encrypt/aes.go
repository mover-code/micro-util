/***************************
@File        : aes.go
@Time        : 2022/07/04 17:34:28
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : aes encrypt
****************************/
package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
//key不能泄露
var PwdKey = []byte("ABCDABCDABCDABCD")

//pkcs7Padding 填充
// It pads the data with the number of bytes needed to make the data a multiple of the block size.
func pkcs7Padding(data []byte, blockSize int) []byte {
    //判断缺少几位长度。最少1，最多 blockSize
    padding := blockSize - len(data)%blockSize
    //补足位数。把切片[]byte{byte(padding)}复制padding个
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(data, padText...)
}

// It removes all the null bytes from the end of the input byte array
func NullUnPadding(in []byte) []byte {
    return bytes.TrimRight(in, string([]byte{0}))
}

// If the length of the input is a multiple of 8, return the input, otherwise return the input with
// enough zero bytes appended to make the length a multiple of 8
func ZeroPadding(in []byte) []byte {
    length := len(in)
    if length%8 == 0 {
        return in
    } else {
        blockCount := length / 8
        out := make([]byte, (blockCount+1)*8)
        var i int
        for i = 0; i < length; i++ {
            out[i] = in[i]
        }
        return out
    }
}

//pkcs7UnPadding 填充的反向操作
// If the last byte of the input is less than 16, then the last byte is the number of padding bytes,
// and the last byte is removed from the input
func pkcs7UnPadding(data []byte) ([]byte, error) {
    length := len(data)
    if length == 0 {
        return nil, errors.New("加密字符串错误！")
    }
    //获取填充的个数
    unPadding := int(data[length-1])
    return data[:(length - unPadding)], nil
}

//AesEncrypt 加密
// It takes a byte array and a key, and returns a byte array
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
    //创建加密实例
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    //判断加密快的大小
    blockSize := block.BlockSize()
    //填充
    // encryptBytes := pkcs7Padding(data, blockSize)
    encryptBytes := ZeroPadding(data)
    //初始化加密数据接收切片
    crypted := make([]byte, len(encryptBytes))
    //使用cbc加密模式
    blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    //执行加密
    blockMode.CryptBlocks(crypted, encryptBytes)
    return crypted, nil
}

//AesDecrypt 解密
// It decrypts the data using AES-CBC.
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
    //创建实例
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    //获取块的大小
    blockSize := block.BlockSize()
    //使用cbc
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    //初始化解密数据接收切片
    crypted := make([]byte, len(data))
    //执行解密
    blockMode.CryptBlocks(crypted, data)
    //去除填充
    crypted, err = pkcs7UnPadding(crypted)
    if err != nil {
        return nil, err
    }
    return crypted, nil
}

//EncryptByAes Aes加密 后 base64 再加
// Encrypts the data using AES and returns the encrypted data in base64 format.
func EncryptByAes(data []byte) (string, error) {
    res, err := AesEncrypt(data, PwdKey)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(res), nil
}

//DecryptByAes Aes 解密
// It decrypts the data using AES.
func DecryptByAes(data string) ([]byte, error) {
    dataByte, err := base64.StdEncoding.DecodeString(data)
    if err != nil {
        return nil, err
    }
    return AesDecrypt(dataByte, PwdKey)
}
