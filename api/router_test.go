package api

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)


func Test_str(t *testing.T){ 
	path :="/api-server/hello"
	s := strings.Split(path, `/`)
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("s: %v\n", s)
	x := s[1]
	fmt.Printf("x: %v\n", x)
	i := strings.Index(path,x)
	fmt.Printf("i: %v\n", i)
	fmt.Println(path[i+len(x):])

}

func Test_url(t *testing.T){ 
	u, _:= url.Parse("http://172.16.32.8:8888/")
	fmt.Printf("u.Host: %v\n", u.Host)
	fmt.Printf("u.Port(): %v\n", u.Port())

}