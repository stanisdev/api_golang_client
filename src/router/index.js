import Vue from 'vue'
import Router from 'vue-router'
import Notifications from '@/components/Notifications'
import Login from '@/components/Login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Notifications',
      component: Notifications
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    }
  ]
})
