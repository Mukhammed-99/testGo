// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "id_tt", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "pets_tts_id_tt",
				Columns: []*schema.Column{PetsColumns[2]},

				RefColumns: []*schema.Column{TtsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "pets_users_pets",
				Columns: []*schema.Column{PetsColumns[3]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TtsColumns holds the columns for the "tts" table.
	TtsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TtsTable holds the schema information for the "tts" table.
	TtsTable = &schema.Table{
		Name:        "tts",
		Columns:     TtsColumns,
		PrimaryKey:  []*schema.Column{TtsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PetsTable,
		TtsTable,
		UsersTable,
	}
)

func init() {
	PetsTable.ForeignKeys[0].RefTable = TtsTable
	PetsTable.ForeignKeys[1].RefTable = UsersTable
}
