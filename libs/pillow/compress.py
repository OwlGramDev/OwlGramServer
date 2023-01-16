import sys
import json
from io import BytesIO
from base64 import b64encode, b64decode
from PIL import Image

data = json.loads(sys.stdin.read())
result = {}

margin = data['margin']
for emoji in data['emojis']:
    img = Image.open(BytesIO(b64decode(data['emojis'][emoji])))
    img = img.convert('RGBA')
    width, height = img.size
    img2 = Image.new('RGBA', (width + margin, height + margin))
    img2.paste(img, (margin // 2, margin // 2))
    img2 = img2.resize((66, 66), Image.LANCZOS)
    # img2 = img2.convert('P', palette=Image.WEB)
    img_byte_arr = BytesIO()
    img2.save(img_byte_arr, format='PNG', optimize=True, quality=100)
    result[emoji] = b64encode(img_byte_arr.getvalue()).decode('utf-8')

sys.stdout.write(json.dumps(result))
