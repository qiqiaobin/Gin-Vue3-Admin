import { defineStore } from 'pinia';

/**
 * TagsView 路由列表
 * @methods setTagsViewRoutes 设置 TagsView 路由列表
 */
export const useTagsViewRoutes = defineStore('tagsViewRoutes', {
	state: (): TagsViewRoutesState => ({
		tagsViewRoutes: [],
	}),
	actions: {
		async setTagsViewRoutes(data: Array<string>) {
			this.tagsViewRoutes = data;
		}
	},
});
