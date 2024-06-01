import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  
  const studentState = ref({
    studentName:'',
    studentInfo:{
      studentEmail:''
    },
    studentCourse:[
    ]
  })


  // updateStudentCourses
  function updateStudentCourses(courses) {

  }


  return { studentState, updateStudentCourses }
})
