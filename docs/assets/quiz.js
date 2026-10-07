/* pika course — quiz component.
 * Usage: <div class="quiz" data-quiz='[{"q":"...","opts":["a","b","c"],"a":1}, ...]'></div>
 * Renders one question at a time, immediate feedback, score at end.
 */
(function () {
  document.querySelectorAll(".quiz").forEach(function (root) {
    var questions = JSON.parse(root.dataset.quiz);
    var idx = 0, score = 0;
    function render() {
      if (idx >= questions.length) {
        var msg = score === questions.length ? "Perfect — that's storage strength."
          : score >= questions.length / 2 ? "Solid. Re-read the misses, then try again from memory."
          : "Rough pass — that's fine. Skim the section above and retry; effortful retrieval is the point.";
        root.innerHTML = '<p class="quiz-score">Score: ' + score + ' / ' + questions.length + ' — ' + msg + '</p>' +
          '<button class="quiz-restart">Retry</button>';
        root.querySelector(".quiz-restart").onclick = function () { idx = 0; score = 0; render(); };
        return;
      }
      var q = questions[idx];
      var html = '<p class="quiz-q"><strong>Q' + (idx + 1) + '.</strong> ' + q.q + '</p><div class="quiz-opts">' +
        q.opts.map(function (o, i) { return '<button data-i="' + i + '">' + o + '</button>'; }).join("") +
        '</div><p class="quiz-fb"></p><p class="quiz-progress">' + (idx + 1) + ' / ' + questions.length + '</p>';
      root.innerHTML = html;
      root.querySelectorAll(".quiz-opts button").forEach(function (b) {
        b.onclick = function () {
          var chosen = +b.dataset.i, fb = root.querySelector(".quiz-fb");
          root.querySelectorAll(".quiz-opts button").forEach(function (x) { x.disabled = true; });
          if (chosen === q.a) {
            b.classList.add("correct"); score++;
            fb.textContent = q.why || "Correct."; fb.style.color = "var(--ok)";
          } else {
            b.classList.add("wrong");
            root.querySelector('[data-i="' + q.a + '"]').classList.add("correct");
            fb.textContent = (q.why || "") + " Correct answer is highlighted."; fb.style.color = "var(--bad)";
          }
          setTimeout(function () { idx++; render(); }, chosen === q.a ? 650 : 2200);
        };
      });
    }
    render();
  });
})();
