<template>
  <v-dialog
    v-model="dialog"
    max-width="600"
  >
  <!-- persistent -->
    <template v-slot:activator="{ props: activatorProps}">
      <v-btn v-bind="activatorProps" :color="buttonColor">
        {{ course.courseDisplay }}
      </v-btn>
    </template>
    <v-card class="card-box">
    <v-card-title>確認</v-card-title>
      <slot name="content"></slot>
      <template v-slot:actions>
        <v-spacer></v-spacer>
        <v-btn @click="dialog=false" class="button-space">
          キャンセル
        </v-btn>
        <v-btn class="button-space" @click=registerOrDelCourse(course.CourseID,course.CourseName,course.courseFlg) color="primary">
          確定
        </v-btn>
      </template>
    </v-card>
  </v-dialog>
</template>

<script setup>
import axios from '@/axios-config';
import { reactive, ref,onBeforeMount } from 'vue'
import { useCounterStore } from '@/stores/counter'
import { useRouter } from 'vue-router'
const emit = defineEmits(['reload','fetchdate'])

const store = useCounterStore()
const props=defineProps(['course','buttonColor'])
const router = useRouter()

let dialog= ref(false)
function registerOrDelCourse (id,name,courseFlg){
  if(courseFlg==1){
    //delete
    deleteCourse(id)

  }else if(courseFlg==0){
    //register
    registerCourse(id,name)
  }else{
    console.error('Attribute not found : courseFlg')
    router.push({
      name: 'ErrorPage'
    })
  }
}

// registerCourse
const registerCourse = async(id,name)=>{
  const res={
      courseid: id,
      coursename: name
    }
  // console.log(id)
  // console.log(name)
  console.log(`registerCourse: ${id}${name}`)
  //TODO: verify the save time can not register course

  //get csrf token
  const token=store.studentState.csrftoken
  try{
    const response = await axios.post(`${_BASE_URL_}api/auth/course`,res,{
      headers: {
        'X-CSRF-Token': token
      }
    })
    console.log(response)

    // sucessful
    if(response.data.message && response.data.message=='successful'){
      console.log('Register successful')
      const c=store.getCourseById(id)
      if(!c){
        //if it doesn't exist,save data to store
        const res2= await searchCourse(id)
        store.$patch((state) => state.studentState.studentCourses.push(res2))
        console.log(res2)
      }
    }

    dialog.value = false
    alert(`${name}登録ができました。`)
    emit('fetchdate')
  }catch(error){
    console.error('Register failed:', error.response.data)
    alert(`${name}登録が失敗しました。もう一回お試しください。`)
    dialog.value = false

  }
}
//scheduls validation
function schedulsValidate(id){
  console.log(id)

}

//Get registered course information
const searchCourse =async(id)=>{
  let res2={}
  try{
    const response= await axios.get(`${_BASE_URL_}api/courses/${id}`)
    res2=response.data.CourseDetail
    res2.Description=''
    // console.log(res2)
  }catch(error){
    // API getOneCourse failed
    if(error.response){
      console.error('Store update failed:', error.response.data)
    }
    //refresh
    emit('fetchdate')
  }
  return res2
}

//DeleteCourse
const deleteCourse = async(id)=>{
  console.log(`deleteCourse:${id}`)
   //get csrf token
  const token=store.studentState.csrftoken
  try{
    const response=await axios.delete(`${_BASE_URL_}api/auth/course/${id}`,{
      headers: {
        'X-CSRF-Token': token
      }
    })
    if(response.data.message && response.data.message=='successful'){
      console.log('Delete successful')
      store.deleteCourseFromStore(id)
      dialog.value = false
      alert(`削除ができました。`)
      emit('fetchdate')
    }
  }catch(error){
    console.error('Delete failed:', error.response.data)
    alert(`${name}削除が失敗しました。もう一回お試しください。`)
    dialog.value = false
  }
}

</script>

<style>
.card-box{
  height: auto;
}
/* .register-button {
  color:primary;
  color: white;
} */

/* .delete-button {
  background-color: red;
  color: white;
} */
</style>

