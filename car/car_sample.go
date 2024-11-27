package car

import (
	"log"
	"my-protobuf/protogen/car"

	"github.com/google/uuid"
)



func ValidateCar() {
	car := car.Car{
		CarId: uuid.New().String(),
		Brand: "BMW",
		Model: "432",
		Price: 1000,
		ManufactureYear: 2021,
	}

	err := car.ValidateAll()
	if err != nil {
		log.Fatalln("Validation Failed", err)
	}

	log.Fatalln(&car)
}