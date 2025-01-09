package com.personal.financeTracking.account.entities;

import jakarta.persistence.*;

@Entity
@Table(name = "tb_account")
public class Account {


    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String account_id;

    private String
}
