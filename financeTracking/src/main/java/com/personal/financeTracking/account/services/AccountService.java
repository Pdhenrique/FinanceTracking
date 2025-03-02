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

    private AccountResponseDTO toResponseDTO(Account account){
     AccountResponseDTO dto = new AccountResponseDTO();

     dto.setId(account.getId());
     dto.setAccountNumber(account.getAccountNumber());
     dto.setAccountType(account.getAccountType());
     dto.setBalance(account.getBalance());
     dto.setBankName(account.getBankName());
     dto.setCreatedAt(account.getCreatedAt());
     dto.setUpdatedAt(account.getCreatedAt());
     dto.setUser(account.getUser());
     dto.setUser(account.getUser());
     return dto;
    }
}
