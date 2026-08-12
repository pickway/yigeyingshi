package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-admin-server/internal/auth"
	"github.com/yigeyingshi/yige-admin-server/internal/config"
	"github.com/yigeyingshi/yige-admin-server/internal/handler"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config, mode string) *gin.Engine {
	gin.SetMode(mode)
	router := gin.New()
	_ = router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{AllowOrigins: []string{cfg.AllowedOrigin}, AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}}))
	router.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Next()
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "yige-admin-server"})
	})
	tokens := auth.NewTokenService(cfg.TokenSecret, cfg.TokenTTL)
	authHandler := handler.NewAuthHandler(cfg, tokens)
	resources := handler.NewResourceHandler(db)
	dashboard := handler.NewDashboardHandler(db)
	api := router.Group("/api/admin")
	api.POST("/auth/login", authHandler.Login)
	secured := api.Group("")
	secured.Use(auth.Middleware(tokens))
	secured.GET("/dashboard", dashboard.Get)
	register := func(path string, list, create, get, update, remove gin.HandlerFunc) {
		group := secured.Group(path)
		group.GET("", list)
		group.POST("", create)
		group.GET("/:id", get)
		group.PUT("/:id", update)
		group.DELETE("/:id", remove)
	}
	register("/movies", resources.ListMovies, resources.CreateMovie, resources.GetMovie, resources.UpdateMovie, resources.DeleteMovie)
	register("/articles", resources.ListArticles, resources.CreateArticle, resources.GetArticle, resources.UpdateArticle, resources.DeleteArticle)
	register("/courses", resources.ListCourses, resources.CreateCourse, resources.GetCourse, resources.UpdateCourse, resources.DeleteCourse)
	register("/paths", resources.ListPaths, resources.CreatePath, resources.GetPath, resources.UpdatePath, resources.DeletePath)
	register("/tools", resources.ListTools, resources.CreateTool, resources.GetTool, resources.UpdateTool, resources.DeleteTool)
	secured.GET("/subscribers", resources.ListSubscribers)
	secured.PATCH("/subscribers/:id/status", resources.SetSubscriberStatus)
	secured.DELETE("/subscribers/:id", resources.DeleteSubscriber)
	return router
}
