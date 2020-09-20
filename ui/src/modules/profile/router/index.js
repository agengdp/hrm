import Profile from '../views/Profile.vue'

export default [
    {
        path: '/profile',
        name: 'profile.index',
        component: Profile,
        meta: {
            title: 'Profile Perusahaan',
            layout: 'default'
        }
    }
]