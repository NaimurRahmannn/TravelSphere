
(function () {
  const stats = document.getElementById("dashboard-stats");
  if (!stats) return;

  async function refreshStats() {
    try {
      const resp = await fetch("/api/dashboard/summary");
      if (resp.status === 401) {
      // session expired redirect in login page
        window.location.href = "/login"; 
        return;
      }
      if (!resp.ok) throw new Error("request failed");
      const summary = await resp.json();
      stats.querySelector('[data-stat="total"]').textContent = summary.total;
      stats.querySelector('[data-stat="planned"]').textContent = summary.planned;
      stats.querySelector('[data-stat="visited"]').textContent = summary.visited;
    } catch (err) {
    }
  }

  //Keep the server-rendered values if refresh fails.
  window.refreshDashboardStats = refreshStats;
  refreshStats();
})();