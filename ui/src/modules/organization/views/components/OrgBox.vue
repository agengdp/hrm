<template>
  <div class="flex border w-1/6 text-left p-2 rounded">
    <div class="flex-1">
      <div class="text-lg">
        Direktur Utama
      </div>
      <div class="text-sm">
        dirut broo
      </div>
    </div>
    <div class="relative">
      <div class="border hover:border-blue-500 rounded-sm hover:border-blue-300">
        <svg class="h-4 fill-current text-gray-300 hover:text-blue-500 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M4 12a2 2 0 1 1 0-4 2 2 0 0 1 0 4zm6 0a2 2 0 1 1 0-4 2 2 0 0 1 0 4zm6 0a2 2 0 1 1 0-4 2 2 0 0 1 0 4z"/></svg>
      </div>
      <div class="absolute top-10 right--0 bg-white w-24">
        <ul class="text-xs border bg-gray-100 rounded-sm">
          <li class="border-b p-1 hover:bg-blue-500 hover:text-white">Add Parent</li>
          <li class="border-b p-1 hover:bg-blue-500 hover:text-white">Add Children</li>
          <li class="p-1 text-red-700 hover:bg-red-600 hover:text-white">Delete</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
name: "OrgBox"
}
</script>

<style scoped>

</style>
