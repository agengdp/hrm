<template>
    <div class="container mx-auto py-5">
      <Button
        text="Tambah Unit Baru"
      >
        Tambah unit baru
      </Button>
      <NewUnit v-show="newUnit"></NewUnit>
    </div>
</template>

<script>
import NewUnit from "./NewUnit.vue";
import Button from '../../../components/Forms/Button.vue'
export default {
  name: "Unit",
  components: {
    NewUnit,
    Button
  },
  data(){
    return {
      newUnit: true
    }
  }
}
</script>

<style scoped>

</style>
