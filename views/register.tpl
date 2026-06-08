<section class="auth-page">
	<h1>Create an account</h1>
	<p class="muted">Register to start saving destinations to your own wishlist.</p>

	{{if .Error}}
	<p class="alert">{{.Error}}</p>
	{{end}}

	<!-- Normal form POST — a full navigation, not AJAX. -->
	<form method="post" action="/register" class="auth-form">
		<div class="field">
			<label for="username">Username</label>
			<input type="text" id="username" name="username" value="{{.FormUsername}}" required>
		</div>
		<div class="field">
			<label for="password">Password</label>
			<input type="password" id="password" name="password" minlength="6" required>
		</div>
		<div class="field">
			<label for="confirm">Confirm password</label>
			<input type="password" id="confirm" name="confirm" minlength="6" required>
		</div>
		<button type="submit" class="auth-submit">Register</button>
	</form>

	<p class="muted">Already have an account? <a href="/login">Log in</a>.</p>
</section>
