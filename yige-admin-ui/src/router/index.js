import { createRouter, createWebHistory } from 'vue-router'
import { getToken } from '@/utils/session'

const routes=[
  {path:'/login',name:'login',component:()=>import('@/views/LoginView.vue'),meta:{public:true,title:'登录'}},
  {path:'/',component:()=>import('@/components/AdminLayout.vue'),children:[
    {path:'',redirect:'/dashboard'},
    {path:'dashboard',name:'dashboard',component:()=>import('@/views/DashboardView.vue'),meta:{title:'工作台'}},
    {path:'movies',name:'movies',component:()=>import('@/views/ResourceView.vue'),meta:{title:'影片管理',resource:'movies'}},
    {path:'articles',name:'articles',component:()=>import('@/views/ResourceView.vue'),meta:{title:'文章管理',resource:'articles'}},
    {path:'courses',name:'courses',component:()=>import('@/views/ResourceView.vue'),meta:{title:'课程管理',resource:'courses'}},
    {path:'paths',name:'paths',component:()=>import('@/views/ResourceView.vue'),meta:{title:'学习路径',resource:'paths'}},
    {path:'tools',name:'tools',component:()=>import('@/views/ResourceView.vue'),meta:{title:'AI 工具',resource:'tools'}},
    {path:'subscribers',name:'subscribers',component:()=>import('@/views/SubscribersView.vue'),meta:{title:'订阅用户'}},
  ]},
  {path:'/:pathMatch(.*)*',redirect:'/dashboard'},
]
const router=createRouter({history:createWebHistory(),routes})
router.beforeEach(to=>{if(!to.meta.public&&!getToken())return {name:'login',query:{redirect:to.fullPath}};if(to.name==='login'&&getToken())return {name:'dashboard'}})
router.afterEach(to=>{document.title=`${to.meta.title || '管理后台'} - CineVerse Admin`})
export default router

