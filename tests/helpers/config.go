package helpers

import "os"

func init() {
	os.Setenv("MODE", "test")
	os.Setenv("PORT", "8080")
	os.Setenv("DISCORD_PUBLIC_KEY", "8933e3749b4feb4d76169b26ed372af3c378f4353c2024fee0601f2a2e7918e1")
}
