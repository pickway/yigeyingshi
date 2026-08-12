# CineVerse 管理后台前端

Vue 3、Vite 与 Element Plus 构建的内容管理界面，覆盖影片、文章、课程、学习路径、AI 工具和邮件订阅者。

## 启动

```powershell
npm install
npm run dev
```

默认打开 `http://127.0.0.1:4174`，开发代理把 `/api` 转发到 `http://127.0.0.1:8081`。如需自定义接口地址，可复制 `.env.example` 为 `.env`。

## 验证

```powershell
npm test
npm run build
```

登录令牌仅保存在当前标签页的 `sessionStorage` 中；接口返回 401 时会自动清理会话并回到登录页。
