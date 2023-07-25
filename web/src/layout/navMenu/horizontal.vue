<template>
	<div class="el-menu-horizontal-warp">
		<el-menu router :default-active="state.defaultActive" background-color="transparent" mode="horizontal">
			<template v-for="val in menuLists">
				<el-sub-menu :index="val.path" v-if="val.children && val.children.length > 0" :key="val.path">
					<template #title>
						<!--SvgIcon :name="val.meta.icon" /-->
						<span>{{ val.meta.title }}</span>
					</template>
					<SubItem :chil="val.children" />
				</el-sub-menu>
				<template v-else>
					<el-menu-item :index="val.path" :key="val.path">
						<template #title v-if="!val.meta.isLink || (val.meta.isLink && val.meta.isIframe)">
							<!--SvgIcon :name="val.meta.icon" /-->
							{{ val.meta.title }}
						</template>
						<template #title v-else>
							<a class="w100" @click.prevent="onALinkClick(val)">
								<!--SvgIcon :name="val.meta.icon" /-->
								{{ val.meta.title }}
							</a>
						</template>
					</el-menu-item>
				</template>
			</template>
		</el-menu>
	</div>
</template>

<script setup lang="ts" name="navMenuHorizontal">
import { defineAsyncComponent, reactive, computed, onBeforeMount } from 'vue';
import { useRoute, onBeforeRouteUpdate, RouteRecordRaw } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useRoutesList } from '/@/stores/routesList';
import other from '/@/utils/other';
import mittBus from '/@/utils/mitt';

// 引入组件
const SubItem = defineAsyncComponent(() => import('/@/layout/navMenu/subItem.vue'));

// 定义父组件传过来的值
const props = defineProps({
	// 菜单列表
	menuList: {
		type: Array<RouteRecordRaw>,
		default: () => [],
	},
});

// 定义变量内容
const stores = useRoutesList();
const { routesList } = storeToRefs(stores);
const route = useRoute();
const state = reactive({
	defaultActive: '' as string | undefined,
});

// 获取父级菜单数据
const menuLists = computed(() => {
	return <RouteItems>props.menuList;
});
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
// 传送当前子级数据到菜单中
const setSendClassicChildren = (path: string) => {
	const currentPathSplit = path.split('/');
	let currentData: MittMenu = { children: [] };
	filterRoutesFun(routesList.value).map((v, k) => {
		if (v.path === `/${currentPathSplit[1]}`) {
			v['k'] = k;
			currentData['item'] = { ...v };
			currentData['children'] = [{ ...v }];
			if (v.children) currentData['children'] = v.children;
		}
	});
	return currentData;
};
// 设置页面当前路由高亮
const setCurrentRouterHighlight = (currentRoute: RouteToFrom) => {
	const { path } = currentRoute;
	state.defaultActive = `/${path?.split('/')[1]}`;
};
// 打开外部链接
const onALinkClick = (val: RouteItem) => {
	other.handleOpenLink(val);
};
// 页面加载前
onBeforeMount(() => {
	setCurrentRouterHighlight(route);
});
// 路由更新时
onBeforeRouteUpdate((to) => {
	// 修复：https://gitee.com/lyt-top/vue-next-admin/issues/I3YX6G
	setCurrentRouterHighlight(to);
	// 修复经典布局开启切割菜单时，点击tagsView后左侧导航菜单数据不变的问题
	mittBus.emit('setSendClassicChildren', setSendClassicChildren(to.path));
});
</script>

<style scoped lang="scss">
.el-menu-horizontal-warp {
	flex: 1;
	overflow: hidden;
	margin-right: 30px;

	:deep(a) {
		width: 100%;
	}
}

// 横向菜单
.el-menu--horizontal {
	background: #182132;
    .el-sub-menu {
		height: 48px !important;
		line-height: 48px !important;
        font-size: 14px;
		color: #182132;
		.el-sub-menu__title {
			height: 48px !important;
			line-height: 48px !important;
			color: #182132;
            font-size: 14px;
		}
		.el-popper.is-pure.is-light {
			.el-menu--horizontal {
				.el-sub-menu .el-sub-menu__title {
					background-color: #96a2b9;
					color: #96a2b9;
				}
			}
		}
    .el-menu-item {
	height: 56px !important;
	line-height: 56px !important;
    }
    .el-sub-menu__title {
	color: #96a2b9;
    }
    .el-menu-item:hover {
	background-color: #F6F6F9;
    }
	}
}

// 横向菜单（经典、横向）布局
.el-menu--horizontal {
	border-bottom: none !important;
	width: 100% !important;
	.el-menu-item,
	.el-sub-menu__title {
        border-bottom: none !important;
		height: 48px !important;
		color: #979ba5;
        font-size: 14px;
	}
	.el-menu-item:not(.is-active):hover,
	.el-sub-menu:not(.is-active):hover .el-sub-menu__title {
		color: #e1ecff;
        background: rgba(49, 64, 94, .5)!important;
	}
    .el-menu-item.is-active,
    .el-sub-menu.is-active .el-sub-menu__title,
    .el-sub-menu:not(.is-opened):hover .el-sub-menu__title {
    background: rgba(49, 64, 94, .5)!important;
    color: #ffffff !important;
    border-bottom: none !important;
   }
}
</style>
