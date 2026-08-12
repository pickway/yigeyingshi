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
    path: '/movies/:id',
    name: 'movie-detail',
    component: () => import('@/views/MovieDetailView.vue'),
    meta: { title: '影片详情' },
  },
  {
    path: '/articles/:id',
    name: 'article-detail',
    component: () => import('@/views/ArticleDetailView.vue'),
    meta: { title: '文章详情' },
  },
  {
    path: '/learning',
    name: 'learning',
    component: () => import('@/views/LearningView.vue'),
    meta: { title: '学习资料' },
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('@/views/AboutView.vue'),
    meta: { title: '关于我' },
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: () => import('@/views/NotFoundView.vue'),
    meta: { title: '页面未找到' },
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
