<template>
  <div class="error-page">
    <h1>404</h1>
    <h2>抱歉，页面不存在</h2>
    <el-button type="primary" @click="$router.push('/')">返回首页</el-button>
  </div>
</template>

<script>
export default {
  name: 'NotFoundPage'
}
</script>

<style lang="scss" scoped>
.error-page {
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: #f0f2f5;

  h1 {
    font-size: 120px;
    margin: 0;
    color: #f56c6c;
  }

  h2 {
    margin: 20px 0;
    color: #666;
  }
}
</style>
