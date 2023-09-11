/* eslint-disable */
import * as axios from 'axios';

// 扩展 axios 数据返回类型，可自行扩展
declare module 'axios' {
	export interface AxiosResponse<T = any> {
        /**
        * 响应码
        */
		code: number;
        /**
        * 数据
        */
		data: T;
        /**
        * 响应消息
        */
		msg: string;
		//type?: string;
		//[key: string]: T;
	}
}
