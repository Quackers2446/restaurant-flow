-- name: GetUser :one
select * from user where user.user_id = unhex(replace(sqlc.arg("user_id"),'-',''));

-- name: GetUserSessions :many
select session_id, user_id, ip_addr, user_agent, created_at, last_used_at, expires_at from `session` where user_id = ?;

-- name: GetAuthDetails :one
select * from user_auth where user_id = unhex(replace(sqlc.arg("user_id"),'-',''));

-- name: GetLoginDetails :one
select *, sqlc.embed(user_auth) from user
inner join user_auth on user_auth.uid = user.uid
where user.email = ?;

-- name: CreateSession :exec
insert into `session` set
    `user_id`=unhex(replace(sqlc.arg("user_id"),'-','')),
    `ip_addr`=?,
    `user_agent`=?,
    `expires_at`=?,
    `refresh_key`=?;
