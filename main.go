package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/DarylSerrano/mimicle-dl/api"
	"github.com/DarylSerrano/mimicle-dl/downloader"
	"github.com/DarylSerrano/mimicle-dl/models"
	"github.com/DarylSerrano/mimicle-dl/utils"
	"github.com/manifoldco/promptui"
)

func SelectAlbumDownload(albums []models.AlbumInfo) models.AlbumInfo {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   ">{{ .Title | cyan }} {{ .ID }}",
		Inactive: "  {{ .Title | cyan }} {{ .ID }}",
		Selected: "{{ .Title | red | cyan }} {{ .ID }}",
		Details: `
---------  Mimicle Album ----------
 {{ .Title }}
 {{ .ShortDescription }}
 https://mimicle.com/album/{{ .Nid }}`,
	}

	searcher := func(input string, index int) bool {
		album := albums[index]
		id := strings.Replace(strings.ToLower(album.ID), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(id, input)
	}

	prompt := promptui.Select{
		Label:     "Select Album to Download",
		Items:     albums,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		slog.Error("Prompt failed", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Printf("You choose number %d: %s\n", i+1, albums[i].Title)

	return albums[i]
}

func DownloadAllOrSelect(token string, tmpFolder string, outputFolder string, downloadAll bool) {
	purchasedAlbums := api.GetAlbums(token)
	if purchasedAlbums.Total > 0 {
		if downloadAll {
			var wg sync.WaitGroup
			fmt.Printf("Download all albums purchased: %d\n", purchasedAlbums.Total)

			for _, album := range purchasedAlbums.Items {
				childFolder := utils.RandomString(5)
				wg.Add(1)
				randomTmpFolder := path.Join(tmpFolder, childFolder)

				go func(album models.AlbumInfo, token string, tmpFolder string) {
					defer wg.Done()
					o := filepath.Join(outputFolder, album.Title)
					downloader.DownloadAllTracks(album, token, tmpFolder, o)
				}(album, token, randomTmpFolder)

				time.Sleep(3 * time.Second)
			}

			wg.Wait()
			fmt.Println("Finished downloading all albums")
		} else {
			albumselected := SelectAlbumDownload(purchasedAlbums.Items)
			o := filepath.Join(outputFolder, albumselected.Title)
			downloader.DownloadAllTracks(albumselected, token, tmpFolder, o)
		}
	} else {
		fmt.Println("No albums purchased")
		os.Exit(1)
	}
}

func main() {

	var token = flag.String("token", "", "Bearer Token of the user")
	var tmpFolder = flag.String("tmpFolder", "./tmp", "path of tmp folder [default ./tmp]")
	var outputFolder = flag.String("outputFolder", "./downloads", "path where to save downloaded files [default ./downloads]")
	var logLevel = flag.String("logLevel", "INFO", "Log leves, can be: [INFO, DEBUG] by default its INFO")
	var downloadAll = flag.Bool("all", false, "pass --all true to download all the albums")
	var email = flag.String("email", "", "Email for authentication")
	var password = flag.String("password", "", "Password for authenticaiton")

	flag.Parse()

	opts := &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: utils.RemoveTimeSlog,
	}

	if len(*logLevel) > 0 && *logLevel == "DEBUG" {
		opts.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
	slog.SetDefault(logger)

	if len(*token) > 0 {
		fmt.Println("Token supplied")
		DownloadAllOrSelect(*token, *tmpFolder, *outputFolder, *downloadAll)
	} else {
		fmt.Println("No token supplied")
		if len(*email) > 0 && len(*password) > 0 {
			fmt.Println("Going to use authentication with email and password")
			resp := api.Authenticate(*email, *password)
			DownloadAllOrSelect(resp.IDToken, *tmpFolder, *outputFolder, *downloadAll)
		} else {
			fmt.Println("Email or password not supplied")
			os.Exit(1)
		}
	}
}
