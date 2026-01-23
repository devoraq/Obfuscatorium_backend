package user

import (
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	userpb "github.com/devoraq/Obfuscatorium_backend/pkg/gen/go/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoUser(u *models.User) *userpb.User {
	if u == nil {
		return nil
	}

	out := &userpb.User{
		Id:        u.ID.String(),
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: timestamppb.New(u.CreatedAt),
	}

	out.Avatar = u.AvatarURL
	out.Bio = u.Bio

	return out
}
