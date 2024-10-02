-- name: CreateOrder :one
INSERT INTO orders (
    source_key_id,
    order_uuid,
    plan, 
    amount 
) VALUES (
    $1, $2, $3, $4
) returning *;

-- name: GetOrdersBySourceKeyID :many
SELECT * FROM orders 
WHERE source_key_id = $1;

-- name: GetOrderUUIDsByAnOrderUUIDList :many
SELECT order_uuid FROM orders
WHERE order_uuid = ANY ($1::uuid[]);
