package config

type AppConfig struct {
	APP_NAME    string `defualt:"GINGO"`
	DESCRIPTION string `defualt:"A lightweight and modular Golang framework built with Gin and GORM for fast and scalable web development."`
	AUTHOR      string `defualt:"Sajad pourajam"`
	DEBUG       string `defualt:"true"`
}
