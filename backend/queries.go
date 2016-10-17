package backend

const (
	INSERT_INTO_CINEMA = `
		insert into cinema(id, name, shortName, url, chain, lastUpdate)
		values (default, $1, $2, $3, $4, $5)
		on conflict (shortName) do update
		set name = $1, shortName = $2, url = $3, chain = $4, lastUpdate = $5
	`
	INSER_INTO_IMDB_FILM = `
		insert into imdbFilm(
			id, title, year, poster, imdbRating,
			numberOfVotes, releaseDate, tagline,
			languages, filmDuration, lastUpdate
		)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now())
		on conflict (id) do update
		set id = $1, title = $2, year = $3, poster = $4,
		imdbRating = $5, numberOfVotes = $6, releaseDate = $7,
		tagline = $8, languages = $9, filmDuration = $10, lastUpdate = now()
	`

	INSERT_INTO_FILM_NOT_NULL_IMDBID = `
		insert into film(id, imdbFilmId, title, year,rating, imdbCertainty)
		values (default, $1, $2, $3, $4, $5) 
		on conflict (imdbFilmId) do update
		set imdbFilmId = $1, title = $2, year = $3, rating = $4, imdbCertainty = $5
		returning id
	`

	INSERT_INTO_FILM_NULL_IMDBID = `
		insert into film(id, imdbFilmId, title, year,rating, imdbCertainty)
		values (default, $1, $2, $3, $4, $5) 
		on conflict (title, year) do update
		set imdbFilmId = $1, title = $2, year = $3, rating = $4, imdbCertainty = $5
		returning id
	`

	INSERT_INTO_SCREENING = `
		insert into screening(id, cinemaId, filmId, hall, duration, language, 		 
							  completeTitle, isActive, 	hasSubtitles, hasDub, isImax, is3D, is4D, lastUpdate)
		values(default, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) returning id 		 
	`

	INSERT_INTO_TICKET = `
		insert into ticket(id, screeningId, cinemaId, type, bookingURL, price, currency, lastUpdate)
		values(default, $1, $2, $3, $4, $5, $6, $7)
	`

	GET_CINEMA_BY_SHORT_NAME = `
		select id from cinema where cinema.shortName = $1;
	`

	CHECK_NEED_UPDATE_IMDB_FILM = `
		select id from imdbFilm where id = $1 and DATE_PART('day', now() - lastUpdate) < 7;
	`
)
