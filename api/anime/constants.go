package anime

const (
	queryGraphqlSearch = `query {
		Page {
			media(search: "%s", type: %s) {
				id
				title {
					romaji
					english
					native
				}
				coverImage {
					large
				}
				bannerImage
				description
				averageScore
				popularity
				episodes
				season
				seasonYear
				isAdult
				format
				genres
			}
		}
	}`
	queryGraphqlFullInfo = `query {
	Page {
		media(id: %d, type: %s) {
		id
		title {
			romaji
			english
			native
		}
		staff {
			edges {
			node {
				name {
				first
				last
				}
			}
			role
			id
			}
		}
		characters {
			edges {
			node {
				name {
				native
				full
				}
				gender
				description
				image {
				medium
				}
			}
			role
			id
			}
		}
		coverImage {
			large
		}
		duration
		bannerImage
		description
		averageScore
		popularity
		episodes
		season
		seasonYear
		isAdult
		format
		genres
		}
	}
	}

`
)
