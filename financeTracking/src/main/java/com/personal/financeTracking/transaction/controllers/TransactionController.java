package com.personal.financeTracking.transaction.controllers;

import com.personal.financeTracking.transaction.dto.TransactionResponseDTO;
import com.personal.financeTracking.transaction.services.TransactionService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping(value = "/transactions")
public class TransactionController {

    private final TransactionService service;

    public TransactionController(TransactionService service){
        this.service = service;
    }


    @GetMapping
    public List<TransactionResponseDTO> findAll(){
        return service.findAll;
    }

}
