<template>
  <div>
    <div class="bg-gray-200 p-3 pt-5 mt-2 w-1/3 mx-auto text-left border border-gray-400 rounded-sm">
      <RInput
        placeholder="blablalba"
        label="Nama Organisasi"
      ></RInput>
      <RTextArea
        placeholder="haha"
        label="Deskripsi"
      ></RTextArea>
      <div class="text-right">
        <RButton
          text="Add"
          type="secondary"
          size="small"
        ></RButton>
      </div>
    </div>
    <div class="flex flex-col justify-center items-center">
      <div class="text-center">
        <OrgBox
          :datasource="organizations"
        ></OrgBox>
      </div>
    </div>
  </div>
</template>

<script>

import RInput from '../../../components/Forms/RInput.vue'
import RTextArea from '../../../components/Forms/RTextArea.vue'
import RButton from '../../../components/Forms/RButton.vue'
import OrgBox from "./components/OrgBox.vue";

export default {
  name: "NewOrganization",
  components: {
    RInput,
    RButton,
    RTextArea,
    OrgBox
  },
  data(){
    return{
      'organizations': [
        {'id': '1', 'title': 'Direktur Utama', 'description': 'Pemimpin Perusahaan'},
        {'id': '2', 'parent_id': '1', 'title': 'Wadir Umum', 'description': 'Wakil Direktur bagian umum'},
        {'id': '3', 'parent_id': '1', 'title': 'Wadir Pelayanan', 'description': 'Wakil Direktur bagian pelayanan'},
        {'id': '4', 'parent_id': '2', 'title': 'Keuangan', 'description': 'Bagian Keuangan'},
        {'id': '5', 'parent_id': '2', 'title': 'IT', 'description': 'Bagian IT'},
        {'id': '6', 'parent_id': '5', 'title': 'Software', 'description': 'Bagian Software'},
        {'id': '7', 'parent_id': '5', 'title': 'Hardware', 'description': 'Bagian Hardware'},
        {'id': '8', 'parent_id': '5', 'title': 'Network', 'description': 'Bagian Network'},
      ]
    }
  }
}
</script>

<style scoped>

</style>
