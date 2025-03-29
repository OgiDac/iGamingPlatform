package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func JSON(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.Encode(obj)
}

func MigrateDB(db *sqlx.DB) {
	// TODO add migration script
	_, err := db.Exec(CreatePlayers)
	if err != nil {
		log.Fatal(err)
	}

	var countPlayers int
	var countTournaments int
	err = db.Get(&countPlayers, "SELECT COUNT(*) FROM players")
	if err != nil {
		log.Fatal(err)
	}

	if countPlayers == 0 {
		_, err = db.Exec(InsertPlayers)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = db.Exec(CreateTournaments)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Get(&countTournaments, "SELECT COUNT(*) FROM tournaments")
	if err != nil {
		log.Fatal(err)
	}
	if countTournaments == 0 {
		_, err = db.Exec(InsertTournaments)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = db.Exec(`DROP PROCEDURE IF EXISTS DistributePrizes`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(CreatePrizeDistributionSP)
	if err != nil {
		log.Fatal(err)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("test1234"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`UPDATE players SET players.password = ?`, encryptedPassword)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(CreatePlayerTournaments)
	if err != nil {
		log.Fatal(err)
	}
}

func SetCookie(w http.ResponseWriter, name string, value string) {
	cookie := http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
}
