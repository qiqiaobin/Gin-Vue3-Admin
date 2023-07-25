<template>
	<div class="layout-parent">
		<router-view v-slot="{ Component }">
			<transition :name="setTransitionName" mode="out-in">
				<keep-alive>
					<component :is="Component" :key="state.refreshRouterViewKey" class="w100" v-show="!isIframePage" />
				</keep-alive>
			</transition>
		</router-view>
		<transition :name="setTransitionName" mode="out-in">
			<Iframes class="w100" v-show="isIframePage" :refreshKey="state.iframeRefreshKey" :name="setTransitionName" :list="state.iframeList" />
		</transition>
	</div>
</template>

<script setup lang="ts" name="layoutParentView">
import { defineAsyncComponent, computed, reactive, watch, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';

// 引入组件
const Iframes = defineAsyncComponent(() => import('/@/layout/routerView/iframes.vue'));

// 定义变量内容
const route = useRoute();
const router = useRouter();
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const state = reactive<ParentViewState>({
	refreshRouterViewKey: '', // 非 iframe tagsview 右键菜单刷新时
	iframeRefreshKey: '', // iframe tagsview 右键菜单刷新时
	iframeList: [],
});

// 设置主界面切换动画
const setTransitionName = computed(() => {
	return themeConfig.value.animation;
});
// 设置 iframe 显示/隐藏
const isIframePage = computed(() => {
	return route.meta.isIframe;
});
// 获取 iframe 组件列表(未进行渲染)
const getIframeListRoutes = async () => {
	router.getRoutes().forEach((v) => {
		if (v.meta.isIframe) {
			v.meta.isIframeOpen = false;
			v.meta.loading = true;
			state.iframeList.push({ ...v });
		}
	});
};
// 页面加载时
onMounted(() => {
	getIframeListRoutes();
	// https://gitee.com/lyt-top/vue-next-admin/issues/I58U75
	// https://gitee.com/lyt-top/vue-next-admin/issues/I59RXK
	// https://gitee.com/lyt-top/vue-next-admin/pulls/40
});
// 监听路由变化，防止 tagsView 多标签时，切换动画消失
// https://toscode.gitee.com/lyt-top/vue-next-admin/pulls/38/files
watch(
	() => route.fullPath,
	() => {
		state.refreshRouterViewKey = decodeURI(route.fullPath);
	},
	{
		immediate: true,
	}
);
</script>
