package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/DarylSerrano/mimicle-dl/models"
)

const authenticationAPIKey string = "AIzaSyDYw7-i4nM50cvaVgnYO0kft9mNhRpJ5GA"

func GetAlbums(token string) models.GetAlbumsResponse {
	slog.Info("Get all albums purchased")
	var getAlbums models.GetAlbumsResponse
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://mimicle.com/api/my/purchased/albums", nil)
	if err != nil {
		slog.Error("Error creating request for getting albums", slog.Any("error", err))
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("take", "30")
	q.Add("skip", "0")

	req.URL.RawQuery = q.Encode()
	slog.Debug(req.URL.String())

	req.Header.Add("authority", "mimicle.com")
	req.Header.Add("path", "/api/my/purchased/albums")
	req.Header.Add("scheme", "https")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "en-US,en;q=0.5")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("dnt", "1")
	req.Header.Add("referer", "https://mimicle.com/my/purchased")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-GPC", "1")

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error doing request for Get Albums", slog.Any("error", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&getAlbums)
	if err != nil {
		slog.Error("Error json decode for Albums", slog.Any("error", err))
		os.Exit(1)
	}

	return getAlbums
}

func Authenticate(email string, password string) models.AuthResponseBody {
	slog.Info("Authenticate user")
	var authresp models.AuthResponseBody

	authReqBody := models.AuthRequestBody{
		ReturnSecureToken: true,
		Email:             email,
		Password:          password,
		ClientType:        "CLIENT_TYPE_WEB",
	}

	marshalled, err := json.Marshal(authReqBody)
	if err != nil {
		slog.Error("cannot marshall json authReqBody", slog.Any("error", err))
		os.Exit(1)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword", bytes.NewReader(marshalled))
	if err != nil {
		slog.Error("Error creating request for authenticating user", slog.Any("error", err))
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("key", authenticationAPIKey)

	req.URL.RawQuery = q.Encode()
	slog.Info(req.URL.String())

	req.Header.Add("accept", "*/*")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error doing request for authenticating", slog.Any("error", err))
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		slog.Error("Error doing request for authenticating", "statusCode", resp.StatusCode)
		os.Exit(1)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&authresp)
	if err != nil {
		slog.Error("Error json decode for response on authentication", slog.Any("error", err))
		os.Exit(1)
	}

	return authresp
}
