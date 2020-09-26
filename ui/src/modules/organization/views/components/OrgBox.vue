<template>
  <div class="flex flex-col justify-center items-center">
    <div class="flex border text-left p-2 rounded">
      <div class="flex-1">
        <slot :node-data="datasource">
          <div class="text-lg">
            {{ datasource.title }}
          </div>
          <div class="text-sm">
            {{ datasource.description }}
          </div>
        </slot>
      </div>
      <div class="relative">
        <div class="border hover:border-blue-500 rounded-sm hover:border-blue-300" @click="expandMenu">
          <svg class="h-4 fill-current text-gray-300 hover:text-blue-500 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M4 12a2 2 0 1 1 0-4 2 2 0 0 1 0 4zm6 0a2 2 0 1 1 0-4 2 2 0 0 1 0 4zm6 0a2 2 0 1 1 0-4 2 2 0 0 1 0 4z"/></svg>
        </div>
        <div class="absolute top-10 right--0 bg-white w-24" v-if="expand">
          <ul class="text-xs border bg-gray-100 rounded-sm">
            <li class="border-b p-1 hover:bg-blue-500 hover:text-white cursor-pointer" @click="addParent">Add Parent</li>
            <li class="border-b p-1 hover:bg-blue-500 hover:text-white cursor-pointer" @click="addChildren">Add Children</li>
            <li class="p-1 text-red-700 hover:bg-red-600 hover:text-white cursor-pointer" @click="deleteNode">Delete</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
  <ul class="flex flex-row justify-center mt-10" v-if="datasource.parent_id">
    <div class="-mt-10 border-l-2 absolute h-10 border-gray-400"></div>
    <li class="relative p-6" v-for="child in datasource.children" :key="child.id">
      <div class="border-t-2 border-l-2 absolute h-6 border-gray-400 top-0" style="left: 50%; right: 0px;"></div>
      <OrgBox :datasource="child"></OrgBox>
    </li>
  </ul>
</template>

<script>

export default {
  name: "OrgBox",
  props: [
    'datasource'
  ],
  data(){
    return{
      expand: false
    }
  },
  methods:{
    expandMenu(){
      this.expand = this.expand === false;
    },
    addParent(){
      this.$emit('add-parent', {
        id: this.id
      })
    },
    addChildren(){
      this.$emit('add-children', {
        id: this.id
      })
    },
    deleteNode(){
      this.$emit('delete-node', {
        id: this.id
      })
    }
  }
}
</script>

<style scoped>

</style>
