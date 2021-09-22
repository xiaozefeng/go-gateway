package eureka

import (
	"encoding/json"
	"fmt"
	"testing"
)

var c = NewClient("http://172.16.0.11:9100/eureka")

func Test_GetApps(t *testing.T) {
	apps, err := c.GetApps()
	if err != nil {
		t.Error(err)
	}
	for _, a := range apps.Apps.App {
		fmt.Printf("a.Name: %v\n", a.Name)
	}
}

func TestClient_GetApp(t *testing.T) {
	app, err := c.GetApp("otmgroup-openapi-gateway")
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("pp: %+v\n", app)

	b, err := json.MarshalIndent(app, "", "   ")

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", b)

}
