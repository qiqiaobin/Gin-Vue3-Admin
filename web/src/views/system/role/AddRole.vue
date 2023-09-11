<template>
	<div class="system-role-dialog-container">
		<el-dialog title="添加角色" v-model="state.dialog.isShowDialog" width="769px" @close="closeDialog">
			<el-form ref="AddFormRef" :model="state.AddForm" size="default" label-width="90px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="角色名称" prop="nickname">
							<el-input v-model="state.AddForm.nickname" placeholder="请输入角色名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="角色代码" prop="name">
						<el-input v-model="state.AddForm.name" placeholder="请输入角色代码" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="角色描述" prop="note">
						<el-input v-model="state.AddForm.note" type="textarea" placeholder="请输入角色描述"
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
import { reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import roleApi from '/@/api/system/role';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const AddFormRef = ref();
const state = reactive({
	AddForm: {
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
		title: '',
	},
});

// 打开弹窗
const openDialog = () => {
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
	// 重置表单
	AddFormRef.value.resetFields();
};

// 提交
const onSubmit = () => {
	AddFormRef.value.validate(async (valid: boolean) => {
		if (!valid) {
			return;
		}

		//添加
		roleApi.add(state.AddForm).then(() => {
			ElMessage.success('添加成功');
			emit('refresh'); //刷新页面
			closeDialog();
		})
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
