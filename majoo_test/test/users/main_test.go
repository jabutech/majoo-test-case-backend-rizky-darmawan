package test

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Load env
	godotenv.Load("../../.env")

	m.Run()

}
