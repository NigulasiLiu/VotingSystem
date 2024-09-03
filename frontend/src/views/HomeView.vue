<template>
  <div class="home-container">
    <div class="image-section">
      <img src="@/assets/homepage-image.png" alt="Homepage Image" class="home-image" />
    </div>
    <div class="content-section">
      <p class="subtitle">高效的隐私保护</p>
      <h1 class="main-title">在线投票系统</h1>
      <a-button type="primary" size="large" @click="$router.replace({ name: 'votinglist' })">
        查看选举列表
      </a-button>
    </div>
  </div>
</template>

<style scoped>
.home-container {
  margin-top: 70px;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 72vh;
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
  background-color: #ffffff; /* 背景色为白色 */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* 添加阴影 */
}

.subtitle {
  font-size: 1.5rem; /* 增大字体 */
  margin-bottom: 1rem; /* 调整与标题的间距 */
  color: #999; /* 灰色文字 */
  font-weight: 600; /* 加粗文字 */
}

.main-title {
  font-size: 4rem; /* 增大字体 */
  margin-bottom: 2rem; /* 调整与按钮的间距 */
  color: #333; /* 深色文字 */
  font-weight: 800; /* 更粗字体 */
  //text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2); /* 添加文字阴影 */
}

.content-section .ant-btn-primary {
  background-color: #1890ff;
  border-color: #1890ff;
  width: 200px;
}
</style>
