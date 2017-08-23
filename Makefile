NAME := learn-go
VERSION := $(shell git describe --tags --always --dirty='-dev')

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

.PHONY: run
run: $(NAME)
	./$(NAME)

.PHONY: test
test: $(NAME)
	go test

.PHONY: lint
lint: $(NAME)
	golint -set_exit_status

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
