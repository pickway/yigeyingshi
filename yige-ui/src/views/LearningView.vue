<template>
  <div class="learning-page">
    <!-- Compact Hero -->
    <div class="hero-card">
      <nav class="breadcrumb">
        <router-link to="/" class="breadcrumb-link">首页</router-link>
        <ChevronRight class="breadcrumb-sep" :size="14" />
        <span class="breadcrumb-current">学习资料</span>
      </nav>
      <h1 class="hero-title">学习资料</h1>
      <p class="hero-subtitle">系统化的影视创作课程，从基础理论到实战技能，助你快速成长</p>
    </div>

    <!-- Category Tabs -->
    <div class="tabs-wrapper">
      <div class="tabs-inner">
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
    </div>

    <!-- Two-Column Layout -->
    <div class="main-layout">
      <!-- Left: Resource Grid -->
      <div class="resource-grid">
        <div v-if="loading" class="loading-state">加载中…</div>
        <div
          v-for="course in allCourses"
          v-else
          :key="course.id"
          class="resource-card"
        >
          <div class="resource-header">
            <component
              :is="iconMap[course.icon] || Clapperboard"
              class="resource-icon"
              :size="20"
            />
            <h3 class="resource-title">{{ course.title }}</h3>
          </div>
          <p class="resource-desc">{{ course.description }}</p>
          <div class="resource-tags">
            <span class="tag tag-category">{{ course.category }}</span>
            <span class="tag tag-level">{{ course.level }}</span>
          </div>
          <div class="resource-footer">
            <span class="resource-hours">{{ course.duration }}课时</span>
            <a href="#" class="resource-link">开始学习 <ChevronRight :size="14" /></a>
          </div>
        </div>
      </div>

      <!-- Right: Sidebar -->
      <aside class="sidebar">
        <!-- My Learning Card -->
        <div class="sidebar-card my-learning-card">
          <h3 class="sidebar-card-title">我的学习</h3>
          <div class="stats-row">
            <div class="stat-item">
              <span class="stat-value">{{ pathCourses.length }}</span>
              <span class="stat-label">进行中</span>
            </div>
            <div class="stat-item">
              <span class="stat-value">{{ myLearning.completed }}</span>
              <span class="stat-label">已完成</span>
            </div>
            <div class="stat-item">
              <span class="stat-value">{{ myLearning.totalHours }}h</span>
              <span class="stat-label">总学时</span>
            </div>
          </div>
          <div class="progress-section">
            <div class="progress-label">
              <span class="progress-course">{{ myLearning.currentCourse }}</span>
              <span class="progress-pct">{{ myLearning.currentProgress }}%</span>
            </div>
            <div class="progress-bar">
              <div
                class="progress-fill"
                :style="{ width: myLearning.currentProgress + '%' }"
              />
            </div>
          </div>
        </div>

        <!-- Recommended Path Card -->
        <div class="sidebar-card path-card">
          <h3 class="sidebar-card-title">
            <Route :size="16" class="path-title-icon" />
            推荐学习路径
          </h3>
          <p class="path-name">{{ recommendedPath ? recommendedPath.name : '暂无推荐路径' }}</p>
          <div class="path-steps">
            <div
              v-for="(course, i) in pathCourses"
              :key="course.id"
              class="path-step"
            >
              <Check v-if="i < 2" :size="16" class="step-icon step-done" />
              <span v-else class="step-icon step-dot" />
              <span class="step-text" :class="{ 'step-text-done': i < 2 }">{{ course.title }}</span>
            </div>
          </div>
        </div>
      </aside>
    </div>

    <!-- Bottom CTA -->
    <div class="bottom-cta">
      <p class="cta-text">找不到需要的资源？</p>
      <router-link to="/ai-share" class="cta-btn">
        <MessageCirclePlus :size="16" />
        提交需求
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { learningApi } from '@/api'
import {
  ChevronRight,
  Clapperboard,
  Palette,
  SlidersHorizontal,
  Sparkles,
  Volume2,
  Video,
  Route,
  Check,
  MessageCirclePlus,
} from '@lucide/vue'

