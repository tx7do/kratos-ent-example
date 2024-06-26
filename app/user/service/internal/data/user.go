package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"

	"kratos-ent-example/app/user/service/internal/biz"
	"kratos-ent-example/app/user/service/internal/data/ent"
	"kratos-ent-example/app/user/service/internal/data/ent/user"

	"kratos-ent-example/gen/api/go/common/pagination"
	v1 "kratos-ent-example/gen/api/go/user/service/v1"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/user-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

func (r *UserRepo) convertEntToProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	return &v1.User{
		Id:         in.ID,
		UserName:   in.UserName,
		NickName:   in.NickName,
		Password:   in.Password,
		CreateTime: util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime: util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	return builder.Count(ctx)
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	res, err := r.data.db.Client().User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *UserRepo) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	res, err := r.data.db.Client().User.Create().
		SetNillableUserName(req.User.UserName).
		SetNillableNickName(req.User.NickName).
		SetPassword(cryptoPassword).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *UserRepo) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	builder := r.data.db.Client().User.UpdateOneID(req.Id).
		SetNillableNickName(req.User.NickName).
		SetUpdateTime(time.Now().UnixMilli())

	if req.User.Password != nil {
		cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}

	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *UserRepo) Upsert(ctx context.Context, req *v1.UpdateUserRequest) error {
	builder := r.data.db.Client().User.Create().
		SetNillableNickName(req.User.NickName).
		SetCreateTime(time.Now().UnixMilli())

	if req.User.Password != nil {
		cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}

	builder.
		OnConflict(
			sql.ConflictColumns(user.FieldID),
		).
		Update(func(u *ent.UserUpsert) {
			if req.User.NickName != nil {
				u.SetNickName(req.User.GetNickName())
			}
			if req.User.Password != nil {
				cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
				if err == nil {
					u.SetPassword(cryptoPassword)
				}
			}
			u.SetUpdateTime(time.Now().UnixMilli())
		})

	err := builder.Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (r *UserRepo) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	err := r.data.db.Client().User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}

func (r *UserRepo) SQLDelete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	args := []any{req.GetId()}
	err := r.data.db.Exec(ctx, "DELETE FROM users WHERE id = $1", args, nil)
	return err != nil, err
}

func (r *UserRepo) SQLGet(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	args := []any{req.GetId()}

	var err error
	var rows sql.Rows
	if err = r.data.db.Query(ctx, "SELECT * FROM users WHERE id = $1", args, &rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, ent.MaskNotFound(errors.New("cannot found user"))
	}

	resp := &v1.User{}
	if err = rows.Scan(&resp.UserName, &resp.NickName); err != nil {
		return nil, err
	}

	return resp, err
}
