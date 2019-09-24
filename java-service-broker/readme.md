# Minimal Java Service Broker

This minimal implementation is intended to be a starting place for service brokers and for learning about service brokers. It assumes the user has 
a base understanding of [Services](https://docs.pivotal.io/pivotalcf/services/index.html)
and the [Open Service Broker API](https://github.com/openservicebrokerapi/servicebroker). 

## Endpoints Implemented

This implementation only provides basic API endpoints for:

[Custom Services Overview](https://docs.pivotal.io/pivotalcf/services/overview.html)

### Catalog Controller

- [Catalog Management](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#catalog-management)

### Service Instance Controller
- [Provisioning](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#provisioning)
- [Binding](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#binding)
- [Unbinding](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#unbinding)
- [Deprovisioning](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#deprovisioning)

It currently only implements the APIs, making them do something useful is up to you.

## Other Resources

Some other resources for building more advanced Service Brokers are:

- [Spring Cloud Open Service Broker API](https://spring.io/projects/spring-cloud-open-service-broker)
- [Spring App Broker](https://spring.io/projects/spring-cloud-app-broker)
