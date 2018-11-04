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

export default {
  name: 'PublisherForm',
  props: ['publisher', 'mode'],
  data () {
    const validateName = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the name'))
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
    onSave () {},
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
