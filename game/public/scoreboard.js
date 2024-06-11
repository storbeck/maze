function updateScoreboard() {
    const scoreboardBody = document.getElementById('scoreboardBody');
    scoreboardBody.innerHTML = '';
    const scores = JSON.parse(localStorage.getItem('scores')) || {};
    Object.keys(scores).forEach(player => {
      const row = document.createElement('tr');
      const winCount = scores[player]?.winCount || 0;
      const bestTime = scores[player]?.bestTime || 'N/A';
      row.innerHTML = `<td>${player}</td><td>${winCount}</td><td>${bestTime}</td>`;
      scoreboardBody.appendChild(row);
    });
  }
  