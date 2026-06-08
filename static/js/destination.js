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
          status: "Planned", // default; user can change it on wishlist page
        }),
      });

      const data = await resp.json();

      if (!resp.ok) {
        // The API's standard error shape gives us .error to show.
        feedback.className = "error";
        feedback.textContent = data.error || "Could not add to wishlist.";
        btn.disabled = false;
        return;
      }

      feedback.className = "success";
      feedback.textContent = `${countryName} added to your wishlist.`;
      // Leave the button disabled on success so it reads as "done".
    } catch (err) {
      feedback.className = "error";
      feedback.textContent = "Network error. Please try again.";
      btn.disabled = false;
    }
  });
})();