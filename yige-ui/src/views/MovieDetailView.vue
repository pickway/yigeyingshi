<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ArrowLeft, ArrowRight, Clock3, Film, Star, UserRound } from '@lucide/vue'
import { movieApi } from '@/api'
import UiState from '@/components/UiState.vue'
import { splitList } from '@/utils/content'

const route = useRoute()
const movie = ref(null)
const related = ref([])
const loading = ref(true)
const error = ref('')

const genres = computed(() => splitList(movie.value?.genres))
const gradient = computed(() => {
  const color = movie.value?.posterColor || '#293149'
  return `radial-gradient(circle at 70% 25%, ${color}ee, transparent 35%), linear-gradient(145deg, ${color}, #090a0e 72%)`
})

async function load() {
  loading.value = true
  error.value = ''
  try {
    const response = await movieApi.detail(route.params.id)
    movie.value = response.data
    document.title = `${movie.value.title} - CineVerse`
    const relatedResponse = await movieApi.list({ genre: genres.value[0], pageSize: 4 })
    related.value = (relatedResponse.data || []).filter(item => item.id !== movie.value.id).slice(0, 3)
  } catch (loadError) {
    error.value = loadError.message
  } finally {
    loading.value = false
  }
}

onMounted(load)
watch(() => route.params.id, load)
</script>

<template>
  <div class="detail-page">
    <div class="content-shell detail-shell">
      <router-link to="/movie-recommend" class="back-link"><ArrowLeft :size="16" /> 返回片单</router-link>
      <UiState v-if="loading" type="loading" title="正在准备放映…" />
      <UiState v-else-if="error" type="error" :message="error" @retry="load" />
      <template v-else-if="movie">
        <section class="movie-hero">
          <div class="movie-poster" :style="{ background: gradient }">
            <Film :size="60" stroke-width="1" />
            <span>{{ movie.year }}</span>
          </div>
          <div class="movie-copy">
            <span class="eyebrow">CineVerse Selection</span>
            <h1 class="display-title">{{ movie.title }}</h1>
            <p v-if="movie.originalTitle" class="original-title">{{ movie.originalTitle }}</p>
            <div class="movie-meta">
              <span class="rating"><Star :size="17" fill="currentColor" /> {{ movie.rating.toFixed(1) }}</span>
              <span>{{ movie.year }}</span><span v-if="movie.duration"><Clock3 :size="15" /> {{ movie.duration }} 分钟</span>
            </div>
            <div class="genre-row"><span v-for="genre in genres" :key="genre">{{ genre }}</span></div>
            <p class="movie-description">{{ movie.description || '一部值得放慢速度、留出时间观看的作品。' }}</p>
            <div class="director"><UserRound :size="18" /><span><small>导演</small><strong>{{ movie.director || '未知' }}</strong></span></div>
          </div>
        </section>

        <section class="viewing-note">
          <span class="eyebrow">Viewing Note</span>
          <h2 class="section-title">为什么值得看</h2>
          <div class="note-grid">
            <p>我更愿意把电影当作一次完整的感官经验：画面、声音和时间共同改变我们对故事的理解。这部作品最迷人的地方，是它在类型外壳之下保留了真实的人物温度。</p>
            <blockquote>“好的电影不会替你回答问题，它会让那个问题在散场后继续发光。”</blockquote>
          </div>
        </section>

        <section v-if="related.length" class="related-section">
          <div class="section-heading"><div><span class="eyebrow">Keep Watching</span><h2 class="section-title">相似气质</h2></div><router-link to="/movie-recommend">完整片单 <ArrowRight :size="15" /></router-link></div>
          <div class="related-grid">
            <router-link v-for="item in related" :key="item.id" :to="`/movies/${item.id}`" class="related-card surface-card">
              <div :style="{ background: `linear-gradient(145deg, ${item.posterColor || '#303747'}, #101116)` }"><Film :size="28" /></div>
              <span><strong>{{ item.title }}</strong><small>{{ item.year }} · {{ item.rating.toFixed(1) }}</small></span>
            </router-link>
          </div>
        </section>
      </template>
    </div>
  </div>
