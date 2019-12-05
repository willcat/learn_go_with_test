package pointer

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		got := w.Balance()
		if want != got {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error) {
		if err == nil {
			t.Error("wanted an error but didnt got one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		_ = wallet.WithDraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("WithDraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.WithDraw(Bitcoin(100))
		assertBalance(t, wallet, startBalance)
		assertError(t, err)
	})
}
