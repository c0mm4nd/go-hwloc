BIN_DIR=_output/cmd/bin

all: init build

build: init
	go build -o ${BIN_DIR}/hwloc-ls ./cmd/hwloc-ls

verify:
	hack/verify-gofmt.sh

init:
	mkdir -p ${BIN_DIR}
clean:
	rm -fr ${BIN_DIR}

.PHONY: clean

