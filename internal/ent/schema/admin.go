package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/facebook/ent/schema/mixin"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").NotEmpty().Default("").Comment("用户名称"),
		field.String("true_name").Default("").Comment("真实姓名"),
		field.String("mobile").Default("").Comment("手机号"),
	}
}

// Indexs 定义索引
func (Admin) Indexs() []ent.Index {
	return []ent.Index{
		index.Fields("user_name"),
		index.Fields("mobile"),
	}
}

// Mixin 定义通用字段
func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		StatusMixin{},
		DelStatusMixin{},
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
