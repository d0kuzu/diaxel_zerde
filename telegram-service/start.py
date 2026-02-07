import asyncio
import logging

import coloredlogs
from aiogram import Bot, Dispatcher
from aiogram.fsm.storage.memory import MemoryStorage
from aiogram.fsm.storage.redis import RedisStorage
from aiogram.types import BotCommand
from redis.asyncio import Redis

from config.config import Environ
from telegram.register import TgRegister


async def start(environ: Environ):
    bot = Bot(environ.bot.token)

    dp = Dispatcher()

    await bot.set_my_commands([])

    tg_register = TgRegister(dp, bot, environ)
    await tg_register.register()

    await dp.start_polling(bot)

if __name__ == "__main__":
    env = Environ()

    logging.basicConfig(level=env.bot.logging_level)
    coloredlogs.install()
    asyncio.run(start(env))
