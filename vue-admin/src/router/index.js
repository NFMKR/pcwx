/* eslint-disable */
import { createRouter, createWebHistory } from 'vue-router'
import { ElMessage } from 'element-plus'
import store from '../store'

// 路由配置
const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/LoginPage.vue'),
    meta: { title: '登录' }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/RegisterPage.vue'),
    meta: { title: '注册' }
  },
  {
    path: '/',
    component: () => import('../layout/MainLayout.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../views/DashboardPage.vue'),
        meta: { 
          title: '首页',
          requiresAuth: true
        }
      }
    ]
  },
  {
    path: '/403',
    name: 'Forbidden',
    component: () => import('../views/error/403Page.vue'),
    meta: { title: '403' }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('../views/error/404Page.vue'),
    meta: { title: '404' }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  document.title = to.meta.title ? `${to.meta.title} - 管理系统` : '管理系统'

  const token = store.state.token
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)

  // 打印调试信息
  console.log('Navigation:', {
    to: to.path,
    requiresAuth,
    hasToken: !!token
  })

  // 不需要认证的路由直接通过
  if (!requiresAuth) {
    next()
    return
  }

  // 需要认证但没有token
  if (!token) {
    ElMessage.warning('请先登录')
    next('/login')
    return
  }

  // 检查用户角色权限
  const userRole = store.state.user?.role
  const requiredRoles = to.meta.roles

  if (requiredRoles && !requiredRoles.includes(userRole)) {
    ElMessage.error('没有访问权限')
    next('/403')
    return
  }

  next()
})

export default router
