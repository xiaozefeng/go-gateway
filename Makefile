clean:
	rm -f cmd/gateway/gateway

build: clean
	cd cmd/gateway && go build -v .