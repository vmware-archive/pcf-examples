# copied from http://flask.pocoo.org/snippets/8/

import os
from functools import wraps

from flask import request, Response

# tile-generator auto-creates SECURITY_USER_NAME and SECURITY_USER_PASSWORD as the broker's basic auth creds
api_username = os.getenv('SECURITY_USER_NAME')
api_password = os.getenv('SECURITY_USER_PASSWORD')


def check_auth(username, password):
    return username == api_username and password == api_password


def authenticate():
    return Response(
        'Basic authentication failure', 401,
        {'WWW-Authenticate': 'Basic realm="Login Required"'}
    )


def requires_auth(wrapped):
    @wraps(wrapped)
    def decorated(*args, **kwargs):
        auth = request.authorization
        if not auth or not check_auth(auth.username, auth.password):
            return authenticate()
        return wrapped(*args, **kwargs)

    return decorated
