NO_COLOR=\033[0m
OK_COLOR=\033[0;32m

all: format vet lint

format:
	@echo "$(OK_COLOR)==> Formatting the code $(NO_COLOR)"
	@gofmt -s -w *.go
	@goimports -w *.go || true

test-privileged:
	@echo "$(OK_COLOR)==> Running tests for privileged user $(NO_COLOR)"
	@sudo `which go` test -v -coverprofile=/tmp/priv.out
	@sudo `which go` test -tags static_build -v -coverprofile=/tmp/priv.out

test-privileged-race:
	@echo "$(OK_COLOR)==> Running tests with -race flag for privileged user $(NO_COLOR)"
	@sudo `which go` test -race -v
	@sudo `which go` test -tags static_build -race -v

test:
	@echo "$(OK_COLOR)==> Running tests for unprivileged user $(NO_COLOR)"
	@`which go` test -v -coverprofile=/tmp/unpriv.out
	@`which go` test -tags static_build -v -coverprofile=/tmp/unpriv.out

test-race:
	@echo "$(OK_COLOR)==> Running tests with -race flag for unprivileged user $(NO_COLOR)"
	@`which go` test -race -v
	@`which go` test -tags static_build -race -v

cover:
	@echo "$(OK_COLOR)==> Running cover for privileged user $(NO_COLOR)"
	@`which go` tool cover -func=/tmp/priv.out || true
	@echo "$(OK_COLOR)==> Running cover for unprivileged user $(NO_COLOR)"
	@`which go` tool cover -func=/tmp/unpriv.out || true

doc:
	@`which go` doc github.com/lxc/go-lxc | less

vet:
	@echo "$(OK_COLOR)==> Running go vet $(NO_COLOR)"
	@`which go` vet .

lint:
	@echo "$(OK_COLOR)==> Running golint $(NO_COLOR)"
	@`which golint` . || true

escape-analysis:
	@`which go` build -gcflags -m

ctags:
	@ctags -R --languages=c,go

update-gomod:
	go get -t -v -d -u ./...
	go mod tidy

scope:
	@echo "$(OK_COLOR)==> Exported container calls in container.go $(NO_COLOR)"
	@/bin/grep -E "\bc+\.([A-Z])\w+" container.go || true

.PHONY: all format test doc vet lint ctags
