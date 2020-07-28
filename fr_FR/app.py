import os

from translator import fr_FR
import http
from flask import Flask, request, Response

app = Flask(__name__)


@app.route('/translate')
def translate():
    word = request.args['word']
    _word_translate = fr_FR.en_to_fr(word)

    return {"word": _word_translate}, http.HTTPStatus.OK


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=os.getenv('PORT', '3000'))
