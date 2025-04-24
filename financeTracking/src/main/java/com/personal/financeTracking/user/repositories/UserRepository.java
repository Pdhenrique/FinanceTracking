package com.personal.financeTracking.user.repositories;

import org.springframework.data.jpa.repository.JpaRepository;

import com.personal.financeTracking.user.entities.User;

import java.util.Optional;
import java.util.UUID;

public interface UserRepository extends JpaRepository<User, UUID> {
    Optional<User> findByEmailOrCpf(String email, String cpf);
    Optional<User> findByEmail(String email);
}
