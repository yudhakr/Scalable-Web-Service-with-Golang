package v1

import (
	engine "github.com/Calmantara/go-fga/config/gin"
	"github.com/Calmantara/go-fga/pkg/domain/user"
	"github.com/Calmantara/go-fga/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) get() {
	// all path for get method are here
}

func (u *UserRouterImpl) post() {
	// all path for post method are here
	u.routerGroup.POST("", u.userHandler.InsertUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.post()
}
