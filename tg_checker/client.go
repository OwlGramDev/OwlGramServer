package tg_checker

import (
	"OwlGramServer/tg_checker/types"
	"log"
)

func Client() *Context {
	var listStatus []types.DCStatus
	log.Println("Started Telegram DC Checker!")
	return &Context{
		listStatus,
		0,
		true,
	}
}
