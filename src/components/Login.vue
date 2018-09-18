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
        <el-button type="primary" @click="submitForm" :loading="formSending">Login</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import Vue from 'vue'

export default {
  name: 'Login',
  data () {
    const validatePass = (rule, value, callback) => {
      if (!value) {
        callback(new Error('Please input the password'))
      } else {
        callback()
      }
    }
    const validateUsername = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the username'))
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
    submitForm () {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          this.formSending = true
          Vue.axios
            .get('https://yesno.wtf/api')
            .then((res) => {
              console.log(res)
              this.formSending = false
            })
            .catch(console.log)
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
