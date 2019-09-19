package io.pivotal.example.servicebroker.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import io.pivotal.example.servicebroker.model.Catalog;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.File;
import java.io.IOException;

@RestController
@RequestMapping("/v2/catalog")
public class CatalogController {

    private static final Logger LOG = LoggerFactory.getLogger(CatalogController.class);

    @GetMapping
    public Catalog catalog() {
        Catalog catalog = new Catalog();
        try {
            File yml = new File("catalog.yml");
            catalog = new ObjectMapper(new YAMLFactory()).readValue(yml, Catalog.class);
        } catch (IOException ioe) {
            LOG.error("catalog.yml NOT FOUND!");
        }

        return catalog;
    }

}
