/***************************
@File        : types.go
@Time        : 2022/07/06 17:35:50
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : Request send
****************************/

package request

import (
    "net/http"
    "time"
)

type Cli struct {
    Client  *http.Client
    Timeout int64
    Heder   map[string]string
}

type Req struct {
    Method string
    Url    string
    Data   []byte
}

// NewCli returns a new Cli with a default http.Client.
func NewCli() *Cli {
    return &Cli{
        Client: &http.Client{},
    }
}

// Setting the timeout for the http client.
func (c *Cli) SetTimeout(timeout int64) {
    c.Timeout = timeout
    c.Client = &http.Client{Timeout: time.Duration(timeout)}
}

// NewReq returns a pointer to a Req with the given method, url, and data.
func NewReq(method, url string, data []byte) *Req {
    return &Req{
        Method: method,
        Url:    url,
        Data:   data,
    }
}
