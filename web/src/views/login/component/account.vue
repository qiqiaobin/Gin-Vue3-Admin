<template>
	<el-form size="large" class="login-content-form">
		<el-form-item class="login-animation1">
			<el-input text placeholder="用户名 admin 或不输均为 common" v-model="state.loginForm.userName" clearable autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-User /></el-icon>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation2">
			<el-input :type="state.isShowPassword ? 'text' : 'password'" placeholder="密码：123456" v-model="state.loginForm.password" autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-Unlock /></el-icon>
				</template>
				<template #suffix>
					<i
						class="iconfont el-input__icon login-content-password"
						:class="state.isShowPassword ? 'icon-yincangmima' : 'icon-xianshimima'"
						@click="state.isShowPassword = !state.isShowPassword"
					>
					</i>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation3">
			<el-col :span="15">
				<el-input text maxlength="4" placeholder="请输入验证码" v-model="state.loginForm.captchaCode" clearable autocomplete="off">
					<template #prefix>
						<el-icon class="el-input__icon"><ele-Position /></el-icon>
					</template>
				</el-input>
			</el-col>
			<el-col :span="1"></el-col>
			<el-col :span="8">
				<el-button class="login-content-code" v-waves>
					<img :src="state.captchaBase64"  class="login-content-code-img" @click="getCaptchaImage" style="width: 100%; height: 100%" title="看不清？点击更换" />
				</el-button>
			</el-col>
		</el-form-item>
		<el-form-item class="login-animation4">
			<el-button type="primary" class="login-content-submit" round v-waves @click="onSignIn" :loading="state.loading.signIn">
				<span>登 录</span>
			</el-button>
		</el-form-item>
	</el-form>
</template>

<script setup lang="ts" name="loginAccount">
import { reactive, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import Cookies from 'js-cookie';
import { initBackEndControlRoutes } from '/@/router/backEnd';
import { Session } from '/@/utils/storage';
import { formatAxis } from '/@/utils/formatTime';
import { NextLoading } from '/@/utils/loading';
import authApi from '/@/api/system/auth';

// 定义变量内容
const route = useRoute();
const router = useRouter();
const state = reactive({
	isShowPassword: false,
	captchaBase64: '',
	loginForm: {
		userName: 'superAdmin',
		password: '123456',
		captchaCode: '',
		captchaId: '',
	},
	loading: {
		signIn: false,
	},
});

// 页面加载时
onMounted(() => {
	getCaptchaImage();
});

// 获取验证码图片
const getCaptchaImage = () => {
	authApi.getCaptcha().then((res: any) => {
		const data = res.data;
		state.loginForm.captchaId = data.captchaid;
		const img = data.imgdata;
		if (img.indexOf('data:image') > -1) {
			state.captchaBase64 = img;
		} else {
			state.captchaBase64 = 'data:image/png;base64,' + img;
		}
	});
};

// 时间获取
const currentTime = computed(() => {
	return formatAxis(new Date());
});
// 登录
const onSignIn = async () => {
  state.loading.signIn = true;
  let res;
  try {
    res = await authApi.login(state.loginForm);
  } catch (e) {
    state.loading.signIn = false;
    state.loginForm.captchaId = "";
    getCaptchaImage();
    return;
  }
	//var res = await authApi.login(state.loginForm);
  //let loginRes = loginRespon.data;
	// 存储 token 到浏览器缓存
	Session.set('token', res.data);
	// 模拟数据，对接接口时，记得删除多余代码及对应依赖的引入。用于 `/src/stores/userInfo.ts` 中不同用户登录判断（模拟数据）
	Cookies.set('username', state.loginForm.userName);
	// 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
	// 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
	const isNoPower = await initBackEndControlRoutes();
  signInSuccess(isNoPower);
};
// 登录成功后的跳转
const signInSuccess = (isNoPower: boolean | undefined) => {
  if (isNoPower) {
		ElMessage.warning('抱歉，您没有登录权限');
		Session.clear();
	} else {
    // 初始化登录成功时间问候语
    let currentTimeInfo = currentTime.value;
	  // TODO：带参数、重定向路径页面测试
	  // 登录成功，跳到转首页
	  // 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
	  //console.log(route);
	  if (route.query?.redirect) {
		  router.push({
			  path: <string>route.query?.redirect,
			  query: Object.keys(<string>route.query?.params).length > 0 ? JSON.parse(<string>route.query?.params) : '',
		  });
	  } else {
		  router.push('/');
	  }

	  // 登录成功提示
	  const signInText = '欢迎回来！';
	  ElMessage.success(`${currentTimeInfo}，${signInText}`);
	  // 添加 loading，防止第一次进入界面时出现短暂空白
	  NextLoading.start();
  }
	state.loading.signIn = false;
};
</script>

<style scoped lang="scss">
.login-content-form {
	margin-top: 20px;
	@for $i from 1 through 4 {
		.login-animation#{$i} {
			opacity: 0;
			animation-name: error-num;
			animation-duration: 0.5s;
			animation-fill-mode: forwards;
			animation-delay: calc($i/10) + s;
		}
	}
	.login-content-password {
		display: inline-block;
		width: 20px;
		cursor: pointer;
		&:hover {
			color: #909399;
		}
	}
	.login-content-submit {
		width: 100%;
		letter-spacing: 2px;
		font-weight: 300;
		margin-top: 15px;
	}
  .login-content-code {
    width: 100%;
		padding: 0;
		font-weight: bold;
		letter-spacing: 5px;

    .login-content-code-img {
      width: 100%;
      height: 40px;
      line-height: 40px;
      background-color: #ffffff;
      border: 1px solid rgb(220, 223, 230);
      color: #333;
      font-size: 16px;
      font-weight: 700;
      letter-spacing: 5px;
      text-indent: 5px;
      text-align: center;
      cursor: pointer;
      transition: all ease 0.2s;
      border-radius: 4px;
      user-select: none;

      &:hover {
        border-color: #c0c4cc;
        transition: all ease 0.2s;
      }
    }
  }
}
</style>
