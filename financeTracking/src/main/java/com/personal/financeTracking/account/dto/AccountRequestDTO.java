package com.personal.financeTracking.account.dto;

import com.personal.financeTracking.enums.AccountType;
import com.personal.financeTracking.user.entities.User;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.Data;

@Data
public class AccountRequestDTO {

    @NotBlank(message = "Bank name cannot be empty")
    private String bankName;

    @NotBlank(message = "Account number cannot be empty")
    private String accountNumber;

    @NotNull(message = "Account type cannot be null")
    private AccountType accountType;

    @NotBlank(message = "User id cannot be null")
    private String userId;
}
