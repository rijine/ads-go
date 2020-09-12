package config

import "time"

type JwtConfig struct {
	Expiry time.Duration
	Issuer string
	Secret string
}

var JwtConf = JwtConfig{
	Expiry: 24 * time.Hour,
	Issuer: "companyname",
	Secret: "this is my secret key, and it should be changed to 512",
}

/*type config struct {
	JwtConfig
}*/

/*// TODO: READ from config file or env
var Config config = config{JwtConfig{
	Expiry: 24 * time.Hour,
	Issuer: "companyname",
	Secret: "this is my secret key, and it should be changed to 512",
}}*/
