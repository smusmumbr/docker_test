import nltk


LANGUAGE_MAP = {
    'en': 'english',
    'de': 'german',
}
DEFAULT_LANGUAGE='english'


def tokenize(text: str, language: str) -> list[str]:
    return nltk.tokenize.sent_tokenize(text, LANGUAGE_MAP.get(language, DEFAULT_LANGUAGE))
