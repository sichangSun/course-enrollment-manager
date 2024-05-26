<template>
    <v-container >
      <div class="tittle">
        <h2>
          ログイン
        </h2>
      </div>
      <div v-if="errMsg" class="err-msg">{{ errMsg }}</div>
      <div class="login-form">
        <v-form @submit.prevent="$event.preventDefault()">
          <v-row>
            <v-col>
              <v-text-field
              v-model="state.email"
              :value="state.email"
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
                :value="state.password"
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
              <v-btn type="submit" :disabled="v$.$invalid" color="primary" @click="$emit('login-sumbit',state,v$)">Login</v-btn>
            </div>
        </v-form>
      </div>
    </v-container>
</template>

<script setup>
  import { reactive, computed, ref, watch} from 'vue'
  import { useVuelidate } from '@vuelidate/core'
  import { required, email,minLength } from '@vuelidate/validators'
  const props = defineProps(['errMsg'])
  const state=reactive({
    email:'',
    password:''
  })

  // validators setting
  const rules = {
    email: { required,email},
    password: { required, minLengthValue: minLength(6)},
  }
  const v$ = useVuelidate(rules, state)

//   watch(
//   () => state.email,
//   (newValue, oldValue) => {
//     console.log('new'+newValue)
//     console.log('old'+oldValue)
//   },
//   { deep: true }
// )

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
.err-msg{
  color: red;
  text-align: center;
}
</style>