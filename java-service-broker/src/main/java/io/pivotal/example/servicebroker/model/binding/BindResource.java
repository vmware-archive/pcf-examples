package io.pivotal.example.servicebroker.model.binding;

import java.net.URI;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class BindResource {
    
    String appGuid;

    URI route;

    public String getAppGuid() {
        return this.appGuid;
    }

    public void setAppGuid(String appGuid) {
        this.appGuid = appGuid;
    }

    public URI getRoute() {
        return this.route;
    }

    public void setRoute(URI route) {
        this.route = route;
    }

    @Override
    public String toString() {
        return "{" +
            " appGuid='" + getAppGuid() + "'" +
            ", route='" + getRoute() + "'" +
            "}";
    }
    
}