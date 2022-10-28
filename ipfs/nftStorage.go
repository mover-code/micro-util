/***************************
@File        : nftStorage.go
@Time        : 2022/10/28 13:52:02
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : upload nftStorage
****************************/

package ipfs

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/mover-code/micro-util/request"
)

var (
    apiUri = "https://api.nft.storage/upload"
    cli    = request.NewCli()
)

func Add(r io.Reader, auth string) (string, error) {
    cli.Heder = map[string]string{
        "Authorization": "Bearer " + auth,
    }
    result := map[string]interface{}{}
    params := map[string]interface{}{
        "file": r,
    }
    data, _ := json.Marshal(&params)
    req := request.NewReq("POST", apiUri, data)
    res, err := cli.Do(req)
    json.Unmarshal(res, &result)
    if err == nil {
        fmt.Println(result)
        return result["value"].(map[string]interface{})["cid"].(string), nil
    }
    return result["error"].(map[string]interface{})["message"].(string), err
}
