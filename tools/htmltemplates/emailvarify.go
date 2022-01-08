package htmltemplates

func EmailVarify(user string, logo string, link string) string {
	return `
	<!DOCTYPE html>
	<html lang="en">
	
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			.rell-image {
				margin: 0;
				padding: 0;
				width: 200px;
			}
	
			.rell-title,
			.rell-user,
			.rell-info {
				text-align: center;
			}
	
			.rell-info-b {
				text-align: center;
			}
	
			.rell-btn {
				background: #6772e5;
				padding: 5px;
				color: #fff;
				text-decoration: none;
				border-radius: 5px;
				transition: .2s;
			}
	
			.rell-btn:hover {
				background: #777fdb;
			}
		</style>
	</head>
	
	<body>
		<div class="mail-b">
	
	
			<h1 class="rell-title">
				<img class="rell-image" src="` + logo + `" alt="">
				
				<br>
				Verify Your Email Address
			</h1>
	
	
			<h3 class="rell-user">
				Hello, ` + user + `!
			</h3>
			<p class="rell-info">
				Please Comform that You Want To Use This As Your Rellitel.ink Account emaill adsress. Once it's done You
				will be
				able to Teck Payment!
				<br>
				<br>
				<a class="rell-btn" target="_blank" href="` + link + `">
					Verify My Email
				</a>
			</p>
	
	
			<p class="rell-info">
				Or Paste this link your browser:
				<br>
				<a href="` + link + `"
					target="_blank">` + link + `</a>
				<br><br>
				&copy;Rellitel.ink 2022
			</p>
		</div>
	
	</body>
	
	</html>
	`
}
