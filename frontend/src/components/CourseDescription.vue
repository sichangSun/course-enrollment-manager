<template>
  <v-container v-if="course && schedulesShow">
    <div class="base-header-info">
      <h4>{{ course.CourseName  }}</h4>
    </div>
    <v-row class="base-info-box">
      <v-col cols="3" >
        <div>担当教員：</div>
        <v-divider></v-divider>
        <div>{{  course.Position }}  {{ course.TeacherName }}</div>
      </v-col>
      <v-col cols="3" >
        <div>実施時間：</div>
        <v-divider></v-divider>
        <!-- todo time fomart -->
        <div v-for="time in schedulesShow" key="index">
          {{ time }}
        </div>
      </v-col>
      <v-col cols="3" >
        <div>単位数：</div>
        <v-divider></v-divider>
        <div>{{ course.Semester }}</div>
      </v-col>
      <v-col cols="3" >
        <div>分類：</div>
        <v-divider></v-divider>
        <div>xxx</div>
      </v-col>
    </v-row>
    <v-row>
      <v-col class="col-text-area">
        <v-card outlined class="text-area">
          <v-card-title>
            概要
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            {{ course.Description }}

          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row >
      <div class="btn">
      <ConfirmBox
      :course="course"
      :buttonColor="getButtonColor(course)"
      @fetchdate="handleDataUpdated">
        <template v-slot:content>
          {{ course.CourseName }}この授業{{ course.courseDisplay }}してよろしでしょうか？
        </template>
      </ConfirmBox>
    </div>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import ConfirmBox from  './ConfirmBox.vue'

const props=defineProps(['course','schedulesShow'])
const emit=defineEmits(['fetchdateFromComfirm'])

function handleDataUpdated() {
  emit('fetchdateFromComfirm')
}


const getButtonColor = (course) => {
  return course.courseFlg === 1 ? 'red' : 'primary'
}


</script>
<style>
.v-container{
  padding: 5px;
}
.col-text-area{
  height: 610px;
  margin: auto;
  border: none;
}
.text-area{
  height: 600px;
  margin: auto;
}
.base-info-box{
  margin-top: 3px;
}
.base-header-info{
  margin: 10px auto;
}
.btn{
  margin: auto;
    margin-top: 10px;
    /* width: 20%; */
}
</style>