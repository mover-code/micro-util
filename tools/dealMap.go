package tools

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

//获取map所有key的方式
func GetMapKeys(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetAloneGradeEquity(aloneGradeEquity string) map[string]string {
	var result map[string]string
	_ = json.Unmarshal([]byte(aloneGradeEquity), &result)
	return result
}

//判断key是否在map中
func KeyInMap(myMap map[string]string, key string) bool {
	if _, ok := myMap[key]; ok {
		return true
	}
	return false
}

type MapEntryHandler func(string, string)

// 按字母顺序遍历map
func TraverseMapInStringOrder(params map[string]string, handler MapEntryHandler) {
	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		handler(k, params[k])
	}
}
func HttpBuildQuery(params map[string]string) (paramStr string) {
	paramsArr := make([]string, 0, len(params))
	for k, v := range params {
		if k == "ret_url" || k == "notify_url" || k == "sign" {
			v = url.QueryEscape(v)
		}
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", k, v))
	}
	//fmt.Println(params_arr)
	paramStr = strings.Join(paramsArr, "&")
	return paramStr
}
