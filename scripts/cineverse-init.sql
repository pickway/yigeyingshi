-- ================================================================
-- CineVerse MySQL 初始化脚本
-- 适用版本：MySQL 8.0+
-- 执行方式：
--   mysql -h 118.145.113.88 -P 3306 -u root -proot < scripts/cineverse-init.sql
-- 或在 MySQL 客户端：SOURCE scripts/cineverse-init.sql;
-- ================================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------------------------------------------
-- 1. 创建数据库（若不存在）
-- ----------------------------------------------------------------
CREATE DATABASE IF NOT EXISTS `yigeyingshi`
  DEFAULT CHARACTER SET utf8mb4
  DEFAULT COLLATE utf8mb4_general_ci;

USE `yigeyingshi`;

-- ----------------------------------------------------------------
-- 2. movies — 影片资源
-- ----------------------------------------------------------------
DROP TABLE IF EXISTS `movies`;
CREATE TABLE `movies` (
  `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title`          VARCHAR(128)    NOT NULL                COMMENT '影片标题',
  `original_title` VARCHAR(128)    DEFAULT NULL            COMMENT '原始标题',
  `year`           INT             NOT NULL DEFAULT 0      COMMENT '上映年份',
  `genres`         VARCHAR(128)    DEFAULT NULL            COMMENT '类型（逗号分隔）',
  `rating`         DECIMAL(3,1)    NOT NULL DEFAULT 0.0    COMMENT '评分 0.0-10.0',
  `description`    TEXT            DEFAULT NULL            COMMENT '简介',
  `poster_color`   VARCHAR(32)     DEFAULT NULL            COMMENT '海报背景色',
  `poster_icon`    VARCHAR(32)     DEFAULT NULL            COMMENT '海报图标 key',
  `duration`       INT             NOT NULL DEFAULT 0      COMMENT '片长（分钟）',
  `director`       VARCHAR(64)     DEFAULT NULL            COMMENT '导演',
  `is_featured`    TINYINT(1)      NOT NULL DEFAULT 0      COMMENT '是否首页精选',
  `created_at`     DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at`     DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_movies_featured` (`is_featured`),
  KEY `idx_movies_year` (`year`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='影片资源表';

-- ----------------------------------------------------------------
-- 3. articles — 文章
-- ----------------------------------------------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title`        VARCHAR(256)    NOT NULL                COMMENT '标题',
  `category`     VARCHAR(32)     NOT NULL                COMMENT '分类：movie-review / learning-note / ai-article',
  `tag`          VARCHAR(32)     DEFAULT NULL            COMMENT '二级标签',
  `summary`      TEXT            DEFAULT NULL            COMMENT '摘要',
  `content`      MEDIUMTEXT      DEFAULT NULL            COMMENT '正文',
  `cover_color`  VARCHAR(32)     DEFAULT NULL            COMMENT '封面背景色',
  `author`       VARCHAR(64)     DEFAULT NULL            COMMENT '作者',
  `view_count`   INT             NOT NULL DEFAULT 0      COMMENT '浏览数',
  `published_at` DATETIME        DEFAULT NULL            COMMENT '发布时间',
  `created_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_articles_category` (`category`),
  KEY `idx_articles_published_at` (`published_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='文章表';

-- ----------------------------------------------------------------
-- 4. learning_courses — 学习课程
-- ----------------------------------------------------------------
DROP TABLE IF EXISTS `learning_courses`;
CREATE TABLE `learning_courses` (
  `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title`       VARCHAR(128)    NOT NULL                COMMENT '课程名称',
  `description` TEXT            DEFAULT NULL            COMMENT '课程介绍',
  `category`    VARCHAR(32)     DEFAULT NULL            COMMENT '分类',
  `level`       VARCHAR(16)     DEFAULT NULL            COMMENT '难度：入门/进阶',
  `duration`    VARCHAR(32)     DEFAULT NULL            COMMENT '时长描述',
  `lessons`     INT             NOT NULL DEFAULT 0      COMMENT '课时数',
  `icon`        VARCHAR(32)     DEFAULT NULL            COMMENT '图标 key',
  `is_active`   TINYINT(1)      NOT NULL DEFAULT 1      COMMENT '是否上架',
  `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_learning_courses_category` (`category`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='学习课程表';

-- ----------------------------------------------------------------
-- 5. learning_paths — 学习路径
-- ----------------------------------------------------------------
DROP TABLE IF EXISTS `learning_paths`;
CREATE TABLE `learning_paths` (
  `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(128)    NOT NULL                COMMENT '路径名称',
  `description` TEXT            DEFAULT NULL            COMMENT '路径描述',
  `course_ids`  TEXT            DEFAULT NULL            COMMENT '关联课程 ID（逗号分隔）',
  `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='学习路径表';

-- ----------------------------------------------------------------
-- 6. ai_tools — AI 工具推荐
-- ----------------------------------------------------------------
DROP TABLE IF EXISTS `ai_tools`;
CREATE TABLE `ai_tools` (
  `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(64)     NOT NULL                COMMENT '工具名',
  `description` TEXT            DEFAULT NULL            COMMENT '工具描述',
  `category`    VARCHAR(32)     DEFAULT NULL            COMMENT '分类：video/image/audio',
  `icon`        VARCHAR(32)     DEFAULT NULL            COMMENT '图标 key',
  `tags`        VARCHAR(128)    DEFAULT NULL            COMMENT '标签（逗号分隔）',
  `url`         VARCHAR(256)    DEFAULT NULL            COMMENT '官网链接',
  `is_free`     TINYINT(1)      NOT NULL DEFAULT 1      COMMENT '是否免费',
  `featured`    TINYINT(1)      NOT NULL DEFAULT 0      COMMENT '是否精选推荐',
  `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_ai_tools_category` (`category`),
  KEY `idx_ai_tools_featured` (`featured`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='AI工具推荐表';

-- ----------------------------------------------------------------
-- 7. newsletters — 邮件订阅
-- ----------------------------------------------------------------
DROP TABLE IF EXISTS `newsletters`;
CREATE TABLE `newsletters` (
  `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `email`      VARCHAR(256)    NOT NULL                COMMENT '邮箱地址',
  `active`     TINYINT(1)      NOT NULL DEFAULT 1      COMMENT '订阅状态',
  `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订阅时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_newsletters_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='邮件订阅表';

-- ================================================================
-- 8. 插入测试数据
-- ================================================================

SET @NOW = NOW();

-- 8.1 movies（22 部）
INSERT INTO `movies`
(`title`,`original_title`,`year`,`genres`,`rating`,`description`,`poster_color`,`poster_icon`,`duration`,`director`,`is_featured`,`created_at`,`updated_at`) VALUES
('星际穿越',        'Interstellar',          2014, '科幻,剧情',        9.4, '穿越虫洞的父女之约，时间与爱的相对论史诗。',              '#2C3E6B', 'orbit',          169, '克里斯托弗·诺兰',   1, @NOW, @NOW),
('千与千寻',        '千と千尋の神隠し',      2001, '动画,剧情',        9.3, '奇幻浴场中的成长之旅，宫崎骏的童话世界。',                '#6B4E8B', 'sparkles',       125, '宫崎骏',           1, @NOW, @NOW),
('寄生虫',          '기생충',                2019, '剧情,悬疑',        9.1, '阶层差异下的黑色寓言，人性与欲望的博弈。',                '#4A6741', 'home',           132, '奉俊昊',           1, @NOW, @NOW),
('盗梦空间',        'Inception',             2010, '科幻,悬疑',        9.3, '梦中梦的 labyrinth，意识深处的终极冒险。',              '#3A5A8C', 'layers',         148, '克里斯托弗·诺兰',   0, @NOW, @NOW),
('龙猫',            'となりのトトロ',        1988, '动画,喜剧',        9.2, '乡间田野的神奇邂逅，童年最温暖的守护。',                '#6B8E5A', 'trees',           86, '宫崎骏',           0, @NOW, @NOW),
('杀人回忆',        '살인의 추억',           2003, '悬疑,剧情',        9.0, '韩国小镇的连环悬案，真实事件改编的震撼。',              '#8B6F4E', 'search',         130, '奉俊昊',           0, @NOW, @NOW),
('你的名字',        '君の名は。',             2016, '动画,剧情',        8.9, '跨越时空的灵魂交错，新海诚的视觉盛宴。',                '#E87D6B', 'flame',          106, '新海诚',           0, @NOW, @NOW),
('肖申克的救赎',    'The Shawshank Redemption',1994,'剧情',              9.7, '黑暗牢笼中的希望之光，自由灵魂的永恒颂歌。',            '#5A6E82', 'bird',           142, '弗兰克·德拉邦特',   0, @NOW, @NOW),
('疯狂动物城',      'Zootopia',              2016, '动画,喜剧',        9.2, '狐兔搭档破解阴谋，迪士尼的温情讽刺。',                  '#D4853A', 'rabbit',         108, '拜恩·霍华德',       0, @NOW, @NOW),
('奥本海默',        'Oppenheimer',           2023, '传记,历史',        8.9, '原子弹之父罗伯特·奥本海默的一生，科学与道德的永恒拷问。', '#2A1F1A', 'zap',            180, '克里斯托弗·诺兰',   1, @NOW, @NOW),
('沙丘',            'Dune',                  2021, '科幻,冒险',        8.8, '沙漠星球上的权力更迭，维伦纽瓦的史诗巨制。',            '#2A1F0D', 'sun',            155, '丹尼斯·维伦纽瓦',   0, @NOW, @NOW),
('布达佩斯大饭店',  'The Grand Budapest Hotel',2014,'喜剧,剧情',       8.9, '韦斯·安德森的色彩童话，欧洲旧时光的挽歌。',              '#2D2D1A', 'hotel',           99, '韦斯·安德森',       0, @NOW, @NOW),
('燃烧女子的肖像',  'Portrait de la jeune fille en feu',2019,'剧情,爱情',8.7,'海边悬崖上的凝视与爱，女性视角的古典浪漫。',            '#2A1A1A', 'flame',          122, '瑟琳·席安玛',       0, @NOW, @NOW),
('银翼杀手2049',    'Blade Runner 2049',     2017, '科幻,悬疑',        8.6, '赛博朋克的视觉巅峰，孤独与存在的深邃冥想。',            '#0D0D1A', 'cpu',            164, '丹尼斯·维伦纽瓦',   0, @NOW, @NOW),
('小偷家族',        '万引き家族',            2018, '剧情,家庭',        9.0, '是枝裕和的温情凝视，非血缘家庭的羁绊与告别。',          '#1A1A1A', 'heart',          121, '是枝裕和',           0, @NOW, @NOW),
('爱乐之城',        'La La Land',            2016, '剧情,爱情,音乐',   8.8, '洛杉矶的星光与爵士，梦想与爱情的浪漫交响。',            '#0D1A2A', 'music',          128, '达米恩·查泽雷',     0, @NOW, @NOW),
('降临',            'Arrival',               2016, '科幻,剧情',        8.7, '语言的尽头是时间，维伦纽瓦的科幻诗意之作。',            '#1A1A0D', 'message-circle', 116, '丹尼斯·维伦纽瓦',   0, @NOW, @NOW),
('楚门的世界',      'The Truman Show',       1998, '剧情,科幻',        9.3, '如果你无法在现实世界找到他，那就去他心里找。',          '#1A2A1A', 'tv',             103, '彼得·威尔',         0, @NOW, @NOW),
('黑暗骑士',        'The Dark Knight',       2008, '动作,犯罪',        9.2, '希斯·莱杰的绝唱，秩序与混沌的终极对决。',              '#0D0D0D', 'moon',           152, '克里斯托弗·诺兰',   0, @NOW, @NOW),
('摔跤吧！爸爸',    'Dangal',                2016, '剧情,运动',        9.0, '父亲的摔跤梦，女儿的自由之路，印度励志经典。',          '#2A1A0D', 'trophy',         161, '尼特什·提瓦瑞',     0, @NOW, @NOW),
('霸王别姬',        '霸王别姬',              1993, '剧情,爱情',        9.6, '不疯魔不成活，程蝶衣的一生就是一出戏。',                '#2A0D0D', 'drama',          171, '陈凯歌',           0, @NOW, @NOW),
('头号玩家',        'Ready Player One',      2018, '科幻,冒险',        8.7, '斯皮尔伯格的彩蛋狂欢，致敬流行文化的冒险之旅。',        '#0D1A1A', 'gamepad-2',      140, '史蒂文·斯皮尔伯格', 0, @NOW, @NOW);

-- 8.2 articles（7 篇）
INSERT INTO `articles`
(`title`,`category`,`tag`,`summary`,`content`,`cover_color`,`author`,`view_count`,`published_at`,`created_at`,`updated_at`) VALUES
('诺兰的时间叙事美学',            'movie-review',  '影评',     '从《记忆碎片》到《奥本海默》，诺兰如何用非线性结构重构电影时间。',
  '# 诺兰的时间叙事美学\n\n克里斯托弗·诺兰的电影总是以复杂的时间结构闻名……',
  '#2C3E6B', '一个影视', 0, '2026-07-10 00:00:00', @NOW, @NOW),
('宫崎骏的色彩哲学',              'movie-review',  '观察笔记', '绿色、蓝色与橙色——宫崎骏电影中反复出现的色彩符号及其意义。',
  '# 宫崎骏的色彩哲学\n\n宫崎骏的画面总是充满温柔的色彩……',
  '#6B4E8B', '一个影视', 0, '2026-07-06 00:00:00', @NOW, @NOW),
('AI视频生成的五大趋势',          'ai-article',    '趋势',     '从文本到视频，从2D到3D，AI视频生成正在重塑内容创作的边界。',
  '# AI视频生成的五大趋势\n\n1. 多模态提示\n2. 3D一致性……',
  '#E87D6B', '一个影视', 0, '2025-06-12 00:00:00', @NOW, @NOW),
('如何用AI辅助剧本创作',          'ai-article',    '创作',     '利用大语言模型进行头脑风暴、角色设计与情节推演的实操指南。',
  '# 如何用AI辅助剧本创作\n\n头脑风暴阶段，可以让 LLM 提供 10 个不同的故事钩子……',
  '#D4853A', '一个影视', 0, '2025-05-28 00:00:00', @NOW, @NOW),
('Stable Diffusion在影视中的应用','ai-article',   '应用',     '概念图、分镜、氛围图——AI图像生成如何加速前期制作流程。',
  '# Stable Diffusion 在影视中的应用\n\n前期制作阶段最耗时的就是概念图与分镜绘制……',
  '#6B8E5A', '一个影视', 0, '2025-05-15 00:00:00', @NOW, @NOW),
('AI配音与人工配音对比',          'ai-article',    '对比',     '在什么场景下AI配音可以替代人工，什么场景下仍然有差距。',
  '# AI 配音与人工配音对比\n\n情绪表达、上下文理解、小语种……',
  '#3A5A8C', '一个影视', 0, '2025-04-30 00:00:00', @NOW, @NOW),
('从Sora到Kling：视频模型演进',   'ai-article',    '深度',     '一文看懂视频生成模型的技术路线图与未来发展方向。',
  '# 从 Sora 到 Kling：视频模型演进\n\n扩散模型、DiT、Motion Module……',
  '#8B6F4E', '一个影视', 0, '2025-04-18 00:00:00', @NOW, @NOW);

-- 8.3 learning_courses（6 门）
INSERT INTO `learning_courses`
(`title`,`description`,`category`,`level`,`duration`,`lessons`,`icon`,`is_active`,`created_at`,`updated_at`) VALUES
('Premiere Pro 快速入门',    '从零开始掌握专业剪辑软件，从界面认知到完整短片输出。',           '剪辑教程',  '入门', '12课时', 12, 'clapperboard',      1, @NOW, @NOW),
('电影色彩理论基础',         '理解色彩的情感表达，掌握电影调色的底层逻辑。',                   '色彩学',    '入门', '8课时',   8, 'palette',           1, @NOW, @NOW),
('DaVinci Resolve 调色实战', '专业调色软件深度教程，案例驱动式学习。',                         '色彩学',    '进阶', '16课时', 16, 'sliders-horizontal',1, @NOW, @NOW),
('AI 视频生成入门',          'Runway、Sora、可灵等主流AI视频工具全解析。',                     'AI创作',    '入门', '10课时', 10, 'sparkles',          1, @NOW, @NOW),
('影视音效设计指南',         '从环境音到配乐，打造身临其境的声音世界。',                       '音效设计',  '进阶', '14课时', 14, 'volume-2',          1, @NOW, @NOW),
('短视频拍摄技巧',           '手机也能拍大片？掌握构图、运镜与布光基础。',                     '影视制作',  '入门', '9课时',   9, 'video',             1, @NOW, @NOW);

-- 8.4 learning_paths（2 条）
INSERT INTO `learning_paths`
(`name`,`description`,`course_ids`,`created_at`,`updated_at`) VALUES
('从零到剪辑师', '从拍摄到剪辑到调色，完整的视频制作入门路径。',      '6,1,2,3', @NOW, @NOW),
('AI创作进阶',   '用AI工具重塑你的创作工作流。',                      '4',       @NOW, @NOW);

-- 8.5 ai_tools（3 个）
INSERT INTO `ai_tools`
(`name`,`description`,`category`,`icon`,`tags`,`url`,`is_free`,`featured`,`created_at`,`updated_at`) VALUES
('Runway ML',    '革命性的 AI 视频生成工具，支持文本转视频与智能剪辑。',               'video', 'video',           '免费,视频生成',   'https://runwayml.com',      1, 1, @NOW, @NOW),
('Midjourney',   '顶级 AI 图像生成平台，创造令人惊叹的视觉概念与分镜。',               'image', 'image',           '付费,图像生成',   'https://midjourney.com',    0, 1, @NOW, @NOW),
('ElevenLabs',   'AI 驱动的语音合成与音频处理，精准还原自然人声。',                   'audio', 'audio-waveform',  '免费,音频处理',   'https://elevenlabs.io',     1, 1, @NOW, @NOW);

-- 8.6 newsletters：留空（等待真实订阅）

SET FOREIGN_KEY_CHECKS = 1;

-- ================================================================
-- 9. 数据校验（可手动执行以下 SELECT 验收）
--   SELECT COUNT(*) FROM movies;           -- 期望 22
--   SELECT COUNT(*) FROM articles;         -- 期望 7
--   SELECT COUNT(*) FROM learning_courses; -- 期望 6
--   SELECT COUNT(*) FROM learning_paths;   -- 期望 2
--   SELECT COUNT(*) FROM ai_tools;         -- 期望 3
--   SELECT COUNT(*) FROM newsletters;      -- 期望 0
-- ================================================================
