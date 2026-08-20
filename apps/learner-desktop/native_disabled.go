//go:build !gtk4

package main

import "fmt"

func runNative(_ desktopConfig) int {
	fmt.Println("Learner Studio requires the Ubuntu GTK4 build.")
	fmt.Println("Build with: go build -tags gtk4 ./apps/learner-desktop")
	fmt.Println("The release package includes the GTK4 native build automatically.")
	return 1
}
