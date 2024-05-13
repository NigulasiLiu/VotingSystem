<template>
    <div class="container">
        <a-space>
            <a-button type="link" @click="$router.replace({name:'votinglist'})">
                <ArrowLeftOutlined />返回投票列表页
            </a-button>
            <a-button type="primary" @click="$router.replace({name:'addcandidate'})">
                新增候选人
            </a-button>
        </a-space>
        <a-divider></a-divider>
        <a-list :grid="{ gutter: 2, xs: 1, sm: 2, md: 2, lg: 3, xl: 4, xxl: 6 }" item-layout="vertical" size="large"
            :data-source="candidateData.list">
            <template #renderItem="{ item }">
                <a-list-item :key="item.ID">
                    <a-card :title="`ID&emsp;${item.ID}`">
                        <template #actions>
                            <router-link :to="{ name: 'editcandidate', params: { id: item.ID }}">
                                <EditOutlined />&emsp;修改信息
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
import { EditOutlined, ArrowLeftOutlined } from '@ant-design/icons-vue';

</script>

<style scoped>
.container {
    margin: 0 auto;
    max-width: 800px;
    padding: 0 20px;
}
</style>
