clean:
	rm -f cmd/gateway/gateway
	rm -f gateway.log

build: clean gen
	cd cmd/gateway && go build -v .

gen:
	cd internal/pkg/wire && wire

run: build 
	./cmd/gateway/gateway -c ./conf/tk.yaml

start: build
	nohup ./cmd/gateway/gateway -c /etc/tk.yaml >> /dev/null 2>&1 &

stop:
	ps -ef |grep gateway |grep -v grep  | awk '{print $$2}' |xargs kill -9

restart: stop start


