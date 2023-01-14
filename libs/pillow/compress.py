import sys
import json
from io import BytesIO
from base64 import b64encode, b64decode
from PIL import Image

data = json.loads(sys.stdin.read())
result = {}

for emoji in data:
    img = Image.open(BytesIO(b64decode(data[emoji])))
    img = img.resize((66, 66), Image.ANTIALIAS)
    img_byte_arr = BytesIO()
    img = img.convert("P")
    img.save(img_byte_arr, format='PNG', optimize=True, quality=95)
    result[emoji] = b64encode(img_byte_arr.getvalue()).decode('utf-8')

sys.stdout.write(json.dumps(result))
