package data

import (
	"context"

	"testing"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/go-utils/trans"

	_ "github.com/xiaoqidun/entps"

	"kratos-ent-example/app/user/service/internal/data/ent"

	"kratos-ent-example/api/gen/go/common/conf"
	userV1 "kratos-ent-example/api/gen/go/user/service/v1"
)

func createTestEntClient() *entgo.EntClient[*ent.Client] {
	client := NewEntClient(&conf.Bootstrap{
		Data: &conf.Data{
			Database: &conf.Data_Database{
				Driver:  "sqlite3",
				Source:  "file:ent?mode=memory&cache=shared&_fk=1",
				Migrate: true,
			},
		},
	}, log.DefaultLogger)

	return client
}

func TestUserRepo_CRUD_List_Count(t *testing.T) {
	ctx := context.Background()
	client := createTestEntClient()

	// 构造 Data（依赖包中 Data 的具体类型需包含名为 db 的字段且其类型是接口或可赋值的）
	d := &Data{
		db: client,
	}

	repo := NewUserRepo(d, log.DefaultLogger)

	// 1) Create
	createReq := &userV1.CreateUserRequest{
		Data: &userV1.User{
			UserName: trans.Ptr("alice"),
			NickName: trans.Ptr("Alice"),
			Password: trans.Ptr("password123"),
		},
	}

	created, err := repo.Create(ctx, createReq)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if created == nil || created.GetUserName() != "alice" {
		t.Fatalf("unexpected created user: %#v", created)
	}
	if created.Password == nil || *created.Password == "password123" {
		t.Fatalf("password should be hashed")
	}

	// 2) Count
	cnt, err := repo.Count(ctx, nil)
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}
	if cnt != 1 {
		t.Fatalf("expected count 1, got %d", cnt)
	}

	// 3) Get by id
	getReq := &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{Id: created.GetId()},
	}
	got, err := repo.Get(ctx, getReq)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if got == nil || got.GetUserName() != "alice" {
		t.Fatalf("unexpected get result: %#v", got)
	}

	// 4) Update nickname and password
	updatedNick := "Alice-Updated"
	updateReq := &userV1.UpdateUserRequest{
		Data: &userV1.User{
			Id:       created.GetId(),
			NickName: &updatedNick,
			Password: trans.Ptr("newpass"),
		},
		UpdateMask: nil, // 依赖 repository 的实现，传 nil 表示整体更新
	}
	updated, err := repo.Update(ctx, updateReq)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updated == nil || updated.GetNickName() != updatedNick {
		t.Fatalf("unexpected updated user: %#v", updated)
	}
	if updated.Password == nil || *updated.Password == "newpass" {
		t.Fatalf("updated password should be hashed")
	}

	// 5) List with paging
	listReq := &pagination.PagingRequest{
		Page:  trans.Ptr(uint32(1)),
		Limit: trans.Ptr(uint32(10)),
	}
	listResp, err := repo.List(ctx, listReq)
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if listResp == nil || listResp.Total == 0 || len(listResp.Items) == 0 {
		t.Fatalf("unexpected list response: %#v", listResp)
	}

	// 6) Delete
	delReq := &userV1.DeleteUserRequest{Id: created.GetId()}
	ok, err := repo.Delete(ctx, delReq)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if !ok {
		t.Fatalf("expected delete ok")
	}

	// ensure gone
	cntAfter, err := repo.Count(ctx, nil)
	if err != nil {
		t.Fatalf("Count after delete failed: %v", err)
	}
	if cntAfter != 0 {
		t.Fatalf("expected count 0 after delete, got %d", cntAfter)
	}
}
