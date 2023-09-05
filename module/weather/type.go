package weather

type Action interface {
	GetQuery() string
}

type Geolocation struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float32           `json:"lat"`
	Lon        float32           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state"`
}

type Weather struct {
	Lat            float32        `json:"lat"`
	Lon            float32        `json:"lon"`
	Timezone       string         `json:"timezone"`
	TimezoneOffset int            `json:"timezone_offset"`
	Current        CurrentWeather `json:"current"`
	//Minutely       MinutelyWeather `json:"minutely"`
	Hourly []HourlyWeather `json:"hourly"`
	Daily  []DailyWeather  `json:"daily"`
}

type CurrentWeather struct {
	Dt         int     `json:"dt"`
	Sunrise    int     `json:"sunrise"`
	Sunset     int     `json:"sunset"`
	Temp       float32 `json:"temp"`
	FeelsLike  float32 `json:"feels_like"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
	DewPoint   float32 `json:"dew_point"`
	Uvi        float32 `json:"uvi"`
	Clouds     int     `json:"clouds"`
	Visibility int     `json:"visibility"`
	WindSpeed  float32 `json:"wind_speed"`
	WindDeg    int     `json:"wind_deg"`
	WindGust   float32 `json:"wind_gust"`
	Weather    []Info  `json:"weather"`
}

// Info Additional weather data
type Info struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MinutelyWeather struct {
	Dt            int `json:"dt"`
	Precipitation int `json:"precipitation"`
}

type HourlyWeather struct {
	Dt         int     `json:"dt"`
	Temp       float32 `json:"temp"`
	FeelsLike  float32 `json:"feels_like"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
	DewPoint   float32 `json:"dew_point"`
	Uvi        float32 `json:"uvi"`
	Clouds     int     `json:"clouds"`
	Visibility int     `json:"visibility"`
	WindSpeed  float32 `json:"wind_speed"`
	WindDeg    int     `json:"wind_deg"`
	WindGust   float32 `json:"wind_gust"`
	Weather    []Info  `json:"weather"`
	Pop        float32 `json:"pop"`
}

type DailyWeather struct {
	Dt        int                  `json:"dt"`
	Sunrise   int                  `json:"sunrise"`
	Sunset    int                  `json:"sunset"`
	Moonrise  int                  `json:"moonrise"`
	Moonset   int                  `json:"moonset"`
	MoonPhase float32              `json:"moon_phase"`
	Summary   string               `json:"summary"`
	Temp      Temperature          `json:"temp"`
	FeelsLike FeelsLikeTemperature `json:"feels_like"`
	Pressure  int                  `json:"pressure"`
	Humidity  int                  `json:"humidity"`
	DewPoint  float32              `json:"dew_point"`
	WindSpeed float32              `json:"wind_speed"`
	WindDeg   int                  `json:"wind_deg"`
	WindGust  float32              `json:"wind_gust"`
	Weather   []Info               `json:"weather"`
	Clouds    int                  `json:"clouds"`
	Pop       float32              `json:"pop"`
	Rain      float32              `json:"rain"`
	Uvi       float32              `json:"uvi"`
}

type Temperature struct {
	Day   float32 `json:"day"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
	Night float32 `json:"night"`
	Eve   float32 `json:"eve"`
	Morn  float32 `json:"morn"`
}

type FeelsLikeTemperature struct {
	Day   float32 `json:"day"`
	Night float32 `json:"night"`
	Eve   float32 `json:"eve"`
	Morn  float32 `json:"morn"`
}
