<template>
  <header class="bg-blue-600 text-white">
    <div class="container mx-auto flex flex-wrap flex-col md:flex-row items-center py-1">
      <a class="flex title-font font-medium items-center mb-4 md:mb-0">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" class="w-10 h-10 text-white p-2 bg-teal-400 rounded-full" viewBox="0 0 24 24">
          <path d="M12 2L2 7l10 5 10-5-10-5zM2 20l10 5 10-5M2 12l10 5 10-5"></path>
        </svg>
        <span class="ml-3 text-xl">HRM</span>
      </a>
      <nav class="md:mr-auto md:ml-4 md:py-1 md:pl-4 md:border-l md:border-blue-200	flex flex-wrap items-center text-base justify-center font-bold">
        <router-link :to="{name: 'profile.index'}" class="mr-5 hover:text-gray-800">Profil</router-link>
        <router-link :to="{name: 'unit.index'}" class="mr-5 hover:text-gray-800">Unit</router-link>
        <router-link :to="{name: 'employee.index'}" class="mr-5 hover:text-gray-800">Data Pegawai</router-link>
      </nav>
      <button class="inline-flex items-center bg-gray-200 border-0 py-1 px-3 focus:outline-none hover:bg-gray-300 rounded text-base mt-4 md:mt-0">Button
        <svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" class="w-4 h-4 ml-1" viewBox="0 0 24 24">
          <path d="M5 12h14M12 5l7 7-7 7"></path>
        </svg>
      </button>
    </div>
  </header>
</template>

<script>
export default {
  name: "Header"
}
</script>

<style scoped>

</style>