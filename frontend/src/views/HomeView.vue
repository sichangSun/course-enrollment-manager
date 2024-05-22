<script setup>
import MyCourseSchedule from '../components/MyCourseSchedule.vue'
</script>

<template>
  <main>
    <v-app>
      <v-app-bar>
      <v-app-bar-title>
        Application
      </v-app-bar-title>
    </v-app-bar>
    <MyCourseSchedule class="schedule"/>
  </v-app>
  </main>
</template>
<script>
import { defineComponent } from 'vue';
import { VApp, VBtn } from 'vuetify/components';

export default defineComponent({
  components: {
    VApp,
    VBtn
  }
});
</script>
<style>
.schedule{
  position: relative;
  top: 30px;
}
</style>
