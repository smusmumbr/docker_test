import io

import click

@click.group()
def cli():
    pass

@cli.command()
@click.option('-l', '--language', type=click.Choice(('en', 'de')), default='en')
@click.option('-f','--filename',type=click.File(), required=True)
def tokenize(language: str, filename: io.IOBase) -> None:
    _print_tokenized(filename.read(), language)

@cli.command()
@click.argument('text')
def test(text: str) -> None:
    _print_tokenized(text, language='en')
    

def _print_tokenized(text: str, language: str) -> None:
    print('To be implemented...')

if __name__ == "__main__":
    cli()
