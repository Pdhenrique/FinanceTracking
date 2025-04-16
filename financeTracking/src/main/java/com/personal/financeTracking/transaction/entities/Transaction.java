package com.personal.financeTracking.transaction.entities;

import com.personal.financeTracking.account.entities.Account;
import jakarta.persistence.*;

import java.time.LocalDateTime;

public class Transaction {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @Column(nullable = false)
    private double value;

    @Column(nullable = false)
    private String sender;

    @Column(nullable = false)
    private String addressee;

    @Column(nullable = false)
    private String description;

    @Column(nullable = false)
    private LocalDateTime transactionDate;

    @Column(nullable = false)
    private String transactionType;


    @Column(nullable = false)
    @OneToMany(mappedBy = "transaction", orphanRemoval = true)
    private Account account;
}
