document.addEventListener("DOMContentLoaded", function () {
  console.log("calendarButton.js loaded"); // Confirma que el archivo se carga correctamente

  var openButton = document.getElementById("openCalendarButton");

  if (openButton) {
    openButton.addEventListener("click", function (event) {
      event.preventDefault();

      // Obtener la URL desde el atributo data-calendar-url
      var calendarUrl = openButton.dataset.calendarUrl;

      if (calendarUrl) {
        // Abrir una ventana emergente con las dimensiones especificadas
        window.open(calendarUrl, "popup", "width=600,height=600");
      } else {
        console.error("Calendar URL not found in data attribute");
      }
    });
  } else {
    console.error("openCalendarButton not found");
  }
});
