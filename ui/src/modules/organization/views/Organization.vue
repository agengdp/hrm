<template>
    <div class="container mx-auto py-5 text-center">
      <RButton
        text="Tambah Organisasi Baru"
        @click="this.showNewUnit"
      >
      </RButton>
      <NewOrganization v-show="newUnit"></NewOrganization>
    </div>
</template>

<script>
import NewOrganization from "./NewOrganization.vue";
import RButton from '../../../components/Forms/RButton.vue'
export default {
  name: "Organization",
  components: {
    NewOrganization,
    RButton
  },
  data(){
    return {
      newUnit: true
    }
  },
  methods: {
    showNewUnit(){
      if(this.newUnit){
        return this.newUnit = false
      }
      return this.newUnit = true
    }
  }
}
</script>

<style scoped>

</style>
