/***************************
@File        : init.go
@Time        : 2022/03/30 13:52:45
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : ipfs 服务
****************************/

package ipfs

import (
	"io"

	shell "github.com/ipfs/go-ipfs-api"
)

// var sh *shell.Shell
type IPFS struct {
	Sh *shell.Shell
}

func Init(addr string) *IPFS {
	return &IPFS{
		Sh: shell.NewShell(addr + ":5001"),
	}
}

func (sh *IPFS) Add(r io.Reader) (string, error) {
	// log.Println("ipfs 上传文件中")
	return sh.Sh.Add(r)
}
