package database

import (
	// "context"
	"testing"
)

// func mustStartPostgresContainer() (func(context.Context) error, error)

// func TestMain(m *testing.M)

func TestNew(t *testing.T) {
	srv := New()
	if srv == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New()

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error to not be present")
	}
}

func TestClose(t *testing.T) {
	srv := New()

	if srv.Close() == nil {
		t.Fatalf("expected Close() to return nil")
	}
}
