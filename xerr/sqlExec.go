/***************************
@File        : cache.go
@Time        : 2022/02/14 16:17:51
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 缓存错误
****************************/

package xerr

import "errors"

var (
	CacheNotFound = errors.New("cache data not find") // 没有缓存
	SearchError   = errors.New("search error")        // 查询出错
)
