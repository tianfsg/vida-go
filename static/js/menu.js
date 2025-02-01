// Check for <dialog> element compatibility
if (!window.HTMLDialogElement) {
  var contactModal = document.getElementById("contactModal");
  if (contactModal) {
    contactModal.style.display = "none"; // Hide the modal
    alert(
      "El navegador no soporta la etiqueta <dialog>. Considera actualizar tu navegador.",
    );
  }
}

// Toggle the language menu
var languageButton = document.getElementById("languageButton");
if (languageButton) {
  languageButton.addEventListener("click", function () {
    var languageMenu = document.getElementById("languageMenu");
    if (languageMenu) {
      languageMenu.classList.toggle("hidden");
      languageButton.classList.toggle("open"); // Add 'open' class when the menu is visible
    }
  });
}

// Close the language menu when clicking outside
window.addEventListener("click", function (e) {
  var languageMenu = document.getElementById("languageMenu");
  var languageButton = document.getElementById("languageButton");

  if (
    languageMenu &&
    languageButton &&
    !languageButton.contains(e.target) &&
    !languageMenu.contains(e.target)
  ) {
    languageMenu.classList.add("hidden");
    languageButton.classList.remove("open"); // Remove the 'open' class when hidden
  }
});

// Toggle the mobile menu
var menuButton = document.getElementById("menuButton");
var menu = document.getElementById("menu");

if (menuButton && menu) {
  menuButton.addEventListener("click", function () {
    menu.classList.toggle("hidden"); // Toggle visibility
  });
}

// Close both menus when clicking outside
window.addEventListener("click", function (e) {
  var languageMenu = document.getElementById("languageMenu");
  var menu = document.getElementById("menu");
  var languageButton = document.getElementById("languageButton");
  var menuButton = document.getElementById("menuButton");

  // Close the menus if clicked outside
  if (
    languageMenu &&
    menu &&
    !languageButton.contains(e.target) &&
    !menuButton.contains(e.target)
  ) {
    languageMenu.classList.add("hidden");
    menu.classList.add("hidden");
  }
});
