install_deps_only:
	poetry install -qn --only main --no-root

install_project:
	poetry install --only-root

post_install:
	python -c "import nltk; nltk.download('punkt')"

full_install: install_deps_only install_project post_install

.PHONY: install_deps_only install_project post_install full_install
