<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { Menu, Search, X } from '@lucide/vue'
import SearchOverlay from '@/components/SearchOverlay.vue'

const route = useRoute()
const mobileOpen = ref(false)
const searchOpen = ref(false)

const navItems = [
  { name: 'index', label: '首页', path: '/' },
  { name: 'movie-recommend', label: '片单', path: '/movie-recommend' },
  { name: 'learning', label: '学习', path: '/learning' },
  { name: 'ai-share', label: 'AI 实验室', path: '/ai-share' },
  { name: 'about', label: '关于', path: '/about' },
]

function isActive(name) {
  if (name === 'movie-recommend' && route.name === 'movie-detail') return true
  return route.name === name
}

function openSearch() {
  searchOpen.value = true
  mobileOpen.value = false
}

watch(() => route.fullPath, () => {
  mobileOpen.value = false
})
</script>

<template>
  <header class="site-header">
    <nav class="site-nav" aria-label="主导航">
      <router-link to="/" class="brand" aria-label="CineVerse 首页">
        <span class="brand-mark" aria-hidden="true"><i></i><i></i></span>
        <span><strong>CineVerse</strong><small>一个影视 · 影像与创造</small></span>
      </router-link>

      <div class="nav-links">
        <router-link v-for="item in navItems" :key="item.name" :to="item.path" :class="{ active: isActive(item.name) }">
          {{ item.label }}
        </router-link>
      </div>

      <div class="nav-actions">
        <button class="search-trigger" type="button" aria-label="打开全站搜索" @click="openSearch">
          <Search :size="17" /><span>搜索</span><kbd>⌘ K</kbd>
        </button>
        <button class="menu-trigger" type="button" :aria-expanded="mobileOpen" aria-label="打开菜单" @click="mobileOpen = !mobileOpen">
          <X v-if="mobileOpen" :size="22" /><Menu v-else :size="22" />
        </button>
      </div>
    </nav>

    <Transition name="drawer">
      <div v-if="mobileOpen" class="mobile-nav">
        <router-link v-for="item in navItems" :key="item.name" :to="item.path" :class="{ active: isActive(item.name) }">{{ item.label }}</router-link>
        <button type="button" @click="openSearch"><Search :size="17" /> 搜索全站内容</button>
      </div>
    </Transition>
  </header>
  <SearchOverlay :open="searchOpen" @close="searchOpen = false" />
</template>

<style scoped>
.site-header{position:fixed;inset:0 0 auto;z-index:60;border-bottom:1px solid rgba(255,255,255,.07);background:rgba(9,10,14,.84);backdrop-filter:blur(18px) saturate(130%)}
.site-nav{width:min(var(--container-max),calc(100% - 40px));height:76px;margin:auto;display:flex;align-items:center;justify-content:space-between;gap:28px}.brand{display:flex;align-items:center;gap:12px;color:inherit;text-decoration:none}.brand>span:last-child{display:grid;gap:1px}.brand strong{font-family:var(--font-display);font-size:1.05rem;letter-spacing:.03em;color:var(--color-text-primary)}.brand small{font-size:.63rem;letter-spacing:.14em;color:var(--color-text-tertiary)}
.brand-mark{position:relative;width:34px;height:34px;display:grid;place-items:center;border:1px solid rgba(212,168,83,.4);border-radius:50%;background:rgba(212,168,83,.08)}.brand-mark i{position:absolute;width:5px;height:5px;border:1px solid var(--color-primary);border-radius:50%}.brand-mark i:first-child{transform:translate(-5px,-3px)}.brand-mark i:last-child{transform:translate(5px,3px)}
.nav-links{display:flex;align-items:center;gap:2px}.nav-links a{position:relative;padding:10px 13px;color:var(--color-text-secondary);font-size:.84rem;text-decoration:none;transition:.18s ease}.nav-links a::after{position:absolute;right:13px;bottom:4px;left:13px;height:1px;background:var(--color-primary);content:'';transform:scaleX(0);transition:.18s ease}.nav-links a:hover,.nav-links a.active{color:var(--color-text-primary)}.nav-links a.active::after{transform:scaleX(1)}
.nav-actions{display:flex;align-items:center}.search-trigger{display:flex;align-items:center;gap:8px;padding:8px 10px;border:1px solid var(--color-border-default);border-radius:11px;background:rgba(255,255,255,.025);color:var(--color-text-secondary);cursor:pointer}.search-trigger span{font-size:.76rem}.search-trigger kbd{padding:2px 5px;border:1px solid var(--color-border-default);border-radius:5px;color:var(--color-text-tertiary);font:600 .6rem var(--font-body)}.menu-trigger{display:none;border:0;background:none;color:var(--color-text-primary)}
.mobile-nav{display:none}.drawer-enter-active,.drawer-leave-active{transition:.2s ease}.drawer-enter-from,.drawer-leave-to{opacity:0;transform:translateY(-8px)}
@media(max-width:850px){.site-nav{height:68px}.nav-links,.search-trigger{display:none}.menu-trigger{display:grid;place-items:center}.mobile-nav{display:grid;padding:10px 20px 20px;border-top:1px solid var(--color-border-subtle);background:var(--color-bg-elevated)}.mobile-nav a,.mobile-nav button{display:flex;align-items:center;gap:9px;padding:14px 12px;border:0;border-bottom:1px solid var(--color-border-subtle);background:none;color:var(--color-text-secondary);text-decoration:none;text-align:left}.mobile-nav a.active{color:var(--color-primary)}}
@media(max-width:430px){.site-nav{width:calc(100% - 28px)}.brand small{display:none}}
</style>
