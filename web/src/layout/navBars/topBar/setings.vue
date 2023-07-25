<template>
	<div class="layout-breadcrumb-seting">
		<el-drawer title="布局配置" direction="rtl" destroy-on-close size="260px">
			<el-scrollbar class="layout-breadcrumb-seting-bar">
				<!-- 界面设置 -->
				<el-divider content-position="left">界面设置</el-divider>
				<div class="layout-breadcrumb-seting-bar-flex" :style="{ opacity: getThemeConfig.layout === 'transverse' ? 0.5 : 1 }">
					<div class="layout-breadcrumb-seting-bar-flex-label">菜单水平折叠</div>
					<div class="layout-breadcrumb-seting-bar-flex-value">
						<el-switch
							v-model="getThemeConfig.isCollapse"
							:disabled="getThemeConfig.layout === 'transverse'"
							size="small"
							@change="onThemeConfigChange"
						></el-switch>
					</div>
				</div>

				<!-- 其它设置 -->
				<el-divider content-position="left">其它设置</el-divider>
				<div class="layout-breadcrumb-seting-bar-flex mt15">
					<div class="layout-breadcrumb-seting-bar-flex-label">主页面切换动画</div>
					<div class="layout-breadcrumb-seting-bar-flex-value">
						<el-select v-model="getThemeConfig.animation" placeholder="请选择" size="default" style="width: 90px" @change="setLocalThemeConfig">
							<el-option label="slide-right" value="slide-right"></el-option>
							<el-option label="slide-left" value="slide-left"></el-option>
							<el-option label="opacitys" value="opacitys"></el-option>
						</el-select>
					</div>
				</div>
			</el-scrollbar>
		</el-drawer>
	</div>
</template>

<script setup lang="ts" name="layoutBreadcrumbSeting">
import { nextTick, onUnmounted, onMounted, computed, reactive } from 'vue';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import { Local } from '/@/utils/storage';
import commonFunction from '/@/utils/commonFunction';
import other from '/@/utils/other';
import mittBus from '/@/utils/mitt';

// 定义变量内容
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const { copyText } = commonFunction();
const state = reactive({
	isMobile: false,
});

// 获取布局配置信息
const getThemeConfig = computed(() => {
	return themeConfig.value;
});

// 3、界面设置 --> 菜单水平折叠
const onThemeConfigChange = () => {
	setDispatchThemeConfig();
};

// 触发 store 布局配置更新
const setDispatchThemeConfig = () => {
	setLocalThemeConfig();
	setLocalThemeConfigStyle();
};
// 存储布局配置
const setLocalThemeConfig = () => {
	Local.remove('themeConfig');
	Local.set('themeConfig', getThemeConfig.value);
};
// 存储布局配置全局主题样式（html根标签）
const setLocalThemeConfigStyle = () => {
	Local.set('themeConfigStyle', document.documentElement.style.cssText);
};

onMounted(() => {
	nextTick(() => {
		// 判断当前布局是否不相同，不相同则初始化当前布局的样式，防止监听窗口大小改变时，布局配置logo、菜单背景等部分布局失效问题
		//if (!Local.get('frequency')) initLayoutChangeFun();
		//Local.set('frequency', 1);
		// 监听窗口大小改变，非默认布局，设置成默认布局（适配移动端）
		mittBus.on('layoutMobileResize', (res: LayoutMobileResize) => {
			state.isMobile = other.isMobile();
		});
	});
});
onUnmounted(() => {
	mittBus.off('layoutMobileResize', () => {});
});

</script>

<style scoped lang="scss">
.layout-breadcrumb-seting-bar {
	height: calc(100vh - 50px);
	padding: 0 15px;
	:deep(.el-scrollbar__view) {
		overflow-x: hidden !important;
	}
	.layout-breadcrumb-seting-bar-flex {
		display: flex;
		align-items: center;
		margin-bottom: 5px;
		&-label {
			flex: 1;
			color: var(--el-text-color-primary);
		}
	}
	.layout-drawer-content-flex {
		overflow: hidden;
		display: flex;
		flex-wrap: wrap;
		align-content: flex-start;
		margin: 0 -5px;
		.layout-drawer-content-item {
			width: 50%;
			height: 70px;
			cursor: pointer;
			border: 1px solid transparent;
			position: relative;
			padding: 5px;
			.el-container {
				height: 100%;
				.el-aside-dark {
					background-color: var(--next-color-seting-header);
				}
				.el-aside {
					background-color: var(--next-color-seting-aside);
				}
				.el-header {
					background-color: var(--next-color-seting-header);
				}
				.el-main {
					background-color: var(--next-color-seting-main);
				}
			}
			.el-circular {
				border-radius: 2px;
				overflow: hidden;
				border: 1px solid transparent;
				transition: all 0.3s ease-in-out;
			}
			.drawer-layout-active {
				border: 1px solid;
				border-color: var(--el-color-primary);
			}
			.layout-tips-warp,
			.layout-tips-warp-active {
				transition: all 0.3s ease-in-out;
				position: absolute;
				left: 50%;
				top: 50%;
				transform: translate(-50%, -50%);
				border: 1px solid;
				border-color: var(--el-color-primary-light-5);
				border-radius: 100%;
				padding: 4px;
				.layout-tips-box {
					transition: inherit;
					width: 30px;
					height: 30px;
					z-index: 9;
					border: 1px solid;
					border-color: var(--el-color-primary-light-5);
					border-radius: 100%;
					.layout-tips-txt {
						transition: inherit;
						position: relative;
						top: 5px;
						font-size: 12px;
						line-height: 1;
						letter-spacing: 2px;
						white-space: nowrap;
						color: var(--el-color-primary-light-5);
						text-align: center;
						transform: rotate(30deg);
						left: -1px;
						background-color: var(--next-color-seting-main);
						width: 32px;
						height: 17px;
						line-height: 17px;
					}
				}
			}
			.layout-tips-warp-active {
				border: 1px solid;
				border-color: var(--el-color-primary);
				.layout-tips-box {
					border: 1px solid;
					border-color: var(--el-color-primary);
					.layout-tips-txt {
						color: var(--el-color-primary) !important;
						background-color: var(--next-color-seting-main) !important;
					}
				}
			}
			&:hover {
				.el-circular {
					transition: all 0.3s ease-in-out;
					border: 1px solid;
					border-color: var(--el-color-primary);
				}
				.layout-tips-warp {
					transition: all 0.3s ease-in-out;
					border-color: var(--el-color-primary);
					.layout-tips-box {
						transition: inherit;
						border-color: var(--el-color-primary);
						.layout-tips-txt {
							transition: inherit;
							color: var(--el-color-primary) !important;
							background-color: var(--next-color-seting-main) !important;
						}
					}
				}
			}
		}
	}
	.copy-config {
		margin: 10px 0;
		.copy-config-btn {
			width: 100%;
			margin-top: 15px;
		}
		.copy-config-btn-reset {
			width: 100%;
			margin: 10px 0 0;
		}
	}
}
</style>
