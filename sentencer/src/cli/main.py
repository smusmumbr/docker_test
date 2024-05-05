import click
import io

@click.group()
def cli():
    pass

@cli.command()
@click.option('-l', '--language', type=click.Choice(('en', 'de')), default='en')
@click.option('-f','--filename',type=click.File(), required=True)
def sentencize(language: str, filename: io.IOBase) -> None:
    print(f'lang: {language}', filename.read())

if __name__ == "__main__":
    cli()
