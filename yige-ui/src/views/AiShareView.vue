<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { aiApi, articleApi } from '@/api'
import UiState from '@/components/UiState.vue'
import { normalizeExternalUrl } from '@/utils/content'
import { Video, Image, AudioWaveform, ArrowRight } from '@lucide/vue'

const tools = ref([])
const articles = ref([])
const loading = ref(true)
const error = ref('')
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
  error.value = ''
  try {
    const [toolsRes, articlesRes] = await Promise.all([
      aiApi.featuredTools(),
      articleApi.list({ category: 'ai-article' }),
    ])
    tools.value = toolsRes.data || []
    articles.value = articlesRes.data || []
  } catch (e) {
    error.value = e.message
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
      <div class="hero-bg">
        <img src="/ai-bright.jpg" alt="AI abstract visual" class="hero-bg-image" />
        <div class="hero-bg-overlay"></div>
      </div>
      <div class="container hero-content">
        <h1 class="cinema-display hero-title">智慧光影</h1>
        <p class="cinema-subheading hero-subtitle">AI 赋能创作，探索影视与技术的交汇点</p>
        <a href="#tool-showcase" class="cta-button">
          开始探索
        </a>
      </div>
    </section>

    <!-- Tool Showcase -->
    <section id="tool-showcase" class="tool-showcase">
      <div class="container">
        <h2 class="cinema-heading section-title">推荐工具</h2>
        <UiState v-if="error" type="error" :message="error" @retry="load" />
        <div v-else-if="loading" class="tool-grid">
          <div v-for="i in 3" :key="i" class="tool-card">
            <div class="tool-icon-box">
              <Video :size="32" />
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
              <component :is="iconMap[tool.icon] || Video" :size="32" />
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
            <a :href="normalizeExternalUrl(tool.url)" class="tool-link" target="_blank" rel="noopener noreferrer">
              查看详情
              <ArrowRight :size="14" />
            </a>
          </div>
        </div>
      </div>
    </section>

    <!-- Article List -->
    <section class="article-list">
      <div class="container">
        <h2 class="cinema-heading section-title">技术文章</h2>
        <div class="article-card">
          <div v-if="loading" class="article-row">
            <div class="article-number">…</div>
            <div class="article-content">
              <h3 class="article-title">加载中…</h3>
              <p class="article-desc">正在获取技术文章</p>
            </div>
          </div>
          <router-link
            v-else
            v-for="(article, index) in articles"
            :key="article.id || index"
            :to="`/articles/${article.id}`"
            class="article-row"
            style="color:inherit;text-decoration:none;"
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
          </router-link>
        </div>
      </div>
    </section>

    <!-- Newsletter CTA -->
    <section class="newsletter">
      <div class="newsletter-glow"></div>
      <div class="container newsletter-content">
        <h2 class="cinema-heading newsletter-title">订阅更新</h2>
        <p class="cinema-body-sm newsletter-subtitle">
          每周精选 AI 工具和影视创作技巧，直达你的邮箱
        </p>
        <form class="newsletter-form" @submit.prevent="subscribe">
          <input
            v-model="form.email"
            type="email"
            placeholder="输入你的邮箱"
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
        <p v-if="subscribeMsg" class="cinema-body-sm subscribe-msg">{{ subscribeMsg }}</p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.ai-share-page {
  /* Cinema Personal Site - Dark Theater Theme */
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

  --color-border-default: #2A2A30;
  --color-border-subtle: #1F1F25;
  --color-border-strong: #3A3A42;

  --color-text-primary: #F0EDE6;
  --color-text-secondary: #9A9AA0;
  --color-text-tertiary: #6A6A72;
  --color-text-inverse: #0C0C0E;

  --radius-sm: 4px;
  --radius-md: 8px;
  --radius-lg: 12px;
  --radius-full: 9999px;

  --shadow-sm: 0 1px 3px rgba(0,0,0,0.3);
  --shadow-md: 0 4px 12px rgba(0,0,0,0.3);
  --shadow-lg: 0 8px 24px rgba(0,0,0,0.4);
  --shadow-glow: 0 0 20px rgba(212,168,83,0.15);

  --container-max: 1200px;
  --content-padding: 24px;

  --font-display: 'Playfair Display', Georgia, 'Noto Serif SC', serif;
  --font-body: 'Inter', -apple-system, 'PingFang SC', 'Microsoft YaHei', sans-serif;

  background: var(--color-bg-base);
  color: var(--color-text-primary);
  font-family: var(--font-body);
}

.container {
  max-width: var(--container-max);
  margin-inline: auto;
  padding-inline: var(--content-padding);
}

/* === Typography Classes === */
.cinema-display {
  font-family: var(--font-display);
  font-weight: 700;
  letter-spacing: -0.02em;
  line-height: 1.25;
}

.cinema-heading {
  font-family: var(--font-display);
  font-weight: 600;
  font-size: 24px;
  letter-spacing: -0.02em;
  line-height: 1.35;
}

.cinema-subheading {
  font-family: var(--font-body);
  font-weight: 500;
  font-size: 17px;
  line-height: 1.35;
  color: var(--color-text-secondary);
}

.cinema-body-sm {
  font-family: var(--font-body);
  font-size: 13px;
  line-height: 1.6;
  color: var(--color-text-secondary);
}

/* ── Hero ── */
.hero {
  position: relative;
  overflow: hidden;
  padding-block: 80px;
}

.hero-bg {
  position: absolute;
  inset: 0;
  z-index: 0;
}

.hero-bg-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.35;
}

.hero-bg-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, var(--color-bg-base) 0%, rgba(12,12,14,0.6) 40%, rgba(12,12,14,0.85) 100%);
}

