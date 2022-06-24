// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/admin/admins"
	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/mail"
	"ultimatedivision/seasons"
	"ultimatedivision/store/lootboxes"
	"ultimatedivision/users"
)

// SeedDB provides access to accounts db.
type SeedDB struct {
	users     *usersDB
	clubs     *clubsDB
	cards     *cardsDB
	matches   *matchesDB
	divisions *divisionsDB
}

// NewSeedDB is a constructor for seed db.
func NewSeedDB(conn *sql.DB) *SeedDB {
	return &SeedDB{
		users:     &usersDB{conn: conn},
		clubs:     &clubsDB{conn: conn},
		cards:     &cardsDB{conn: conn},
		matches:   &matchesDB{conn: conn},
		divisions: &divisionsDB{conn: conn},
	}
}

// CreateUser creates a user and writes to the database.
func CreateUser(ctx context.Context, db *sql.DB) error {
	testUser1 := users.User{
		ID:           uuid.New(),
		Email:        "testUser1@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin1",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b1"),
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser2 := users.User{
		ID:           uuid.New(),
		Email:        "testUser2@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin2",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b2"),
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser3 := users.User{
		ID:           uuid.New(),
		Email:        "testUser3@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin3",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b3"),
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser4 := users.User{
		ID:           uuid.New(),
		Email:        "testUser4@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin4",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b4"),
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser5 := users.User{
		ID:           uuid.New(),
		Email:        "testUser5@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin5",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b5"),
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser6 := users.User{
		ID:           uuid.New(),
		Email:        "testUser6@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin6",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b6"),
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUsers := []users.User{testUser1, testUser2, testUser3, testUser4, testUser5, testUser6}

	for _, user := range testUsers {
		err := user.EncodePass()
		if err != nil {
			return Error.Wrap(err)
		}

		emailNormalized := mail.Normalize(user.Email)
		query := `INSERT INTO users(
                  id, 
                  email, 
                  email_normalized, 
                  password_hash, 
                  nick_name, 
                  first_name, 
                  last_name,
                  wallet_address,
                  last_login, 
                  status, 
                  created_at) 
                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

		_, err = db.ExecContext(ctx, query, user.ID, user.Email, emailNormalized, user.PasswordHash,
			user.NickName, user.FirstName, user.LastName, user.Wallet, user.LastLogin, user.Status, user.CreatedAt)
		if err != nil {
			return ErrUsers.Wrap(err)
		}
	}

	return nil
}

// CreateAdmin inserts admin to DB.
func CreateAdmin(ctx context.Context, conn *sql.DB) error {
	testAdmin := admins.Admin{
		ID:           uuid.New(),
		Email:        "test@test.com",
		PasswordHash: []byte("Qwerty123-"),
		CreatedAt:    time.Now().UTC(),
	}
	err := testAdmin.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}
	_, err = conn.ExecContext(ctx,
		`INSERT INTO admins(id,email,password_hash,created_at)
                VALUES($1,$2,$3,$4)`, testAdmin.ID, testAdmin.Email, testAdmin.PasswordHash, testAdmin.CreatedAt)

	return ErrAdmins.Wrap(err)
}

// CreateDivisions creates a divisions and writes to the database.
func CreateDivisions(ctx context.Context, conn *sql.DB) error {
	var allDivisions []divisions.Division

	for i := 1; i < 11; i++ {
		division := divisions.Division{
			ID:             uuid.New(),
			Name:           i,
			PassingPercent: 10,
			CreatedAt:      time.Now().UTC(),
		}

		allDivisions = append(allDivisions, division)
	}

	query := `INSERT INTO divisions(id, name, passing_percent, created_at) 
	VALUES ($1, $2, $3, $4)`

	for _, division := range allDivisions {
		_, err := conn.ExecContext(ctx, query, division.ID, division.Name, division.PassingPercent, division.CreatedAt)
		if err != nil {
			return ErrDivisions.Wrap(err)
		}
	}

	return nil
}

// CreateClubs creates clubs for each users and writes to the database.
func CreateClubs(ctx context.Context, conn *sql.DB) error {
	allUsers, err := ListUsers(ctx, conn)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	lastDivision, err := GetLastDivision(ctx, conn)
	if err != nil {
		return ErrDivisions.Wrap(err)
	}

	for _, user := range allUsers {
		club := clubs.Club{
			ID:         uuid.New(),
			OwnerID:    user.ID,
			Name:       user.NickName,
			Status:     clubs.StatusActive,
			DivisionID: lastDivision.ID,
			CreatedAt:  time.Now(),
		}

		_, err = conn.ExecContext(ctx, "INSERT INTO clubs(id, owner_id, club_name, status, division_id, created_at)VALUES($1,$2,$3,$4,$5,$6)",
			club.ID, club.OwnerID, club.Name, club.Status, club.DivisionID, club.CreatedAt)

		if err != nil {
			return ErrClubs.Wrap(err)
		}
	}

	return nil
}

// CreateSquads creates squads for all clubs from the database.
func CreateSquads(ctx context.Context, conn *sql.DB) error {
	allClubs, err := ListClubs(ctx, conn)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	for _, club := range allClubs {
		squad := clubs.Squad{
			ID:        uuid.New(),
			Name:      "",
			ClubID:    club.ID,
			Formation: clubs.FourFourTwo,
			Tactic:    clubs.Balanced,
			CaptainID: uuid.Nil,
		}

		_, err = conn.ExecContext(ctx, "INSERT INTO squads(id, squad_name, club_id, formation, tactic, captain_id)VALUES($1,$2,$3,$4,$5,$6)",
			squad.ID, squad.Name, squad.ClubID, squad.Formation, squad.Tactic, squad.CaptainID)

		if err != nil {
			return ErrClubs.Wrap(err)
		}
	}

	return nil
}

// CreateSquadCards creates and inserts squad cards to the database.
func (seedDB *SeedDB) CreateSquadCards(ctx context.Context, conn *sql.DB, cardsConfig cards.Config, lootboxesConfig lootboxes.Config) error {
	cardsService := cards.NewService(seedDB.cards, cardsConfig)

	allClubs, err := ListClubs(ctx, conn)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	var squadCards []clubs.SquadCard

	for _, club := range allClubs {
		squad, err := ListSquadByClubID(ctx, conn, club.ID)
		if err != nil {
			return ErrClubs.Wrap(err)
		}
		for i := 0; i < 11; i++ {
			probabilities := []int{lootboxesConfig.RegularBoxConfig.Wood, lootboxesConfig.RegularBoxConfig.Silver, lootboxesConfig.RegularBoxConfig.Gold, lootboxesConfig.RegularBoxConfig.Diamond}
			card, err := cardsService.Create(ctx, club.OwnerID, probabilities, cards.TypeWon)
			if err != nil {
				return ErrClubs.Wrap(err)
			}

			squadCard := clubs.SquadCard{
				SquadID:  squad.ID,
				CardID:   card.ID,
				Position: clubs.Position(i),
			}

			squadCards = append(squadCards, squadCard)
		}
	}

	for _, card := range squadCards {
		query := `INSERT INTO squad_cards(id, card_id, card_position)
		          VALUES($1,$2,$3)`

		_, err := conn.ExecContext(ctx, query, card.SquadID, card.CardID, card.Position)
		if err != nil {
			return ErrClubs.Wrap(err)
		}
	}

	return nil
}

// ListUsers returns all users from the database.
func ListUsers(ctx context.Context, conn *sql.DB) ([]users.User, error) {
	rows, err := conn.QueryContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, last_login, status, created_at FROM users")
	if err != nil {
		return nil, ErrUsers.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.LastLogin, &user.Status, &user.CreatedAt)
		if err != nil {
			return nil, ErrUsers.Wrap(err)
		}

		data = append(data, user)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrUsers.Wrap(err)
	}

	return data, ErrUsers.Wrap(err)
}

// GetLastDivision returns last division from the database.
func GetLastDivision(ctx context.Context, conn *sql.DB) (divisions.Division, error) {
	query := `SELECT * FROM divisions WHERE name=(SELECT MAX(name) FROM divisions)`
	var division divisions.Division

	row := conn.QueryRowContext(ctx, query)

	err := row.Scan(&division.ID, &division.Name, &division.PassingPercent, &division.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return division, divisions.ErrNoDivision.Wrap(err)
		}

		return division, ErrDivisions.Wrap(err)
	}

	return division, ErrDivisions.Wrap(err)
}

// ListClubs returns all clubs from the database.
func ListClubs(ctx context.Context, conn *sql.DB) ([]clubs.Club, error) {
	query := `SELECT id, owner_id, club_name, status, division_id, created_at
			  FROM clubs`

	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allClubs []clubs.Club

	for rows.Next() {
		var club clubs.Club
		err = rows.Scan(&club.ID, &club.OwnerID, &club.Name, &club.Status, &club.DivisionID, &club.CreatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return allClubs, clubs.ErrNoClub.Wrap(err)
			}
			return allClubs, clubs.ErrClubs.Wrap(err)
		}

		allClubs = append(allClubs, club)
	}

	return allClubs, nil
}

// ListSquadByClubID returns squad by club id from the database.
func ListSquadByClubID(ctx context.Context, conn *sql.DB, clubID uuid.UUID) (clubs.Squad, error) {
	query := `SELECT id, squad_name, club_id, tactic, formation, captain_id
			  FROM squads
			  WHERE club_id = $1`

	row := conn.QueryRowContext(ctx, query, clubID)

	var squad clubs.Squad

	err := row.Scan(&squad.ID, &squad.Name, &squad.ClubID, &squad.Tactic, &squad.Formation, &squad.CaptainID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return squad, clubs.ErrNoSquad.Wrap(err)
		}

		return squad, ErrClubs.Wrap(err)
	}

	return squad, nil
}

// CreateMatches creates matches in the database.
func (seedDB *SeedDB) CreateMatches(ctx context.Context, conn *sql.DB, matchesConfig matches.Config, cardsConfig cards.Config) error {
	usersService := users.NewService(seedDB.users)
	cardsService := cards.NewService(seedDB.cards, cardsConfig)
	clubsService := clubs.NewService(seedDB.clubs, usersService, cardsService, seedDB.divisions)
	matchesService := matches.NewService(seedDB.matches, matchesConfig, clubsService, cardsService)

	type player struct {
		userID   uuid.UUID
		squadID  uuid.UUID
		seasonID int
	}

	var players []player

	allClubs, err := ListClubs(ctx, conn)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	for _, club := range allClubs {
		squad, err := ListSquadByClubID(ctx, conn, club.ID)
		if err != nil {
			return ErrClubs.Wrap(err)
		}
		season, err := GetSeasonByDivisionID(ctx, club.DivisionID, conn)
		if err != nil {
			return Error.Wrap(err)
		}
		player := player{
			userID:   club.OwnerID,
			squadID:  squad.ID,
			seasonID: season.ID,
		}

		players = append(players, player)
	}

	index := 1

	for _, player1 := range players {
		for _, player2 := range players[:len(players)-index] {
			_, err := matchesService.Create(ctx, player1.squadID, player2.squadID, player1.userID, player2.userID, player1.seasonID)
			if err != nil {
				return Error.Wrap(err)
			}
		}
		index++
	}

	return nil
}

// GetSeasonByDivisionID returns season by division id from the data base.
func GetSeasonByDivisionID(ctx context.Context, divisionID uuid.UUID, conn *sql.DB) (seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons WHERE division_id=$1 AND ended_at=$2`
	var season seasons.Season

	row := conn.QueryRowContext(ctx, query, divisionID, time.Time{})

	err := row.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return season, seasons.ErrNoSeason.Wrap(err)
		}

		return season, ErrSeasons.Wrap(err)
	}

	return season, ErrSeasons.Wrap(err)
}
