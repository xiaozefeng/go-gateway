package auth

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz/domain"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/schema"
)

type URLRepo struct {
	*sql.DB
}

func NewURLRepo(db *sql.DB) biz.AuthRepo {
	return &URLRepo{db}
}

var cache []*domain.AuthURL

func (repo *URLRepo) List() ([]*domain.AuthURL, error) {
	if cache != nil {
		logrus.Info("命中缓存")
		return cache, nil
	}
	rows, err := repo.Query("select service_id, url, force_auth, prefix from auth_url")
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
