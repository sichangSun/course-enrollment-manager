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
  function deleteCourseFromStore(courseId) {
    const index = studentState.value.studentCourses.findIndex(course => course.CourseID === courseId);
    if (index !== -1) {
      studentState.value.studentCourses.splice(index, 1);
    }
  }



  return { studentState, updateStudentCourses,getCourseById,updateToken,deleteCourseFromStore }
})
