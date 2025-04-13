package downloader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DarylSerrano/mimicle-dl/models"
	"github.com/DarylSerrano/mimicle-dl/utils"
)

func GenericHttpDownload(url string, filepath string) {
	slog.Debug(fmt.Sprintf("Do a GET https requests and save the information into: %s, %s \n", filepath, url))
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Error creating request generic http download", slog.Any("error", err))
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error doing request generic http download", slog.Any("error", err))
		os.Exit(1)
	}

	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		slog.Error("Error create file", "filepath", filepath, slog.Any("error", err))
		os.Exit(1)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		slog.Error("Error copy http request response to file", slog.Any("error", err))
		os.Exit(1)
	}
}

func DownloadContentURL(contentHost string, contentPath string, streamToken string, outpath string) {
	slog.Debug(fmt.Sprintf("Do a GET https requests and save the information into: %s, %s %s \n", outpath, contentHost, contentPath))

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", contentHost, contentPath), nil)
	if err != nil {
		slog.Error("Error creating request for DownloadContentURL ", slog.Any("error", err))
		os.Exit(1)
	}

	req.Header.Add("origin", "https://mimicle.com")
	req.Header.Add("referer", "https://mimicle.com/")
	// req.Header.Add("path", "/api/my/purchased/albums")
	req.Header.Add("scheme", "https")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "en-US,en;q=0.5")
	// req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("dnt", "1")
	// req.Header.Add("referer", "https://mimicle.com/my/purchased")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-GPC", "1")
	req.Header.Add("x-stream-token", streamToken)

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error doing request DownloadContentURL", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Debug(fmt.Sprintf("Status code for downloading %s%s is: %d \n", contentHost, contentPath, resp.StatusCode))

	// Save the contents to a file
	defer resp.Body.Close()
	out, err := os.Create(outpath)
	if err != nil {
		slog.Error("Error create file", "outpath", outpath, slog.Any("error", err))
		os.Exit(1)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		slog.Error("Error copy http request response DownloadContentURL to file", slog.Any("error", err))
		os.Exit(1)
	}
}

func OptionsContentURL(contentHost string, contentPath string, streamToken string, outpath string) {
	slog.Debug(fmt.Sprintf("Do a Options https requests and save the information into: %s, %s %s \n", outpath, contentHost, contentPath))

	client := &http.Client{}
	req, err := http.NewRequest("OPTIONS", fmt.Sprintf("%s%s", contentHost, contentPath), nil)
	if err != nil {
		slog.Error("Error creating request for OptionsContentURL ", slog.Any("error", err))
		os.Exit(1)
	}

	req.Header.Add("origin", "https://mimicle.com")
	req.Header.Add("referer", "https://mimicle.com/")
	// req.Header.Add("path", "/api/my/purchased/albums")
	req.Header.Add("scheme", "https")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "en-US,en;q=0.5")
	// req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("dnt", "1")
	// req.Header.Add("referer", "https://mimicle.com/my/purchased")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-GPC", "1")
	req.Header.Add("x-stream-token", streamToken)

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error doing request OptionsContentURL", slog.Any("error", err))
		os.Exit(1)
	}

	// Save the contents to a file
	defer resp.Body.Close()
	out, err := os.Create(outpath)
	if err != nil {
		slog.Error("Error create file", "outpath", outpath, slog.Any("error", err))
		os.Exit(1)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		slog.Error("Error copy http request response OptionsContentURL to file", slog.Any("error", err))
		os.Exit(1)
	}

}

func GetTrackDownloadToken(token string, trackid string) models.TrackTokenInfo {
	slog.Debug("Get download track token")
	var tokeninfo models.TrackTokenInfo

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://mimicle.com/api/permission/generate/%s", trackid), nil)
	if err != nil {
		slog.Error("Error creating request for GetTrackDownloadToken ", slog.Any("error", err))
		os.Exit(1)
	}
	req.Header.Add("authority", "mimicle.com")
	// req.Header.Add("path", "/api/my/purchased/albums")
	req.Header.Add("scheme", "https")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "en-US,en;q=0.5")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("dnt", "1")
	// req.Header.Add("referer", "https://mimicle.com/my/purchased")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-GPC", "1")

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error doing request GetTrackDownloadToken", slog.Any("error", err))
		os.Exit(1)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&tokeninfo)
	if err != nil {
		slog.Error("Error doing decode json for GetTrackDownloadToken reponse body", slog.Any("error", err))
		os.Exit(1)
	}

	return tokeninfo
}

