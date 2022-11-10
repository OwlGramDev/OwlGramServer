import os.path
from asyncio import subprocess

from pyrogram import Client
from typing import List, Dict

import asyncio
from pyrogram.types import InputMediaDocument

version_list = [
    '1.0',
    '1.1',
    '1.5',
    '1.6',
    '2.0',
    '2.0',
    '2.1',
    '2.2',
    '2.3',
    '2.3.7',
    '3.0',
    '3.1',
    '3.2',
    '4.0.0',
    '4.0.4',
    '4.1',
    '4.2',
    '4.3',
    '4.4',
    '4.4.4',
    '5.0',
    '5.1',
    '6.0',
    '7.0',
    '7.1',
    '8.0',
    '8.1',
    '9',
    '10',
    '11',
    '12'
]


class Utils:
    def __init__(self, path: str):
        self._java_process = None
        self._path = path

    async def md5(self, filename):
        self._java_process = await asyncio.create_subprocess_exec(
            'sha256sum',
            filename,
            stdout=subprocess.PIPE,
        )
        stdout, _ = await self._java_process.communicate()
        return stdout.decode().split(' ')[0]

    async def generate_thumb(self, version: str):
        self._java_process = await asyncio.create_subprocess_exec(
            'php',
            os.path.join(self._path, 'getThumb.php'),
            version,
        )
        await self._java_process.communicate()

    async def copy_apk(self, from_folder, to_folder):
        self._java_process = await asyncio.create_subprocess_exec(
            'cp',
            from_folder,
            to_folder,
        )
        await self._java_process.communicate()

    async def upload_apks(self, client: Client, apks_descriptor: Dict[str, any]):
        await self.generate_thumb(
            apks_descriptor['apks'][0]['version_name'].split(' ')[0]
        )
        list_files: List[InputMediaDocument] = []
        list_temp_files = []
        for apk in apks_descriptor['apks']:
            text = f'<b>Version: </b><code>{apk["version_name"]}</code>\n'
            text += f'<b>Minimum OS: </b><code>{version_list[int(apk["min_sdk_version"]) - 1]}</code>\n'
            text += f'<b>Base: </b><code>{apks_descriptor["base"]}</code>\n'
            text += f'<b>SHA256: </b>\n<code>{await self.md5(apk["path"])}</code>\n'
            text += f'\n#{apk["abi_name"]} #{apks_descriptor["channel"]}'
            apk_name = f'OwlGram-{apk["version_name"]}-{apk["version_code"]}-{apk["abi_name"]}.apk'
            new_folder = os.path.join(self._path, 'cache', apk_name)
            await self.copy_apk(apk['path'], new_folder)
            list_temp_files.append(new_folder)
            list_files.append(InputMediaDocument(
                media=new_folder,
                thumb=os.path.join(self._path, 'output.jpg'),
                caption=text,
            ))
        await client.send_media_group(
            chat_id=int(apks_descriptor["chat_id"]),
            media=list_files,
        )
        for file in list_temp_files:
            os.remove(file)
