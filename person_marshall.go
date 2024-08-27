package main

import (
	"go_protobuf/models"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	person := &models.Person{
		Name:  "John Wick",
		Id:    1234,
		Email: "wick@codeheim.io",
		Phones: []*models.PhoneNumber{
			{Number: "123-456-7891", Type: models.PhoneType_MOBILE},
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}

	os.WriteFile("tmp/person.bin", data, 0644)
}
