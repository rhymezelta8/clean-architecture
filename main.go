package main

import (
	"log"
	"arquitectura/bootstrap"
)

func main(){
	if err:= bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
