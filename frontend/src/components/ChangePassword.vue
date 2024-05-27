<template>
  <!-- <v-app> -->
    <v-container >
      <div class="tittle">
        <h2>
          パスワード変更
        </h2>
      </div>
      <div class="login-form">
    <!-- <h2>Login</h2> -->
      <v-form>
        <v-row>
          <v-col>
            <v-text-field
              v-model="state.oldPassword"
              :value="state.oldPassword"
              label="Please enter your old password."
              type="password"
              @blur="v$.oldPassword.$touch"
              :error-messages ="v$.oldPassword.$errors.map((e) => e.$message)"
              @input="v$.oldPassword.$touch"
              >
            </v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field
              v-model="state.newPassword"
              :value="state.newPassword"
              label="Please enter a new password."
              type="password"
              @blur="v$.newPassword.$touch"
              :error-messages ="v$.newPassword.$errors.map((e) => e.$message)"
              @input="v$.newPassword.$touch"
              >
            </v-text-field>

          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field
              v-model="state.confirmPassword"
              :value="state.confirmPassword"
              label="Please enter your new password again."
              type="password"
              @blur="v$.confirmPassword.$touch"
              :error-messages ="v$.confirmPassword.$errors.map((e) => e.$message)"
              @input="v$.confirmPassword.$touch"
              >
            </v-text-field>

          </v-col>
        </v-row>
    </v-form>
      <div class="button" >
          <v-btn type="submit" :disabled="v$.$invalid" color="primary"  @click="$emit('change-password',state,v$)">提出</v-btn>
      </div>
    </div>
  </v-container>
  <!-- </v-app> -->
</template>

<script setup>
  import { reactive, computed, ref} from 'vue'
  import { useVuelidate } from '@vuelidate/core'
  import { required, email,minLength,sameAs } from '@vuelidate/validators'



  const state=reactive({
    oldPassword:'',
    newPassword:'',
    confirmPassword:''
  })

  // validators setting
  const rules = {
    oldPassword: { required },
    newPassword: { required, minLengthValue: minLength(6)},
    confirmPassword: {
      required,
      sameAsPassword:sameAs(state.newPassword)},
  }
  const v$ = useVuelidate(rules, state)

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
