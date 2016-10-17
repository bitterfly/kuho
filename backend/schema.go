package backend

/*
Cinema(_id, name<string>, url<string>, chain<string>, lastUpdate<timestamp>)

Screening(_id, cinemaId<int>, filmId<int>, hall<string>, duration<duration>, language<string>, completeTitle<string>, isActive<bool>, HasSubtitles<bool>, hasDub<bool>, isImax<bool>, is3D<bool>, is4D<bool>, lastUpdate<duration>)

Film(_id, imdbFilmId<int>, title<string>, year<int>, rating<string>, imdbCertainty<float32>)

IMDBFilm(_id, title<string>, year<int>, poster<string>, imdbRating<float32>, numberOfVotes<int>, releaseDate<timestamp>, tagline<string>, laguages string>, imdbFilmDuration<duration>, lastUpdate<timestamp>)

Plot(_id, IMDBFilmId<bigint>, length<string>, description<string>)

ForegnTitle(_id, IMDBFilmId<bigint>, country<string>, title<string>)

Ticket(_id, screeningId<int>, cinemaId<int>, type<string>, bookingURL<string>, price<float32>, currency<string>, lastUpdate<timestamp>)
*/

const schema = `
	create table cinema(
		id 			 bigserial primary key,
		name 		 varchar(256),
		shortName	 varchar(25) unique,
		url 		 varchar(512),
		chain 		 varchar(256),
		lastUpdate   timestamp
	);

	create table imdbFilm(
		id 			  bigint primary key,
		title 		  varchar(256),
		year 		  integer,
		poster 		  varchar(512),
		imdbRating 	  real,
		numberOfVotes integer,
		releaseDate   timestamp,
		tagline       varchar(256),
		laguages      varchar(512),
		filmDuration  bigint,
		lastUpdate    timestamp
	);

	create table film(
		id 			  bigserial primary key,
		imdbFilmId    bigint references imdbFilm(id) unique, 
		title 		  varchar(256),
		year 		  integer,
		rating 		  varchar(256),
		imdbCertainty real,
		
		unique(title, year)
	);

	create table screening(
		id                bigserial primary key,
		cinemaId 	 	  integer references cinema(id), 
		filmId 		 	  bigint references film(id),
		hall		 	  varchar(256),
		duration 		  bigint,
		language 		  varchar(256),
		completeTitle 	  varchar(256),
		isActive 		  boolean,
		hasSubtitles 	  boolean,
		hasDub 			  boolean,
		isImax 			  boolean,
		is3D 			  boolean,
		is4D 			  boolean,
		lastUpdate 		  timestamp
	);

	
	create table ticket(
		id 				 bigserial primary key,
		screeningId  bigint references screening(id),
		cinemaId     bigint references cinema(id),
		type 		 varchar(256),
		bookingURL   varchar(512),
		price 		 real,
		currency 	 varchar(256),
		lastUpdate 	 timestamp
	);


	create table foreignTitle(
		id bigserial primary key,
		country 	 varchar(256),
		title 	     varchar(256),
		imdbFilmId   bigint references imdbFilm(id)
	);

	create table plot(
		id bigserial primary key,
		length 		 varchar(256),
		description  varchar(2048),
		imdbFilmId 	 bigint references imdbFilm(id)
	);
`

const dropSchema = `
	drop table ticket;
	drop table foreignTitle;
	drop table plot;
	drop table screening;
	drop table film;
	drop table imdbFilm;
	drop table cinema;;
`
