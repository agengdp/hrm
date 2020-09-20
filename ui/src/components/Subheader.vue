<template>
  <div class="bg-gray-200 text-gray-800 border-b">
      <div class="container mx-auto py-2 flex flex-wrap flex-col flex-row">
          <h1 class="text-2xl">{{ title }}</h1>
      </div>
  </div>
</template>

<script>
export default {
  name: "Subheader",
  props: [
    'title'
  ],
}
</script>

<style scoped>

</style>