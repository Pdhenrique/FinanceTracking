package com.personal.financeTracking.user.controllers;

import com.personal.financeTracking.user.repositories.UserRepository;
import com.personal.financeTracking.user.entities.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value = "/users")
public class UserController {

    @Autowired
    private UserRepository repository;

    @GetMapping
    public List<User> findAll() {
       return repository.findAll();
    }

    @GetMapping(value = "/{id}")
    public User findById(@PathVariable Long id) {
        return repository.findById(id).get();

    }

    @PostMapping(value = "/create")
    public User create(@RequestBody User user) {
        return repository.save(user);

    }
}
