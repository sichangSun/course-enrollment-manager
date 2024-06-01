<template>
  <h1>授業一覧</h1>
 <BackToStudentHome/>
  <CoursesList
    :courseList="courseList"
    @getDetail="getDetail"
    @deleteCourse="deleteCourse"
    @registerCourse="registerCourse">
  </CoursesList>
</template>

<script setup>

  import { onBeforeMount, reactive, ref } from 'vue'
  import CoursesList from '../components/CoursesList.vue'
  import BackToStudentHome from '../components/BackToStudentHome.vue'
  import { useRouter } from 'vue-router'
  import axios from '../axios-config'
  import { useCounterStore } from '@/stores/counter'

  const store = useCounterStore()

  const router = useRouter()
  // const btn=reactive({
  //   btnTitle:'履修登録',
  //   btnName:'時間割に戻る'
  // })

  const courseList=reactive([])
  const res=reactive([])



  onBeforeMount(async()=>{
    try{
      await axios.get(`${_BASE_URL_}api/courses`)
      .then(response=>{
        response.data.CourseList.forEach(course=> {
          res.push(course)
        });
        console.log(`All courseList is`)
        console.log(res)

      })
    }catch(error){
      console.error('Get courses failed:', error.response.data)
      router.push({
        name: 'ErrorPage'
      })
    }
    const processedCourses=res.map(course => {
      const c = store.getCourseById(course.ID)
      if(c){
        return { ...course, courseFlg: 1 ,courseDisplay:'削除'};
      }else{
        return { ...course, courseFlg: 0 ,courseDisplay:'登録'};
      }
    });
    courseList.push(...processedCourses);
    console.log(`CoursesList compoment courseList is`)
    console.log(courseList)


  })
  // get course detail
const getDetail = async(id)=>{
  console.log(id)
}
// registerCourse
const registerCourse = async(id)=>{
  console.log(id)
  console.log('registerCourse')

}

const deleteCourse = async(id)=>{
  console.log(id)
  console.log('deleteCourse')

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
