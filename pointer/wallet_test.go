package pointer

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, want Bitcoin, got Bitcoin) {
		if want != got {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, want, got)
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.WithDraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, want, got)
	})
}
