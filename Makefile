NAME=npuzzle

all:
	go build -o npuzzle

fclean:
	rm -f $(NAME)

re: fclean all
