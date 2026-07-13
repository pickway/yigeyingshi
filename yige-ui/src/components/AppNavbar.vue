<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { Search, Menu, X } from '@lucide/vue'

const route = useRoute()
const mobileOpen = ref(false)

const navItems = [
  { name: 'index', label: '首页', path: '/' },
  { name: 'movie-recommend', label: '影视推荐', path: '/movie-recommend' },
  { name: 'learning', label: '学习资料', path: '/learning' },
  { name: 'ai-share', label: 'AI 分享', path: '/ai-share' },
]

function isActive(name) {
  return route.name === name
}

function toggleMobile() {
  mobileOpen.value = !mobileOpen.value
}
</script>

<template>
  <header
    class="fixed top-0 left-0 w-full z-50"
    style="background: var(--color-bg-elevated); border-bottom: 1px solid var(--color-border-subtle)"
  >
    <nav
      class="mx-auto flex items-center justify-between"
      style="max-width: var(--container-max); height: 64px; padding: 0 var(--content-padding)"
    >
      <!-- Logo -->
      <router-link
        to="/"
        class="shrink-0 cine-display no-underline"
        style="color: var(--color-primary); font-size: 1.25rem"
      >
        CineVerse
      </router-link>

      <!-- Desktop Nav Links -->
      <div class="hidden md:flex items-center gap-1">
        <router-link
          v-for="item in navItems"
          :key="item.name"
          :to="item.path"
          class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-150 whitespace-nowrap no-underline"
          :style="
            isActive(item.name)
              ? 'color: var(--color-primary); background: var(--color-primary-tint-1)'
              : 'color: var(--color-text-secondary)'
          "
        >
          {{ item.label }}
        </router-link>
      </div>

      <!-- Right: Search + Mobile -->
      <div class="flex items-center gap-3">
        <!-- Search pill -->
        <div
          class="hidden md:flex items-center gap-2 px-4 py-2 rounded-full transition-colors duration-150 cursor-pointer"
          style="background: var(--color-bg-surface); border: 1px solid var(--color-border-subtle)"
        >
          <Search :size="16" style="color: var(--color-text-tertiary)" />
          <span class="text-sm whitespace-nowrap" style="color: var(--color-text-tertiary)">搜索</span>
        </div>

        <!-- Mobile hamburger -->
        <button
          class="md:hidden p-2 rounded-lg"
          style="color: var(--color-text-secondary)"
          aria-label="菜单"
          @click="toggleMobile"
        >
          <Menu v-if="!mobileOpen" :size="22" />
          <X v-else :size="22" />
        </button>
      </div>
    </nav>

    <!-- Mobile Drawer -->
    <div
      v-if="mobileOpen"
      class="md:hidden border-t"
      style="background: var(--color-bg-elevated); border-color: var(--color-border-subtle)"
    >
      <div class="flex flex-col" style="padding: 16px var(--content-padding)">
        <router-link
          v-for="item in navItems"
          :key="item.name"
          :to="item.path"
          class="px-3 py-3 rounded-lg text-sm font-medium no-underline"
          :style="isActive(item.name) ? 'color: var(--color-primary)' : 'color: var(--color-text-secondary)'"
          @click="mobileOpen = false"
        >
          {{ item.label }}
        </router-link>
      </div>
    </div>
  </header>
</template>
