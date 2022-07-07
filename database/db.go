// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // using postgres driver.
	"github.com/zeebo/errs"

	"ultimatedivision"
	"ultimatedivision/admin/admins"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/cards/waitlist"
	"ultimatedivision/clubs"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/gameplay/queue"
	"ultimatedivision/marketplace"
	"ultimatedivision/seasons"
	"ultimatedivision/store"
	"ultimatedivision/store/lootboxes"
	"ultimatedivision/udts"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

// ensures that database implements ultimatedivision.DB.
var _ ultimatedivision.DB = (*database)(nil)

var (
	// Error is the default ultimatedivision error class.
	Error = errs.Class("ultimatedivision db error")
)

// database combines access to different database tables with a record
// of the db driver, db implementation, and db source URL.
//
// architecture: Master Database
type database struct {
	conn *sql.DB
}

// New returns ultimatedivision.DB postgresql implementation.
func New(databaseURL string) (ultimatedivision.DB, error) {
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	return &database{conn: conn}, nil
}

// Hub entity describes hub of queue for clients.
type Hub struct {
	Queue []queue.Client
}

// NewHub is a constructor for hub entity.
func NewHub() *Hub {
	return &Hub{
		Queue: []queue.Client{},
	}
}

// CreateSchema create schema for all tables and databases.
func (db *database) CreateSchema(ctx context.Context) (err error) {
	createTableQuery :=
		`CREATE TABLE IF NOT EXISTS users (
            id                    BYTEA     PRIMARY KEY    NOT NULL,
            email                 VARCHAR                  NOT NULL,
            email_normalized      VARCHAR                  NOT NULL,
            password_hash         BYTEA                    NOT NULL,
            nick_name             VARCHAR                  NOT NULL,
            first_name            VARCHAR                  NOT NULL,
            last_name             VARCHAR                  NOT NULL,
            wallet_address        BYTEA,
            casper_wallet_address VARCHAR,
            wallet_type           VARCHAR,
            nonce                 BYTEA,
            public_key            VARCHAR,
            private_key           VARCHAR,
            last_login            TIMESTAMP WITH TIME ZONE NOT NULL,
            status                INTEGER                  NOT NULL,
            created_at            TIMESTAMP WITH TIME ZONE NOT NULL
        );
        CREATE TABLE IF NOT EXISTS cards (
            id                BYTEA         PRIMARY KEY NOT NULL,
            player_name       VARCHAR                   NOT NULL,
            quality           VARCHAR                   NOT NULL,  
            height            NUMERIC(16,2)             NOT NULL,
            weight            NUMERIC(16,2)             NOT NULL,
            dominant_foot     VARCHAR                   NOT NULL,
            is_tattoo         BOOLEAN                   NOT NULL,
            status            INTEGER                   NOT NULL,
            type              VARCHAR                   NOT NULL,
            user_id           BYTEA                     NOT NULL,
            tactics           INTEGER                   NOT NULL,
            positioning       INTEGER                   NOT NULL,
            composure         INTEGER                   NOT NULL,
            aggression        INTEGER                   NOT NULL,
            vision            INTEGER                   NOT NULL,
            awareness         INTEGER                   NOT NULL,
            crosses           INTEGER                   NOT NULL,
            physique          INTEGER                   NOT NULL,
            acceleration      INTEGER                   NOT NULL,
            running_speed     INTEGER                   NOT NULL,
            reaction_speed    INTEGER                   NOT NULL,
            agility           INTEGER                   NOT NULL,
            stamina           INTEGER                   NOT NULL,
            strength          INTEGER                   NOT NULL,
            jumping           INTEGER                   NOT NULL,
            balance           INTEGER                   NOT NULL,
            technique         INTEGER                   NOT NULL,
            dribbling         INTEGER                   NOT NULL,
            ball_control      INTEGER                   NOT NULL,
            weak_foot         INTEGER                   NOT NULL,
            skill_moves       INTEGER                   NOT NULL,
            finesse           INTEGER                   NOT NULL,
            curve             INTEGER                   NOT NULL,
            volleys           INTEGER                   NOT NULL,
            short_passing     INTEGER                   NOT NULL,
            long_passing      INTEGER                   NOT NULL,
            forward_pass      INTEGER                   NOT NULL,
            offense           INTEGER                   NOT NULL,
            finishing_ability INTEGER                   NOT NULL,
            shot_power        INTEGER                   NOT NULL,
            accuracy          INTEGER                   NOT NULL,
            distance          INTEGER                   NOT NULL,
            penalty           INTEGER                   NOT NULL,
            free_kicks        INTEGER                   NOT NULL,
            corners           INTEGER                   NOT NULL,
            heading_accuracy  INTEGER                   NOT NULL,
            defence           INTEGER                   NOT NULL,
            offside_trap      INTEGER                   NOT NULL,
            sliding           INTEGER                   NOT NULL,
            tackles           INTEGER                   NOT NULL,
            ball_focus        INTEGER                   NOT NULL,
            interceptions     INTEGER                   NOT NULL,
            vigilance         INTEGER                   NOT NULL,
            goalkeeping       INTEGER                   NOT NULL,
            reflexes          INTEGER                   NOT NULL,
            diving            INTEGER                   NOT NULL,
            handling          INTEGER                   NOT NULL,
            sweeping          INTEGER                   NOT NULL,
            throwing          INTEGER                   NOT NULL
        );
        CREATE TABLE IF NOT EXISTS avatars (
            card_id          BYTEA   PRIMARY KEY REFERENCES cards(id) ON DELETE CASCADE NOT NULL,
            picture_type     INTEGER                                                    NOT NULL,
            face_color       INTEGER                                                    NOT NULL,
            face_type        INTEGER                                                    NOT NULL,
            eyebrows_type    INTEGER                                                    NOT NULL,
            eyebrows_color   INTEGER                                                    NOT NULL,
            eyelaser_type    INTEGER                                                    NOT NULL,
            hairstyle_color  INTEGER                                                    NOT NULL,
            hairstyle_type   INTEGER                                                    NOT NULL,
            nose             INTEGER                                                    NOT NULL,
            tshirt           INTEGER                                                    NOT NULL,
            beard            INTEGER                                                    NOT NULL,
            lips             INTEGER                                                    NOT NULL,
            tattoo           INTEGER                                                    NOT NULL,
            original_url     VARCHAR                                                    NOT NULL,
            preview_url      VARCHAR                                                    NOT NULL
        );
        CREATE TABLE IF NOT EXISTS admins (
            id            BYTEA     PRIMARY KEY    NOT NULL,
            email         VARCHAR                  NOT NULL,
            password_hash BYTEA                    NOT NULL,
            created_at    TIMESTAMP WITH TIME ZONE NOT NULL
        );
        CREATE TABLE IF NOT EXISTS divisions (
            id              BYTEA   PRIMARY KEY      NOT NULL,
            name            INTEGER UNIQUE           NOT NULL,
            passing_percent INTEGER                  NOT NULL,
            created_at      TIMESTAMP WITH TIME ZONE NOT NULL
        );
        CREATE TABLE IF NOT EXISTS clubs (
            id          BYTEA     PRIMARY KEY                                NOT NULL,
            owner_id    BYTEA     REFERENCES users(id) ON DELETE CASCADE     NOT NULL,
            club_name   VARCHAR                                              NOT NULL,
            status      INTEGER                                              NOT NULL,
            division_id BYTEA     REFERENCES divisions(id) ON DELETE CASCADE NOT NULL,
            created_at  TIMESTAMP WITH TIME ZONE                             NOT NULL
        );
        CREATE TABLE IF NOT EXISTS squads (
            id            BYTEA   PRIMARY KEY                            NOT NULL,
            squad_name    VARCHAR,
            club_id       BYTEA   REFERENCES clubs(id) ON DELETE CASCADE NOT NULL,
            tactic        INTEGER                                        NOT NULL,
            formation     INTEGER                                        NOT NULL,
            captain_id    BYTEA
        );
        CREATE TABLE IF NOT EXISTS squad_cards (
            id            BYTEA   REFERENCES squads(id) ON DELETE CASCADE NOT NULL,
            card_id       BYTEA   REFERENCES cards(id) ON DELETE CASCADE  NOT NULL, 
            card_position INTEGER                                         NOT NULL,
            PRIMARY KEY(id, card_id)
        );
        CREATE TABLE IF NOT EXISTS lootboxes(
            user_id      BYTEA   REFERENCES users(id) ON DELETE CASCADE NOT NULL,
            lootbox_id   BYTEA                                          NOT NULL,
            lootbox_type VARCHAR                                        NOT NULL,
            PRIMARY KEY(user_id, lootbox_id)
        );
        CREATE TABLE IF NOT EXISTS lots (
            id            BYTEA                    PRIMARY KEY                            NOT NULL,
            item_id       BYTEA                                                           NOT NULL,
            type          VARCHAR                                                         NOT NULL,
            user_id       BYTEA                    REFERENCES users(id) ON DELETE CASCADE NOT NULL,
            shopper_id    BYTEA,
            status        VARCHAR                                                         NOT NULL,
            start_price   BYTEA                                                           NOT NULL,
            max_price     BYTEA,
            current_price BYTEA,
            start_time    TIMESTAMP WITH TIME ZONE                                        NOT NULL,
            end_time      TIMESTAMP WITH TIME ZONE                                        NOT NULL,
            period        INTEGER                                                         NOT NULL
        );
        CREATE TABLE IF NOT EXISTS seasons(
            id          SERIAL PRIMARY KEY       NOT NULL,
            division_id BYTEA                    NOT NULL,
            started_at  TIMESTAMP WITH TIME ZONE NOT NULL,
            ended_at    TIMESTAMP WITH TIME ZONE NOT NULL,
            FOREIGN KEY (division_id) REFERENCES divisions (id) ON DELETE CASCADE
        );
        CREATE TABLE IF NOT EXISTS matches (
            id           BYTEA   PRIMARY KEY                              NOT NULL,
            user1_id     BYTEA   REFERENCES users(id) ON DELETE CASCADE   NOT NULL,
            squad1_id    BYTEA   REFERENCES squads(id) ON DELETE CASCADE  NOT NULL,
            user1_points INTEGER                                          NOT NULL,
            user2_id     BYTEA   REFERENCES users(id) ON DELETE CASCADE   NOT NULL,
            squad2_id    BYTEA   REFERENCES squads(id) ON DELETE CASCADE  NOT NULL,
            user2_points INTEGER                                          NOT NULL,
            season_id    INTEGER REFERENCES seasons(id) ON DELETE CASCADE NOT NULL
        );
        CREATE TABLE IF NOT EXISTS match_results(
            id       BYTEA   PRIMARY KEY                              NOT NULL,
            match_id BYTEA   REFERENCES matches(id) ON DELETE CASCADE NOT NULL,
            user_id  BYTEA   REFERENCES users(id) ON DELETE CASCADE   NOT NULL,
            card_id  BYTEA   REFERENCES cards(id) ON DELETE CASCADE   NOT NULL,
            minute   INTEGER                                          NOT NULL
        );
        CREATE TABLE IF NOT EXISTS waitlist(
            token_id       SERIAL                                                     NOT NULL,
            card_id        BYTEA   PRIMARY KEY REFERENCES cards(id) ON DELETE CASCADE NOT NULL,
            wallet_address BYTEA                                                      NOT NULL,
            value          BYTEA                                                      NOT NULL,
            password       VARCHAR                                                    NOT NULL,
            wallet_type    VARCHAR                                                    NOT NULL
        );
        CREATE TABLE IF NOT EXISTS nfts(
            card_id        BYTEA   PRIMARY KEY REFERENCES cards(id) NOT NULL,
            token_id       INTEGER                                  NOT NULL,
            chain          VARCHAR                                  NOT NULL,
            wallet_address BYTEA                                    NOT NULL
        );
        CREATE TABLE IF NOT EXISTS currency_waitlist(
            wallet_address BYTEA   NOT NULL,
            value          BYTEA   NOT NULL,
            nonce          INTEGER NOT NULL,
            signature      VARCHAR NOT NULL,
            PRIMARY KEY(wallet_address, nonce)
        );
        CREATE TABLE IF NOT EXISTS udts(
            user_id BYTEA   PRIMARY KEY REFERENCES users(id) NOT NULL,
            nonce   INTEGER                                  NOT NULL
        );
        CREATE TABLE IF NOT EXISTS store(
            id           INTEGER PRIMARY KEY NOT NULL,
            cards_amount INTEGER             NOT NULL,
            is_renewal   BOOLEAN             NOT NULL,
            hour_renewal INTEGER             NOT NULL,
            price        BYTEA               NOT NULL
        );`

	_, err = db.conn.ExecContext(ctx, createTableQuery)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// Close closes underlying db connection.
func (db *database) Close() error {
	return Error.Wrap(db.conn.Close())
}

// Admins provides access to accounts db.
func (db *database) Admins() admins.DB {
	return &adminsDB{conn: db.conn}
}

// Users provides access to accounts db.
func (db *database) Users() users.DB {
	return &usersDB{conn: db.conn}
}

// Cards provides access to accounts db.
func (db *database) Cards() cards.DB {
	return &cardsDB{conn: db.conn}
}

// Avatars provides access to accounts db.
func (db *database) Avatars() avatars.DB {
	return &avatarsDB{conn: db.conn}
}

// Clubs provides access to accounts db.
func (db *database) Clubs() clubs.DB {
	return &clubsDB{conn: db.conn}
}

// LootBoxes provides access to accounts db.
func (db *database) LootBoxes() lootboxes.DB {
	return &lootboxesDB{conn: db.conn}
}

// Marketplace provides access to accounts db.
func (db *database) Marketplace() marketplace.DB {
	return &marketplaceDB{conn: db.conn}
}

// Matches provides access to accounts db.
func (db *database) Matches() matches.DB {
	return &matchesDB{conn: db.conn}
}

// Queue provides access to accounts db.
func (db *database) Queue() queue.DB {
	return &queueHub{hub: NewHub()}
}

// Divisions provides access to accounts db.
func (db *database) Divisions() divisions.DB {
	return &divisionsDB{conn: db.conn}
}

// Seasons provides access to accounts db.
func (db *database) Seasons() seasons.DB {
	return &seasonsDB{conn: db.conn}
}

// WaitList provides access to accounts db.
func (db *database) WaitList() waitlist.DB {
	return &waitlistDB{conn: db.conn}
}

// NFTs provides access to accounts db.
func (db *database) NFTs() nfts.DB {
	return &nftsDB{conn: db.conn}
}

// CurrencyWaitList provides access to accounts db.
func (db *database) CurrencyWaitList() currencywaitlist.DB {
	return &currencywaitlistDB{conn: db.conn}
}

// UDTs provides access to accounts db.
func (db *database) UDTs() udts.DB {
	return &udtsDB{conn: db.conn}
}

// Store provides access to accounts db.
func (db *database) Store() store.DB {
	return &storeDB{conn: db.conn}
}
