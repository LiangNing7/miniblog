package conversion

import (
	"github.com/LiangNing7/goutils/pkg/core"

	"github.com/LiangNing7/miniblog/internal/apiserver/model"
	apiv1 "github.com/LiangNing7/miniblog/pkg/api/apiserver/v1"
)

// UserModelToUserV1 将模型层的 UserM（用户模型对象）转换为 Protobuf 层的 User（v1 用户对象）.
func UserModelToUserV1(userModel *model.UserM) *apiv1.User {
	var protoUser apiv1.User
	_ = core.CopyWithConverters(&protoUser, userModel)
	return &protoUser
}

// UserV1ToUserModel 将 Protobuf 层的 User（v1 用户对象）转换为模型层的 UserM（用户模型对象）.
func UserV1ToUserModel(protoUser *apiv1.User) *model.UserM {
	var userModel model.UserM
	_ = core.CopyWithConverters(&userModel, protoUser)
	return &userModel
}