.hero-content {
  position: relative;
  z-index: 10;
  text-align: center;
}

.hero-title {
  font-size: clamp(40px, 6vw, 72px);
  color: var(--color-text-primary);
  margin-bottom: 24px;
}

.hero-subtitle {
  font-size: 17px;
  color: var(--color-text-secondary);
  max-width: 520px;
  margin-inline: auto;
  margin-bottom: 40px;
}

.cta-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 12px 28px;
  background: var(--color-primary);
  color: var(--color-text-inverse);
  font-family: var(--font-body);
  font-size: 14px;
  font-weight: 600;
  border-radius: var(--radius-md);
  text-decoration: none;
  box-shadow: var(--shadow-glow);
  transition: all 0.25s ease;
}

.cta-button:hover {
  background: var(--color-primary-light);
  transform: translateY(-1px);
}

/* ── Tool Showcase ── */
.tool-showcase {
  padding-block: 80px;
}

.section-title {
  font-size: 30px;
  color: var(--color-text-primary);
  margin-bottom: 48px;
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
  background: var(--color-bg-surface);
  border: 1px solid var(--color-border-default);
  border-radius: var(--radius-lg);
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  transition: all 0.25s ease;
  cursor: pointer;
}

.tool-card:hover {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-glow);
}

.tool-icon-box {
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(212,168,83,0.25) 0%, rgba(212,168,83,0.05) 100%);
  border-radius: var(--radius-lg);
  margin-bottom: 8px;
  color: var(--color-primary);
}

.tool-title {
  font-family: var(--font-display);
  font-size: 20px;
  font-weight: 600;
  color: var(--color-text-primary);
}

.tool-desc {
  font-size: 13px;
  line-height: 1.6;
  color: var(--color-text-secondary);
  flex: 1;
}

.tool-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 4px;
}

.tool-tag {
  display: inline-flex;
  align-items: center;
  height: 22px;
  padding: 2px 10px;
  font-size: 12px;
  font-weight: 500;
  background: var(--color-bg-muted);
  color: var(--color-text-secondary);
  border-radius: var(--radius-full);
}

.tool-tag:first-child {
  background: var(--color-primary-tint-1);
  color: var(--color-primary);
}

.tool-link {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
  font-weight: 500;
  color: var(--color-primary);
  text-decoration: none;
  transition: color 0.15s ease;
}

.tool-link:hover {
  color: var(--color-primary-light);
}

/* ── Article List ── */
.article-list {
  padding-block: 80px;
}

.article-card {
  display: flex;
  flex-direction: column;
  border: 1px solid var(--color-border-subtle);
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.article-row {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 24px;
  border-bottom: 1px solid var(--color-border-subtle);
  transition: background 0.15s ease;
}

.article-row:last-child {
  border-bottom: none;
}

.article-row:hover {
  background: var(--color-bg-surface);
}

.article-number {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  background: var(--color-primary-tint-1);
  color: var(--color-primary);
  font-family: var(--font-display);
  font-weight: 700;
  font-size: 13px;
  border-radius: var(--radius-full);
}

.article-content {
  flex: 1;
  min-width: 0;
}

.article-title {
  font-family: var(--font-display);
  font-size: 15px;
  font-weight: 600;
  color: var(--color-text-primary);
  margin-bottom: 4px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.article-desc {
  font-size: 13px;
  line-height: 1.6;
  color: var(--color-text-tertiary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.article-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 8px;
  flex-shrink: 0;
}

.article-tag {
  display: inline-flex;
  align-items: center;
  height: 22px;
  padding: 2px 10px;
  font-size: 12px;
  font-weight: 500;
  background: var(--color-bg-muted);
  color: var(--color-text-secondary);
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
    padding-left: 60px;
    gap: 12px;
  }
}

/* ── Newsletter ── */
.newsletter {
  position: relative;
  padding-block: 80px;
  background: var(--color-bg-elevated);
}

.newsletter-glow {
  position: absolute;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 600px;
  height: 1px;
  background: linear-gradient(90deg, transparent 0%, var(--color-primary-tint-1) 30%, var(--color-primary) 50%, var(--color-primary-tint-1) 70%, transparent 100%);
}

.newsletter-content {
  text-align: center;
}

.newsletter-title {
  font-size: 30px;
  color: var(--color-text-primary);
  margin-bottom: 16px;
}

.newsletter-subtitle {
  max-width: 480px;
  margin-inline: auto;
  margin-bottom: 32px;
}

.newsletter-form {
  display: flex;
  align-items: center;
  gap: 12px;
  justify-content: center;
  max-width: 440px;
  margin-inline: auto;
}

.newsletter-input {
  flex: 1;
  min-width: 0;
  height: 42px;
  padding: 0 20px;
  font-size: 14px;
  font-family: var(--font-body);
  background: var(--color-bg-surface);
  border: 1px solid var(--color-border-default);
  border-radius: var(--radius-full);
  color: var(--color-text-primary);
  outline: none;
  transition: border-color 0.15s ease;
}

.newsletter-input::placeholder {
  color: var(--color-text-tertiary);
}

.newsletter-input:focus {
  border-color: var(--color-primary);
}

.newsletter-button {
  flex-shrink: 0;
  height: 42px;
  padding: 0 24px;
  font-size: 14px;
  font-weight: 600;
  font-family: var(--font-body);
  background: var(--color-primary);
  color: var(--color-text-inverse);
  border: none;
  border-radius: var(--radius-full);
  cursor: pointer;
  box-shadow: var(--shadow-glow);
  transition: all 0.25s ease;
}

.newsletter-button:hover {
  background: var(--color-primary-light);
  transform: translateY(-1px);
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
  .newsletter-input,
  .newsletter-button {
    width: 100%;
  }
}
</style>
