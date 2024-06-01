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
      <slot name="courseName"></slot>
      <template v-slot:actions>
        <v-spacer></v-spacer>
        <v-btn @click="dialog=flase" class="button-space">
          キャンセル
        </v-btn>
        <v-btn class="button-space" @click="$emit('registerOrDelCourse',course.ID,course.courseFlg)" color="primary">
          確定
        </v-btn>
      </template>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { reactive, ref,onBeforeMount } from 'vue'
const props=defineProps(['course','dialog','buttonColor'])

let dialog= ref(false)

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

