// gallery.js — infinite scroll for the photo gallery

(function () {
  let offset = 0;
  let loading = false;
  let exhausted = false;

  const loader = document.getElementById('loader');
  const container = document.getElementById('photo-container');

  function loadPhotos() {
    if (loading || exhausted) return;
    loading = true;

    if (loader) loader.style.display = 'block';

    fetch(`/photos?offset=${offset}`)
      .then(response => response.json())
      .then(data => {
        const photos = data.photos || [];

        if (photos.length === 0) {
          exhausted = true;
          return;
        }

        photos.forEach(photo => {
          const item = document.createElement('div');
          item.classList.add('photo-item');

          const picture = document.createElement('picture');

          if (photo.WebP) {
            const source = document.createElement('source');
            source.srcset = photo.WebP;
            source.type = 'image/webp';
            picture.appendChild(source);
          }

          const img = document.createElement('img');
          img.src = photo.URL;
          img.alt = 'Photo';
          img.loading = 'lazy';
          img.classList.add('photo');

          picture.appendChild(img);
          item.appendChild(picture);
          container.appendChild(item);
        });

        offset = data.offset;
      })
      .catch(err => console.error('Error loading photos:', err))
      .finally(() => {
        loading = false;
        if (loader) loader.style.display = 'none';
      });
  }

  // Load more when near bottom
  window.addEventListener('scroll', () => {
    if (window.innerHeight + window.scrollY >= document.body.offsetHeight - 300) {
      loadPhotos();
    }
  });

  // Initial load
  loadPhotos();
})();