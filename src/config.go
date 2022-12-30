package src

import (
	"os"
)

// Get APPID AppSecret from env
var (
	appID         = os.Getenv("APP_ID")
	appSecret     = os.Getenv("APP_SECRET")
	githubToken   = os.Getenv("GITHUB_TOKEN")
	githubOrgName = os.Getenv("GITHUB_ORG_NAME")
	port          = os.Getenv("PORT")
)
