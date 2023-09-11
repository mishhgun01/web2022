<template>
  <div id="app">
    <div v-if="$router.currentRoute.path !== '/sign-in'">
      <HeaderComponent ref="header"/>
    </div>
    <div class="app_body">
      <router-view/>
    </div>
  </div>
</template>

<script>

import HeaderComponent from './components/HeaderComponent.vue';

export default {
  name: 'App',
  components: {
    HeaderComponent
  },
  created() {
    const path = document.cookie.split(";").find(el=> el.includes("lastPath=")).split("=")[1]
    if (path) {
      this.$router.push(path)
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
</style>
