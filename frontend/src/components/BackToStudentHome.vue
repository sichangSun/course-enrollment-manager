<template>
  <v-btn class="btn-back" @click="backToStudentHome">
    時間割に戻る
  </v-btn>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

function backToStudentHome(){
  router.push({
    name: 'StudentHome'
  })
}



</script>
<style>

.btn-back{
  margin: 100px 40px;
  position: absolute!important;
}
</style>