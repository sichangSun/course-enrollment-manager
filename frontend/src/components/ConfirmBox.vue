<template>
  <v-dialog
    v-model="dialog"
    max-width="600"
  >
  <!-- persistent -->
    <template v-slot:activator="{ props: activatorProps}">
      <v-btn v-bind="activatorProps" color="primary">
        登録
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
        <v-btn class="button-space" @click="$emit('registerCourse',course.ID,course.courseFlg)" color="primary">
          確定
        </v-btn>
      </template>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { reactive, ref } from 'vue'
const props=defineProps(['course','dialog'])


let dialog= ref(false)

</script>

<style>
.card-box{
  height: auto;
}
</style>

