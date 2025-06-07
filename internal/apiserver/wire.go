//go:build wireinject
// +build wireinject

package apiserver

import (
	"github.com/LiangNing7/goutils/pkg/authz"
	"github.com/google/wire"

	"github.com/LiangNing7/miniblog/internal/apiserver/biz"
	"github.com/LiangNing7/miniblog/internal/apiserver/pkg/validation"
	"github.com/LiangNing7/miniblog/internal/apiserver/store"
	ginmw "github.com/LiangNing7/miniblog/internal/pkg/middleware/gin"
	"github.com/LiangNing7/miniblog/internal/pkg/server"
)

func InitializeWebServer(*Config) (server.Server, error) {
	wire.Build(
		wire.NewSet(NewWebServer, wire.FieldsOf(new(*Config), "ServerMode")),
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
		wire.NewSet(store.ProviderSet, biz.ProviderSet),
		ProvideDB, // 提供数据库实例
		validation.ProviderSet,
		wire.NewSet(
			wire.Struct(new(UserRetriever), "*"),
			wire.Bind(new(ginmw.UserRetriever), new(*UserRetriever)),
		),
		authz.ProviderSet,
	)
	return nil, nil
}
