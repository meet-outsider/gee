package pkg

import "log"

func Todo() string {
	log.Default().Println("be careful, this is a TODO")
	return "TODO"
}
