// Search autocomplete on the home page
(function () {
  const input = document.getElementById("destination-search");
  const box = document.getElementById("search-suggestions");

  if (!input || !box) return;

  let timer;
  let currentRequestId = 0;

  function resetSearch() {
    input.value = "";
    box.innerHTML = "";
  }

  // Clear search on normal page load
  resetSearch();

  // Clear search when returning with browser back button
  window.addEventListener("pageshow", function () {
    resetSearch();
  });

  function clearSuggestions() {
    box.innerHTML = "";
  }

  function renderSuggestions(countries) {
    if (!countries || countries.length === 0) {
      clearSuggestions();
      return;
    }

    const top = countries.slice(0, 8);

    box.innerHTML = top
      .map(
        (c) =>
          `<a class="suggestion" href="/countries/${c.Slug}">${c.Name} — ${c.Capital}</a>`
      )
      .join("");
  }

  async function fetchSuggestions() {
    const q = input.value.trim();

    if (!q) {
      clearSuggestions();
      return;
    }

    const requestId = ++currentRequestId;

    try {
      const resp = await fetch(`/api/countries?search=${encodeURIComponent(q)}`);

      if (!resp.ok) {
        throw new Error("request failed");
      }

      const countries = await resp.json();

      if (requestId !== currentRequestId || input.value.trim() !== q) {
        return;
      }

      renderSuggestions(countries);
    } catch (err) {
      if (requestId === currentRequestId) {
        clearSuggestions();
      }
    }
  }

  input.addEventListener("input", function () {
    clearTimeout(timer);

    const q = input.value.trim();

    if (!q) {
      currentRequestId++;
      clearSuggestions();
      return;
    }

    timer = setTimeout(fetchSuggestions, 300);
  });
})();