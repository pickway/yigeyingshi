<template>
  <main class="min-h-screen" style="background:var(--color-bg-base);">
    <div class="max-w-[var(--container-max)] mx-auto" style="padding-left:var(--content-padding); padding-right:var(--content-padding);">

      <!-- Page Header + Filter Bar -->
      <section class="pt-8 pb-6 border-b" style="border-color:var(--color-border-subtle);">
        <span class="eyebrow">Curated Filmography</span>
        <h1 class="cinema-heading mb-3" style="font-size:clamp(2.5rem,6vw,4.8rem); text-wrap:balance; word-break:keep-all; overflow-wrap:break-word;">私人片单</h1>
        <p class="mb-8" style="max-width:620px;color:var(--color-text-tertiary);line-height:1.7;">在类型、年代和作者之间漫游。每一部都来自真实观看，而不是算法自动生成的榜单。</p>

        <!-- Filter + Search Row -->
        <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <!-- Genre Pills -->
          <div class="flex flex-nowrap overflow-x-auto gap-2 no-scrollbar">
            <button
              v-for="genre in genreList"
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
                @keydown.enter="submitSearch"
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
        <UiState v-if="loading" type="loading" title="正在整理片库…" />
        <UiState v-else-if="error" type="error" :message="error" @retry="loadMovies" />
        <UiState v-else-if="allMovies.length === 0" type="empty" title="没有找到相关影片" message="换一个关键词或类型试试。" />
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <router-link
            v-for="movie in allMovies"
            :key="movie.id"
            :to="`/movies/${movie.id}`"
            class="group rounded-lg border transition-all duration-250 movie-card"
            style="background:var(--color-bg-surface); border-color:var(--color-border-default); overflow:hidden; color:inherit;text-decoration:none;"
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
              <div class="flex items-center gap-1 mb-2">
                <Star class="star-icon" :size="14" />
                <span class="text-sm font-medium" style="color:var(--color-primary);">{{ movie.rating }}</span>
              </div>
              <p class="text-xs line-clamp-2" style="color:var(--color-text-secondary); line-height:var(--leading-normal);">
                {{ movie.description }}
              </p>
            </div>
          </router-link>
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
            <!-- Jump to page -->
            <div class="flex items-center gap-2" style="padding-right:12px; border-right:1px solid var(--color-border-subtle);">
              <span class="cinema-body-sm" style="color:var(--color-text-tertiary); white-space:nowrap;">跳至</span>
              <input
                v-model.number="jumpPage"
                type="number"
                min="1"
                :max="totalPages"
                class="text-center outline-none transition-colors duration-150"
                style="width:40px; height:28px; border-radius:var(--radius-sm); border:1px solid var(--color-border-default); background:var(--color-bg-base); color:var(--color-text-primary); font-family:var(--font-body); font-size:var(--text-sm); -moz-appearance:textfield;"
                @focus="$event.target.style.borderColor='var(--color-primary)'"
                @blur="$event.target.style.borderColor='var(--color-border-default)'"
                @keydown.enter="doJump"
              />
              <span class="cinema-body-sm" style="color:var(--color-text-tertiary); white-space:nowrap;">页</span>
              <button
                class="jump-btn"
                aria-label="跳转"
                @click="doJump"
              >
                <ArrowRight :size="14" />
              </button>
            </div>

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
                v-for="p in visiblePages"
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
              <div class="flex items-center gap-3 mb-4">
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
                <span class="text-xs" style="color:var(--color-text-tertiary);">导演</span>
                <span class="text-sm" style="color:var(--color-text-secondary);">{{ editorPick.director }}</span>
              </div>
              <div class="flex items-center gap-1 mb-4">
                <Star class="star-icon" :size="14" />
                <span class="text-sm font-medium" style="color:var(--color-primary);">{{ editorPick.rating }}</span>
              </div>
              <p class="text-sm line-clamp-4 mb-6" style="color:var(--color-text-secondary); line-height:var(--leading-normal);">
                {{ editorPick.description || '诺兰以独特的叙事结构，将"原子弹之父"奥本海默的生平编织成一部震撼人心的传记史诗。影片在科学与道德的交汇处，探寻人类文明的终极困境。' }}
              </p>
              <router-link
                :to="`/movies/${editorPick.id}`"
                class="inline-flex items-center justify-center px-6 py-2.5 rounded-lg text-sm font-medium whitespace-nowrap transition-all duration-150 no-underline cta-link"
                style="background:var(--color-primary); color:var(--color-text-inverse); font-family:var(--font-body);"
              >
                查看影片笔记
              </router-link>
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
import UiState from '@/components/UiState.vue'
import { buildMovieQuery, createPageRange } from '@/utils/content'
import {
  Star, Search, ChevronDown, ChevronLeft, ChevronRight, ArrowRight
} from '@lucide/vue'

