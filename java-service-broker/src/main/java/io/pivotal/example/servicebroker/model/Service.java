package io.pivotal.example.servicebroker.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

import javax.validation.constraints.NotEmpty;
import java.util.List;
import java.util.Map;

@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Service {

    @NotEmpty
    private String id;

    @NotEmpty
    private String name;

    @NotEmpty
    private  String description;

    private List<String> tags;

    private List<String> requires;

    private Boolean bindable;

    private Boolean instancesRetrievable;

    private Boolean bindingsRetrievable;

    private Boolean allowContextUpdates;

    private Map<String, Object> metadata;

    private Boolean planUpdateable;

    @NotEmpty
    private List<Plan> plans;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public List<String> getTags() {
        return tags;
    }

    public void setTags(List<String> tags) {
        this.tags = tags;
    }

    public List<String> getRequires() {
        return requires;
    }

    public void setRequires(List<String> requires) {
        this.requires = requires;
    }

    public Boolean getBindable() {
        return bindable;
    }

    public void setBindable(Boolean bindable) {
        this.bindable = bindable;
    }

    public Boolean getInstancesRetrievable() {
        return instancesRetrievable;
    }

    public void setInstancesRetrievable(Boolean instancesRetrievable) {
        this.instancesRetrievable = instancesRetrievable;
    }

    public Boolean getBindingsRetrievable() {
        return bindingsRetrievable;
    }

    public void setBindingsRetrievable(Boolean bindingsRetrievable) {
        this.bindingsRetrievable = bindingsRetrievable;
    }

    public Boolean getAllowContextUpdates() {
        return allowContextUpdates;
    }

    public void setAllowContextUpdates(Boolean allowContextUpdates) {
        this.allowContextUpdates = allowContextUpdates;
    }

    public Map<String, Object> getMetadata() {
        return metadata;
    }

    public void setMetadata(Map<String, Object> metadata) {
        this.metadata = metadata;
    }

    public Boolean getPlanUpdateable() {
        return planUpdateable;
    }

    public void setPlanUpdateable(Boolean planUpdateable) {
        this.planUpdateable = planUpdateable;
    }

    public List<Plan> getPlans() {
        return plans;
    }

    public void setPlans(List<Plan> plans) {
        this.plans = plans;
    }
}
