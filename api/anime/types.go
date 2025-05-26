package anime

type InfoResponse struct {
	Data struct {
		Page struct {
			Media []struct {
				ID    int `json:"id"`
				Title struct {
					Romaji  string `json:"romaji"`
					English string `json:"english"`
					Native  string `json:"native"`
				} `json:"title"`
				Staff struct {
					Edges []struct {
						Node struct {
							Name struct {
								First string `json:"first"`
								Last  string `json:"last"`
							} `json:"name"`
						} `json:"node"`
						Role string `json:"role"`
						ID   int    `json:"id"`
					} `json:"edges"`
				} `json:"staff"`
				Characters struct {
					Edges []struct {
						Node struct {
							Name struct {
								Native string `json:"native"`
								Full   string `json:"full"`
							} `json:"name"`
						} `json:"node"`
						Role string `json:"role"`
						ID   int    `json:"id"`
					} `json:"edges"`
				} `json:"characters"`
				CoverImage struct {
					Large string `json:"large"`
				} `json:"coverImage"`
				Duration     int      `json:"duration"`
				BannerImage  string   `json:"bannerImage"`
				Description  string   `json:"description"`
				AverageScore int      `json:"averageScore"`
				Popularity   int      `json:"popularity"`
				Episodes     int      `json:"episodes"`
				Season       string   `json:"season"`
				SeasonYear   int      `json:"seasonYear"`
				IsAdult      bool     `json:"isAdult"`
				Format       string   `json:"format"`
				Genres       []string `json:"genres"`
			} `json:"media"`
		} `json:"Page"`
	} `json:"data"`
}

type SearchResult struct {
	Data struct {
		Page struct {
			Media []struct {
				ID    int `json:"id"`
				Title struct {
					Romaji  string `json:"romaji"`
					English string `json:"english"`
					Native  string `json:"native"`
				} `json:"title"`
				CoverImage struct {
					Large string `json:"large"`
				} `json:"coverImage"`
				Duration     int      `json:"duration"`
				BannerImage  string   `json:"bannerImage"`
				Description  string   `json:"description"`
				AverageScore int      `json:"averageScore"`
				Episodes     int      `json:"episodes"`
				Season       string   `json:"season"`
				SeasonYear   int      `json:"seasonYear"`
				IsAdult      bool     `json:"isAdult"`
				Format       string   `json:"format"`
				Genres       []string `json:"genres"`
			} `json:"media"`
		} `json:"Page"`
	} `json:"data"`
}
