package dto

type AuthMenuVo struct {
	Path      string           `json:"path"` //路由地址
	Redirect  string           `json:"redirect"`
	Name      string           `json:"name"`      //路由名称（TODO：暂时不清楚前端这个字段有什么用处）
	Meta      MenuRouterMetaVo `json:"meta"`      //Mate
	Component string           `json:"component"` //组件路径
	Children  []AuthMenuVo     `json:"children"`  //子菜单
}

type MenuRouterMetaVo struct {
	Title       string `json:"title"`       //菜单标题
	Icon        string `json:"icon"`        //菜单图标
	IsLink      string `json:"isLink"`      //外部地址
	IsIframe    bool   `json:"isIframe"`    //是否内嵌
	IsHide      bool   `json:"isHide"`      //是否隐藏
	IsKeepAlive bool   `json:"isKeepAlive"` //是否缓存
	IsAffix     bool   `json:"isAffix"`     //是否固定在 tagsView 栏上
}

type UserPermissionVo struct {
	Permission    string `json:"permission"`    //按钮权限
	RequestMethod string `json:"requestMethod"` //请求方法
}
