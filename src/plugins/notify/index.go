package main

import (
	"log"
)

func Init(){
	hooks.OnConnect(func() {
		log.Println("- ON CONNECT HOOK")
	})

	hooks.OnMessage(func() {
		log.Println("- ON MESSAGE HOOK")
	})
}
