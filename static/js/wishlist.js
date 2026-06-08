(function () {
  const rows = document.getElementById("wishlist-rows");
  const message = document.getElementById("wishlist-message");

  if (!rows) return;

  // Briefly show a status message, then clear it.
  function flash(text, kind) {
    message.textContent = text;
    message.className = kind || "";
    setTimeout(() => {
      message.textContent = "";
      message.className = "";
    }, 2500);
  }


  rows.addEventListener("click", async function (e) {
    const row = e.target.closest("tr");
    if (!row) return;
    const id = row.dataset.id;

    // Save update note and status
    if (e.target.classList.contains("btn-save")) {
      const note = row.querySelector(".row-note").value;
      const status = row.querySelector(".row-status").value;

      e.target.disabled = true;
      try {
        const resp = await fetch(`/api/wishlist/${id}`, {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ note: note, status: status }),
        });
        const data = await resp.json();
        if (!resp.ok) {
          flash(data.error || "Could not save changes.", "error");
        } else {
          flash("Saved.", "success");
        }
      } catch (err) {
        flash("Network error. Please try again.", "error");
      } finally {
        e.target.disabled = false;
      }
    }

    //  Delete 
    if (e.target.classList.contains("btn-delete")) {
      e.target.disabled = true;
      try {
        const resp = await fetch(`/api/wishlist/${id}`, { method: "DELETE" });
        if (!resp.ok) {
          const data = await resp.json();
          flash(data.error || "Could not delete.", "error");
          e.target.disabled = false;
          return;
        }
        // Remove just this row from the DOM — no reload, no refetch needed.
        row.remove();
        flash("Removed.", "success");

        // If that was the last row, show the empty state.
        if (!rows.querySelector("tr[data-id]")) {
          rows.innerHTML =
            '<tr class="empty-row"><td colspan="4">Your wishlist is empty. Add countries from their detail pages.</td></tr>';
        }
      } catch (err) {
        flash("Network error. Please try again.", "error");
        e.target.disabled = false;
      }
    }
  });
})();