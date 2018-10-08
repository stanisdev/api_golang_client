<template>
  <div>
    <menu-navigation></menu-navigation>

  <el-dialog
    title="Change Password"
    :visible.sync="dialogVisible"
    width="30%">

    <el-form :rules="rules" ref="PForm" :model="form" label-width="150px">
      <el-form-item label="Current password" prop="old_password">
        <el-input type="password" v-model="form.old_password"></el-input>
      </el-form-item>
      <el-form-item label="New password" prop="new_password">
        <el-input type="password" v-model="form.new_password"></el-input>
      </el-form-item>
      <el-form-item label="Confirm" prop="new_password_confirm">
        <el-input type="password" v-model="form.new_password_confirm"></el-input>
      </el-form-item>
    </el-form>

    <div slot="footer" v-bind:style="{ marginTop: '-20px' }" class="dialog-footer">
      <el-button @click="dialogVisible = false">Cancel</el-button>
      <el-button type="primary" :loading="formSending" @click="handleSavePassword">Save</el-button>
    </div>
  </el-dialog>

    <el-form label-width="120px">

      <el-form-item label="Username">
        <span>mr.admin</span>
      </el-form-item>

      <el-form-item label="Password">
        <el-button @click="handleStartChangePswd" type="primary">Change</el-button>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'
import ApiService from '@/common/api.service'
let errors = {}

export default {
  name: 'UserProfile',
  components: {
    MenuNavigation
  },
  data () {
    const validateOldPassword = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the current password'))
      } else if (errors instanceof Object && typeof errors.old_password === 'string') {
        callback(new Error(errors.old_password))
      } else {
        callback()
      }
    }
    const validateNewPassword = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the new password'))
      } else if (this.form.old_password === value) {
        callback(new Error('The current password and new one should not be equal'))
      } else {
        callback()
      }
    }
    const validateNewPasswordConfirm = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the new password confirm'))
      } else if (this.form.new_password !== value) {
        callback(new Error('The new password and confirm are not equal'))
      } else {
        callback()
      }
    }
    return {
      formSending: false,
      firstShowing: false,
      dialogVisible: false,
      form: {
        old_password: '',
        new_password: '',
        new_password_confirm: ''
      },
      rules: {
        old_password: [
          { validator: validateOldPassword, trigger: 'blur' }
        ],
        new_password: [
          { validator: validateNewPassword, trigger: 'blur' }
        ],
        new_password_confirm: [
          { validator: validateNewPasswordConfirm, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleClose (done) {
      this.$confirm('Are you sure to close this dialog?')
        .then(() => {
          done()
        })
        .catch(Symbol)
    },
    handleSavePassword () {
      this.$refs['PForm'].validate((valid) => {
        if (valid) {
          const {form} = this
          const url = '/user/change/password'
          this.formSending = true
          ApiService
            .setAuth().post
            .call(this, url, {
              old_password: form.old_password,
              new_password: form.new_password
            })
            .then((data) => {
              if (data instanceof Object && !data.ok) { // Wrong password
                errors = data.errors
                this.$refs['PForm'].validate(() => {
                  errors = {}
                })
              } else if (data instanceof Object && data.ok) { // Success
                this.$message({
                  message: 'Password was successfully changed',
                  type: 'success'
                })
                this.dialogVisible = false
              } else {
                throw new Error('Undefined error')
              }
            })
            .catch(Symbol)
            .finally(() => {
              this.formSending = false
            })
        }
      })
    },
    handleStartChangePswd () {
      if (this.firstShowing) {
        this.$refs['PForm'].resetFields()
      }
      this.dialogVisible = true
      this.firstShowing = true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
