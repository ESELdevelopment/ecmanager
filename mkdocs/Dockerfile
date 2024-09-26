FROM squidfunk/mkdocs-material:latest
WORKDIR /docs
COPY . /docs
RUN pip3 install --no-cache-dir  -r requirements.txt
EXPOSE 8000
CMD ["serve", "--dev-addr=0.0.0.0:8000"]
