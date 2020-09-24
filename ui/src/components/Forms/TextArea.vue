<template>
  <div class="mb-2">
    <label class="block">{{ label }}</label>
    <textarea class="border border-gray-400 hover:border-blue-500 text-base w-full py-1 px-2"
              :placeholder="placeholder"
    >
    </textarea>
  </div>
</template>

<script>
export default {
  name: "TextArea",
  props: [
    'label',
    'placeholder'
  ]
}
</script>

<style scoped>

</style>
