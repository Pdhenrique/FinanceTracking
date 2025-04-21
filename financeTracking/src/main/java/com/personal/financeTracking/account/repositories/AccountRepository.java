package com.personal.financeTracking.account.repositories;

import com.personal.financeTracking.account.entities.Account;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface AccountRepository extends JpaRepository<Account, String> {
    Optional<Account> findByUserId(UUID userId);
}
