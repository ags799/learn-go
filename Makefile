NAME := learn-go

$(NAME): dependencies
	godep go build

.PHONY: dependencies
dependencies:
	go get -u github.com/tools/godep
	go get -u github.com/golang/lint/golint

.PHONY: run
run: $(NAME)
	./$(NAME)

.PHONY: test
test: $(NAME)
	godep go test

.PHONY: style
style: $(NAME)
	golint -set_exit_status
