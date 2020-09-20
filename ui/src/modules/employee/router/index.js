import Employee from '../views/Employee.vue'

export default [
    {
        path: '/employee',
        name: 'employee.index',
        component: Employee,
        meta: {
            title: 'Data Pegawai',
            layout: 'default'
        }
    }
]