package models

import "time"

type AuthResponseBody struct {
	Kind         string `json:"kind"`
	LocalID      string `json:"localId"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	IDToken      string `json:"idToken"` // This is the token we need for authenticated api calls
	Registered   bool   `json:"registered"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

type AuthRequestBody struct {
	ReturnSecureToken bool   `json:"returnSecureToken"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ClientType        string `json:"clientType"`
}

type TrackTokenInfo struct {
	Type        string `json:"type"`
	Token       string `json:"token"`
	ContentPath string `json:"contentPath"`
}

type FFmpegTrackMetadata struct {
	Title      string
	Artist     string
	Album      string
	Genre      string
	Date       string
	TrackCount string
}

type TrackInfo struct {
	ID          string    `json:"id"`
	VoiceID     any       `json:"voiceId"`
	AlbumID     string    `json:"albumId"`
	Status      string    `json:"status"`
	FileID      string    `json:"fileId"`
	Price       any       `json:"price"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	I18N        any       `json:"i18n"`
	TrackNumber int       `json:"trackNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	File        struct {
		FileType string  `json:"fileType"`
		Duration float64 `json:"duration"`
		Formats  struct {
			RV1 struct {
				Type     string `json:"type"`
				FilePath string `json:"filePath"`
			} `json:"r_v1"`
			SV1H struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				FilePath string `json:"filePath"`
			} `json:"s_v1_h"`
			SV1L struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				FilePath string `json:"filePath"`
			} `json:"s_v1_l"`
			SV1M struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				FilePath string `json:"filePath"`
			} `json:"s_v1_m"`
			SV2H struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				FilePath string `json:"filePath"`
			} `json:"s_v2_h"`
			SV2L struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				FilePath string `json:"filePath"`
			} `json:"s_v2_l"`
			SV2M struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				FilePath string `json:"filePath"`
			} `json:"s_v2_m"`
			TV1H struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				Duration int    `json:"duration"`
				FilePath string `json:"filePath"`
			} `json:"t_v1_h"`
			TV1L struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				Duration int    `json:"duration"`
				FilePath string `json:"filePath"`
			} `json:"t_v1_l"`
			TV1M struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				Duration int    `json:"duration"`
				FilePath string `json:"filePath"`
			} `json:"t_v1_m"`
			TV2H struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				Duration int    `json:"duration"`
				FilePath string `json:"filePath"`
			} `json:"t_v2_h"`
			TV2L struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				Duration int    `json:"duration"`
				FilePath string `json:"filePath"`
			} `json:"t_v2_l"`
			TV2M struct {
				Type     string `json:"type"`
				BitRate  string `json:"bitRate"`
				Duration int    `json:"duration"`
				FilePath string `json:"filePath"`
			} `json:"t_v2_m"`
		} `json:"formats"`
		NoSample bool `json:"noSample"`
	} `json:"file"`
}

type AlbumInfo struct {
	ID               string    `json:"id"`
	Nid              string    `json:"nid"`
	Status           string    `json:"status"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	Price            int       `json:"price"`
	PublishAt        time.Time `json:"publishAt"`
	PreOrderStartAt  any       `json:"preOrderStartAt"`
	ReleaseAt        time.Time `json:"releaseAt"`
	LotteryEndAt     time.Time `json:"lotteryEndAt"`
	CoverArtID       string    `json:"coverArtId"`
	BookletID        string    `json:"bookletId"`
	AvailableCountry []any     `json:"availableCountry"`
	I18N             any       `json:"i18n"`
	TargetGender     string    `json:"targetGender"`
	Metadata         struct {
		Banners []struct {
			URI  string `json:"uri"`
			Link string `json:"link"`
		} `json:"banners"`
		Copyright []struct {
			Name string `json:"name"`
		} `json:"copyright"`
		Transcript bool `json:"transcript"`
	} `json:"metadata"`
	SeriesID       string    `json:"seriesId"`
	IsTradeHalted  bool      `json:"isTradeHalted"`
	IsDownloadable bool      `json:"isDownloadable"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	User           struct {
		ID          string `json:"id"`
		Username    string `json:"username"`
		Nickname    string `json:"nickname"`
		Description string `json:"description"`
		AvatarID    string `json:"avatarId"`
		BannerID    string `json:"bannerId"`
		Avatar      struct {
			Path     string `json:"path"`
			FileType string `json:"fileType"`
			Width    int    `json:"width"`
			Height   int    `json:"height"`
			Dominant string `json:"dominant"`
		} `json:"avatar"`
		Banner struct {
			Path     string `json:"path"`
			FileType string `json:"fileType"`
			Width    int    `json:"width"`
			Height   int    `json:"height"`
			Dominant string `json:"dominant"`
		} `json:"banner"`
	} `json:"user"`
	CoverArt struct {
		Path     string `json:"path"`
		FileType string `json:"fileType"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		Dominant string `json:"dominant"`
	} `json:"coverArt"`
	Booklet struct {
		Path     string `json:"path"`
		FileType string `json:"fileType"`
		FileSize int    `json:"fileSize"`
		Numpages int    `json:"numpages"`
	} `json:"booklet"`
	Tracks []TrackInfo `json:"tracks"`
	Series struct {
		ID         string    `json:"id"`
		SeriesName string    `json:"seriesName"`
		UserID     string    `json:"userId"`
		I18N       any       `json:"i18n"`
		CreatedAt  time.Time `json:"createdAt"`
	} `json:"series"`
	Tags []struct {
		AlbumID   string    `json:"albumId"`
		TagID     string    `json:"tagId"`
		CreatedAt time.Time `json:"createdAt"`
		Tag       struct {
			ID        string    `json:"id"`
			TagName   string    `json:"tagName"`
			Hiragana  any       `json:"hiragana"`
			I18N      any       `json:"i18n"`
			CreatedAt time.Time `json:"createdAt"`
		} `json:"tag"`
	} `json:"tags"`
	CastTags []struct {
		AlbumID       string    `json:"albumId"`
		TagID         string    `json:"tagId"`
		AttributeType string    `json:"attributeType"`
		AttributeName string    `json:"attributeName"`
		SortOrder     int       `json:"sortOrder"`
		CreatedAt     time.Time `json:"createdAt"`
		Tag           struct {
			ID        string    `json:"id"`
			CastName  string    `json:"castName"`
			Hiragana  any       `json:"hiragana"`
			I18N      any       `json:"i18n"`
			CreatedAt time.Time `json:"createdAt"`
		} `json:"tag"`
	} `json:"castTags"`
}

type GetAlbumsResponse struct {
	Skip  int         `json:"skip"`
	Take  int         `json:"take"`
	Total int         `json:"total"`
	Items []AlbumInfo `json:"items"`
}
