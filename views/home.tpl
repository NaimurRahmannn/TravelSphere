<section class="hero">
	<h1>Discover your next destination</h1>
	<p>Search countries, explore attractions, and curate your personal travel wishlist.</p>
	<div class="search-box">
		<label for="destination-search">Where to next?</label>
		<input type="text" id="destination-search" placeholder="Start typing a country..." autocomplete="off">
		<!-- Autocomplete suggestions get injected here by the search script.
		     Empty on first load; populated only as the user types. -->
		<div id="search-suggestions"></div>
	</div>
</section>

<section class="featured">
	<h2>Featured destinations</h2>
	<div class="country-grid">
		{{range .Featured}}
		<a href="/countries/{{.Slug}}" class="country-card">
			<img src="{{.FlagPNG}}" alt="{{.FlagAlt}}" class="flag">
			<div class="card-body">
				<h3>{{.Name}}</h3>
				<p class="meta">{{.Capital}} · {{.Region}}</p>
			</div>
		</a>
		{{end}}
	</div>
</section>

<section class="attractions">
	<h2>Popular attractions</h2>
	<ul class="attraction-list">
		{{range .PopularAttractions}}
		<li>
			<span class="attraction-name">{{.Name}}</span>
			<span class="attraction-kinds">{{.Kinds}}</span>
		</li>
		{{end}}
	</ul>
</section>