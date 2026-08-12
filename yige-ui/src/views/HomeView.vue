<script setup>
import { computed, onMounted, ref } from 'vue'
import { ArrowRight, BookOpen, Clapperboard, Play, Sparkles, Star } from '@lucide/vue'
import { aiApi, articleApi, movieApi } from '@/api'
import UiState from '@/components/UiState.vue'
import { formatChineseDate, normalizeExternalUrl, splitList } from '@/utils/content'

const featuredMovies = ref([])
const latestPosts = ref([])
const aiHighlights = ref([])
const loading = ref(true)
const error = ref('')

const leadMovie = computed(() => featuredMovies.value[0] || null)
const secondaryMovies = computed(() => featuredMovies.value.slice(1, 3))

function movieGradient(movie) {
  const color = movie?.posterColor || '#263049'
  return `radial-gradient(circle at 70% 20%, ${color}dd, transparent 30%), linear-gradient(150deg, ${color}, #0b0c11 75%)`
}

async function load() {
  loading.value = true
  error.value = ''
  try {
    const [movies, articles, tools] = await Promise.all([
      movieApi.featured(), articleApi.latest(3), aiApi.featuredTools(),
    ])
    featuredMovies.value = movies.data || []
    latestPosts.value = articles.data || []
    aiHighlights.value = tools.data || []
  } catch (loadError) {
    error.value = loadError.message
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>

<template>
  <div class="home-page">
    <section class="hero">
      <img src="/cineverse-hero.png" alt="暖金色放映室与影像工作台" class="hero-image" />
      <div class="hero-shade"></div>
      <div class="hero-content content-shell">
        <span class="eyebrow">Cinema · Learning · Creative AI</span>
        <h1 class="display-title">在光影之间，<br /><em>保持好奇。</em></h1>
        <p>一个影视创作者的数字花园。收藏值得重看的电影，整理真正有用的学习路径，也记录 AI 时代的新创作方式。</p>
        <div class="hero-actions">
          <router-link to="/movie-recommend" class="btn-primary"><Play :size="15" fill="currentColor" />开始探索</router-link>
          <router-link to="/about" class="btn-secondary">认识我 <ArrowRight :size="15" /></router-link>
        </div>
        <div class="hero-index"><span>ISSUE 01</span><i></i><span>持续更新</span></div>
      </div>
      <a href="#weekly" class="scroll-cue" aria-label="向下阅读"><span>SCROLL</span><i></i></a>
    </section>

    <section id="weekly" class="weekly content-shell">
      <div class="section-heading">
        <div><span class="eyebrow">This Week</span><h2 class="section-title">本周放映</h2></div>
        <p>不是排行榜，只是最近仍在脑海里回响的作品。</p>
      </div>
      <UiState v-if="loading" type="loading" />
      <UiState v-else-if="error" type="error" :message="error" @retry="load" />
      <template v-else-if="leadMovie">
        <router-link :to="`/movies/${leadMovie.id}`" class="lead-movie">
          <div class="lead-visual" :style="{ background: movieGradient(leadMovie) }"><span>01</span><Clapperboard :size="54" stroke-width="1" /></div>
          <div class="lead-copy">
            <span class="film-label">EDITOR'S CHOICE</span>
            <h3>{{ leadMovie.title }}</h3>
            <p>{{ leadMovie.description }}</p>
            <div class="film-meta"><span><Star :size="15" fill="currentColor" /> {{ leadMovie.rating.toFixed(1) }}</span><span>{{ leadMovie.year }}</span><span>{{ splitList(leadMovie.genres).join(' / ') }}</span></div>
            <strong>进入放映 <ArrowRight :size="16" /></strong>
          </div>
        </router-link>
        <div class="film-row">
          <router-link v-for="(movie, index) in secondaryMovies" :key="movie.id" :to="`/movies/${movie.id}`" class="film-mini surface-card">
            <div :style="{ background: movieGradient(movie) }"><span>0{{ index + 2 }}</span></div>
            <section><small>{{ movie.year }} · {{ splitList(movie.genres)[0] }}</small><h3>{{ movie.title }}</h3><p>{{ movie.description }}</p><b><Star :size="13" fill="currentColor" />{{ movie.rating.toFixed(1) }}</b></section>
          </router-link>
        </div>
      </template>
    </section>

    <section class="journal">
      <div class="content-shell">
        <div class="section-heading light"><div><span class="eyebrow">Notes & Essays</span><h2 class="section-title">最近写下</h2></div><router-link to="/about">关于这个数字花园 <ArrowRight :size="15" /></router-link></div>
        <div class="journal-layout">
          <router-link v-if="latestPosts[0]" :to="`/articles/${latestPosts[0].id}`" class="featured-post">
            <div class="post-art" :style="{ '--cover': latestPosts[0].coverColor || '#25365c' }"><BookOpen :size="34" /><span>FEATURED NOTE</span></div>
            <span class="post-tag">{{ latestPosts[0].tag }}</span><h3>{{ latestPosts[0].title }}</h3><p>{{ latestPosts[0].summary }}</p><small>{{ formatChineseDate(latestPosts[0].publishedAt) }} · 6 分钟阅读</small>
          </router-link>
          <div class="post-list">
            <router-link v-for="(post, index) in latestPosts.slice(1)" :key="post.id" :to="`/articles/${post.id}`">
              <span>0{{ index + 2 }}</span><div><small>{{ post.tag }} · {{ formatChineseDate(post.publishedAt) }}</small><h3>{{ post.title }}</h3><p>{{ post.summary }}</p></div><ArrowRight :size="18" />
            </router-link>
            <router-link to="/learning" class="learning-callout"><span><BookOpen :size="22" /></span><div><small>LEARNING PATH</small><h3>把兴趣变成可以完成的学习路线</h3></div><ArrowRight :size="18" /></router-link>
          </div>
        </div>
      </div>
    </section>

    <section class="lab content-shell">
      <div class="lab-intro"><span class="eyebrow">Creative AI Lab</span><h2 class="section-title">工具会更新，<br />判断力不会过时。</h2><p>这里不做工具堆砌。每一个推荐，都先放进真实的影像工作流里试一遍。</p><router-link to="/ai-share" class="btn-secondary">进入 AI 实验室 <ArrowRight :size="15" /></router-link></div>
      <div class="tool-stack">
        <a v-for="(tool, index) in aiHighlights" :key="tool.id" :href="normalizeExternalUrl(tool.url)" target="_blank" rel="noopener noreferrer" class="tool-line">
          <span>0{{ index + 1 }}</span><div><small>{{ tool.category }}</small><strong>{{ tool.name }}</strong><p>{{ tool.description }}</p></div><Sparkles :size="20" />
        </a>
      </div>
    </section>

    <section class="closing"><div class="content-shell"><span>KEEP WATCHING</span><h2>下一次灵感，<br />也许就在下一帧。</h2><router-link to="/movie-recommend" class="btn-primary">打开完整片单 <ArrowRight :size="16" /></router-link></div></section>
  </div>
</template>

<style scoped>
.hero{position:relative;min-height:min(820px,calc(100vh - 20px));display:flex;align-items:center;overflow:hidden;background:#07080b}.hero-image{position:absolute;inset:0;width:100%;height:100%;object-fit:cover;object-position:center}.hero-shade{position:absolute;inset:0;background:linear-gradient(90deg,rgba(6,7,10,.98) 0%,rgba(6,7,10,.9) 33%,rgba(6,7,10,.24) 68%,rgba(6,7,10,.2)),linear-gradient(0deg,rgba(6,7,10,.82),transparent 45%)}.hero-content{position:relative;z-index:2;padding-block:90px}.hero h1{margin:25px 0 24px;max-width:760px}.hero h1 em{color:var(--color-primary);font-weight:500}.hero-content>p{max-width:590px;color:#b8b3aa;font-family:var(--font-display);font-size:1.1rem;line-height:1.85}.hero-actions{display:flex;gap:12px;margin-top:35px}.hero-index{display:flex;align-items:center;gap:12px;margin-top:70px;color:var(--color-text-tertiary);font-size:.62rem;letter-spacing:.15em}.hero-index i{width:60px;height:1px;background:var(--color-border-strong)}.scroll-cue{position:absolute;right:35px;bottom:35px;z-index:2;display:flex;align-items:center;gap:12px;color:var(--color-text-tertiary);font-size:.58rem;letter-spacing:.2em;text-decoration:none;transform:rotate(90deg);transform-origin:right}.scroll-cue i{width:45px;height:1px;background:var(--color-text-tertiary)}
.weekly{padding-block:120px}.section-heading{display:flex;align-items:end;justify-content:space-between;margin-bottom:50px}.section-heading h2{margin:13px 0 0}.section-heading>p{max-width:360px;color:var(--color-text-tertiary);line-height:1.7}.lead-movie{display:grid;grid-template-columns:1.1fr 1fr;min-height:470px;border:1px solid var(--color-border-subtle);color:inherit;text-decoration:none}.lead-visual{position:relative;display:grid;place-items:center;color:rgba(255,255,255,.42);overflow:hidden}.lead-visual span{position:absolute;top:-50px;left:25px;font-family:var(--font-display);font-size:13rem;color:rgba(255,255,255,.045)}.lead-copy{display:flex;flex-direction:column;justify-content:center;padding:70px}.film-label,.post-tag{color:var(--color-primary);font-size:.65rem;letter-spacing:.17em}.lead-copy h3{margin:16px 0;font-family:var(--font-display);font-size:clamp(2.4rem,5vw,4.7rem);font-weight:500}.lead-copy>p{color:var(--color-text-secondary);line-height:1.8}.film-meta{display:flex;flex-wrap:wrap;gap:14px;margin:22px 0;color:var(--color-text-tertiary);font-size:.73rem}.film-meta span{display:flex;align-items:center;gap:5px}.film-meta span:first-child{color:var(--color-primary)}.lead-copy>strong{display:flex;align-items:center;gap:8px;margin-top:12px;color:var(--color-primary);font-size:.78rem}.film-row{display:grid;grid-template-columns:1fr 1fr;gap:18px;margin-top:18px}.film-mini{display:grid;grid-template-columns:150px 1fr;color:inherit;text-decoration:none;overflow:hidden}.film-mini>div{position:relative;min-height:210px}.film-mini>div span{position:absolute;right:12px;bottom:5px;font-family:var(--font-display);font-size:3rem;color:rgba(255,255,255,.08)}.film-mini section{padding:30px}.film-mini small{color:var(--color-text-tertiary);font-size:.68rem}.film-mini h3{margin:8px 0;font-family:var(--font-display);font-size:1.5rem}.film-mini p{color:var(--color-text-tertiary);font-size:.78rem;line-height:1.6}.film-mini b{display:flex;align-items:center;gap:4px;color:var(--color-primary);font-size:.72rem}
.journal{padding-block:120px;background:#eee8dc;color:#171614}.section-heading.light .eyebrow{color:#8b6327}.section-heading.light>a{display:flex;align-items:center;gap:7px;color:#675e52;font-size:.75rem;text-decoration:none}.journal-layout{display:grid;grid-template-columns:1fr 1.1fr;gap:80px}.featured-post{color:inherit;text-decoration:none}.post-art{position:relative;aspect-ratio:16/10;display:grid;place-items:center;margin-bottom:30px;background:radial-gradient(circle at 70% 30%,var(--cover),transparent 25%),linear-gradient(135deg,#15171d,var(--cover));color:rgba(255,255,255,.6)}.post-art span{position:absolute;right:15px;bottom:12px;color:rgba(255,255,255,.38);font-size:.58rem;letter-spacing:.17em}.featured-post h3{margin:12px 0;font-family:var(--font-display);font-size:2rem;font-weight:600}.featured-post>p{color:#696157;line-height:1.75}.featured-post>small{color:#8b8277}.post-list>a{display:grid;grid-template-columns:45px 1fr auto;gap:20px;align-items:center;padding:27px 0;border-bottom:1px solid #d2c9bb;color:inherit;text-decoration:none}.post-list>a>span{font-family:var(--font-display);font-size:1.4rem;color:#9e8b6c}.post-list small{color:#8b6327;font-size:.64rem;letter-spacing:.1em}.post-list h3{margin:7px 0;font-family:var(--font-display);font-size:1.35rem}.post-list p{margin:0;color:#766e64;font-size:.8rem;line-height:1.6}.learning-callout>span{display:grid!important;place-items:center;width:38px;height:38px;border:1px solid #b7aa98;border-radius:50%;font-size:inherit!important;color:#8b6327!important}
.lab{display:grid;grid-template-columns:.85fr 1.1fr;gap:100px;padding-block:130px}.lab-intro h2{margin:20px 0}.lab-intro>p{max-width:450px;color:var(--color-text-tertiary);line-height:1.8}.lab-intro .btn-secondary{margin-top:25px}.tool-line{display:grid;grid-template-columns:45px 1fr auto;gap:20px;align-items:center;padding:27px 0;border-bottom:1px solid var(--color-border-subtle);color:inherit;text-decoration:none}.tool-line>span{font-family:var(--font-display);color:var(--color-primary)}.tool-line div{display:grid;gap:4px}.tool-line small{color:var(--color-primary);font-size:.6rem;letter-spacing:.14em;text-transform:uppercase}.tool-line strong{font-family:var(--font-display);font-size:1.25rem}.tool-line p{margin:0;color:var(--color-text-tertiary);font-size:.77rem;line-height:1.5}.tool-line>svg{color:var(--color-text-tertiary)}.tool-line:hover>svg{color:var(--color-primary)}
.closing{padding:120px 0;background:linear-gradient(120deg,#17140f,#0e1015);text-align:center}.closing span{color:var(--color-primary);font-size:.65rem;letter-spacing:.22em}.closing h2{margin:20px 0 35px;font-family:var(--font-display);font-size:clamp(2.8rem,6vw,5.7rem);font-weight:500;line-height:1.15}.closing .btn-primary{display:inline-flex}
@media(max-width:850px){.hero{min-height:760px}.hero-shade{background:linear-gradient(90deg,rgba(6,7,10,.95),rgba(6,7,10,.45)),linear-gradient(0deg,rgba(6,7,10,.9),transparent)}.lead-movie,.journal-layout,.lab{grid-template-columns:1fr}.lead-visual{min-height:400px}.lead-copy{padding:45px}.journal-layout,.lab{gap:55px}.film-mini{grid-template-columns:120px 1fr}}
@media(max-width:620px){.hero-content{padding-block:70px}.hero-image{object-position:62% center}.hero-actions{align-items:stretch;flex-direction:column;width:190px}.scroll-cue{display:none}.weekly,.lab{padding-block:80px}.section-heading{align-items:flex-start;flex-direction:column;gap:18px}.film-row{grid-template-columns:1fr}.lead-visual{min-height:300px}.lead-copy{padding:32px 24px}.journal{padding-block:80px}.post-list>a{grid-template-columns:35px 1fr}.post-list>a>svg{display:none}.film-mini{grid-template-columns:105px 1fr}.film-mini section{padding:22px 18px}}
</style>
