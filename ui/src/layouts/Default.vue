<template>
  <div>
    <Header></Header>
    <Subheader :title="title"></Subheader>
    <slot/>
  </div>

</template>

<script>
import Header from '../components/Header.vue'
import Subheader from '../components/Subheader.vue'
export default {
  name: "Default",
  components: {
    Header,
    Subheader
  },
  data(){
    return {
      title: 'aa'
    }
  },
  watch: {
    '$route' (to, from) {
      this.title = to.meta.title || 'Your Website'
    }
  },
}
</script>

<style scoped>

</style>