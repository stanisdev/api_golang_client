<template>
  <el-form :rules="rules" ref="form" :model="publisher" label-width="120px">

    <el-form-item label="Name" prop="name">
      <el-input class="inpCustom" v-model="publisher.name"></el-input>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSave">Save</el-button>
      <el-button @click="onCancel">Cancel</el-button>
    </el-form-item>

  </el-form>
</template>

<script>
import ApiService from '@/common/api.service'
let errors = {}

export default {
  name: 'PublisherForm',
  props: ['publisher', 'mode'],
  data () {
    const validateName = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the name'))
      } else if (errors instanceof Object && typeof errors.name === 'string') {
        callback(new Error(errors.name))
      } else {
        callback()
      }
    }
    return {
      rules: {
        name: [
          { validator: validateName, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    onSave () {
      let url = '/publisher'
      const {mode} = this
      if (mode === 'create') {
        url += '/create'
      } else {
        url += '/edit/' + this.publisher.id
      }
      this.$refs['form'].validate((valid) => {
        if (valid) {
          ApiService
            .setAuth().post
            .call(this, url, {
              name: this.publisher.name
            })
            .then((data) => {
              if (data instanceof Object) {
                if (data.ok === true) { // Success
                  this.$router.push({
                    name: 'Publishers'
                  })
                } else if (data.errors instanceof Object) { // Show errors
                  errors = data.errors
                  this.$refs['form'].validate(() => {
                    errors = {}
                  })
                } else {
                  throw new Error(`Wrong response`)
                }
              } else {
                throw new Error(`Publisher was not ${mode}d`)
              }
            })
            .catch(() => {
              this.$message.error('Unexpected Error')
            })
            .finally(() => {})
        } else {
          return false
        }
      })
    },
    onCancel () {
      this.$router.push({
        name: 'Publishers'
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .inpCustom {
    width: 70%
  }
</style>
