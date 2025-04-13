package utils

import (
	"errors"
	"log/slog"
	"math/rand"
	"os"
	"path"
	"strings"

	"github.com/DarylSerrano/mimicle-dl/models"
)

func GetArtistsMetdata(album models.AlbumInfo) string {
	out := ""
	var actors []string
	for _, tagsCast := range album.CastTags {
		if tagsCast.AttributeType == "voiceActor" {
			actors = append(actors, tagsCast.Tag.CastName)
		}
	}

	out = strings.Join(actors, "/")

	return out

}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func Mkdir(folder string) {

	if !checkFileExists(folder) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			slog.Error("Error when creating a folder", "folder", folder, slog.Any("error", err))
			os.Exit(1)
		}
	}

}

// Cleanup all contents inside of the folder
func CleanupFolder(folder string) {
	dir, err := os.ReadDir(folder)
	if err != nil {
		slog.Error("Error reading tmpFolder", "tmpFolder", folder, slog.Any("error", err))
		os.Exit(1)
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{folder, d.Name()}...))
	}
}

// https://golangdocs.com/generate-random-string-in-golang
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// https://cs.opensource.google/go/go/+/refs/tags/go1.24.2:src/log/slog/internal/slogtest/slogtest.go;l=13
func RemoveTimeSlog(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}
