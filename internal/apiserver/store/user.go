// Copyright 2025 LiangNing7 <LiangNing7@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/LiangNing7/miniblog.

// nolint: dupl
package store

import (
	"context"

	genericstore "github.com/LiangNing7/goutils/pkg/store"
	"github.com/LiangNing7/goutils/pkg/store/where"

	"github.com/LiangNing7/miniblog/internal/apiserver/model"
)

type UserStore interface {
	Create(ctx context.Context, obj *model.UserM) error
	Update(ctx context.Context, obj *model.UserM) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.UserM, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.UserM, error)

	UserExpansion
}

// UserExpansion 定义了用户操作的附加方法.
// nolint: iface
type UserExpansion interface{}

// userStore 是 UserStore 接口的实现.
type userStore struct {
	*genericstore.Store[model.UserM]
}

// 确保 userStore 实现了 UserStore 接口.
var _ UserStore = (*userStore)(nil)

// newUserStore 创建 userStore 的实例.
func newUserStore(store *datastore) *userStore {
	return &userStore{
		Store: genericstore.NewStore[model.UserM](store, NewLogger()),
	}
}
