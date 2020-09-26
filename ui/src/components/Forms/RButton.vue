<template>
  <button :class="this.kelas">{{ text }}</button>
</template>

<script>
export default {
  name: "Button",
  props:[
    'text',
    'style',
    'type',
    'size'
  ],
  computed: {
    kelas(){

      let _class = ''
      _class = 'bg-blue-500 border border-blue-700 hover:bg-blue-700 hover:text-white'
      if(this.type == 'secondary'){
        _class = 'bg-indigo-600 border border-indigo-700 hover:bg-indigo-700 hover:text-white'
      }

      if(this.size == 'small'){
        _class += ' px-2 py-1 text-sm'
      }else{
        _class += ' px-4 py-1'
      }

      return 'text-gray-100 rounded-sm ' + _class
    }
  }
}
</script>

<style scoped>

</style>
