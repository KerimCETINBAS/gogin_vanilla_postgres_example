package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/repositories"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

func AuthMiddleWare(r repositories.IUserRepository) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			id   string
			idx  uint64
			err  error
			user types.User
		)
		id, err = ctx.Cookie("session")

		idx, err = strconv.ParseUint(id, 10, 32)

		if err != nil {
			ctx.Next()
			return
		}
		user, err = r.FindUserById(uint(idx))

		if err != nil {
			ctx.Next()
			return
		}
		ctx.Set("session", user)
		ctx.Next()
	}
}
