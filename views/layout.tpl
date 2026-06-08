<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.Title}} | TravelSphere</title>
	<link rel="stylesheet" href="/static/css/style.css?v=3">
</head>
<body>
	{{template "partials/header.tpl" .}}

	<main class="container">
		{{.LayoutContent}}
	</main>

	{{template "partials/footer.tpl" .}}
</body>
</html>