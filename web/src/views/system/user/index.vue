<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<el-form :model="state.tableData.param" ref="queryFormRef" :inline="true" label-width="70px">
                <el-form-item label="用户账号" prop="userName">
				<el-input size="default"
						v-model="state.tableData.param.userName"
						placeholder="请输入用户账号/手机/昵称"
						style="width: 180px"
						@keyup.enter="getUserList"
						clearable
					></el-input>
                </el-form-item>
                <el-form-item>
				<el-button size="default" class="ml10" @click="resetQueryForm()">
					<el-icon>
						<ele-Refresh />
					</el-icon>
					重置
				</el-button>
				<el-button size="default" type="primary" class="ml10" @click="getUserList">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
                <el-button size="default" type="primary" @click="openAddDialog()" v-permission="['system_user_add']" plain>
					<el-icon>
						<ele-Plus />
					</el-icon>
					新增
				</el-button>
            </el-form-item>
			</el-form>
			<el-table 
            :data="state.tableData.data" 
            v-loading="state.tableData.loading" 
            style="width: 100%"
            @selection-change="handleSelectionChange" >
                <el-table-column type="selection" width="55" align="center"/>
				<el-table-column prop="id" label="编号" width="60" fixed></el-table-column>
				<el-table-column prop="username" label="账户账号" fixed show-overflow-tooltip></el-table-column>
				<el-table-column prop="nickname" label="用户昵称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="phone" label="手机号" show-overflow-tooltip></el-table-column>
				<el-table-column prop="email" label="邮箱" show-overflow-tooltip></el-table-column>
                <el-table-column prop="roles" label="角色" show-overflow-tooltip></el-table-column>
				<el-table-column prop="create_at" label="创建时间" show-overflow-tooltip>
                    <template #default="scope">
                        {{ dateFormat(scope.row.create_at) }}
                    </template>
                </el-table-column>
                <el-table-column prop="update_at" label="更新时间" show-overflow-tooltip>
                    <template #default="scope">
                        {{ dateFormat(scope.row.update_at) }}
                    </template>
                </el-table-column>
				<el-table-column label="操作" width="200" fixed="right">
					<template #default="scope">
						<el-button text type="primary" @click="openEditDialog(scope.row.id)" v-permission="['system_user_update']">修改</el-button>
                        <el-button text type="primary"  @click="handleResetPwd(scope.row)" v-permission="['system_user_delete']">重置密码</el-button>
						<el-button text type="primary" @click="handleDel(scope.row)" v-permission="['system_user_delete']">删除</el-button>

					</template>
				</el-table-column>
			</el-table>
			<el-pagination
				@size-change="onHandleSizeChange"
				@current-change="onHandleCurrentChange"
				class="mt15"
				:pager-count="5"
				:page-sizes="[10, 20, 30]"
				v-model:current-page="state.tableData.param.pageNum"
				background
				v-model:page-size="state.tableData.param.pageSize"
				layout="total, sizes, prev, pager, next, jumper"
				:total="state.tableData.total"
			>
			</el-pagination>
		</el-card>
    <!--新增用户弹出框 start-->
    <AddUser ref="userAddRef" @refresh="getUserList()" />
    <!--新增用户弹出框 end-->


    <!--编辑用户弹出框 start-->
	<EditUser :uid="state.EditDialog.id"  ref="userEditRef" @refresh="getUserList()" />
    <!--编辑用户弹出框 end-->
	</div>
</template>

<script setup lang="ts" name="systemUser">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import userApi from '/@/api/system/user';
import { dateFormat } from '/@/utils/formatTime';

// 引入组件
const EditUser = defineAsyncComponent(() => import('/@/views/system/user/EditUser.vue'));
const AddUser = defineAsyncComponent(() => import('/@/views/system/user/AddUser.vue'));

// 定义变量内容
const queryFormRef = ref();
const userEditRef = ref();
const userAddRef = ref();
const state = reactive({
    // 选中数组
    ids: [],
    // 非单个禁用
    single: true,
    // 非多个禁用
    multiple: true,
	tableData: {
		data: [],
		total: 0,
		loading: false,
		param: {
			pageNum: 1,
			pageSize: 10,
            userName: undefined,
		},
	},
    EditDialog: {
        id: 0,
    },
});

// 页面加载时
onMounted(() => {
	getUserList();
});


// 重置表单
const resetQueryForm = () => {
	queryFormRef.value?.resetFields();
	getUserList();
};

// 多选框选中数据
const handleSelectionChange = (selection: any) => {
  state.ids = selection.map((item: any) => item.roleId);
  state.single = selection.length != 1;
  state.multiple = !selection.length;
};

/** 查询用户列表 */
const getUserList = async () => {
	state.tableData.loading = true;
	userApi.query(state.tableData.param).then((res) => {
		state.tableData.data = res.data.list;
		state.tableData.total = res.data.total;
		state.tableData.loading = false;
	});
};

// 打开新增用户弹窗
const openAddDialog = () => {
	userAddRef.value.openDialog();
};

// 打开编辑弹窗
const openEditDialog = (id:number) => {
    state.EditDialog.id = id;
	userEditRef.value.openDialog(id);
};


// 删除用户
const handleDel = (row: any) => {
	if (row.username === "superadmin") {
		ElMessage.error('超级管理员不允许删除');
		return;
	}
	ElMessageBox.confirm(`此操作将永久删除账号：${row.username}`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			//userApi.delete({ id: row.id }).then(() => {
            userApi.userdel(row.id).then(() => {
				getUserList();
				ElMessage.success('删除成功');
			});
		})
		.catch(() => {});
};


//重置密码
const handleResetPwd = (row: any ) => {
	ElMessageBox.prompt(`请输入账号【${row.username}】的新密码 `, '修改密码', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		inputPattern: /^.{6,20}$/,
		inputErrorMessage: '请输入6-20位密码',
	})
        .then(({ value }) => {
        userApi.RsetPwd( row.id, { password: value } )
            ElMessage.success('重置密码成功');
    })
};

// 分页改变
const onHandleSizeChange = (val: number) => {
	state.tableData.param.pageSize = val;
	getUserList();
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
	state.tableData.param.pageNum = val;
	getUserList();
};
</script>

<style scoped lang="scss">
.system-user-container {
	:deep(.el-card__body) {
		display: flex;
		flex-direction: column;
		flex: 1;
		overflow: auto;
		.el-table {
			flex: 1;
		}
	}
}
</style>
