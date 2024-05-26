  <template>
    <v-app>
        <v-container>
            <v-row>
                <h3>時間割</h3>
            </v-row>
            <!-- Header Row: Weekday-->
            <v-row>
                <v-col cols="2"></v-col>
                <v-col v-for="day in days" :key="day.index" cols="2">{{ day }}</v-col>
            </v-row>
            <!-- Main Rows: TimeSlots-->
            <v-row v-for="timeSlot in timeSlots" :key="timeSlot">
                <v-col cols="2">{{ timeSlot.name }}
                    <span>{{ timeSlot.timeFrame }}</span>
                </v-col>
                <!-- Course List -->
                <v-col v-for="day in days" :key="day" cols="2">
                    <v-card class="pa-2 font-in-list" outlined>
                        <v-card-title v-if="gridData[timeSlot.name] && gridData[timeSlot.name][day]">
                          <router-link to="/detail">{{ gridData[timeSlot.name][day] }}</router-link>
                        </v-card-title>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </v-app>

  </template>
  <script setup>
  import { reactive } from 'vue'

  const props = defineProps(['days','timeSlots','gridData'])

  </script>

  <style>
  .v-container {
    max-width: 1200px;
  }
  .v-row {
    margin-bottom: 10px;
  }
  .v-col {
    text-align: center;
    border: 1px solid #ccc;
    padding: 10px;
    height: 100px;

  }
  .v-card {
    height: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .v-card-title {
    font-size: 10px;
  }
  h3{
    margin-top: -30px;
  }
  </style>
