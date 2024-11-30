package com.personal.financeTracking.user.repositories;

import org.springframework.data.jpa.repository.JpaRepository;

import com.personal.financeTracking.user.entities.User;

public interface UserRepository extends JpaRepository<User, Long> {

}
