package main

import (
	"encoding/json"
	logger "github.com/sirupsen/logrus"
)


//https://www.sohamkamani.com/golang/2018-07-19-golang-omitempty/
//https://golang.org/pkg/encoding/json/
func main()  {
	// TODO: Example print struct of pointer
	t := Customer{
		Id:   createInt(4),
		Name: createString("name"),
		Age:  createString("age"),
	}
	s, _ := json.MarshalIndent(t, "", "\t")
	logger.Info(string(s))
	// End example
}

type Customer struct {
	Id 		*int 		`json:"id"`		// pointer to convert: null(input, db) <-> nil (go)
	Name 	*string 	`json:"name"`
	Age 	*string 	`json:"age"`
}

func createInt(x int) *int {
	return &x
}

func createString(x string) *string{
	return &x
}