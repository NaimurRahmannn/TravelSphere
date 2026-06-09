
(function () {
  const searchInput = document.getElementById("country-search");
  const regionSelect = document.getElementById("region-filter");
  const results = document.getElementById("country-results");

  
  if (!searchInput || !regionSelect || !results) return;

 // Match the server-rendered country card layout.
  function cardHTML(country) {
    const currencies = (country.Currencies || []).join(" ");
    const languages = (country.Languages || []).join(" ");
    return `
      <a href="/countries/${country.Slug}" class="country-card">
        <img src="${country.FlagPNG}" alt="${country.FlagAlt}" class="flag">
        <div class="card-body">
          <h3>${country.Name}</h3>
          <p><strong>Capital:</strong> ${country.Capital}</p>
          <p><strong>Population:</strong> ${country.Population}</p>
          <p><strong>Currency:</strong> ${currencies}</p>
          <p><strong>Languages:</strong> ${languages}</p>
        </div>
      </a>`;
  }

  async function runSearch() {
    const params = new URLSearchParams();
    if (searchInput.value.trim()) params.set("search", searchInput.value.trim());
    if (regionSelect.value) params.set("region", regionSelect.value);

  
    results.innerHTML = `<p class="muted">Loading...</p>`;

    try {
      const resp = await fetch(`/api/countries?${params.toString()}`);
      if (!resp.ok) throw new Error("request failed");
      const countries = await resp.json();

      if (!countries || countries.length === 0) {
        results.innerHTML = `<p class="muted">No countries match your search.</p>`;
        return;
      }
      results.innerHTML = countries.map(cardHTML).join("");
    } catch (err) {
      results.innerHTML = `<p class="alert">Could not load results. Please try again.</p>`;
    }
  }

  // // avoid request on every keypress
  let debounceTimer;
  function debouncedSearch() {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(runSearch, 300);
  }

  searchInput.addEventListener("input", debouncedSearch);
  regionSelect.addEventListener("change", runSearch);
})();