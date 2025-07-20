package user

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type User struct {
	Id           uuid.UUID `json:"id"`
	UserName     string    `json:"user_name"`
	UserPassword string
}

func (u User) New() (User, error) {
	var userName, userPassword, userChoice string
	var user User

	fmt.Println("Please enter your username")
	fmt.Scanln(&userName)
	fmt.Printf("Thanks. you choosed : '%v' as username. Please enter now your password.", userName)
	fmt.Scanln(&userPassword)
	fmt.Println("Thanks. You wish to create that account now?\nPlease enter: [y/n]")
	fmt.Scanln(&userChoice)
	switch userChoice {
	case "y":
		fmt.Println("User created.")
		id, err := uuid.NewV4()
		if err != nil {
			return User{}, nil
		}
		user = User{
			Id:           id,
			UserName:     userName,
			UserPassword: userPassword,
		}
		return user, nil
	case "n":
		{
			fmt.Println("You choosed 'no'. If you wish create a new user now.")
		}
	}
	return user, nil
}
