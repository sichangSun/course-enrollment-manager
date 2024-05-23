<template>
  <v-container>
    <div class="base-header-info">
      <h4>name</h4>
    </div>
    <v-row class="base-info-box">
      <v-col cols="3" >
        <div>担当教員：</div>
        <v-divider></v-divider>
        <div>{{  }} {{  }}</div>
      </v-col>
      <v-col cols="3" >
        <div>実施時間：</div>
        <v-divider></v-divider>
        <div>
          {{  }}
        </div>
      </v-col>
      <v-col cols="3" >
        <div>単位数：</div>
        <v-divider></v-divider>
        <div>{{  }}</div>
      </v-col>
      <v-col cols="3" >
        <div>分類：</div>
        <v-divider></v-divider>
        <div>{{  }}</div>
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


          </v-card-text>
        </v-card>
        

      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>


</script>
<style>
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
</style>