package io.pivotal.example.servicebroker.controller;

import io.pivotal.example.servicebroker.model.ServiceInstance;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

/**
 *
 */
@RestController
@RequestMapping("/v2/service_instances")
public class ServiceInstanceController {

    private static final Logger LOG = LoggerFactory.getLogger(CatalogController.class);

    /**
     * Create / Provision a Service Instance
     * https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#provisioning
     */
    @PutMapping(value = "/{id}")
    @ResponseStatus(HttpStatus.OK)
    public String provision(@PathVariable( "id" ) String id,
                            @RequestBody ServiceInstance serviceInstance) {
        LOG.debug(serviceInstance.toString());

        // Do provision things here!

        return id;
    }
}
