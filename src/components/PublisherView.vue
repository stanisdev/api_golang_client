<template>
  <div>
    <menu-navigation></menu-navigation>

    <loader v-bind:is_showing="loading"></loader>
    <transition name="fade">
      <div v-if="done">
        <publisherForm v-bind:mode="'edit'" v-bind:publisher="publisher"></publisherForm>
      </div>
    </transition>
  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'
import Loader from '@/components/Loader.vue'
import PublisherForm from '@/components/PublisherForm.vue'
import ApiService from '@/common/api.service'

export default {
  name: 'PublisherView',
  components: {
    MenuNavigation,
    Loader,
    PublisherForm
  },
  created () {
    this.fetchPublisher()
  },
  methods: {
    fetchPublisher () {
      const {id} = this.$route.params
      ApiService
        .setAuth().get
        .call(this, `/publisher/get/${id}`)
        .then((data) => {
          this.done = true
          this.publisher = data.payload
        })
        .catch(Symbol)
        .finally(() => {
          this.loading = false
        })
    }
  },
  data () {
    return {
      loading: true,
      done: false,
      publisher: {
        id: 19,
        name: 'Flickr'
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
