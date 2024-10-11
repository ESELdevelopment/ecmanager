FROM python:3-alpine

COPY requirements.txt /app/requirements.txt
COPY fixture.py app/fixture.py

RUN pip install -r /app/requirements.txt

CMD ["python", "/app/fixture.py"]
