<template>
  <div id="app">
    <div v-if="auth">
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
    HeaderComponent,
  },
  data() {
    return {
      auth: false
    }
  },
  created() {
    console.log(this.$router)
  },
  methods: {
    loggedIn(param) {
      console.log(param)
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
}

.MainView {
  margin-top: 100px;
}
</style>
