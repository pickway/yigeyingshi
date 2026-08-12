package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-admin-server/internal/model"
	"gorm.io/gorm"
)

type ResourceHandler struct{ db *gorm.DB }

func NewResourceHandler(db *gorm.DB) *ResourceHandler { return &ResourceHandler{db: db} }

func pagination(c *gin.Context) (int, int, bool) {
	page, e1 := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, e2 := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if e1 != nil || e2 != nil || page < 1 || size < 1 || size > 100 {
		Error(c, http.StatusBadRequest, "INVALID_PAGINATION", "分页参数无效")
		return 0, 0, false
	}
	return page, size, true
}

func (h *ResourceHandler) list(c *gin.Context, target interface{}, table string, fields []string) {
	page, size, ok := pagination(c)
	if !ok {
		return
	}
	query := h.db.Table(table)
	keyword := strings.TrimSpace(c.Query("keyword"))
	if keyword != "" && len(fields) > 0 {
		clauses := make([]string, len(fields))
		args := make([]interface{}, len(fields))
		for i, f := range fields {
			clauses[i] = f + " LIKE ?"
			args[i] = "%" + keyword + "%"
		}
		query = query.Where(strings.Join(clauses, " OR "), args...)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		Error(c, 500, "DATABASE_ERROR", "加载失败")
		return
	}
	if err := query.Order("id DESC").Offset((page - 1) * size).Limit(size).Find(target).Error; err != nil {
		Error(c, 500, "DATABASE_ERROR", "加载失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": target, "total": total, "page": page, "pageSize": size})
}

func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, 400, "INVALID_ID", "编号无效")
		return 0, false
	}
	return uint(id), true
}
func (h *ResourceHandler) create(c *gin.Context, item interface{}) {
	if err := c.ShouldBindJSON(item); err != nil {
		Error(c, 422, "VALIDATION_ERROR", "请检查必填字段")
		return
	}
	if err := h.db.Create(item).Error; err != nil {
		Error(c, 500, "DATABASE_ERROR", "保存失败")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": item})
}
func (h *ResourceHandler) update(c *gin.Context, item interface{}) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := c.ShouldBindJSON(item); err != nil {
		Error(c, 422, "VALIDATION_ERROR", "请检查必填字段")
		return
	}
	result := h.db.Model(item).Where("id = ?", id).Select("*").Omit("id", "created_at").Updates(item)
	if result.Error != nil {
		Error(c, 500, "DATABASE_ERROR", "保存失败")
		return
	}
	if result.RowsAffected == 0 {
		Error(c, 404, "NOT_FOUND", "内容不存在")
		return
	}
	h.db.First(item, id)
	c.JSON(200, gin.H{"data": item})
}
func (h *ResourceHandler) get(c *gin.Context, item interface{}) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.db.First(item, id).Error; err != nil {
		Error(c, 404, "NOT_FOUND", "内容不存在")
		return
	}
	c.JSON(200, gin.H{"data": item})
}
func (h *ResourceHandler) remove(c *gin.Context, item interface{}) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	result := h.db.Delete(item, id)
	if result.Error != nil {
		Error(c, 500, "DATABASE_ERROR", "删除失败")
		return
	}
	if result.RowsAffected == 0 {
		Error(c, 404, "NOT_FOUND", "内容不存在")
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *ResourceHandler) ListMovies(c *gin.Context) {
	h.list(c, &[]model.Movie{}, "movies", []string{"title", "director", "genres"})
}
func (h *ResourceHandler) CreateMovie(c *gin.Context) { h.create(c, &model.Movie{}) }
func (h *ResourceHandler) GetMovie(c *gin.Context)    { h.get(c, &model.Movie{}) }
func (h *ResourceHandler) UpdateMovie(c *gin.Context) { h.update(c, &model.Movie{}) }
func (h *ResourceHandler) DeleteMovie(c *gin.Context) { h.remove(c, &model.Movie{}) }
func (h *ResourceHandler) ListArticles(c *gin.Context) {
	h.list(c, &[]model.Article{}, "articles", []string{"title", "author", "category", "tag"})
}
func (h *ResourceHandler) CreateArticle(c *gin.Context) { h.create(c, &model.Article{}) }
func (h *ResourceHandler) GetArticle(c *gin.Context)    { h.get(c, &model.Article{}) }
func (h *ResourceHandler) UpdateArticle(c *gin.Context) { h.update(c, &model.Article{}) }
func (h *ResourceHandler) DeleteArticle(c *gin.Context) { h.remove(c, &model.Article{}) }
func (h *ResourceHandler) ListCourses(c *gin.Context) {
	h.list(c, &[]model.LearningCourse{}, "learning_courses", []string{"title", "category", "level"})
}
func (h *ResourceHandler) CreateCourse(c *gin.Context) { h.create(c, &model.LearningCourse{}) }
func (h *ResourceHandler) GetCourse(c *gin.Context)    { h.get(c, &model.LearningCourse{}) }
func (h *ResourceHandler) UpdateCourse(c *gin.Context) { h.update(c, &model.LearningCourse{}) }
func (h *ResourceHandler) DeleteCourse(c *gin.Context) { h.remove(c, &model.LearningCourse{}) }
func (h *ResourceHandler) ListPaths(c *gin.Context) {
	h.list(c, &[]model.LearningPath{}, "learning_paths", []string{"name", "description"})
}
func (h *ResourceHandler) CreatePath(c *gin.Context) { h.create(c, &model.LearningPath{}) }
func (h *ResourceHandler) GetPath(c *gin.Context)    { h.get(c, &model.LearningPath{}) }
func (h *ResourceHandler) UpdatePath(c *gin.Context) { h.update(c, &model.LearningPath{}) }
func (h *ResourceHandler) DeletePath(c *gin.Context) { h.remove(c, &model.LearningPath{}) }
func (h *ResourceHandler) ListTools(c *gin.Context) {
	h.list(c, &[]model.AiTool{}, "ai_tools", []string{"name", "description", "category", "tags"})
}
func (h *ResourceHandler) CreateTool(c *gin.Context) { h.create(c, &model.AiTool{}) }
func (h *ResourceHandler) GetTool(c *gin.Context)    { h.get(c, &model.AiTool{}) }
func (h *ResourceHandler) UpdateTool(c *gin.Context) { h.update(c, &model.AiTool{}) }
func (h *ResourceHandler) DeleteTool(c *gin.Context) { h.remove(c, &model.AiTool{}) }

func (h *ResourceHandler) ListSubscribers(c *gin.Context) {
	h.list(c, &[]model.Newsletter{}, "newsletters", []string{"email"})
}
func (h *ResourceHandler) SetSubscriberStatus(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var input struct {
		Active bool `json:"active"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		Error(c, 422, "VALIDATION_ERROR", "状态无效")
		return
	}
	result := h.db.Model(&model.Newsletter{}).Where("id = ?", id).Update("active", input.Active)
	if result.Error != nil {
		Error(c, 500, "DATABASE_ERROR", "状态更新失败")
		return
	}
	if result.RowsAffected == 0 {
		Error(c, 404, "NOT_FOUND", "订阅者不存在")
		return
	}
	var item model.Newsletter
	if err := h.db.First(&item, id).Error; err != nil {
		Error(c, 500, "DATABASE_ERROR", "状态已更新，但读取结果失败")
		return
	}
	c.JSON(200, gin.H{"data": item})
}
func (h *ResourceHandler) DeleteSubscriber(c *gin.Context) { h.remove(c, &model.Newsletter{}) }
