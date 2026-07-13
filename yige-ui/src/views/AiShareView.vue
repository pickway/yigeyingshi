<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { aiApi, articleApi } from '@/api'
import { Video, Image, AudioWaveform, ArrowRight } from '@lucide/vue'

const tools = ref([])
const articles = ref([])
const loading = ref(true)
const subscribing = ref(false)
const subscribeMsg = ref('')

const iconMap = {
  video: Video,
  image: Image,
  'audio-waveform': AudioWaveform,
}

const form = reactive({ email: '' })

async function load() {
  loading.value = true
  try {
    const [toolsRes, articlesRes] = await Promise.all([
      aiApi.featuredTools(),
      articleApi.list({ category: 'ai-article' }),
    ])
    tools.value = toolsRes.data || []
    articles.value = articlesRes.data || []
  } catch (e) {
    console.error('加载 AI 数据失败', e)
  } finally {
    loading.value = false
  }
}

async function subscribe() {
  if (!form.email) return
  subscribing.value = true
  subscribeMsg.value = ''
  try {
    const res = await aiApi.subscribe(form.email)
    subscribeMsg.value = res.message || '订阅成功'
    form.email = ''
  } catch (e) {
    subscribeMsg.value = e.message || '订阅失败'
  } finally {
    subscribing.value = false
  }
}

function formatDate(iso) {
  if (!iso) return ''
  const d = new Date(iso)
  return `${d.getFullYear()}-${String(d.getMonth()+1).padStart(2,'0')}-${String(d.getDate()).padStart(2,'0')}`
}

function primaryTag(tags) {
  if (!tags) return ''
  return tags.split(',')[0]
}

onMounted(load)
</script>

