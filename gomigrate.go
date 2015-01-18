package main

import (
	"fmt"
	"github.com/softinio/gomigrate/conf"
)

func main() {
	// Test
	fmt.Printf("Using DB: %s\n", conf.DbxConfig.DbConnectionString())
}
