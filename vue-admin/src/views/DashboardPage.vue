<template>
  <div class="dashboard-container">
    <el-card class="welcome-card">
      <template #header>
        <div class="card-header">
          <span>欢迎使用管理系统</span>
        </div>
      </template>
      <div class="content">
        <h3>当前用户：{{ username }}</h3>
        <p>角色：{{ roleName }}</p>
      </div>
    </el-card>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'DashboardPage',

  setup () {
    const store = useStore()

    const username = computed(() => store.state.user?.username || '')
    const roleName = computed(() => {
      const roleMap = {
        super_admin: '超级管理员',
        second_admin: '二级管理员',
        third_admin: '三级管理员',
        user: '普通用户'
      }
      return roleMap[store.state.user?.role] || '未知角色'
    })

    return {
      username,
      roleName
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-container {
  padding: 20px;

  .welcome-card {
    max-width: 600px;
    margin: 0 auto;

    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .content {
      text-align: center;
      padding: 20px;

      h3 {
        margin-bottom: 10px;
      }
    }
  }
}
</style>
