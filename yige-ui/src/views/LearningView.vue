<template>
  <div class="learning-page">
    <!-- Compact Hero Section -->
    <section class="hero-section">
      <nav class="breadcrumb" aria-label="面包屑导航">
        <router-link to="/" class="breadcrumb-link">首页</router-link>
        <ChevronRight class="breadcrumb-sep" :size="14" />
        <span class="breadcrumb-current">学习资料</span>
      </nav>
      <h1 class="cinema-display-lg">学习资料</h1>
      <p class="hero-subtitle">影视制作、剪辑技巧、AI创作 — 系统化的学习路径</p>
    </section>

    <!-- Category Tabs -->
    <section class="tabs-section">
      <div class="tabs-container">
        <button
          v-for="tab in tabs"
          :key="tab"
          class="tab-btn"
          :class="{ active: activeTab === tab }"
          @click="activeTab = tab"
        >
          {{ tab }}
        </button>
      </div>
    </section>

    <!-- Two-Column Layout -->
    <section class="content-layout">
      <!-- Left: Resource Grid -->
      <div class="resource-grid-container">
        <UiState v-if="loading" type="loading" title="正在整理学习资料…" />
        <UiState v-else-if="error" type="error" :message="error" @retry="load" />
        <div v-else class="resource-grid">
          <article
            v-for="course in allCourses"
            :key="course.id"
            class="resource-card"
          >
            <div class="resource-header">
              <div class="resource-icon-wrapper">
                <component
                  :is="iconMap[course.icon] || Clapperboard"
                  class="resource-icon"
                  :size="20"
                />
              </div>
              <div class="resource-tags-top">
                <span class="tag tag-category">{{ course.category }}</span>
                <span class="tag tag-level">{{ course.level }}</span>
              </div>
            </div>
            <h3 class="cinema-heading resource-title">{{ course.title }}</h3>
            <p class="resource-desc">{{ course.description }}</p>
            <div class="resource-footer">
              <div class="resource-meta">
                <span class="meta-item">
                  <BookOpen :size="13" />
                  {{ course.lessons || course.duration }}{{ course.lessons ? ' 课时' : '' }}
                </span>
                <span class="meta-item">
                  <Clock :size="13" />
                  约 {{ Math.max(1, Math.round((course.lessons || 2) * 0.5)) }} 小时
                </span>
              </div>
              <button type="button" class="resource-link" @click="selectCourse(course)">
                开始学习
              </button>
            </div>
          </article>
        </div>
      </div>

      <!-- Right: Sidebar -->
      <aside class="sidebar">
        <!-- My Learning Card -->
        <div class="sidebar-card my-learning-card">
          <h3 class="cinema-heading sidebar-card-title">我的学习</h3>
          <div class="stats-row">
            <div class="stat-item">
              <span class="cinema-display stat-value">{{ pathCourses.length }}</span>
              <span class="stat-label">进行中</span>
            </div>
            <div class="stat-divider"></div>
            <div class="stat-item">
              <span class="cinema-display stat-value stat-value-secondary">{{ myLearning.completed }}</span>
              <span class="stat-label">已完成</span>
            </div>
          </div>
          <div class="current-course-section">
            <p class="current-course-label">当前课程</p>
            <p class="current-course-name">{{ myLearning.currentCourse }}</p>
            <div class="progress-wrapper">
              <div class="progress-bar">
                <div
                  class="progress-fill"
                  :style="{ width: myLearning.currentProgress + '%' }"
                />
              </div>
              <span class="progress-pct">{{ myLearning.currentProgress }}%</span>
            </div>
          </div>
        </div>

        <!-- Recommended Path Card -->
        <div class="sidebar-card path-card">
          <h3 class="cinema-heading sidebar-card-title">推荐路径</h3>
          <p class="path-name">{{ recommendedPath ? recommendedPath.name : '暂无推荐路径' }}</p>
          <div class="path-steps">
            <div
              v-for="(course, i) in pathCourses"
              :key="course.id"
              class="path-step"
            >
              <div
                class="step-indicator"
                :class="{
                  'step-completed': i < 2,
                  'step-in-progress': i === 2,
                  'step-locked': i > 2
                }"
              >
                <Check v-if="i < 2" :size="13" class="step-check-icon" />
                <span v-else-if="i === 2" class="step-number">{{ i + 1 }}</span>
                <span v-else class="step-number step-number-locked">{{ i + 1 }}</span>
              </div>
              <div class="step-content">
                <p class="step-text" :class="{ 'step-text-completed': i < 2, 'step-text-locked': i > 2 }">
                  {{ course.title }}
                </p>
                <p class="step-duration">约 {{ Math.max(1, Math.round((course.lessons || 2) * 0.5)) }} 小时</p>
              </div>
            </div>
          </div>
        </div>
      </aside>
    </section>

    <!-- Bottom CTA Strip -->
    <section class="bottom-cta">
      <div class="cta-content">
        <HelpCircle :size="20" class="cta-icon" />
        <p class="cta-text">
          找不到需要的<span class="cta-highlight">资源</span>？
        </p>
      </div>
      <router-link to="/ai-share" class="cta-btn">
        提交需求
      </router-link>
    </section>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { learningApi } from '@/api'
