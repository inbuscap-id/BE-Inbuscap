package routes

import (
	"BE-Inbuscap/config"
	invest "BE-Inbuscap/features/invest"
	post "BE-Inbuscap/features/post"
	user "BE-Inbuscap/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uc user.Controller, pc post.Controller, cc invest.Controller) {
	config := echojwt.WithConfig(echojwt.Config{SigningKey: []byte(config.JWTSECRET)})

	userRoute(c, uc, config)
	postRoute(c, pc, config)
	investRoute(c, cc, config)
}

func userRoute(c *echo.Echo, uc user.Controller, config echo.MiddlewareFunc) {
	c.POST("/login", uc.Login())
	c.POST("/users", uc.Register())
	c.GET("/users", uc.Profile(), config)
	c.PUT("/users", uc.Update(), config)
	c.DELETE("/users", uc.Delete(), config)
}

func postRoute(c *echo.Echo, pc post.Controller, config echo.MiddlewareFunc) {
	// c.POST("/posts", pc.Create(), config)
	// c.PUT("/posts/:postID", pc.Edit(), config)
	// c.GET("/posts", pc.Posts())
	// c.GET("/posts/:postID", pc.PostById())
	// c.DELETE("/posts/:postID", pc.Delete(), config)
}

func investRoute(c *echo.Echo, cc invest.Controller, config echo.MiddlewareFunc) {
	// c.POST("/invests", cc.Create(), config)
	// c.DELETE("/invests/:investID", cc.Delete(), config)
}
