package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"kratos-ent-example/app/user/service/internal/data/ent"
	"kratos-ent-example/app/user/service/internal/data/ent/user"

	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	"github.com/tx7do/go-utils/mapper"

	pagination "github.com/tx7do/go-curd/api/gen/go/pagination/v1"

	userV1 "kratos-ent-example/api/gen/go/user/service/v1"
)

type UserRepo struct {
	data *Data
	log  *log.Helper

	mapper *mapper.CopierMapper[userV1.User, ent.User]
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/user-service"))
	repo := &UserRepo{
		data:   data,
		log:    l,
		mapper: mapper.NewCopierMapper[userV1.User, ent.User](),
	}

	repo.init()

	return repo
}

func (r *UserRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
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

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		int32(req.GetPage()), int32(req.GetPageSize()), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreatedAt,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]*userV1.User, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListUserResponse{
		Total: int32(count),
		Items: dtos,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	entity, err := r.data.db.Client().User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.mapper.ToDTO(entity), err
}

func (r *UserRepo) Create(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	entity, err := r.data.db.Client().User.Create().
		SetNillableUserName(req.User.UserName).
		SetNillableNickName(req.User.NickName).
		SetPassword(cryptoPassword).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapper.ToDTO(entity), err
}

func (r *UserRepo) Update(ctx context.Context, req *userV1.UpdateUserRequest) (*userV1.User, error) {
	builder := r.data.db.Client().User.UpdateOneID(req.Id).
		SetNillableNickName(req.User.NickName).
		SetUpdatedAt(time.Now())

	if req.User.Password != nil {
		cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}

	entity, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *UserRepo) Upsert(ctx context.Context, req *userV1.UpdateUserRequest) error {
	builder := r.data.db.Client().User.Create().
		SetNillableNickName(req.User.NickName).
		SetCreatedAt(time.Now())

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
			u.SetUpdatedAt(time.Now())
		})

	err := builder.Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (r *UserRepo) Delete(ctx context.Context, req *userV1.DeleteUserRequest) (bool, error) {
	err := r.data.db.Client().User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}

func (r *UserRepo) SQLDelete(ctx context.Context, req *userV1.DeleteUserRequest) (bool, error) {
	args := []any{req.GetId()}
	err := r.data.db.Exec(ctx, "DELETE FROM users WHERE id = $1", args, nil)
	return err != nil, err
}

func (r *UserRepo) SQLGet(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
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

	resp := &userV1.User{}
	if err = rows.Scan(&resp.UserName, &resp.NickName); err != nil {
		return nil, err
	}

	return resp, err
}
