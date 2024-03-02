package html

func WelcomePageHTML() string {
    return `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>CoinSnark</title>
			</head>
			<body>
				<h1>CoinSnark Go API</h1>
				<ul>
					<li><a href="https://rmottanet.github.io/coinsnark">Frontend</a></li>
					<li><a href="https://rmottanet.gitbook.io/coinsnark">Documentation</a></li>
				</ul>
			</body>
		</html>
    `
}
