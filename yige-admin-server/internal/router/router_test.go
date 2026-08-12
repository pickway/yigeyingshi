package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/yigeyingshi/yige-admin-server/internal/auth"
	"github.com/yigeyingshi/yige-admin-server/internal/config"
	"github.com/yigeyingshi/yige-admin-server/internal/model"
	"gorm.io/gorm"
)

func setupTestRouter(t *testing.T) http.Handler {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}
	if err := db.AutoMigrate(&model.Movie{}, &model.Article{}, &model.LearningCourse{}, &model.LearningPath{}, &model.AiTool{}, &model.Newsletter{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	cfg := &config.Config{Username: "admin", Password: "correct-password", TokenSecret: "test-secret-that-is-long-enough", TokenTTL: time.Hour, AllowedOrigin: "http://localhost:4174"}
	return Setup(db, cfg, "test")
}

func requestJSON(t *testing.T, router http.Handler, method, path string, body interface{}, token string) *httptest.ResponseRecorder {
	t.Helper()
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(method, path, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func loginToken(t *testing.T, router http.Handler) string {
	t.Helper()
	res := requestJSON(t, router, http.MethodPost, "/api/admin/auth/login", map[string]string{"username": "admin", "password": "correct-password"}, "")
	if res.Code != http.StatusOK {
		t.Fatalf("login failed: %d %s", res.Code, res.Body.String())
	}
	var payload struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	if err := json.Unmarshal(res.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode login: %v", err)
	}
	return payload.Data.Token
}

func TestProtectedRouteRejectsAnonymousRequest(t *testing.T) {
	router := setupTestRouter(t)
	res := requestJSON(t, router, http.MethodGet, "/api/admin/dashboard", nil, "")
	if res.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", res.Code)
	}
}

func TestHealthEndpoint(t *testing.T) {
	router := setupTestRouter(t)
	res := requestJSON(t, router, http.MethodGet, "/health", nil, "")
	if res.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.Code)
	}
}

func TestLoginRejectsWrongPasswordAndAcceptsCorrectPassword(t *testing.T) {
	router := setupTestRouter(t)
	bad := requestJSON(t, router, http.MethodPost, "/api/admin/auth/login", map[string]string{"username": "admin", "password": "wrong"}, "")
	if bad.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", bad.Code)
	}
	if token := loginToken(t, router); token == "" {
		t.Fatal("expected token")
	}
}

func TestMovieCRUDAndDashboard(t *testing.T) {
	router := setupTestRouter(t)
	token := loginToken(t, router)
	create := requestJSON(t, router, http.MethodPost, "/api/admin/movies", map[string]interface{}{"title": "后台测试影片", "year": 2026, "rating": 8.8, "genres": "剧情"}, token)
	if create.Code != http.StatusCreated {
		t.Fatalf("create failed: %d %s", create.Code, create.Body.String())
	}
	var created struct {
		Data model.Movie `json:"data"`
	}
	_ = json.Unmarshal(create.Body.Bytes(), &created)
	id := strconv.FormatUint(uint64(created.Data.ID), 10)
	update := requestJSON(t, router, http.MethodPut, "/api/admin/movies/"+id, map[string]interface{}{"title": "已更新影片", "year": 2026, "rating": 9.1, "genres": "剧情"}, token)
	if update.Code != http.StatusOK {
		t.Fatalf("update failed: %d %s", update.Code, update.Body.String())
	}
	dashboard := requestJSON(t, router, http.MethodGet, "/api/admin/dashboard", nil, token)
	if dashboard.Code != http.StatusOK {
		t.Fatalf("dashboard failed: %d", dashboard.Code)
	}
	deleteRes := requestJSON(t, router, http.MethodDelete, "/api/admin/movies/"+id, nil, token)
	if deleteRes.Code != http.StatusNoContent {
		t.Fatalf("delete failed: %d", deleteRes.Code)
	}
}

func TestTokenServiceRejectsTamperedToken(t *testing.T) {
	service := auth.NewTokenService("a sufficiently long test secret", time.Hour)
	token, err := service.Issue("admin")
	if err != nil {
		t.Fatalf("issue token: %v", err)
	}
	if _, err := service.Verify(token + "x"); err == nil {
		t.Fatal("expected tampered token rejection")
	}
}
