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
import { reactive, ref ,onBeforeMount} from 'vue'
import MyCourseSchedule from '../components/MyCourseSchedule.vue'
import BaseInfo from '../components/BaseInfo.vue'
import { useRouter } from 'vue-router'
import axios from '../axios-config'
const router = useRouter();

  // Todo need to add get studentInfo API
  let student=reactive({
     studentName:'student name',
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
        First: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''},
        Second: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''},
        Third: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''},
        Fourth: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: '' },
        Fifth: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''}
      });

      let btn=reactive({
          btnTitle:'履修登録',
          btnName:'授業一覧'
      })

  onBeforeMount(async () => {
    let res ={};
    try{
      await axios.get(`${_BASE_URL_}api/auth/courses`)
      .then(response => {
        console.log('Get courses successful');
        console.log(response.data)
        if(response.data.message && response.data.message=='successful but 0 row'){
          return
        }
        res = response.data
      })
    }catch(error){
      if(error.response){
      console.error('Get courses failed:', error.response.data);
      }
      router.push({
        name: 'ErrorPage'
      })
    }
    //console.log(res)

    // dayMap & periodMap
    const dayMap = {1: 'Monday', 2: 'Tuesday', 3: 'Wednesday', 4: 'Thursday', 5: 'Friday'};
    const periodMap = {1: 'First', 2: 'Second', 3: 'Third', 4: 'Fourth', 5: 'Fifth'};
    // gridData setting
    res.CoursesList.forEach(course => {
      course.Schedules.forEach(schedule => {
        const day = dayMap[schedule.DayOfWeek];   //DayOfWeek
        const period = periodMap[schedule.Period]; // Period
          gridData[period][day] = course.CourseName;
      });
    });
  })
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
