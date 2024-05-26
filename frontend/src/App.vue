<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <!-- <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" /> -->
    <!-- <h1>CourseEnrollmentManager</h1> -->
    <div >
      <nav>
        <RouterLink to="/"></RouterLink>
      </nav>
    </div>
  </header>
<main>
  <v-app>
    <v-app-bar>
      <v-app-bar-title>
        学生履修授業登録システム
      </v-app-bar-title>
    <!-- <v-btn v-if="">PasswordChange</v-btn> -->
    <!-- <template v-slot:append></template> -->
    </v-app-bar>
    <RouterView />
  </v-app>
</main>


</template>

<style scoped>
</style>
