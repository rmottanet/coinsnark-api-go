package html

func WelcomePageHTML() string {
    return `
		<!DOCTYPE html>
		<html lang="en">
			<head>
			    <meta charset="UTF-8">
			    <meta name="viewport" content="width=device-width, initial-scale=1.0">
			    <title>Coin Snark</title>
			</head>
		    <body>
		        <h1>Welcome to the Coin Snark API</h1>
		        <p>Thank you for visiting CoinSnark! For more information, see the <a href="https://rmottanet.gitbook.io/coinsnark">documentation</a>.</p>
		    </body>
		</html>
    `
}
