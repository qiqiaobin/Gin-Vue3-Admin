<template>
	<div class="system-user-dialog-container">
        <el-dialog title="创建用户" v-model="state.dialog.isShowDialog" width="769px" @close="closeDialog">
			<el-form ref="dataFormRef" :model="state.dataForm" :rules="state.rules" size="default" label-width="90px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="账户名称" prop="username">
							<el-input v-model="state.dataForm.username" placeholder="请输入账号名称" :disabled="state.dialog.isEdit"
							clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="用户昵称" prop="nickname">
						<el-input v-model="state.dataForm.nickname" placeholder="请输入用户昵称" clearable></el-input>
					</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="手机号码" prop="phone">
						<el-input v-model="state.dataForm.phone" placeholder="请输入手机号" clearable></el-input>
					</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="邮箱" prop="email">
						<el-input v-model="state.dataForm.email" placeholder="请输入邮箱" clearable></el-input>
					</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="账户密码" prop="password">
						<el-input v-model="state.dataForm.password" placeholder="请输入账户密码" type="password"
							show-password></el-input>
					</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="用户角色" prop="roleIds">
						<el-select v-model="state.dataForm.roleIds" multiple collapse-tags collapse-tags-tooltip
							:max-collapse-tags="2" placeholder="请选择角色" style="width: 100%">
							<el-option v-for="item in state.roleOption" :key="item.id" :label="item.name"
								:value="item.nickname" />
						</el-select>
					</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="closeDialog" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">确 定</el-button>
				</span>
			</template>
        </el-dialog>
	</div>
</template>

<script setup lang="ts" name="systemUserDialog">
import { reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import userApi from '/@/api/system/user';
import roleApi from '/@/api/system/role';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const dataFormRef = ref();
const state = reactive({
	dataForm: {
		id: 0, //用户id
		username: '', // 账号名称
		nickname: '', // 用户昵称
		phone: '', // 手机号
		email: '', // 邮箱
		password: '123456', // 账户密码
		roleIds: [],
	},
	rules: {
		username: [{ required: true, message: '账号名称不能为空', trigger: 'blur' }],
		nickname: [{ required: true, message: '用户昵称不能为空', trigger: 'blur' }],
		email: [{ type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }],
		phone: [{ pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: '请输入正确的手机号码', trigger: 'blur' }],
		password: [{ pattern: /^(?=.*[0-9])(.{6,18})$/, message: '请输入6-18位密码', trigger: 'blur' }],
	},

	dialog: {
        isShowDialog: false,
		isEdit: false,
		title: '',
	},
    roleOption: [] as any,
});

// 打开弹窗
const openDialog = () => {
	getRoleOption();
    state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
    state.dialog.isShowDialog = false;
    // 重置表单
	dataFormRef.value.resetFields();
};

// 提交
const onSubmit = () => {
	dataFormRef.value.validate(async (valid: boolean) => {
		if (!valid) {
			return;
		}

		//添加
		userApi.add(state.dataForm).then(() => {
			ElMessage.success('添加成功');
			emit('refresh'); //刷新页面
			closeDialog();
		});
	});
};
const getRoleOption = () => {
	roleApi.list().then((res) => {
		state.roleOption = res.data;
	});
};

// 暴露变量
defineExpose({
	openDialog,
});
</script>
