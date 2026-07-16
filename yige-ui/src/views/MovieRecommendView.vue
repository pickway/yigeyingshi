<template>
  <main class="min-h-screen" style="background:var(--color-bg-base);">
    <div class="max-w-[var(--container-max)] mx-auto" style="padding:var(--content-padding);">

      <!-- Page Header -->
      <section class="pt-8 pb-6 border-b" style="border-color:var(--color-border-subtle);">
        <h1 class="cinema-heading" style="font-size:var(--text-3xl); text-wrap:balance; word-break:keep-all; overflow-wrap:break-word;">
          影视推荐
        </h1>
      </section>

      <!-- Filter Bar -->
      <section class="pt-8 pb-6">
        <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <!-- Genre Pills -->
          <div class="flex flex-nowrap overflow-x-auto gap-2 no-scrollbar">
            <button
              v-for="genre in genres"
              :key="genre"
              class="shrink-0 inline-flex items-center justify-center px-4 py-2 rounded-lg text-sm font-medium whitespace-nowrap transition-colors duration-150"
              :style="activeGenre === genre
                ? { background: 'var(--color-primary)', color: 'var(--color-text-inverse)' }
                : { background: 'transparent', color: 'var(--color-text-secondary)', border: '1px solid var(--color-border-default)' }"
              @click="activeGenre = genre"
            >
              {{ genre }}
            </button>
          </div>

          <!-- Search + Sort -->
          <div class="flex items-center gap-3 shrink-0">
            <div class="relative">
              <Search class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none" style="width:16px; height:16px; color:var(--color-text-tertiary);" />
              <input
                v-model="searchQuery"
                type="text"
                placeholder="搜索影片..."
                class="pl-9 pr-4 py-2 rounded-lg text-sm outline-none transition-colors duration-150"
                style="width:180px; height:36px; background:var(--color-bg-muted); border:1px solid var(--color-border-default); color:var(--color-text-primary); font-family:var(--font-body);"
                @focus="$event.target.style.borderColor='var(--color-primary)'"
                @blur="$event.target.style.borderColor='var(--color-border-default)'"
              />
            </div>
            <div class="relative">
              <select
                v-model="sort"
                class="appearance-none pl-4 pr-8 py-2 rounded-lg text-sm outline-none cursor-pointer transition-colors duration-150"
                style="height:36px; background:var(--color-bg-muted); border:1px solid var(--color-border-default); color:var(--color-text-primary); font-family:var(--font-body);"
              >
                <option v-for="opt in sortOptions" :key="opt.value" :value="opt.value">
                  {{ opt.label }}
                </option>
              </select>
              <ChevronDown class="absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none" style="width:14px; height:14px; color:var(--color-text-tertiary);" />
            </div>
          </div>
        </div>
      </section>

      <!-- Movie Grid -->
      <section class="pt-8 pb-12">
        <div v-if="loading" class="loading-state">加载中...</div>
        <div v-else-if="allMovies.length === 0" class="loading-state">暂无数据</div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <article
            v-for="movie in allMovies"
            :key="movie.id"
            class="group rounded-lg border transition-all duration-250 movie-card"
            style="background:var(--color-bg-surface); border-color:var(--color-border-default); overflow:hidden;"
          >
            <!-- Poster Area -->
            <div
              class="relative overflow-hidden"
              style="aspect-ratio:2/3;"
              :style="{ background: movie.gradient || posterGradient(getMovieColor(movie)) }"
            >
              <div
                class="absolute inset-0 transition-transform duration-500 group-hover:scale-105"
                style="background:radial-gradient(circle at 70% 30%, rgba(212,168,83,0.15) 0%, transparent 60%);"
              ></div>
              <component :is="getIcon(movie)" class="poster-icon" />
              <div
                class="absolute bottom-0 left-0 right-0 p-4"
                style="background:linear-gradient(to top, rgba(12,12,14,0.9) 0%, transparent 100%);"
              >
                <h3 class="cinema-heading truncate" style="font-size:var(--text-lg); color:var(--color-text-primary);">
                  {{ movie.title }}
                </h3>
              </div>
            </div>

            <!-- Info Section -->
            <div class="p-4">
              <div class="flex items-center gap-2 mb-2">
                <span class="text-xs whitespace-nowrap" style="color:var(--color-text-tertiary);">{{ movie.year }}</span>
                <span
                  v-for="g in splitGenres(movie.genres).slice(0, 2)"
                  :key="g"
                  class="inline-flex items-center justify-center px-2 py-0.5 rounded-lg text-xs whitespace-nowrap"
                  style="background:var(--color-primary-tint-1); color:var(--color-primary);"
                >
                  {{ g }}
                </span>
              </div>
              <div class="flex items-center gap-1">
                <Star class="star-icon" :size="14" />
                <span class="text-sm font-medium" style="color:var(--color-primary);">{{ movie.rating }}</span>
              </div>
            </div>
          </article>
        </div>

        <!-- Pagination -->
        <div
          v-if="!loading && totalPages > 1"
          class="flex items-center justify-between mt-8 pt-6 border-t"
          style="border-color:var(--color-border-subtle);"
        >
          <span class="cinema-body-sm" style="color:var(--color-text-tertiary);">
            第 {{ page }} 页 / 共 {{ totalPages }} 页，共 {{ total }} 部影片
          </span>
          <div class="flex items-center gap-3">
            <button
              class="page-btn"
              :disabled="page === 1"
              @click="goToPage(page - 1)"
            >
              <ChevronLeft :size="16" />
              上一页
            </button>
            <div class="flex gap-2 items-center">
              <button
                v-for="p in totalPages"
                :key="p"
                class="page-dot"
                :class="{ active: p === page }"
                @click="goToPage(p)"
              />
            </div>
            <button
              class="page-btn"
              :disabled="page === totalPages"
              @click="goToPage(page + 1)"
            >
              下一页
              <ChevronRight :size="16" />
            </button>
          </div>
        </div>
      </section>

      <!-- Editor Pick Section -->
      <section v-if="editorPick" class="pb-16">
        <div class="rounded-lg border" style="background:var(--color-bg-surface); border-color:var(--color-border-default); overflow:hidden;">
          <!-- Section heading -->
          <div class="px-6 pt-6 pb-4">
            <h2 class="cinema-heading" style="font-size:var(--text-2xl); text-wrap:balance; word-break:keep-all;">编辑推荐</h2>
          </div>

          <!-- Featured card -->
          <div class="flex flex-col md:flex-row">
            <!-- Poster -->
            <div
              class="relative w-full md:w-2/3 overflow-hidden"
              style="min-height:280px;"
              :style="{ background: editorPick.gradient || posterGradient('#5A4A6B') }"
            >
              <div class="absolute inset-0" style="background:radial-gradient(circle at 60% 40%, rgba(212,168,83,0.2) 0%, transparent 50%);"></div>
              <Sparkles class="poster-icon" style="width:80px; height:80px;" />
              <div
                class="absolute bottom-0 left-0 right-0 p-6 md:hidden"
                style="background:linear-gradient(to top, rgba(12,12,14,0.9) 0%, transparent 100%);"
              >
                <h3 class="cinema-heading" style="font-size:var(--text-2xl); color:var(--color-text-primary);">{{ editorPick.title }}</h3>
              </div>
            </div>

            <!-- Info -->
            <div
              class="w-full md:w-1/3 p-6 flex flex-col justify-center border-l"
              style="border-color:var(--color-primary); background:var(--color-bg-elevated);"
            >
              <div class="hidden md:block mb-3">
                <h3 class="cinema-heading" style="font-size:var(--text-2xl); color:var(--color-text-primary); text-wrap:balance; word-break:keep-all;">
                  {{ editorPick.title }}
                </h3>
              </div>
              <div class="flex items-center gap-3 mb-4 flex-wrap">
                <span class="text-xs whitespace-nowrap" style="color:var(--color-text-tertiary);">{{ editorPick.year }}</span>
                <span
                  v-for="g in splitGenres(editorPick.genres)"
                  :key="g"
                  class="inline-flex items-center justify-center px-2 py-0.5 rounded-lg text-xs whitespace-nowrap"
                  style="background:var(--color-primary-tint-1); color:var(--color-primary);"
                >
                  {{ g }}
                </span>
              </div>
              <div class="flex items-center gap-2 mb-4">
                <Star class="star-icon" :size="14" />
                <span class="text-sm font-medium" style="color:var(--color-primary);">{{ editorPick.rating }}</span>
              </div>
              <p class="text-sm line-clamp-4 mb-6" style="color:var(--color-text-secondary); line-height:var(--leading-normal);">
                {{ editorPick.description || '诺兰以独特的叙事结构，将"原子弹之父"奥本海默的生平编织成一部震撼人心的传记史诗。影片在科学与道德的交汇处，探寻人类文明的终极困境。' }}
              </p>
              <button class="cta-button">
                <Play :size="16" />
                <span>立即观看</span>
              </button>
            </div>
          </div>
        </div>
      </section>

    </div>
  </main>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { movieApi } from '@/api'
