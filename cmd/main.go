package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/proto"
	pro "github.com/golang/protobuf/proto"
)

func encode(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(jsonData)
}

func main() {
	book1 := &response.BookResponse{
		ID:          1,
		Title:       "Book 1",
		Quantity:    4,
		Description: "Lala",
		Cover:       "Sampul",
	}
	book2 := &response.BookResponse{
		ID:          2,
		Title:       "Book 2",
		Quantity:    8,
		Description: "Lala",
		Cover:       "Sampul",
	}

	books := []response.BookResponse{
		*book1,
		*book2,
	}

	jsonRes := encode(books)

	fmt.Println(jsonRes)

	bookProto1 := &proto.BookResponse{
		Id:          1,
		Title:       "Book 1",
		Description: "Lala",
		Quantity:    4,
		Cover:       "Sampul",
		Author:      nil,
	}

	bookProto2 := &proto.BookResponse{
		Id:          2,
		Title:       "Book 2",
		Description: "Lala",
		Quantity:    8,
		Cover:       "Sampul",
		Author:      nil,
	}

	booksProto := &proto.BookResponses{
		Books: []*proto.BookResponse{
			bookProto1, bookProto2,
		},
	}

	data, err := pro.Marshal(booksProto)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	err = ioutil.WriteFile("response.json", []byte(jsonRes), 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("response-proto", data, 0644)
	if err != nil {
		panic(err)
	}
}
