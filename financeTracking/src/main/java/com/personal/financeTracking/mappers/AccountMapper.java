package com.personal.financeTracking.mappers;

import com.personal.financeTracking.account.dto.SimpleAccountDto;
import com.personal.financeTracking.account.entities.Account;
import org.springframework.stereotype.Component;

@Component
public class AccountMapper {

    public SimpleAccountDto toSimpleAccountDTO(Account account) {
        if (account == null) return null;

        return SimpleAccountDto.builder()
                .id(account.getId())
                .bankName(account.getBankName())
                .accountType(account.getAccountType())
                .balance(account.getBalance())
                .build();
    }
}
