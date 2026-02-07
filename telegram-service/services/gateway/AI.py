import requests


def send_request_to_endpoint(token: str, user_message: str, target_endpoint: str) -> dict:
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json",
    }

    payload = {
        "user_message": user_message
    }

    response = requests.post(
        target_endpoint,
        json=payload,
        headers=headers,
        timeout=10
    )

    response.raise_for_status()
    return response.json()