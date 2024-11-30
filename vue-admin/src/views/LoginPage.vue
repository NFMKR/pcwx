<template>
  <div class="login-container">
    <el-card class="login-card">
      <div class="title">系统登录</div>
      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="loginRules"
        class="login-form"
      >
        <el-form-item prop="Username">
          <el-input
            v-model="loginForm.Username"
            :prefix-icon="User"
            placeholder="用户名"
          />
        </el-form-item>

        <el-form-item prop="Password">
          <el-input
            v-model="loginForm.Password"
            :prefix-icon="Lock"
            type="password"
            placeholder="密码"
            @keyup.enter="handleLogin"
          />
        </el-form-item>

        <el-form-item>
          <el-button
            :loading="loading"
            type="primary"
            class="login-button"
            @click="handleLogin"
          >
            登录
          </el-button>
          <el-button class="register-link" type="text" @click="goToRegister">
            没有账号？去注册
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { User, Lock } from '@element-plus/icons-vue'

const store = useStore()
const router = useRouter()
const loginFormRef = ref(null)
const loading = ref(false)

const loginForm = reactive({
  Username: '',
  Password: ''
})

const loginRules = {
  Username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  Password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ]
}

const handleLogin = async () => {
  if (!loginFormRef.value) return

  try {
    // 表单验证
    await loginFormRef.value.validate()
    loading.value = true

    // 调用登录接口
    const { message } = await store.dispatch('login', loginForm)

    // 登录成功提示
    ElMessage({
      message: message || '登录成功',
      type: 'success',
      duration: 1500
    })

    // 立即跳转到主页
    await router.push('/dashboard')
  } catch (error) {
    console.error('Login error:', error)
    ElMessage.error(error.message || '登录失败')
  } finally {
    loading.value = false
  }
}

const goToRegister = () => {
  router.push('/register')
}
</script>

<style lang="scss" scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f3f3f3;

  .login-card {
    width: 400px;

    .title {
      text-align: center;
      font-size: 24px;
      font-weight: bold;
      margin-bottom: 30px;
    }

    .login-form {
      .login-button {
        width: 100%;
        margin-bottom: 10px;
      }

      .register-link {
        width: 100%;
      }
    }
  }
}
</style>
