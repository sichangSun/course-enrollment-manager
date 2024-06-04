<template>
    <v-container>
      <v-list lines="one"
        v-if="courseList.length > 0">
        <v-list-item
          v-for="course in courseList"
          :key="course.CourseID"
          :border="true"
        >
          <div>
            <v-list-item-title>{{ course.CourseName }}</v-list-item-title>
            <v-list-item-subtitle>{{ course.TeacherName }}</v-list-item-subtitle>
          </div>

          <v-list-item-action
          :end="true" class="list-button">
            <v-btn class="button-space" @click="$emit('getDetail',course.ID)">
              詳細
            </v-btn>
            <ConfirmBox
            :course="course"
            :buttonColor="getButtonColor(course)">
              <template v-slot:content>
                {{ course.CourseName }}この授業{{ course.courseDisplay }}してよろしでしょうか？
              </template>
          </ConfirmBox>
            <!-- <v-btn class="button-space" color="primary" @click="$emit('registerCourse',course.CourseID)">
              登録
            </v-btn> -->

          </v-list-item-action>
        </v-list-item>
      </v-list>
    </v-container>
</template>
<script setup>
import { reactive, ref } from 'vue'
import ConfirmBox from  './ConfirmBox.vue'
// import { useRouter } from 'vue-router';
// const router=useRouter()
const props=defineProps(['courseList','chooseFlg'])
const emit=defineEmits(['registerCourse'])

// const registerOrDelCourse = async(id,courseFlg)=>{
//   if(courseFlg==1){
//     //delete
//     emit('deleteCourse',id)

//   }else if(courseFlg==0){
//     //register
//     emit('registerCourse',id)
//   }else{
//     console.error('Attribute not found : courseFlg')
//     router.push({
//       name: 'ErrorPage'
//     })
//   }
// }

const getButtonColor = (course) => {
    return course.courseFlg === 1 ? 'red' : 'primary'
}


</script>
<style>
.option-info{
  text-align: right;
}
.list-button{
  float: right;
    margin-top: -10px;
    padding: 10px;
}
.button-space{
  margin: auto 5px;
}

</style>