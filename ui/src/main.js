import { createApp } from 'vue'

import Default from "./layouts/Default.vue";
import Blank from "./layouts/Blank.vue";

import router from './router'
// import store from "./store"

import Vuelidate from 'vuelidate'
import App from './App.vue'

import 'tailwindcss/tailwind.css'
import './index.css'

createApp(App)
    .use(Vuelidate)
    .use(router)
    // .use(store)
    .component('default-layout', Default)
    .component('blank-layout', Blank)
    .mount('#app')
