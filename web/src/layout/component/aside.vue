<template>
	<div class="h100">
		<el-aside class="layout-aside" :class="setCollapseStyle">
			<el-scrollbar class="flex-auto" ref="layoutAsideScrollbarRef">
				<Vertical :menuList="state.menuList" />
			</el-scrollbar>
            <div class="nav-option">
                <SvgIcon class="nav-stick" :name="themeConfig.isCollapse ? 'ele-Expand' :'ele-Fold'" :size="16" @click="onThemeConfigChange"/>
            </div>
		</el-aside>
	</div>
</template>

<script setup lang="ts" name="layoutAside">
import { defineAsyncComponent, reactive, computed, watch, onBeforeMount, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useRoutesList } from '/@/stores/routesList';
import { useThemeConfig } from '/@/stores/themeConfig';
import mittBus from '/@/utils/mitt';

// 引入组件
const Vertical = defineAsyncComponent(() => import('/@/layout/navMenu/vertical.vue'));

// 定义变量内容
const layoutAsideScrollbarRef = ref();
const stores = useRoutesList();
const storesThemeConfig = useThemeConfig();
const { routesList } = storeToRefs(stores);
const { themeConfig } = storeToRefs(storesThemeConfig);
const state = reactive<AsideState>({
	menuList: [],
	clientWidth: 0,
});

// 展开/收起左侧菜单点击
const onThemeConfigChange = () => {
	themeConfig.value.isCollapse = !themeConfig.value.isCollapse;
};

// 设置菜单展开/收起时的宽度
const setCollapseStyle = computed(() => {
	const { isCollapse } = themeConfig.value;
	const asideBrColor = 'layout-el-aside-br-color';
	// 判断是否是手机端
	if (state.clientWidth <= 1000) {
		if (isCollapse) {
			document.body.setAttribute('class', 'el-popup-parent--hidden');
			const asideEle = document.querySelector('.layout-container') as HTMLElement;
			const modeDivs = document.createElement('div');
			modeDivs.setAttribute('class', 'layout-aside-mobile-mode');
			asideEle.appendChild(modeDivs);
			modeDivs.addEventListener('click', closeLayoutAsideMobileMode);
			return [asideBrColor, 'layout-aside-mobile', 'layout-aside-mobile-open'];
		} else {
			// 关闭弹窗
			closeLayoutAsideMobileMode();
			return [asideBrColor, 'layout-aside-mobile', 'layout-aside-mobile-close'];
		}
	} else {
		//经典布局分割菜单只有一项子级时，菜单收起时宽度给 1px，防止切换动画消失
        if (state.menuList.length <= 1)
        {
            return [asideBrColor, 'layout-aside-pc-1'];
		} else {
			// 其它布局给 64px，显示图标
			if (isCollapse) {
                return [asideBrColor, 'layout-aside-pc-64'];
            }else {
                return [asideBrColor, 'layout-aside-pc-220'];
            }
		}
	}
});
// 关闭移动端蒙版
const closeLayoutAsideMobileMode = () => {
	const el = document.querySelector('.layout-aside-mobile-mode');
	el?.setAttribute('style', 'animation: error-img-two 0.3s');
	setTimeout(() => {
		el?.parentNode?.removeChild(el);
	}, 300);
	const clientWidth = document.body.clientWidth;
	if (clientWidth < 1000) themeConfig.value.isCollapse = false;
	document.body.setAttribute('class', '');
};
// 设置/过滤路由（非静态路由/是否显示在菜单中）
const setFilterRoutes = () => {
	state.menuList = filterRoutesFun(routesList.value);
};
// 路由过滤递归函数
const filterRoutesFun = <T extends RouteItem>(arr: T[]): T[] => {
	return arr
		.filter((item: T) => !item.meta?.isHide)
		.map((item: T) => {
			item = Object.assign({}, item);
			if (item.children) item.children = filterRoutesFun(item.children);
			return item;
		});
};
// 设置菜单导航是否固定（移动端）
const initMenuFixed = (clientWidth: number) => {
	state.clientWidth = clientWidth;
};
// 页面加载前
onBeforeMount(() => {
	initMenuFixed(document.body.clientWidth);
	setFilterRoutes();
    // 开启经典布局分割菜单时，设置菜单数据
	mittBus.on('setSendClassicChildren', (res: MittMenu) => {
		// 经典布局分割菜单只有一项子级时，收起左侧导航菜单
		res.children.length <= 1 ? (themeConfig.value.isCollapse = true) : (themeConfig.value.isCollapse = false);
		state.menuList = [];
		state.menuList = res.children;
	});
    // 开启经典布局分割菜单时，重新处理菜单数据
	mittBus.on('getBreadcrumbIndexSetFilterRoutes', () => {
		setFilterRoutes();
	});
    // 监听窗口大小改变时(适配移动端)
	mittBus.on('layoutMobileResize', (res: LayoutMobileResize) => {
		initMenuFixed(res.clientWidth);
		closeLayoutAsideMobileMode();
	});
});
// 监听 pinia 值的变化，动态赋值给菜单中
watch(
	() => routesList.value,
	() => {
		setFilterRoutes();
	}
);
</script>

<style lang="scss" scoped>
.icon {
  vertical-align: middle;
}
.nav-option {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 50px;
  line-height: 49px;
  border-top: 1px solid #dcdee5;
  font-size: 0;
  color: #63656e;
  &:before {
    content: "";
    display: inline-block;
    height: 100%;
    width: 0;
    vertical-align: middle;
  }
  .nav-stick {
    display: inline-block;
    vertical-align: middle;
    width: 32px;
    height: 32px;
    margin: 0 0 0 13px;
    line-height: 32px;
    text-align: center;
    font-size: 14px;
    cursor: pointer;
    //transition: transform $duration $cubicBezier;
    &:hover {
      opacity: 0.8;
    }
    &.sticked {
      transform: rotate(180deg);
    }
  }
}
</style>