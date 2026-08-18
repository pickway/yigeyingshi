# CineVerse MySQL 数据库设计方案

## 一、概述

将原有 SQLite 数据源切换为 MySQL 8.x，供前台 `yige-server` 和后台 `yige-admin-server` 共享访问。全部表使用 InnoDB 引擎、`utf8mb4` 字符集与 `utf8mb4_general_ci` 排序规则，保证中文与 emoji 的稳定存储与检索。

- 数据库名：`yigeyingshi`
- 连接方式：TCP `118.145.113.88:3306`
- DSN 格式：`user:pass@tcp(host:port)/yigeyingshi?charset=utf8mb4&parseTime=True&loc=Local`

## 二、ER 关系

```
movies               articles            learning_courses      ai_tools
  ├─ id(PK)           ├─ id(PK)            ├─ id(PK)             ├─ id(PK)
  └─ is_featured      ├─ category(IDX)     ├─ category(IDX)      ├─ category(IDX)
                      ├─ published_at         ├─ is_active        ├─ featured
                                                              newsletters
                                                                 ├─ id(PK)
                                                      learning_paths ─ email(UNIQ)
                                                           ├─ id(PK)
```

六张表之间没有外键硬约束，保持与原 SQLite 相同的松耦合设计（学习路径 `course_ids` 仍以逗号分隔字符串形式存储），便于 GORM 直接映射与内容管理后台编辑。

## 三、表结构明细

### 3.1 movies — 影片资源

| 字段 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT UNSIGNED | PRIMARY KEY AUTO_INCREMENT | 主键 |
| title | VARCHAR(128) | NOT NULL | 影片标题 |
| original_title | VARCHAR(128) | NULL | 原始标题（外文原名） |
| year | INT | DEFAULT 0 | 上映年份 |
| genres | VARCHAR(128) | NULL | 类型，逗号分隔（科幻,剧情） |
| rating | DECIMAL(3,1) | DEFAULT 0.0 | 评分 0.0–10.0 |
| description | TEXT | NULL | 简介 |
| poster_color | VARCHAR(32) | NULL | 海报背景色 #RRGGBB |
| poster_icon | VARCHAR(32) | NULL | 海报图标 key |
| duration | INT | DEFAULT 0 | 片长（分钟） |
| director | VARCHAR(64) | NULL | 导演 |
| is_featured | TINYINT(1) | DEFAULT 0 | 是否首页精选 |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |

**索引**：`idx_movies_featured` (is_featured)，`idx_movies_year` (year)

### 3.2 articles — 文章（影评/学习笔记/AI 文章）

| 字段 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT UNSIGNED | PRIMARY KEY AUTO_INCREMENT | 主键 |
| title | VARCHAR(256) | NOT NULL | 标题 |
| category | VARCHAR(32) | NOT NULL | 分类枚举：movie-review / learning-note / ai-article |
| tag | VARCHAR(32) | NULL | 二级标签 |
| summary | TEXT | NULL | 摘要 |
| content | MEDIUMTEXT | NULL | 正文内容 |
| cover_color | VARCHAR(32) | NULL | 封面背景色 |
| author | VARCHAR(64) | NULL | 作者 |
| view_count | INT | DEFAULT 0 | 浏览数 |
| published_at | DATETIME | NULL | 发布时间 |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |

**索引**：`idx_articles_category` (category)，`idx_articles_published_at` (published_at)

### 3.3 learning_courses — 学习课程

| 字段 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT UNSIGNED | PRIMARY KEY AUTO_INCREMENT | 主键 |
| title | VARCHAR(128) | NOT NULL | 课程名称 |
| description | TEXT | NULL | 课程介绍 |
| category | VARCHAR(32) | NULL | 分类（剪辑教程 / 色彩学 / AI创作 等） |
| level | VARCHAR(16) | NULL | 难度：入门 / 进阶 |
| duration | VARCHAR(32) | NULL | 时长描述（如"12课时"） |
| lessons | INT | DEFAULT 0 | 课时数 |
| icon | VARCHAR(32) | NULL | 图标 key |
| is_active | TINYINT(1) | DEFAULT 1 | 是否上架 |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |

**索引**：`idx_learning_courses_category` (category)

### 3.4 learning_paths — 学习路径

