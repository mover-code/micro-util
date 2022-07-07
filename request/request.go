/***************************
@File        : request.go
@Time        : 2022/07/06 17:47:49
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : req send
****************************/

package request

import (
    "bytes"
    "io/ioutil"
    "net/http"
)

// Do
// A simple http client.
func (c *Cli) Do(r *Req) (body []byte, err error) {
    req, err := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(r.Data))
    if err != nil {
        return nil, err
    }
    for k, v := range c.Heder {
        req.Header.Add(k, v)
    }
    resp, err := c.Client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    // if resp.StatusCode != 200 {
    //     return nil, err
    // }
    body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return body, nil
}
