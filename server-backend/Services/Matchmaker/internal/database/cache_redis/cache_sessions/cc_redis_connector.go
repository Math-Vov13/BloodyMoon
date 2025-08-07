package cache_sessions

import "github.com/Math-Vov13/BloodyMoon/models/cache/sessions_models"

var fake_cache = map[string]*sessions_models.User{
	"1234": {
		ID:       "1234",
		Username: "test1",
		Status:   "offline",
	},
	"5678": {
		ID:       "5678",
		Username: "test2",
		Status:   "offline",
	},
	"567891": {
		ID:       "567891",
		Username: "test3",
		Status:   "offline",
	},
	"56781564": {
		ID:       "56781564",
		Username: "test4",
		Status:   "offline",
	},
}
