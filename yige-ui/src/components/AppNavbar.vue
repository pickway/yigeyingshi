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
  <nav class="navbar">
    <div class="navbar-container">
      <!-- Logo -->
      <router-link to="/" class="logo">
        CineVerse
      </router-link>

      <!-- Desktop Nav Links -->
      <ul class="nav-links">
        <li v-for="item in navItems" :key="item.name">
          <router-link
            :to="item.path"
            class="nav-link"
            :class="{ 'nav-link-active': isActive(item.name) }"
          >
            {{ item.label }}
          </router-link>
        </li>
      </ul>

      <!-- Right: Search Icon + Mobile Hamburger -->
      <div class="nav-right">
        <button class="search-btn" aria-label="搜索">
          <Search :size="16" />
        </button>
        <button class="hamburger-btn" aria-label="菜单" @click="toggleMobile">
          <Menu v-if="!mobileOpen" :size="22" />
          <X v-else :size="22" />
        </button>
      </div>
    </div>

    <!-- Mobile Nav Drawer -->
    <div v-if="mobileOpen" class="mobile-drawer">
      <div class="mobile-drawer-content">
        <router-link
          v-for="item in navItems"
          :key="item.name"
          :to="item.path"
          class="mobile-link"
          :class="{ 'mobile-link-active': isActive(item.name) }"
          @click="mobileOpen = false"
        >
          {{ item.label }}
        </router-link>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 50;
  border-bottom: 1px solid;
  background: var(--color-bg-elevated);
  border-color: var(--color-border-subtle);
}

.navbar-container {
  max-width: var(--container-max);
  margin-left: auto;
  margin-right: auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
  padding-left: var(--content-padding);
  padding-right: var(--content-padding);
}

/* Logo */
.logo {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-decoration: none;
  font-family: var(--font-display);
  font-weight: 700;
  font-size: var(--text-xl);
  color: var(--color-primary);
  letter-spacing: var(--tracking-tight);
}

/* Desktop Nav Links */
.nav-links {
  display: none;
  align-items: center;
  gap: 0.25rem;
  list-style: none;
  margin: 0;
  padding: 0;
}

@media (min-width: 768px) {
  .nav-links {
    display: flex;
  }
}

.nav-link {
  display: inline-flex;
  align-items: center;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  font-family: var(--font-body);
  color: var(--color-text-secondary);
  text-decoration: none;
  transition: color 150ms, background-color 150ms;
}

.nav-link:hover {
  color: var(--color-text-primary);
  background: var(--color-bg-muted);
}

.nav-link-active {
  color: var(--color-primary);
  background: var(--color-primary-tint-1);
}

.nav-link-active:hover {
  color: var(--color-primary);
  background: var(--color-primary-tint-1);
}

/* Right: Search + Mobile */
.nav-right {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.search-btn {
  display: none;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  border-radius: 9999px;
  border: 1px solid var(--color-border-default);
  background: transparent;
  color: var(--color-text-tertiary);
  cursor: pointer;
  transition: color 150ms, border-color 150ms;
}

@media (min-width: 768px) {
  .search-btn {
    display: inline-flex;
  }
}

.search-btn:hover {
  color: var(--color-text-primary);
  border-color: var(--color-text-tertiary);
}

.hamburger-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  width: 40px;
  height: 40px;
  border-radius: 0.5rem;
  border: none;
  background: transparent;
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: color 150ms, background-color 150ms;
}

@media (min-width: 768px) {
  .hamburger-btn {
    display: none;
  }
}

.hamburger-btn:hover {
  color: var(--color-text-primary);
  background: var(--color-bg-muted);
}

/* Mobile Nav Drawer */
.mobile-drawer {
  border-top: 1px solid var(--color-border-subtle);
  background: var(--color-bg-elevated);
}

.mobile-drawer-content {
  display: flex;
  flex-direction: column;
  padding: var(--space-4) var(--content-padding);
}

.mobile-link {
  display: inline-flex;
  align-items: center;
  padding: 0.75rem 0.75rem;
  border-radius: 0.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--color-text-secondary);
  text-decoration: none;
  transition: color 150ms, background-color 150ms;
}

.mobile-link:hover {
  color: var(--color-text-primary);
  background: var(--color-bg-muted);
}

.mobile-link-active {
  color: var(--color-primary);
  background: var(--color-primary-tint-1);
}

.mobile-link-active:hover {
  color: var(--color-primary);
  background: var(--color-primary-tint-1);
}
</style>