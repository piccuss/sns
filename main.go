package main

import (
	"log"
)

func init() {
	_logInit()
}

func _logInit() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetPrefix("[SNS]")
	log.Println("log style init end.")
}

func main() {
	Start()
}