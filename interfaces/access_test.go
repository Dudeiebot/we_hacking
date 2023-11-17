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

// Table driven approach because i prefer it also
// it is very easy to understand also

//func checkBalance(t *testing.T, access Access, want float64) {
// 	t.Helper()
// 	got := access.balance
//
// 	if got != want {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }
//
// func TestAccess(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		initial   float64
// 		operation func(*Access)
// 		want      float64
// 	}{
// 		{"Get balance", 100.00, nil, 100.00},
// 		{"Deposit", 120.50, func(a *Access) { a.Deposit(50) }, 170.50},
// 		{"Withdraw", 50.00, func(a *Access) { a.Withdraw(20) }, 30.00},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			access := Access{balance: test.initial}
// 			if test.operation != nil {
// 				test.operation(&access)
// 			}
// 			checkBalance(t, access, test.want)
// 		})
// 	}
// }
//
// func TestGetName(t *testing.T) {
// 	t.Run("Get Name", func(t *testing.T) {
// 		got := NewAccess().GetName()
// 		want := "Access"
//
// 		if got != want {
// 			t.Errorf("got %v want %v", got, want)
// 		}
// 	})
// }
