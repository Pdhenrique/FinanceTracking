package com.personal.financeTracking.user.services;

import com.personal.financeTracking.user.dto.UserRequestDTO;
import com.personal.financeTracking.user.dto.UserResponseDTO;
import com.personal.financeTracking.user.entities.User;
import com.personal.financeTracking.user.repositories.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Optional;
import java.util.stream.Collectors;
import java.util.List;

@Service
public class UserService {

    @Autowired
    private final UserRepository repository;
    private final BCryptPasswordEncoder encoder;

    public UserService(UserRepository repository){
        this.repository = repository;
        this.encoder = new BCryptPasswordEncoder();
    }

    public List<UserResponseDTO> findAll(){
        return repository.findAll().stream()
                .map(this::toResponseDTO).collect(Collectors.toList());
    }

    public UserResponseDTO findById(String id) {
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
        UserResponseDTO dto = new UserResponseDTO();
        dto.setId(user.getId());
        dto.setName(user.getName());
        dto.setEmail(user.getEmail());
        dto.setCpf(user.getCpf());
        dto.setCreatedAt(user.getCreatedAt());
        dto.setUpdatedAt(user.getUpdatedAt());
        dto.setAccounts(user.getAccounts());
        return dto;
    }
}
