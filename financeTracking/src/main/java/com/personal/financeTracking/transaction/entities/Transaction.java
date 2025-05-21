package com.personal.financeTracking.transaction.entities;

import com.fasterxml.jackson.annotation.JsonBackReference;
import com.personal.financeTracking.account.entities.Account;
import com.personal.financeTracking.enums.TransactionType;
import jakarta.persistence.*;
import jakarta.validation.constraints.Positive;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

@Entity
@Table(name = "tb_transactions", indexes = @Index(columnList = "account_id, occurredAt"))
@EntityListeners(AuditingEntityListener.class)
@Data
@NoArgsConstructor
public class Transaction {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Positive
    @Column(nullable = false)
    private BigDecimal amount;

    @Column(nullable = false)
    private String sender;

    @Column(nullable = false)
    private String addressee;

    @Column(nullable = false)
    private String description;

    @Column(nullable = false)
    private LocalDateTime occurredAt;

    @Column(nullable = false)
    private TransactionType transactionType;

    @Column(nullable = false)
    private String category; //Update after create category


    @ManyToOne(fetch = FetchType.LAZY, optional = false)
    @JoinColumn(name = "account_id")
    @JsonBackReference
    private Account account;
}
