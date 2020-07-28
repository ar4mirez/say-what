FROM python:3.6-slim

ENV LANG="C.UTF-8"
ENV LC_ALL="C.UTF-8"
ENV FLASK_APP=app.py
ENV FLASK_ENV=production
ENV FLASK_DEBUG=0
ENV PORT=3000

RUN mkdir -p /home/python/code
WORKDIR /home/python/code

ADD ./requirements.txt .
RUN pip install -r requirements.txt

ADD . .

CMD python app.py
EXPOSE $PORT