const allCourses = ref([])
const paths = ref([])
const loading = ref(true)
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
    console.error('加载学习数据失败', e)
  } finally {
    loading.value = false
  }
}

onMounted(load)
watch(activeTab, load)

// First path as the recommended one
const recommendedPath = computed(() => paths.value[0] || null)

// path.courseIds is comma-separated string "6,1,2,3"
const pathCourses = computed(() => {
  if (!recommendedPath.value || !recommendedPath.value.courseIds) return []
  const ids = recommendedPath.value.courseIds.split(',').map(s => parseInt(s.trim()))
  return ids.map(id => allCourses.value.find(c => c.id === id)).filter(Boolean)
})

const myLearning = ref({
  completed: 12,
  totalHours: 28,
  currentCourse: 'Premiere Pro 快速入门',
  currentProgress: 67,
})
</script>

<style scoped>
.learning-page {
  min-height: 100vh;
  background: var(--color-bg-base);
  padding-bottom: 64px;
}

/* ---- Hero ---- */
.hero-card {
  max-width: var(--container-max);
  margin: 0 auto;
  padding: 24px var(--content-padding);
  background: var(--color-bg-elevated);
  border-top: 3px solid var(--color-primary);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
  margin-top: 24px;
}

.breadcrumb {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-bottom: 12px;
  font-size: 13px;
}

.breadcrumb-link {
  color: var(--color-text-tertiary);
  text-decoration: none;
  transition: color 0.2s;
}

.breadcrumb-link:hover {
  color: var(--color-primary);
}

.breadcrumb-sep {
  color: var(--color-text-tertiary);
}

.breadcrumb-current {
  color: var(--color-text-secondary);
}

.hero-title {
  font-family: var(--font-display);
  font-size: 28px;
  font-weight: 700;
  color: var(--color-text-primary);
  margin: 0 0 6px;
}

.hero-subtitle {
  font-family: var(--font-body);
  font-size: 14px;
  color: var(--color-text-secondary);
  margin: 0;
}

/* ---- Tabs ---- */
.tabs-wrapper {
  max-width: var(--container-max);
  margin: 0 auto;
  padding: 0 var(--content-padding);
  margin-top: 20px;
}

.tabs-inner {
  display: flex;
  gap: 0;
  border-bottom: 1px solid var(--color-border-default);
}

.tab-btn {
  padding: 10px 20px;
  font-family: var(--font-body);
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text-secondary);
  background: transparent;
  border: none;
  border-bottom: 2px solid transparent;
  cursor: pointer;
  transition: color 0.2s, border-color 0.2s;
  white-space: nowrap;
}

.tab-btn:hover {
  color: var(--color-text-primary);
}

.tab-btn.active {
  color: var(--color-primary);
  border-bottom-color: var(--color-primary);
}

/* ---- Main Layout ---- */
.main-layout {
  max-width: var(--container-max);
  margin: 0 auto;
  padding: 24px var(--content-padding) 0;
  display: flex;
  gap: 24px;
}

/* ---- Resource Grid ---- */
.resource-grid {
  flex: 1;
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.loading-state {
  grid-column: 1 / -1;
  text-align: center;
  padding: 40px;
  color: var(--color-text-tertiary);
  font-size: 14px;
}

.resource-card {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border-subtle);
  border-radius: var(--radius-md);
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  transition: box-shadow 0.2s, border-color 0.2s;
}

.resource-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--color-border-default);
}

.resource-header {
  display: flex;
  align-items: center;
  gap: 8px;
}

.resource-icon {
  color: var(--color-primary);
  flex-shrink: 0;
}

.resource-title {
  font-family: var(--font-body);
  font-size: 15px;
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0;
}

.resource-desc {
  font-family: var(--font-body);
  font-size: 13px;
  color: var(--color-text-secondary);
  margin: 0;
  line-height: 1.5;
}

