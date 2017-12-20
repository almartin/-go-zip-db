GENERATOR=generate_zip.go
TOOL_DIR=tools
FILE=zip.go


all: build test install

.PHONY: all

build:
	go run tools/${GENERATOR}
	go fmt ${FILE}
	go build -i *.go

.PHONY: build

test: 
	go test -v -cover

.PHONY: test

install:
	go install 

.PHONY: install

# Cleans our project: deletes generated file
clean:
	if [ -f ${FILE} ] ; then rm ${FILE} ; fi

.PHONY: clean
