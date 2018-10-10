<template>
  <el-form :rules="rules" ref="form" :model="notification" label-width="120px">

    <el-form-item label="Header" prop="header">
      <el-input class="inpCustom" v-model="notification.header"></el-input>
    </el-form-item>

    <el-form-item label="Button" prop="button">
      <el-input class="inpCustom" v-model="notification.button"></el-input>
    </el-form-item>

    <el-form-item label="Destination" prop="link">
      <el-input class="inpCustom" v-model="notification.link"></el-input>
    </el-form-item>

    <el-form-item label="Company" prop="company">
      <el-input class="inpCustom" v-model="notification.company"></el-input>
    </el-form-item>

    <el-form-item label="Expired" prop="expired">
      <el-col v-bind:style="{ width: '70%' }" :span="11">
        <el-date-picker @focus="datepickerFocused" type="date" v-model="notification.expired" placeholder="Pick an expired date" style="width: 100%;"></el-date-picker>
      </el-col>
    </el-form-item>

    <el-form-item label="Priority" prop="priority">
      <el-input class="inpCustom" v-model="notification.priority"></el-input>
    </el-form-item>

    <el-form-item label="Message" prop="message">
      <el-input class="inpCustom" type="textarea" v-model="notification.message"></el-input>
    </el-form-item>

    <el-form-item label="Image">

      <form style="display: none" action="" method="post" enctype="multipart/form-data">
        <input type="submit" value="uploadImage" id="uploadImage">
        <input type="file" name="image" id="selectImage">
      </form>
      <el-button @click="onUploadImage" type="primary">Upload</el-button>
      <div v-if="imgInvalid" class="imgError">Please select the image</div>

      <div style="margin-top: 23px" id="imagePreview"></div>
    </el-form-item>

    <el-form-item v-if="mode != 'create'" label="Created">
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
import Env from '@/env.js'
import UrlValidator from 'url-regex'

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
    const validateExpired = (role, value, callback) => { // eslint-disable-next-line
      const validateLessThen = (d) => {
        const now = new Date().getTime()
        if (d.getTime() <= now) {
          callback(new Error('Please input the date greater than now'))
          return false
        }
        return true
      } // eslint-disable-next-line
      if (typeof value === 'string' && value.length > 0 && /^[0-9]{4}\-[0-9]{1,2}\-[0-9]{1,2}$/.test(value) && new Date(value).toString() !== 'Invalid Date') {
        const d = new Date(value)
        if (!validateLessThen(d)) {
          return
        }
      } else if (typeof value === 'object') {
        if (!validateLessThen(value)) {
          return
        }
      }
      if (!value) {
        callback(new Error('Please input the expired date'))
      } else {
        callback()
      }
    }
    const validateLink = (role, value, callback) => {
      if (!value) {
        callback(new Error('Please input the destination'))
      } else if (typeof value === 'string' && !UrlValidator({exact: true}).test(value)) {
        callback(new Error('Wrong URL'))
      } else {
        callback()
      }
    }
    return {
      imgInvalid: false,
      selectedImage: null,
      rules: {
        header: [
          { validator: validateHeader, trigger: 'blur' }
        ],
        button: [
          { validator: validateButton, trigger: 'blur' }
        ],
        link: [
          { validator: validateLink, trigger: 'blur' }
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
  beforeMount () {
    if (this.mode === 'edit') {
      let ca = this.notification.created_at + '000'
      ca = new Date(+ca)
      if (ca.toString() === 'Invalid Date') {
        return
      }
      const m = ca.getMonth() + 1
      const d = ca.getDate()
      this.notification.created_at = ca.getFullYear() + '-' + (m < 10 ? '0' + m : m) + '-' + (d < 10 ? '0' + d : d)
    }
  },
  mounted () {
    if (this.mode === 'edit') {
      const img = this.notification.image
      document.getElementById('imagePreview').innerHTML = `<img src="${Env.UPLOADS_URL}/${img}">`
      this.selectedImage = img
    }
    const ta = document.getElementsByTagName('textarea')[0]
    if (ta instanceof Object) {
      ta.style.height = '150px'
      ta.style['font-size'] = '14px'
      ta.style['font-family'] = "'Avenir', Helvetica, Arial, sans-serif"
    }
  },
  methods: {
    datepickerFocused () {
      setTimeout(() => {
        document.querySelectorAll('[class*=el-date-picker]').forEach((elem) => {
          elem.style['font-family'] = 'Avenir, Helvetica, Arial, sans-serif'
        })
      }, 70)
    },
    onUploadImage () {
      const selBtn = document.getElementById('selectImage')
      selBtn.click()
      selBtn.onchange = () => {
        const file = selBtn.files[0]
        const fd = new FormData()
        fd.append('image', file)
        fetch(Env.API_URL + '/image/upload', {
          method: 'POST',
          body: fd,
          headers: {
            authorization: 'Bearer ' + localStorage.getItem('id_token')
          }
        }).then((response) => {
          return response.json()
        }).then((data) => {
          data = data instanceof Object ? data : {}
          const {ok, payload} = data
          if (ok !== true || typeof payload !== 'string') {
            throw new Error('Image not uploaded correctly')
          }
          const path = Env.UPLOADS_URL + '/' + payload
          this.selectedImage = payload
          this.imgInvalid = false

          document.getElementById('imagePreview').innerHTML = `<img src="${path}">`
        }).catch((err) => {
          console.log(err)
          this.$message.error('Image was not uploaded. Try again')
        })
      }
    },
    onSave () {
      let url = '/notification'
      if (this.mode === 'create') {
        url += '/create'
      } else {
        url += '/edit/' + this.notification.id
      }
      this.$refs['form'].validate((valid) => {
        let imgError
        if (typeof this.selectedImage === 'string' && this.selectedImage.length > 0) {
          imgError = false
        } else {
          imgError = true
        }
        this.imgInvalid = imgError

        let exp = this.notification.expired
        if (exp instanceof Date) {
          const m = exp.getMonth() + 1
          const d = exp.getDate()
          exp = exp.getFullYear() + '/' + (m < 10 ? '0' + m : m) + '/' + (d < 10 ? '0' + d : d)
        } else if (typeof exp === 'string') { // eslint-disable-next-line
          exp = exp.replace(/\-/g, '/')
          exp = exp.split('/').map((e) => {
            if (+e < 10) {
              return '0' + e
            }
            return e
          }).join('/')
        }
        if (valid && !this.imgInvalid) {
          const ntf = this.notification
          ApiService
            .setAuth().post
            .call(this, url, {
              message: ntf.message,
              image: this.selectedImage,
              header: ntf.header,
              priority: ntf.priority.toString(),
              expired: exp,
              button: ntf.button,
              link: ntf.link,
              company: ntf.company
            })
            .then((data) => {
              if (data instanceof Object && data.ok) {
                this.$router.push({
                  name: 'Notifications'
                })
                this.$message({
                  message: 'Notification was created successfully',
                  type: 'success'
                })
              } else {
                throw new Error('Notification was not created')
              }
            })
            .catch(Symbol)
            .finally(() => {})
        } else {
          return false
        }
      })
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
  .imgError {
    color: #f56c6c;
    font-size: 12px;
    position: absolute;
    margin-top: -9px;
  }
  .inpCustom {
    width: 70%
  }
</style>
