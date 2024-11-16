package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
	h := basic.User{
		Id: 99,
		Username: "superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Email: 	  []string{"superman@movie.com", "superman@dc.com"},
		Gender: basic.Gender_GENDER_MALE,
	}

	log.Println(&h)
}