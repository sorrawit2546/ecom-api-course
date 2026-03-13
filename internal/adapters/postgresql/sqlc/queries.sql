-- name: ListProducts :many
SELECT
    *
FROM
    products;

-- name: ListProduct :one
SELECT
    *
FROM
    products
WHERE
    id = $1;