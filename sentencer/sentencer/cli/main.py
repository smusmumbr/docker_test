import io

import click


@click.group()
def cli():
    pass


@cli.command()
@click.option('-t', '--stype', type=click.Choice(('naive',)), default='naive')
@click.option('-l', '--language', type=click.Choice(('en', 'de')), default='en')
@click.option('-f','--filename',type=click.File(), required=True)
def tokenize(language: str, filename: io.IOBase, stype: str) -> None:
    _print_tokenized(filename.read(), stype, language)


@cli.command()
@click.option('-t', '--stype', type=click.Choice(('naive',)), default='naive')
@click.argument('text')
def test(text: str, stype: str) -> None:
    _print_tokenized(text, stype, language='en')
    

def _print_tokenized(text: str, stype: str, language: str) -> None:
    if stype == 'naive':
        from ..tokenizers.naive import tokenize as impl_tokenize
    else:
        raise ValueError(f'Unknown sentencer type: {stype}')
    print(impl_tokenize(text, language))
