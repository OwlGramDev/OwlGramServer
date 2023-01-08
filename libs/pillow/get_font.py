from io import BytesIO
from base64 import b64decode
from PIL import ImageFont


def get_font(data: dict) -> ImageFont:
    input_data = b64decode(data['emoji_file'])
    tmp_font = None
    for i in range(0, 200):
        try:
            tmp_font = ImageFont.truetype(BytesIO(input_data), i)
        except OSError:
            continue
    return tmp_font
