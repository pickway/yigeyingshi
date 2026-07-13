<template>
  <div class="movie-recommend-page">
    <!-- Page Title -->
    <div class="page-header">
      <div class="container">
        <h1 class="cine-display-lg">影视推荐</h1>
        <p class="page-subtitle">精选佳片，从经典到新锐，发现属于你的光影故事</p>
      </div>
    </div>

    <!-- Filter Bar -->
    <div class="container filter-bar-wrapper">
      <div class="filter-bar">
        <div class="genre-pills">
          <button
            v-for="genre in genres"
            :key="genre"
            class="genre-pill"
            :class="{ active: activeGenre === genre }"
            @click="activeGenre = genre"
          >
            {{ genre }}
          </button>
        </div>
        <div class="sort-select-wrapper">
          <select v-model="sort" class="sort-select">
            <option v-for="opt in sortOptions" :key="opt.value" :value="opt.value">
              {{ opt.label }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <!-- Movie Poster Grid -->
    <div class="container movie-grid-section">
      <div v-if="loading" class="loading-state">加载中...</div>
      <div v-else-if="allMovies.length === 0" class="loading-state">暂无数据</div>
      <div v-else class="movie-grid">
        <div
          v-for="movie in allMovies"
          :key="movie.id"
          class="movie-card"
        >
          <div class="poster-area" :style="{ background: movie.gradient || posterGradient(getMovieColor(movie)) }">
            <component :is="getIcon(movie)" class="poster-icon" />
          </div>
          <div class="movie-info">
            <h3 class="movie-title">{{ movie.title }}</h3>
            <span class="movie-year">{{ movie.year }}</span>
            <div class="movie-genres">
              <span v-for="g in splitGenres(movie.genres)" :key="g" class="genre-tag">{{ g }}</span>
            </div>
            <div class="movie-rating">
              <Star class="star-icon" :size="14" />
              <span class="rating-value">{{ movie.rating }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Editor Pick Section -->
    <div v-if="editorPick" class="container editor-pick-section">
      <div class="editor-pick-container">
        <h2 class="cine-heading editor-pick-title">编辑推荐</h2>
        <div class="editor-pick-content">
          <div class="editor-pick-poster" :style="{ background: editorPick.gradient || posterGradient('#5A4A6B') }">
            <Sparkles class="poster-icon" />
          </div>
          <div class="editor-pick-info">
            <div class="editor-pick-meta">
              <span class="editor-pick-year">{{ editorPick.year }}</span>
              <span class="meta-divider">/</span>
              <span v-for="(g, i) in splitGenres(editorPick.genres)" :key="g" class="editor-pick-genre">
                {{ g }}<span v-if="i < splitGenres(editorPick.genres).length - 1" class="meta-divider"> /</span>
              </span>
            </div>
            <h3 class="editor-pick-name cine-subheading">{{ editorPick.title }}</h3>
            <div class="editor-pick-rating">
              <Star class="star-icon" :size="16" />
              <span class="rating-value">{{ editorPick.rating }}</span>
            </div>
            <p class="editor-pick-desc">
              {{ editorPick.description || '诺兰以独特的叙事结构，将"原子弹之父"奥本海默的生平编织成一部震撼人心的传记史诗。影片在科学与道德的交汇处，探寻人类文明的终极困境。' }}
            </p>
            <button class="cta-button">
              <Play :size="16" />
              <span>立即观看</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { movieApi } from '@/api'
import {
  Star, Play, Orbit, Sparkles, Home, Layers, Trees, Search,
  Flame, Bird, Rabbit, Zap
} from '@lucide/vue'

const allMovies = ref([])
const loading = ref(true)
const activeGenre = ref('全部')
const sort = ref('rating_desc')

const genres = ['全部', '科幻', '动画', '剧情', '悬疑', '喜剧', '传记', '历史', '纪录片']

const sortOptions = [
  { value: 'rating_desc', label: '评分优先' },
  { value: 'year_desc', label: '年份最新' },
  { value: 'year_asc', label: '年份最早' },
]

const iconMap = { orbit, sparkles, home, layers, trees, search, flame, bird, rabbit, zap }
const iconKeys = Object.keys(iconMap)
const colorPalette = [
  '#2C3E6B', '#6B4E8B', '#4A6741', '#3A5A8C', '#6B8E5A',
  '#8B6F4E', '#E87D6B', '#5A6E82', '#D4853A', '#5A4A6B'
]

function getMovieColor(movie) {
  const idx = movie.id ? Number(movie.id) % colorPalette.length : 0
  return colorPalette[idx]
}

function getIcon(movie) {
  const idx = movie.id ? Number(movie.id) % iconKeys.length : 0
  return iconMap[iconKeys[idx]] || Sparkles
}

function splitGenres(genres) {
  if (!genres) return []
  return String(genres).split(',').map(g => g.trim()).filter(Boolean)
}

function posterGradient(color) {
  return `linear-gradient(135deg, ${color} 0%, ${color}aa 60%, ${color}55 100%)`
}

function primaryGenre(genres) {
  if (!genres) return ''
  return String(genres).split(',')[0]
}

async function loadMovies() {
  loading.value = true
  try {
    const params = { pageSize: 50 }
    if (activeGenre.value !== '全部') params.genre = activeGenre.value
    params.sort = sort.value
    const res = await movieApi.list(params)
    allMovies.value = (res.data || []).map(m => ({
      ...m,
      gradient: posterGradient(getMovieColor(m)),
    }))
  } catch (e) {
    console.error('加载电影失败', e)
    allMovies.value = []
  } finally {
    loading.value = false
  }
}

onMounted(loadMovies)
watch([activeGenre, sort], loadMovies)

const editorPick = computed(() => {
  if (!allMovies.value.length) return null
  return allMovies.value.find(m => m.title === '奥本海默') || allMovies.value[0]
})
</script>

<style scoped>
:root {
  --color-primary: #E85D3A;
  --color-primary-light: #F2845F;
  --color-primary-dark: #C44A2B;
  --color-primary-tint-1: rgba(232, 93, 58, 0.10);
  --color-bg-base: #FAFAF8;
  --color-bg-elevated: #FFFFFF;
  --color-bg-surface: #F3F2EF;
  --color-bg-muted: #EAEAE6;
  --color-border-subtle: #EDECE8;
  --color-border-default: #E0DFDB;
  --color-text-primary: #1A1A1A;
  --color-text-secondary: #6B6B6B;
  --color-text-tertiary: #9A9A96;
  --color-text-inverse: #FFFFFF;
  --shadow-sm: 0 1px 2px rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.07), 0 2px 4px -2px rgba(0, 0, 0, 0.05);
  --radius-sm: 6px;
  --radius-md: 8px;
  --radius-lg: 12px;
  --radius-full: 9999px;
  --container-max: 1200px;
  --content-padding: 24px;
  --font-display: 'Georgia', 'Noto Serif SC', serif;
  --font-body: 'Inter', 'Noto Sans SC', sans-serif;
  --transition-fast: 150ms ease;
  --transition-base: 250ms ease;
}

.movie-recommend-page {
  background-color: var(--color-bg-base);
  font-family: var(--font-body);
  color: var(--color-text-primary);
  min-height: 100vh;
}

.container {
  max-width: var(--container-max);
  margin-inline: auto;
  padding-inline: 24px;
}

/* Page Header */
.page-header {
  padding-top: 40px;
  padding-bottom: 16px;
}

.cine-display-lg {
  font-family: var(--font-display);
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--color-text-primary);
  letter-spacing: -0.02em;
  line-height: 1.2;
}

