package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/crypto"
	"github.com/tx7do/go-utils/mapper"

	pagination "github.com/tx7do/go-curd/api/gen/go/pagination/v1"
	entCurd "github.com/tx7do/go-curd/entgo"

	"kratos-ent-example/app/user/service/internal/data/ent"
	"kratos-ent-example/app/user/service/internal/data/ent/predicate"
	"kratos-ent-example/app/user/service/internal/data/ent/user"

	userV1 "kratos-ent-example/api/gen/go/user/service/v1"
)

type UserRepo struct {
	data *Data
	log  *log.Helper

	mapper     *mapper.CopierMapper[userV1.User, ent.User]
	repository *entCurd.Repository[
		ent.UserQuery, ent.UserSelect, ent.UserCreate, ent.UserCreateBulk, ent.UserUpdate, ent.UserUpdateOne, ent.UserDelete,
		predicate.User,
		userV1.User, ent.User,
	]
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/user-service"))
	repo := &UserRepo{
		data:   data,
		log:    l,
		mapper: mapper.NewCopierMapper[userV1.User, ent.User](),
	}

	repo.repository = entCurd.NewRepository[
		ent.UserQuery, ent.UserSelect, ent.UserCreate, ent.UserCreateBulk, ent.UserUpdate, ent.UserUpdateOne, ent.UserDelete,
		predicate.User,
		userV1.User, ent.User,
	](repo.mapper)

	repo.init()

	return repo
}

func (r *UserRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Debug().User.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	return builder.Count(ctx)
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	builder := r.data.db.Client().Debug().User.Query()

	ret, err := r.repository.ListWithPaging(ctx, builder, builder.Clone(), req)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return &userV1.ListUserResponse{Total: 0, Items: nil}, nil
	}

	return &userV1.ListUserResponse{
		Total: ret.Total,
		Items: ret.Items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	var whereCond []func(s *sql.Selector)
	switch req.QueryBy.(type) {
	case *userV1.GetUserRequest_Id:
		whereCond = append(whereCond, user.IDEQ(req.GetId()))
	case *userV1.GetUserRequest_Username:
		whereCond = append(whereCond, user.UserNameEQ(req.GetUsername()))
	default:
		whereCond = append(whereCond, user.IDEQ(req.GetId()))
	}

	builder := r.data.db.Client().Debug().User.Query()
	dto, err := r.repository.Get(ctx, builder, whereCond, req.GetViewMask())
	if err != nil {
		return nil, err
	}

	return dto, err
}

func (r *UserRepo) Create(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.User, error) {
	if req == nil || req.User == nil {
		return nil, errors.New("request is nil")
	}

	if req.User.Password != nil && req.User.GetPassword() != "" {
		cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
		if err != nil {
			return nil, err
		}
		req.User.Password = &cryptoPassword
	}

	builder := r.data.db.Client().Debug().User.Create()
	result, err := r.repository.Create(ctx, builder, req.User, nil, func(dto *userV1.User) {
		builder.
			SetNillableUserName(req.User.UserName).
			SetNillableNickName(req.User.NickName).
			SetCreatedAt(time.Now())

		if req.User.Password != nil {
			builder.SetPassword(req.User.GetPassword())
		}
	})

	return result, err
}

func (r *UserRepo) Update(ctx context.Context, req *userV1.UpdateUserRequest) (*userV1.User, error) {
	if req == nil || req.User == nil {
		return nil, errors.New("request is nil")
	}

	if req.User.Password != nil && req.User.GetPassword() != "" {
		cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
		if err != nil {
			return nil, err
		}
		req.User.Password = &cryptoPassword
	}

	builder := r.data.db.Client().Debug().User.UpdateOneID(req.User.GetId())
	result, err := r.repository.UpdateOne(ctx, builder, req.User, req.GetUpdateMask(),
		[]predicate.User{
			func(s *sql.Selector) {
				s.Where(sql.EQ(user.FieldID, req.User.GetId()))
			},
		},
		func(dto *userV1.User) {
			builder.
				SetNillableNickName(req.User.NickName).
				SetUpdatedAt(time.Now())

			if req.User.Password != nil {
				builder.SetPassword(req.User.GetPassword())
			}
		},
	)

	return result, err
}

func (r *UserRepo) Upsert(ctx context.Context, req *userV1.UpdateUserRequest) error {
	if req == nil || req.User == nil {
		return errors.New("request is nil")
	}

	if req.User.Password != nil && req.User.GetPassword() != "" {
		cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
		if err != nil {
			return err
		}
		req.User.Password = &cryptoPassword
	}

	builder := r.data.db.Client().Debug().User.Create().
		SetNillableNickName(req.User.NickName).
		SetCreatedAt(time.Now())

	builder.
		OnConflict(
			sql.ConflictColumns(user.FieldID),
		).
		Update(func(u *ent.UserUpsert) {
			if req.User.NickName != nil {
				u.SetNickName(req.User.GetNickName())
			}
			if req.User.Password != nil {
				u.SetPassword(req.User.GetPassword())
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
	if req == nil {
		return false, errors.New("request is nil")
	}

	builder := r.data.db.Client().Debug().User.Delete()
	affected, err := r.repository.Delete(ctx, builder, []predicate.User{
		func(s *sql.Selector) {
			s.Where(sql.EQ(user.FieldID, req.GetId()))
		},
	})

	return err == nil && affected > 0, err
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
