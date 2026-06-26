package main

import "log"

func logUser(user string) {
	log.Printf("event=user_lookup value=%s", user)
}
