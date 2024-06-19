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
        <v-btn class="button-space" @click=registerOrDelCourse(course.ID,course.CourseName,course.courseFlg) color="primary">
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
const emit = defineEmits(['reload'])

const store = useCounterStore()
const props=defineProps(['course','buttonColor'])

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
  console.log(id)
  console.log(name)
  console.log('registerCourse')
  const token=store.studentState.csrftoken
  try{
    const response = await axios.post(`${_BASE_URL_}api/auth/course`,res,{
      headers: {
        'X-CSRF-Token': token
      }
    })
    if(response.data.message){
      console.log('Register successful')
    const c=store.getCourseById(id)
      if(!c){
        //if it doesn't exist,save data to store
        store.$patch((state) => state.studentState.studentCourses.push(course))
      }
    }
    console.log(response)
  }catch(error){
    console.error('Register failed:', error.response.data);
    alert(`${name}登録が失敗しました。もう一回お試しください。`)
    dialog.value = false
  }

}

//deleteCourse
const deleteCourse = async(id)=>{
  console.log(id)
  console.log('deleteCourse')
  const token=store.studentState.csrftoken
  try{
    await axios.delete(`${_BASE_URL_}api/auth/course/${id}`,{
      headers: {
        'X-CSRF-Token': token
      }
    })
  .then(response=>console.log('Delete successful'))
  }catch(error){
    console.error('Register failed:', error.response.data);
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

