package {{cookiecutter.app_name}}

import (
	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao"
	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

/*
一个业务单元聚合:
	- 聚合一堆业务逻辑: 所有业务层的 logic, 都在此层, 与 dao 层区别
	- g: 公共数据资源对象, 包含数据层所有可操作资源
	- bizScope list: 根据业务拆分的单元, 聚合相关 biz logic
*/
type Domain struct {
	// inner global use:
	g *dao.MetaResource

	// biz:
	Hello *HelloScope
	Block *BlockScope
}

//
func NewDomain(cfg config.Config, isRpcServer bool) *Domain {
	// global:
	g := dao.NewMetaResource(cfg, isRpcServer)

	return &Domain{
		g: g,

		// biz:
		Hello: newAuthScope(g),
		Block: newBlockScope(g),
	}
}

func (m *Domain) Close() {
	m.g.Close()
	return
}