| 字段 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT UNSIGNED | PRIMARY KEY AUTO_INCREMENT | 主键 |
| name | VARCHAR(128) | NOT NULL | 路径名称 |
| description | TEXT | NULL | 路径描述 |
| course_ids | TEXT | NULL | 关联课程 ID，逗号分隔字符串 |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |

### 3.5 ai_tools — AI 工具推荐

| 字段 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT UNSIGNED | PRIMARY KEY AUTO_INCREMENT | 主键 |
| name | VARCHAR(64) | NOT NULL | 工具名 |
| description | TEXT | NULL | 工具描述 |
| category | VARCHAR(32) | NULL | 分类：video / image / audio |
| icon | VARCHAR(32) | NULL | 图标 key |
| tags | VARCHAR(128) | NULL | 标签，逗号分隔 |
| url | VARCHAR(256) | NULL | 官网链接 |
| is_free | TINYINT(1) | DEFAULT 1 | 是否免费 |
| featured | TINYINT(1) | DEFAULT 0 | 是否精选推荐 |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |

**索引**：`idx_ai_tools_category` (category)，`idx_ai_tools_featured` (featured)

### 3.6 newsletters — 邮件订阅

| 字段 | 类型 | 约束 | 说明 |
|---|---|---|---|
| id | BIGINT UNSIGNED | PRIMARY KEY AUTO_INCREMENT | 主键 |
| email | VARCHAR(256) | NOT NULL UNIQUE | 邮箱地址 |
| active | TINYINT(1) | DEFAULT 1 | 订阅状态 |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | 订阅时间 |

**索引**：`uk_newsletters_email` (email) UNIQUE

## 四、字段类型与 SQLite 的映射差异

| SQLite 类型 | MySQL 选择 | 原因 |
|---|---|---|
| `INTEGER PRIMARY KEY` | `BIGINT UNSIGNED AUTO_INCREMENT` | 预留足够 ID 空间，避免 int32 上限风险 |
| `REAL` / `float64` rating | `DECIMAL(3,1)` | 评分定点存储，避免浮点精度问题 |
| `TEXT` content 正文 | `MEDIUMTEXT` | 文章正文可能超过 64KB，MEDIUMTEXT 最大 16MB |
| `INTEGER` bool | `TINYINT(1)` | 与 GORM 默认 MySQL 布尔映射一致 |
| `DATETIME` (Go time.Time) | `DATETIME` | DSN 中 `parseTime=True` 直接映射 |
| SQLite `AUTOINCREMENT` | MySQL `AUTO_INCREMENT` | 语义等价 |

## 五、字符集与排序规则

- `CREATE DATABASE yigeyingshi CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;`
- 所有表 `ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;`
- 连接串追加 `charset=utf8mb4`，GORM 会自动在建连接后设置 `SET NAMES utf8mb4`

## 六、GORM DSN 配置

```
# yige-server
DB_DRIVER=mysql
DB_DSN=root:root@tcp(118.145.113.88:3306)/yigeyingshi?charset=utf8mb4&parseTime=True&loc=Local

# yige-admin-server
ADMIN_DB_DSN=root:root@tcp(118.145.113.88:3306)/yigeyingshi?charset=utf8mb4&parseTime=True&loc=Local
```

两个服务的 `parseTime=True` 和 `loc=Local` 必不可少，否则 `time.Time` 字段无法正确扫描。

## 七、测试数据范围（随脚本一并插入）

- movies：22 部精选影片（含 is_featured 标记 4 部）
- articles：7 篇（影评 2 篇 + AI 文章 5 篇）
- learning_courses：6 门课程（剪辑 / 色彩 / 调色 / AI / 音效 / 拍摄）
- learning_paths：2 条（"从零到剪辑师"、"AI创作进阶"）
- ai_tools：3 个精选工具（Runway / Midjourney / ElevenLabs）
- newsletters：留空，仅建表，由真实用户订阅行为写入

## 八、连接注意事项

1. `yige-server` 与 `yige-admin-server` 不再各自持有本地 SQLite 文件，改为共同连接远程 MySQL，无需 WAL 与 busy_timeout 配置（MySQL 通过 MVCC + 行锁解决并发写入）。
2. `newsletters` 表的 `email` 列加了唯一索引，前台订阅接口重复提交时会报唯一键冲突，需在 Go handler 中捕获并返回友好提示（已有 SQLite 场景下的验证逻辑基本可复用）。
