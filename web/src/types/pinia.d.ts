/**
 * pinia 类型定义
 */

// 后端返回原始路由(未处理时)
declare interface RequestOldRoutesState {
	requestOldRoutes: string[];
}

// TagsView 路由列表
declare interface TagsViewRoutesState<T = any> {
	tagsViewRoutes: T[];
}

// 路由列表
declare interface RoutesListState<T = any> {
	routesList: T[];
}

// 布局配置
declare interface ThemeConfigState {
	themeConfig: {
		isCollapse: boolean;
		animation: string;
		globalTitle: string;
		globalViceTitle: string;
		globalViceTitleMsg: string;
		globalI18n: string;
		globalComponentSize: string;
	};
}
