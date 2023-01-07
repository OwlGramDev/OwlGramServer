package main

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin"
	"OwlGramServer/gopy"
	"OwlGramServer/http/webserver"
	"OwlGramServer/telegram/bot"
	"OwlGramServer/telegram/checker"
	"OwlGramServer/telegram/emoji"
	"OwlGramServer/updates"
	"OwlGramServer/utilities"
	"os"
	"os/exec"
)

var botClient *bot.Context
var tgChecker *checker.Context
var crowdinClient *crowdin.Context
var updatesClient *updates.Context
var emojiClient *emoji.Context
var pythonClient *gopy.Context

func main() {
	consts.LoadEnv()
	if _, err := os.Stat(consts.UploadsFolder); os.IsNotExist(err) {
		_ = os.Mkdir(consts.UploadsFolder, 0775)
	}
	_ = exec.Command("chmod", "+x", consts.AAPT2ToolPath).Run()
	pythonClient = utilities.SmartPythonBuild()
	emojiClient = emoji.Client()
	crowdinClient = crowdin.Client()
	updatesClient = updates.Client()
	botClient = bot.Bot(updatesClient, pythonClient)
	if !consts.IsDebug {
		tgChecker = checker.Client()
		go tgChecker.Run()
	}
	go botClient.Run()
	go emojiClient.Run()
	go crowdinClient.Run()
	go updatesClient.Run()
	webserver.Client(handler).Run()
}
