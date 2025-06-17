package com.personal.financeTracking.transaction.services;

import com.personal.financeTracking.transaction.dto.TransactionResponseDTO;
import com.personal.financeTracking.transaction.repositories.TransactionRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;

import java.util.List;

@Service
public class TransactionService {

    @Autowired
    private TransactionRepository transactionRepository;

    @Autowired
    private WebClient c6WebClient;

    public List<TransactionResponseDTO> findAll(String consentId, String accessToken) {
        return c6WebClient.get()
                .uri("/open-banking/accounts/v1/", consentId)
                .headers(h -> h.setBearerAuth(accessToken))
                .retrieve()
                .bodyToMono()
    }

}
