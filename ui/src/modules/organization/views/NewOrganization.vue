<template>
    <div class="bg-gray-200 p-3 pt-5 mt-2 w-1/3 mx-auto text-left border border-gray-400 rounded-sm">
      <Input
        placeholder="blablalba"
        label="Nama Organisasi"
      ></Input>
      <TextArea
        placeholder="haha"
        label="Deskripsi"
      ></TextArea>
      <div class="text-right">
        <Button
          text="Add"
          type="secondary"
          size="small"
        ></Button>
      </div>
    </div>
</template>

<script>

import Input from '../../../components/Forms/Input.vue'
import TextArea from '../../../components/Forms/TextArea.vue'
import Button from '../../../components/Forms/Button.vue'

export default {
  name: "NewOrganization",
  components: {
    Input,
    Button,
    TextArea
  }
}
</script>

<style scoped>

</style>
