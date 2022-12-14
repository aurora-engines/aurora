package route

import "gitee.com/aurora-engine/aurora/web"

const (
	DefaultType = iota
	RESTFulType
)

// node 路由节点
type node struct {
	Path       string           //当前节点路径
	FullPath   string           //当前处理器全路径
	NodeType   int              //节点类型
	Count      int              //路径数量
	middleware []web.Middleware //中间处理函数
	Control    *Controller      //服务处理函数
	Child      []*node          //子节点
}
