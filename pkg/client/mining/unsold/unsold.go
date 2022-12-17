//nolint:dupl
package unsold

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/unsold"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get unsold connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateUnsold(ctx context.Context, in *npool.UnsoldReq) (*npool.Unsold, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateUnsold(ctx, &npool.CreateUnsoldRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create unsold: %v", err)
	}
	return info.(*npool.Unsold), nil
}

func CreateUnsolds(ctx context.Context, in []*npool.UnsoldReq) ([]*npool.Unsold, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateUnsolds(ctx, &npool.CreateUnsoldsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create unsolds: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create unsolds: %v", err)
	}
	return infos.([]*npool.Unsold), nil
}

func GetUnsold(ctx context.Context, id string) (*npool.Unsold, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetUnsold(ctx, &npool.GetUnsoldRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get unsold: %v", err)
	}
	return info.(*npool.Unsold), nil
}

func GetUnsoldOnly(ctx context.Context, conds *npool.Conds) (*npool.Unsold, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetUnsoldOnly(ctx, &npool.GetUnsoldOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get unsold: %v", err)
	}
	return info.(*npool.Unsold), nil
}

func GetUnsolds(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Unsold, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetUnsolds(ctx, &npool.GetUnsoldsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get unsolds: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get unsolds: %v", err)
	}
	return infos.([]*npool.Unsold), total, nil
}

func ExistUnsold(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistUnsold(ctx, &npool.ExistUnsoldRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get unsold: %v", err)
	}
	return infos.(bool), nil
}

func ExistUnsoldConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistUnsoldConds(ctx, &npool.ExistUnsoldCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get unsold: %v", err)
	}
	return infos.(bool), nil
}

func CountUnsolds(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountUnsolds(ctx, &npool.CountUnsoldsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count unsold: %v", err)
	}
	return infos.(uint32), nil
}
