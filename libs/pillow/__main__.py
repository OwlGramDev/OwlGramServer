import sys
import json
from get_font import get_font
from get_glyph import get_glyph
from base64 import b64encode

stdin = sys.stdin.read()
data = json.loads(stdin)
emojiFont = get_font(data)
glyphList = data['glyphs']
glyphs = {}

for page in glyphList:
    glyphs[page] = {}
    for section in glyphList[page]:
        emojiInfo = glyphList[page][section]
        glyphRes = None
        if emojiInfo is not None:
            byteGlyph = get_glyph(emojiFont, emojiInfo['emoji'])
            if byteGlyph is not None:
                glyphRes = b64encode(byteGlyph).decode('utf-8')
        glyphs[page][section] = glyphRes

sys.stdout.write(json.dumps(glyphs))
