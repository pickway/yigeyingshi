<script setup>
import { ArrowRight, BookOpen, Clapperboard, Sparkles } from '@lucide/vue'
</script>
<template>
  <div class="about-page">
    <section class="about-hero content-shell">
      <div><span class="eyebrow">About the Journal</span><h1 class="display-title">你好，<br />我是“一个影视”。</h1></div>
      <div class="about-lead"><p>我把这里当作一间常亮的小放映室：收藏真正喜欢的电影，整理走过的学习路径，也诚实记录 AI 如何改变影像创作。</p><p>不追逐每一个热点，更关心那些值得反复观看、长期练习和慢慢想明白的事。</p></div>
    </section>
    <section class="manifesto"><div class="content-shell"><span>01</span><h2>观看不是消费，<br />而是一种练习。</h2><p>练习理解不同的人，辨认画面背后的选择，并对自己的感受保持诚实。</p></div></section>
    <section class="pillars content-shell">
      <article><Clapperboard :size="28" /><small>FILM</small><h3>影视片单</h3><p>以个人观看经验筛选作品，不追求大而全，只留下值得一次次重看的电影。</p><router-link to="/movie-recommend">浏览片单 <ArrowRight :size="15" /></router-link></article>
      <article><BookOpen :size="28" /><small>LEARN</small><h3>创作学习</h3><p>从剪辑、色彩到声音，把零散资料组织成可以真正走完的学习路线。</p><router-link to="/learning">开始学习 <ArrowRight :size="15" /></router-link></article>
      <article><Sparkles :size="28" /><small>LAB</small><h3>AI 实验</h3><p>把新工具放进真实的创作流程里检验，记录能力，也记录边界。</p><router-link to="/ai-share">进入实验室 <ArrowRight :size="15" /></router-link></article>
    </section>
    <section class="contact content-shell"><span class="eyebrow">Say Hello</span><h2 class="section-title">如果你也在观看、学习或创造，<br />欢迎写信聊聊。</h2><a href="mailto:hello@cineverse.site" class="btn-primary">hello@cineverse.site <ArrowRight :size="16" /></a></section>
  </div>
</template>
<style scoped>
.about-page{padding:80px 0 120px}.about-hero{display:grid;grid-template-columns:1.1fr .8fr;gap:100px;align-items:end}.about-hero h1{margin:25px 0 0}.about-lead p{color:var(--color-text-secondary);font-family:var(--font-display);font-size:1.15rem;line-height:1.9}.manifesto{margin:110px 0;padding:120px 0;background:linear-gradient(120deg,#111218,#17150f)}.manifesto .content-shell{position:relative}.manifesto span{position:absolute;top:-50px;right:5%;font-family:var(--font-display);font-size:11rem;color:rgba(217,173,95,.05)}.manifesto h2{margin:0;font-family:var(--font-display);font-size:clamp(2.8rem,6vw,5.8rem);font-weight:500;line-height:1.15}.manifesto p{max-width:500px;margin:32px 0 0;color:var(--color-text-tertiary);line-height:1.8}.pillars{display:grid;grid-template-columns:repeat(3,1fr);gap:1px;background:var(--color-border-subtle);border:1px solid var(--color-border-subtle)}.pillars article{padding:42px;background:var(--color-bg-base)}.pillars article>svg{color:var(--color-primary)}.pillars small{display:block;margin:35px 0 12px;color:var(--color-text-tertiary);letter-spacing:.18em}.pillars h3{font-family:var(--font-display);font-size:1.7rem}.pillars p{min-height:100px;color:var(--color-text-tertiary);line-height:1.75}.pillars a{display:flex;align-items:center;gap:7px;color:var(--color-primary);font-size:.8rem;text-decoration:none}.contact{margin-top:130px;text-align:center}.contact h2{margin:20px 0 35px}.contact .btn-primary{display:inline-flex}
@media(max-width:800px){.about-hero{grid-template-columns:1fr;gap:40px}.pillars{grid-template-columns:1fr}.pillars p{min-height:0}.manifesto{padding:80px 0}}
</style>

