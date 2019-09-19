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
 * The `catalog` must exist in order to create a service broker.
 *
 * If it does not you will receive this error:
 *
 * → cf create-service-broker service-broker user pass https://java-service-broker.apps.oldsilver.cf-app.com
 * Creating service broker service-broker as admin...
 * The service broker rejected the request to https://java-service-broker.apps.oldsilver.cf-app.com/v2/catalog. Status Code: 404 Not Found, Body: {"timestamp":"2019-09-19T14:44:56.961+0000","status":404,"error":"Not Found","message":"No message available","path":"/v2/catalog"}
 * FAILED
 *
 * When it exists:
 * → cf create-service-broker service-broker user pass https://java-service-broker.apps.oldsilver.cf-app.com
 * Creating service broker service-broker as admin...
 * OK
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
