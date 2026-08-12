<script setup>
import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'
</script>

<template>
  <AppNavbar />
  <main class="site-main">
    <router-view />
  </main>
  <AppFooter />
</template>
