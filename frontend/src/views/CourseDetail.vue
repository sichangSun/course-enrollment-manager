<template>
  <h1>授業詳細</h1>
<BackToStudentHome/>
  <CourseDescription
  :course="course"
  :schedulesShow="schedulesShow"
  @fetchdateFromComfirm="fetchData"/>
</template>

<script setup>
import { reactive, ref,watchEffect,onBeforeMount } from 'vue'
import CourseDescription from '../components/CourseDescription.vue'
import { useRouter } from 'vue-router'
import { useRoute } from 'vue-router'
import BackToStudentHome from '../components/BackToStudentHome.vue'
import axios from '@/axios-config'
import { useCounterStore } from '@/stores/counter'

const store = useCounterStore()

const router = useRouter()
const route =useRoute()
// const props =defineProps(['course'])
let course=ref({})
let schedulesShow=reactive([])

onBeforeMount(fetchData)

async function fetchData(){
  const courseId = route.params.courseId
  console.log(courseId)
  let res ={}

 try{
  const response= await axios.get(`${_BASE_URL_}api/courses/${courseId}`)

    console.log(response.data.CourseDetail)
    if(response.data.CourseDetail){
      res = response.data.CourseDetail
    }
  }
  catch(error){
    console.error('Get detail failed:', error.response.data)
    router.push({
      name: 'ErrorPage'
    })

  }
  // console.log(res)
  const c=store.getCourseById(res.CourseID)

  if(c){
    res.courseFlg = 1
    res.courseDisplay = '削除'
  }else{
    res.courseFlg = 0
    res.courseDisplay = '登録'
  }
  course.value=res
  console.log(course)

  const days={1:'月曜日' , 2:'火曜日',  3:'水曜日',  4:'木曜日',  5:'金曜日'}
  const peridos={ 1:'一限', 2:'二限',  3:'三限',  4:'四限', 5:'五限' }

  if(res.Schedules != 0){
    res.Schedules.splice(0)
  }
  res.Schedules.forEach(trim=>{
    const day=days[trim.Period]
    const perido=peridos[trim.DayOfWeek]
    const schedule=`${day}  ${perido}`
    schedulesShow.push(schedule)
  })

}

</script>
<style>
h1{
  margin: 90px auto 3px auto;
}
</style>