from aiogram import Router
from aiogram.types import Message

from config.config import Environ
from services.gateway.AI import send_request_to_endpoint
from services.jwt.service_jwt import create_service_jwt

router = Router()

@router.message()
async def on_message(message: Message, env: Environ):
    token = create_service_jwt(env.service.secret, env.service.exp)
    response = send_request_to_endpoint(token, message.text, env.service.target_endpoint)
    await message.answer(str(response))