package main

import "testing"

func checkAll(t testing.TB, access Access, want float64) {
	t.Helper()
	got := access.balance

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAccess(t *testing.T) {
	t.Run("Get balance", func(t *testing.T) {
		access := Access{balance: 100.00}
		checkAll(t, access, 100.00)
	})

	t.Run("Deposit", func(t *testing.T) {
		access := Access{120.5}
		access.Deposit(50)
		checkAll(t, access, 170.50)
	})

	t.Run("Withdraw", func(t *testing.T) {
		access := Access{balance: 50}
		access.Withdraw(20)
		checkAll(t, access, 30)
	})
}

func TestGetName(t *testing.T) {
	got := NewAccess().GetName()
	want := "Access"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
