<header class="site-header">
	<nav class="nav">
		<a href="/" class="logo">TravelSphere</a>
		<ul class="nav-links">
			<li><a href="/">Home</a></li>
			<li><a href="/countries">Countries</a></li>
			<li><a href="/wishlist">Wishlist</a></li>
			<li><a href="/dashboard">Dashboard</a></li>
		</ul>
		<div class="nav-auth">
			{{if .IsLoggedIn}}
				<span class="nav-user">Hi, {{.Username}}</span>
				<a href="/logout">Logout</a>
			{{else}}
				<a href="/login">Login</a>
			{{end}}
		</div>
	</nav>
</header>