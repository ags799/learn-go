NAME := learn-go
VERSION := $(shell git describe --tags --always --dirty='-dev')

$(NAME): dependencies
	go build

.PHONY: clean
clean:
	(test -f $(NAME) && rm $(NAME)) || true

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

.PHONY: docker
docker: dependencies
	docker build -t $(NAME):$(VERSION) .

.PHONY: docker-run
docker-run: docker
	docker-compose up

.PHONY: docker-stop
docker-stop:
	docker-compose stop

.PHONY: integration-test
integration-test: dependencies
	go test -integration
