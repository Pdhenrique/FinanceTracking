package com.personal.financeTracking.user.controllers;
import com.personal.financeTracking.user.dto.LoginRequestDTO;
import com.personal.financeTracking.user.dto.UserRequestDTO;
import com.personal.financeTracking.user.dto.UserResponseDTO;
import com.personal.financeTracking.user.services.UserService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import jakarta.validation.Valid;
import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping(value = "/users")
public class UserController {

    private final UserService service;

    public UserController(UserService service){
        this.service = service;
    }

    @GetMapping
    public List<UserResponseDTO> findAll(){
        return service.findAll();
    }

    @GetMapping(value = "/{id}")
    public UserResponseDTO findById(@PathVariable UUID id) {
        return service.findById(id);
    }

    @PostMapping
    public UserResponseDTO create(@Valid @RequestBody UserRequestDTO userRequestDTO){
        return service.create(userRequestDTO);
    }

    @PostMapping(value = "/login")
    public ResponseEntity<UserResponseDTO> login(@RequestBody @Valid LoginRequestDTO dto){
        UserResponseDTO user = service.login(dto);
        return ResponseEntity.ok(user);
    }
}
