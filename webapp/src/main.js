import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'


import PrimeVue from 'primevue/config';
Vue.use(PrimeVue)

import "primevue/resources/themes/saga-blue/theme.css"
import "primevue/resources/primevue.min.css"
import "primeicons/primeicons.css"
import {i18n} from "@/plugins/internationalization";

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
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)
Vue.use(VueRouter)
new Vue({
  i18n,
  router,
  render: h => h(App),
}).$mount('#app')
