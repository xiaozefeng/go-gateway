package auth

import (
	"database/sql"

	"github.com/go-gateway/internal/app/gateway/biz"
	"github.com/go-gateway/internal/app/gateway/biz/domain"
	"github.com/go-gateway/internal/app/gateway/data/schema"
)

type AuthURLRepo struct {
	*sql.DB
}

func NewAuthURLRepo (db *sql.DB) biz.AuthRepo {
	return &AuthURLRepo{db}
}

func (repo *AuthURLRepo) List() ([]*domain.AuthURL, error) {
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

	return convert(list), nil
}

func convert(list []*schema.AuthURL) []*domain.AuthURL {
	var result []*domain.AuthURL
	for _, au := range list {
		var authURL *domain.AuthURL
		authURL.ServiceId = au.ServiceId
		authURL.Url = au.Url
		authURL.Prefix = au.Prefix
		authURL.ForceAuth = au.ForceAuth
		result = append(result, authURL)
	}
	return result
}
