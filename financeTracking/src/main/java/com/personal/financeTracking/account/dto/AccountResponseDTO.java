package com.personal.financeTracking.account.dto;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class AccountResponseDTO {
    private String id;
    private String bankName;
    private String accountNumer;
    private String accountType;
    private String userId;
    private Double balance;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
}
