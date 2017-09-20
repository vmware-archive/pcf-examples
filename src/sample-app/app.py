import json
import os
import sys
import traceback

import flask
import requests

app = flask.Flask(__name__)


@app.route("/")
def index():
    error = flask.request.args.get("error")
    db_response = requests.get(
        app.config['sb_uri'],
        auth=(app.config['sb_username'], app.config['sb_password']),
        verify=False
    )
    if db_response.status_code != 200:
        print(db_response)
        return "", 500

    return flask.render_template('index.j2.html', key_values=db_response.json(), error=error)


@app.route("/put", methods=['POST'])
def put():
    key = flask.request.form.get("key", "").strip()
    value = flask.request.form.get("value", "").strip()
    if key == "" or value == "":
        # todo: error message
        return flask.redirect("/?error=Key and value are required")
    db_url = app.config['sb_uri'] + "/" + key
    db_response = requests.put(
        db_url,
        auth=(app.config['sb_username'], app.config['sb_password']),
        verify=False,
        data=value
    )
    if db_response.status_code != 200:
        print(db_response)
        # todo: error message
        return flask.redirect("/")
    return flask.redirect("/")


@app.route("/delete", methods=['POST'])
def delete():
    key = flask.request.form.get("key", "").strip()
    db_url = app.config['sb_uri'] + "/" + key
    print("making delete request to", db_url)
    db_response = requests.delete(
        db_url,
        auth=(app.config['sb_username'], app.config['sb_password']),
        verify=False
    )
    if db_response.status_code != 200:
        print(db_response)
        # todo: error message
        return flask.redirect("/")
    return flask.redirect("/")


def configure_app(app: flask.Flask):
    vcap_services = json.loads(os.getenv('VCAP_SERVICES', '{}'))
    service_instance = vcap_services.get('spacebears-db', [])
    if len(service_instance) != 1:
        print("Missing 1 'spacebears-db' in bound services ")
        sys.exit(1)

    credentials = service_instance[0].get('credentials')
    app.config['sb_uri'] = credentials['uri']
    app.config['sb_username'] = credentials['username']
    app.config['sb_password'] = credentials['password']


if __name__ == "__main__":
    try:
        configure_app(app)
        app.run(host='0.0.0.0', port=os.getenv('PORT', '8080'))
        print("Exited normally")
    except:
        print("* Exited with exception")
        traceback.print_exc()
