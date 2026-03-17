"""
check_orientation.py — diagnóstico de orientación EXIF
=======================================================
Corre desde la raíz del proyecto:
    python check_orientation.py

Muestra la orientación EXIF de cada foto y sus dimensiones reales.
"""

from PIL import Image
import os, sys

IMAGE_DIR = "static/images/data"

def get_exif_orientation(img):
    try:
        exif = img._getexif()
        if exif:
            # Tag 274 es Orientation en EXIF
            return exif.get(274, "No orientation tag")
        return "No EXIF data"
    except Exception:
        return "EXIF read error"

files = [f for f in os.listdir(IMAGE_DIR)
         if f.lower().endswith((".jpg", ".jpeg", ".png"))]

print(f"{'File':<30} {'Width':>8} {'Height':>8} {'Orientation':>15} {'Status'}")
print("-" * 75)

for filename in sorted(files):
    fpath = os.path.join(IMAGE_DIR, filename)
    try:
        img = Image.open(fpath)
        orientation = get_exif_orientation(img)
        w, h = img.size
        status = "VERTICAL" if h > w else "horizontal"
        print(f"{filename:<30} {w:>8} {h:>8} {str(orientation):>15}  {status}")
    except Exception as ex:
        print(f"{filename:<30} ERROR: {ex}")