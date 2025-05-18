package com.personal.financeTracking.user.repositories;

import org.springframework.data.jpa.repository.JpaRepository;

import com.personal.financeTracking.user.entities.User;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

@Repository
public interface UserRepository extends JpaRepository<User, UUID> {
    Optional<User> findByEmailOrCpf(String email, String cpf);
}
