package service

import (
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open test database: %v", err)
	}
	if err := db.AutoMigrate(
		&model.Movie{},
		&model.Article{},
		&model.LearningCourse{},
		&model.LearningPath{},
		&model.AiTool{},
		&model.Newsletter{},
	); err != nil {
		t.Fatalf("migrate test database: %v", err)
	}
	return db
}

func TestMovieServiceListFiltersKeywordAndSortsByNewest(t *testing.T) {
	db := newTestDB(t)
	movies := []model.Movie{
		{Title: "星际穿越", Director: "克里斯托弗·诺兰", Genres: "科幻,剧情", Year: 2014, Rating: 9.4},
		{Title: "奥本海默", Director: "克里斯托弗·诺兰", Genres: "传记,历史", Year: 2023, Rating: 8.9},
		{Title: "千与千寻", Director: "宫崎骏", Genres: "动画,剧情", Year: 2001, Rating: 9.3},
	}
	if err := db.Create(&movies).Error; err != nil {
		t.Fatalf("seed movies: %v", err)
	}

	service := NewMovieService(db)
	list, total, err := service.List("", "诺兰", "year_desc", 1, 10)
	if err != nil {
		t.Fatalf("list movies: %v", err)
	}
	if total != 2 {
		t.Fatalf("expected 2 matches, got %d", total)
	}
	if len(list) != 2 || list[0].Title != "奥本海默" {
		t.Fatalf("expected newest Nolan movie first, got %#v", list)
	}
}

func TestMovieServiceListPopularSortIsDeterministic(t *testing.T) {
	db := newTestDB(t)
	movies := []model.Movie{
		{Title: "较早高分", Year: 2010, Rating: 9.2},
		{Title: "较新同分", Year: 2020, Rating: 9.2},
		{Title: "低分", Year: 2024, Rating: 8.1},
	}
	if err := db.Create(&movies).Error; err != nil {
		t.Fatalf("seed movies: %v", err)
	}

	service := NewMovieService(db)
	list, _, err := service.List("", "", "popular", 1, 10)
	if err != nil {
		t.Fatalf("list movies: %v", err)
	}
	if len(list) != 3 || list[0].Title != "较新同分" {
		t.Fatalf("expected rating then year ordering, got %#v", list)
	}
}

func TestArticleServiceGetByIDReturnsArticle(t *testing.T) {
	db := newTestDB(t)
	article := model.Article{Title: "一篇影像笔记", Category: "movie-review", PublishedAt: time.Now()}
	if err := db.Create(&article).Error; err != nil {
		t.Fatalf("seed article: %v", err)
	}

	service := NewArticleService(db)
	actual, err := service.GetByID(article.ID)
	if err != nil {
		t.Fatalf("get article: %v", err)
	}
	if actual.Title != article.Title {
		t.Fatalf("expected %q, got %q", article.Title, actual.Title)
	}
}

func TestSearchServiceReturnsGroupedLimitedResults(t *testing.T) {
	db := newTestDB(t)
	if err := db.Create(&model.Movie{Title: "AI 电影实验", Genres: "科幻", Rating: 8.8}).Error; err != nil {
		t.Fatalf("seed movie: %v", err)
	}
	if err := db.Create(&model.Article{Title: "AI 与影像", Category: "ai-article", Summary: "创作观察", PublishedAt: time.Now()}).Error; err != nil {
		t.Fatalf("seed article: %v", err)
	}
	if err := db.Create(&model.LearningCourse{Title: "AI 视频生成", Category: "AI创作", IsActive: true}).Error; err != nil {
		t.Fatalf("seed course: %v", err)
	}
	if err := db.Create(&model.AiTool{Name: "AI Studio", Description: "视频生成", Featured: true}).Error; err != nil {
		t.Fatalf("seed AI tool: %v", err)
	}

	service := NewSearchService(db)
	results, err := service.Search("AI", 3)
	if err != nil {
		t.Fatalf("search site: %v", err)
	}
	if len(results.Movies) != 1 || len(results.Articles) != 1 || len(results.Courses) != 1 || len(results.Tools) != 1 {
		t.Fatalf("expected one result per group, got %#v", results)
	}
}
