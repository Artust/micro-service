package util

func PageIndex(page int64, perPage int64) (int64, int64) {
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 20
	}
	skip := (page - 1) * perPage
	return perPage, skip
}
