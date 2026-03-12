(function () {

  // ── Toggle dropdown on button click ───────────────────────────────────────
  document.addEventListener("click", function (e) {
    var btn = e.target.closest("#languageButton");
    if (!btn) return;
    var menu = document.getElementById("languageMenu");
    if (!menu) return;
    var isHidden = menu.classList.contains("hidden");
    // Close all first, then open if it was closed
    menu.classList.add("hidden");
    btn.classList.remove("open");
    if (isHidden) {
      menu.classList.remove("hidden");
      btn.classList.add("open");
    }
  });


  // ── Close when clicking outside ────────────────────────────────────────────
  // Uses mousedown (not click) so it doesn't race with the toggle above
  document.addEventListener("mousedown", function (e) {
    var langBtn  = document.getElementById("languageButton");
    var langMenu = document.getElementById("languageMenu");
    if (!langMenu || !langBtn) return;
    if (!langBtn.contains(e.target) && !langMenu.contains(e.target)) {
      langMenu.classList.add("hidden");
      langBtn.classList.remove("open");
    }

    var navBtn  = document.getElementById("menuButton");
    var navMenu = document.getElementById("menu");
    if (!navMenu || !navBtn) return;
    if (!navBtn.contains(e.target) && !navMenu.contains(e.target)) {
      navMenu.classList.add("hidden");
    }
  });


  // ── Mobile menu toggle ─────────────────────────────────────────────────────
  document.addEventListener("click", function (e) {
    if (!e.target.closest("#menuButton")) return;
    var menu = document.getElementById("menu");
    if (menu) menu.classList.toggle("hidden");
  });


  // ── Lang option click — swap DOM without reload ────────────────────────────
  document.addEventListener("click", function (e) {
    var link = e.target.closest("a.lang-option[data-lang]");
    if (!link) return;

    e.preventDefault();
    e.stopImmediatePropagation();

    var targetLang = link.getAttribute("data-lang");
    var url = new URL(window.location.href);
    url.searchParams.set("lang", targetLang);

    var menu = document.getElementById("languageMenu");
    var btn  = document.getElementById("languageButton");
    if (menu) menu.classList.add("hidden");
    if (btn)  btn.classList.remove("open");

    fetch(url.toString())
      .then(function (r) { return r.text(); })
      .then(function (html) {
        var parser = new DOMParser();
        var newDoc = parser.parseFromString(html, "text/html");
        [["main","main"], ["header","header"], ["footer","footer"]].forEach(function (p) {
          var o = document.querySelector(p[0]);
          var n = newDoc.querySelector(p[1]);
          if (o && n) o.replaceWith(n);
        });
        window.history.pushState({}, "", url.toString());
        if (typeof initCV       === "function") initCV();
        if (typeof initCalendar === "function") initCalendar();
      })
      .catch(function () {
        window.location.href = url.toString();
      });
  }, true); // capture phase

})();