<template>
  <div class="ai-share-page">
    <!-- AI Hero -->
    <section class="hero">
      <div class="container hero-grid">
        <div class="hero-left">
          <h1 class="cine-display hero-title">智慧光影</h1>
          <p class="cine-body-sm hero-subtitle">AI 赋能创作，探索影视与技术的交汇点</p>
          <a href="#tool-showcase" class="cta-button">
            开始探索
            <ArrowRight :size="18" />
          </a>
        </div>
        <div class="hero-right">
          <img src="/ai-bright.jpg" alt="AI Bright" class="hero-image" />
        </div>
      </div>
    </section>

    <!-- Tool Showcase -->
    <section id="tool-showcase" class="tool-showcase">
      <div class="container">
        <h2 class="cine-heading section-title">推荐工具</h2>
        <div v-if="loading" class="tool-grid">
          <div v-for="i in 3" :key="i" class="tool-card">
            <div class="tool-icon-box">
              <Video :size="24" />
            </div>
            <h3 class="tool-title">加载中…</h3>
            <p class="tool-desc">正在获取推荐工具</p>
          </div>
        </div>
        <div v-else class="tool-grid">
          <div
            v-for="tool in tools.slice(0, 3)"
            :key="tool.id || tool.name"
            class="tool-card"
          >
            <div class="tool-icon-box">
              <component :is="iconMap[tool.icon] || Video" :size="24" />
            </div>
            <h3 class="tool-title">{{ tool.name }}</h3>
            <p class="tool-desc">{{ tool.description }}</p>
            <div v-if="tool.tags" class="tool-tags">
              <span
                v-for="tag in tool.tags.split(',')"
                :key="tag"
                class="tool-tag"
              >
                {{ tag.trim() }}
              </span>
            </div>
            <a :href="tool.url || '#'" class="tool-link" target="_blank" rel="noopener">查看详情 →</a>
          </div>
        </div>
      </div>
    </section>

    <!-- Article List -->
    <section class="article-list">
      <div class="container">
        <h2 class="cine-heading section-title">技术文章</h2>
        <div class="article-card">
          <div v-if="loading" class="article-row">
            <div class="article-number">…</div>
            <div class="article-content">
              <h3 class="article-title">加载中…</h3>
              <p class="article-desc">正在获取技术文章</p>
            </div>
          </div>
          <div
            v-else
            v-for="(article, index) in articles"
            :key="article.id || index"
            class="article-row"
          >
            <div class="article-number">{{ index + 1 }}</div>
            <div class="article-content">
              <h3 class="article-title">{{ article.title }}</h3>
              <p class="article-desc">{{ article.summary || article.content }}</p>
            </div>
            <div class="article-meta">
              <span class="article-tag">{{ primaryTag(article.tag) }}</span>
              <span class="article-date">{{ formatDate(article.publishedAt) }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Newsletter CTA -->
    <section class="newsletter">
      <div class="container">
        <div class="newsletter-card">
          <div class="newsletter-decoration"></div>
          <h2 class="cine-heading newsletter-title">订阅更新</h2>
          <p class="cine-body-sm newsletter-subtitle">
            获取最新的 AI 影视工具资讯与技术文章，直达你的收件箱
          </p>
          <form class="newsletter-form" @submit.prevent="subscribe">
            <input
              v-model="form.email"
              type="email"
              placeholder="输入你的邮箱地址"
              class="newsletter-input"
            />
            <button
              type="submit"
              class="newsletter-button"
              :disabled="subscribing"
            >
              {{ subscribing ? '订阅中…' : '订阅' }}
            </button>
          </form>
          <p v-if="subscribeMsg" class="cine-body-sm subscribe-msg">{{ subscribeMsg }}</p>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.ai-share-page {
  --color-primary: #E85D3A;
  --color-primary-light: #F2845F;
  --color-primary-dark: #C44A2B;
  --color-primary-tint-1: rgba(232, 93, 58, 0.10);
  --color-primary-tint-2: rgba(232, 93, 58, 0.05);
  --color-bg-base: #FAFAF8;
  --color-bg-elevated: #FFFFFF;
  --color-bg-surface: #F3F2EF;
  --color-border-subtle: #EDECE8;
  --color-border-default: #E0DFDB;
  --color-text-primary: #1A1A1A;
  --color-text-secondary: #6B6B6B;
  --color-text-tertiary: #9A9A96;
  --color-text-inverse: #FFFFFF;
  --shadow-sm: 0 1px 2px rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 12px rgba(0, 0, 0, 0.08);
  --radius-sm: 6px;
  --radius-md: 10px;
  --radius-lg: 16px;
  --radius-full: 9999px;
  --container-max: 1200px;
  --content-padding: 24px;
  --font-display: 'Playfair Display', 'Noto Serif SC', Georgia, serif;
  --font-body: 'Inter', 'Noto Sans SC', system-ui, sans-serif;

  background: var(--color-bg-base);
  color: var(--color-text-primary);
  font-family: var(--font-body);
}

.container {
  max-width: var(--container-max);
  margin-inline: auto;
  padding-inline: var(--content-padding);
}

.cine-display {
  font-family: var(--font-display);
  font-weight: 700;
  line-height: 1.1;
  letter-spacing: -0.02em;
}

.cine-heading {
  font-family: var(--font-display);
  font-weight: 600;
  line-height: 1.25;
  letter-spacing: -0.01em;
}

.cine-body-sm {
  font-family: var(--font-body);
  font-size: 15px;
  line-height: 1.6;
  color: var(--color-text-secondary);
}

/* ── Hero ── */
.hero {
  padding-block: 80px 60px;
}

.hero-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 40px;
  align-items: center;
}

.hero-title {
  font-size: clamp(40px, 6vw, 72px);
  color: var(--color-text-primary);
  margin-bottom: 20px;
}

.hero-subtitle {
  margin-bottom: 36px;
  max-width: 480px;
}

.cta-button {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 14px 32px;
  background: var(--color-primary);
  color: var(--color-text-inverse);
  font-family: var(--font-body);
  font-size: 15px;
  font-weight: 600;
  border-radius: var(--radius-md);
  text-decoration: none;
  transition: background 0.2s;
}

.cta-button:hover {
  background: var(--color-primary-dark);
}

.hero-right {
  display: flex;
  justify-content: center;
}

.hero-image {
  width: 100%;
  aspect-ratio: 4 / 3;
  object-fit: cover;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-md);
}

@media (min-width: 1024px) {
  .hero-grid {
    grid-template-columns: repeat(5, 1fr);
  }
  .hero-left {
    grid-column: span 3;
  }
  .hero-right {
    grid-column: span 2;
  }
}

/* ── Tool Showcase ── */
.tool-showcase {
  padding-block: 60px;
}

.section-title {
  font-size: 32px;
  margin-bottom: 36px;
}

.tool-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 24px;
}

@media (min-width: 768px) {
  .tool-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

.tool-card {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border-subtle);
  border-radius: var(--radius-lg);
  padding: 28px;
  display: flex;
  flex-direction: column;
  gap: 14px;
  transition: box-shadow 0.2s;
}

