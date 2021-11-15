package data

import (
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/auth"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/db"
)

var ProviderSet = wire.NewSet(db.New, auth.NewURLRepo)
