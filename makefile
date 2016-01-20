version=0.2.0

.PHONY: all

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - go build, vet & lint"
	@echo "  coverage      - generate a test coverage report"
	@echo "  deps          - pull and setup dependencies"
	@echo "  install       - run go install for all sub packages"
	@echo "  test          - run tests"
	@echo "  tools         - go gets a bunch of tools for development"
	@echo "  update_deps   - update deps lock file"

build:
	@go build ./...
	@go vet ./...
	@golint ./...

coverage:
	@go test -cover -v ./...

deps:
	@glock sync -n github.com/crowdriff/golog < Glockfile

install:
	@go install ./...

test:
	@ginkgo -r

tools:
	go get github.com/robfig/glock
	go get github.com/golang/lint/golint
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

update_deps:
	@glock save -n github.com/crowdriff/golog > Glockfile
