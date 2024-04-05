package routes

import (
	"BE-Inbuscap/config"
	invest "BE-Inbuscap/features/invest"
	proposal "BE-Inbuscap/features/proposal"
	"BE-Inbuscap/features/transaction"
	user "BE-Inbuscap/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uc user.Controller, pc proposal.Controller, cc invest.Controller, tc transaction.Controller) {
	config := echojwt.WithConfig(echojwt.Config{SigningKey: []byte(config.JWTSECRET)})

	userRoute(c, uc, config)
	proposalRoute(c, pc, config)
	investRoute(c, cc, config)
	transactionRoute(c, tc, config)
}

func userRoute(c *echo.Echo, uc user.Controller, config echo.MiddlewareFunc) {
	c.POST("/login", uc.Login())
	c.POST("/users", uc.Register())
	c.GET("/users", uc.Profile(), config)
	c.PUT("/users", uc.Update(), config)
	c.DELETE("/users", uc.Delete(), config)
}

func proposalRoute(c *echo.Echo, pc proposal.Controller, config echo.MiddlewareFunc) {
	c.POST("/proposals", pc.Create(), config)
	c.PUT("/proposals/:proposal_id", pc.Update(), config)
	c.GET("/proposals", pc.GetAll())
	c.GET("/proposals/:proposal_id", pc.GetDetail())
	c.DELETE("/proposals/:proposal_id", pc.Delete(), config)
	// c.POST("/proposals/:proposal_id", pc.Create(), config)
}

func investRoute(c *echo.Echo, cc invest.Controller, config echo.MiddlewareFunc) {
	// c.POST("/invests", cc.Create(), config)
	// c.DELETE("/invests/:investID", cc.Delete(), config)
}

func transactionRoute(c *echo.Echo, cc transaction.Controller, config echo.MiddlewareFunc) {
	c.POST("/transactions/topup", cc.AddTransaction(), config)
	c.POST("/transactions/callback", cc.CallBack())
	c.GET("/transactions/topup/:id", cc.CheckTransaction())
}
