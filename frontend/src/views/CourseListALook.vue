<template>
  <h1>授業一覧</h1>
 <BackToStudentHome/>
<!-- totdo frezz search  -->
  <CoursesList></CoursesList>
</template>

<script setup>

import { reactive } from 'vue'
import CoursesList from '../components/CoursesList.vue'
import BackToStudentHome from '../components/BackToStudentHome.vue'

let btn=reactive({
          btnTitle:'履修登録',
          btnName:'時間割に戻る'
      })

</script>
<style>
h1{
  margin: 90px auto;
}
/* .btn-back{
  margin: 100px 40px;
  position: absolute!important;
} */
</style>
