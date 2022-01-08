package htmltemplates

func Forgot(logo string, link string) string {
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
				Reset Password
			</h1>
	
			<p class="rell-info">
				We Receved a request to You Are Forgot password For Rellitel.ink Account 
				<br>
				<br>
				<a class="rell-btn" target="_blank" href="` + link + `">
					Resete Password
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
