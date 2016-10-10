package backend

const (
	INSERT_INTO_CINEMA = `
		insert into cinema(id, name, url, chain, lastUpdate)
		values (default, $1, $2, $3, $4)
	`

	INSERT_INTO_FILM = `
		insert into film(id, imdbFilmId, title, year,rating, imdbCertainty)
		values (default, $1, $2, $3, 4, $5) returning id
	`

	INSERT_INTO_SCREENING = `
		insert into screening(id, cinemaId, filmId, hall, filmDuration, screeningDuration, language, 		 
							  completeTitle, isActive, 	hasSubtitles, hasDub, isImax, is3D, is4D, lastUpdate)
		values(default, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning id 		 
	`

	INSERT_INTO_TICKET = `
		insert into ticket(id, screeningId, cinemaId, type, bookingURL, price, currency, lastUpdate)
		values(default, $1, $2, $3, $4, $5, $6, $7)
	`
)
