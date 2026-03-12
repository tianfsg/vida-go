function initCalendar() {
  var btn = document.getElementById("openCalendarButton");
  if (!btn) return;
  var newBtn = btn.cloneNode(true);
  btn.parentNode.replaceChild(newBtn, btn);
  newBtn.addEventListener("click", function (e) {
    e.preventDefault();
    var url = this.dataset.calendarUrl;
    if (url) window.open(url, "popup", "width=600,height=600");
  });
}

if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", initCalendar);
} else {
  initCalendar();
}