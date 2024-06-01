import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {

  const studentState = ref({
    studentName:'',
    studentInfo:{
      studentEmail:''
    },
    studentCourses:[
    ]
  })


  // updateStudentCourses
  function updateStudentCourses(courseList) {
    studentState.value.studentCourses = courseList;

  }

  function getCourseById(courseId) {
    return studentState.value.studentCourses.find(course => course.CourseID === courseId);
  }


  return { studentState, updateStudentCourses,getCourseById }
})
