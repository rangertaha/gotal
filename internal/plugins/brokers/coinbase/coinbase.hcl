broker "coinbase" {
	api_key = "YOUR_API_KEY"
	base_url = "https://api.coinbase.com"
	version = "v2"
	account {
		id = "YOUR_ACCOUNT_ID"
		name = "YOUR_ACCOUNT_NAME"
		currency = "YOUR_ACCOUNT_CURRENCY"
		balance = "YOUR_ACCOUNT_BALANCE"
		allowances {
			buy = "0" // max amount of money that can be bought	
			sell = "0" // max amount of money that can be sold
			transfer = "0" // max amount of money that can be transferred
			deposit = "0" // max amount of money that can be deposited
			withdraw = "0" // max amount of money that can be withdrawn
		}
		limits {
			buy = "0" // max amount of money that can be bought	
			sell = "0" // max amount of money that can be sold
			transfer = "0" // max amount of money that can be transferred
			deposit = "0" // max amount of money that can be deposited
			withdraw = "0" // max amount of money that can be withdrawn
		}
		fees {
			buy = "0" // fee for buying
			sell = "0" // fee for selling	
			transfer = "0" // fee for transferring
			deposit = "0" // fee for depositing
			withdraw = "0" // fee for withdrawing
		}
	}
}
