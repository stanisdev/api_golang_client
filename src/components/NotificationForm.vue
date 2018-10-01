<template>
  <el-form :rules="rules" ref="form" :model="notification" label-width="120px">

    <el-form-item label="Header" prop="header">
      <el-input v-model="notification.header"></el-input>
    </el-form-item>

    <el-form-item label="Button" prop="button">
      <el-input v-model="notification.button"></el-input>
    </el-form-item>

    <el-form-item label="Company" prop="company">
      <el-input v-model="notification.company"></el-input>
    </el-form-item>

    <el-form-item label="Expired" prop="expired">
      <el-col :span="11">
        <el-date-picker type="date" v-model="notification.expired" placeholder="Pick an expired date" style="width: 100%;"></el-date-picker>
      </el-col>
    </el-form-item>

    <el-form-item label="Priority" prop="priority">
      <el-input v-model="notification.priority"></el-input>
    </el-form-item>

    <el-form-item label="Message" prop="message">
      <el-input type="textarea" v-model="notification.message"></el-input>
    </el-form-item>

    <el-form-item label="Image">
      <span>{{ notification.image }}</span>
    </el-form-item>

    <el-form-item label="Created">
      <span>{{ notification.created_at }}</span>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSave">Save</el-button>
      <el-button @click="onCancel">Cancel</el-button>
    </el-form-item>

  </el-form>
</template>

<script>
import ApiService from '@/common/api.service'

export default {
  name: 'NotificationForm',
  props: ['notification', 'mode'],
  data () {
    const validateHeader = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the header'))
      } else {
        callback()
      }
    }
    const validateButton = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the button caption'))
      } else {
        callback()
      }
    }
    const validateCompany = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the company name'))
      } else {
        callback()
      }
    }
    const validatePriority = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the priority'))
      } else if (!/^\d+$/ig.test(value)) {
        callback(new Error('Wrong priority value'))
      } else {
        callback()
      }
    }
    const validateMessage = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the message'))
      } else {
        callback()
      }
    }
    const validateExpired = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the expired date'))
      } else {
        callback()
      }
    }
    return {
      rules: {
        header: [
          { validator: validateHeader, trigger: 'blur' }
        ],
        button: [
          { validator: validateButton, trigger: 'blur' }
        ],
        company: [
          { validator: validateCompany, trigger: 'blur' }
        ],
        priority: [
          { validator: validatePriority, trigger: 'blur' }
        ],
        message: [
          { validator: validateMessage, trigger: 'blur' }
        ],
        expired: [
          { validator: validateExpired, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    onSave () {
      let url = '/notification'
      if (this.mode === 'create') {
        url += '/create'
      } else {
        url += '/edit/' + this.notification.id
      }
      ApiService
        .setAuth().post
        .call(this, url, {
          message: '',
          image: '',
          header: '',
          priority: '',
          expired: '',
          button: '',
          link: '',
          company: ''
        })
        .then((data) => {
          console.log(data)
        })
        .catch(Symbol)
        .finally(() => {})
    },
    onCancel () {
      this.$router.push({
        name: 'Notifications'
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