.page-subtitle {
  margin-top: 8px;
  font-size: 1rem;
  color: var(--color-text-secondary);
  line-height: 1.6;
}

/* Filter Bar */
.filter-bar-wrapper {
  margin-top: 8px;
}

.filter-bar {
  background: var(--color-bg-elevated);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
  border-bottom: 2px solid var(--color-primary);
  padding: 16px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 12px;
}

.genre-pills {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.genre-pill {
  padding: 6px 16px;
  border-radius: var(--radius-full);
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  border: 1px solid transparent;
  transition: all var(--transition-fast);
  background: var(--color-bg-muted);
  color: var(--color-text-secondary);
  border-color: var(--color-border-default);
}

.genre-pill.active {
  background: var(--color-primary);
  color: var(--color-text-inverse);
  border-color: var(--color-primary);
}

.genre-pill:not(.active):hover {
  background: var(--color-primary-tint-1);
  color: var(--color-primary);
  border-color: var(--color-primary-light);
}

.sort-select-wrapper {
  flex-shrink: 0;
}

.sort-select {
  appearance: none;
  background: var(--color-bg-muted);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border-default);
  border-radius: var(--radius-md);
  padding: 6px 36px 6px 12px;
  font-size: 0.875rem;
  font-family: var(--font-body);
  cursor: pointer;
  transition: all var(--transition-fast);
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 24 24' fill='none' stroke='%236B6B6B' stroke-width='2'%3E%3Cpath d='m6 9 6 6 6-6'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 12px center;
}

