package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// constructor function
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("please fill all the fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// method
func (u *User) OutputUserDetails() {

	fmt.Println((u).firstName, (u).lastName)
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)

}
