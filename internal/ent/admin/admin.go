// Code generated by entc, DO NOT EDIT.

package admin

import (
	"time"
)

const (
	// Label holds the string label denoting the admin type in the database.
	Label = "admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDelStatus holds the string denoting the del_status field in the database.
	FieldDelStatus = "del_status"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldTrueName holds the string denoting the true_name field in the database.
	FieldTrueName = "true_name"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"

	// Table holds the table name of the admin in the database.
	Table = "admins"
)

// Columns holds all SQL columns for admin fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStatus,
	FieldDelStatus,
	FieldUserName,
	FieldTrueName,
	FieldMobile,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// DefaultDelStatus holds the default value on creation for the "del_status" field.
	DefaultDelStatus int
	// DefaultUserName holds the default value on creation for the "user_name" field.
	DefaultUserName string
	// DefaultTrueName holds the default value on creation for the "true_name" field.
	DefaultTrueName string
	// DefaultMobile holds the default value on creation for the "mobile" field.
	DefaultMobile string
)
