(function () {
  const btn = document.getElementById("add-wishlist-btn");
  const feedback = document.getElementById("wishlist-feedback");

  if (!btn || !feedback) return;

  btn.addEventListener("click", async function () {
    const countryName = btn.dataset.country;

    btn.disabled = true;
    feedback.className = "";
    feedback.textContent = "Adding...";

    try {
      const resp = await fetch("/api/wishlist", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          country_name: countryName,
          note: "",
          status: "Planned", //user can change it on wishlist page
        }),
      });

      if (resp.status === 401) {
        feedback.textContent = "Please log in to save to your wishlist...";
        window.location.href = "/login";
        return;
      }

      const data = await resp.json();

      if (!resp.ok) {
        feedback.className = "error";
        feedback.textContent = data.error || "Could not add to wishlist.";
        btn.disabled = false;
        return;
      }

      feedback.className = "success";
      feedback.textContent = `${countryName} added to your wishlist.`;
    } catch (err) {
      feedback.className = "error";
      feedback.textContent = "Network error. Please try again.";
      btn.disabled = false;
    }
  });
})();