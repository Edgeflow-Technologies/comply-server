-- name: CreateFramework :one
INSERT INTO frameworks (
  name, description, version, locked, editable, category
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetFramework :one
SELECT * FROM frameworks
WHERE id = $1;

-- name: ListFrameworks :many
SELECT * FROM frameworks
ORDER BY id;

-- name: UpdateFramework :one
UPDATE frameworks
SET
  name = COALESCE($2, name),
  description = COALESCE($3, description),
  version = COALESCE($4, version),
  locked = COALESCE($5, locked),
  editable = COALESCE($6, editable),
  category = COALESCE($7, category),
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteFramework :exec
DELETE FROM frameworks WHERE id = $1;