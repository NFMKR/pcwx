<template>
  <div class="profile-container">
    <el-card class="profile-card">
      <template #header>
        <div class="card-header">
          <span>个人信息</span>
          <el-button type="primary" @click="handleEdit">编辑</el-button>
        </div>
      </template>

      <div class="profile-info">
        <div class="info-item">
          <label>用户名：</label>
          <span>{{ userInfo.username }}</span>
        </div>

        <div class="info-item">
          <label>姓名：</label>
          <span>{{ userInfo.name }}</span>
        </div>

        <div class="info-item">
          <label>手机号：</label>
          <span>{{ userInfo.phone }}</span>
        </div>

        <div class="info-item">
          <label>邮箱：</label>
          <span>{{ userInfo.email }}</span>
        </div>

        <div class="info-item">
          <label>角色：</label>
          <span>{{ roleDisplay }}</span>
        </div>

        <div class="info-item">
          <label>创建时间：</label>
          <span>{{ formatDate(userInfo.createdAt) }}</span>
        </div>
      </div>
    </el-card>

    <!-- 修改信息对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="修改个人信息"
      width="500px"
    >
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item label="姓名" prop="name">
          <el-input v-model="form.name" />
        </el-form-item>

        <el-form-item label="手机号" prop="phone">
          <el-input v-model="form.phone" />
        </el-form-item>

        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" />
        </el-form-item>

        <el-form-item label="旧密码" prop="oldPassword">
          <el-input
            v-model="form.oldPassword"
            type="password"
            placeholder="不修改密码请留空"
          />
        </el-form-item>

        <el-form-item
          label="新密码"
          prop="newPassword"
          :rules="form.oldPassword ? [{ required: true, message: '请输入新密码', trigger: 'blur' }] : []"
        >
          <el-input
            v-model="form.newPassword"
            type="password"
            placeholder="不修改密码请留空"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/date'

const store = useStore()
const userInfo = computed(() => store.state.user || {})

const roleDisplay = computed(() => {
  const roleMap = {
    super_admin: '超级管理员',
    second_admin: '二级管理员',
    third_admin: '三级管理员',
    user: '普通用户'
  }
  return roleMap[userInfo.value.role] || '未知角色'
})

const dialogVisible = ref(false)
const formRef = ref(null)
const form = reactive({
  name: '',
  phone: '',
  email: '',
  oldPassword: '',
  newPassword: ''
})

const rules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ]
}

const handleEdit = () => {
  Object.assign(form, {
    name: userInfo.value.name,
    phone: userInfo.value.phone,
    email: userInfo.value.email,
    oldPassword: '',
    newPassword: ''
  })
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()

    const updateData = {
      name: form.name,
      phone: form.phone,
      email: form.email
    }

    if (form.oldPassword && form.newPassword) {
      updateData.oldPassword = form.oldPassword
      updateData.newPassword = form.newPassword
    }

    await store.dispatch('updateProfile', updateData)
    ElMessage.success('更新成功')
    dialogVisible.value = false
  } catch (error) {
    ElMessage.error(error.response?.data?.message || '更新失败')
  }
}
</script>

<style lang="scss" scoped>
.profile-container {
  padding: 20px;

  .profile-card {
    max-width: 800px;
    margin: 0 auto;

    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .profile-info {
      .info-item {
        margin-bottom: 20px;

        label {
          display: inline-block;
          width: 100px;
          color: #606266;
        }

        span {
          color: #303133;
        }
      }
    }
  }
}
</style>
