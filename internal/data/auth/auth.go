package auth

import (
	"github.com/go-gateway/internal/data/db"
	"github.com/go-gateway/internal/data/schema"
)

type AuthURLRepo struct {

}

func (authRepo *AuthURLRepo) ListAuthURL() ([]*schema.AuthURL, error) {
	rows, err := db.DB.Query("select service_id, url, force_auth, prefix from auth_url")
	if err != nil {
		return nil, err
	}
	var v []*schema.AuthURL
	for rows.Next() {
		var r schema.AuthURL
		err = rows.Scan(&r.ServiceId, &r.Url, &r.ForceAuth, &r.Prefix)
		if err != nil {
			return nil, err
		}
		v = append(v, &r)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return v, nil
}
