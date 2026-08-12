# Spec: CineVerse 个人内容网站完善

## Objective

将现有影视、学习与 AI 内容展示 MVP 完善为无需登录即可使用的公开个人网站。核心用户是希望发现优质影片、学习影视创作并了解 AI 创作工具的中文读者。

成功体验包括：清晰的个人品牌表达、可搜索和筛选的内容、可阅读的电影与文章详情、可靠的空/错/加载状态、完整的移动端适配，以及可持续扩展的数据与接口边界。

## Tech Stack

- Frontend: Vue 3.5, Vue Router 4.6, Vite 8, Tailwind CSS 4, Lucide Vue
- Backend: Go 1.22, Gin 1.9, GORM 1.25, SQLite
- Testing: Go standard library + `httptest`; frontend pure-logic tests via Node test runner

## Commands

执行构建、测试和启动前，PowerShell 进程必须清空大小写代理环境变量。

- Frontend dev: `npm run dev`
- Frontend test: `npm test`
- Frontend build: `npm run build`
- Backend test: `go test ./...`
- Backend dev: `go run ./cmd/server`

## Project Structure

- `yige-ui/src/views`: route-level pages
- `yige-ui/src/components`: reusable presentation and state components
- `yige-ui/src/api`: API boundary and error normalization
- `yige-ui/src/utils`: pure view/domain helpers
- `yige-server/internal/handler`: HTTP validation and response mapping
- `yige-server/internal/service`: query and domain behavior
- `yige-server/internal/model`: persistence and JSON contracts
- `yige-server/internal/router`: routes and cross-cutting middleware
- `docs`: product and implementation documentation

## Code Style

```go
func (s *MovieService) GetByID(id uint) (*model.Movie, error) {
	var movie model.Movie
	err := s.db.First(&movie, id).Error
	return &movie, err
}
```

- Go methods use explicit `this.` equivalent receiver (`s.` / `h.`) for instance calls.
- Vue uses `<script setup>` and named domain variables; shared behavior lives in small pure helpers.
- API fields use camelCase and additive changes; errors use `{ error: { code, message } }`.
- Avoid unsafe HTML rendering; all external input is validated at the handler boundary.

## Testing Strategy

- Unit tests cover parsing, search/sort helpers, pagination validation, and API error behavior.
- Integration tests use in-memory SQLite and `httptest` for list/detail/search routes.
- Production build proves SFC templates and styles compile.
- Browser checks cover desktop and mobile layout, console errors, navigation, search, details, and newsletter validation.

## Boundaries

- Always: validate query/body input, parameterize database queries, preserve backward-compatible response fields, show user-visible error states, test new behavior.
- Ask first: authentication, private/admin data, new third-party runtime integrations, deployment credentials.
- Never: commit secrets, render untrusted HTML, expose internal errors, commit/push/create PR without explicit instruction.

## Success Criteria

- Home page communicates author identity and routes users into the three content pillars.
- Global search finds movies, articles, courses, and AI tools from one interface.
- Movie and article cards link to meaningful detail pages with related content.
- Movie filtering supports keyword, genre, valid sorting, and bounded pagination.
- All pages have responsive, accessible loading/empty/error feedback.
- Backend returns consistent public errors and rejects invalid pagination/subscription input.
- Frontend build, frontend tests, backend tests, and browser smoke tests pass.
- README explains local setup and production assumptions.

## Open Questions

- Real social/contact URLs and the author's real biography are not available; polished placeholders remain clearly editable.
- A content-management login is intentionally out of scope for this public-site iteration.

