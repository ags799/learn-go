NAME := learn-go

$(NAME):
	go build

.PHONY: run
run: $(NAME)
	./$(NAME)
