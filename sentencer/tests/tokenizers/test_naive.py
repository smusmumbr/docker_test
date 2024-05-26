from sentencer.tokenizers.naive import tokenize

import pytest


def test_separators():
    text = 'a regular sentence. a question? a stressed one!'
    expected = ['a regular sentence', 'a question', 'a stressed one']
    assert tokenize(text, 'en') == expected


@pytest.mark.parametrize(
    'lang',
    (('en',), ('de',), ('fr')),
)
def test_languages(lang):
    text = "whatever text. e.g. should be separated"
    expected = ['whatever text', 'e', 'g', 'should be separated']
    assert tokenize(text, lang) == expected
