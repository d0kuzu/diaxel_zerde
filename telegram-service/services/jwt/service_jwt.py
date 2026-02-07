import jwt
import time


def create_service_jwt(secret: str, exp) -> str:
    payload = {
        "sub": "telegram-service",
        "aud": "ai-service",
        "exp": exp,
        "scope": "internal"
    }

    token = jwt.encode(payload, secret, algorithm="HS256")
    return token