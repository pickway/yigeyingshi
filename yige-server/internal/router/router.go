package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/handler"
	"github.com/yigeyingshi/yige-server/internal/service"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
		c.Next()
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")

	movieSvc := service.NewMovieService(db)
	movieH := handler.NewMovieHandler(movieSvc)
	movies := api.Group("/movies")
	{
		movies.GET("", movieH.List)
		movies.GET("/featured", movieH.Featured)
		movies.GET("/:id", movieH.GetByID)
	}

	articleSvc := service.NewArticleService(db)
	articleH := handler.NewArticleHandler(articleSvc)
	articles := api.Group("/articles")
	{
		articles.GET("", articleH.List)
		articles.GET("/latest", articleH.Latest)
		articles.GET("/:id", articleH.GetByID)
	}

	searchSvc := service.NewSearchService(db)
	searchH := handler.NewSearchHandler(searchSvc)
	api.GET("/search", searchH.Search)

	learningSvc := service.NewLearningService(db)
	learningH := handler.NewLearningHandler(learningSvc)
	learning := api.Group("/learning")
	{
		learning.GET("/courses", learningH.ListCourses)
		learning.GET("/paths", learningH.ListPaths)
	}

	aiToolSvc := service.NewAiToolService(db)
	newsletterSvc := service.NewNewsletterService(db)
	aiH := handler.NewAiHandler(aiToolSvc, newsletterSvc)
	ai := api.Group("/ai")
	{
		ai.GET("/tools", aiH.ListTools)
		ai.GET("/tools/featured", aiH.FeaturedTools)
		ai.POST("/subscribe", aiH.Subscribe)
	}

	return r
}
