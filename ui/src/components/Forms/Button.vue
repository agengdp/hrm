<template>
  <button class="border">{{ $text }}</button>
</template>

<script>
export default {
  name: "Button",
  props:[
    'text'
  ]
}
</script>

<style scoped>

</style>
