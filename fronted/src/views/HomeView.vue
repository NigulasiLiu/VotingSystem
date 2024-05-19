<template>
  <div class="home-container">
    <div class="image-section">
      <img src="@/assets/homepage-image.png" alt="Homepage Image" class="home-image" />
    </div>
    <div class="content-section">
      <h1>欢迎使用电子投票系统</h1>
      <p>快速、安全、便捷的在线投票体验</p>
      <a-button type="primary" size="large" @click="$router.replace({ name: 'votinglist' })">
        查看投票列表
      </a-button>
    </div>
  </div>
</template>

<style scoped>
.home-container {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 70vh;
  text-align: center;
  background-color: #f0f2f5;
  padding: 0 2rem;
}

.image-section {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.home-image {
  max-width: 100%;
  height: auto;
  border-radius: 10px;
}

.content-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 2rem;
}

.content-section h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: #333;
}

.content-section p {
  font-size: 1.2rem;
  margin-bottom: 2rem;
  color: #666;
}

.content-section .ant-btn-primary {
  background-color: #1890ff;
  border-color: #1890ff;
  width: 200px;
}
</style>
