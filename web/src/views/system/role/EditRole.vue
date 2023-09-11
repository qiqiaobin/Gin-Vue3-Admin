<template>
	<div class="system-role-dialog-container">
		<el-dialog title="编辑角色" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="dataFormRef" :model="state.dataForm" size="default" label-width="90px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="角色名称">
							<el-input v-model="state.dataForm.nickname" placeholder="请输入角色名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="角色代码" prop="name">
						<el-input v-model="state.dataForm.name" placeholder="请输入角色代码" :disabled="state.dataForm.name" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="角色描述" prop="note">
						<el-input v-model="state.dataForm.note" type="textarea" placeholder="请输入角色描述"
							maxlength="150"></el-input>
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

<script setup lang="ts" name="systemRoleDialog">
import { defineProps, reactive, ref, nextTick } from 'vue';
import { ElMessage } from 'element-plus';
import roleApi from '/@/api/system/role';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
const props = defineProps(["rid"]);

// 定义变量内容
const dataFormRef = ref();
const state = reactive({
	dataForm: {
		nickname: '', // 显示名称
		name: '', // 角色名称
		note: '', // 角色描述
	},
	rules: {
		nickname: [{ required: true, message: '角色显示名称不能为空', trigger: 'blur' }],
		name: [{ required: true, message: '角色名称不能为空', trigger: 'blur' }],
	},
	dialog: {
		isShowDialog: false,
	},
});

// 打开弹窗
const openDialog = (id: any) => {
	nextTick(() => {
		getRoleDetail(id);
	});

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

		//更新
		roleApi.update(props.rid,state.dataForm).then(() => {
			ElMessage.success('更新成功');
			emit('refresh'); //刷新页面
			closeDialog();
		});
    });
};

// 获取角色详情
const getRoleDetail = (roleId: number) => {
    roleApi.getdetail(roleId).then((res) => {
		state.dataForm = res.data;
	});
};

// 暴露变量
defineExpose({
	openDialog,
});
</script>

<style scoped lang="scss">
.system-role-dialog-container {
	.menu-data-tree {
		width: 100%;
		border: 1px solid var(--el-border-color);
		border-radius: var(--el-input-border-radius, var(--el-border-radius-base));
		padding: 5px;
	}
}
</style>