</template>

<style scoped>
.detail-page{padding:54px 0 110px;background:radial-gradient(circle at 80% 0,rgba(217,173,95,.06),transparent 30%)}.back-link{display:inline-flex;align-items:center;gap:8px;margin-bottom:32px;color:var(--color-text-tertiary);font-size:.8rem;text-decoration:none}.back-link:hover{color:var(--color-primary)}
.movie-hero{display:grid;grid-template-columns:minmax(260px,390px) 1fr;gap:clamp(42px,8vw,100px);align-items:center}.movie-poster{position:relative;aspect-ratio:2/3;display:grid;place-items:center;border:1px solid rgba(255,255,255,.12);border-radius:8px;color:rgba(255,255,255,.55);box-shadow:0 35px 80px rgba(0,0,0,.5);overflow:hidden}.movie-poster::after{position:absolute;inset:0;background:linear-gradient(90deg,transparent 49.7%,rgba(255,255,255,.04) 50%,transparent 50.3%);content:''}.movie-poster span{position:absolute;right:20px;bottom:18px;font-family:var(--font-display);font-size:4rem;color:rgba(255,255,255,.06)}
.movie-copy h1{margin:22px 0 5px;font-size:clamp(3.2rem,7vw,6.5rem)}.original-title{margin:0 0 25px;color:var(--color-text-tertiary);font-family:var(--font-display);font-style:italic}.movie-meta{display:flex;align-items:center;gap:18px;color:var(--color-text-secondary);font-size:.84rem}.movie-meta span{display:flex;align-items:center;gap:6px}.movie-meta .rating{color:var(--color-primary)}.genre-row{display:flex;flex-wrap:wrap;gap:8px;margin:24px 0}.genre-row span{padding:6px 11px;border:1px solid var(--color-border-default);border-radius:999px;color:var(--color-text-secondary);font-size:.72rem}.movie-description{max-width:670px;color:var(--color-text-secondary);font-family:var(--font-display);font-size:1.2rem;line-height:1.9}.director{display:flex;align-items:center;gap:12px;margin-top:28px;color:var(--color-primary)}.director span{display:grid}.director small{color:var(--color-text-tertiary);font-size:.65rem}.director strong{color:var(--color-text-primary);font-size:.88rem}
.viewing-note{margin-top:120px;padding:64px;border:1px solid var(--color-border-subtle);border-radius:24px;background:linear-gradient(145deg,rgba(255,255,255,.025),rgba(217,173,95,.025))}.viewing-note h2{margin:15px 0 38px}.note-grid{display:grid;grid-template-columns:1fr 1fr;gap:70px}.note-grid p,.note-grid blockquote{margin:0;color:var(--color-text-secondary);line-height:1.9}.note-grid blockquote{padding-left:25px;border-left:1px solid var(--color-primary);font-family:var(--font-display);font-size:1.25rem;color:var(--color-text-primary)}
.related-section{margin-top:100px}.section-heading{display:flex;align-items:end;justify-content:space-between;margin-bottom:30px}.section-heading h2{margin:12px 0 0}.section-heading>a{display:flex;align-items:center;gap:6px;color:var(--color-primary);font-size:.8rem;text-decoration:none}.related-grid{display:grid;grid-template-columns:repeat(3,1fr);gap:18px}.related-card{display:flex;align-items:center;gap:16px;padding:12px;color:inherit;text-decoration:none}.related-card>div{width:72px;aspect-ratio:2/3;display:grid;place-items:center;border-radius:8px;color:rgba(255,255,255,.45)}.related-card>span{display:grid;gap:6px}.related-card small{color:var(--color-text-tertiary)}
@media(max-width:800px){.movie-hero{grid-template-columns:1fr}.movie-poster{width:min(340px,85vw);margin:auto}.movie-copy{text-align:center}.movie-meta,.genre-row,.director{justify-content:center}.note-grid{grid-template-columns:1fr;gap:35px}.viewing-note{padding:36px 24px}.related-grid{grid-template-columns:1fr}}
</style>

