import request from '/@/utils/request';

/**
 * 用户接口
 */
class UserApi {
	/**
	 * 用户查询
	 * @param parameter 
	 * @returns 
	 */
	query(parameter: any) {
		return request({
			url: '/api/n9e/users',
			method: 'get',
			params: parameter,
		});
	}
	/**
	 * 新增用户
	 * @param parameter 
	 * @returns 
	 */
	add(parameter: any) {
		return request({
			url: '/api/n9e/users',
			method: 'post',
			data: parameter,
		});
	}
	/**
	 * 更新用户
	 * @param parameter 
	 * @returns 
	 */
    userupdate(id:number, data:any) {
		return request({
			url: `/api/n9e/${id}/update`,
			method: 'put',
			data: data,
		});
	}
	/**
	 * 删除用户
	 * @param parameter 
	 * @returns 
	 */
    userdel(id:number) {
		return request({
			url: `/api/n9e/${id}`,
			method: 'delete',
		});
	}
	/**
	 * 用户详情
	 * @param parameter 
	 * @returns 
	 */
    getdetail(id:number) {
		return request({
			url: `/api/n9e/${id}`,
			method: 'get',
			//params: parameter
		});
	}
	/**
	 * 用户列表
	 * @returns 
	 */
	list() {
		return request({
			url: '/api/n9e/list',
			method: 'get',
		});
	}
    /**
	 * 重置密码
	 * @param parameter 
	 * @returns 
	 */
    RsetPwd(id:number, data:any) {
        return request({
            //url: '/system/user/'+uid+ '/password',
            url: `/api/n9e/${id}/password`, 
            method: 'put',
            data: data,
        });
    }
}

export default new UserApi()