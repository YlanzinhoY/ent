// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"github.com/facebookincubator/ent/examples/start/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the user in the database.
	Table = "users"
	// CarsTable is the table the holds the cars relation/edge.
	CarsTable = "cars"
	// CarsInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarsInverseTable = "cars"
	// CarsColumn is the table column denoting the cars relation/edge.
	CarsColumn = "owner_id"
	// GroupsTable is the table the holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_users"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "user_id"}
)

var (
	fields = schema.User{}.Fields()

	// descAge is the schema descriptor for age field.
	descAge = fields[0].Descriptor()
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator = descAge.Validators[0].(func(int) error)

	// descName is the schema descriptor for name field.
	descName = fields[1].Descriptor()
	// DefaultName holds the default value on creation for the name field.
	DefaultName = descName.Default.(string)
)