import {
  Star, Play, Orbit, Sparkles, Home, Layers, Trees, Search,
  Flame, Bird, Rabbit, Zap, ChevronLeft, ChevronRight, ChevronDown
} from '@lucide/vue'

const allMovies = ref([])
const loading = ref(true)
const activeGenre = ref('全部')
const sort = ref('rating_desc')
const page = ref(1)
const pageSize = ref(6)
const total = ref(0)

// Local state for search input (no backend integration yet)
const searchQuery = ref('')

const genres = computed(() => {
  const set = new Set()
  allMovies.value.forEach(m => {
    if (m.genres) {
      m.genres.split(',').forEach(g => {
        const trimmed = g.trim()
        if (trimmed) set.add(trimmed)
      })
    }
  })
  return ['全部', ...Array.from(set).sort()]
})

const sortOptions = [
  { value: 'rating_desc', label: '评分最高' },
  { value: 'year_desc', label: '最近上映' },
  { value: 'rating_asc', label: '评分最低' },
]

const iconMap = { Orbit, Sparkles, Home, Layers, Trees, Search, Flame, Bird, Rabbit, Zap }
const iconKeys = Object.keys(iconMap)
const colorPalette = [
  '#1a1a2e', '#16213e', '#0f3460', '#2d1b2e', '#3b1f3b',
  '#5c2d5c', '#1a2a1a', '#2b3d2b', '#3a5a3a', '#0d1b2a'
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
  return `linear-gradient(160deg, ${color} 0%, ${color}cc 40%, ${color}88 100%)`
}

