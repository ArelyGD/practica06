package config

import "os"

var USERNAME = os.Getenv("USERNAME")
var PASSWORD = os.Getenv("PASSWORD")
