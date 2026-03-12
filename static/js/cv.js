// cv.js — handles CV download via hCaptcha overlay

document.addEventListener('DOMContentLoaded', function () {
  const cvButton = document.getElementById('cv-button');
  const overlay = document.getElementById('hcaptcha-overlay');
  const closeButton = document.getElementById('close-overlay');

  if (!cvButton || !overlay || !closeButton) return;

  // Open overlay and store selected language
  cvButton.addEventListener('click', function () {
    window._cvLang = this.dataset.lang || 'en';
    overlay.style.display = 'block';
  });

  // Close overlay
  closeButton.addEventListener('click', function () {
    overlay.style.display = 'none';
  });
});

// Called by hCaptcha widget on successful verification
function onCaptchaSuccess(token) {
  const lang = window._cvLang || 'en';
  const cvPath = `/static/content/${lang}-cv.pdf`;

  const a = document.createElement('a');
  a.href = cvPath;
  a.target = '_blank';
  a.rel = 'noopener noreferrer';
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);

  document.getElementById('hcaptcha-overlay').style.display = 'none';
}