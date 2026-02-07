import requests

FALLBACK_RESPONSE = {"answer": "service unavailable"}


def send_request_to_endpoint(token: str, user_message: str, target_endpoint: str) -> dict:
    try:
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
    except requests.RequestException:
        return FALLBACK_RESPONSE


    if response.status_code != 200:
        return FALLBACK_RESPONSE

    if "application/json" not in response.headers.get("Content-Type", ""):
        return FALLBACK_RESPONSE

    try:
        return response.json()
    except ValueError:
        return FALLBACK_RESPONSE