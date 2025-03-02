package com.personal.financeTracking.user.controllers;
import com.personal.financeTracking.user.dto.UserRequestDTO;
import com.personal.financeTracking.user.dto.UserResponseDTO;
import com.personal.financeTracking.user.services.UserService;
import org.springframework.web.bind.annotation.*;
import jakarta.validation.Valid;
import java.util.List;

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
    public UserResponseDTO findById(@PathVariable String id) {
        return service.findById(id);
    }

    @PostMapping
    public UserResponseDTO create(@Valid @RequestBody UserRequestDTO userRequestDTO){
        return service.create(userRequestDTO);
    }
}
