<template>
	<div class="layout-navbars-container">
		<BreadcrumbIndex />
	</div>
</template>

<script setup lang="ts" name="layoutNavBars">
import { defineAsyncComponent } from 'vue';

// 引入组件
const BreadcrumbIndex = defineAsyncComponent(() => import('/@/layout/navBars/breadcrumb/index.vue'));

// 定义变量内容
</script>

<style scoped lang="scss">
.layout-navbars-container {
	display: flex;
	flex-direction: column;
	width: 100%;
	height: 100%;
}
</style>
