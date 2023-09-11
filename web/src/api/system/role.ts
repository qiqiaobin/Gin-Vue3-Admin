import request from '/@/utils/request';

/**
 * 角色接口
 */
 class RoleApi {
	/**
	 * 角色查询
	 * @param parameter 
	 * @returns 
	 */
	query() {
		return request({
			url: '/system/role/query',
			method: 'get',
			//params: parameter
		});
	}
	/**
	 * 新增角色
	 * @param parameter 
	 * @returns 
	 */
	add(parameter:any) {
		return request({
			url: '/system/role/add',
			method: 'post',
			data: parameter
		});
	}
	/**
	 * 更新角色
	 * @param parameter 
	 * @returns 
	 */
    update(id:number, data:any) {
		return request({
			url: `/system/role/${id}/update`,
			method: 'put',
			data: data,
		});
	}
	/**
	 * 删除角色
	 * @param parameter 
	 * @returns 
	 */
    roledel(id:number) {
		return request({
			url: `/system/role/${id}`,
			method: 'delete',
		});
	}
	/**
	 * 角色详情
	 * @param parameter 
	 * @returns 
	 */
    getdetail(id:number) {
		return request({
			url: `/system/role/${id}`,
			method: 'get',
			//params: parameter
		});
	}
	/**
	 * 角色列表
	 * @returns 
	 */
	list() {
		return request({
			url: '/system/role/list',
			method: 'get',
		});
	}
	/**
	 * 角色菜单权限
	 * @returns 
	 */
	authMenu(parameter:any) {
		return request({
			url: '/system/role/authMenu',
			method: 'get',
			params: parameter
		});
	}
    /**
	 * 保存角色菜单权限
	 * @returns 
	 */
    saveRoleMenu(id:number, data:any) {
		return request({
			url: `/system/role/${id}/update`,
			method: 'put',
			data: data,
		});
	}
}

export default new RoleApi()