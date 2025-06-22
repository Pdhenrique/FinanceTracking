package com.personal.financeTracking.enums;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

public enum AccountType {
    CONTA_CORRENTE,
    POUPANCA,
    CARTAO_CREDITO,
    INVESTIMENTO;

    @JsonCreator
    public static AccountType fromString(String value) {
        try {
            return AccountType.valueOf(value.toUpperCase());
        } catch (IllegalArgumentException e) {
            throw new RuntimeException("Invalid account type: " + value);
        }
    }

    @JsonValue
    public String toValue() {
        return this.name();
    }
}
