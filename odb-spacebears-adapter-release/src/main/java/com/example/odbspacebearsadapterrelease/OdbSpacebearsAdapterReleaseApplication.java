package com.example.odbspacebearsadapterrelease;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class OdbSpacebearsAdapterReleaseApplication {

	public static void main(String[] args) {

//		SpringApplication.run(OdbSpacebearsAdapterReleaseApplication.class, args);

		System.out.println(args);
		if (args != null && args.length > 0 && "generate-manifest".equals(args[0])) {
			GenerateManifest m = new GenerateManifest();
			m.generateManifest();
		}

	}
}
