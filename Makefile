NAME=npuzzle
SRC = astar.go \
	  board.go \
	  board_utils.go \
	  generator.go \
	  heuristics.go \
	  main.go \
	  parsing.go \
	  queue.go \
	  state.go \
	  Makefile



all: $(NAME)

$(NAME): $(SRC)
	go build -o npuzzle

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
