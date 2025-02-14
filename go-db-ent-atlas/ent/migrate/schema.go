// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "model", Type: field.TypeString},
		{Name: "registered_at", Type: field.TypeTime},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
	}
	// FoosColumns holds the columns for the "foos" table.
	FoosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "foo_foo", Type: field.TypeString},
	}
	// FoosTable holds the schema information for the "foos" table.
	FoosTable = &schema.Table{
		Name:       "foos",
		Columns:    FoosColumns,
		PrimaryKey: []*schema.Column{FoosColumns[0]},
	}
	// HogesColumns holds the columns for the "hoges" table.
	HogesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// HogesTable holds the schema information for the "hoges" table.
	HogesTable = &schema.Table{
		Name:       "hoges",
		Columns:    HogesColumns,
		PrimaryKey: []*schema.Column{HogesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		FoosTable,
		HogesTable,
	}
)

func init() {
}
