<template>
	<Classic />
</template>

<script setup lang="ts" name="layout">
import { onBeforeMount, onUnmounted } from 'vue';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import mittBus from '/@/utils/mitt';
import Classic from '/@/layout/main/classic.vue';

// 定义变量内容
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);

// 窗口大小改变时(适配移动端)
const onLayoutResize = () => {
	const clientWidth = document.body.clientWidth;
	if (clientWidth < 1000) {
		themeConfig.value.isCollapse = false;
		mittBus.emit('layoutMobileResize', {
			clientWidth,
		});
	} else {
		mittBus.emit('layoutMobileResize', {
			clientWidth,
		});
	}
};
// 页面加载前
onBeforeMount(() => {
	onLayoutResize();
	window.addEventListener('resize', onLayoutResize);
});
// 页面卸载时
onUnmounted(() => {
	window.removeEventListener('resize', onLayoutResize);
});
</script>
