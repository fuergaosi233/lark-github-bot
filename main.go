package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"regexp"

	"github.com/chyroc/lark"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Get APPID AppSecret from env
var (
	appID         = os.Getenv("APP_ID")
	appSecret     = os.Getenv("APP_SECRET")
	githubToken   = os.Getenv("GITHUB_TOKEN")
	githubOrgName = os.Getenv("GITHUB_ORG_NAME")
	port          = os.Getenv("PORT")
)

func sendOrganizationInviteRequest(email string) string {
	print("send invite request to: ", email)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	_, _, err := client.Organizations.CreateOrgInvitation(ctx, githubOrgName, &github.CreateOrgInvitationOptions{
		Email:  &email,
		Role:   github.String("direct_member"),
		TeamID: []int64{},
	})
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	username := strings.Split(email, "@")[0]
	return "Invite " + username + " success"
}

// Use regex git first email
func getFirstEmail(msg string) string {
	email := regexp.MustCompile(`(?m)[\w\.-]+@[\w\.-]+\.\w+`).FindString(msg)
	return email
}
func ReciverMessage(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
	content, err := lark.UnwrapMessageContent(event.Message.MessageType, event.Message.Content)
	if err != nil {
		return "", err
	}
	msg := ""
	switch event.Message.MessageType {
	case lark.MsgTypeText:
		msg = content.Text.Text
	default:
		return "", nil
	}
	email := getFirstEmail(
		msg,
	)
	if email == "" {
		return "", nil
	}
	result := sendOrganizationInviteRequest(email)
	_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, result)
	return "", err
}
func main() {
	cli := lark.New(lark.WithAppCredential(appID, appSecret))
	cli.EventCallback.HandlerEventV2IMMessageReceiveV1(ReciverMessage)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		cli.EventCallback.ListenCallback(r.Context(), r.Body, w)
	})

	fmt.Println("start server ... 9726")
	if port == "" {
		port = "9726"
	}
	log.Fatal(http.ListenAndServe(":9726", nil))
}
