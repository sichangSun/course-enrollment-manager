<template>
  <main>
      <div class="base-box">
        <div class="base-info">
        <v-card>
          <v-card-title>
            {{  student.studentName }}
          </v-card-title>
          <v-card-subtitle>{{ student.studentInfo }}</v-card-subtitle>
        </v-card>
      </div>
      <div class="base-regu">
        <h3>{{ btn.btnTitle }}</h3>
        <v-btn class="base-regu regu">{{btn.btnName}}</v-btn>
      </div>
      </div>
  </main>
</template>

<script setup>
import { reactive } from 'vue';

const props = defineProps(['student','btn'])

</script>
<style>
.base-box{
  /* margin-top: 90px;
    width: 600px;
    margin-left: 60px; */
    padding: 20px;
    margin: 90px 20px auto;
}
.base-info{
  width: 50%;
    margin: auto 14%;
}

.base-regu{
  position: absolute;
    margin: auto 65%;
    margin-top: -70px;
}
.regu{
  margin: 30px;
}
</style>
