package thirdparty

import (
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member"
)

var ProviderSet = wire.NewSet(eureka.NewClient, member.NewUserCase)
