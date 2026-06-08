<section class="page-head">
	<h1>Travel Dashboard</h1>
	<p>Your saved trips at a glance. Stats refresh automatically when your
	   wishlist changes.</p>
</section>

{{if .LoadError}}
<p class="alert">{{.LoadError}}</p>
{{end}}

<!-- AJAX target: these three counters are re-fetched from
     /api/dashboard/summary after wishlist changes. Server-rendered on load. -->
<div id="dashboard-stats" class="stats-grid">
	<div class="stat-card">
		<span class="stat-label">Total Saved</span>
		<span class="stat-value" data-stat="total">{{.Summary.Total}}</span>
	</div>
	<div class="stat-card">
		<span class="stat-label">Planned</span>
		<span class="stat-value" data-stat="planned">{{.Summary.Planned}}</span>
	</div>
	<div class="stat-card">
		<span class="stat-label">Visited</span>
		<span class="stat-value" data-stat="visited">{{.Summary.Visited}}</span>
	</div>
</div>

<section class="saved-list">
	<h2>Saved destinations</h2>
	<ul class="destination-list">
		{{range .Items}}
		<li>
			<span class="dest-name">{{.CountryName}}</span>
			<span class="dest-status">{{.Status}}</span>
			{{if .Note}}<span class="dest-note">· {{.Note}}</span>{{end}}
		</li>
		{{else}}
		<li class="muted">No saved destinations yet.</li>
		{{end}}
	</ul>
</section>

<script src="/static/js/dashboard.js"></script>