<template>
    <div class="mb-2">
      <label class="block">{{ label }}</label>
      <input type="text" class="border border-gray-400 hover:border-blue-500 text-base w-full py-1 px-2"
             :placeholder="placeholder"
             :value="value"
      >
    </div>
</template>

<script>
export default {
    name: 'Input',
    props: [
      'placeholder',
      'value',
      'label'
    ]
}
</script>

<style>

</style>
