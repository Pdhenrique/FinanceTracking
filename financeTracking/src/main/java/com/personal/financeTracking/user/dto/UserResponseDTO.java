package com.personal.financeTracking.user.dto;

import com.personal.financeTracking.account.entities.Account;
import lombok.Data;

import java.time.LocalDateTime;
import java.util.List;

@Data
public class UserResponseDTO {
    private String id;
    private String name;
    private String email;
    private String cpf;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    private List<Account> accounts;
}
