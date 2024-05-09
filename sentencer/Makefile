install_deps_only:
	poetry install -qn --only main --no-root

install_project:
	poetry install --only-root

post_install: install_nltk install_spacy

install_nltk:
	python -c "import nltk; nltk.download('punkt')"

install_spacy:
	python -m spacy download en_core_web_sm
	python -m spacy download de_core_news_sm

full_install: install_deps_only install_project post_install

.PHONY: install_deps_only install_project post_install full_install
