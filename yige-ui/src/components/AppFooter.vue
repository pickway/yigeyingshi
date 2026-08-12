<script setup>
import { ArrowUpRight, Mail } from '@lucide/vue'
const currentYear = new Date().getFullYear()
</script>

<template>
  <footer class="site-footer">
    <div class="footer-grid content-shell">
      <div class="footer-intro">
        <span class="eyebrow">CineVerse Journal</span>
        <h2>愿每一次观看，<br />都成为下一次创作的开始。</h2>
        <p>一个关于电影、学习与 AI 创作的个人数字花园。保持好奇，持续记录。</p>
      </div>
      <div class="footer-links">
        <div><strong>探索</strong><router-link to="/movie-recommend">影视片单</router-link><router-link to="/learning">学习资料</router-link><router-link to="/ai-share">AI 实验室</router-link></div>
        <div><strong>关于</strong><router-link to="/about">关于我</router-link><a href="mailto:hello@cineverse.site">联系我</a><router-link to="/about">创作理念</router-link></div>
      </div>
      <a class="footer-mail" href="mailto:hello@cineverse.site"><Mail :size="16" /> hello@cineverse.site <ArrowUpRight :size="15" /></a>
    </div>
    <div class="footer-bottom content-shell"><span>© {{ currentYear }} CineVerse</span><span>在光影与代码之间持续更新</span></div>
  </footer>
</template>

<style scoped>
.site-footer{position:relative;padding:88px 0 30px;border-top:1px solid var(--color-border-subtle);background:#08090c;overflow:hidden}.site-footer::before{position:absolute;top:-120px;left:8%;width:380px;height:220px;border-radius:50%;background:rgba(212,168,83,.06);filter:blur(70px);content:''}
.footer-grid{position:relative;display:grid;grid-template-columns:1.4fr .8fr;gap:80px}.eyebrow{color:var(--color-primary);font-size:.68rem;letter-spacing:.2em;text-transform:uppercase}.footer-intro h2{margin:18px 0;font-family:var(--font-display);font-size:clamp(2rem,4vw,3.5rem);font-weight:500;line-height:1.15}.footer-intro p{max-width:520px;color:var(--color-text-tertiary);line-height:1.8}.footer-links{display:grid;grid-template-columns:1fr 1fr;gap:32px}.footer-links div{display:flex;flex-direction:column;gap:13px}.footer-links strong{margin-bottom:8px;color:var(--color-text-primary);font-size:.75rem;letter-spacing:.12em}.footer-links a{color:var(--color-text-tertiary);font-size:.85rem;text-decoration:none}.footer-links a:hover{color:var(--color-primary)}.footer-mail{grid-column:2;display:flex;align-items:center;justify-content:space-between;padding:16px 0;border-bottom:1px solid var(--color-border-default);color:var(--color-text-primary);text-decoration:none}.footer-bottom{display:flex;justify-content:space-between;margin-top:64px;padding-top:22px;border-top:1px solid var(--color-border-subtle);color:var(--color-text-tertiary);font-size:.7rem;letter-spacing:.08em}
@media(max-width:760px){.site-footer{padding-top:64px}.footer-grid{grid-template-columns:1fr;gap:44px}.footer-mail{grid-column:1}.footer-bottom{gap:20px;flex-direction:column;margin-top:44px}}
</style>
