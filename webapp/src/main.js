import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'


import PrimeVue from 'primevue/config';
Vue.use(PrimeVue)

import "primevue/resources/themes/saga-blue/theme.css"
import "primevue/resources/primevue.min.css"
import "primeicons/primeicons.css"

import http from './plugins/http';

Vue.use(http, {
  baseUrl: "https://localhost:8081"
})
Vue.prototype.$baseUrl = "https://localhost:8081"
export const url = Vue.prototype.$baseUrl

import routes from '@/routes'

const router = new VueRouter({
  routes
})
import ToastService from 'primevue/toastservice';
Vue.use(ToastService);
import { EmojiPickerPlugin } from 'vue-emoji-picker'
Vue.use(EmojiPickerPlugin)

Vue.use(VueRouter)
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
