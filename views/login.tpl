<section class="auth-page">
	<h1>Login</h1>
	<p class="muted">Log in to access your wishlist and dashboard.</p>

	{{if .Error}}
	<p class="alert">{{.Error}}</p>
	{{end}}

	<!-- Normal form POST — this is a real navigation, not an AJAX action, so a
	     full submit is correct here. -->
	<form method="post" action="/login" class="auth-form">
		<div class="field">
			<label for="username">Username</label>
			<input type="text" id="username" name="username" value="{{.FormUsername}}" required>
		</div>
		<div class="field">
			<label for="password">Password</label>
			<input type="password" id="password" name="password" required>
		</div>
		<button type="submit" class="auth-submit">Log in</button>
	</form>

	<p class="muted">No account yet? <a href="/register">Register</a>.</p>
</section>