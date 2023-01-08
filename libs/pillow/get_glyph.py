from io import BytesIO
from typing import Optional

from PIL import Image, ImageDraw, ImageFont


def get_glyph(font: ImageFont, emoji: str) -> Optional[bytes]:
    box_data = font.getbbox(emoji)
    width, height = box_data[2:]
    box_size = max(width, height)
    img = Image.new('RGBA', (box_size, box_size))
    d = ImageDraw.Draw(img)
    x = (box_size - width) / 2
    y = (box_size - height) / 2
    d.text((x, y), emoji, font=font, embedded_color=True)
    is_transparent = True
    for d in img.getextrema():
        is_transparent = is_transparent and d[0] == 0 and d[1] == 0
    if is_transparent:
        return None
    img = img.resize((66, 66), Image.ANTIALIAS)
    img_byte_arr = BytesIO()
    img = img.convert("P")
    img.save(img_byte_arr, format='PNG', optimize=True, quality=95)
    return img_byte_arr.getvalue()
