package com.personal.financeTracking.account.dto;

import com.personal.financeTracking.enums.AccountType;
import com.personal.financeTracking.user.entities.User;
import lombok.Builder;
import lombok.Data;

import java.time.LocalDateTime;
import java.util.UUID;

@Data
@Builder
public class AccountResponseDTO {
    private UUID id;
    private String bankName;
    private String accountNumber;
    private AccountType accountType;
    private Double balance;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    private boolean active;
    private UUID userId;
}
