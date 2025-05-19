// Copyright 2025 LiangNing7 <LiangNing7@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/LiangNing7/miniblog.

package gin

import (
	"context"

	"github.com/LiangNing7/goutils/pkg/core"
	"github.com/LiangNing7/goutils/pkg/token"
	"github.com/gin-gonic/gin"

	"github.com/LiangNing7/miniblog/internal/apiserver/model"
	"github.com/LiangNing7/miniblog/internal/pkg/contextx"
	"github.com/LiangNing7/miniblog/internal/pkg/errno"
	"github.com/LiangNing7/miniblog/internal/pkg/log"
)

// UserRetriever 用于根据用户名获取用户的接口.
type UserRetriever interface {
	// GetUser 根据用户ID获取用户信息
	GetUser(ctx context.Context, userID string) (*model.UserM, error)
}

// AuthnMiddleware 是一个认证中间件，用于从 gin.Context 中提取 token 并验证 token 是否合法.
func AuthnMiddleware(retriever UserRetriever) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		userID, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, nil, errno.ErrTokenInvalid.WithMessage("%s", err.Error()))
			c.Abort()
			return
		}

		log.Debugw("Token parsing successful", "userID", userID)

		user, err := retriever.GetUser(c, userID)
		if err != nil {
			core.WriteResponse(c, nil, errno.ErrUnauthenticated.WithMessage("%s", err.Error()))
			c.Abort()
			return
		}

		ctx := contextx.WithUserID(c.Request.Context(), user.UserID)
		ctx = contextx.WithUsername(ctx, user.Username)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
