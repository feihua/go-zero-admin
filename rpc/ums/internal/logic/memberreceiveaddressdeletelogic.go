package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberReceiveAddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressDeleteLogic {
	return &MemberReceiveAddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressDeleteLogic) MemberReceiveAddressDelete(in *ums.MemberReceiveAddressDeleteReq) (*ums.MemberReceiveAddressDeleteResp, error) {
	err := l.svcCtx.UmsMemberReceiveAddressModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberReceiveAddressDeleteResp{}, nil
}
