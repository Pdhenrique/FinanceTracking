package com.personal.financeTracking.user.dto;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class UserResponseDTO {
    private String id;
    private String name;
    private String email;
    private String cpf;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
}
