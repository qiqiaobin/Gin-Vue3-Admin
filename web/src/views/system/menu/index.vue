<template>
	<div class="system-menu-container layout-pd">
		<el-card shadow="hover">
			<div class="system-menu-search mb15">
				<!--el-input size="default" placeholder="请输入菜单名称" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button-->
				<el-button size="default" type="primary" class="ml10" @click="openEditDialog()" v-permission="['system_menu_add']" plain>
					<el-icon>
						<ele-Plus />
					</el-icon>
					新增菜单
				</el-button>
			</div>
			<el-table
				:data="state.tableData.data"
				v-loading="state.tableData.loading"
				style="width: 100%"
				row-key="path"
				:tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
			>
				<el-table-column label="菜单名称" show-overflow-tooltip>
					<template #default="scope">
						<SvgIcon :name="scope.row.icon" />
						<span class="ml10">{{ scope.row.menuName }}</span>
					</template>
				</el-table-column>
				<el-table-column prop="path" label="路由路径" show-overflow-tooltip></el-table-column>
                <el-table-column prop="redirect" label="重定向路径" show-overflow-tooltip></el-table-column>
				<el-table-column prop="component" label="组件路径" show-overflow-tooltip></el-table-column>
				<el-table-column prop="permission" label="权限标识" show-overflow-tooltip></el-table-column>
				<el-table-column prop="sort" label="排序" show-overflow-tooltip></el-table-column>
				<el-table-column prop="status" label="菜单状态" show-overflow-tooltip>
					<template #default="scope">
						<el-tag v-if="scope.row.status == 0">启用</el-tag>
						<el-tag type="info" v-else>禁用</el-tag>
					</template>
				</el-table-column>
				<el-table-column label="操作" show-overflow-tooltip width="140">
					<template #default="scope">
						<el-button text type="primary" @click="openEditDialog(scope.row)" v-permission="['system_menu_update']">修改</el-button>
						<el-button text type="primary" @click="handleDel(scope.row)" v-permission="['system_menu_delete']">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-card>
		<EditDialog ref="editFormRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="systemMenu">
import { defineAsyncComponent, ref, onMounted, reactive } from 'vue';
import { RouteRecordRaw } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import menuApi from '/@/api/system/menu';

// 引入组件
const EditDialog = defineAsyncComponent(() => import('/@/views/system/menu/dialog.vue'));

// 定义变量内容
const editFormRef = ref();
const state = reactive({
	tableData: {
		data: [],
		loading: true,
	},
});

// 获取路由数据，真实请从接口获取
const getTableData = () => {
    state.tableData.loading = true;
	menuApi.tree().then((res) => {
		if (res.success) {
			state.tableData.data = res.data;
			state.tableData.loading = false;
		}
	});
};
// 打开新增菜单弹窗
const openEditDialog = (row?: any) => {
	editFormRef.value.openDialog(row);
};
// 打开编辑菜单弹窗
const onOpenEditMenu = (type: string, row: RouteRecordRaw) => {
	editFormRef.value.openDialog(type, row);
};
const handleDel = (row: any) => {
	ElMessageBox.confirm(`此操作将永久删除菜单：${row.menuName}`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			menuApi.delete({ id: row.id }).then((res) => {
				if (res.success) {
					getTableData();
					ElMessage.success('删除成功');
				}
			});
		})
		.catch(() => {});
};
// 页面加载时
onMounted(() => {
	getTableData();
});
</script>
