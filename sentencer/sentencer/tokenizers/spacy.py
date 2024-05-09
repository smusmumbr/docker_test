import spacy

MODELS_MAP={
    'en': 'en_core_web_sm',
    'de': 'de_core_news_sm',
}


def tokenize(text: str, language: str) -> list[str]:
    model_name=MODELS_MAP.get(language)
    if model_name is None:
        raise ValueError(f'Unknown language: {language}')
    nlp = spacy.load(model_name)
    doc = nlp(text)
    assert doc.has_annotation("SENT_START")
    return [sent.text for sent in doc.sents]
