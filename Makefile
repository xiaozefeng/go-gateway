clean:
	rm -f cmd/gateway/gateway
	rm -f gateway.log

build: clean generate
	cd cmd/gateway && go build -v .

generate:
	cd internal/pkg/app && wire

start: build kill
	nohup ./cmd/gateway/gateway -c ./conf/tk.yaml >> /dev/null 2>&1 &

kill:
	ps -ef |grep gateway |grep -v grep  | awk '{print $2}' |xargs kill -9

try: build
	./cmd/gateway/gateway  -c ./conf/config.yaml
