<script setup>
import { Mail, Globe, MessageCircle } from '@lucide/vue'

const currentYear = new Date().getFullYear()
</script>

<template>
  <footer class="border-t" style="background:var(--color-bg-elevated); border-color:var(--color-border-subtle); padding-top:var(--space-12); padding-bottom:var(--space-12);">
    <div class="max-w-[var(--container-max)] mx-auto"
         style="padding-left:var(--content-padding); padding-right:var(--content-padding);">

      <div class="grid grid-cols-1 md:grid-cols-3 gap-10 md:gap-8">
        <!-- Brand -->
        <div>
          <router-link to="/"
             class="inline-block no-underline"
             style="font-family:var(--font-display); font-weight:700; font-size:var(--text-xl); color:var(--color-primary); letter-spacing:var(--tracking-tight);">
            CineVerse
          </router-link>
          <p class="mt-3" style="font-family:var(--font-body); font-size:var(--text-sm); color:var(--color-text-tertiary); line-height:var(--leading-normal); max-width:280px;">
            探索光影世界，分享智慧火花
          </p>
        </div>

        <!-- Quick Links -->
        <div>
          <h3 class="mb-4" style="font-family:var(--font-body); font-weight:600; font-size:var(--text-sm); color:var(--color-text-primary); letter-spacing:var(--tracking-wide); text-transform:uppercase;">
            快速导航
          </h3>
          <ul class="flex flex-col gap-3 list-none m-0 p-0">
            <li>
              <router-link to="/" class="text-sm transition-colors duration-150 no-underline" style="color:var(--color-text-tertiary);">
                首页
              </router-link>
            </li>
            <li>
              <router-link to="/movie-recommend" class="text-sm transition-colors duration-150 no-underline" style="color:var(--color-text-tertiary);">
                影视推荐
              </router-link>
            </li>
            <li>
              <router-link to="/learning" class="text-sm transition-colors duration-150 no-underline" style="color:var(--color-text-tertiary);">
                学习资料
              </router-link>
            </li>
            <li>
              <router-link to="/ai-share" class="text-sm transition-colors duration-150 no-underline" style="color:var(--color-text-tertiary);">
                AI 分享
              </router-link>
            </li>
          </ul>
        </div>

        <!-- Contact -->
        <div>
          <h3 class="mb-4" style="font-family:var(--font-body); font-weight:600; font-size:var(--text-sm); color:var(--color-text-primary); letter-spacing:var(--tracking-wide); text-transform:uppercase;">
            联系方式
          </h3>
          <ul class="flex flex-col gap-3 list-none m-0 p-0">
            <li class="flex items-center gap-2">
              <Mail :size="16" style="color:var(--color-text-tertiary);" />
              <span class="text-sm" style="color:var(--color-text-tertiary);">hello@cineverse.site</span>
            </li>
            <li class="flex items-center gap-2">
              <Globe :size="16" style="color:var(--color-text-tertiary);" />
              <a href="#" class="text-sm transition-colors duration-150 no-underline" style="color:var(--color-text-tertiary);">GitHub</a>
            </li>
            <li class="flex items-center gap-2">
              <MessageCircle :size="16" style="color:var(--color-text-tertiary);" />
              <a href="#" class="text-sm transition-colors duration-150 no-underline" style="color:var(--color-text-tertiary);">Twitter</a>
            </li>
          </ul>
        </div>
      </div>

      <div class="mt-10 pt-6 border-t" style="border-color:var(--color-border-subtle);">
        <p class="text-center" style="font-family:var(--font-body); font-size:var(--text-xs); color:var(--color-text-tertiary);">
          &copy; {{ currentYear }} CineVerse. All rights reserved.
        </p>
      </div>
    </div>
  </footer>
</template>

<style scoped>
a:hover {
  color: var(--color-primary);
}
</style>