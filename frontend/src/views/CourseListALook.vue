<template>
  <h1>授業一覧</h1>
 <BackToStudentHome/>
  <CoursesList
    :courseList="paginatedData"
    @getDetail="getDetail"
    @fetchdateFromComfirm="fetchData"
    >
  </CoursesList>
  <div class="text-center">
    <v-container class="max-width">
      <v-pagination
        v-model="page"
        :length="pageLength"
        :total-visible="totalVisible"
        class="my-4"
      ></v-pagination>
    </v-container>
  </div>
</template>

<script setup>

  import { onBeforeMount, reactive, onUpdated,ref, computed } from 'vue'
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
  onBeforeMount(fetchData)
  const page = ref(1)
  const pageItems=10
  const totalVisible=7


  async function fetchData() {
    let res=[]
    //clear courseList to refresh
    if(courseList.length != 0){
      courseList.splice(0)
    }
    //call API
    try{
      const response= await axios.get(`${_BASE_URL_}api/courses`)
      response.data.CourseList.forEach(course=> {
        res.push(course)
      })
      // console.log(response)
    }catch(error){
      router.push({
        name: 'ErrorPage'
      })
      console.error('Get courses failed:', error.response.data)
    }
    //edit button
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

  const paginatedData=computed(()=>{
    const start=(page.value-1)*pageItems
    const end=start+pageItems
    return courseList.slice(start, end)
  })

  const pageLength = computed(() => {
    return Math.ceil(courseList.length / pageItems)
  })
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
