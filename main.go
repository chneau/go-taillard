package main

import (
	"log"

	"github.com/chneau/go-taillard/pfsp"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func checkError(err error, message string) {
	if err != nil {
		log.Panicln(message, err)
	}
}

func main() {
	instance, err := pfsp.NewMakespan(20, 5, 0)
	checkError(err, "pfsp.New")
	log.Printf("%+v", instance)
}
