package com.personal.financeTracking.user.dto;

import com.personal.financeTracking.account.dto.SimpleAccountDto;

import lombok.Builder;
import lombok.Data;

import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

@Data
@Builder
public class UserResponseDTO {
    private UUID id;
    private String name;
    private String email;
    private String cpf;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    private List<SimpleAccountDto> accounts;
    private boolean active;
}
