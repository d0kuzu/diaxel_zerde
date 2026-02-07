from dataclasses import dataclass
from pathlib import Path


@dataclass(frozen=True)
class Paths:
    _BASE_DIR: Path = Path(__file__).resolve().parent.parent.parent


PATHS = Paths()
