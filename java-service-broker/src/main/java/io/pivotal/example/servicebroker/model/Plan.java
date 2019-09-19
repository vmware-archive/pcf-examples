package io.pivotal.example.servicebroker.model;

import com.fasterxml.jackson.annotation.JsonInclude;

import javax.validation.constraints.NotEmpty;
import java.util.Map;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Plan {

    @NotEmpty
    private String id;

    @NotEmpty
    private String name;

    @NotEmpty
    private String description;

    private Map<String, Object> metadata;

    private Boolean free;

    private Boolean bindable;

    private Boolean planUpdateable;

    private Integer maximumPollingDuration;

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

    public Map<String, Object> getMetadata() {
        return metadata;
    }

    public void setMetadata(Map<String, Object> metadata) {
        this.metadata = metadata;
    }

    public Boolean getFree() {
        return free;
    }

    public void setFree(Boolean free) {
        this.free = free;
    }

    public Boolean getBindable() {
        return bindable;
    }

    public void setBindable(Boolean bindable) {
        this.bindable = bindable;
    }

    public Boolean getPlanUpdateable() {
        return planUpdateable;
    }

    public void setPlanUpdateable(Boolean planUpdateable) {
        this.planUpdateable = planUpdateable;
    }

    public Integer getMaximumPollingDuration() {
        return maximumPollingDuration;
    }

    public void setMaximumPollingDuration(Integer maximumPollingDuration) {
        this.maximumPollingDuration = maximumPollingDuration;
    }
}
