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
	fitness, err := instance.Evaluate([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19})
	checkError(err, "instance.Evaluate")
	log.Println("fitness:", fitness)
}
