package controller

import (
	"datasource"
	"logger"
	"net/http"
	"route"
	"view"
)

type FriendsConteoller struct{}

func (self FriendsConteoller) RegisterRoutes() {
	route.HandleFunc("/friends", self.Detail)
}

func (FriendsConteoller) Detail(w http.ResponseWriter, r *http.Request) {
	friends, err := datasource.DefaultDataSourcer.GetFriends()
	if err == nil {
		view.Render(w, r, "friends.html", map[string]interface{}{"friends": friends})
	} else {
		logger.Instance().Error("get friends.yaml error " + err.Error())
		http.NotFound(w, r)
	}
}
