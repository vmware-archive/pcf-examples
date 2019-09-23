package io.pivotal.example.servicebroker.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import io.pivotal.example.servicebroker.model.Catalog;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * Implementation of the Open Service Broker API Spec:
 * https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#catalog-management
 *
 * The `catalog` must exist in order to create a service broker.
 *
 * If it does not you will receive this error:

    → cf create-service-broker service-broker user pass https://MY-DOMAIN.com
    Creating service broker service-broker as admin...
    The service broker rejected the request to https://MY-DOMAIN.com/v2/catalog. Status Code: 404 Not Found, Body: {"timestamp":"2019-09-19T14:44:56.961+0000","status":404,"error":"Not Found","message":"No message available","path":"/v2/catalog"}
    FAILED

 * When a catalog  exists:

    → cf create-service-broker service-broker user pass https://MY-DOMAIN.com
    Creating service broker service-broker as admin...
    OK

 * Check the list for it
    → cf service-brokers
    Getting service brokers as admin...
    name             url
    service-broker   https://MY-DOMAIN.com *

 * Check the service access
    → cf service-access
    Getting service access as admin...
    broker: service-broker
    service               plan       access   orgs
    java-service-broker   standard   none

 * Enable service access
    → cf enable-service-access java-service-broker
    Enabling access to all plans of service java-service-broker for all orgs as admin...
    OK

 * Ensure it shows up in the marketplace
    → cf marketplace
    Getting services from marketplace in org peter / space broker as admin...
    OK

    service               plans      description                    broker
    java-service-broker   standard   A simple java-service-broker   service-broker

 */

@RestController
@RequestMapping("/v2/catalog")
public class CatalogController {

    private static final Logger LOG = LoggerFactory.getLogger(CatalogController.class);

    @Autowired Catalog catalog;

    @GetMapping
    public Catalog catalog() {
        LOG.debug(catalog.toString());

        return catalog;
    }
}
