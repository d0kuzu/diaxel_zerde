from dataclasses import dataclass


@dataclass(frozen=True)
class WelcomeTexts:
    welcome_text: str = """asd"""


@dataclass(frozen=True)
class Texts:
    welcome_texts: WelcomeTexts = WelcomeTexts()


TEXTS: Texts = Texts()