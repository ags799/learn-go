NAME := learn-go
VERSION := $(shell git describe --tags --always --dirty='-dev')

$(NAME): dependencies
	go build

.PHONY: dependencies
dependencies:
	go get -u github.com/FiloSottile/gvt
	go get -u github.com/golang/lint/golint

.PHONY: run
run: $(NAME)
	./$(NAME)

.PHONY: test
test: $(NAME)
	go test

.PHONY: style
style: $(NAME)
	golint -set_exit_status

docker: dependencies
	docker build -t $(NAME):$(VERSION) .

docker-run: docker
	docker run --rm --publish 8080:8080 --name $(NAME)-test $(NAME):$(VERSION)

docker-stop:
	docker stop $(NAME)-test

integration-test: dependencies
	go test -integration
