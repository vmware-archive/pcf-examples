package io.pivotal.example.servicebroker.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

import java.net.URI;
import java.util.Map;

@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class ProvisionResponse {

    String serviceId;

    String planId;

    URI dashboardUrl;

    String operation;

    Map<String, Object> parameters;

    public String getServiceId() {
        return serviceId;
    }

    public void setServiceId(String serviceId) {
        this.serviceId = serviceId;
    }

    public String getPlanId() {
        return planId;
    }

    public void setPlanId(String planId) {
        this.planId = planId;
    }

    public URI getDashboard_url() {
        return dashboardUrl;
    }

    public void setDashboard_url(URI dashboard_url) {
        this.dashboardUrl = dashboard_url;
    }

    public String getOperation() {
        return operation;
    }

    public void setOperation(String operation) {
        this.operation = operation;
    }

    public Map<String, Object> getParameters() {
        return parameters;
    }

    public void setParameters(Map<String, Object> parameters) {
        this.parameters = parameters;
    }

    @Override
    public String toString() {
        return "ProvisionResponse{" +
                "serviceId='" + serviceId + '\'' +
                ", planId='" + planId + '\'' +
                ", dashboardUrl=" + dashboardUrl +
                ", operation='" + operation + '\'' +
                ", parameters=" + parameters +
                '}';
    }
}
