package backend

import (
	"fmt"

	"github.com/bitterfly/kuho/spiderdata"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Backend struct {
	db *sqlx.DB
}

func New(dbURN string) (*Backend, error) {
	var db *sqlx.DB
	db, err := sqlx.Connect("postgres", dbURN)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %s", err)
	}

	return &Backend{
		db: db,
	}, nil
}

func (b *Backend) Foo() (string, error) {
	return "this is foo.", nil
}

func (b *Backend) InitDB() error {
	_, err := b.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("unable to execute schema: %s", err)
	}
	return nil
}

func (b *Backend) DropDB() error {
	_, err := b.db.Exec(dropSchema)
	if err != nil {
		return fmt.Errorf("unable to drop schema: %s", err)
	}
	return nil
}

func (b *Backend) Fill(data *spiderdata.Request) error {
	_, err = b.Wrap(func(tx *sqlx.Tx) (interface{}, error) {
		// first clear all of the old data

		cinemas := data.Cinemas
		films := data.Films

		// insert cinemas
		for i := range cinamas {
			err = insertCinema(tx, cinemas[i])
			if err != nil {
				return nil, fmt.Errorf("unable to insert cinema: %s", err)
			}
		}

		// insert films
		for i := range films {
			err = insertFilm(tx, films[i])
			if err != nil {
				return nil, fmt.Errorf("unable to insert timetable: %s", err)
			}
		}
		return nil, nil
	})
	return err
}

func insertCinema(tx *sqlx.Tx, cinema spiderdata.Cinema) error {
	//name, url, chain, lastUpdate
	_, err := tx.NamedExec(INSERT_INTO_CINEMA, cinema.Name, cinema.URL, cinema.Chain, cinema.Acquired)

	return err
}

func inserFilm(tx *sqlx.Tx, film spiderdata.Film) error {
	screenings := film.Screenings

	var filmID int64
	//imdbFilmId, title, year,rating, imdbCertainty
	err := tx.Get(&filmID, INSERT_INTO_FILM, nil, film.Title, film.Year, film.Rating, film.ImdbIDCertainty)

	if err != nil {
		return err
	}

	var screeningID int64
	for i := range screenings {
		screening := screenings[i]

		// cinemaId, filmId, hall, duration, language,
		//					  completeTitle, isActive, 	hasSubtitles, hasDub, isImax, is3D, is4D, lastUpdate
		err = tx.Get(&screeningID, INSERT_INTO_SCREENING,
			screening.CinemaID, filmID, screening.Hall, screening.Duration, screening.Language,
			screening.Variant, screening.Active, screening.IsSubtitled, screening.IsDubbed,
			screening.IsImax, screening.Is3D, screening.Is4D, screening.Acquired,
		)

		if err != nil {
			return err
		}

		screeningTickets := screening.Tickets

		for j := range screeningTickets {
			ticket := screeningTickets[j]

			//screeningId, cinemaId, type, bookingURL, price, currency, lastUpdate
			_, err = tx.NamedExec(INSERT_INTO_TICKET, screeningID, screening.CinemaID,
				ticket.Type, ticket.BookingURL, ticket.Price, ticket.Currency, ticket.Acquired,
			)

			if err != nil {
				return err
			}

		}
	}
	return nil
}
