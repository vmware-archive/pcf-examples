package io.pivotal.example.servicebroker.model.binding;

import java.util.Map;

import javax.validation.constraints.NotEmpty;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class BindingRequest {

    Map<String, Object> context;

    @NotEmpty
    String serviceId;

    @NotEmpty
    String planId;

    String appGuid;

    BindResource bindResource;

    Map<String, Object> parameters;

    public Map<String,Object> getContext() {
        return this.context;
    }

    public void setContext(Map<String,Object> context) {
        this.context = context;
    }

    public String getServiceId() {
        return this.serviceId;
    }

    public void setServiceId(String serviceId) {
        this.serviceId = serviceId;
    }

    public String getPlanId() {
        return this.planId;
    }

    public void setPlanId(String planId) {
        this.planId = planId;
    }

    public String getAppGuid() {
        return this.appGuid;
    }

    public void setAppGuid(String appGuid) {
        this.appGuid = appGuid;
    }

    public BindResource getBindResource() {
        return this.bindResource;
    }

    public void setBindResource(BindResource bindResource) {
        this.bindResource = bindResource;
    }

    public Map<String,Object> getParameters() {
        return this.parameters;
    }

    public void setParameters(Map<String,Object> parameters) {
        this.parameters = parameters;
    }

    @Override
    public String toString() {
        return "{" +
            " context='" + getContext() + "'" +
            ", serviceId='" + getServiceId() + "'" +
            ", planId='" + getPlanId() + "'" +
            ", appGuid='" + getAppGuid() + "'" +
            ", bindResource='" + getBindResource() + "'" +
            ", parameters='" + getParameters() + "'" +
            "}";
    }

}