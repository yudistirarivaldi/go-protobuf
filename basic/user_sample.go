package basic

import (
	"my-protobuf/protogen/basic"
	"log"
)

func BasicUser() {
	h := basic.User{
		Id: 99,
		Username: "superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
	}

	log.Println(&h)
}