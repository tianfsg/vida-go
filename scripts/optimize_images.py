"""
optimize_images.py — vida-go gallery optimizer
================================================
Corre desde la raíz del proyecto:
    pip install Pillow
    python optimize_images.py

Qué hace:
  1. Corrige orientación EXIF (fotos de móvil que aparecen rotadas)
  2. Recomprime cada JPG en static/images/data/ a calidad 82
  3. Genera .webp en static/images/data/webp/ con orientación correcta
  4. Redimensiona imágenes que superen 1920px de ancho
  5. Muestra cuánto peso se ahorró

Estructura resultante:
  static/images/data/photo1.jpg
  static/images/data/webp/photo1.webp  ← sube esto al servidor
"""

from PIL import Image, ImageOps
import os, sys

IMAGE_DIR    = "static/images/data"
WEBP_DIR     = "static/images/data/webp"
MAX_WIDTH    = 1920
JPG_QUALITY  = 82
WEBP_QUALITY = 80

def fmt(n):
    return f"{n/1024:.1f} KB"

def process():
    os.makedirs(WEBP_DIR, exist_ok=True)

    files = [f for f in os.listdir(IMAGE_DIR)
             if f.lower().endswith((".jpg", ".jpeg", ".png"))]

    if not files:
        print("No images found in", IMAGE_DIR)
        return

    print(f"Processing {len(files)} images\n")
    print(f"{'File':<30} {'Original':>10} {'JPG':>10} {'WebP':>10} {'Saved':>10}")
    print("-" * 65)

    total_original = total_jpg = total_webp = 0

    for filename in sorted(files):
        fpath    = os.path.join(IMAGE_DIR, filename)
        base     = os.path.splitext(filename)[0]
        webp_out = os.path.join(WEBP_DIR, base + ".webp")

        original_size = os.path.getsize(fpath)

        try:
            img = Image.open(fpath)

            # Fix EXIF orientation — critical for photos taken with phones
            img = ImageOps.exif_transpose(img)
            img = img.convert("RGB")

            # Resize if too wide
            if img.width > MAX_WIDTH:
                ratio = MAX_WIDTH / img.width
                img   = img.resize((MAX_WIDTH, int(img.height * ratio)), Image.LANCZOS)

            # Recompress JPG in-place
            img.save(fpath, "JPEG", quality=JPG_QUALITY, optimize=True, progressive=True)
            jpg_size = os.path.getsize(fpath)

            # Save WebP to webp/ subfolder
            img.save(webp_out, "WEBP", quality=WEBP_QUALITY, method=6)
            webp_size = os.path.getsize(webp_out)

            saving = original_size - min(jpg_size, webp_size)
            total_original += original_size
            total_jpg      += jpg_size
            total_webp     += webp_size

            print(f"{filename:<30} {fmt(original_size):>10} {fmt(jpg_size):>10} {fmt(webp_size):>10} {fmt(saving):>10}")

        except Exception as ex:
            print(f"  ERROR {filename}: {ex}")

    print("-" * 65)
    print(f"\n{'Total original:':<35} {fmt(total_original)}")
    print(f"{'Total after JPG recompression:':<35} {fmt(total_jpg)}")
    print(f"{'Total WebP:':<35} {fmt(total_webp)}")
    print(f"{'Saved (WebP vs original):':<35} {fmt(total_original - total_webp)}")
    print(f"\nWebP files written to: {WEBP_DIR}/")
    print("\nNow upload the webp folder to the server:")
    print(f"  rsync -avz --progress {WEBP_DIR}/ server:/home/vidanuun/public_html/static/images/data/")

if __name__ == "__main__":
    if not os.path.isdir(IMAGE_DIR):
        print(f"ERROR: '{IMAGE_DIR}' not found. Run from the root of vida-go/")
        sys.exit(1)
    process()