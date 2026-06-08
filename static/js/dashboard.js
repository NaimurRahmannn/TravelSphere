// Refreshes the dashboard stat counters from the JSON API, updating only the
// values inside #dashboard-stats — the rest of the page stays put.
(function () {
  const stats = document.getElementById("dashboard-stats");
  if (!stats) return;

  async function refreshStats() {
    try {
      const resp = await fetch("/api/dashboard/summary");
      if (resp.status === 401) {
        // Session expired while the page was open — send them back to login.
        window.location.href = "/login";
        return;
      }
      if (!resp.ok) throw new Error("request failed");
      const summary = await resp.json();

      // Update each counter in place by its data-stat key.
      stats.querySelector('[data-stat="total"]').textContent = summary.total;
      stats.querySelector('[data-stat="planned"]').textContent = summary.planned;
      stats.querySelector('[data-stat="visited"]').textContent = summary.visited;
    } catch (err) {
    }
  }

  // Refresh once on load to exercise the AJAX path. Exposed globally so other
  window.refreshDashboardStats = refreshStats;
  refreshStats();
})();