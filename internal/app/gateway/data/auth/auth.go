package auth

import (
	"github.com/go-gateway/internal/app/gateway/biz"
	"github.com/go-gateway/internal/app/gateway/biz/domain"
	"github.com/go-gateway/internal/app/gateway/data/db"
	"github.com/go-gateway/internal/app/gateway/data/schema"
	"github.com/sirupsen/logrus"
)

type AuthURLRepo struct {
	// *sql.DB
}

func NewAuthURLRepo() biz.AuthRepo {
	return &AuthURLRepo{}
}

var cache []*domain.AuthURL

func (repo *AuthURLRepo) List() ([]*domain.AuthURL, error) {
	if cache != nil {
		logrus.Info("命中缓存")
		return cache, nil
	}
	rows, err := db.DB.Query("select service_id, url, force_auth, prefix from auth_url")
	if err != nil {
		return nil, err
	}
	var list []*schema.AuthURL
	for rows.Next() {
		var r schema.AuthURL
		err = rows.Scan(&r.ServiceId, &r.Url, &r.ForceAuth, &r.Prefix)
		if err != nil {
			return nil, err
		}
		list = append(list, &r)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	cache = convert(list)
	return cache, nil
}

func convert(list []*schema.AuthURL) []*domain.AuthURL {
	var result []*domain.AuthURL
	for _, au := range list {
		var authURL = &domain.AuthURL{}
		authURL.ServiceId = au.ServiceId
		authURL.Url = au.Url
		authURL.Prefix = au.Prefix
		authURL.ForceAuth = au.ForceAuth
		result = append(result, authURL)
	}
	return result
}
