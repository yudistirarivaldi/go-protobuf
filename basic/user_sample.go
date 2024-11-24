package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
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

	comm := randomCommunicationChannel()

	h := basic.User{
		Id: 99,
		Username: "superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Email: 	  []string{"superman@movie.com", "superman@dc.com"},
		Gender: basic.Gender_GENDER_MALE,
		Address: &addr,
		CommunicationChannel: &comm,
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

func randomCommunicationChannel() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "instagram",
		SocialMediaUsername: "ydstrarvldiii",

	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct: "whatsup",
		InstantMessagingUsername: "crypto",
	}

	a := anypb.Any{}

	switch r := rand.Intn(10) % 3; r {
	case 0:
			anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (social media)
	sm := basic.SocialMedia{}

	err := a.UnmarshalTo(&sm)
	if err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))

}

func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()
	if err != nil {
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))

}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		err := a.UnmarshalTo(&pm)
		if err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not papermain, but : ", a.TypeUrl)
	}
}

func BasicOneOf() {
	// socialMedia := basic.SocialMedia{
	// 	SocialMediaPlatform: "twitter",
	// 	SocialMediaUsername: "peter",
	// }

	// ecomm := basic.User_SocialMedia{
	// 	SocialMedia: &socialMedia,
	// }

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct: "IG",
		InstantMessagingUsername: "Spiderman",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMessaging,
	}

	u := basic.User{
		Id: 96,
		Username: "peter",
		IsActive: true,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}