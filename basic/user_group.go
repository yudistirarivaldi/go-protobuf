package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := basic.User{
		Id: 97,
		Username: "batman",
		IsActive: true,
		Password: []byte("password"),
		Gender: basic.Gender_GENDER_MALE,
	}

	joker := basic.User{
		Id: 98,
		Username: "Joker",
		IsActive: true,
		Password: []byte("thejoker"),
		Gender: basic.Gender_GENDER_MALE,
	}

	vilFamily := basic.UserGroup{
		GroupId: 999,
		GroupName: "Villain Family",
		Users: []*basic.User{&batman, &joker},
		Description: "The best villain ever",
	}

	jsonBytes, _ := protojson.Marshal(&vilFamily)

	log.Println(string(jsonBytes))
}