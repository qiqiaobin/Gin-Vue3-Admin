<template>
	<el-header class="layout-header">
		<NavBarsIndex />
	</el-header>
</template>

<script setup lang="ts" name="layoutHeader">
import { defineAsyncComponent } from 'vue';

// 引入组件
const NavBarsIndex = defineAsyncComponent(() => import('/@/layout/navBars/index.vue'));

// 定义变量内容
</script>
