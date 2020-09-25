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
    <div>
      <OrgBox
        :datasource="organizations"
      ></OrgBox>
    </div>
</template>

<script>

import Input from '../../../components/Forms/Input.vue'
import TextArea from '../../../components/Forms/TextArea.vue'
import Button from '../../../components/Forms/Button.vue'
import OrgBox from "./components/OrgBox.vue";

export default {
  name: "NewOrganization",
  components: {
    Input,
    Button,
    TextArea,
    OrgBox
  },
  data(){
    return{
      'organizations': {
        'id': '1',
        'title': 'Direktur Utama',
        'description': 'Pemimpin Perusahaan',
        'children': [
          {
            'id': '2',
            'title': 'Wadir Umum',
            'description': 'Wakil Direktur bagian umum',
            'children': [
              {
                'id': '4',
                'title': 'Keuangan',
                'description': 'Bagian Keuangan'
              },
              {
                'id': '5',
                'title': 'IT',
                'description': 'Bagian IT',
                'children': [
                  {
                    'id': '6',
                    'title': 'Hardware',
                    'description': 'Bagian IT Hardware',
                  },
                  {
                    'id': '7',
                    'title': 'Software',
                    'description': 'Bagian Software',
                  }
                ]
              }
            ]
          },
          {
            'id': '3',
            'title': 'Wadir Pelayanan',
            'description': 'Wakil Direktur bagian pelayanan'
          }
        ]
      }
    }
  }
}
</script>

<style scoped>

</style>
