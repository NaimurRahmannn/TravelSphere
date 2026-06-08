<section class="page-head">
	<h1>Country Explorer</h1>
	<p>Browse every destination on first load. Search and filter update only the
	   results below — no full page reload.</p>
</section>

<div class="filters">
	<div class="field">
		<label for="country-search">Search</label>
		<input type="text" id="country-search" placeholder="Country or capital...">
	</div>
	<div class="field">
		<label for="region-filter">Region</label>
		<select id="region-filter">
			<option value="">All regions</option>
			{{range .Regions}}
			<option value="{{.}}">{{.}}</option>
			{{end}}
		</select>
	</div>
</div>

{{if .LoadError}}
<p class="alert">{{.LoadError}}</p>
{{end}}

<!-- The search/filter script swaps the inner HTML of this container only.
     The server fills it on first load so the grid works without JS. -->
<div id="country-results" class="country-grid">
	{{range .Countries}}
	<a href="/countries/{{.Slug}}" class="country-card">
		<img src="{{.FlagPNG}}" alt="{{.FlagAlt}}" class="flag">
		<div class="card-body">
			<h3>{{.Name}}</h3>
			<p><strong>Capital:</strong> {{.Capital}}</p>
			<p><strong>Population:</strong> {{.Population}}</p>
			<p><strong>Currency:</strong> {{range .Currencies}}{{.}} {{end}}</p>
			<p><strong>Languages:</strong> {{range .Languages}}{{.}} {{end}}</p>
		</div>
	</a>
	{{end}}
</div>

 <script src="/static/js/countries.js"></script>