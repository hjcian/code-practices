.PHONY: venv, source

venv:
	@virtualenv -p python3 env

source:
	ln -s env/bin/activate