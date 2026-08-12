<script setup>
import { onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ArrowLeft, ArrowRight, CalendarDays, Clock3 } from '@lucide/vue'
import { articleApi } from '@/api'
import UiState from '@/components/UiState.vue'
import { formatChineseDate } from '@/utils/content'

const route = useRoute()
const article = ref(null)
const related = ref([])
const loading = ref(true)
const error = ref('')

async function load() {
  loading.value = true
  error.value = ''
  try {
    const response = await articleApi.detail(route.params.id)
    article.value = response.data
    related.value = response.related || []
    document.title = `${article.value.title} - CineVerse`
  } catch (loadError) { error.value = loadError.message } finally { loading.value = false }
}

onMounted(load)
watch(() => route.params.id, load)
</script>

<template>
  <div class="article-page">
    <div class="article-shell">
      <router-link to="/" class="back-link"><ArrowLeft :size="16" /> 返回首页</router-link>
      <UiState v-if="loading" type="loading" title="正在展开笔记…" />
      <UiState v-else-if="error" type="error" :message="error" @retry="load" />
      <article v-else-if="article">
        <header>
          <span class="eyebrow">{{ article.tag || 'Journal' }}</span>
          <h1>{{ article.title }}</h1>
          <p class="lead">{{ article.summary }}</p>
          <div class="meta"><span>{{ article.author || '一个影视' }}</span><span><CalendarDays :size="15" />{{ formatChineseDate(article.publishedAt) }}</span><span><Clock3 :size="15" />6 分钟阅读</span></div>
        </header>
        <div class="cover" :style="{ '--cover': article.coverColor || '#2c3e6b' }"><span>{{ article.category === 'ai-article' ? 'AI / IMAGE' : 'CINEMA / NOTE' }}</span></div>
        <div class="prose">
          <p>{{ article.content || article.summary }}</p>
          <h2>从观看开始</h2>
          <p>真正打动人的作品，往往不是因为它给出了一个无懈可击的答案，而是因为它让我们重新留意那些习以为常的细节：一次停顿、一束光、一个角色没有说出口的话。</p>
          <blockquote>影像的价值，不只在于记录发生过什么，也在于提醒我们该如何观看。</blockquote>
          <h2>把观察带回创作</h2>
          <p>无论使用摄影机、剪辑软件，还是新一代生成式工具，技术最终都应该服务于感受。先确认自己真正想表达什么，再决定使用哪一种方法。工具越强，克制与选择就越重要。</p>
          <p>这篇笔记仍会随着新的观看和实践继续更新。数字花园的意义，不是把每篇文字封存成结论，而是允许想法保留生长的余地。</p>
        </div>
      </article>
      <section v-if="related.length" class="related"><span class="eyebrow">Related Notes</span><h2>继续阅读</h2><router-link v-for="item in related" :key="item.id" :to="`/articles/${item.id}`"><span>{{ item.title }}<small>{{ item.summary }}</small></span><ArrowRight :size="17" /></router-link></section>
    </div>
  </div>
</template>

<style scoped>
.article-page{padding:55px 0 110px;background:linear-gradient(180deg,rgba(217,173,95,.035),transparent 28%)}.article-shell{width:min(860px,calc(100% - 40px));margin:auto}.back-link{display:inline-flex;align-items:center;gap:7px;margin-bottom:60px;color:var(--color-text-tertiary);font-size:.8rem;text-decoration:none}article header{text-align:center}article h1{margin:22px auto;max-width:820px;font-family:var(--font-display);font-size:clamp(2.8rem,7vw,5.5rem);font-weight:500;letter-spacing:-.045em;line-height:1.16}.lead{max-width:650px;margin:auto;color:var(--color-text-secondary);font-family:var(--font-display);font-size:1.15rem;line-height:1.8}.meta{display:flex;justify-content:center;gap:20px;margin-top:28px;color:var(--color-text-tertiary);font-size:.75rem}.meta span{display:flex;align-items:center;gap:6px}.cover{position:relative;aspect-ratio:16/7;margin:64px -80px;background:radial-gradient(circle at 70% 40%,color-mix(in srgb,var(--cover) 80%,white),transparent 20%),linear-gradient(130deg,#0b0c11,var(--cover),#17181e);border-radius:4px;overflow:hidden}.cover::after{position:absolute;inset:0;background:repeating-linear-gradient(90deg,transparent 0 80px,rgba(255,255,255,.025) 81px);content:''}.cover span{position:absolute;right:24px;bottom:20px;color:rgba(255,255,255,.4);font-size:.65rem;letter-spacing:.2em}.prose{width:min(680px,100%);margin:auto}.prose p{margin:0 0 28px;color:#c2beb5;font-family:var(--font-display);font-size:1.06rem;line-height:2}.prose h2{margin:64px 0 22px;font-family:var(--font-display);font-size:1.8rem;font-weight:500}.prose blockquote{margin:55px 0;padding:10px 0 10px 30px;border-left:2px solid var(--color-primary);font-family:var(--font-display);font-size:1.45rem;line-height:1.7;color:var(--color-text-primary)}.related{margin-top:100px;padding-top:50px;border-top:1px solid var(--color-border-subtle)}.related>h2{margin:13px 0 25px;font-family:var(--font-display);font-size:2rem}.related>a{display:flex;align-items:center;justify-content:space-between;padding:20px 0;border-bottom:1px solid var(--color-border-subtle);color:var(--color-text-primary);text-decoration:none}.related>a>span{display:grid;gap:5px}.related small{max-width:650px;color:var(--color-text-tertiary);font-weight:400}.related a:hover{color:var(--color-primary)}
@media(max-width:900px){.cover{margin:50px 0}}@media(max-width:600px){.meta{flex-wrap:wrap}.cover{aspect-ratio:4/3}.article-page{padding-top:35px}}
</style>

