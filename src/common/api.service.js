import Vue from 'vue'
import Env from '@/env.js'

const GET_REQUEST_ERROR = 'Unknown error'
const POST_REQUEST_ERROR = 'Unknown error during data transfer'

const connectToServer = (queryFunc) => {
  return new Promise((resolve, reject) => {
    let done = false
    const tmId = setTimeout(() => {
      if (!done) {
        reject(new Error('Code #1'))
      }
      clearTimeout(tmId)
    }, Env.TIMEOUT_SERVER_RESPOND)
    queryFunc
      .then(resolve)
      .catch(reject)
      .finally(() => {
        done = true
        clearTimeout(tmId)
      })
  })
}

const logout = function () {
  localStorage.removeItem('id_token')
  this.$router.push({
    name: 'Login'
  })
}

const ApiService = {
  setAuth () {
    Vue.axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('id_token')
    return this
  },
  get (slug = '') {
    return new Promise((resolve, reject) => {
      connectToServer(Vue.axios.get(Env.API_URL + slug))
        .then((response) => {
          const {data, status} = response
          if (status !== 200 || !(data instanceof Object) || (!data.ok && !(data.errors instanceof Object))) {
            throw new Error('Wrong GET query')
          }
          resolve(data)
        })
        .catch((err) => {
          err = err.toString()
          if (err.endsWith('code 401')) {
            logout.call(this)
            return
          }
          let msg = GET_REQUEST_ERROR
          if (err.endsWith('Code #1')) {
            msg = 'Server does not respond'
          }
          this.$message.error(msg)
          reject(new Error(GET_REQUEST_ERROR))
        })
    })
  },
  post (slug = '', params) {
    return new Promise((resolve, reject) => {
      connectToServer(Vue.axios.post(Env.API_URL + slug, params))
        .then((response) => {
          const {data, status} = response
          if (status !== 200 || !(data instanceof Object) || (!data.ok && !(data.errors instanceof Object))) {
            throw new Error('Wrong POST query')
          }
          resolve(data)
        })
        .catch((err) => {
          err = err.toString()
          if (err.endsWith('code 401')) {
            logout.call(this)
            return
          }
          let msg = POST_REQUEST_ERROR
          if (err.toString().endsWith('Code #1')) {
            msg = 'Server does not respond'
          }
          this.$message.error(msg)
          reject(new Error(POST_REQUEST_ERROR))
        })
    })
  },
  put () {},
  delete () {}
}

export default ApiService
