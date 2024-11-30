<template>
  <div class="error-page">
    <h1>403</h1>
    <h2>抱歉，您没有权限访问此页面</h2>
    <el-button type="primary" @click="$router.push('/')">返回首页</el-button>
  </div>
</template>

<script>
export default {
  name: 'ForbiddenPage'
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
    color: #409EFF;
  }

  h2 {
    margin: 20px 0;
    color: #666;
  }
}
</style>
