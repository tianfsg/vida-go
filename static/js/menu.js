// menu.js — all menu logic handled by lang-switch.js via event delegation.
// This file is kept for compatibility but intentionally left empty.

// Dialog compatibility check only
if (!window.HTMLDialogElement) {
  var contactModal = document.getElementById("contactModal");
  if (contactModal) {
    contactModal.style.display = "none";
    alert("El navegador no soporta la etiqueta <dialog>. Considera actualizar tu navegador.");
  }
}