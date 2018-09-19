import Vue from 'vue'

const GET_REQUEST_ERROR = 'Unknown error'
const POST_REQUEST_ERROR = 'Unknown error during data transfer'
const SERVER_URL = 'http://localhost:7000'

const ApiService = {
  get (slug = '') {
    return new Promise((resolve, reject) => {
      Vue.axios
        .get(SERVER_URL + slug)
        .then(resolve)
        .catch(() => {
          this.$message.error(GET_REQUEST_ERROR)
          reject(new Error(GET_REQUEST_ERROR))
        })
    })
  },
  post (slug = '', params) {
    return new Promise((resolve, reject) => {
      Vue.axios
        .post(SERVER_URL + slug, params)
        .then((response) => {
          const {data, status} = response
          if (status !== 200 || !(data instanceof Object) || (!data.ok && !(data.errors instanceof Object))) {
            throw new Error('Wrong post query')
          }
          resolve(data)
        })
        .catch(() => {
          this.$message.error(POST_REQUEST_ERROR)
          reject(new Error(POST_REQUEST_ERROR))
        })
    })
  },
  put () {},
  delete () {}
}

export default ApiService
