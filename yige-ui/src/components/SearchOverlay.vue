<script setup>
import { computed, nextTick, onBeforeUnmount, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { BookOpen, Bot, Clapperboard, GraduationCap, Search, X } from '@lucide/vue'
import { searchApi } from '@/api'
import UiState from '@/components/UiState.vue'

const props = defineProps({ open: Boolean })
const emit = defineEmits(['close'])
const router = useRouter()
const input = ref(null)
const query = ref('')
const loading = ref(false)
const error = ref('')
const results = ref(null)
let timer

const sections = computed(() => results.value ? [
  { key: 'movies', title: '影片', icon: Clapperboard, items: results.value.movies || [], label: item => item.title, detail: item => `${item.year || '年份未知'} · ${item.genres || '未分类'}`, to: item => `/movies/${item.id}` },
  { key: 'articles', title: '文章', icon: BookOpen, items: results.value.articles || [], label: item => item.title, detail: item => item.summary || item.tag || '', to: item => `/articles/${item.id}` },
  { key: 'courses', title: '课程', icon: GraduationCap, items: results.value.courses || [], label: item => item.title, detail: item => `${item.category || '课程'} · ${item.level || '通用'}`, to: () => '/learning' },
  { key: 'tools', title: 'AI 工具', icon: Bot, items: results.value.tools || [], label: item => item.name, detail: item => item.description || '', to: () => '/ai-share' },
].filter(section => section.items.length) : [])

const hasResults = computed(() => sections.value.length > 0)

watch(() => props.open, async (open) => {
  document.body.style.overflow = open ? 'hidden' : ''
  if (open) await nextTick(() => input.value?.focus())
})

watch(query, (value) => {
  clearTimeout(timer)
  error.value = ''
  if (value.trim().length < 2) {
    results.value = null
    loading.value = false
    return
  }
  timer = setTimeout(() => performSearch(value), 260)
})

async function performSearch(value) {
  loading.value = true
  try {
    const response = await searchApi.search(value.trim())
    if (query.value.trim() === value.trim()) results.value = response.data
  } catch (searchError) {
    error.value = searchError.message
  } finally {
    loading.value = false
  }
}

function navigate(path) {
  router.push(path)
  emit('close')
}

function onKeydown(event) {
  if (event.key === 'Escape' && props.open) emit('close')
}

window.addEventListener('keydown', onKeydown)
onBeforeUnmount(() => {
  clearTimeout(timer)
  document.body.style.overflow = ''
  window.removeEventListener('keydown', onKeydown)
})
</script>

<template>
  <Transition name="search-fade">
    <div v-if="open" class="search-overlay" role="dialog" aria-modal="true" aria-label="全站搜索" @click.self="$emit('close')">
      <div class="search-panel">
        <div class="search-panel__bar">
          <Search :size="21" />
          <input ref="input" v-model="query" type="search" placeholder="搜索影片、文章、课程或 AI 工具" aria-label="搜索关键词" />
          <span class="search-panel__hint">ESC</span>
          <button type="button" aria-label="关闭搜索" @click="$emit('close')"><X :size="20" /></button>
        </div>
        <div class="search-panel__body">
          <p v-if="query.trim().length < 2" class="search-intro">输入至少两个字符。试试“诺兰”“剪辑”或“AI”。</p>
          <UiState v-else-if="loading" compact type="loading" title="正在检索片库与笔记…" />
          <UiState v-else-if="error" compact type="error" :message="error" @retry="performSearch(query)" />
          <UiState v-else-if="results && !hasResults" compact type="empty" title="没有找到相关内容" message="换一个更简短的关键词试试。" />
          <div v-else class="search-results">
            <section v-for="section in sections" :key="section.key" class="search-group">
              <h2><component :is="section.icon" :size="17" />{{ section.title }}</h2>
              <button v-for="item in section.items" :key="item.id" type="button" class="search-result" @click="navigate(section.to(item))">
                <span><strong>{{ section.label(item) }}</strong><small>{{ section.detail(item) }}</small></span>
                <span aria-hidden="true">↗</span>
              </button>
            </section>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.search-overlay{position:fixed;inset:0;z-index:100;display:flex;align-items:flex-start;justify-content:center;padding:9vh 20px 20px;background:rgba(4,5,8,.82);backdrop-filter:blur(14px)}
.search-panel{width:min(760px,100%);max-height:78vh;overflow:hidden;border:1px solid var(--color-border-strong);border-radius:24px;background:rgba(18,19,24,.98);box-shadow:0 30px 90px rgba(0,0,0,.65)}
.search-panel__bar{display:flex;align-items:center;gap:13px;padding:18px 20px;border-bottom:1px solid var(--color-border-subtle);color:var(--color-primary)}
.search-panel__bar input{flex:1;min-width:0;border:0;outline:0;background:transparent;color:var(--color-text-primary);font-size:1.05rem}.search-panel__bar input::placeholder{color:var(--color-text-tertiary)}
.search-panel__bar button{display:grid;place-items:center;border:0;background:transparent;color:var(--color-text-secondary);cursor:pointer}.search-panel__hint{padding:4px 7px;border:1px solid var(--color-border-default);border-radius:6px;color:var(--color-text-tertiary);font-size:.67rem}
.search-panel__body{max-height:calc(78vh - 64px);overflow:auto;padding:20px}.search-intro{padding:36px 14px;text-align:center;color:var(--color-text-tertiary)}
.search-results{display:grid;gap:24px}.search-group h2{display:flex;align-items:center;gap:8px;margin-bottom:9px;color:var(--color-primary);font-size:.78rem;letter-spacing:.12em;text-transform:uppercase}
.search-result{width:100%;display:flex;align-items:center;justify-content:space-between;gap:20px;padding:13px 12px;border:0;border-radius:12px;background:transparent;color:var(--color-text-secondary);text-align:left;cursor:pointer;transition:.18s ease}.search-result:hover{background:var(--color-bg-muted);color:var(--color-primary)}
.search-result span:first-child{min-width:0;display:grid;gap:4px}.search-result strong{color:var(--color-text-primary);font-size:.92rem}.search-result small{overflow:hidden;color:var(--color-text-tertiary);font-size:.78rem;text-overflow:ellipsis;white-space:nowrap}
.search-fade-enter-active,.search-fade-leave-active{transition:opacity .2s ease}.search-fade-enter-from,.search-fade-leave-to{opacity:0}
@media(max-width:600px){.search-overlay{padding:0}.search-panel{height:100%;max-height:none;border:0;border-radius:0}.search-panel__body{max-height:calc(100vh - 64px)}.search-panel__hint{display:none}}
</style>
