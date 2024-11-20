package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {

	addr := basic.Address{
		Street: "Jl Haji Noin",
		City: "Jakarta",
		Country: "Indonesia",
		Coordinate: &basic.Address_Coordinate{
			Latitude: 40.70129319,
			Longitude: -92.02302030,
		},
	}

	h := basic.User{
		Id: 99,
		Username: "superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Email: 	  []string{"superman@movie.com", "superman@dc.com"},
		Gender: basic.Gender_GENDER_MALE,
		Address: &addr,
	}

	jsonBytes, _ := protojson.Marshal(&h)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id: 98,
		Username: "batman",
		IsActive: true,
		Password: []byte("batmanpassword"),
		Email: []string{"batman@gmail.co.id", "batman@gmail.com"},
		Gender: basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := 
	`{
		"id":98,
		"username":"batman",
		"is_active":true,
		"password":"YmF0bWFucGFzc2dvcmQ=",
		"email":[
			"batman@gmail.co.id",
			"batman@gmail.com"
		],
		"gender":"GENDER_FEMALE"
	}`

	p := basic.User{}

	err := protojson.Unmarshal([]byte(json), &p)
	if err != nil {
		log.Fatalln("Got error : ", err)
	}

	log.Println(&p)

}
