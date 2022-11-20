package consts

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

var CrowdinAuthToken string
var BotToken string
var GithubBotToken string
var GithubToken string
var SecretDCKey string
var HuaweiClientId string
var HuaweiClientSecret string
var PublisherToken string
var WebAppLink string
var ApiID string
var ApiHash string
var IsDebug bool
var SshIP string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	CrowdinAuthToken = os.Getenv("CROWDIN_AUTH_TOKEN")
	BotToken = os.Getenv("BOT_TOKEN")
	GithubBotToken = os.Getenv("GITHUB_BOT_TOKEN")
	GithubToken = os.Getenv("GITHUB_TOKEN")
	SecretDCKey = os.Getenv("SECRET_DC_KEY")
	HuaweiClientId = os.Getenv("HUAWEI_CLIENT_ID")
	HuaweiClientSecret = os.Getenv("HUAWEI_CLIENT_SECRET")
	PublisherToken = os.Getenv("PUBLISHER_TOKEN")
	ApiID = os.Getenv("API_ID")
	ApiHash = os.Getenv("API_HASH")
	IsDebug = strings.Contains(os.Args[0], "tmp")
	WebAppLink = "https://app.owlgram.org/webapp"
	if IsDebug {
		BotToken = os.Getenv("BOT_TOKEN_DEBUG")
		WebAppLink = "https://app-test.owlgram.org/webapp"
	}
	SshIP = strings.Split(os.Getenv("SSH_CLIENT"), " ")[0]
}