async function loadMovies() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (activeGenre.value !== '全部') params.genre = activeGenre.value
    params.sort = sort.value
    const res = await movieApi.list(params)
    allMovies.value = (res.data || []).map(m => ({
      ...m,
      gradient: posterGradient(getMovieColor(m)),
    }))
    total.value = res.total || 0
  } catch (e) {
    console.error('加载电影失败', e)
    allMovies.value = []
    total.value = 0
  } finally {
    loading.value = false
  }
}

const totalPages = computed(() => Math.ceil(total.value / pageSize.value) || 1)

function goToPage(p) {
  if (p < 1 || p > totalPages.value || p === page.value) return
  page.value = p
}

watch([activeGenre, sort], () => {
  page.value = 1
  loadMovies()
})

watch(page, loadMovies)

onMounted(loadMovies)

const editorPick = computed(() => {
  if (!allMovies.value.length) return null
  return allMovies.value.find(m => m.title === '奥本海默') || allMovies.value[0]
})
</script>

<style scoped>
/* Cinema Personal Site - Dark Theater Theme */
:root {
  /* === Primary Color System === */
  --color-primary: #D4A853;
  --color-primary-light: #E8C97A;
  --color-primary-dark: #B08930;
  --color-primary-tint-1: rgba(212,168,83,0.12);
  --color-primary-tint-2: rgba(212,168,83,0.06);
  --color-primary-tint-3: rgba(212,168,83,0.03);

  /* === Neutral Scale (Dark Theater) === */
  --color-bg-base: #0C0C0E;
  --color-bg-elevated: #151518;
  --color-bg-surface: #1C1C20;
  --color-bg-muted: #242428;
  --color-bg-overlay: rgba(12,12,14,0.85);
  --color-border-default: #2A2A30;
  --color-border-subtle: #1F1F25;
  --color-border-strong: #3A3A42;

  --color-text-primary: #F0EDE6;
  --color-text-secondary: #9A9AA0;
  --color-text-tertiary: #6A6A72;
  --color-text-inverse: #0C0C0E;
  --color-text-link: #D4A853;

  /* === Shape System === */
  --radius-sm: 4px;
  --radius-md: 8px;
  --radius-lg: 12px;
  --radius-full: 9999px;

  /* === Typography === */
  --font-display: 'Playfair Display', Georgia, 'Noto Serif SC', serif;
  --font-body: 'Inter', -apple-system, 'PingFang SC', 'Microsoft YaHei', sans-serif;
  --font-mono: 'JetBrains Mono', 'Fira Code', Consolas, monospace;

  --text-xs: 0.75rem;
  --text-sm: 0.8125rem;
  --text-base: 0.9375rem;
  --text-lg: 1.0625rem;
  --text-xl: 1.25rem;
  --text-2xl: 1.5rem;
  --text-3xl: 1.875rem;
  --text-4xl: 2.25rem;
  --text-5xl: 3rem;

  --leading-tight: 1.25;
  --leading-snug: 1.35;
  --leading-normal: 1.6;
  --leading-relaxed: 1.75;

  --tracking-tight: -0.02em;
  --tracking-normal: 0;
  --tracking-wide: 0.04em;

  /* === Spacing Scale === */
  --space-1: 4px;
  --space-2: 8px;
  --space-3: 12px;
  --space-4: 16px;
  --space-5: 20px;
  --space-6: 24px;
  --space-8: 32px;
  --space-10: 40px;
  --space-12: 48px;
  --space-16: 64px;

  /* === Shadows === */
  --shadow-sm: 0 1px 3px rgba(0,0,0,0.3);
  --shadow-md: 0 4px 12px rgba(0,0,0,0.3);
  --shadow-lg: 0 8px 24px rgba(0,0,0,0.4);
  --shadow-glow: 0 0 20px rgba(212,168,83,0.15);

  /* === Transitions === */
  --transition-fast: 150ms ease;
  --transition-base: 250ms ease;
  --transition-slow: 400ms ease;

  /* === Container === */
  --container-max: 1200px;
  --content-padding: 24px;
}

