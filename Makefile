.PHONY: build gen run start stop restart clean help

## build: clean and generate code and build
build: clean gen
	cd cmd/gateway && go build -tags=nomsgpack  -v .

## gen: generate code
gen:
	cd cmd/gateway && wire

## run: build and run config file: `./conf/tk.yaml`
run: build 
	./cmd/gateway/gateway -c ./conf/tk.yaml

## start: build and start server , config file: `/etc/tk.yaml`
start: build
	nohup ./cmd/gateway/gateway -c /etc/tk.yaml >> /dev/null 2>&1 &

## stop: stop the server
stop:
	ps -ef |grep gateway |grep -v grep  | awk '{print $$2}' |xargs kill -9

## restart: stop server then  start server
restart: stop start


## clean: remove gateway binary and gateway.log
clean:
	rm -f cmd/gateway/gateway
	rm -f gateway.log

## help: show this helo info
help: Makefile
	@echo  "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
