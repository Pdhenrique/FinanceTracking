package com.personal.financeTracking.user.repositories;

import org.springframework.data.jpa.repository.JpaRepository;

import com.personal.financeTracking.user.entities.User;

import java.util.Optional;

public interface UserRepository extends JpaRepository<User, String> {
    Optional<User> findByEmailOrCpf(String email, String cpf);
}
