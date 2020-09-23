<template>
  <button class="border {{ kelas }}">{{ text }}</button>
</template>

<script>
export default {
  name: "Button",
  props:[
    'text',
    'style',
  ],
  computed: {
    kelas(){
      return 'p-5'
    }
  }
}
</script>

<style scoped>

</style>