.resource-tags {
  display: flex;
  gap: 6px;
}

.tag {
  font-family: var(--font-body);
  font-size: 11px;
  font-weight: 500;
  padding: 2px 8px;
  border-radius: var(--radius-full);
}

.tag-category {
  background: var(--color-primary-tint-1);
  color: var(--color-primary);
}

.tag-level {
  background: var(--color-bg-surface);
  color: var(--color-text-secondary);
}

.resource-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: auto;
  padding-top: 8px;
  border-top: 1px solid var(--color-border-subtle);
}

.resource-hours {
  font-family: var(--font-body);
  font-size: 12px;
  color: var(--color-text-tertiary);
}

.resource-link {
  display: inline-flex;
  align-items: center;
  gap: 2px;
  font-family: var(--font-body);
  font-size: 13px;
  font-weight: 500;
  color: var(--color-primary);
  text-decoration: none;
  transition: color 0.2s;
}

.resource-link:hover {
  color: var(--color-primary-dark);
}

/* ---- Sidebar ---- */
.sidebar {
  width: 320px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.sidebar-card {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border-subtle);
  border-radius: var(--radius-md);
  padding: 20px;
  box-shadow: var(--shadow-sm);
}

.sidebar-card-title {
  font-family: var(--font-body);
  font-size: 15px;
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0 0 16px;
  display: flex;
  align-items: center;
  gap: 6px;
}

/* My Learning */
.my-learning-card {
  border-top: 3px solid var(--color-primary);
}

.stats-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 16px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
}

.stat-value {
  font-family: var(--font-display);
  font-size: 20px;
  font-weight: 700;
  color: var(--color-primary);
}

.stat-label {
  font-family: var(--font-body);
  font-size: 12px;
  color: var(--color-text-tertiary);
}

.progress-section {
  margin-top: 4px;
}

.progress-label {
  display: flex;
  justify-content: space-between;
  margin-bottom: 6px;
}

.progress-course {
  font-family: var(--font-body);
  font-size: 13px;
  color: var(--color-text-primary);
}

.progress-pct {
  font-family: var(--font-body);
  font-size: 13px;
  font-weight: 600;
  color: var(--color-primary);
}

.progress-bar {
  height: 6px;
  background: var(--color-bg-surface);
  border-radius: var(--radius-full);
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: var(--color-primary);
  border-radius: var(--radius-full);
  transition: width 0.3s ease;
}

/* Recommended Path */
.path-title-icon {
  color: var(--color-primary);
}

.path-name {
  font-family: var(--font-body);
  font-size: 14px;
  font-weight: 600;
  color: var(--color-primary);
  margin: 0 0 12px;
}

.path-steps {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.path-step {
  display: flex;
  align-items: center;
  gap: 8px;
}

.step-icon {
  flex-shrink: 0;
}

.step-done {
  color: var(--color-primary);
}

.step-dot {
  width: 16px;
  height: 16px;
  border: 2px solid var(--color-border-default);
  border-radius: 50%;
  display: inline-block;
}

.step-text {
  font-family: var(--font-body);
  font-size: 13px;
  color: var(--color-text-primary);
}

.step-text-done {
  color: var(--color-text-tertiary);
  text-decoration: line-through;
}

/* ---- Bottom CTA ---- */
.bottom-cta {
  max-width: var(--container-max);
  margin: 40px auto 0;
  padding: 0 var(--content-padding);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
}

.cta-text {
  font-family: var(--font-body);
  font-size: 14px;
  color: var(--color-text-secondary);
  margin: 0;
}

.cta-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-family: var(--font-body);
  font-size: 13px;
  font-weight: 500;
  color: var(--color-primary);
  background: transparent;
  border: 1px solid var(--color-primary);
  border-radius: var(--radius-sm);
  padding: 6px 14px;
  text-decoration: none;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}

.cta-btn:hover {
  background: var(--color-primary-tint-1);
}
</style>
