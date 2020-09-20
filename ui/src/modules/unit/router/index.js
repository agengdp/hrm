import Unit from '../views/Unit.vue'
import NewUnit from '../views/NewUnit.vue'

export default [
    {
        path: '/unit',
        name: 'unit.index',
        component: Unit,
        meta: {
            title: 'Unit',
            layout: 'default'
        }
    },{
        path: '/unit/new',
        name: 'unit.new',
        component: NewUnit,
        meta:{
            title: 'Tambah Unit Baru',
            layout: 'default'
        }
    }
]