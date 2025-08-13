package auth

import (
	"encoding/base64"
	"fmt"
	"misw/db"
	"misw/model"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
)

func ValidateRequestingUser(r *http.Request, conn *pgx.Conn) (*model.User, error) {
	header := r.Header.Get("Authorization")
	if header == "" || len(header) < 6 || header[:6] != "Basic " {
		return nil, fmt.Errorf("Unauthorized")
	}

	payload, err := base64.StdEncoding.DecodeString(header[6:])
	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return nil, fmt.Errorf("Unauthorized")
	}

	username := pair[0]
	password := pair[1]

	user, err := db.GetUser(conn, username, password)
	if(err != nil){
		return nil, err;
	}

	return user, nil;
}