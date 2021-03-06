package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/zveinn/golang-lessons-for-beginners/src/files"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("You need to specify a search word ...")
		os.Exit(1)
	}

	files.LoadConfig()
	rand.Seed(time.Now().UTC().UnixNano())
	files.InitSearchBuffers(os.Args[1])
	files.InitPrintBuffers(os.Args[1])

	files.WalkDirectories(".")

	// Wait group
	files.GlobalWaitGroup.Wait()

	// meow
	// meow
	// meow
	// meow
	// meow
}

func something() (error, chan int) {
	return nil, make(chan int)
}
