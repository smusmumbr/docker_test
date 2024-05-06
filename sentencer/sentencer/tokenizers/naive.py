import re


SEPARATORS = re.compile(r'[.!?]')


def tokenize(text: str, _: str) -> list[str]:
    return [stripped for s in re.split(SEPARATORS, text) if (stripped := s.strip())]
