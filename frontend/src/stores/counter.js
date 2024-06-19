import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {

  const studentState = ref({
    studentName:'',
    studentInfo:{
      studentEmail:''
    },
    studentCourses:[
    ],
    csrftoken:''
  })


  // updateStudentCourses
  function updateStudentCourses(courseList) {
    studentState.value.studentCourses = courseList;

  }

  function getCourseById(courseId) {
    return studentState.value.studentCourses.find(course => course.CourseID === courseId);
  }

  function updateToken(token) {
    studentState.value.csrftoken = token;
  }



  return { studentState, updateStudentCourses,getCourseById,updateToken }
})