import UiState from '@/components/UiState.vue'
import {
  ChevronRight,
  Clapperboard,
  Palette,
  SlidersHorizontal,
  Sparkles,
  Volume2,
  Video,
  BookOpen,
  Clock,
  Check,
  HelpCircle,
} from '@lucide/vue'

const allCourses = ref([])
const paths = ref([])
const loading = ref(true)
const error = ref('')
const activeTab = ref('全部')

const tabs = computed(() => {
  const set = new Set()
  allCourses.value.forEach(c => {
    if (c.category) set.add(c.category)
  })
  return ['全部', ...Array.from(set).sort()]
})

const iconMap = {
  clapperboard: Clapperboard,
  palette: Palette,
  'sliders-horizontal': SlidersHorizontal,
  sparkles: Sparkles,
  'volume-2': Volume2,
  video: Video,
}

async function load() {
  loading.value = true
  error.value = ''
  try {
    const params = {}
    if (activeTab.value !== '全部') params.category = activeTab.value
    const [courses, pathsRes] = await Promise.all([
      learningApi.courses(params),
      learningApi.paths(),
    ])
    allCourses.value = courses.data || []
    paths.value = pathsRes.data || []
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

function selectCourse(course) {
  myLearning.value.currentCourse = course.title
  myLearning.value.currentProgress = 0
  document.querySelector('.my-learning-card')?.scrollIntoView({ behavior: 'smooth', block: 'center' })
}

onMounted(load)
watch(activeTab, load)

const recommendedPath = computed(() => paths.value[0] || null)

const pathCourses = computed(() => {
  if (!recommendedPath.value || !recommendedPath.value.courseIds) return []
  const ids = recommendedPath.value.courseIds.split(',').map(s => parseInt(s.trim()))
  return ids.map(id => allCourses.value.find(c => c.id === id)).filter(Boolean)
})

const myLearning = ref({
  completed: 12,
  totalHours: 28,
  currentCourse: 'Premiere Pro 快速入门',
  currentProgress: 65,
})
</script>

<style scoped>
.learning-page {
  min-height: 100vh;
  background: var(--color-bg-base);
  padding-bottom: 64px;
}

/* === Hero Section === */
.hero-section {
  max-width: var(--container-max);
  margin: 0 auto;
  padding: var(--space-12) var(--content-padding) var(--space-8);
  border-top: 2px solid var(--color-primary);
  background: linear-gradient(180deg, var(--color-bg-elevated) 0%, var(--color-bg-base) 100%);
}

.breadcrumb {
  display: flex;
  align-items: center;
  gap: var(--space-2);
  margin-bottom: var(--space-6);
}

.breadcrumb-link {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  color: var(--color-text-tertiary);
  text-decoration: none;
  transition: color var(--transition-fast);
}

.breadcrumb-link:hover {
  color: var(--color-primary);
}

.breadcrumb-sep {
  color: var(--color-text-tertiary);
}

.breadcrumb-current {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  color: var(--color-text-secondary);
}

.hero-title {
  color: var(--color-text-primary);
  text-wrap: balance;
  word-break: keep-all;
  overflow-wrap: break-word;
}

.hero-subtitle {
  font-family: var(--font-body);
  font-size: var(--text-lg);
  color: var(--color-text-secondary);
  line-height: var(--leading-relaxed);
  max-width: 600px;
  margin-top: var(--space-3);
  margin-bottom: 0;
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
  color: var(--color-text-primary);
}

.cinema-heading {
  font-family: var(--font-display);
  font-weight: 600;
  font-size: var(--text-xl);
  letter-spacing: var(--tracking-tight);
  line-height: var(--leading-snug);
}

/* === Tabs Section === */
.tabs-section {
  max-width: var(--container-max);
  margin: 0 auto;
  padding: var(--space-6) var(--content-padding) var(--space-8);
}

.tabs-container {
  display: flex;
  align-items: center;
  gap: var(--space-1);
  overflow-x: auto;
  border-bottom: 1px solid var(--color-border-subtle);
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.tabs-container::-webkit-scrollbar {
  display: none;
}

.tab-btn {
  flex-shrink: 0;
  padding: var(--space-3) var(--space-4);
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 500;
  color: var(--color-text-secondary);
  background: transparent;
  border: none;
  border-bottom: 2px solid transparent;
  cursor: pointer;
  transition: color var(--transition-fast), border-color var(--transition-fast);
  white-space: nowrap;
}

.tab-btn:hover {
  color: var(--color-text-primary);
  border-bottom-color: var(--color-border-default);
}

.tab-btn.active {
  color: var(--color-primary);
  border-bottom-color: var(--color-primary);
}

/* === Content Layout === */
.content-layout {
  max-width: var(--container-max);
  margin: 0 auto;
  padding: 0 var(--content-padding) var(--space-16);
  display: grid;
  grid-template-columns: 1fr;
  gap: var(--space-8);
}

@media (min-width: 1024px) {
  .content-layout {
    grid-template-columns: 2fr 1fr;
  }
}

/* === Resource Grid === */
.resource-grid-container {
  min-width: 0;
}

.loading-state {
  text-align: center;
  padding: var(--space-10);
  color: var(--color-text-tertiary);
  font-family: var(--font-body);
  font-size: var(--text-base);
}

.resource-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: var(--space-6);
}

@media (min-width: 768px) {
  .resource-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

.resource-card {
  background: var(--color-bg-surface);
  border: 1px solid var(--color-border-default);
  border-radius: var(--radius-lg);
  padding: var(--space-6);
  display: flex;
  flex-direction: column;
  transition: border-color var(--transition-base), box-shadow var(--transition-base);
}

.resource-card:hover {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-glow);
}

.resource-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  margin-bottom: var(--space-4);
}

.resource-icon-wrapper {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border-radius: var(--radius-md);
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
}

.resource-icon {
  color: var(--color-text-inverse);
}

.resource-tags-top {
  display: flex;
  align-items: center;
  gap: var(--space-2);
}

.tag {
  display: inline-flex;
  align-items: center;
  padding: 2px var(--space-2);
  border-radius: var(--radius-md);
  font-family: var(--font-body);
  font-size: var(--text-xs);
  white-space: nowrap;
}

.tag-category {
  background: var(--color-bg-muted);
  color: var(--color-text-secondary);
}

.tag-level {
  background: var(--color-primary-tint-1);
  color: var(--color-primary);
}

.resource-title {
  color: var(--color-text-primary);
  margin: 0 0 var(--space-2);
  text-wrap: balance;
  word-break: keep-all;
}

.resource-desc {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  color: var(--color-text-secondary);
  line-height: var(--leading-normal);
  margin: 0;
  flex: 1;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.resource-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: var(--space-4);
  padding-top: var(--space-4);
  border-top: 1px solid var(--color-border-subtle);
}

.resource-meta {
  display: flex;
  align-items: center;
  gap: var(--space-4);
}

.meta-item {
  display: flex;
  align-items: center;
  gap: var(--space-1);
  font-family: var(--font-body);
  font-size: var(--text-xs);
  color: var(--color-text-tertiary);
  white-space: nowrap;
}

.resource-link {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 500;
  color: var(--color-primary);
  text-decoration: none;
  white-space: nowrap;
  transition: color var(--transition-fast);
}

.resource-link:hover {
  color: var(--color-primary-light);
}

/* === Sidebar === */
.sidebar {
  display: flex;
  flex-direction: column;
  gap: var(--space-6);
}

.sidebar-card {
  background: var(--color-bg-surface);
  border: 1px solid var(--color-border-default);
  border-radius: var(--radius-lg);
  padding: var(--space-6);
}

.sidebar-card-title {
  font-size: var(--text-lg);
  color: var(--color-text-primary);
  margin: 0 0 var(--space-5);
  text-wrap: balance;
  word-break: keep-all;
}

/* My Learning Card */
.my-learning-card {
  border-top: 2px solid var(--color-primary);
}

.stats-row {
  display: flex;
  align-items: center;
  gap: var(--space-6);
  margin-bottom: var(--space-6);
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 64px;
}

.stat-value {
  font-size: var(--text-2xl);
  color: var(--color-primary);
}

.stat-value-secondary {
  color: var(--color-text-primary);
}

.stat-label {
  font-family: var(--font-body);
  font-size: var(--text-xs);
  color: var(--color-text-tertiary);
  white-space: nowrap;
}

.stat-divider {
  width: 1px;
  height: 36px;
  background: var(--color-border-subtle);
}

.current-course-section {
  padding-top: var(--space-4);
  border-top: 1px solid var(--color-border-subtle);
}

.current-course-label {
  font-family: var(--font-body);
  font-size: var(--text-xs);
  color: var(--color-text-tertiary);
  margin: 0 0 var(--space-2);
  white-space: nowrap;
}

.current-course-name {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 500;
  color: var(--color-text-primary);
  margin: 0 0 var(--space-3);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.progress-wrapper {
  display: flex;
  align-items: center;
  gap: var(--space-3);
}

.progress-bar {
  flex: 1;
  height: 6px;
  background: var(--color-bg-muted);
  border-radius: var(--radius-full);
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: var(--color-primary);
  border-radius: var(--radius-full);
  transition: width var(--transition-slow);
}

.progress-pct {
  font-family: var(--font-body);
  font-size: var(--text-xs);
  font-weight: 500;
  color: var(--color-primary);
  white-space: nowrap;
}

/* Recommended Path Card */
.path-name {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 500;
  color: var(--color-primary);
  margin: 0 0 var(--space-4);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.path-steps {
  display: flex;
  flex-direction: column;
  gap: var(--space-4);
}

.path-step {
  display: flex;
  align-items: flex-start;
  gap: var(--space-3);
}

.step-indicator {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 22px;
  height: 22px;
  border-radius: 50%;
  margin-top: 2px;
}

.step-completed {
  background: var(--color-primary);
}

.step-in-progress {
  border: 2px solid var(--color-primary);
  background: var(--color-primary-tint-1);
}

.step-locked {
  background: var(--color-bg-muted);
}

.step-check-icon {
  color: var(--color-text-inverse);
}

.step-number {
  font-family: var(--font-body);
  font-size: var(--text-xs);
  font-weight: 700;
  color: var(--color-primary);
}

.step-number-locked {
  font-weight: 500;
  color: var(--color-text-tertiary);
}

.step-content {
  flex: 1;
  min-width: 0;
}

.step-text {
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 500;
  color: var(--color-text-primary);
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.step-text-completed {
  color: var(--color-text-primary);
}

.step-text-locked {
  color: var(--color-text-tertiary);
}

.step-duration {
  font-family: var(--font-body);
  font-size: var(--text-xs);
  color: var(--color-text-tertiary);
  margin: 0;
  white-space: nowrap;
}

/* === Bottom CTA === */
.bottom-cta {
  max-width: var(--container-max);
  margin: 0 auto var(--space-16);
  padding: 0 var(--content-padding);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  gap: var(--space-6);
  background: var(--color-bg-surface);
  border: 1px solid var(--color-border-default);
  border-radius: var(--radius-lg);
  padding: var(--space-8);
}

@media (min-width: 768px) {
  .bottom-cta {
    flex-direction: row;
  }
}

.cta-content {
  display: flex;
  align-items: center;
  gap: var(--space-3);
}

.cta-icon {
  color: var(--color-primary);
  flex-shrink: 0;
}

.cta-text {
  font-family: var(--font-body);
  font-size: var(--text-base);
  color: var(--color-text-primary);
  margin: 0;
}

.cta-highlight {
  color: var(--color-primary);
}

.cta-btn {
  flex-shrink: 0;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: var(--space-2) var(--space-5);
  border: 1px solid var(--color-primary);
  border-radius: var(--radius-md);
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 600;
  color: var(--color-primary);
  background: transparent;
  text-decoration: none;
  white-space: nowrap;
  transition: background var(--transition-fast), color var(--transition-fast);
}

.cta-btn:hover {
  background: var(--color-primary);
  color: var(--color-text-inverse);
}
</style>
