package src

import (
	"context"

	"github.com/chyroc/lark"
)

var LarkServer = lark.New(lark.WithAppCredential(appID, appSecret))

func init() {
	LarkServer.EventCallback.HandlerEventV2IMMessageReceiveV1(ReciverMessage)
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
