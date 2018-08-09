package example;

import org.json.*;

public class OdbSpacebearsAdapter {

	public static void main(String[] args) {
		if (args != null && args.length > 0 && "generate-manifest".equals(args[0])) {
            generateManifest(args[1], args[2], args[3]);
            return;
		}
        if (args != null && args.length > 0 && "create-binding".equals(args[0])) {
            createBinding(args[1], args[2], args[3]);
            return;
        }
        System.exit(10);
    }
    
    public static void generateManifest(String serviceDeployment, String a, String b) {
        JSONObject obj = new JSONObject(serviceDeployment);
        String name = obj.getString("deployment_name");

        String s = "---\n" 
        + "name: " + name + "\n"
        + "releases: \n"
        + "  - name: bosh-simple-spacebears\n"
        + "    version: '0+dev.14'\n"
        + "stemcells:\n"
        + "  - alias: bosh-warden-boshlite-ubuntu-trusty-go_agent\n"
        + "    os: ubuntu-trusty\n"
        + "    version: latest\n"
        + "update:\n"
        + "  canaries: 1\n"
        + "  max_in_flight: 10\n"
        + "  canary_watch_time: 1000-30000\n"
        + "  update_watch_time: 1000-30000\n"
        + "instance_groups:\n"
        + "  - name: spacebears_db_node\n"
        + "    instances: 1\n"
        + "    azs: [z1]\n"
        + "    jobs:\n"
        + "     - name: spacebears\n"
        + "       release: bosh-simple-spacebears\n"
        + "    properties:\n"
        + "      spacebears:\n"
        + "        password: symphony27_Trailers\n"
        + "    vm_type: minimal\n"
        + "    cloud_properties:\n"
        + "      tags:\n"
        + "       - allow-ssh\n"
        + "    stemcell: bosh-warden-boshlite-ubuntu-trusty-go_agent\n"
        + "    persistent_disk_type: 5GB\n"
        + "    networks:\n"
        + "    - name: default\n";

        System.out.println(s);

    }

    public static void createBinding(String serviceDeployment, String a, String b) {

	    String bindingResponse = "{\n" +
                "  \"credentials\": {\n" +
                "    \"uri\": \"http://localhost:9000/api\",\n" +
                "    \"username\": \"admin\",\n" +
                "    \"password\": \"symphony27_Trailers\"\n" +
                "  }\n" +
                "}";

	    System.out.println(bindingResponse);

    }


}
