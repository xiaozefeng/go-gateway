package auth

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-gateway/internal/data/schema"
	_ "github.com/go-sql-driver/mysql"
)

func Test_db(t *testing.T){ 
	user := "u_super_hero"
	passwd := "u*P1e5r@hE2r"
	host := "172.16.32.3"
	port := 3306
	database := "api_gate"
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, passwd, host, port, database)

	db, err := getDataSource(url)
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

	// columns, err := rows.Columns()
	// if err != nil {
	// 	t.Error(err)
	// }

	// values := make([]sql.RawBytes, len(columns))
	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }
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


func getDataSource(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to the mysql server")
	return db, nil
}