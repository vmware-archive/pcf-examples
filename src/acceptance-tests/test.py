import os
import subprocess
import time


def main(cf_path, api_url, admin_user, admin_pass, skip_ssl):
    # echo '' is to deal with cf command prompting for org/space after login
    cmd_login = "echo '' | {cf} login -a {api_url} -u {admin_user} -p {admin_pass}".format(
        cf=cf_path, api_url=api_url, admin_user=admin_user, admin_pass=admin_pass
    )
    if skip_ssl:
        cmd_login = cmd_login + " --skip-ssl-validation"
    run_cmd(cmd_login)

    discriminator = int(time.time())
    org = "simple_smoke_test_org_{}".format(discriminator)
    run_cmd("{cf} create-org {org}".format(cf=cf_path, org=org))
    run_cmd("{cf} target -o {org}".format(cf=cf_path, org=org))
    run_cmd("{cf} create-space tests".format(cf=cf_path))
    run_cmd("{cf} target -s tests".format(cf=cf_path))

    run_cmd("{cf} create-service spacebears-db plan1 my-spacebears".format(cf=cf_path))
    run_cmd("{cf} create-service-key my-spacebears spacebears-key".format(cf=cf_path))

    # todo: here. how to capture output?
    run_cmd("{cf} service-key my-spacebears spacebears-key".format(cf=cf_path))


def run_cmd(cmd):
    print("running:\n    {}".format(cmd))
    exit_code = subprocess.call(cmd, shell=True)
    if exit_code != 0:
        exit(exit_code)


if __name__ == "__main__":
    cf_path = os.getenv('CF_PATH')
    if not cf_path:
        print("CF_PATH is required")
        exit(1)

    api_url = os.getenv('API_URL')
    if not api_url:
        print("API_URL is required")
        exit(1)
    admin_user = os.getenv('ADMIN_USER')

    if not cf_path:
        print("ADMIN_USER is required")
        exit(1)
    admin_pass = os.getenv('ADMIN_PASS')

    if not admin_pass:
        print("ADMIN_PASS is required")
        exit(1)

    skip_ssl = os.getenv('SKIP_SSL', 'false') == 'true'

    main(cf_path, api_url, admin_user, admin_pass, skip_ssl)
