<template>
  <h1>授業一覧</h1>
 <BackToStudentHome/>
  <CoursesList
    :courseList="courseList"
    @getDetail="getDetail"
    @fetchdateFromComfirm="fetchData"
    >
  </CoursesList>
</template>

<script setup>

  import { onBeforeMount, reactive, onUpdated,ref } from 'vue'
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
  // const res=reactive([])
  onBeforeMount(fetchData)
  // onUpdated(fetchData)

  async function fetchData() {
    let res=[]
    if(courseList.length != 0){
      courseList.splice(0)
    }
    try{
      const response= await axios.get(`${_BASE_URL_}api/courses`)

      response.data.CourseList.forEach(course=> {
        res.push(course)
      })
      console.log(response)

    }catch(error){
      router.push({
        name: 'ErrorPage'
      })
      console.error('Get courses failed:', error.response.data)
    }
    const processedCourses=res.map(course => {
      const c = store.getCourseById(course.CourseID)
      if(c){
        return { ...course, courseFlg: 1 ,courseDisplay:'削除'};
      }else{
        return { ...course, courseFlg: 0 ,courseDisplay:'登録'};
      }
    });
    courseList.push(...processedCourses);
    console.log(courseList)
  }
  // get course detail
const getDetail = async(id)=>{
  console.log(id)
  router.push({
    name:'CourseDetail',
    params: {
    courseId: id
  }
  })

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
