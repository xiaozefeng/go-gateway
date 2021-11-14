package biz

import (
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member"
)

var ProviderSet = wire.NewSet(NewTokenUserCase, NewAuthUserCase,eureka.NewClient, member.NewUserCase)

