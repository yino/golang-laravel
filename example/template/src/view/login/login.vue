/* eslint-disable eqeqeq */
<style lang="less">
  @import './login.less';
</style>

<template>
  <div class="login">
    <div class="login-con">
      <Card icon="log-in" title="欢迎登录" :bordered="false">
        <div class="form-con">
          <login-form @on-success-valid="handleSubmit"></login-form>
          <p class="login-tip">输入任意用户名和密码即可</p>
        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import LoginForm from '_c/login-form'
import { setToken } from '@/libs/util'
import { login } from '@/api/user'
import store from '@/store'
export default {
  components: {
    LoginForm
  },
  methods: {
    handleSubmit ({ account, password }) {
      login({ account, password }).then(res => {
        if (res.code === 200) {
          setToken(res.data.token)
          store.commit("setToken", res.data.token)
          store.state.user.hasGetInfo = true
          store.commit("setAvatar", "")
          store.commit("setUserName", res.data.username)
          this.$router.push({name:"_home"})
        } else {
          this.$Message.error(res.msg)
        }
      })
    }
  }
}
</script>

<style>

</style>
