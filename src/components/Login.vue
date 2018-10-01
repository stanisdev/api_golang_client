<template>
  <div>
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm">
      <h2>Login Page</h2>

      <el-form-item label="Username" prop="username">
        <el-input v-model="ruleForm.username" autocomplete="off"></el-input>
      </el-form-item>

      <el-form-item label="Password" prop="pass">
        <el-input v-model="ruleForm.pass" type="password" autocomplete="off"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm(ruleForm.username, ruleForm.pass)" :loading="formSending">Login</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import ApiService from '@/common/api.service'
let errors = {}

export default {
  name: 'Login',
  data () {
    const validatePass = (rule, value, callback) => {
      if (!value) {
        callback(new Error('Please input the password'))
      } else if (errors instanceof Object && typeof errors.pass === 'string') {
        callback(new Error(errors.pass))
      } else {
        callback()
      }
    }
    const validateUsername = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the username'))
      } else if (errors instanceof Object && typeof errors.username === 'string') {
        callback(new Error(errors.username))
      } else {
        callback()
      }
    }
    return {
      formSending: false,
      ruleForm: {
        username: '',
        pass: ''
      },
      rules: {
        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        username: [
          { validator: validateUsername, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm (username, pass) {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          this.formSending = true
          ApiService.post
            .call(this, '/user/login', { username, password: pass })
            .then((data) => {
              if (!data.ok) { // Show errors
                errors = data.errors
                return this.$refs['ruleForm'].validate(() => {
                  errors = {}
                })
              }
              localStorage.setItem('id_token', data.payload.id_token)
              this.$router.push({ name: 'Notifications' })
            })
            .catch(Symbol)
            .finally(() => {
              this.formSending = false
            })
        } else {
          return false
        }
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
