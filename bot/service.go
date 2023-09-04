package bot

var Commands []CMD = []CMD{
	CMD{
		Command:    CMDRegister,
		Parameters: []string{"password"},
		Handler:    Register,
	},
	CMD{
		Command: CMDNow,
		Handler: Now,
	},
	CMD{
		Command: CMDWeather,
		Handler: Weather,
	},
}
