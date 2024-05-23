<template>
  <!-- <v-app> -->
    <v-container >
      <div class="tittle">
        <h2>
          ログイン
        </h2>
      </div>
      <div class="login-form">
    <!-- <h2>Login</h2> -->
    <v-form>
      <v-row>
        <v-col>
          <v-text-field
          v-model="state.email"
          label="E-mail"
          @blur="v$.email.$touch"
          :error-messages ="v$.email.$errors.map((e) => e.$message)"
          @input="v$.email.$touch">
          </v-text-field>

        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            v-model="state.password"
            label="Password"
            type="password"
            @blur="v$.password.$touch"
             :error-messages ="v$.password.$errors.map((e) => e.$message)"
             @input="v$.password.$touch"
            >
          </v-text-field>

        </v-col>
      </v-row>
        <div class="button" >
          <v-btn type="submit" :disabled="v$.$invalid" color="primary" @click.prevent="login">Login</v-btn>
        </div>
    </v-form>
  </div>
    </v-container>
  <!-- </v-app> -->
</template>

<script setup>
import { reactive, computed, ref, watch} from 'vue'
import { useVuelidate } from '@vuelidate/core'
import { required, email,minLength } from '@vuelidate/validators'
  let state = reactive({
        email: '',
        password:''
  })
  // validators setting
  const rules = {
        email: { required,email },
        password: { required, minLengthValue: minLength(6)},
  }
  const v$ = useVuelidate(rules, state)
  // validate data
  function validateForm(){

  }
  //login
  function login () {
    validateForm()

  }
</script>
<style>
.v-container {
    max-width: 1000px;
  }
.login-form {
  padding: 30px 20px;
  margin: auto;
  width: 500px;
}
.button{
  margin-top: 40px;
}
.tittle{
  margin: auto;
    padding: 10px 30%;
    text-align: center;

}
</style>