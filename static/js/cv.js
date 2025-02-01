let selectedLang = ''; // Variable global para almacenar el idioma seleccionado

// Cuando el captcha es validado correctamente
function onCaptchaSuccess(token) {
    // Utilizar la variable global selectedLang para obtener el idioma
    const cvPath = `/static/content/${selectedLang}-cv.pdf`;

    // Crear un enlace temporal
    const a = document.createElement('a');
    a.href = cvPath;
    a.target = '_blank'; // Abrir en una nueva pestaña
    a.rel = 'noopener noreferrer'; // Seguridad adicional

    // Agregar el enlace al DOM, hacer clic en él, y luego eliminarlo
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);

    // Ocultar el overlay después de la verificación
    document.getElementById('hcaptcha-overlay').style.display = 'none';
}

// Muestra el captcha flotante cuando se hace clic en el botón de CV
document.getElementById('cv-button').addEventListener('click', function() {
    // Establecer el idioma seleccionado en la variable global
    selectedLang = this.dataset.lang;

    // Mostrar el overlay y el captcha
    document.getElementById('hcaptcha-overlay').style.display = 'block';
});

document.addEventListener('DOMContentLoaded', function () {
    // Botón "X" que cierra el overlay
    const closeButton = document.getElementById('close-overlay');
    const hcaptchaOverlay = document.getElementById('hcaptcha-overlay');

    closeButton.addEventListener('click', function () {
        hcaptchaOverlay.style.display = 'none';
    });

    // Abrir el overlay al hacer clic en el botón de CV
    const cvButton = document.getElementById('cv-button');
    cvButton.addEventListener('click', function () {
        hcaptchaOverlay.style.display = 'block';
    });
});
