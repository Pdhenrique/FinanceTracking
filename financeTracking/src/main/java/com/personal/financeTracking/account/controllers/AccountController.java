package com.personal.financeTracking.account.controllers;


import com.personal.financeTracking.account.dto.AccountRequestDTO;
import com.personal.financeTracking.account.dto.AccountResponseDTO;
import com.personal.financeTracking.account.services.AccountService;
import com.personal.financeTracking.user.controllers.UserController;
import jakarta.validation.Valid;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value = "/accounts")
public class AccountController {

    private final AccountService service;

    public AccountController(AccountService service){
        this.service = service;
    }

    @GetMapping
    public List<AccountResponseDTO> findAll(){
        return service.findAll();
    }

    @PostMapping
    public AccountResponseDTO create(@Valid @RequestBody AccountRequestDTO accountRequestDTO){
        return service.create(accountRequestDTO);
    }

}
