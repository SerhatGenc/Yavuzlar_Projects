-- name: AddBook :exec
INSERT INTO books (bookname, writer, user_id, deadline)
VALUES ($1, $2, $3, $4);

-- name: ListBooks :many
SELECT * FROM books;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: UpdateBook :exec
UPDATE books
SET
  bookname = CASE WHEN $2 = '' THEN bookname ELSE $2 END,
  writer = CASE WHEN $3 = '' THEN writer ELSE $3 END,
  user_id = CASE WHEN $4 = 0 THEN user_id ELSE $4 END,
  deadline = CASE WHEN $5 = '' THEN deadline ELSE $5 END
WHERE id = $1;


-- name: ListDeadlineBooks :many
SELECT * FROM books
WHERE deadline < $1;

-- name: AddUser :exec
INSERT INTO users (email, password, name, surname, roles)
VALUES ($1, $2, $3, $4, $5);

-- name: ListUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET
  email = CASE WHEN $2 = '' THEN email ELSE $2 END,
  password = CASE WHEN $3 = '' THEN password ELSE $3 END,
  name = CASE WHEN $4 = '' THEN name ELSE $4 END,
  surname = CASE WHEN $5 = '' THEN surname ELSE $5 END,
  roles = CASE WHEN $6 = '' THEN roles ELSE $6 END
WHERE id = $1;

-- name: ControlUser :one
SELECT * FROM users
WHERE email = $1 AND password = $2;
