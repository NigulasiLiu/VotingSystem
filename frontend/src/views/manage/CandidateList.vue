<template>
    <div class="container">
        <a-space>
            <a-button type="link"
                      size="large"
                      @click="$router.replace({name:'votinglist'})">
                <ArrowLeftOutlined />返回投票列表页
            </a-button>
            <a-button type="primary"
                      size="large"
                      @click="$router.replace({name:'addcandidate'})">
              <PlusOutlined /> 新增候选人
            </a-button>
        </a-space>
        <a-divider></a-divider>
        <a-list v-if="candidateData.list" :grid="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 3, xxl: 4 }" item-layout="vertical"
            size="large" :data-source="candidateData.list">
            <template #renderItem="{ item }">
                <a-list-item :key="item.ID">
                    <a-card :title="`ID&emsp;${item.ID}`">
                        <template #actions>
                          <router-link :to="{ name: 'editcandidate', params: { id: item.ID }}">
                            <EditOutlined />&emsp;<span class="edit-text">修改信息</span>
                          </router-link>
                        </template>
                        <a-list-item-meta>
                            <template #title>
                                <a>{{ item.Name }}</a>
                            </template>
                            <template #avatar>
                                <a-avatar v-if="candidateData" :src="`data:image/png;base64,${item.Photo}`" />
                            </template>
                        </a-list-item-meta>
                        {{ item.Detail }}
                    </a-card>
                </a-list-item>
            </template>
        </a-list>
    </div>
</template>

<script setup>
import { candidateData } from '@/data/candidatedata';
import { EditOutlined, ArrowLeftOutlined, PlusOutlined } from '@ant-design/icons-vue';

</script>

<style scoped>
.container {
  margin: 0px auto; /* 调整上边距 */
    max-width: 1350px;
    padding: 0 20px;
}
.edit-text {
  //font-weight: bold;
  color: #000000;
  text-decoration: none;
  cursor: pointer; /* 使鼠标悬停时显示为指针形状 */
  //transition: color 0.3s ease; /* 渐变效果 */
}

.edit-text:hover {
  color: #0056b3; /* 鼠标悬停时的颜色变化 */
}
</style>
