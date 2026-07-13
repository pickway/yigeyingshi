package service

import (
	"log"
	"time"

	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&model.Movie{}).Count(&count)
	if count > 0 {
		log.Println("[seed] data already exists, skip seeding")
		return
	}

	log.Println("[seed] inserting initial data...")
	now := time.Now()

	movies := []model.Movie{
		{Title: "星际穿越", Year: 2014, Genres: "科幻,剧情", Rating: 9.4, Description: "一场穿越虫洞的星际冒险，关于爱、时间与人类命运的恢弘史诗。", PosterColor: "#2C3E6B", PosterIcon: "orbit", Director: "克里斯托弗·诺兰", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
		{Title: "千与千寻", Year: 2001, Genres: "动画,剧情", Rating: 9.4, Description: "少女千寻误入神灵世界，在油屋中成长与寻找自我的奇幻旅程。", PosterColor: "#6B4E8B", PosterIcon: "sparkles", Director: "宫崎骏", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
		{Title: "寄生虫", Year: 2019, Genres: "剧情,悬疑", Rating: 8.8, Description: "一户底层家庭逐步渗透富裕家庭的黑色寓言，揭露社会阶层鸿沟。", PosterColor: "#4A6741", PosterIcon: "home", Director: "奉俊昊", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
		{Title: "盗梦空间", Year: 2010, Genres: "科幻,悬疑", Rating: 9.3, Description: "在梦境的层层嵌套中，盗取与植入意念的高智商动作悬疑片。", PosterColor: "#3A5A8C", PosterIcon: "layers", Director: "克里斯托弗·诺兰", CreatedAt: now, UpdatedAt: now},
		{Title: "龙猫", Year: 1988, Genres: "动画,喜剧", Rating: 9.2, Description: "姐妹俩在乡间邂逅森林精灵龙猫的温暖童话。", PosterColor: "#6B8E5A", PosterIcon: "trees", Director: "宫崎骏", CreatedAt: now, UpdatedAt: now},
		{Title: "杀人回忆", Year: 2003, Genres: "悬疑,剧情", Rating: 8.9, Description: "基于真实悬案改编，讲述韩国小镇连环杀人案的调查过程。", PosterColor: "#8B6F4E", PosterIcon: "search", Director: "奉俊昊", CreatedAt: now, UpdatedAt: now},
		{Title: "你的名字", Year: 2016, Genres: "动画,剧情", Rating: 8.4, Description: "少年少女在梦中交换身体，跨越时空追寻彼此的青春故事。", PosterColor: "#E87D6B", PosterIcon: "flame", Director: "新海诚", CreatedAt: now, UpdatedAt: now},
		{Title: "肖申克的救赎", Year: 1994, Genres: "剧情", Rating: 9.7, Description: "银行家安迪在冤狱之中用二十年凿出自由之路，关于希望与友情的经典。", PosterColor: "#5A6E82", PosterIcon: "bird", Director: "弗兰克·德拉邦特", CreatedAt: now, UpdatedAt: now},
		{Title: "疯狂动物城", Year: 2016, Genres: "动画,喜剧", Rating: 9.2, Description: "兔子警官朱迪与狐狸尼克联手破获动物城失踪案的爆笑冒险。", PosterColor: "#D4853A", PosterIcon: "rabbit", Director: "拜恩·霍华德", CreatedAt: now, UpdatedAt: now},
		{Title: "奥本海默", Year: 2023, Genres: "传记,历史", Rating: 8.9, Description: "原子弹之父罗伯特·奥本海默的一生，科学与道德的永恒拷问。", PosterColor: "#2A1F1A", PosterIcon: "zap", Director: "克里斯托弗·诺兰", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
	}
	db.Create(&movies)

	articles := []model.Article{
		{Title: "诺兰的时间叙事美学", Category: "movie-review", Tag: "影评", Summary: "从《记忆碎片》到《奥本海默》，诺兰如何用非线性结构重构电影时间。", CoverColor: "#2C3E6B", Author: "一个影视", PublishedAt: time.Date(2026, 7, 10, 0, 0, 0, 0, time.UTC), CreatedAt: now, UpdatedAt: now},
		{Title: "宫崎骏的色彩哲学", Category: "movie-review", Tag: "观察笔记", Summary: "绿色、蓝色与橙色——宫崎骏电影中反复出现的色彩符号及其意义。", CoverColor: "#6B4E8B", Author: "一个影视", PublishedAt: time.Date(2026, 7, 6, 0, 0, 0, 0, time.UTC), CreatedAt: now, UpdatedAt: now},
		{Title: "AI视频生成的五大趋势", Category: "ai-article", Tag: "趋势", Summary: "从文本到视频，从2D到3D，AI视频生成正在重塑内容创作的边界。", CoverColor: "#E87D6B", Author: "一个影视", PublishedAt: time.Date(2025, 6, 12, 0, 0, 0, 0, time.UTC), CreatedAt: now, UpdatedAt: now},
		{Title: "如何用AI辅助剧本创作", Category: "ai-article", Tag: "创作", Summary: "利用大语言模型进行头脑风暴、角色设计与情节推演的实操指南。", CoverColor: "#D4853A", Author: "一个影视", PublishedAt: time.Date(2025, 5, 28, 0, 0, 0, 0, time.UTC), CreatedAt: now, UpdatedAt: now},
		{Title: "Stable Diffusion在影视中的应用", Category: "ai-article", Tag: "应用", Summary: "概念图、分镜、氛围图——AI图像生成如何加速前期制作流程。", CoverColor: "#6B8E5A", Author: "一个影视", PublishedAt: time.Date(2025, 5, 15, 0, 0, 0, 0, time.UTC), CreatedAt: now, UpdatedAt: now},
		{Title: "AI配音与人工配音对比", Category: "ai-article", Tag: "对比", Summary: "在什么场景下AI配音可以替代人工，什么场景下仍然有差距。", CoverColor: "#3A5A8C", Author: "一个影视", PublishedAt: time.Date(2025, 4, 30, 0, 0, 0, 0, time.UTC), CreatedAt: now, UpdatedAt: now},
		{Title: "从Sora到Kling：视频模型演进", Category: "ai-article", Tag: "深度", Summary: "一文看懂视频生成模型的技术路线图与未来发展方向。", CoverColor: "#8B6F4E", Author: "一个影视", PublishedAt: time.Date(2025, 4, 18, 0, 0, 0, 0, time.UTC), CreatedAt: now, UpdatedAt: now},
	}
	db.Create(&articles)

	courses := []model.LearningCourse{
		{Title: "Premiere Pro 快速入门", Description: "从零开始掌握专业剪辑软件，从界面认知到完整短片输出。", Category: "剪辑教程", Level: "入门", Duration: "12课时", Lessons: 12, Icon: "clapperboard", CreatedAt: now, UpdatedAt: now},
		{Title: "电影色彩理论基础", Description: "理解色彩的情感表达，掌握电影调色的底层逻辑。", Category: "色彩学", Level: "入门", Duration: "8课时", Lessons: 8, Icon: "palette", CreatedAt: now, UpdatedAt: now},
		{Title: "DaVinci Resolve 调色实战", Description: "专业调色软件深度教程，案例驱动式学习。", Category: "色彩学", Level: "进阶", Duration: "16课时", Lessons: 16, Icon: "sliders-horizontal", CreatedAt: now, UpdatedAt: now},
		{Title: "AI 视频生成入门", Description: "Runway、Sora、可灵等主流AI视频工具全解析。", Category: "AI创作", Level: "入门", Duration: "10课时", Lessons: 10, Icon: "sparkles", CreatedAt: now, UpdatedAt: now},
		{Title: "影视音效设计指南", Description: "从环境音到配乐，打造身临其境的声音世界。", Category: "音效设计", Level: "进阶", Duration: "14课时", Lessons: 14, Icon: "volume-2", CreatedAt: now, UpdatedAt: now},
		{Title: "短视频拍摄技巧", Description: "手机也能拍大片？掌握构图、运镜与布光基础。", Category: "影视制作", Level: "入门", Duration: "9课时", Lessons: 9, Icon: "video", CreatedAt: now, UpdatedAt: now},
	}
	db.Create(&courses)

	paths := []model.LearningPath{
		{Name: "从零到剪辑师", Description: "从拍摄到剪辑到调色，完整的视频制作入门路径。", CourseIDs: "6,1,2,3", CreatedAt: now, UpdatedAt: now},
		{Name: "AI创作进阶", Description: "用AI工具重塑你的创作工作流。", CourseIDs: "4", CreatedAt: now, UpdatedAt: now},
	}
	db.Create(&paths)

	tools := []model.AiTool{
		{Name: "Runway ML", Description: "革命性的 AI 视频生成工具，支持文本转视频与智能剪辑。", Category: "video", Icon: "video", Tags: "免费,视频生成", Url: "https://runwayml.com", IsFree: true, Featured: true, CreatedAt: now, UpdatedAt: now},
		{Name: "Midjourney", Description: "顶级 AI 图像生成平台，创造令人惊叹的视觉概念与分镜。", Category: "image", Icon: "image", Tags: "付费,图像生成", Url: "https://midjourney.com", IsFree: false, Featured: true, CreatedAt: now, UpdatedAt: now},
		{Name: "ElevenLabs", Description: "AI 驱动的语音合成与音频处理，精准还原自然人声。", Category: "audio", Icon: "audio-waveform", Tags: "免费,音频处理", Url: "https://elevenlabs.io", IsFree: true, Featured: true, CreatedAt: now, UpdatedAt: now},
	}
	db.Create(&tools)

	log.Println("[seed] done")
}
