package com.personal.financeTracking.account.dto;

import com.personal.financeTracking.enums.AccountType;
import lombok.Builder;
import lombok.Data;

import java.util.UUID;

@Data
@Builder
public class SimpleAccountDto {
    private UUID id;
    private String bankName;
    private AccountType accountType;
    private Double balance;
}
