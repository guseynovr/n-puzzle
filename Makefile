NAME	=	npuzzle

SRC		=	$(wildcard *.go) $(wildcard */*.go)

$(NAME):	$(SRC)
			go build $(NAME)

all: 		$(NAME)

compile:	
			go tool compile $(SRC)

clean:
			go clean

fclean:		clean
			rm -rf $(NAME)

re:			fclean all

.PHONY:		all clean fclean re compile
