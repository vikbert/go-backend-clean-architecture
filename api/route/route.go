package route

import (
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/api/middleware"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {

	// All Public APIs without bearer token
	publicRouter := gin.Group("api/public")
	RouterSignup(env, timeout, db, publicRouter)
	RouterLogin(env, timeout, db, publicRouter)
	RouterRefreshToken(env, timeout, db, publicRouter)

	// Middleware to verify AccessToken
	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Private APIs
	RouterProfile(env, timeout, db, protectedRouter)
	RouterTask(env, timeout, db, protectedRouter)
}
