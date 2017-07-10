NAME := learn-go

$(NAME):
	go build

.PHONY: run
run: $(NAME)
	./$(NAME)

.PHONY: test
test: $(NAME)
	go test
