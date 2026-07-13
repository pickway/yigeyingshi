import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'index',
    component: () => import('@/views/HomeView.vue'),
    meta: { title: '首页' },
  },
  {
    path: '/movie-recommend',
    name: 'movie-recommend',
    component: () => import('@/views/MovieRecommendView.vue'),
    meta: { title: '影视推荐' },
  },
  {
    path: '/learning',
    name: 'learning',
    component: () => import('@/views/LearningView.vue'),
    meta: { title: '学习资料' },
  },
  {
    path: '/ai-share',
    name: 'ai-share',
    component: () => import('@/views/AiShareView.vue'),
    meta: { title: 'AI 分享' },
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

router.afterEach((to) => {
  document.title = `${to.meta.title} - CineVerse`
})

export default router
