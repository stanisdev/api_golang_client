<template>
  <div class="top-nav-menu">
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
      <el-menu-item index="0-publishers">Publishers</el-menu-item>
      <el-menu-item index="1-notifications">Notifications</el-menu-item>
      <el-submenu index="2">
        <template slot="title">Actions</template>
        <el-menu-item v-bind:style="{ fontFamily: 'Avenir, Helvetica, Arial, sans-serif' }" index="2-create-notification">Create notification</el-menu-item>
        <el-menu-item v-bind:style="{ fontFamily: 'Avenir, Helvetica, Arial, sans-serif' }" index="2-create-publisher">Create publisher</el-menu-item>
      </el-submenu>
      <el-submenu v-bind:style="{ 'float': 'right' }" index="3">
        <template slot="title">Settings</template>
        <el-menu-item v-bind:style="{fontFamily: 'Avenir, Helvetica, Arial, sans-serif'}" index="3-user-profile">User profile</el-menu-item>
        <el-menu-item v-bind:style="{fontFamily: 'Avenir, Helvetica, Arial, sans-serif'}" index="3-logout">Logout</el-menu-item>
      </el-submenu>
    </el-menu>
  </div>
</template>

<script>
const routes = {
  'create-notification': 'CreateNotification',
  'create-publisher': 'CreatePublisher',
  notifications: 'Notifications',
  'user-profile': 'UserProfile',
  publishers: 'Publishers'
}

const urls = {
  '/publishers': '0-publishers',
  '/notifications': '1-notifications',
  '/notifications/create': '2-create-notification',
  '/publishers/create': '2-create-publisher',
  '/user/profile': '3-user-profile'
}

/* eslint-disable */
export default {
  name: 'MenuNavigation',
  data () {
    let url = window.location.href.split('#')[1];
    if (typeof url === 'string' && url.startsWith('/publishers/edit')) {
      url = '0-publishers';
    }
    else {
      url = urls[url]
      if (!url) {
        url = '1-notifications'
      }
    }
    return {
      activeIndex: url
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      let path = key.split(/^\d+\-/ig)[1]
      if (path === 'logout') {
        localStorage.removeItem('id_token')
        this.$router.push({ name: 'Login' })
        return
      }
      if (typeof path === 'string' && routes.hasOwnProperty(path)) {
        path = routes[path]
      } else {
        path = routes.notifications
      }
      this.$router.push({ name: path })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .top-nav-menu {
    margin-bottom: 37px;
  }
</style>
