package src

import (
	"context"
	"fmt"
	"strings"

	"regexp"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Use regex git first email
func getFirstEmail(msg string) string {
	email := regexp.MustCompile(`(?m)[\w\.-]+@[\w\.-]+\.\w+`).FindString(msg)
	return email
}
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
