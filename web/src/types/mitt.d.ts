/**
 * mitt 事件类型定义
 *
 * @method setSendClassicChildren 经典布局，开启切割菜单时，菜单数据传入到 navMenu 下的菜单中
 * @method getBreadcrumbIndexSetFilterRoutes 布局设置弹窗，开启切割菜单时，菜单数据传入到 navMenu 下的菜单中
 * @method layoutMobileResize 浏览器窗口改变时，用于适配移动端界面显示
 */
declare type MittType<T = any> = {
	setSendClassicChildren: T;
	getBreadcrumbIndexSetFilterRoutes?: string;
	layoutMobileResize: T;
};

// mitt 参数类型定义
declare type LayoutMobileResize = {
	layout: string;
	clientWidth: number;
};

// mitt 参数菜单类型
declare type MittMenu = {
	children: RouteRecordRaw[];
	item?: RouteItem;
};
