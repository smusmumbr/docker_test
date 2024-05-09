import io

from importlib import import_module

import click

from ..constants import SUPPORTED_LANGUAGES, TOKENIZERS


@click.group()
def cli():
    pass


@cli.command()
@click.option('-t', '--stype', type=click.Choice(TOKENIZERS), default='naive')
@click.option('-l', '--language', type=click.Choice(SUPPORTED_LANGUAGES), default='en')
@click.option('-f','--filename',type=click.File(), required=True)
def tokenize(language: str, filename: io.IOBase, stype: str) -> None:
    _print_tokenized(filename.read(), stype, language)


@cli.command()
@click.option('-t', '--stype', type=click.Choice(TOKENIZERS), default='naive')
@click.argument('text')
def test(text: str, stype: str) -> None:
    _print_tokenized(text, stype, language='en')
    

def _print_tokenized(text: str, stype: str, language: str) -> None:
    try:
        module = import_module(f'sentencer.tokenizers.{stype}')
    except ImportError:
        raise ValueError(f'Unknown sentencer type: {stype}')
    print(module.tokenize(text, language))
