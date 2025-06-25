// Package orm provides reusable database-related structures and functionalities.
// It defines common base model structures, such as `MModel` and `PModel`,
// to standardize fields like unique identifiers, timestamps, and soft deletion
// across different database backends.
package orm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MModel represents a base model for MySQL database tables.
// It defines a standard set of fields, including a UUID primary key,
// timestamps for creation and updates, and support for soft deletion.
// This struct can be embedded in other models to ensure consistency
// and reusability across MySQL-based applications.
//
//	User represents a specific application model that embeds MModel.
//	type User struct {
//		MModel    // Embeds the common fields from MModel
//		Name      string `json:"name" gorm:"column:NAME;type:varchar(255);not null"`
//		Email     string `json:"email" gorm:"column:EMAIL;type:varchar(255);unique;not null"`
//		IsActive  bool   `json:"isActive" gorm:"column:IS_ACTIVE;type:tinyint(1);default:1"`
//	}
type MModel struct {
	// ID is the primary key for the record, represented as a universally unique identifier (UUID).
	// It is stored as a string in varchar(36) format and is automatically generated
	// using a UUID generator function like MySQL's UUID() or UUID_TO_BIN(UUID()) for binary storage.
	ID string `json:"id" gorm:"column:ID;primaryKey;type:varchar(36);not null"`

	// CreatedAt stores the timestamp indicating when the record was created.
	// This value is automatically set to the current timestamp at creation time.
	// It is indexed to optimize queries based on record creation time.
	CreatedAt time.Time `json:"createdAt" gorm:"column:CREATED_AT;type:datetime(6);default:CURRENT_TIMESTAMP(6);index:IDX_CREATED_AT;<-:create"`

	// UpdatedAt stores the timestamp of the most recent modification to the record.
	// This value is automatically updated to the current timestamp whenever the record is modified.
	// It is indexed to allow efficient queries for recently updated records.
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:UPDATED_AT;type:datetime(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);index:IDX_UPDATED_AT"`

	// DeletedAt is used for soft deletion of records. When a record is soft-deleted,
	// this field is populated with the deletion timestamp, but the record itself remains in the database.
	// By default, GORM excludes records with a non-null DeletedAt field from queries.
	// It is indexed to optimize filtering of active or soft-deleted records.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:DELETED_AT;type:datetime(6);default:NULL;index:IDX_DELETED_AT"`
}

// BeforeCreate is a GORM hook that runs before a new record is inserted into the database.
// This function generates a new UUID for the ID field of the MModel struct.
func (mModel *MModel) BeforeCreate(*gorm.DB) error {
	mModel.ID = uuid.New().String()
	return nil
}

// PModel is a base model for PostgreSQL database tables.
// Similar to `MModel`, it provides a standard set of fields, including a UUID
// primary key, timestamps for creation and updates, and soft deletion support.
// This struct can be embedded in other models to enforce consistency and
// streamline development for PostgreSQL-based applications.
type PModel struct {
	// ID is a universally unique identifier (UUID) that serves as the primary key for the record.
	// It is stored as a `uuid` type in PostgreSQL and is automatically generated
	// using the `uuid_generate_v4()` function. This ensures uniqueness and efficient indexing.
	ID uuid.UUID `json:"id" gorm:"column:ID;primaryKey;type:uuid;not null;default:uuid_generate_v4()"`

	// CreatedAt is a timestamp indicating when the record was created.
	// It is stored in PostgreSQL as a `timestamp` with microsecond precision (`timestamp(6)`).
	// This field is automatically set to the current timestamp at the time of creation
	// and is indexed for efficient queries based on creation time.
	CreatedAt time.Time `json:"createdAt" gorm:"column:CREATED_AT;type:timestamp(6);default:CURRENT_TIMESTAMP(6);index"`

	// UpdatedAt is a timestamp indicating the last time the record was updated.
	// Like `CreatedAt`, it uses `timestamp(6)` for microsecond precision and is
	// automatically updated whenever the record is modified. This field is indexed
	// for optimized queries involving recently updated records.
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:UPDATED_AT;type:timestamp(6);default:CURRENT_TIMESTAMP(6);index"`

	// DeletedAt is a field used for soft deletion of records. When a record is soft-deleted,
	// this field stores the timestamp of deletion, while the record remains in the database.
	// GORM's built-in soft delete functionality automatically manages this field, enabling
	// queries to include or exclude soft-deleted records as needed. It is indexed to support
	// efficient filtering of active and deleted records.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:DELETED_AT;type:timestamp(6);index"`
}
