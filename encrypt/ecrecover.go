/***************************
@File        : ecrecover.go
@Time        : 2022/03/29 17:59:39
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : web3 sign
****************************/
package encrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// Ecrecover returns the uncompressed public key that created the given signature.
func Ecrecover(hash, sig []byte) ([]byte, error) {
	return secp256k1.RecoverPubkey(hash, sig)
}

// SigToPub returns the public key that created the given signature.
func SigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
	s, err := Ecrecover(hash, sig)
	if err != nil {
		return nil, err
	}

	x, y := elliptic.Unmarshal(secp256k1.S256(), s)
	return &ecdsa.PublicKey{Curve: secp256k1.S256(), X: x, Y: y}, nil
}

// It takes a message, a signature, and a public key, and returns true if the signature is valid for
// the message under the public key
func VerifySig(from string, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	// fmt.Println(sig[64])
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// fmt.Println(sigHex, "-------", hexutil.Encode(signHash(msg)), "输出信息")
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	// fmt.Println("签名地址", fromAddr, recoveredAddr)
	// vI, _ := strconv.Atoi(sigHex[130:])
	// fmt.Println("r", sigHex[:66], "s", "0x"+sigHex[66:130], "v", uint8(vI+27))
	// fmt.Println(fmt.Sprintf(`[%v,"%v","%v"]`, uint8(vI+27), sigHex[:66], "0x"+sigHex[66:130]))
	return fromAddr == recoveredAddr
}

// It takes a byte array, converts it to a string, prepends a string to it, and then hashes the result
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

// It takes a string and returns a byte array
func signString(data string) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", 32)
	return []byte(msg)
}

// It takes a string and returns a byte array that contains the string's length
func signLen(data string) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data))
	return []byte(msg)
}

// It takes the address of the owner, the tokenId of the token to be minted, and the private key of the
// owner, and returns the signature of the minting transaction
func SignMint(address string, tokenId int64, key string) map[string]interface{} {
	dataSign := map[string]interface{}{}
	privateKey, err := crypto.HexToECDSA(key)
	token_id := big.NewInt(tokenId)

	hash := solsha3.SoliditySHA3(
		// types
		[]string{"address", "uint256"},
		// values
		[]interface{}{
			address,
			token_id.String(),
		},
	)
	hashs := crypto.Keccak256Hash(signString(""), hash)
	signature, err := crypto.Sign(hashs.Bytes(), privateKey)
	signatures := hexutil.Encode(signature)
	// fmt.Println(signatures)
	dataSign["r"] = signatures[:66]
	dataSign["s"] = "0x" + signatures[66:130]
	vI, _ := strconv.Atoi(signatures[130:])
	dataSign["v"] = uint8(vI + 27)
	if err == nil {
		return dataSign
	}
	return nil
}

// It takes a message, hashes it, signs the hash with a private key, and returns the signature
func SignTest() {
	dataSign := map[string]interface{}{}

	msg := "\x19Ethereum Signed Message:\nhello world"
	privateKey, _ := crypto.HexToECDSA("57f558afb87fa20414ae84aeec05569f944ce7357cc41fb4d076b97b88a362bd")
	hash := crypto.Keccak256([]byte(msg))
	signature, _ := crypto.Sign(hash, privateKey)
	signatures := hexutil.Encode(signature)
	// log.Println(signatures)
	dataSign["r"] = signatures[:66]
	dataSign["s"] = "0x" + signatures[66:130]
	vI, _ := strconv.Atoi(signatures[130:])
	dataSign["v"] = uint8(vI + 27)
	// log.Println(fmt.Sprintf("%+v", dataSign))
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)
	// log.Println("r", hexutil.Encode(R[:]), "s", hexutil.Encode(S[:]))
	return R, S, V
}

// discard  & 弃用
func SignOrder(data []byte) string {
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"[]byte"}, //  "string"

		// values
		[]interface{}{
			data,
		},
	)
	fmt.Println(hexutil.Encode(hash), hexutil.Encode(crypto.Keccak256(data)))
	return hexutil.Encode(hash)
}

func SignBuyOrder(hash string, key string) map[string]interface{} {
	dataSign := map[string]interface{}{}
	privateKey, err := crypto.HexToECDSA(key)
	hashs := signHash([]byte(hash))
	signature, err := crypto.Sign(hashs, privateKey)
	signatures := hexutil.Encode(signature)
	// fmt.Println(signatures)
	dataSign["signatures"] = signatures
	dataSign["r"] = signatures[:66]
	dataSign["s"] = "0x" + signatures[66:130]
	vI, _ := strconv.Atoi(signatures[130:])
	dataSign["v"] = uint8(vI + 27)
	if err == nil {
		return dataSign
	}
	return nil
}

// It takes a string of the signature, splits it into the R, S, and V values, and then returns a map of
// the values
func stringToRSV(signatures string) {
	dataSign := map[string]interface{}{}
	dataSign["r"] = signatures[:66]
	dataSign["s"] = "0x" + signatures[66:130]
	vI, _ := strconv.Atoi(signatures[130:])
	dataSign["v"] = uint8(vI + 27)
	fmt.Println(fmt.Sprintf("%+v", dataSign))
}
