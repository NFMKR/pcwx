<template>
  <div class="user-management">
    <div class="header">
      <h2>用户管理</h2>
      <el-button type="primary" @click="showCreateDialog">创建用户</el-button>
    </div>

    <!-- 用户列表 -->
    <el-table :data="users" style="width: 100%">
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="role" label="角色">
        <template #default="{ row }">
          <el-tag :type="getRoleType(row.role)">{{ getRoleLabel(row.role) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="createdAt" label="创建时间">
        <template #default="{ row }">
          {{ formatDate(row.createdAt) }}
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态">
        <template #default="{ row }">
          <el-tag :type="row.status ? 'success' : 'danger'">
            {{ row.status ? '启用' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建用户对话框 -->
    <el-dialog v-model="dialogVisible" title="创建用户" width="500px">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" />
        </el-form-item>
        <el-form-item label="权限" prop="permissions">
          <el-checkbox-group v-model="form.permissions">
            <el-checkbox label="view">浏览权限</el-checkbox>
            <el-checkbox label="edit">编辑权限</el-checkbox>
            <el-checkbox label="manage">管理权限</el-checkbox>
            <el-checkbox label="people">人员管理</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="createUser">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'
import { getSubUsers, createSubUser } from '@/api/user'

export default {
  name: 'UserManagement',
  setup () {
    const store = useStore()
    const users = ref([])
    const dialogVisible = ref(false)
    const formRef = ref(null)

    const form = ref({
      username: '',
      password: '',
      permissions: []
    })

    const rules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
      ]
    }

    // 获取用户角色对应的API路径
    const getUserRolePath = () => {
      const role = store.getters.userRole
      switch (role) {
        case 'super_admin':
          return 'super'
        case 'second_admin':
          return 'second'
        case 'third_admin':
          return 'third'
        default:
          return ''
      }
    }

    // 加载用户列表
    const loadUsers = async () => {
      try {
        const rolePath = getUserRolePath()
        const response = await getSubUsers(rolePath)
        users.value = response
      } catch (error) {
        ElMessage.error('获取用户列表失败')
      }
    }

    // 显示创建用户对话框
    const showCreateDialog = () => {
      form.value = {
        username: '',
        password: '',
        permissions: []
      }
      dialogVisible.value = true
    }

    // 创建用户
    const createUser = async () => {
      if (!formRef.value) return

      try {
        await formRef.value.validate()
        const rolePath = getUserRolePath()
        await createSubUser(rolePath, form.value)
        ElMessage.success('创建用户成功')
        dialogVisible.value = false
        loadUsers()
      } catch (error) {
        ElMessage.error('创建用户失败')
      }
    }

    // 格式化日期
    const formatDate = (date) => {
      return new Date(date).toLocaleString()
    }

    // 获取角色标签类型
    const getRoleType = (role) => {
      const types = {
        super_admin: 'danger',
        second_admin: 'warning',
        third_admin: 'success',
        user: 'info'
      }
      return types[role] || 'info'
    }

    // 获取角色标签文本
    const getRoleLabel = (role) => {
      const labels = {
        super_admin: '超级管理员',
        second_admin: '二级管理员',
        third_admin: '三级管理员',
        user: '普通用户'
      }
      return labels[role] || '未知角色'
    }

    onMounted(() => {
      loadUsers()
    })

    return {
      users,
      dialogVisible,
      form,
      rules,
      formRef,
      showCreateDialog,
      createUser,
      formatDate,
      getRoleType,
      getRoleLabel
    }
  }
}
</script>

<style scoped>
.user-management {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>
