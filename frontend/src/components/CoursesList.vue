<template>
    <v-container>
      <v-list lines="one">
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
                <v-btn class="button-space" @click="getCourse(course.CourseID)">
                詳細
              </v-btn>
              <v-btn class="button-space" color="primary" @click="registerCourse(course.CourseID)">
                登録
              </v-btn>
            </v-list-item-action>
          </v-list-item>
      </v-list>
    </v-container>
</template>
<script setup>
import { reactive } from 'vue'

const courseList = reactive([
{
    'StudentID': 1,
    'CourseID': 1,
    'CourseName': '経済学(上)',
    'Semester': 1,
    'InstructorID': 1,
    'Credits': 2,
    'TeacherName': '飯田涼',
    'Position': '准教授',
    'Description': '',
    'Schedules': [{'DayOfWeek': 3, 'Period': 2}]
  },
  {
    'StudentID': 1,
    'CourseID': 2,
    'CourseName': '経済学(下)',
    'Semester': 2,
    'InstructorID': 1,
    'Credits': 2,
    'TeacherName': '飯田涼',
    'Position': '准教授',
    'Description': '',
    'Schedules': [{'DayOfWeek': 3, 'Period': 4}]
  },

])
function getCourse(int){
  console.log('getcourse')
}
function registerCourse(int){
  console.log('registerCourse')
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