<template>
  <h1>授業一覧</h1>
 <BackToStudentHome/>
  <CoursesList
    :courseList="courseList"
    @getDetail="getDetail"
    @registerCourse="registerCourse">
  </CoursesList>
</template>

<script setup>

  import { onBeforeMount, reactive, ref } from 'vue'
  import CoursesList from '../components/CoursesList.vue'
  import BackToStudentHome from '../components/BackToStudentHome.vue'
  import { useRouter } from 'vue-router'
  import axios from '../axios-config'

  const router = useRouter()
  let btn=reactive({
    btnTitle:'履修登録',
    btnName:'時間割に戻る'
  })
  const courseList=reactive([])

  onBeforeMount(async()=>{
    try{
      await axios.get(`${_BASE_URL_}api/courses`)
      .then(response=>{
        console.log(response.data.CourseList)

        response.data.CourseList.forEach(course => {
          courseList.push(course)
        });
      })
    }catch(error){
      if(error.response){
      console.error('Get courses failed:', error.response.data);
      }
      router.push({
        name: 'ErrorPage'
      })
    }
  })
  // get course detail
const getDetail = async(id)=>{
  console.log(id)
}
// registerCourse
const registerCourse = async(id)=>{
  console.log(id)


}



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
