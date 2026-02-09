from aiogram import Router
from aiogram.types import Message

from config.config import Environ
from services.gateway.AI import send_request_to_endpoint

router = Router()

@router.message()
async def on_message(message: Message, env: Environ):
    response = send_request_to_endpoint(env.service.gateway_token, str(message.from_user.id), message.text, env.service.target_endpoint)
    await message.answer(str(response))