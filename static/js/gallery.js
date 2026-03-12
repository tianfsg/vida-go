(function () {
  let offset = 0;
  let loading = false;
  let exhausted = false;

  const loader = document.getElementById("loader");
  const container = document.getElementById("photo-container");

  function appendPhotos(photos) {
    photos.forEach((photo) => {
      const item = document.createElement("div");
      item.className = "photo-item";
      const img = document.createElement("img");
      img.src = photo.URL;
      img.alt = "Photo";
      img.loading = "lazy";
      if (photo.WebP) {
        const picture = document.createElement("picture");
        const source = document.createElement("source");
        source.srcset = photo.WebP;
        source.type = "image/webp";
        picture.appendChild(source);
        picture.appendChild(img);
        item.appendChild(picture);
      } else {
        item.appendChild(img);
      }
      container.appendChild(item);
    });
  }

  function loadPhotos() {
    if (loading || exhausted) return;
    loading = true;
    loader.style.display = "block";

    fetch(`/photos?offset=${offset}`)
      .then((r) => r.json())
      .then((data) => {
        const photos = data.photos || [];
        if (photos.length === 0) {
          exhausted = true;
          loader.style.display = "none";
          return;
        }
        appendPhotos(photos);
        offset = data.offset;
        loader.style.display = "none";
        loading = false;
      })
      .catch((err) => {
        console.error("Error loading photos:", err);
        loader.style.display = "none";
        loading = false;
      });
  }

  window.addEventListener("scroll", () => {
    if (exhausted || loading) return;
    const threshold = 300;
    if (window.innerHeight + window.scrollY >= document.body.offsetHeight - threshold) {
      loadPhotos();
    }
  });

  // Initial load — don't load again if server already rendered photos
  if (container.children.length === 0) {
    loadPhotos();
  } else {
    offset = container.children.length;
  }
})();