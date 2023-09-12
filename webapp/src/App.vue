<template>
  <div id="app">
    <div v-if="$router.currentRoute.path !== '/sign-in'">
      <HeaderComponent ref="header"/>
    </div>
    <div class="btn-group">
      <div  class="ml-2">
        <img src="../src/assets/flag-ru.svg" alt="flag"/>
      </div>
      <InputSwitch v-model="lang"/>
      <div  class="ml-2">
        <img src="../src/assets/flag-en.svg" alt="flag"/>
      </div>
    </div>
    <div class="app_body">
      <router-view/>
    </div>
  </div>
</template>

<script>

import HeaderComponent from './components/HeaderComponent.vue';
import InputSwitch from 'primevue/inputswitch';
import {i18n} from "@/plugins/internationalization";

export default {
  name: 'App',
  components: {
    HeaderComponent,
    InputSwitch
  },
  data() {
    return {
      lang: false
    }
  },
  watch:{
    lang(val) {
      i18n.locale = val ? "en" : "ru"
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}


body {
  margin: 0;
}

.app_body{
  display: flex;
  justify-content: center;
  margin-top: 100px;
}


.btn-group {
  display: flex;
  justify-content: space-between;
  flex-direction: row;
  width: 120px;
  position: fixed;
  bottom: 5px;
  left: 90%;
  right: 10%;
  transform: translatex(-5px);
  z-index: 1000;
  opacity: 1;
}
</style>
