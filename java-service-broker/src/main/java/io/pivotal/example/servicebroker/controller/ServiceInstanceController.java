package io.pivotal.example.servicebroker.controller;

import io.pivotal.example.servicebroker.model.ProvisionResponse;
import io.pivotal.example.servicebroker.model.ProvisionRequest;
import io.pivotal.example.servicebroker.model.binding.BindingRequest;
import io.pivotal.example.servicebroker.model.binding.BindingResponse;
import io.pivotal.example.servicebroker.model.binding.Credentials;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.UUID;

/**
 * Provision, deprovision, bind and unbind operations.
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

    /**
     * Bind
     *
     * https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#binding
     *

     → cf bind-service spring-music my-java-service-broker --binding-name my-java-binding
     Binding service my-java-service-broker to app spring-music with binding name my-java-binding in org test / space dev as admin...
     OK

     → cf services
     Getting services in org test / space dev as admin...

     name                     service               plan       bound apps     last operation     broker           upgrade available
     my-java-service-broker   java-service-broker   standard   spring-music   create succeeded   service-broker

     * @param instanceId
     * @param bindingId
     * @param request
     * @return
     */
    /*
        TODO: Implement `accepts_incomplete` parameter

        TODO: Implement the other response codes:
        https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#response-5
     */
    @PutMapping(value = "/{instance_id}/service_bindings/{binding_id}")
    @ResponseStatus(HttpStatus.OK)
    public BindingResponse bind(@PathVariable( "instance_id" ) String instanceId,
                                @PathVariable( "binding_id" ) String bindingId,
                                @RequestBody BindingRequest request) {
        LOG.debug("BIND: instanceId=" + instanceId +
                        " bindingId=" + bindingId +
                        " request=" + request.toString());

        // Bind here

        BindingResponse response = new BindingResponse();
        response.setCredentials(new Credentials(
                "my_username", UUID.randomUUID().toString()
        ));

        LOG.debug("Binding Response: " + response);

        return response;
    }

    /**
     * Unbind
     *
     * https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#unbinding
     *
     → cf unbind-service spring-music my-java-service-broker
     Unbinding app spring-music from service my-java-service-broker in org test / space dev as admin...
     OK

     → cf services
     Getting services in org test / space dev as admin...

     name                     service               plan       bound apps   last operation     broker           upgrade available
     my-java-service-broker   java-service-broker   standard                create succeeded   service-broker

     *
     * @param instanceId
     * @param bindingId
     * @param serviceId
     * @param planId
     * @return
     */
    /*
        TODO: Implement `accepts_incomplete` parameter

        TODO: Implement the other response codes:
        https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#response-8
     */
    @DeleteMapping(value = "/{instance_id}/service_bindings/{binding_id}")
    @ResponseStatus(HttpStatus.OK)
    public BindingResponse unbind(@PathVariable( "instance_id" ) String instanceId,
                                  @PathVariable( "binding_id" ) String bindingId,
                                  @RequestParam( "service_id" ) String serviceId,
                                  @RequestParam( "plan_id" ) String planId) {

        LOG.debug("UNBIND: instance=" + instanceId + ", bindingId=" + bindingId +
                ", service=" + serviceId + ", plan=" + planId );

        // Unbind here
        BindingResponse response = new BindingResponse();
        response.setOperation("my-unbinding-operation");

        LOG.debug("Binding Response: " + response);

        return response;
    }
}
