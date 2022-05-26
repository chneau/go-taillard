package main

import (
	"log"

	"github.com/chneau/go-taillard/pfsp"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

// checkError
func ce(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}

func main() {
	instance, err := pfsp.NewMakespan(20, 5, 0)
	ce(err, "pfsp.New")
	log.Printf("%+v", instance)
}
