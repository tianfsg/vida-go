let offset = 0;
const loader = document.getElementById("loader");

// Función para cargar más fotos
function loadPhotos() {
  loader.style.display = "block";
  fetch(`/photos?offset=${offset}`)
    .then((response) => response.json())
    .then((data) => {
      // Insertar las nuevas fotos
      const photoContainer = document.getElementById("photo-container");
      data.photos.forEach((photo) => {
        const photoElement = document.createElement("div");
        photoElement.classList.add("photo-item");
        if (photo.IsWide) {
          photoElement.classList.add("horizontal");
        } else {
          photoElement.classList.add("vertical");
        }
        photoElement.innerHTML = `<img src="${photo.URL}" alt="Photo" loading="lazy" />`;
        photoContainer.appendChild(photoElement);
      });

      // Actualizar el offset
      offset = data.offset;
      loader.style.display = "none";
    })
    .catch((error) => {
      console.error("Error loading photos:", error);
      loader.style.display = "none";
    });
}

// Detectar cuando el usuario llega al final de la página
window.addEventListener("scroll", () => {
  // Si el usuario está cerca del final de la página, cargar más fotos
  if (window.innerHeight + window.scrollY >= document.body.offsetHeight - 200) {
    // 200px antes del final
    loadPhotos();
  }
});

// Cargar las primeras fotos al iniciar
loadPhotos();
