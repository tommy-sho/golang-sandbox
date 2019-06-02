package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/tommy-sho/golang-sandbox/test/error-testing/repository"

	"github.com/tommy-sho/golang-sandbox/test/error-testing/domain"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	repo := repository.NewRepository()

L:
	for s.Scan() {
		n := s.Text()
		switch n {
		case "exit":
			break L
		case "new":
			u, err := Ask(s)
			if err != nil {
				fmt.Println("invalid input: ", err)
				continue
			}

			err = repo.CreateUser(u)
			if err != nil {
				fmt.Println(err)
			}
		case "show":
			fmt.Println("userID : ", "> ")
			users, err := repo.GetUsers()
			if err != nil {
				fmt.Println(err)
			}
			for _, v := range users {
				fmt.Println(v)
			}
		default:
		}
		fmt.Print("> ")
	}

}

func Ask(s *bufio.Scanner) (*domain.User, error) {
	field := make([]string, 3)
	msg := []string{":First name", ":Last name", ":Age"}
	for i, v := range msg {
		fmt.Println(v, "> ")
		s.Scan()
		field[i] = s.Text()
	}

	age, err := strconv.Atoi(field[2])
	if err != nil {
		return nil, err
	}
	return &domain.User{
		FirstName: field[0],
		LastName:  field[1],
		Age:       age,
	}, nil
}
