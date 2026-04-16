package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"testing"
	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"
)

func TestDelGoodLogic_DelGood(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *users.DelGoodRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.DelGoodResponse
		wantErr bool
	}{
		{
			name: "测试成功",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: nil,
				Logger: nil,
			},
			args: args{in: &users.DelGoodRequest{
				UserId:  1,
				GoodsId: 1,
			}},
			want: &users.DelGoodResponse{
				Success: true,
			},
			wantErr: false,
		},
		{
			name: "测试失败",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: nil,
				Logger: nil,
			},
			args: args{in: &users.DelGoodRequest{
				UserId:  1,
				GoodsId: 1,
			}},
			want: &users.DelGoodResponse{
				Success: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DelGoodLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.DelGood(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelGood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelGood() got = %v, want %v", got, tt.want)
			}
		})
	}
}
