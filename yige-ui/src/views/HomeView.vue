<template>
  <div class="home-page">
    <!-- Hero Section -->
    <section class="hero-section">
      <div class="container hero-inner">
        <div class="hero-text">
          <h1 class="cinema-display hero-title">光影之间</h1>
          <p class="cinema-subheading hero-subtitle">探索影视世界，分享学习心得，洞见AI未来</p>
          <div class="hero-actions">
            <router-link to="/movie-recommend" class="btn-primary">
              探索推荐
              <ArrowRight :size="18" class="btn-icon" />
            </router-link>
            <button class="btn-outline">了解更多</button>
          </div>
        </div>
        <div class="hero-image">
          <img src="/hero-bright.jpg" alt="光影之间" class="hero-img" />
        </div>
      </div>
    </section>

    <!-- Featured Movies Strip -->
    <section class="featured-section">
      <div class="container">
        <div class="section-header">
          <h2 class="cinema-heading section-title">本周推荐</h2>
          <span class="title-accent"></span>
        </div>
        <div v-if="loading" class="cinema-body-sm loading-text">加载中...</div>
        <div v-else class="movies-grid">
          <div
            v-for="movie in featuredMovies"
            :key="movie.id"
            class="movie-card group"
          >
            <div class="movie-poster" :style="{ background: posterGradient(movie.posterColor) }">
              <span class="genre-pill">{{ primaryGenre(movie.genres) }}</span>
            </div>
            <div class="movie-info">
              <h3 class="cinema-subheading movie-title">{{ movie.title }}</h3>
              <div class="movie-rating">
                <Star :size="14" class="star-icon" />
                <span class="cinema-body-sm">{{ movie.rating }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Latest Posts Section -->
    <section class="posts-section">
      <div class="container">
        <div class="section-header">
          <h2 class="cinema-heading section-title">最新动态</h2>
          <span class="title-accent"></span>
        </div>
        <div class="posts-layout">
          <div class="posts-articles">
            <div v-if="loading" class="cinema-body-sm loading-text">加载中...</div>
            <div
              v-else
              v-for="post in latestPosts"
              :key="post.id"
              class="article-card group"
            >
              <div class="article-meta">
                <span class="cinema-caption article-date">{{ formatDate(post.publishedAt) }}</span>
                <span class="article-tag">{{ post.tag }}</span>
              </div>
              <h3 class="cinema-subheading article-title">{{ post.title }}</h3>
              <p class="cinema-body-sm article-desc">{{ post.summary }}</p>
            </div>
          </div>
          <div class="posts-sidebar">
            <div class="learning-card">
              <div class="learning-top-accent"></div>
              <div class="learning-content">
                <GraduationCap :size="24" class="learning-icon" />
                <h3 class="cinema-subheading">从零开始的剪辑课</h3>
                <div class="progress-bar-track">
                  <div class="progress-bar-fill" style="width: 35%"></div>
                </div>
                <span class="cinema-caption progress-label">已完成 35%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- AI Highlights -->
    <section class="ai-section">
      <div class="container">
        <div class="section-header-row">
          <div class="section-header-left">
            <h2 class="cinema-heading section-title">AI 前沿</h2>
            <span class="title-accent"></span>
          </div>
          <router-link to="/ai-share" class="view-more-link">
            查看更多
            <ArrowRight :size="16" />
          </router-link>
        </div>
        <div class="ai-grid">
          <div
            v-for="item in aiHighlights"
            :key="item.id"
            class="ai-card group"
          >
            <div class="ai-icon-wrapper">
              <component :is="iconMap[item.icon] || Video" :size="24" class="ai-icon" />
            </div>
            <h3 class="cinema-subheading ai-card-title">{{ item.name }}</h3>
            <p class="cinema-body-sm ai-card-desc">{{ item.description }}</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { movieApi, articleApi, aiApi } from '@/api'
import { ArrowRight, Star, GraduationCap, Video, Subtitles, Brain, Image } from '@lucide/vue'

const featuredMovies = ref([])
const latestPosts = ref([])
const aiHighlights = ref([])
const loading = ref(true)

function posterGradient(color) {
  // posterColor from API is single hex like "#2C3E6B"
  // Build a gradient from that color
  return `linear-gradient(135deg, ${color} 0%, ${color}aa 50%, ${color}55 100%)`
}

function formatDate(iso) {
  if (!iso) return ''
  const d = new Date(iso)
  return `${d.getFullYear()}年${d.getMonth()+1}月${d.getDate()}日`
}

function primaryGenre(genres) {
  if (!genres) return ''
  return genres.split(',')[0]
}

const iconMap = {
  video: Video,
  image: Image,
  'audio-waveform': Brain,
}

onMounted(async () => {
  try {
    const [movies, articles, aiTools] = await Promise.all([
      movieApi.featured(),
      articleApi.latest(2),
      aiApi.featuredTools(),
    ])
    featuredMovies.value = movies.data || []
    latestPosts.value = articles.data || []
    aiHighlights.value = (aiTools.data || []).slice(0, 3)
  } catch (e) {
    console.error('加载首页数据失败', e)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.home-page {
  background-color: var(--color-bg-base);
  color: var(--color-text-primary);
}

.container {
  max-width: var(--container-max);
  margin-inline: auto;
  padding-inline: var(--content-padding);
}

/* ===== Hero Section ===== */
.hero-section {
  padding: 80px 0;
  background-color: var(--color-bg-base);
}

.hero-inner {
  display: flex;
  align-items: center;
  gap: 48px;
}

.hero-text {
  flex: 3;
  min-width: 0;
}

.hero-title {
  font-size: clamp(2.5rem, 5vw, 4rem);
  line-height: 1.15;
  margin-bottom: 16px;
  color: var(--color-text-primary);
}

.hero-subtitle {
  margin-bottom: 32px;
  color: var(--color-text-secondary);
}

.hero-actions {
  display: flex;
  gap: 16px;
  align-items: center;
}

.btn-primary {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 28px;
  background-color: var(--color-primary);
  color: var(--color-text-inverse);
  border-radius: var(--radius-md);
  font-weight: 600;
  text-decoration: none;
  transition: background-color var(--transition-fast), transform var(--transition-fast);
}

.btn-primary:hover {
  background-color: var(--color-primary-dark);
  transform: translateY(-1px);
}

.btn-icon {
  transition: transform var(--transition-fast);
}

.btn-primary:hover .btn-icon {
  transform: translateX(2px);
}

.btn-outline {
  display: inline-flex;
  align-items: center;
  padding: 12px 28px;
  background: transparent;
  color: var(--color-primary);
  border: 1.5px solid var(--color-primary);
  border-radius: var(--radius-md);
  font-weight: 600;
  cursor: pointer;
  transition: background-color var(--transition-fast), color var(--transition-fast);
}

.btn-outline:hover {
  background-color: var(--color-primary-tint-1);
}

.hero-image {
  flex: 2;
  min-width: 0;
}

.hero-img {
  width: 100%;
  aspect-ratio: 4 / 3;
  object-fit: cover;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-md);
}

/* ===== Section Header ===== */
.section-header {
  margin-bottom: 32px;
}

.section-title {
  margin-bottom: 0;
}

.title-accent {
  display: block;
  width: 40px;
  height: 3px;
  background-color: var(--color-primary);
  border-radius: 2px;
  margin-top: 8px;
}

.loading-text {
  color: var(--color-text-tertiary);
  padding: 24px 0;
}

/* ===== Featured Movies ===== */
.featured-section {
  padding: 80px 0;
  background-color: var(--color-bg-surface);
}

.movies-grid {
  display: grid;
  grid-template-columns: repeat(1, 1fr);
  gap: 24px;
}

@media (min-width: 640px) {
  .movies-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 900px) {
  .movies-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

.movie-card {
  background-color: var(--color-bg-elevated);
  border-radius: var(--radius-lg);
  overflow: hidden;
  transition: transform var(--transition-base), box-shadow var(--transition-base);
  cursor: pointer;
}

.movie-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.movie-poster {
  aspect-ratio: 3 / 4;
  position: relative;
  padding: 12px;
}

.genre-pill {
  position: absolute;
  top: 12px;
  left: 12px;
  padding: 4px 12px;
  background-color: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(8px);
  color: var(--color-text-inverse);
  border-radius: var(--radius-full);
  font-size: 12px;
  font-weight: 500;
}

.movie-info {
  padding: 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.movie-title {
  color: var(--color-text-primary);
}

.movie-rating {
  display: flex;
  align-items: center;
  gap: 4px;
  color: var(--color-primary);
}

.star-icon {
  fill: var(--color-primary);
  color: var(--color-primary);
}

/* ===== Latest Posts ===== */
.posts-section {
  padding: 80px 0;
  background-color: var(--color-bg-base);
}

.posts-layout {
  display: flex;
  gap: 32px;
}

.posts-articles {
  flex: 3;
  display: flex;
  flex-direction: column;
  gap: 24px;
  min-width: 0;
}

.posts-sidebar {
  flex: 2;
  min-width: 0;
}

.article-card {
  background-color: var(--color-bg-elevated);
  border-left: 3px solid var(--color-primary);
  border-radius: var(--radius-md);
  padding: 24px;
  transition: transform var(--transition-base), box-shadow var(--transition-base);
  cursor: pointer;
}

.article-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.article-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
}

.article-date {
  color: var(--color-text-tertiary);
}

.article-tag {
  padding: 2px 10px;
  background-color: var(--color-primary-tint-1);
  color: var(--color-primary);
  border-radius: var(--radius-full);
  font-size: 12px;
  font-weight: 500;
}

.article-title {
  color: var(--color-text-primary);
  margin-bottom: 8px;
}

.article-desc {
  color: var(--color-text-secondary);
  line-height: 1.6;
}

/* Learning Card */
.learning-card {
  background-color: var(--color-bg-elevated);
  border-radius: var(--radius-lg);
  overflow: hidden;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.learning-top-accent {
  height: 4px;
  background: linear-gradient(90deg, var(--color-primary), var(--color-primary-light));
}

.learning-content {
  padding: 24px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.learning-icon {
  color: var(--color-primary);
}

.progress-bar-track {
  height: 6px;
  background-color: var(--color-bg-surface);
  border-radius: var(--radius-full);
  overflow: hidden;
}

.progress-bar-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--color-primary), var(--color-primary-light));
  border-radius: var(--radius-full);
  transition: width var(--transition-base);
}

.progress-label {
  color: var(--color-text-tertiary);
}

/* ===== AI Highlights ===== */
.ai-section {
  padding: 80px 0;
  background-color: var(--color-bg-surface);
}

.section-header-row {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-bottom: 32px;
}

.section-header-left {
  display: flex;
  flex-direction: column;
}

.view-more-link {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  color: var(--color-text-link);
  font-weight: 500;
  text-decoration: none;
  transition: color var(--transition-fast);
}

.view-more-link:hover {
  color: var(--color-primary-dark);
}

.ai-grid {
  display: grid;
  grid-template-columns: repeat(1, 1fr);
  gap: 24px;
}

@media (min-width: 640px) {
  .ai-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 900px) {
  .ai-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

.ai-card {
  background-color: var(--color-bg-elevated);
  border-radius: var(--radius-lg);
  padding: 28px;
  transition: transform var(--transition-base), box-shadow var(--transition-base);
  cursor: pointer;
}

.ai-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.ai-icon-wrapper {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  background-color: var(--color-primary-tint-1);
  border-radius: var(--radius-md);
  margin-bottom: 16px;
}

.ai-icon {
  color: var(--color-primary);
}

.ai-card-title {
  color: var(--color-text-primary);
  margin-bottom: 8px;
}

.ai-card-desc {
  color: var(--color-text-secondary);
  line-height: 1.6;
}

/* ===== Responsive ===== */
@media (max-width: 768px) {
  .hero-inner {
    flex-direction: column;
    gap: 32px;
  }

  .hero-text,
  .hero-image {
    flex: none;
    width: 100%;
  }

  .posts-layout {
    flex-direction: column;
  }

  .posts-articles,
  .posts-sidebar {
    flex: none;
    width: 100%;
  }
}
</style>