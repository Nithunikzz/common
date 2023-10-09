package interceptor

import (
	"context"

	userpb "github.com/Nithunikzz/common/userproto"
)

func (a Adapter) GetUserByOryID(ctx context.Context, in *userpb.GetUserByOryIDRequest) (*userpb.GetUserInfo, error) {
	response, err := a.userapi.GetUserByOryID(in.OryId)
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserInfo{
		Userid:    response.ID,
		Firstname: response.Firstname,
		Lastname:  response.Lastname,
		Email:     response.Email,
		OryId:     response.OryID,
	}, nil
}
