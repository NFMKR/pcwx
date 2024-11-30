<template>
  <el-container class="app-wrapper">
    <!-- 侧边栏 -->
    <el-aside width="200px" class="sidebar-container">
      <div class="logo">
        <img src="../assets/logo.png" alt="Logo">
        <span>管理系统</span>
      </div>

      <el-menu
        :default-active="route.path"
        :router="true"
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
      >
        <el-menu-item index="/dashboard">
          <el-icon><Monitor /></el-icon>
          <template #title>首页</template>
        </el-menu-item>

        <!-- 超级管理员菜单 -->
        <el-menu-item
          v-if="userRole === 'super_admin'"
          index="/second-admin"
        >
          <el-icon><User /></el-icon>
          <template #title>二级管理员</template>
        </el-menu-item>

        <!-- 二级管理员菜单 -->
        <el-menu-item
          v-if="userRole === 'second_admin'"
          index="/third-admin"
        >
          <el-icon><UserFilled /></el-icon>
          <template #title>三级管理员</template>
        </el-menu-item>

        <!-- 三级管理员菜单 -->
        <el-menu-item
          v-if="userRole === 'third_admin'"
          index="/users"
        >
          <el-icon><Avatar /></el-icon>
          <template #title>用户管理</template>
        </el-menu-item>

        <el-menu-item index="/profile">
          <el-icon><Setting /></el-icon>
          <template #title>个人信息</template>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <el-container>
      <!-- 顶部导航 -->
      <el-header height="60px" class="navbar">
        <div class="right-menu">
          <el-dropdown @command="handleCommand">
            <span class="el-dropdown-link">
              {{ username }}
              <el-icon class="el-icon--right"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 主要内容区 -->
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'
import {
  Monitor,
  User,
  UserFilled,
  Avatar,
  Setting,
  ArrowDown
} from '@element-plus/icons-vue'

const store = useStore()
const router = useRouter()
const route = useRoute()

const userRole = computed(() => store.state.user?.role)
const username = computed(() => store.state.user?.username)

const handleCommand = (command) => {
  if (command === 'profile') {
    router.push('/profile')
  } else if (command === 'logout') {
    store.dispatch('logout')
    router.push('/login')
  }
}
</script>

<style lang="scss" scoped>
.app-wrapper {
  height: 100vh;

  .sidebar-container {
    background-color: #304156;

    .logo {
      height: 60px;
      display: flex;
      align-items: center;
      padding: 0 15px;
      color: #fff;

      img {
        width: 32px;
        height: 32px;
        margin-right: 10px;
      }

      span {
        font-size: 20px;
        font-weight: bold;
      }
    }
  }

  .navbar {
    background-color: #fff;
    box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
    display: flex;
    align-items: center;
    justify-content: flex-end;
    padding: 0 20px;

    .right-menu {
      .el-dropdown-link {
        cursor: pointer;
        color: #333;
        display: flex;
        align-items: center;

        &:hover {
          color: #409EFF;
        }
      }
    }
  }

  .el-main {
    background-color: #f0f2f5;
    padding: 20px;
  }
}

:deep(.el-menu) {
  border-right: none;

  .el-menu-item {
    .el-icon {
      margin-right: 10px;
    }
  }
}
</style>
