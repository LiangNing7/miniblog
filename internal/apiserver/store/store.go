// Copyright 2025 LiangNing7 <LiangNing7@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/LiangNing7/miniblog.

package store

//go:generate mockgen -destination mock_store.go -package store github.com/LiangNing7/miniblog/internal/apiserver/store IStore,UserStore,PostStore,ConcretePostStore

import (
	"context"
	"sync"

	"github.com/LiangNing7/goutils/pkg/store/where"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet 是一个 Wire 的 Provider 集合，用于声明依赖注入的规则.Add commentMore actions
// 包含 NewStore 构造函数，用于生成 datastore 实例.
// wire.Bind 用于将接口 IStore 与具体实现 *datastore 绑定，
// 从而在依赖 IStore 的地方，能够自动注入 *datastore 实例.
var ProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

var (
	once sync.Once
	// 全局变量，方便其他包直接调用已初始化好的 datastore 实例.
	S *datastore
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	// 返回 Store 层的 *gorm.DB 实例，在少数场景下会被用到.
	DB(ctx context.Context, wheres ...where.Where) *gorm.DB
	TX(ctx context.Context, fn func(ctx context.Context) error) error

	User() UserStore
	Post() PostStore
	// ConcretePost 是一个示例的 store 实现，用来演示在 Go 中如何直接与 DB 交互.
	ConcretePost() ConcretePostStore
}

// transactionKey 用于在 context.Context 中存储事务上下文的键.
type transactionKey struct{}

// datastore 是 IStore 的具体实现.
type datastore struct {
	core *gorm.DB

	// 可以根据需要添加其他数据库实例.
	// fake *gorm.DB
}

// 确保 datastore 实现了 IStore 接口.
var _ IStore = (*datastore)(nil)

// NewStore 创建一个 IStore 类型的实例.
func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次.
	once.Do(func() {
		S = &datastore{db}
	})
	return S
}

// DB 根据传入的条件（wheres）对数据库实例进行筛选.
// 如果未传入任何条件，则返回上下文中的数据库实例（事务实例或核心数据库实例）.
func (store *datastore) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
	db := store.core
	// 从上下文中提取事务实例
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}

	// 遍历所有传入的条件并逐一叠加到数据库查询对象上
	for _, whr := range wheres {
		db = whr.Where(db)
	}
	return db
}

// TX 返回一个新的事务实例.
// nolint: fatcontext
func (store *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	// WithContext 用于在 context 中注入事务实例.
	return store.core.WithContext(ctx).
		// Transaction 用于开启事务，传入回调 fn.
		Transaction(func(tx *gorm.DB) error {
			// 上层传入的 ctx 被打包上了这个事务 tx,
			// 如果后续调用 store.DB 则会优先使用这个事务连接.
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		})
}

// Users 返回一个实现了 UserStore 接口的实例.
func (store *datastore) User() UserStore {
	return newUserStore(store)
}

// Posts 返回一个实现了 PostStore 接口的实例.
func (store *datastore) Post() PostStore {
	return newPostStore(store)
}

// ConcretePost 返回一个实现了 ConcretePostStore 接口的实例.
func (store *datastore) ConcretePost() ConcretePostStore {
	return newConcretePostStore(store)
}
