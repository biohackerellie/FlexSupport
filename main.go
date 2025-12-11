package main

import (
	"context"
	"fmt"
	"os"

	// "os/exec"

	"github.com/joho/godotenv"
)

var Environment = "development"

func init() {
	os.Setenv("env", Environment)
	if Environment == "development" {
		// 	// exec.Command("bunx", "tailwindcss", "-i", "./static/assets/css/input.css", "-o", "./static/assets/css/output.min.css", "-m").Run()
		// 	// exec.Command("go", "tool", "templ", "generate")
		// 	out, err := exec.Command("task", "gen").Output()
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	fmt.Printf("%s\n", out)
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	ctx := context.Background()
	if err := App(ctx, os.Stdout, getEnv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
