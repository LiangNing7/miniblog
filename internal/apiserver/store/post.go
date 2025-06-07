// Copyright 2025 LiangNing7 <LiangNing7@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/LiangNing7/miniblog.

// nolint: dupl
package store

import (
	"context"

	gerericstore "github.com/LiangNing7/goutils/pkg/store"
	"github.com/LiangNing7/goutils/pkg/store/where"

	"github.com/LiangNing7/miniblog/internal/apiserver/model"
)

// PostStore 定义了 post 模块在 store 层所实现的方法.
type PostStore interface {
	Create(ctx context.Context, obj *model.PostM) error
	Update(ctx context.Context, obj *model.PostM) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.PostM, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.PostM, error)

	PostExpansion
}

// PostExpansion 定义了帖子操作的附加方法.
type PostExpansion interface{}

// postStore 是 PostStore 接口的实现.
type postStore struct {
	*gerericstore.Store[model.PostM]
}

// 确保 postStore 实现了 PostStore 接口.
var _ PostStore = (*postStore)(nil)

// newPostStore 创建 postStore 的实例.
func newPostStore(store *datastore) *postStore {
	return &postStore{
		Store: gerericstore.NewStore[model.PostM](store, NewLogger()),
	}
}
