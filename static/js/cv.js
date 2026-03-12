function onCaptchaSuccess(token) {
  var lang = window._cvLang || "en";
  var cvPath = "/static/content/" + lang + "-cv.pdf";
  var a = document.createElement("a");
  a.href = cvPath;
  a.target = "_blank";
  a.rel = "noopener noreferrer";
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  var overlay = document.getElementById("hcaptcha-overlay");
  if (overlay) overlay.style.display = "none";
}

function initCV() {
  var cvBtn = document.getElementById("cv-button");
  var overlay = document.getElementById("hcaptcha-overlay");
  var closeBtn = document.getElementById("close-overlay");

  if (cvBtn) {
    // Clone to remove old listeners
    var newBtn = cvBtn.cloneNode(true);
    cvBtn.parentNode.replaceChild(newBtn, cvBtn);
    newBtn.addEventListener("click", function () {
      window._cvLang = (this.dataset.lang || "en").toLowerCase();
      if (overlay) overlay.style.display = "block";
    });
  }

  if (closeBtn) {
    var newClose = closeBtn.cloneNode(true);
    closeBtn.parentNode.replaceChild(newClose, closeBtn);
    newClose.addEventListener("click", function () {
      if (overlay) overlay.style.display = "none";
    });
  }
}

if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", initCV);
} else {
  initCV();
}