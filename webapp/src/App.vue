<template>
  <div id="app">
    <div>
      <HeaderComponent ref="header"/>
    </div>
    <MainView/>
  </div>
</template>

<script>

import HeaderComponent from './components/HeaderComponent.vue';
import MainView from './views/MainView.vue';

export default {
  name: 'App',
  components: {
    MainView,
    HeaderComponent
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
  margin-top: 60px;
}


body {
  margin: 0;
}
</style>
