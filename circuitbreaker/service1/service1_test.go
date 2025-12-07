package service1

import (
    "testing"

    "example/circuitbreaker/entity"
)

func TestRetryIfFunc(t *testing.T) {
    if retryIfFunc(entity.ERRForever) {
        t.Fatalf("expected retryIfFunc to return false for ERRForever")
    }

    if retryIfFunc(entity.ERRPermission) {
        t.Fatalf("expected retryIfFunc to return false for ERRPermission")
    }

    if !retryIfFunc(entity.ERRUnusual) {
        t.Fatalf("expected retryIfFunc to return true for ERRUnusual")
    }
}
