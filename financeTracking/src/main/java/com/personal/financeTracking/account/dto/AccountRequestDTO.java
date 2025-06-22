package com.personal.financeTracking.account.dto;

import com.personal.financeTracking.enums.AccountType;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.Data;

import java.util.UUID;

@Data
public class AccountRequestDTO {

    @NotBlank(message = "Bank name cannot be empty")
    private String bankName;

    @NotBlank(message = "Account number cannot be empty")
    private String accountNumber;

    @NotNull(message = "Account type cannot be null")
    private AccountType accountType;

    @NotNull(message = "User id cannot be null")
    private UUID userId;
}
