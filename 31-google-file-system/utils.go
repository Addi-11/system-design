package main

import (
	"log"
)

func errorlog(msg string, err error){
	if err != nil{
		log.Fatal(msg, err)
	}
}