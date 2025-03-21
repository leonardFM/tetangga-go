package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

// InitSentry menginisialisasi Sentry
func InitSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://e1676bad3f26db6f4aa8547060f9145a@o1337716.ingest.us.sentry.io/4509014921379840",
	})
	if err != nil {
		log.Fatalf("❌ Sentry initialization failed: %v", err)
	}
}

// CaptureError menangkap error dan mengirim ke Sentry
func CaptureError(err error) {
	if err != nil {
		sentry.CaptureException(err)
		sentry.Flush(2 * 1000) // Flush dengan timeout 2 detik
	}
}
