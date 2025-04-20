package com.personal.financeTracking.account.dto;

import com.personal.financeTracking.enums.AccountType;
import com.personal.financeTracking.user.entities.User;
import lombok.Data;

import java.time.LocalDateTime;

@Data
public class AccountResponseDTO {
    private String id;
    private String bankName;
    private String accountNumber;
    private AccountType accountType;
    private User user;
    private Double balance;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    private boolean active;
}
