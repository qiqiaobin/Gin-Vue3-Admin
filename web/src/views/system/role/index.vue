<template>
	<div class="system-role-container layout-padding">
		<el-row :gutter="20">
            <!--角色数据-->
            <el-col :span="4" :xs="24">
                <el-card shadow="always">
                <div class="sub-title">
                    角色列表
                    <el-button size="samll" class="ml10" @click="openAddDialog()" text>
					<span>添加角色</span>
				</el-button>
                </div>
                <div class="head-container">
                    <el-input
                    v-model="state.nickname"
                    placeholder="请输入角色名称"
                    clearable
                    prefix-icon="el-icon-search"
                    style="margin-bottom: 20px"
                    />
                </div>
                <div class="head-container">
                    <el-tree
                        :data="state.roleOptions"
                        :props="state.defaultProps"
                        node-key="id"
                        :expand-on-click-node="false"
                        :filter-node-method="filterNode"
                        highlight-current
                        ref="treeRef"
                        @node-click="handleNodeClick"
                    />
                </div>
                </el-card>
            </el-col>

            <el-col :span="20" :xs="24">
		            <el-card shadow="hover">
			            <div class="system-menu-search mb15">
                            <div class="team-info">
                                角色名称
                            <el-link @click="openEditDialog(state.roleinfo.id)" :icon="Edit" class="ml10"  ></el-link>
                            <el-link @click="handleDel(state.roleinfo.id)" :icon="Delete" class="ml10" ></el-link>
                             </div>
			            </div>
 
                        <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
                        <div class="sub-title">
                            菜单列表  
                        </div>
			                <el-tree
                                ref="roleMenuRef" 
                                :data="state.menuData" 
                                :props="state.menuProps" 
                                class="menu-data-tree"		
                                node-key="menuid" 
                                show-checkbox
                                check-on-click-node
                                :height="550"
                                :expand-on-click-node="false"
                            />
                        </el-col>
                        <div >
					<el-button type="primary" @click="savePremission" size="default">保存权限</el-button>
			        </div>
                </el-card>
            </el-col>
		</el-row>
		<EditRole :rid="state.roleinfo.id" ref="roleEditRef" @refresh="getTableData()" />
        <AddRole  ref="roleAddRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="systemRole">
import { defineAsyncComponent, reactive, onMounted, ref, watch } from 'vue';
import { ElMessageBox, ElMessage,ElTree } from 'element-plus';

import { Delete, Edit } from '@element-plus/icons-vue'
import roleApi from '/@/api/system/role';
import menuApi from '/@/api/system/menu';
import AuthApi from '/@/api/system/auth';

// 引入组件
const EditRole = defineAsyncComponent(() => import('/@/views/system/role/EditRole.vue'));
const AddRole = defineAsyncComponent(() => import('/@/views/system/role/AddRole.vue'));

// 定义变量内容
const roleEditRef = ref();
const roleAddRef = ref();
const state = reactive({
    // 角色选项
    roleOptions: [],
    defaultProps: {
        children: "children",
        label: "nickname",
    },
    // 角色名称
    nickname: undefined,
    roleinfo:{
        id: 0,
        name:undefined,
    },
	menuData: [] as any,
	menuProps: {
        value: 'menuid',
        label: 'label',
		children: 'children',
	},
    tableData: {
		data: [],
		total: 0,
		loading: false,
		param: {
			pageNum: 1,
			pageSize: 10,
            id: undefined,
		},
	},
    menuIds: [], // 角色菜单
});

const treeRef = ref<InstanceType<typeof ElTree>>();
const roleMenuRef = ref<InstanceType<typeof ElTree>>();

watch(() => state.nickname, (val) => {
    (treeRef as any).value!.filter(val);
})

// 页面加载时
onMounted(() => {
	getTableData();
    getMenuData();
    //handleNodeClick;
});

// 节点单击事件
const handleNodeClick = (data: any) => {
  state.roleinfo.id = data.id;
  state.roleinfo.name = data.nickname;
  getPremission(state.roleinfo.id);
};

// 获取当前权限
const getPremission = (id:any) => {
	// 请求接口返回权限
    AuthApi.getMenuIds({ id: id }).then((res) => {
        (roleMenuRef as any).value!.setCheckedKeys(res.data);
	});
};

// 筛选节点
const filterNode = (value: string, data: any) => {
  if (!value) return true;
  return data.nickname.includes(value);
};

// 初始化表格数据
const getTableData = () => {
	state.tableData.loading = true;
	roleApi.query().then((res) => {
        state.roleOptions = res.data.list; 
	});
    menuApi.tree().then((res) => {
		state.tableData.data = res.data;
		state.tableData.loading = false;
	});
};
// 打开编辑弹窗
const openEditDialog = (id?:number) => {
	roleEditRef.value.openDialog(id);
};

// 打开新增角色弹窗
const openAddDialog = () => {
	roleAddRef.value.openDialog();
};


// 获取菜单结构数据
const getMenuData = () => {
    menuApi.list().then((res) => {
		state.menuData = buildMenuTree(res.data, 0);
	});
};

// 构建菜单树状（递归）
const buildMenuTree = (menuData: any, parentId: number): any => {
	var menuTree = [];
	for (let index = 0; index < menuData.length; index++) {
		const menu = menuData[index];
		if (menu.parentId == parentId) {
			menuTree.push({
				menuid: menu.id,
				label: menu.menuName,
				children: buildMenuTree(menuData, menu.id),
			});
		}
	}
	return menuTree;
};


// 获取当前权限
const savePremission = () => {
    // 选中节点
    var checkedKeys = roleMenuRef.value!.getCheckedKeys() as any;
	// 半选中节点
	var halfCheckedKeys = roleMenuRef.value!.getHalfCheckedKeys() as any;

    // 需要将半选keys放到前面，否则回显的半选会变成选中
	state.menuIds = halfCheckedKeys.concat(checkedKeys);

    roleApi.saveRoleMenu(state.roleinfo.id,state.menuIds).then((res) => {
        if (res.msg !== "请求成功"){
            ElMessage.error(res.msg);
        }else{
			ElMessage.success('保存权限成功');
        }
	});
}

// 删除角色
const handleDel = (id:number) => {
	ElMessageBox.confirm(`此操作将永久删除角色名称：“${state.roleinfo.name}”?`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			//roleApi.delete({ id: id }).then((res) => {
            roleApi.roledel(id).then((res) => {
				getTableData();
                if (res.msg !== "请求成功"){
                    ElMessage.error(res.msg);
                }else{
					ElMessage.success('删除成功');
                }
			});			
		})
		.catch(() => {});
};

</script>

<style scoped lang="scss">
.system-role-container {
	.system-role-padding {
		padding: 15px;
		.el-table {
			flex: 1;
		}
	}
    .sub-title {
        color: #333;
        font-size: 14px;
        font-weight: 700;
        line-height: 36px;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .team-info {
        border: 1px solid #e0dee2;
        padding: 10px;
        margin-bottom: 10px;
        background: #f6f6f6;
    }
}
</style>
