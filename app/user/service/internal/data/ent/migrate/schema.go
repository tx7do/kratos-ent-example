// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "user_name", Type: field.TypeString, Unique: true, Nullable: true, Size: 50, Comment: "用户名"},
		{Name: "nick_name", Type: field.TypeString, Nullable: true, Size: 128, Comment: "昵称"},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 255, Comment: "登陆密码"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "用户账号",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[0]},
			},
			{
				Name:    "user_id_user_name",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0], UsersColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
	UsersTable.Annotation = &entsql.Annotation{
		Table:     "users",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}