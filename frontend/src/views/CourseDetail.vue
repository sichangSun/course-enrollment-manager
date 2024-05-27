<template>
  <h1>授業詳細</h1>
  <v-btn class="btn-back">
    時間割に戻る
  </v-btn>
  <CourseDescription/>
</template>

<script setup>
import CourseDescription from "../components/CourseDescription.vue";



</script>
<style>
h1{
  margin: 90px auto 3px auto;
}
.btn-back{
  margin: 100px 40px;
  position: absolute!important;
}
</style>