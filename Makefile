clean:
	rm -f cmd/gateway/gateway

build: clean
	cd cmd/gateway && go build -v .


run:
	nohup ./cmd/gateway/gateway -c ./conf/tk.yaml >> /dev/null 2>&1 &