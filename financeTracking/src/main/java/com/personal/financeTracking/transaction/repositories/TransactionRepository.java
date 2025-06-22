package com.personal.financeTracking.transaction.repositories;

import com.personal.financeTracking.transaction.entities.Transaction;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface TransactionRepository extends JpaRepository<Transaction, UUID> {

}
