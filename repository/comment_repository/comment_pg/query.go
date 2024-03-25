package comment_pg

const (
	addCommentQuery = `
		INSERT INTO
			comments
				(
					user_id,
					photo_id,
					message
				)
		VALUES
				(
					$1, 
					$2, 
					$3
				)
		RETURNING
			id, message, photo_id, user_id, created_at
	`

	getCommentQuery = `
		SELECT 
			c.id,
			c.user_id,
			c.photo_id,
			c.message,
			c.created_at,
			c.updated_at,
			u.id,
			u.username,
			u.email,
			p.id,
			p.title,
			p.caption,
			p.photo_url,
			p.user_id
		FROM 
			comments AS c
		LEFT JOIN
			users AS u
		ON
			c.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			c.photo_id = p.id
		ORDER BY 
			c.id
		ASC
	`

	getCommentByIdQuery = `
		SELECT 
			c.id,
			c.user_id,
			c.photo_id,
			c.message,
			c.created_at,
			c.updated_at,
			u.id,
			u.username,
			u.email,
			p.id,
			p.title,
			p.caption,
			p.photo_url,
			p.user_id
		FROM 
			comments AS c
		LEFT JOIN
			users AS u
		ON
			c.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			c.photo_id = p.id
		WHERE c.id = $1
	`

	deleteCommentQuery = `
		DELETE FROM
			comments
		WHERE
			id = $1
	`

	updateCommentQuery = `
		UPDATE 
			comments 
		SET
			message = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, message, photo_id, user_id, updated_at
	`
)
