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
		{Title: "星际穿越", Year: 2014, Genres: "科幻,剧情", Rating: 9.4, Description: "穿越虫洞的父女之约，时间与爱的相对论史诗。", PosterColor: "#2C3E6B", PosterIcon: "orbit", Director: "克里斯托弗·诺兰", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
		{Title: "千与千寻", Year: 2001, Genres: "动画,剧情", Rating: 9.3, Description: "奇幻浴场中的成长之旅，宫崎骏的童话世界。", PosterColor: "#6B4E8B", PosterIcon: "sparkles", Director: "宫崎骏", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
		{Title: "寄生虫", Year: 2019, Genres: "剧情,悬疑", Rating: 9.1, Description: "阶层差异下的黑色寓言，人性与欲望的博弈。", PosterColor: "#4A6741", PosterIcon: "home", Director: "奉俊昊", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
		{Title: "盗梦空间", Year: 2010, Genres: "科幻,悬疑", Rating: 9.3, Description: "梦中梦的 labyrinth，意识深处的终极冒险。", PosterColor: "#3A5A8C", PosterIcon: "layers", Director: "克里斯托弗·诺兰", CreatedAt: now, UpdatedAt: now},
		{Title: "龙猫", Year: 1988, Genres: "动画,喜剧", Rating: 9.2, Description: "乡间田野的神奇邂逅，童年最温暖的守护。", PosterColor: "#6B8E5A", PosterIcon: "trees", Director: "宫崎骏", CreatedAt: now, UpdatedAt: now},
		{Title: "杀人回忆", Year: 2003, Genres: "悬疑,剧情", Rating: 9.0, Description: "韩国小镇的连环悬案，真实事件改编的震撼。", PosterColor: "#8B6F4E", PosterIcon: "search", Director: "奉俊昊", CreatedAt: now, UpdatedAt: now},
		{Title: "你的名字", Year: 2016, Genres: "动画,剧情", Rating: 8.9, Description: "跨越时空的灵魂交错，新海诚的视觉盛宴。", PosterColor: "#E87D6B", PosterIcon: "flame", Director: "新海诚", CreatedAt: now, UpdatedAt: now},
		{Title: "肖申克的救赎", Year: 1994, Genres: "剧情", Rating: 9.7, Description: "黑暗牢笼中的希望之光，自由灵魂的永恒颂歌。", PosterColor: "#5A6E82", PosterIcon: "bird", Director: "弗兰克·德拉邦特", CreatedAt: now, UpdatedAt: now},
		{Title: "疯狂动物城", Year: 2016, Genres: "动画,喜剧", Rating: 9.2, Description: "狐兔搭档破解阴谋，迪士尼的温情讽刺。", PosterColor: "#D4853A", PosterIcon: "rabbit", Director: "拜恩·霍华德", CreatedAt: now, UpdatedAt: now},
		{Title: "奥本海默", Year: 2023, Genres: "传记,历史", Rating: 8.9, Description: "原子弹之父罗伯特·奥本海默的一生，科学与道德的永恒拷问。", PosterColor: "#2A1F1A", PosterIcon: "zap", Director: "克里斯托弗·诺兰", IsFeatured: true, CreatedAt: now, UpdatedAt: now},
		{Title: "沙丘", Year: 2021, Genres: "科幻,冒险", Rating: 8.8, Description: "沙漠星球上的权力更迭，维伦纽瓦的史诗巨制。", PosterColor: "#2A1F0D", PosterIcon: "sun", Director: "丹尼斯·维伦纽瓦", CreatedAt: now, UpdatedAt: now},
		{Title: "布达佩斯大饭店", Year: 2014, Genres: "喜剧,剧情", Rating: 8.9, Description: "韦斯·安德森的色彩童话，欧洲旧时光的挽歌。", PosterColor: "#2D2D1A", PosterIcon: "hotel", Director: "韦斯·安德森", CreatedAt: now, UpdatedAt: now},
		{Title: "燃烧女子的肖像", Year: 2019, Genres: "剧情,爱情", Rating: 8.7, Description: "海边悬崖上的凝视与爱，女性视角的古典浪漫。", PosterColor: "#2A1A1A", PosterIcon: "flame", Director: "瑟琳·席安玛", CreatedAt: now, UpdatedAt: now},
		{Title: "银翼杀手2049", Year: 2017, Genres: "科幻,悬疑", Rating: 8.6, Description: "赛博朋克的视觉巅峰，孤独与存在的深邃冥想。", PosterColor: "#0D0D1A", PosterIcon: "cpu", Director: "丹尼斯·维伦纽瓦", CreatedAt: now, UpdatedAt: now},
		{Title: "小偷家族", Year: 2018, Genres: "剧情,家庭", Rating: 9.0, Description: "是枝裕和的温情凝视，非血缘家庭的羁绊与告别。", PosterColor: "#1A1A1A", PosterIcon: "heart", Director: "是枝裕和", CreatedAt: now, UpdatedAt: now},
		{Title: "爱乐之城", Year: 2016, Genres: "剧情,爱情,音乐", Rating: 8.8, Description: "洛杉矶的星光与爵士，梦想与爱情的浪漫交响。", PosterColor: "#0D1A2A", PosterIcon: "music", Director: "达米恩·查泽雷", CreatedAt: now, UpdatedAt: now},
		{Title: "降临", Year: 2016, Genres: "科幻,剧情", Rating: 8.7, Description: "语言的尽头是时间，维伦纽瓦的科幻诗意之作。", PosterColor: "#1A1A0D", PosterIcon: "message-circle", Director: "丹尼斯·维伦纽瓦", CreatedAt: now, UpdatedAt: now},
		{Title: "楚门的世界", Year: 1998, Genres: "剧情,科幻", Rating: 9.3, Description: "如果你无法在现实世界找到他，那就去他心里找。", PosterColor: "#1A2A1A", PosterIcon: "tv", Director: "彼得·威尔", CreatedAt: now, UpdatedAt: now},
		{Title: "黑暗骑士", Year: 2008, Genres: "动作,犯罪", Rating: 9.2, Description: "希斯·莱杰的绝唱，秩序与混沌的终极对决。", PosterColor: "#0D0D0D", PosterIcon: "moon", Director: "克里斯托弗·诺兰", CreatedAt: now, UpdatedAt: now},
		{Title: "摔跤吧！爸爸", Year: 2016, Genres: "剧情,运动", Rating: 9.0, Description: "父亲的摔跤梦，女儿的自由之路，印度励志经典。", PosterColor: "#2A1A0D", PosterIcon: "trophy", Director: "尼特什·提瓦瑞", CreatedAt: now, UpdatedAt: now},
		{Title: "霸王别姬", Year: 1993, Genres: "剧情,爱情", Rating: 9.6, Description: "不疯魔不成活，程蝶衣的一生就是一出戏。", PosterColor: "#2A0D0D", PosterIcon: "drama", Director: "陈凯歌", CreatedAt: now, UpdatedAt: now},
		{Title: "头号玩家", Year: 2018, Genres: "科幻,冒险", Rating: 8.7, Description: "斯皮尔伯格的彩蛋狂欢，致敬流行文化的冒险之旅。", PosterColor: "#0D1A1A", PosterIcon: "gamepad-2", Director: "史蒂文·斯皮尔伯格", CreatedAt: now, UpdatedAt: now},
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
