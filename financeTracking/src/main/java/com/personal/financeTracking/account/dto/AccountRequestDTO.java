package com.personal.financeTracking.account.dto;

import com.personal.financeTracking.user.entities.User;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.validation.constraints.NotBlank;
import lombok.Data;

@Data
public class AccountRequestDTO {

    @NotBlank(message = "Bank name cannot be empty")
    private String bankName;

    @NotBlank(message = "Account number cannot be empty")
    private String accountNumber;

    @NotBlank(message = "Account type cannot be empty")
    @Enumerated(EnumType.STRING)
    private String accountType;

    @NotBlank(message = "User id cannot be null")
    private String userId;
}
