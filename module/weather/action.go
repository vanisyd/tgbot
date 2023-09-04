package weather

import "fmt"

type ActionGeolocation struct {
	Query string
	Limit int
}

func (act ActionGeolocation) GetQuery() string {
	limit := act.Limit
	if limit == 0 {
		limit = 5
	}

	return fmt.Sprintf("%s?q=%s&limit=%d", GeoUrl, act.Query, limit)
}