.sort-select:focus {
  outline: 2px solid var(--color-primary-light);
  outline-offset: 2px;
}

/* Movie Grid */
.movie-grid-section {
  padding-top: 40px;
  padding-bottom: 80px;
}

.movie-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 24px;
}

@media (min-width: 640px) {
  .movie-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .movie-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

/* Movie Card */
.movie-card {
  border-radius: var(--radius-lg);
  overflow: hidden;
  background: var(--color-bg-elevated);
  box-shadow: var(--shadow-sm);
  transition: transform var(--transition-base), box-shadow var(--transition-base);
  cursor: pointer;
}

.movie-card:hover {
  transform: translateY(-4px);
  box-shadow:
    0 8px 16px -4px rgba(232, 93, 58, 0.15),
    0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.poster-area {
  aspect-ratio: 2 / 3;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.poster-icon {
  width: 64px;
  height: 64px;
  opacity: 0.15;
  color: #FFFFFF;
}

.movie-info {
  padding: 12px 16px 16px;
}

.movie-title {
  font-family: var(--font-display);
  font-size: 1.05rem;
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0;
  line-height: 1.3;
}

.movie-year {
  font-size: 0.8rem;
  color: var(--color-text-tertiary);
  margin-top: 2px;
  display: inline-block;
}

.movie-genres {
  display: flex;
  gap: 6px;
  margin-top: 6px;
  flex-wrap: wrap;
}

.genre-tag {
  font-size: 0.7rem;
  padding: 2px 8px;
  border-radius: var(--radius-full);
  background: var(--color-primary-tint-1);
  color: var(--color-primary);
  font-weight: 500;
}

.movie-rating {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-top: 8px;
}

.star-icon {
  color: var(--color-primary);
  fill: var(--color-primary);
}

.rating-value {
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

/* Editor Pick Section */
.editor-pick-section {
  padding-bottom: 80px;
}

.editor-pick-container {
  background: var(--color-bg-surface);
  border-radius: var(--radius-lg);
  padding: 32px;
}

.cine-heading {
  font-family: var(--font-display);
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--color-text-primary);
  letter-spacing: -0.01em;
  line-height: 1.3;
}

.cine-subheading {
  font-family: var(--font-display);
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text-primary);
  line-height: 1.4;
}

.editor-pick-title {
  margin-bottom: 24px;
}

.editor-pick-content {
  display: flex;
  gap: 32px;
  align-items: flex-start;
}

.editor-pick-poster {
  max-width: 200px;
  width: 100%;
  aspect-ratio: 2 / 3;
  border-radius: var(--radius-md);
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.editor-pick-poster .poster-icon {
  width: 56px;
  height: 56px;
  opacity: 0.15;
  color: #FFFFFF;
}

.editor-pick-info {
  flex: 1;
  padding-left: 24px;
  border-left: 3px solid var(--color-primary);
}

.editor-pick-meta {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.85rem;
  color: var(--color-text-secondary);
  margin-bottom: 8px;
}

.meta-divider {
  color: var(--color-text-tertiary);
}

.editor-pick-name {
  margin: 0 0 8px;
}

.editor-pick-rating {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-bottom: 16px;
}

.editor-pick-desc {
  font-size: 0.95rem;
  color: var(--color-text-secondary);
  line-height: 1.7;
  margin: 0 0 24px;
}

.cta-button {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 24px;
  background: var(--color-primary);
  color: var(--color-text-inverse);
  border: none;
  border-radius: var(--radius-full);
  font-size: 0.9rem;
  font-weight: 600;
  font-family: var(--font-body);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.cta-button:hover {
  background: var(--color-primary-dark);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(232, 93, 58, 0.3);
}

/* Loading State */
.loading-state {
  text-align: center;
  padding: 60px 0;
  color: var(--color-text-secondary);
  font-size: 0.95rem;
}

@media (max-width: 640px) {
  .editor-pick-content {
    flex-direction: column;
  }

  .editor-pick-poster {
    max-width: 160px;
  }

  .editor-pick-info {
    padding-left: 0;
    padding-top: 16px;
    border-left: none;
    border-top: 3px solid var(--color-primary);
  }

  .filter-bar {
    flex-direction: column;
    align-items: flex-start;
  }

  .cine-display-lg {
    font-size: 1.75rem;
  }
}
</style>
