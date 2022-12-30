package handle

import (
	"net/http"

	"github.com/bytemate/lark-github-bot/src"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	src.LarkServer.EventCallback.ListenCallback(r.Context(), r.Body, w)
}
