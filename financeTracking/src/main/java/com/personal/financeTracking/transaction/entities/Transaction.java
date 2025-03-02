package com.personal.financeTracking.transaction.entities;

import jakarta.persistence.*;
import jakarta.validation.constraints.NotNull;
import lombok.Data;

import java.util.UUID;

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

}
