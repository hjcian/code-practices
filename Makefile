.PHONY: venv

venv:
	@virtualenv -p python3 env
	@source ./env/bin/activate