<template>
    <div class="container mx-auto py-5 text-center">
      <Button
        text="Tambah Organisasi Baru"
        @click="this.showNewUnit"
      >
      </Button>
      <NewOrganization v-show="newUnit"></NewOrganization>
      <OrgBox></OrgBox>
    </div>
</template>

<script>
import NewOrganization from "./NewOrganization.vue";
import OrgBox from "./components/OrgBox.vue";
import Button from '../../../components/Forms/Button.vue'
export default {
  name: "Organization",
  components: {
    OrgBox,
    NewOrganization,
    Button
  },
  data(){
    return {
      newUnit: false
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
