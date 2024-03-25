package photo_pg

const (
	addNewPhotoQuery = `
		INSERT INTO
			photos
				(
					title,
					caption,
					photo_url,
					user_id
				)
		VALUES
				($2, $3, $4, $1)
		RETURNING
				id, title, caption, photo_url, user_id, created_at
	`

	getUserAndPhotos = `
				SELECT
					p.id,
					p.title,
					p.caption,
					p.photo_url,
					p.user_id,
					p.created_at,
					p.updated_at,
					u.email,
					u.username
				FROM
					photos as p
				LEFT JOIN
					users AS u
				ON
					p.user_id = u.id
				ORDER BY
					p.id
				ASC
	`
	getUserAndPhotosById = `
		SELECT
			p.id,
			p.title,
			p.caption,
			p.photo_url,
			p.user_id,
			p.created_at,
			p.updated_at,
			u.email,
			u.username
		FROM
			photos as p
		LEFT JOIN
			users AS u
		ON
			p.user_id = u.id
		WHERE
				p.id = $1
		ORDER BY
			p.id
		ASC
	`

	UpdatePhotoQuery = `
		UPDATE
			photos
		SET
			title = $2,
			caption = $3,
			photo_url = $4,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
				id, title, caption, photo_url, user_id, updated_at
	`

	deletePhotoById = `
		DELETE FROM
			photos
		WHERE
			id = $1		
	`
)
