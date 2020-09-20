<template>
  <component :is="layout">
    <router-view></router-view>
  </component>
</template>

<script>

import Header from './components/Header.vue'
const default_layout = 'default'

export default {
  name: 'App',
  components: {
    Header
  },
  computed: {
    layout(){
      return (this.$route.meta.layout || default_layout) + '-layout'
    }
  },
  created(){
    console.log(this.$route, this.$route.meta.layout)
  },
  updated() {
    console.log(this.$route, this.$route.meta.layout);
  },
  watch: {
    '$route' (to, from) {
      document.title = to.meta.title || 'Your Website'
    }
  },
}
</script>

<style>
html{
  //background:#eee
}
</style>
