<script setup>
const currentYear = new Date().getFullYear()
</script>

<template>
  <footer style="background: var(--color-bg-elevated); border-top: 1px solid var(--color-border-subtle)">
    <div class="mx-auto" style="max-width: var(--container-max); padding: 48px var(--content-padding) 32px">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-10">
        <!-- Brand -->
        <div>
          <span class="cine-display" style="color: var(--color-primary); font-size: 1.25rem">CineVerse</span>
          <p class="cine-body-sm mt-3" style="line-height: 1.7; max-width: 280px">
            探索影视世界，分享学习心得，洞见AI未来。一个关于光影与技术的个人空间。
          </p>
        </div>

        <!-- Quick Links -->
        <div>
          <h4
            class="text-sm font-semibold mb-4 whitespace-nowrap"
            style="color: var(--color-text-primary); font-family: var(--font-body)"
          >
            快速导航
          </h4>
          <div class="flex flex-col gap-2">
            <router-link
              v-for="item in [
                { label: '首页', path: '/' },
                { label: '影视推荐', path: '/movie-recommend' },
                { label: '学习资料', path: '/learning' },
                { label: 'AI 分享', path: '/ai-share' },
              ]"
              :key="item.path"
              :to="item.path"
              class="text-sm whitespace-nowrap transition-colors duration-150 no-underline"
              style="color: var(--color-text-secondary)"
            >
              {{ item.label }}
            </router-link>
          </div>
        </div>

        <!-- Contact -->
        <div>
          <h4
            class="text-sm font-semibold mb-4 whitespace-nowrap"
            style="color: var(--color-text-primary); font-family: var(--font-body)"
          >
            联系方式
          </h4>
          <div class="flex flex-col gap-2">
            <span class="text-sm whitespace-nowrap" style="color: var(--color-text-secondary); font-family: var(--font-body)">
              hello@cineverse.example
            </span>
            <span class="text-sm whitespace-nowrap" style="color: var(--color-text-secondary); font-family: var(--font-body)">
              GitHub / Twitter
            </span>
          </div>
        </div>
      </div>

      <!-- Bottom bar -->
      <div class="pt-6 flex items-center justify-between" style="border-top: 1px solid var(--color-border-subtle)">
        <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary); font-family: var(--font-body)">
          {{ currentYear }} CineVerse. All rights reserved.
        </span>
      </div>
    </div>
  </footer>
</template>
