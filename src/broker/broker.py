import flask
import json
import os
import requests
import traceback

app = flask.Flask(__name__)

db_admin_username = os.getenv('DB_ADMIN_USERNAME')
db_admin_password = os.getenv('DB_ADMIN_PASSWORD')
db_url = os.getenv('DB_URL')


# todo: missing basic auth check on broker api

@app.route("/health")
def health():
    return "healthy"


@app.route("/v2/catalog")
def broker_catalog():
    # catalog ids were randomly generated guids, per best practices
    catalog = {
        "services": [{
            "id": 'c084b262-b733-45e2-974b-ed8ad94e808a',
            "name": 'example-db-service',
            "description": "Simple key/value services",
            "bindable": True,
            "plans": [{
                "id": '30f7be98-dc0b-4fce-91bc-aeb87c864ecc',
                "name": "first-plan",
                "description": "A first, free, service plan"
            }]
        }]
    }
    return json.dumps(catalog, indent=4)


@app.route("/v2/service_instances/<instance_id>", methods=['PUT'])
def broker_provision_instance(instance_id):
    api_response = requests.post("{}/api/admin/bucket/{}".format(db_url, instance_id), verify=False)
    if api_response.status_code > 299:
        print(api_response)
        return "{}", 500
    else:
        return "{}", 201


@app.route("/v2/service_instances/<instance_id>", methods=['DELETE'])
def broker_deprovision_instance(instance_id):
    # delete bucket
    response_body = json.dumps({}, indent=4)
    return response_body, 200


@app.route("/v2/service_instances/<instance_id>/service_bindings/<binding_id>", methods=['PUT'])
def broker_bind_instance(instance_id, binding_id):
    # create credentials
    response_body = json.dumps({}, indent=4)
    return response_body, 201


@app.route("/v2/service_instances/<instance_id>/service_bindings/<binding_id>", methods=['DELETE'])
def broker_unbind_instance(instance_id, binding_id):
    # delete credentials
    response_body = json.dumps({}, indent=4)
    return response_body, 200


@app.errorhandler(500)
def internal_error(error):
    print(error)
    return "Internal server error", 500


if __name__ == "__main__":
    try:
        app.run(host='0.0.0.0', port=os.getenv('PORT', '8080'))
        print("Exited normally")
    except:
        print("* Exited with exception")
        traceback.print_exc()