const allMovies = ref([])
const loading = ref(true)
const error = ref('')
const activeGenre = ref('全部')
const sort = ref('rating_desc')
const page = ref(1)
const pageSize = ref(6)
const total = ref(0)
const jumpPage = ref(1)

const searchQuery = ref('')

// Hard-coded genre list aligned with design
const genreList = ['全部', '科幻', '动画', '剧情', '悬疑', '喜剧', '纪录片']

const sortOptions = [
  { value: 'rating_desc', label: '评分最高' },
  { value: 'year_desc', label: '最近上映' },
  { value: 'popular', label: '最受欢迎' },
]

const colorPalette = [
  '#1a1a2e', '#16213e', '#0f3460', '#2d1b2e', '#3b1f3b',
  '#5c2d5c', '#1a2a1a', '#2b3d2b', '#3a5a3a', '#0d1b2a'
]

function getMovieColor(movie) {
  const idx = movie.id ? Number(movie.id) % colorPalette.length : 0
  return colorPalette[idx]
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
  error.value = ''
  try {
    const params = { ...buildMovieQuery({ keyword: searchQuery.value, genre: activeGenre.value, sort: sort.value }), page: page.value, pageSize: pageSize.value }
    const res = await movieApi.list(params)
    allMovies.value = (res.data || []).map(m => ({
      ...m,
      gradient: posterGradient(getMovieColor(m)),
    }))
    total.value = res.total || 0
  } catch (e) {
    error.value = e.message
    allMovies.value = []
    total.value = 0
  } finally {
    loading.value = false
  }
}

function submitSearch() {
  page.value = 1
  jumpPage.value = 1
  loadMovies()
}

const totalPages = computed(() => Math.ceil(total.value / pageSize.value) || 1)
const visiblePages = computed(() => createPageRange(totalPages.value, page.value))

function goToPage(p) {
  if (p < 1 || p > totalPages.value || p === page.value) return
  page.value = p
  jumpPage.value = p
}

function doJump() {
  const p = parseInt(jumpPage.value)
  if (p >= 1 && p <= totalPages.value) {
    goToPage(p)
  }
}

watch([activeGenre, sort], () => {
  page.value = 1
  jumpPage.value = 1
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
  --color-primary: #D4A853;
  --color-primary-light: #E8C97A;
  --color-primary-dark: #B08930;
  --color-primary-tint-1: rgba(212,168,83,0.12);
  --color-primary-tint-2: rgba(212,168,83,0.06);
  --color-primary-tint-3: rgba(212,168,83,0.03);

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

  --radius-sm: 4px;
  --radius-md: 8px;
  --radius-lg: 12px;
  --radius-full: 9999px;

  --shadow-sm: 0 1px 3px rgba(0,0,0,0.3);
  --shadow-md: 0 4px 12px rgba(0,0,0,0.3);
  --shadow-lg: 0 8px 24px rgba(0,0,0,0.4);
  --shadow-glow: 0 0 20px rgba(212,168,83,0.15);

  --transition-fast: 150ms ease;
  --transition-base: 250ms ease;
  --transition-slow: 400ms ease;

  --container-max: 1200px;
  --content-padding: 24px;

  --text-xs: 0.75rem;
  --text-sm: 0.8125rem;
  --text-base: 0.9375rem;
  --text-lg: 1.0625rem;
  --text-xl: 1.25rem;
  --text-2xl: 1.5rem;
  --text-3xl: 1.875rem;
  --text-4xl: 2.25rem;

  --leading-tight: 1.25;
  --leading-snug: 1.35;
  --leading-normal: 1.6;
  --leading-relaxed: 1.75;

  --tracking-tight: -0.02em;
  --tracking-normal: 0;
  --tracking-wide: 0.04em;

  --font-display: Georgia, 'Songti SC', 'STSong', serif;
  --font-body: -apple-system, 'PingFang SC', 'Microsoft YaHei', system-ui, sans-serif;
}

.cinema-heading {
  font-family: var(--font-display);
  font-weight: 600;
  font-size: var(--text-2xl);
  letter-spacing: var(--tracking-tight);
  line-height: var(--leading-snug);
}

.cinema-body-sm {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  line-height: var(--leading-normal);
  color: var(--color-text-secondary);
}

.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

/* Remove number input arrows */
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
input[type="number"] {
  -moz-appearance: textfield;
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
  gap: 8px;
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
  width: 6px;
  height: 6px;
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
  width: 14px;
  height: 6px;
  border-radius: var(--radius-full);
  background: var(--color-primary);
}

/* Jump Button */
.jump-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: var(--radius-sm);
  border: 1px solid var(--color-border-default);
  background: var(--color-bg-elevated);
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: all var(--transition-fast);
}
.jump-btn:hover {
  border-color: var(--color-primary);
  color: var(--color-primary);
}

/* CTA Link */
.cta-link {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: all var(--transition-fast);
}
.cta-link:hover {
  background: var(--color-primary-light) !important;
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
