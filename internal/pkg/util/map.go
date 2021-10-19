package util

import "github.com/go-gateway/internal/pkg/client/eureka"


func LoadBalance(instances []string) string {
	if len(instances) > 0 {
		return instances[0]
	}
	return ""
}


func MapToString(instances []eureka.Instance, apply func(eureka.Instance) string) []string {
	var res []string
	for _, v := range instances {
		res = append(res, apply(v))
	}
	return res
}