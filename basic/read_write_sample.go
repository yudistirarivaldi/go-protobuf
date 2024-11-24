package basic

import (
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(msg proto.Message, fname string) {
	b, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Can not marshal message", err)
	}

	err = ioutil.WriteFile(fname, b, 0644)
	if err != nil {
		log.Fatalln("Can not write to file", err)
	}


}

func WriteProtoToJSON(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatalln("Can not marshal message", err)
	}

	err = ioutil.WriteFile(fname, b, 0644)
	if err != nil {
		log.Fatalln("Can not write to file", err)
	}


}

func ReadProtoFromFile(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can not read file", err)
	}

	err = proto.Unmarshal(in, dest)
	if err != nil {
		log.Fatalln("Can not unmarshal", err)
	}

}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can not read file", err)
	}

	err = protojson.Unmarshal(in, dest)
	if err != nil {
		log.Fatalln("Can not unmarshal", err)
	}

}

func dummyUser() basic.User {
	addr := basic.Address{
		Street: "Daily Planet",
		City: "Metropolis",
		Country: "US",
		Coordinate: &basic.Address_Coordinate{
			Latitude: 40.12312,
			Longitude: -74.828193,
		},
	}

	comm := randomCommunicationChannel()
	sr := map[string]uint32{
		"fly": 5,
		"speed": 5,
		"durability": 5,
	}

	return basic.User{
		Id: 99,
		Username: "superman",
		IsActive: true,
		Password: []byte("password"),
		Address: &addr,
		CommunicationChannel: &comm,
		SkillRating: sr,
	}

}

func WriteToFileSample() {
	u := dummyUser()
	WriteProtoToFile(&u, "superman_file.bin")
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtoToJSON(&u, "superman_file.json")
}


func ReadFromFileSample() {
	dest := basic.User{}

	ReadProtoFromFile("superman_file.bin", &dest)

	log.Println(&dest)
}

func ReadFromJSONSample() {
	dest := basic.User{}

	ReadProtoFromJSON("superman_file.json", &dest)

	log.Println(&dest)
}