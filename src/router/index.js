import Vue from 'vue'
import Router from 'vue-router'
import Notifications from '@/components/Notifications'
import Login from '@/components/Login'
import CreateNotification from '@/components/CreateNotification'
import NotificationView from '@/components/NotificationView'
import UserProfile from '@/components/UserProfile'
import Publishers from '@/components/Publishers'
import CreatePublisher from '@/components/CreatePublisher'
import NotFound from '@/components/NotFound'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: {
        guest: true
      }
    },
    {
      path: '/notifications',
      name: 'Notifications',
      component: Notifications,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/notifications/create',
      name: 'CreateNotification',
      component: CreateNotification,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/notifications/edit/:id',
      name: 'NotificationView',
      component: NotificationView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/user/profile',
      name: 'UserProfile',
      component: UserProfile,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/publishers',
      name: 'Publishers',
      component: Publishers,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/publishers/create',
      name: 'CreatePublisher',
      component: CreatePublisher,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/notFound',
      name: 'NotFound',
      component: NotFound
    }
  ]
})

router.beforeEach((to, from, next) => {
  const hasToken = typeof localStorage.getItem('id_token') === 'string'
  if (to.path === '/') {
    return next({
      path: '/login'
    })
  }
  if (!to.matched.length) {
    return next('/notFound')
  }
  if (to.path !== '/login' && to.meta instanceof Object && to.meta.requiresAuth && !hasToken) {
    return next({
      path: '/login',
      params: { nextUrl: to.fullPath }
    })
  } else if (hasToken && to.path === '/login') {
    return next({
      path: '/notifications'
    })
  }
  next()
})

export default router
