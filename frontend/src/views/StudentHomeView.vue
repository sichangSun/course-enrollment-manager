<template>
  <main>
    <v-btn @click="toChangePasswordPage" class="password-change-btn">パスワード変更</v-btn>

    <BaseInfo
      :student="student"
      :btn="btn"/>
    <MyCourseSchedule
      class="schedule"
      :days="days"
      :timeSlots="timeSlots"
      :gridData="gridData"/>

  </main>
</template>

<script setup>
import { reactive, ref ,onBeforeMount,} from 'vue'
import MyCourseSchedule from '../components/MyCourseSchedule.vue'
import BaseInfo from '../components/BaseInfo.vue'
import { useRouter } from 'vue-router'
import axios from '../axios-config'
import { useCounterStore } from '@/stores/counter'

const store = useCounterStore()
const router = useRouter()

onBeforeMount(fetchData)
// onBeforeUpdate(fetchData)

  // Todo need to add get studentInfo API
  let student=reactive({
     studentName:store.studentState.studentName,
     studentInfo:'student Info'

  })

      const days = reactive(['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'])

      const timeSlots = [
        {name:'First',timeFrame:'9:00 ~ 10:30'},
        {name:'Second',timeFrame:'10:40 ~ 12:10'},
        {name:'Third',timeFrame:'13:10 ~ 14:40'},
        {name:'Fourth',timeFrame:'14:50 ~ 16:20'},
        {name:'Fifth',timeFrame:'16:30 ~ 18:00'}
    ];
      let gridData = reactive({
        First: { Monday: {name:'',id:''}, Tuesday: {name:'',id:''}, Wednesday: {name:'',id:''}, Thursday: {name:'',id:''}, Friday: {name:'',id:''}},
        Second: { Monday: {name:'',id:''}, Tuesday: {name:'',id:''}, Wednesday: {name:'',id:''}, Thursday: {name:'',id:''}, Friday: {name:'',id:''}},
        Third: { Monday: {name:'',id:''}, Tuesday: {name:'',id:''}, Wednesday: {name:'',id:''}, Thursday: {name:'',id:''}, Friday: {name:'',id:''}},
        Fourth: { Monday: {name:'',id:''}, Tuesday: {name:'',id:''}, Wednesday: {name:'',id:''}, Thursday: {name:'',id:''}, Friday: {name:'',id:''} },
        Fifth: { Monday: {name:'',id:''}, Tuesday: {name:'',id:''}, Wednesday: {name:'',id:''}, Thursday: {name:'',id:''}, Friday: {name:'',id:''}}
      });

      let btn=reactive({
          btnTitle:'履修登録',
          btnName:'授業一覧'
      })

  async function fetchData(){
    let res ={}
    try{
      const response = await axios.get(`${_BASE_URL_}api/auth/courses`)
        console.log('Get courses successful');
        // console.log(response.data)
        if(response.data.message && response.data.message=='successful but 0 row'){
          return
        }
        res = response.data
    }catch(error){
      if(error.response){
      console.error('Get courses failed:', error.response.data)
      }
      router.push({
        name: 'ErrorPage'
      })
    }

    // dayMap & periodMap
    const dayMap = {1: 'Monday', 2: 'Tuesday', 3: 'Wednesday', 4: 'Thursday', 5: 'Friday'}
    const periodMap = {1: 'First', 2: 'Second', 3: 'Third', 4: 'Fourth', 5: 'Fifth'}
    // gridData & store setting
    res.courses.CoursesList.forEach(course => {
      //find course by Id in store
      const c=store.getCourseById(course.CourseID)
      if(!c){
        //if it doesn't exist,save data to store
        store.$patch((state) => state.studentState.studentCourses.push(course))
      }
      course.Schedules.forEach(schedule => {
        const day = dayMap[schedule.DayOfWeek]  //DayOfWeek
        const period = periodMap[schedule.Period] // Period
          gridData[period][day].name = course.CourseName
          gridData[period][day].id = course.CourseID
      })
    })
    //Save csrftoken
    store.updateToken(res.csrfToken)
  }
  function toChangePasswordPage(){
    router.push({
    name: 'ChangePassword'
  })
  }
</script>
<style>
.schedule{
  position: relative;
  top: 30px;
}
.password-change-btn{
  margin: 90px 10px auto;
}
</style>
