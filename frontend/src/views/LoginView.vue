<template>
  <v-app>
  <Login class="login"></Login>
  </v-app>

</template>

<script setup>
  import Login from  '../components/Login.vue'
  // import { VApp, VBtn } from 'vuetify/components';
  import { useRouter } from 'vue-router'

  const router = useRouter()

    // function toCourse(){
    //   console.log('OK')
    //   $router.push({ name:'courseList'})
    // }

</script>
<style>
.login{
  margin: 90px auto;
}

</style>