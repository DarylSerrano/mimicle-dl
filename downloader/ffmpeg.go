package downloader

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/DarylSerrano/mimicle-dl/models"
)

func CreateMetadataFile(trackWorkingDirectory string, trackmetadata models.FFmpegTrackMetadata) string {
	// Information: on how to construct metadata and use it: https://jmesb.com/how_to/dump_and_load_metadata_with_ffmpeg
	// Header for metadata.
	// ;FFMETADATA1

	metadataPath := filepath.Join(trackWorkingDirectory, "metadata.txt")
	f, err := os.Create(metadataPath)
	if err != nil {
		slog.Error("Error creating metadata file", slog.Any("error", err))
		os.Exit(1)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	content := fmt.Sprintf(";FFMETADATA1\nartist=%s\ngenre=%s\ndate=%s\nalbum=%s\ntitle=%s\ntrack=%s\nalbum_artist=%s\n",
		trackmetadata.Artist,
		trackmetadata.Genre,
		trackmetadata.Date,
		trackmetadata.Album,
		trackmetadata.Title,
		trackmetadata.TrackCount,
		trackmetadata.Artist)

	n4, err := w.WriteString(content)
	if err != nil {
		slog.Error("Error writing into metadata file", slog.Any("error", err))
		os.Exit(1)
	}
	slog.Debug("Bytes writen", "bytes", n4)

	w.Flush()

	r, err := filepath.Abs(metadataPath)

	if err != nil {
		slog.Error("Error resolving absolute path for metadata.txt", slog.Any("error", err))
		os.Exit(1)
	}

	return r
}

func RunFffmpeg(trackWorkingDirectory string, trackm3u8 string, outputTrack string, metadataPath string, coverpath string) {
	var cmd *exec.Cmd
	var args []string

	slog.Debug("Current working directory: ", "trackWorkingDirectory", trackWorkingDirectory)
	slog.Info(fmt.Sprintf("Going to convert track: %s into m4a and save it into: %s", trackm3u8, outputTrack))

	args = append(args, "-y",
		"-allowed_extensions",
		"ALL",
		"-i",
		trackm3u8,
		"-i",
		coverpath,
		"-i",
		metadataPath,
		"-map_metadata",
		"2",
		"-c",
		"copy",
		"-map",
		"0",
		"-map",
		"1",
		"-disposition:v:0",
		"attached_pic",
		outputTrack)

	// -y -allowed_extensions ALL -i v1_h.m3u8 -c copy  test.m4a

	cmd = exec.Command("ffmpeg", args...)
	cmd.Dir = trackWorkingDirectory

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		slog.Error("Error waiting ffmpeg command start execution", slog.Any("error", err))
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		slog.Error("Error ffmpeg command run wait", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("converted m3u8 into m4a", "outputTrack", outputTrack)
}
