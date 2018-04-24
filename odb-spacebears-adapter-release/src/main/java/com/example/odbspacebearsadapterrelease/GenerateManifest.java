package com.example.odbspacebearsadapterrelease;

import org.springframework.util.FileCopyUtils;

import java.io.File;
import java.util.Scanner;

public class GenerateManifest {

    public void generateManifest() {

        ClassLoader classLoader = getClass().getClassLoader();
        File file = new File(classLoader.getResource("manifest.yml").getFile());

        try (Scanner scanner = new Scanner(file)) {

            while (scanner.hasNextLine()) {
                String line = scanner.nextLine();
                System.out.println(line);
            }

            scanner.close();

        } catch (Exception e) {
            e.printStackTrace();
        }

        System.out.println();

    }
}
