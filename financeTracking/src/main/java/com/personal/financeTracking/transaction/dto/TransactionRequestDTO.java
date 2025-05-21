package com.personal.financeTracking.transaction.dto;

import com.personal.financeTracking.enums.TransactionType;
import jakarta.validation.constraints.NotBlank;
import lombok.Data;

import java.time.LocalDateTime;

@Data
public class TransactionRequestDTO {

    @NotBlank(message = "id cannot be empty")
    private String id;
    private LocalDateTime createdAt;
    private String sender;
    private TransactionType transactionType;
    private String category;
}
