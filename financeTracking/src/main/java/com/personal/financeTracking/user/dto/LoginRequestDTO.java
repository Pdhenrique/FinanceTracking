package com.personal.financeTracking.user.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.Data;

@Data
public class LoginRequestDTO {
  @NotBlank(message = "Email cannot be empty")
  @Email(message = "Invalid Email")
  private String email;

  @NotBlank(message = "Password cannot be empty")
  @Size(min = 6, message = "Password must contain at least 6 characters")
  private String password;
}
