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
	"io"
	"io/ioutil"

	"github.com/mover-code/micro-util/request"
)

var (
    apiUri = "https://api.nft.storage/upload"
    cli    = request.NewCli()
)

func AddStorage(f io.Reader, auth string) (string, error) {
    cli.Heder = map[string]string{
        "Authorization": "Bearer " + auth,
    }
    result := map[string]interface{}{}
    data, _ := ioutil.ReadAll(f)
    req := request.NewReq("POST", apiUri, data)
    res, err := cli.Do(req)
    json.Unmarshal(res, &result)
    if err == nil {
        if result["ok"].(bool) == true {
            return result["value"].(map[string]interface{})["cid"].(string), nil
        }
    }
    return result["error"].(map[string]interface{})["message"].(string), err
}