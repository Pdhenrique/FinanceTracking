package com.personal.financeTracking.user.services;

import com.personal.financeTracking.mappers.AccountMapper;
import com.personal.financeTracking.user.dto.UserRequestDTO;
import com.personal.financeTracking.user.dto.UserResponseDTO;
import com.personal.financeTracking.user.entities.User;
import com.personal.financeTracking.user.repositories.UserRepository;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.UUID;
import java.util.stream.Collectors;
import java.util.List;

@Service
public class UserService {

    private final UserRepository repository;
    private final PasswordEncoder encoder;

    @Autowired
    private AccountMapper accountMapper;

    @Transactional
    public void deactivate(UUID id) {
        User user = repository.findById(id)
                .orElseThrow(() -> new RuntimeException("User not found"));
        user.setActive(false);
        user.getAccounts().forEach(acc -> acc.setActive(false));
    }

    public UserService(UserRepository repository){
        this.repository = repository;
        this.encoder = new BCryptPasswordEncoder();
    }

    public List<UserResponseDTO> findAll(){
        return repository.findAll().stream()
                .map(this::toResponseDTO).collect(Collectors.toList());
    }

    public UserResponseDTO findById(UUID id) {
        User user = repository.findById(id).orElseThrow(() -> new RuntimeException("User not found with ID: " + id));
        return toResponseDTO(user);
    }

    public UserResponseDTO create(UserRequestDTO dto) {
        if (repository.findByEmailOrCpf(dto.getEmail(), dto.getCpf()).isPresent()) {
            throw new IllegalArgumentException("User with email or CPF already exists");
        }

        User user = new User();
        user.setName(dto.getName());
        user.setEmail(dto.getEmail());
        user.setCpf(dto.getCpf());
        user.setPassword(encoder.encode(dto.getPassword()));

        user = repository.save(user);
        return toResponseDTO(user);
    }

    private UserResponseDTO toResponseDTO(User user) {
        return UserResponseDTO.builder()
                .id(user.getId())
                .name(user.getName())
                .email(user.getEmail())
                .cpf(user.getCpf())
                .createdAt(user.getCreatedAt())
                .updatedAt(user.getUpdatedAt())
                .accounts(user.getAccounts().stream()
                        .map(accountMapper::toSimpleAccountDTO)
                        .collect(Collectors.toList()))
                .build();
    }
}
