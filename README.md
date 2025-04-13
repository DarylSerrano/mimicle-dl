# mimicle-dl
Downloads an album or all albums from a users purchased albums. It can also download subtitles

## Requirements
This cli needs ffmpeg to be on the `PATH`

## Usage
Download the binaries from releases tab according to your os
```bash
mimicle-dl --help
# Usage of mimicle-dl:
#   -all
#         pass --all true to download all the albums
#   -email string
#         Email for authentication
#   -logLevel string
#         Log leves, can be: [INFO, DEBUG] by default its INFO (default "INFO")
#   -outputFolder string
#         path where to save downloaded files [default ./downloads] (default "./downloads")
#   -password string
#         Password for authenticaiton
#   -tmpFolder string
#         path of tmp folder [default ./tmp] (default "./tmp")
#   -token string
#         Bearer Token of the user

# Will fetch all purchased albums and you can select one to download
mimicle-dl --email "your@email.here" --password "yourpasswordhere"

# You can also use the bearer token obtained from doing a inspect and checking api requests headers in your web browser
# --all will download all the albums the user has purchased
mimicle-dl --token "bearer token from network tab here" --all

# You can specify a different outputFolder and a different tmpFolder
mimicle-dl --token "tokenhere" --outputFolder "some/other/folder" --tmpFolder "some/other/tmp/folder" 
```
