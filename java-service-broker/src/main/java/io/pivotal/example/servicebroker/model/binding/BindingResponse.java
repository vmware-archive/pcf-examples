package io.pivotal.example.servicebroker.model.binding;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class BindingResponse {

    String operation;

    Credentials credentials;

    String syslogDrainUrl;

    String routeServiceUrl;

    public String getOperation() {
        return operation;
    }

    public void setOperation(String operation) {
        this.operation = operation;
    }

    public Credentials getCredentials() {
        return credentials;
    }

    public void setCredentials(Credentials credentials) {
        this.credentials = credentials;
    }

    public String getSyslogDrainUrl() {
        return syslogDrainUrl;
    }

    public void setSyslogDrainUrl(String syslogDrainUrl) {
        this.syslogDrainUrl = syslogDrainUrl;
    }

    public String getRouteServiceUrl() {
        return routeServiceUrl;
    }

    public void setRouteServiceUrl(String routeServiceUrl) {
        this.routeServiceUrl = routeServiceUrl;
    }

    @Override
    public String toString() {
        return "BindingResponse{" +
                "operation='" + operation + '\'' +
                ", credentials=" + credentials +
                ", syslogDrainUrl='" + syslogDrainUrl + '\'' +
                ", routeServiceUrl='" + routeServiceUrl + '\'' +
                '}';
    }
}