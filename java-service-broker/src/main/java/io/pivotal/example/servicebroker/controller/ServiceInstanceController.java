package io.pivotal.example.servicebroker.controller;

import io.pivotal.example.servicebroker.model.ServiceInstance;
import io.pivotal.example.servicebroker.model.ProvisionResponse;
import io.pivotal.example.servicebroker.model.ProvisionRequest;
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

     → cf create-service java-service-broker standard my-java-service-broker
     Creating service instance my-java-service-broker in org peter / space broker as admin...
     OK

     */
    /*
        TODO: Implement `accepts_incomplete` parameter

        TODO: Implement the other response codes:
        https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#response-3
     */
    @PutMapping(value = "/{instance_id}")
    @ResponseStatus(HttpStatus.OK)
    public ProvisionResponse provision(@PathVariable( "instance_id" ) String instanceId,
                                       @RequestBody ProvisionRequest request) {
        LOG.debug("PROVISION: " + request.toString());

        // Provision your service

        ProvisionResponse provisionResponse = new ProvisionResponse();
        provisionResponse.setOperation("my-provision-operation");

        LOG.debug(provisionResponse.toString());

        return provisionResponse;
    }

    /**
     * Delete / deprovision a service instance
     *
     * https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#deprovisioning
     → cf delete-service my-jsb

     Really delete the service my-jsb?> y
     Deleting service my-jsb in org test / space dev as admin...
     OK

     * @param instanceId
     * @param serviceId
     * @param planId
     * @return
     */
    /*
        TODO: Implement `accepts_incomplete` parameter

        TODO: Implement the other response codes:
        https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#response-9
     */
    @DeleteMapping(value = "/{instance_id}")
    @ResponseStatus(HttpStatus.OK)
    public ProvisionResponse deprovision(@PathVariable( "instance_id" ) String instanceId,
                                         @RequestParam( "service_id" ) String serviceId,
                                         @RequestParam( "plan_id" ) String planId) {
        LOG.debug("DEPROVISION: instance=" + instanceId + ", service=" + serviceId + ", plan=" + planId );

        // Deprovision your service

        ProvisionResponse provisionResponse = new ProvisionResponse();
        provisionResponse.setOperation("my-deprovision-operation");

        LOG.debug(provisionResponse.toString());

        return provisionResponse;
    }
}
