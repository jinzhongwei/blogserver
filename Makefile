GOPATH:=$(CURDIR)
export GOPATH
$(warning $(GOPATH))
all: build

fmt:
	gofmt -l -w -s src/

build:dep
	go build -o bin/blogserver blogserver

clean:
	@rm -f bin/blogserver
	@rm -rf ./pkg
	@rm -rf status
	@rm -f  log/*log*
	@rm -rf output

dep:fmt
output: build
	mkdir -p output/bin
	mkdir -p output/conf
	mkdir -p output/log
	mkdir -p output/status
	cp -r bin/* output/bin/
	cp -r conf/* output/conf/
cleanlog:
	@rm -f log/*log*

