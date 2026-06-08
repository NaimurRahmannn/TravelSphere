<section class="page-head">
	<h1>Travel Wishlist</h1>
	<p>Edit notes, update trip status, or remove destinations. Changes save
	   without reloading the page.</p>
</section>

{{if .LoadError}}
<p class="alert">{{.LoadError}}</p>
{{end}}

<table class="wishlist-table">
	<thead>
		<tr>
			<th>Country</th>
			<th>Note</th>
			<th>Status</th>
			<th>Actions</th>
		</tr>
	</thead>
	<!-- AJAX target: rows are rebuilt here after add/edit/delete. The server
	     fills it on first load so the table works without JavaScript. -->
	<tbody id="wishlist-rows">
		{{range .Items}}
		<tr data-id="{{.ID}}">
			<td>{{.CountryName}}</td>
			<td>
				<input type="text" class="row-note" value="{{.Note}}" placeholder="Add a note...">
			</td>
			<td>
				<select class="row-status">
					<option value="Planned" {{if eq .Status "Planned"}}selected{{end}}>Planned</option>
					<option value="Visited" {{if eq .Status "Visited"}}selected{{end}}>Visited</option>
				</select>
			</td>
			<td class="row-actions">
				<button class="btn-save">Save</button>
				<button class="btn-delete">Delete</button>
			</td>
		</tr>
		{{else}}
		<tr class="empty-row">
			<td colspan="4">Your wishlist is empty. Add countries from their detail pages.</td>
		</tr>
		{{end}}
	</tbody>
</table>

<!-- A page-level area for messages not tied to a single row (e.g. save success). -->
<div id="wishlist-message"></div>
<script src="/static/js/wishlist.js"></script>