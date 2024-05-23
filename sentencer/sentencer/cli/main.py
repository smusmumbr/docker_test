import io

from importlib import import_module

import click
import csv

from ..constants import SUPPORTED_LANGUAGES, TOKENIZERS


@click.group()
def cli():
    pass


@cli.command()
@click.option('-t', '--stype', type=click.Choice(TOKENIZERS), default='naive')
@click.option('-l', '--language', type=click.Choice(SUPPORTED_LANGUAGES), default='en')
@click.option('-o','--output',type=click.File('w'), required=True)
@click.option('-i','--input',type=click.File(), required=True)
def tokenize(language: str, input: io.IOBase, output: io.IOBase, stype: str) -> None:
    csvize(_tokenize(input.read(), stype, language), output)


@cli.command()
@click.option('-t', '--stype', type=click.Choice(TOKENIZERS), default='naive')
@click.argument('text')
def test(text: str, stype: str) -> None:
    print(_tokenize(text, stype, language='en'))
    

def _tokenize(text: str, stype: str, language: str) -> list[str]:
    try:
        module = import_module(f'sentencer.tokenizers.{stype}')
    except ImportError:
        raise ValueError(f'Unknown sentencer type: {stype}')
    return module.tokenize(text, language)

def csvize(sentences: list[str], output_filepath: io.IOBase) -> None:
    writer=csv.DictWriter(output_filepath, ('sentence',))
    writer.writeheader()
    writer.writerows([{'sentence': s.strip()} for s in sentences])
