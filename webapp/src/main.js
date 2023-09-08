import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

import PrimeVue from 'primevue/config';
Vue.use(PrimeVue)

import "primevue/resources/themes/saga-blue/theme.css"
import "primevue/resources/primevue.min.css"
import "primeicons/primeicons.css"
new Vue({
  render: h => h(App),
}).$mount('#app')
