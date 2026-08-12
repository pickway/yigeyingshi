package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

func testRouter(t *testing.T) http.Handler {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open database: %v", err)
	}
	if err := db.AutoMigrate(&model.Movie{}, &model.Article{}, &model.LearningCourse{}, &model.LearningPath{}, &model.AiTool{}, &model.Newsletter{}); err != nil {
		t.Fatalf("migrate database: %v", err)
	}
	if err := db.Create(&model.Movie{Title: "星际穿越", Director: "诺兰", Genres: "科幻", Year: 2014, Rating: 9.4}).Error; err != nil {
		t.Fatalf("seed movie: %v", err)
	}
	return Setup(db, "test")
}

func TestMovieListRejectsInvalidPagination(t *testing.T) {
	router := testRouter(t)
	request := httptest.NewRequest(http.MethodGet, "/api/v1/movies?page=0&pageSize=100", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", response.Code)
	}
	var payload map[string]map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload["error"]["code"] != "INVALID_PAGINATION" {
		t.Fatalf("unexpected error response: %#v", payload)
	}
}

func TestSearchReturnsGroupedResultsAndSecurityHeaders(t *testing.T) {
	router := testRouter(t)
	request := httptest.NewRequest(http.MethodGet, "/api/v1/search?q=星际", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}
	if response.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatalf("missing security header")
	}
	var payload struct {
		Data struct {
			Movies []model.Movie `json:"movies"`
		} `json:"data"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Data.Movies) != 1 {
		t.Fatalf("expected one movie, got %#v", payload.Data.Movies)
	}
}
