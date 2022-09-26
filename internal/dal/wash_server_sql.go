// Code generated by mtgroup-generator.
package dal

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

const (
	sqlGetWashServer = `
	SELECT
		id,
		created_at,
		modified_at,
		service_key,
		name,
		description,
		owner_id
	FROM
		wash_servers
	WHERE
		id=:id AND
		NOT deleted
	`

	sqlAddWashServer = `
	INSERT INTO wash_servers(
		name,
		description,
		owner_id,
		created_at,
	) VALUES (
		:name,
		:description,
		:owner_id,
		:created_at
	)
	RETURNING
		id
	`

	sqlDeleteWashServer = `
	UPDATE
		wash_servers
	SET
		deleted=true,
		deleted_at=:deleted_at,
		deleted_by=:deleted_by
	WHERE
		id=:id AND
		NOT deleted
	`

	sqlEditWashServer = `
	UPDATE
		wash_servers
	SET
		service_key=:service_key,
		description=:description,
		owner_id=:owner_id,
		name=:name,
		modified_at=:modified_at,
		modified_by=:modified_by
	WHERE
		id=:id AND
		NOT deleted
	`

	sqlListWashServer = `
	SELECT
		id,
		created_at,
		modified_at,
		service_key,
		name,
		description,
		owner_id
	FROM
		wash_servers
	WHERE
		NOT deleted
	`
)

type (
	argGetWashServer struct {
		ID string `db:"id"`
	}

	argAddWashServer struct {
		Name        string    `db:"name"`
		Description string    `db:"description"`
		OwnerID     string    `db:"owner_id"`
		CreatedAt   time.Time `db:"created_at"`
	}

	argEditWashServer struct {
		ServiceKey  string        `db:"service_key"`
		Name        string        `db:"name"`
		Description string        `db:"description"`
		OwnerID     string        `db:"owner_id"`
		ModifiedAt  time.Time     `db:"modified_at"`
		ModifiedBy  uuid.NullUUID `db:"modified_by"`
	}

	argDeleteWashServer struct {
		ID        string        `db:"id"`
		DeletedAt time.Time     `db:"deleted_at"`
		DeletedBy uuid.NullUUID `db:"deleted_by"`
	}
)
