package main

func loadAll() []*Account {
	return []*Account{
		&Account{balance: balance(1000)},
		&Account{balance: balance(0)},
		&Account{balance: balance(0)},
		&Account{balance: balance(200)},
		&Account{balance: balance(0)},
	}
}

func loadIncomplete() []*Account {
	return []*Account{
		&Account{balance: balance(1000)},
		nil,
		&Account{balance: nil},
		&Account{balance: balance(200)},
		nil,
  }
}
