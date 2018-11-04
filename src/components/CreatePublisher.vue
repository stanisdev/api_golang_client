<template>
  <div>
    <menu-navigation></menu-navigation>
    <publisherForm v-bind:mode="'create'" v-bind:publisher="publisher"></publisherForm>
  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'
import PublisherForm from '@/components/PublisherForm.vue'

export default {
  name: 'CreatePublisher',
  components: {
    MenuNavigation,
    PublisherForm
  },
  data () {
    return {
      publisher: {}
    }
  },
  methods: {}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
