package httputils

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/util/gconv"
)

/**
* 类似php httpbuilder的函数
* @description:
*
* @param {[type]} params
* @param {[type]} parentKey
 */
func HttpBuildQuery(params map[string]interface{}, parentKey string) (param_str string) {
	//fmt.Println("parentKey ", parentKey)
	params_arr := make([]string, 0)
	for k, v := range params {
		if vals, ok := v.(map[string]interface{}); ok {
			if parentKey != "" {
				k = fmt.Sprintf("%s[%s]", parentKey, k)
			}
			params_arr = append(params_arr, HttpBuildQuery(vals, k))
		} else {
			if parentKey != "" {
				params_arr = append(params_arr, fmt.Sprintf("%s[%s]=%s", parentKey, k, gconv.String(v)))
			} else {
				params_arr = append(params_arr, fmt.Sprintf("%s=%s", k, gconv.String(v)))
			}
		}
	}
	param_str = strings.Join(params_arr, "&")
	return param_str
}
