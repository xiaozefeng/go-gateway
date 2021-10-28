package auth

import (
	"fmt"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/db"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/schema"
)

func Test_db(t *testing.T) {
	user := "u_super_hero"
	passwd := "u*P1e5r@hE2r"
	host := "172.16.32.3"
	port := 3306
	database := "api_gate"
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, passwd, host, port, database)

	db, err := db.Init(url)
	if err != nil {
		t.Error(err)
	}
	log.Println("Connected to the MySQL Server")
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("select service_id, url, force_auth, prefix from auth_url")
	if err != nil {
		t.Error(err)
	}

	var result []schema.AuthURL
	for rows.Next() {
		var r schema.AuthURL
		err = rows.Scan(&r.ServiceId, &r.Url, &r.ForceAuth, &r.Prefix)
		if err != nil {
			t.Error(err)
		}
		result = append(result, r)
	}
	if err = rows.Err(); err != nil {
		t.Error(err)
	}
	for _, au := range result {
		fmt.Printf("au: %+v\n", au)
	}

}
