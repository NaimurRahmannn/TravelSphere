<section class="detail-header card">
	<img src="{{.Country.FlagPNG}}" alt="{{.Country.FlagAlt}}" class="detail-flag">
	<div class="detail-info">
		<span class="region-badge">{{.Country.Region}}</span>
		<h1>{{.Country.Name}}</h1>
		<p class="official">{{.Country.OfficialName}}</p>

		<div class="fact-row">
			<div><span class="fact-label">Capital</span>{{.Country.Capital}}</div>
			<div><span class="fact-label">Population</span>{{population .Country.Population}}</div>
			<div><span class="fact-label">Region</span>{{.Country.Region}} · {{.Country.Subregion}}</div>
			<div><span class="fact-label">Currency</span>{{range .Country.Currencies}}{{.}} {{end}}</div>
			<div><span class="fact-label">Languages</span>{{range .Country.Languages}}{{.}} {{end}}</div>
		</div>
	</div>
</section>

<section class="wishlist-action card">
	<!-- data-country feeds the wishlist POST; the script reads it off this button. -->
	<button id="add-wishlist-btn" data-country="{{.Country.Name}}">Add to Wishlist</button>
	<!-- Success/error text lands here only — the rest of the page stays put. -->
	<div id="wishlist-feedback"></div>
</section>

<div class="detail-columns">
	<section class="weather card">
		<h2>Travel weather</h2>
		<!-- Weather is the optional bonus. Without a key configured we show a
		     hint instead of leaving an empty box. -->
		<div id="weather-panel">
			{{if .Weather}}
				<p class="temp">{{.Weather.TempC}}°C — {{.Weather.Condition}}</p>
			{{else}}
				<p class="muted">Weather data is optional. Add WEATHERAPI_KEY to enable live conditions.</p>
			{{end}}
		</div>
	</section>

	<section class="attractions card">
		<h2>Attractions &amp; landmarks</h2>
		{{if .AttractionError}}
			<p class="alert">{{.AttractionError}}</p>
		{{end}}
		<ul class="attraction-list">
			{{range .Attractions}}
			<li>
				<span class="attraction-name">{{.Name}}</span>
				<span class="attraction-kinds">{{.Kinds}}</span>
			</li>
			{{else}}
			<li class="muted">No attractions found for this location.</li>
			{{end}}
		</ul>
	</section>
</div>