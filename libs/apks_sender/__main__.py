import json
import sys

import asyncio
import os

from pyrogram import Client

from utils import Utils


async def __main__():
    print("[*] Starting APKs sender...")
    current_path = os.path.abspath(os.path.join(__file__, "../"))
    cache_path = os.path.join(current_path, "cache")
    if not os.path.exists(cache_path):
        os.mkdir(cache_path)
    data = json.loads(sys.stdin.read())
    utils = Utils(current_path)
    client = Client('OwlGram', data['api_id'], data['api_hash'], bot_token=data['bot_token'])
    await client.start()
    await utils.upload_apks(client, data)


asyncio.new_event_loop().run_until_complete(__main__())
