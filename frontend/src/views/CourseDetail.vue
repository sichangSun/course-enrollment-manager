<template>
  <h1>授業詳細</h1>
<BackToStudentHome/>
  <CourseDescription/>
</template>

<script setup>
import CourseDescription from '../components/CourseDescription.vue'
import { useRouter } from 'vue-router'
import BackToStudentHome from '../components/BackToStudentHome.vue'

const router = useRouter()

// function backToStudentHome(){
//   router.push({
//     name: 'StudentHome'
//   })
// }



</script>
<style>
h1{
  margin: 90px auto 3px auto;
}
</style>