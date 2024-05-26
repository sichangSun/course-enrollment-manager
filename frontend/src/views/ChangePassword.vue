<template>
<BackToStudentHome/>
  <ChangePassword
  @change-password="ChangePassword"
  class="change-password"/>
</template>


<script setup>
import ChangePassword from '../components/ChangePassword.vue'
import BackToStudentHome from '../components/BackToStudentHome.vue'


// const ChangePassword = async (state,$v) => {
//   if ($v.$invalid) {
//       // errMsg.value=''
//       return;
//     }
//     //set res
//     const res={

//     }
//     // call API
// }

</script>

<style>
.change-password{
  margin: 90px auto;
}
</style>