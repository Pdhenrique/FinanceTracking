package com.personal.financeTracking.account.services;

import com.personal.financeTracking.account.dto.AccountRequestDTO;
import com.personal.financeTracking.account.dto.AccountResponseDTO;
import com.personal.financeTracking.account.entities.Account;
import com.personal.financeTracking.account.repositories.AccountRepository;
import com.personal.financeTracking.exceptions.ResourceNotFoundException;
import com.personal.financeTracking.user.entities.User;
import com.personal.financeTracking.user.repositories.UserRepository;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;
import java.util.stream.Collectors;

@Service
public class AccountService {

    @Autowired
    private AccountRepository repository;

    @Autowired
    private UserRepository userRepository;

    public List<AccountResponseDTO> findAll(){

        return repository.findAll().stream().map(this::toResponseDTO).collect(Collectors.toList());
    }

    public AccountResponseDTO findById(UUID id) {
        Account account = repository.findById(id.toString())
                .orElseThrow(() -> new ResourceNotFoundException("Account not found with id: " + id));
        return toResponseDTO(account);
    }

    public AccountResponseDTO create(@Valid AccountRequestDTO dto) {

      Account account = new Account();
      account.setAccountNumber(dto.getAccountNumber());
      account.setAccountType(dto.getAccountType());
      account.setBankName(dto.getBankName());

      User user = userRepository.findById(dto.getUserId())
              .orElseThrow(() -> new ResourceNotFoundException("User not found with id: " + dto.getUserId()));
      account.setUser(user);

      account.setBalance(0.0);

      account = repository.save(account);
      return toResponseDTO(account);
    }

    private AccountResponseDTO toResponseDTO(Account account) {
        return AccountResponseDTO.builder()
                .id(account.getId())
                .bankName(account.getBankName())
                .accountNumber(account.getAccountNumber())
                .accountType(account.getAccountType())
                .balance(account.getBalance())
                .active(account.isActive())
                .createdAt(account.getCreatedAt())
                .updatedAt(account.getUpdatedAt())
                .userId(account.getUser().getId())
                .build();
    }
}
