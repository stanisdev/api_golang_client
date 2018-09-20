import Vue from 'vue'
import Router from 'vue-router'
import Notifications from '@/components/Notifications'
import Login from '@/components/Login'
import CreateNotification from '@/components/CreateNotification'
import NotificationView from '@/components/NotificationView'

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
    }
  ]
})

router.beforeEach((to, from, next) => {
  const hasToken = typeof localStorage.getItem('id_token') === 'string'
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