func DownloadTrack(trackinfo models.TrackInfo, token string, tmpFolder string, outFolder string) {
	slog.Info(fmt.Sprintf("Download track: %s", trackinfo.Title))

	jobs := make(chan string)
	var wg sync.WaitGroup

	filenamem3u8 := filepath.Base(trackinfo.File.Formats.SV1H.FilePath) //"v1_h.m3u8"
	// fmt.Println(filenamem3u8)
	pathm3u8 := trackinfo.File.Formats.SV1H.FilePath
	// fmt.Println(pathm3u8)
	outputpathm3u8 := filepath.Join(tmpFolder, filenamem3u8) //fmt.Sprintf("./downloads/%s", filenamem3u8)
	// fmt.Println(outputpathm3u8)

	// here we only replace the filename by "stream"
	pathstream := strings.Replace(pathm3u8, filenamem3u8, "stream", 1) //"stream/v1_h/stream"

	tracktokeninfo := GetTrackDownloadToken(token, trackinfo.ID)

	subtitlesFilePath := filepath.Join(outFolder, fmt.Sprintf("%d_%s_subtitles.json", trackinfo.TrackNumber, trackinfo.Title))

	// Download subtitles
	slog.Info("Download subtitles")
	DownloadContentURL(tracktokeninfo.ContentPath, "stream/subtitles.json", tracktokeninfo.Token, subtitlesFilePath)

	// Download m3u8
	// OptionsContentURL(tracktokeninfo.ContentPath, pathm3u8, tracktokeninfo.Token, "./downloads/options.txt")
	DownloadContentURL(tracktokeninfo.ContentPath, pathm3u8, tracktokeninfo.Token, outputpathm3u8)

	// Download Stream key
	OptionsContentURL(tracktokeninfo.ContentPath, pathm3u8, tracktokeninfo.Token, filepath.Join(tmpFolder, "stream")) //"./downloads/stream"
	DownloadContentURL(tracktokeninfo.ContentPath, pathstream, tracktokeninfo.Token, filepath.Join(tmpFolder, "stream"))

	file, err := os.Open(outputpathm3u8)
	if err != nil {
		slog.Error("Error opening file to write m3u8", "outputpathm3u8", outputpathm3u8, slog.Any("error", err))
		os.Exit(1)
	}
	defer file.Close()

	// Parallel reading lines from file
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			slog.Debug("scanner", "line", line)
			if !strings.HasPrefix(line, "#") {
				slog.Debug("send to jobs channel")
				jobs <- line
			}
		}
		close(jobs)
		if err := scanner.Err(); err != nil {
			slog.Error("Error scanner fir file", "outputpathm3u8", outputpathm3u8, slog.Any("error", err))
			os.Exit(1)
		}
	}()

	// Paralel download of files
	for fileToDownload := range jobs {
		wg.Add(1)

		slog.Debug("Range over jobs", "fileToDownload", fileToDownload)

		go func(fileToDownload string) {
			defer wg.Done()
			DownloadContentURL(tracktokeninfo.ContentPath,
				strings.Replace(pathm3u8, filenamem3u8, fileToDownload, 1),
				tracktokeninfo.Token,
				filepath.Join(tmpFolder, fileToDownload)) //fmt.Sprintf("./downloads/%s", line)
		}(fileToDownload)
	}

	// Must wait untill all downloads finish
	wg.Wait()
}

func DownloadAllTracks(album models.AlbumInfo, token string, tmpFolder string, outFolder string) {

	utils.Mkdir(tmpFolder)
	utils.Mkdir(outFolder)

	for i, track := range album.Tracks {

		slog.Info("Going to download track", "track.Title", track.Title)

		trackTmpFolder := filepath.Join(tmpFolder, strconv.Itoa(i))
		utils.Mkdir(trackTmpFolder)
		DownloadTrack(track, token, trackTmpFolder, outFolder)

		// Save into m4a
		outfilename := fmt.Sprintf("%d_%s.m4a", track.TrackNumber, strings.TrimSpace(track.Title))
		absOutFolder, err := filepath.Abs(outFolder)
		if err != nil {
			slog.Error("Error on getting absolute path", slog.Any("error", err))
			os.Exit(1)
		}
		cleanOutputPath := filepath.Join(absOutFolder, outfilename)

		// Setup metadata file and download cover for the track
		metadataTrack := models.FFmpegTrackMetadata{
			Title:      track.Title,
			Artist:     utils.GetArtistsMetdata(album),
			Album:      album.Title,
			Genre:      "asmr",
			Date:       strconv.Itoa(album.PublishAt.Year()),
			TrackCount: fmt.Sprintf("%d/%d", track.TrackNumber, len(album.Tracks)),
		}

		metadataFilepath := CreateMetadataFile(trackTmpFolder, metadataTrack)

		// Download cover
		coverFilepath, err := filepath.Abs(filepath.Join(trackTmpFolder, fmt.Sprintf("%s.webp", album.ID)))
		if err != nil {
			slog.Error("Error on getting absolute path for coverFilePath", slog.Any("error", err))
			os.Exit(1)
		}

		// Download the covert art
		GenericHttpDownload(album.CoverArt.Path, coverFilepath)

		RunFffmpeg(trackTmpFolder,
			filepath.Base(track.File.Formats.SV1H.FilePath),
			cleanOutputPath,
			metadataFilepath,
			coverFilepath)

		// Sleep before doing another calls
		slog.Debug("Sleep 10 seconds before downloading the next track")
		time.Sleep(10 * time.Second)
	}

	// Cleanup tmp folder
	utils.CleanupFolder(tmpFolder)
}
