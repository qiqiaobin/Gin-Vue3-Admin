<template>
	<el-main class="layout-main" :style="`minHeight: calc(100% - 51px`">
		<el-scrollbar
			ref="layoutMainScrollbarRef"
			class="layout-main-scroll layout-backtop-header-fixed"
			wrap-class="layout-main-scroll"
			view-class="layout-main-scroll"
		>
			<LayoutParentView />
			<LayoutFooter />
		</el-scrollbar>
		<el-backtop :target="setBacktopClass" />
	</el-main>
</template>

<script setup lang="ts" name="layoutMain">
import { defineAsyncComponent, onMounted, computed, ref } from 'vue';
import { NextLoading } from '/@/utils/loading';

// 引入组件
const LayoutParentView = defineAsyncComponent(() => import('/@/layout/routerView/parent.vue'));
const LayoutFooter = defineAsyncComponent(() => import('/@/layout/footer/index.vue'));

// 定义变量内容
const layoutMainScrollbarRef = ref();

// 设置 Backtop 回到顶部
const setBacktopClass = computed(() => {
	return `.layout-backtop-header-fixed .el-scrollbar__wrap`;
});
// 页面加载前
onMounted(() => {
	NextLoading.done(600);
});

// 暴露变量
defineExpose({
	layoutMainScrollbarRef,
});
</script>
