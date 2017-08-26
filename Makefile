NAME := learn-go
VERSION := $(shell git describe --tags --always --dirty='-dev')
GO_FILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

.PHONY: all
all: test lint

$(NAME):
	go build

.PHONY: clean
clean:
	(test -f $(NAME) && rm $(NAME)) || true

.PHONY: devtools
devtools:
	go get -u github.com/FiloSottile/gvt
	go get -u github.com/golang/lint/golint
	go get -u golang.org/x/tools/cmd/goreturns

.PHONY: run
run: $(NAME)
	./$(NAME)

.PHONY: test
test:
	go test

.PHONY: lint
lint:
	go vet
	test -z '$(shell goreturns -l $(GO_FILES))'
	golint -set_exit_status

.PHONY: format
format:
	$(shell goreturns -w $(GO_FILES))

.PHONY: docker
docker:
	docker build -t $(NAME):$(VERSION) -t $(NAME):latest .

.PHONY: docker-run
docker-run: docker
	docker-compose up

.PHONY: docker-stop
docker-stop:
	docker-compose stop

.PHONY: integration-test
integration-test:
	go test -integration
