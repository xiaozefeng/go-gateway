package mapping

import "github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"

func MapToString(instances []eureka.Instance, apply func(eureka.Instance) string) []string {
	var res []string
	for _, v := range instances {
		res = append(res, apply(v))
	}
	return res
}
