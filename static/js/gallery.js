(function () {
  let offset = 0;
  let loading = false;
  let exhausted = false;

  const loader    = document.getElementById("loader");
  const container = document.getElementById("photo-container");
  const modal     = document.getElementById("photo-modal");
  const modalImg  = document.getElementById("modal-img");
  const modalClose= document.getElementById("modal-close");

  // ── Modal helpers ──────────────────────────────────────────────────────
  function openModal(src) {
    modalImg.src = src;
    modal.style.display = "flex";
    document.body.style.overflow = "hidden";
  }

  function closeModal() {
    modal.style.display = "none";
    modalImg.src = "";
    document.body.style.overflow = "";
  }

  // Close on backdrop click (not on the image itself)
  modal.addEventListener("click", function (e) {
    if (e.target === modal) closeModal();
  });

  // Close button
  modalClose.addEventListener("click", closeModal);

  // Close on Escape key
  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape") closeModal();
  });

  // Swipe-down to close on mobile
  var touchStartY = 0;
  modal.addEventListener("touchstart", function (e) {
    touchStartY = e.touches[0].clientY;
  }, { passive: true });
  modal.addEventListener("touchend", function (e) {
    if (e.changedTouches[0].clientY - touchStartY > 60) closeModal();
  }, { passive: true });

  // ── Click delegation — works for SSR + dynamically loaded photos ───────
  container.addEventListener("click", function (e) {
    const img = e.target.closest("img");
    if (!img) return;
    // Prefer the <source> srcset (WebP) if available, else fall back to img.src
    const picture = img.closest("picture");
    const src = (picture && picture.querySelector("source"))
      ? picture.querySelector("source").srcset
      : img.src;
    openModal(src);
  });

  // ── Infinite scroll ────────────────────────────────────────────────────
  function appendPhotos(photos) {
    photos.forEach((photo) => {
      const item = document.createElement("div");
      item.className = "photo-item";
      const img = document.createElement("img");
      img.src = photo.URL;
      img.alt = "Photo";
      img.loading = "lazy";
      img.style.cursor = "pointer";
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
    if (window.innerHeight + window.scrollY >= document.body.offsetHeight - 300) {
      loadPhotos();
    }
  });

  // Initial load — don't load again if server already rendered photos
  if (container.children.length === 0) {
    loadPhotos();
  } else {
    offset = container.children.length;
    // Make SSR images show pointer cursor
    container.querySelectorAll("img").forEach(function(img) {
      img.style.cursor = "pointer";
    });
  }
})();