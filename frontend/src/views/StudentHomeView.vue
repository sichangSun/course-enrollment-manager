<template>
  <main>

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
import { reactive, ref } from 'vue'
import MyCourseSchedule from '../components/MyCourseSchedule.vue'
import BaseInfo from '../components/BaseInfo.vue'
import { useRouter } from 'vue-router'
  let student=reactive({
     studentName:'huahua',
     studentInfo:'経済学部xx学科xx専攻'

  })

      const days = reactive(['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'])
    //   const timeFrames=['9:00 ~ 10:30','10:40 ~ 12:10','13:10 ~ 14:40','14:50 ~ 16:20','16:30 ~ 18:00'];
      const timeSlots = [
        {name:'First',timeFrame:'9:00 ~ 10:30'},
        {name:'Second',timeFrame:'10:40 ~ 12:10'},
        {name:'Third',timeFrame:'13:10 ~ 14:40'},
        {name:'Fourth',timeFrame:'14:50 ~ 16:20'},
        {name:'Fifth',timeFrame:'16:30 ~ 18:00'}
    ];
      let gridData = reactive({
        First: { Monday: 'Math', Tuesday: '', Wednesday: 'English', Thursday: '', Friday: ''},
        Second: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''},
        Third: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''},
        Fourth: { Monday: 'Gym', Tuesday: 'Art', Wednesday: '', Thursday: '', Friday: '' },
        Fifth: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''}
      });

      let btn=reactive({
          btnTitle:'履修登録',
          btnName:'授業一覧'
      })
  // const router = useRouter();
  //   function pushNew(){
  //   console.log('OK')
  //   router.push({ name:'login'})
  // }
</script>
<style>
.schedule{
  position: relative;
  top: 30px;
}
/* .basic-box{
    padding: 20px;
    margin: 90px 20px auto;
}
.basic-info{
  width: 50%;
    margin: auto 14%;
}

.basic-regu{
  position: absolute;
    margin: auto 65%;
    margin-top: -70px;
}
.regu{
  margin: 30px;
} */
</style>
