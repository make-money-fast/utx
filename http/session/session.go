package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionStart 开启session，使用cookie存储.
func SessionStart(key string, id string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(key))
	return sessions.Sessions(id, store)
}

func SessionGet(ctx *gin.Context) sessions.Session {
	return sessions.Default(ctx)
}
