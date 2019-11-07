package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/yangxinwei/gin-demo-api/controllers"
	"github.com/yangxinwei/gin-demo-api/middlewares"
)

// LoggerWithFormatter instance a Logger middleware with the specified log format function.

func SetupRouter(controller *controllers.BaseController) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	// gin.SetMode(gin.ReleaseMode)
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middlewares.Ginrus(controller.Logger))
	//r.Use(gin.LoggerWithWriter(func(param gin.LogFormatterParams) string {
	//	// your favorites format
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))

	//v1 := r.Group("/api/v1")

	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.GET("/ping", controller.Ping)
	//passport := controllers.Passport{controller}
	//{
	//	v1.POST("/passport/register", passport.Register)
	//	v1.POST("/passport/login", passport.Login)
	//	v1.POST("/passport/logout", passport.Logout)
	//	v1.POST("/passport/changeMobile", passport.ChangeMobile)
	//	v1.POST("/passport/changePassword", passport.ChangePassword)
	//	v1.POST("/passport/getUserInfo", passport.GetUserInfo)
	//}
	//
	//favorites := controllers.Favorites{controller}
	//{
	//	v1.POST("/favorites/add", favorites.Add)
	//	v1.POST("/favorites/remove", favorites.Remove)
	//	v1.POST("/favorites/sort", favorites.Sort)
	//	v1.POST("/favorites/sortset", favorites.SortSet)
	//	v1.POST("/favorites/list", favorites.List)
	//	v1.POST("/favorites/quote", favorites.Quote)
	//}
	//
	//market := controllers.Market{controller}
	//{
	//	v1.POST("/market/search", market.Search)
	//	v1.POST("/market/quotes", market.Quotes)
	//	v1.POST("/market/", market.Snapshot)
	//	v1.POST("/market/status/", market.Status)
	//	v1.POST("/market/indices/", market.Indices)
	//	v1.POST("/market/hotstocks/", market.HotStocks)
	//	v1.POST("/market/chinese/", market.Chinese)
	//
	//	//v1.POST("/market/quote/", market.Quote)
	//	//v1.POST("/market/bar/", market.Bar)
	//	//v1.POST("/market/timeline/", market.TimeLine)
	//
	//	v1.POST("/market/stock", market.Stock)
	//
	//	v1.POST("/market/quote", market.Quote)
	//	v1.POST("/market/bar", market.Bar)
	//	v1.POST("/market/timeline", market.TimeLine)
	//	v1.POST("/market/news", market.News)
	//
	//	v1.POST("/market/index", market.Index)
	//	v1.POST("/market/index/stocks", market.IndexStocks)
	//
	//	v1.POST("/market/stock/fundamental", market.Fundamental)
	//	v1.POST("/market/stock/bulletins", market.Bulletins)
	//	v1.POST("/market/stock/profile", market.Profile)
	//	v1.POST("/market/stock/finance", market.Finance)
	//	}

	//
	r.POST("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	//r.POST("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := db[user]
	//	if ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//	}
	//})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if c.Bind(&json) == nil {
	//		db[user] = json.Value
	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})

	return r
}
