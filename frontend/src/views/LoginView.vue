<template>
  <v-app>
  <Login class="login"
  @login-sumbit="loginSumbit"
  :errMsg="errMsg"
  ></Login>

  </v-app>

</template>

<script setup>
import { reactive, computed, ref, watch} from 'vue'
  import Login from  '../components/Login.vue'
  import { useRouter } from 'vue-router'
  import axios from '../axios-config';

  const router = useRouter()

  let errMsg=ref('')
  const loginSumbit = async (state,$v) => {

    if ($v.$invalid) {
      errMsg.value='ログインが問題ありました。確認した上もう一回お試しください。'
      return;
    }

    const res={
      email: state.email,
      password: state.password,
    }
    // call API
    try {
      await axios.post(`${_BASE_URL_}api/login`, res)
      .then(response => {
        console.log('Login successful');
        // const token = localStorage.getItem('jwtToken');
        // //Save jwt token
        // localStorage.setItem('jwtToken', response.data.token)
        router.push({
          // Successful to StudentHome
          name: 'StudentHome'
          });
      })
    } catch (error) {
      console.error('Login failed:', error.response.data);
      if(error.response.data.error=='user not found' || error.response.data.error =='invalid password'){
        errMsg.value='メールアドレス、またはパスワードが間違っています。'
      }else{
        router.push({
        name: 'ErrorPage'
        })
      }
    }
  }

</script>
<style>
.login{
  margin: 90px auto;
}

</style>