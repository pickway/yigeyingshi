<script setup>
import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'
</script>

<template>
  <AppNavbar />
  <main class="min-h-screen pt-16">
    <router-view />
  </main>
  <AppFooter />
</template>
