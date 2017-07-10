NAME := learn-go

$(NAME): dependencies
	go build

.PHONY: dependencies
dependencies:
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