/* === Typography Classes === */
.cinema-display {
  font-family: var(--font-display);
  font-weight: 700;
  letter-spacing: var(--tracking-tight);
  line-height: var(--leading-tight);
}
.cinema-display-lg {
  font-family: var(--font-display);
  font-weight: 700;
  font-size: var(--text-4xl);
  letter-spacing: var(--tracking-tight);
  line-height: var(--leading-tight);
}
.cinema-heading {
  font-family: var(--font-display);
  font-weight: 600;
  font-size: var(--text-2xl);
  letter-spacing: var(--tracking-tight);
  line-height: var(--leading-snug);
}
.cinema-subheading {
  font-family: var(--font-body);
  font-weight: 500;
  font-size: var(--text-lg);
  line-height: var(--leading-snug);
  color: var(--color-text-secondary);
}
.cinema-body {
  font-family: var(--font-body);
  font-size: var(--text-base);
  line-height: var(--leading-normal);
  color: var(--color-text-primary);
}
.cinema-body-sm {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  line-height: var(--leading-normal);
  color: var(--color-text-secondary);
}
.cinema-caption {
  font-family: var(--font-body);
  font-size: var(--text-xs);
  line-height: var(--leading-normal);
  color: var(--color-text-tertiary);
  letter-spacing: var(--tracking-wide);
  text-transform: uppercase;
}

/* === Component Styles === */
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

/* Movie Card Hover Effects */
.movie-card {
  cursor: pointer;
}
.movie-card:hover {
  border-color: var(--color-primary) !important;
  transform: translateY(-4px);
  box-shadow: var(--shadow-glow);
}

/* Poster Icon */
.poster-icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 64px;
  height: 64px;
  opacity: 0.12;
  color: #FFFFFF;
}

/* Star Icon */
.star-icon {
  color: var(--color-primary);
  fill: var(--color-primary);
}

/* Loading State */
.loading-state {
  text-align: center;
  padding: 60px 0;
  color: var(--color-text-secondary);
  font-size: var(--text-base);
  font-family: var(--font-body);
}

/* Page Button */
.page-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-2);
  padding: 8px 16px;
  border-radius: var(--radius-md);
  border: 1px solid var(--color-border-default);
  background: var(--color-bg-elevated);
  color: var(--color-text-secondary);
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 500;
  cursor: pointer;
  transition: all var(--transition-fast);
}
.page-btn:hover:not(:disabled) {
  border-color: var(--color-primary);
  color: var(--color-primary);
}
.page-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
  pointer-events: none;
}

/* Page Dot */
.page-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  border: none;
  background: var(--color-bg-muted);
  cursor: pointer;
  transition: all var(--transition-fast);
}
.page-dot:hover {
  background: var(--color-primary-tint-2);
}
.page-dot.active {
  width: 24px;
  border-radius: var(--radius-full);
  background: var(--color-primary);
}

/* CTA Button */
.cta-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-2);
  padding: 10px 24px;
  border-radius: var(--radius-lg);
  background: var(--color-primary);
  color: var(--color-text-inverse);
  border: none;
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 600;
  cursor: pointer;
  transition: all var(--transition-fast);
}
.cta-button:hover {
  background: var(--color-primary-light);
  transform: translateY(-1px);
}

/* Select Dropdown */
select {
  background-image: none;
}

/* Responsive Adjustments */
@media (max-width: 768px) {
  .page-btn {
    padding: 6px 12px;
    font-size: var(--text-xs);
  }
}

@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    transition-duration: 0.01ms !important;
    animation-duration: 0.01ms !important;
  }
}
</style>