.tool-card:hover {
  box-shadow: var(--shadow-md);
}

.tool-icon-box {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--color-primary-tint-1);
  color: var(--color-primary);
  border-radius: var(--radius-md);
}

.tool-title {
  font-family: var(--font-display);
  font-size: 20px;
  font-weight: 600;
  color: var(--color-text-primary);
}

.tool-desc {
  font-size: 14px;
  line-height: 1.6;
  color: var(--color-text-secondary);
  flex: 1;
}

.tool-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tool-tag {
  display: inline-block;
  padding: 4px 12px;
  font-size: 12px;
  font-weight: 500;
  background: var(--color-primary-tint-2);
  color: var(--color-primary);
  border-radius: var(--radius-full);
}

.tool-link {
  font-size: 14px;
  font-weight: 600;
  color: var(--color-primary);
  text-decoration: none;
  transition: color 0.2s;
}

.tool-link:hover {
  color: var(--color-primary-dark);
}

/* ── Article List ── */
.article-list {
  padding-block: 60px;
}

.article-card {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border-subtle);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
  overflow: hidden;
}

.article-row {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 20px 28px;
  transition: background 0.2s;
}

.article-row:hover {
  background: var(--color-bg-surface);
}

.article-row:not(:last-child) {
  border-bottom: 1px solid var(--color-border-subtle);
}

.article-number {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  background: var(--color-primary);
  color: var(--color-text-inverse);
  font-weight: 700;
  font-size: 14px;
  border-radius: var(--radius-sm);
}

.article-content {
  flex: 1;
  min-width: 0;
}

.article-title {
  font-family: var(--font-display);
  font-size: 16px;
  font-weight: 600;
  color: var(--color-text-primary);
  margin-bottom: 4px;
}

.article-desc {
  font-size: 13px;
  line-height: 1.5;
  color: var(--color-text-tertiary);
}

.article-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 4px;
  flex-shrink: 0;
}

.article-tag {
  display: inline-block;
  padding: 2px 10px;
  font-size: 12px;
  font-weight: 500;
  background: var(--color-primary-tint-2);
  color: var(--color-primary);
  border-radius: var(--radius-full);
}

.article-date {
  font-size: 12px;
  color: var(--color-text-tertiary);
}

@media (max-width: 640px) {
  .article-row {
    flex-wrap: wrap;
    gap: 12px;
  }
  .article-meta {
    flex-direction: row;
    align-items: center;
    width: 100%;
    padding-left: 56px;
  }
}

/* ── Newsletter ── */
.newsletter {
  padding-block: 60px 80px;
}

.newsletter-card {
  position: relative;
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border-subtle);
  border-radius: var(--radius-lg);
  padding: 48px 40px;
  text-align: center;
  overflow: hidden;
}

.newsletter-decoration {
  position: absolute;
  top: -40px;
  right: -40px;
  width: 160px;
  height: 160px;
  background: var(--color-primary-tint-1);
  border-radius: 50%;
  pointer-events: none;
}

.newsletter-title {
  font-size: 28px;
  color: var(--color-text-primary);
  margin-bottom: 12px;
}

.newsletter-subtitle {
  max-width: 420px;
  margin-inline: auto;
  margin-bottom: 32px;
}

.newsletter-form {
  display: flex;
  gap: 12px;
  max-width: 440px;
  margin-inline: auto;
}

.newsletter-input {
  flex: 1;
  padding: 12px 16px;
  font-size: 14px;
  font-family: var(--font-body);
  background: var(--color-bg-base);
  border: 1px solid var(--color-border-default);
  border-radius: var(--radius-md);
  color: var(--color-text-primary);
  outline: none;
  transition: border-color 0.2s;
}

.newsletter-input::placeholder {
  color: var(--color-text-tertiary);
}

.newsletter-input:focus {
  border-color: var(--color-primary);
}

.newsletter-button {
  padding: 12px 28px;
  font-size: 14px;
  font-weight: 600;
  font-family: var(--font-body);
  background: var(--color-primary);
  color: var(--color-text-inverse);
  border: none;
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: background 0.2s;
}

.newsletter-button:hover {
  background: var(--color-primary-dark);
}

.newsletter-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.subscribe-msg {
  margin-top: 16px;
  color: var(--color-text-secondary);
}

@media (max-width: 480px) {
  .newsletter-form {
    flex-direction: column;
  }
  .newsletter-card {
    padding: 36px 24px;
  }
}
</style>
