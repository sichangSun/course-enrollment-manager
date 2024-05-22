  <template>
    <v-app>
        <v-container>
            <!-- Header Row: Weekday-->
            <v-row>
                <v-col cols="2"></v-col>
                <v-col v-for="day in days" :key="index" cols="2">{{ day }}</v-col>
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
                            {{ gridData[timeSlot.name][day] }}
                        </v-card-title>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </v-app>

  </template>
  <script>
  import { reactive } from 'vue'
  export default {
    setup() {
      const days = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'];
    //   const timeFrames=['9:00 ~ 10:30','10:40 ~ 12:10','13:10 ~ 14:40','14:50 ~ 16:20','16:30 ~ 18:00'];
      const timeSlots = [
        {name:'First',timeFrame:'9:00 ~ 10:30'},
        {name:'Second',timeFrame:'10:40 ~ 12:10'},
        {name:'Third',timeFrame:'13:10 ~ 14:40'},
        {name:'Fourth',timeFrame:'14:50 ~ 16:20'},
        {name:'Fifth',timeFrame:'16:30 ~ 18:00'}
    ];
      let gridData = reactive({
        First: { Monday: 'Math', Tuesday: '', Wednesday: 'English', Thursday: '', Friday: ''},
        Second: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''},
        Third: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''},
        Fourth: { Monday: 'Gym', Tuesday: 'Art', Wednesday: '', Thursday: '', Friday: '' },
        Fifth: { Monday: '', Tuesday: '', Wednesday: '', Thursday: '', Friday: ''}
      });
      return {
        days,
        timeSlots,
        gridData
      };
    }
  }
  </script>

  <style>
  .v-container {
    max-width: 1200px;
    margin: auto;
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
  </style>
