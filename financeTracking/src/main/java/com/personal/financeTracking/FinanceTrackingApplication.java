package com.personal.financeTracking;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

@SpringBootApplication
@EnableJpaAuditing
public class FinanceTrackingApplication {

	public static void main(String[] args) {
		SpringApplication.run(FinanceTrackingApplication.class, args);
	}

}
